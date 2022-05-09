package main

import (
	"container/list"
	"context"
	"fmt"
	"sync"
)

/*
单队列单消费者模式
*/

//定义队列结构体，这里我们用Golang中的List数据结果来实现，List数据结构是一个双向链表，包含了将元素放到链表尾部、将头部元素弹出的操作，符合队列先进先出的特性
type JobQueue struct {
	mu sync.Mutex //队列的操作需要并发安全
	jobList *list.List //List是golang库的双向队列实现，每个元素都是一个job
	noticeChan chan struct{} //入队一个job就往该channel中放入一个消息，以供消费者消费,注意channel中放入的消息不是任务本身，只是消息通知
}

//队列的Push操作,入队
func (queue *JobQueue) PushJob(job Job) {
	queue.jobList.PushBack(job) //将job加到队列尾部
	queue.noticeChan <- struct{}{} //发送通知给消费者有任务待消费
}

//弹出队列的第一个元素
func (queue *JobQueue) PopJob() Job {
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
}

type BaseJob struct {
	Err error
	DoneChan chan struct{} //当作业完成时，或者作业被取消时，通知调用者
	Ctx context.Context
	cancelFunc context.CancelFunc
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
}

func (s *SquareJob) Execute() error {
	result := s.x * s.x
	fmt.Println("the result is ", result)
	return nil
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