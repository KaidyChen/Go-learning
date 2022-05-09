$(function(){
	app.init()
	$(window).resize(function (){
		app.resizeIframe()
	})
})

var config={
	adminPath:"beego"
}

var app={
	init:function(){
		this.slideToggle();
		this.resizeIframe();
		this.confirmDelete();
		this.changeStatus();
		this.changeNum()
	},
	slideToggle: function () {
		$('.aside>li:nth-child(1) ul, .aside>li:nth-child(2) ul').hide()
		$('.aside h4').click(function(){
			$(this).siblings('ul').slideToggle();
		})
	},
	resizeIframe:function () {
		$("#rightMain").height($(window).height() - 80)
	},
	//删除提示
	confirmDelete:function() {
		$(".delete").click(function (){
			var flag=confirm("您确定要删除吗?")
			return flag
		})
	},
	//修改元素状态属性,实行页面局部刷新
	changeStatus:function () {
		$(".chStatus").click(function (){
			var id = $(this).attr("data-id");
			var table = $(this).attr("data-table");
			var field = $(this).attr("data-field");
			var el = $(this)
			$.get("/"+config.adminPath+"/main/changeStatus", {id:id, table:table, field:field}, function(response){
				if(response.success){
					if(el.attr("src").indexOf("yes") != -1){
						el.attr("src", "/static/admin/images/no.gif")
					}else{
						el.attr("src", "/static/admin/images/yes.gif")
					}
				}else{
					console.log(response)
				}
			})
		})
	},
	changeNum:function (){
		$(".chNum").click(function (){
			var id = $(this).attr("data-id");
			var table = $(this).attr("data-table");
			var field = $(this).attr("data-field");
			var num = $(this).attr("data-num");
			var spanElementor = $(this) //保存span这个dom节点

			var input=$("<input value='' style='width: 60px' />");
			$(this).html(input);
			$(input).trigger('focus').val(num);//让输入框获取焦点并设置值
			$(input).click(function (e){
				e.stopPropagation();
			})
			$(input).blur(function (e) {
				var inputNum = $(this).val();
				spanElementor.html(inputNum);
				//ajax异步请求修改数量
				$.get("/"+config.adminPath+"/main/changeNum",{id:id,table:table,field:field,num:inputNum},function(response){
					if(!response.success){
						console.log(response)
					}
				})
			})
		})
	}
}

