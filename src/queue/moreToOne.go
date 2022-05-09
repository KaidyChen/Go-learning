package main

import (
	"container/list"
	"context"
	"fmt"
	"sort"
	"sync"
)

/*
多队列单消费者模式，优先级队列
*/

//对队列进行一次封装
type PriorityQueue struct {
	mu sync.Mutex
	noticeChan chan struct{}
	capacity int
	queues []*JobQueue
	PriorityIdx map[int]int  //该map的key是优先级，value代表的是queues切片的索引
	size int
}


//定义队列结构体，这里我们用Golang中的List数据结果来实现，List数据结构是一个双向链表，包含了将元素放到链表尾部、将头部元素弹出的操作，符合队列先进先出的特性
type JobQueue struct {
	priority int //代表该队列是哪种优先级的队列
	jobList *list.List //List是golang库的双向队列实现，每个元素都是一个job
}

//队列的Push操作,入队
func (priorityQueue *PriorityQueue) PushJob(job Job) {
	priorityQueue.mu.Lock()
	defer  priorityQueue.mu.Unlock()

	//先根据job的优先级找到要入队的目标队列
	var idx int
	var ok bool
	//从优先级切片索引的map中查找该优先级的队列是否存在
	if idx, ok := priorityQueue.PriorityIdx[job.Priority()]; !ok {
		//如果不存在该优先级的队列，则需要初始化一个队列并返回该队列在切片中的索引位置
		idx = priorityQueue.addPriorityQueue(job.Priority())
	}
	//根据获取到的切片索引idx找到具体的队列
	queue := priorityQueue.queues[idx]
	//将job推送到队列的队尾
	queue.jobList.PushBack(job)

	//队列job个数+1
	priorityQueue.size++

	//如果队列job个数超过队列的最大容量，则从优先级最低的队列中移除工作单元
	if priorityQueue.size > priorityQueue.capacity {
		priorityQueue.RemoveLeastPriorityJob()
	} else {
		//通知新进来一个job
		priorityQueue.noticeChan <- struct{}{}
	}
}

func (priorityQueue *PriorityQueue) addPriorityQueue(priority int) int {
	n := len(priorityQueue.queues)
	//通过二分查找找到priority应插入的切片索引
	pos := sort.Search(n, func(i int) bool {
		return priority < priorityQueue.queues[i].priority
	})

	tail := make([]*JobQueue, n -pos)
	copy(tail, priorityQueue.queues[pos:])

	//初始化一个全新的优先级队列并将该元素放到切片的pos位置中
	priorityQueue.queues = append(priorityQueue.queues[0:pos], newJobQueue(priority))
	//将高于priority优先级的元素也拼接到切片后面
	priorityQueue.queues = append(priorityQueue.queues, tail...)
	return pos
}

//弹出队列的第一个元素
func (queue *PriorityQueue) PopJob() Job {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	//判断队列中有没有任务
	if queue.jobList.Len() == 0 {
		fmt.Println("队列中没有任务...")
		return nil
	}

	elements := queue.jobList.Front() //获取队列中的第一个元素
	return queue.jobList.Remove(elements).(Job) //将元素从队列中移除并返回
}

//等待通知操作
func (queue *JobQueue) WaitJob() <-chan struct{} {
	return queue.noticeChan
}

/*
Job接口：定义了所有Job要实现的方法
BaseJob类（结构体）：定义了具体Job的基类。因为具体Job类中的有共同的属性和方法。所以抽象出一个基类，避免重复实现。但该基类对Execute方法没有实现，因为不同的工作单元有具体的执行逻辑
SquareJob和AreaJob类（结构体）：是我们要具体实现的业务工作Job。主要是实现Execute的具体执行逻辑。根据业务的需要定义自己的工作Job和对应的Execute方法即可
*/

type Job interface {
	Execute() error
	WaitDone()
	Done()
	Priority() int
}


/*
优先级的队列，实质上就是根据工作单元Job的优先级属性，将其放到对应的优先级队列中，以便worker可以根据优先级进行消费。
我们要在Job结构体中增加一个Priority属性。因为该属性是所有Job都共有的，因此定义在BaseJob上更合适
*/
type BaseJob struct {
	Err error
	DoneChan chan struct{} //当作业完成时，或者作业被取消时，通知调用者
	Ctx context.Context
	cancelFunc context.CancelFunc
	priority int //工作单元的优先级
}

/**
 * 作业执行完毕，关闭DoneChan，所有监听DoneChan的接收者都能收到关闭的信号
 */
func (job *BaseJob) Done()  {
	close(job.DoneChan)
}

/**
 * 等待job执行完成
 */
func (job *BaseJob) WaitDone() {
	select {
	case <-job.DoneChan:
		return
	}
}

//具体任务结构体
type SquareJob struct {
	*BaseJob
	x int
	priority int
}

func (s *SquareJob) Execute() error {
	result := s.x * s.x
	fmt.Println("the result is ", result)
	return nil
}

func (s *SquareJob) Priority() int {
	return s.priority
}

//定义消费者结构体, Worker主要功能是通过监听队列里的noticeChan是否有通知，如果有通知的话从队列里获取到要处理的元素job，然后执行job的Execute方法
type WorkerManager struct {
	queue *JobQueue
	closeChan chan struct{}
}

func (m *WorkerManager) StartWork() error {
	fmt.Println("start to work")
	for {
		select {
		case <-m.closeChan:
			return nil
		case <-m.queue.noticeChan:
			job := m.queue.PopJob()
			m.ConsumeJob(job)
		}
	}
	return nil
}

func (m *WorkerManager) ConsumeJob(job Job) {
	defer func() {
		job.Done()
	}()

	job.Execute()
}

func main() {
	//初始化一个队列
	queue := &JobQueue{
		jobList: list.New(),
		noticeChan: make(chan struct{}, 10),
	}

	//初始化一个消费者worker
	workerManager := WorkerManager{
		queue: queue,
		closeChan: make(chan struct{}, 1),
	}

	//worker开始监听队列
	go workerManager.StartWork()

	//构造具体任务job
	job := &SquareJob{
		BaseJob: &BaseJob{
			DoneChan: make(chan struct{}, 1),
		},
		x: 5,
	}

	//压入队列中
	queue.PushJob(job)

	//等待job执行完成
	job.WaitDone()

	fmt.Println("Finsh the job")
}