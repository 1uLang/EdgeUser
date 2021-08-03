Tea.context(function () {

    this.page = ""
    this.pageState = 1

    this.editUserName = "zhangxl1"
    this.editPassword = ""
    this.editPasswordConfirm = ""
    this.editFullName = ""
    this.editPhone = ""
    this.editEmail = ""
    this.editRemark=""
    this.editEnabled = 0

    this.selectAuthData = [] //开启的权限ID


    this.onChangeShowState = function (state){
        if(this.pageState!=state){
            this.pageState = state
        }
    }
	this.createUser = function () {
		teaweb.popup(Tea.url(".create"), {
			height: "680px",
            width:"650px",
			callback: function () {
				teaweb.success("保存成功", function () {
					teaweb.reload()
				})
			}
		})
	}

    this.onOpenDetail = function(userId){
        this.onChangeShowState(2)
    }

	this.onDelete = function (userId) {
		let that = this
		teaweb.confirm("确定要删除这个用户吗？", function () {
			// that.$post(".delete")
			// 	.params({
			// 		userId: userId
			// 	})
			// 	.refresh()
		})
	}

    this.getStatus = function (status) {
        switch (status) {
            case 1:
                return "已启用"
            case 0:
                return "已停用"
            default:
                return "已停用"
        }
    }

    this.onListenEditCheckBox=function(){
        let enabled = document.getElementById("editCheckBox").checked
        if(enabled){
            this.editEnabled = 1
        }else{
            this.editEnabled = 0
        }
    }

    this.onSaveEdit = function(){
        // this.editUserName 
        // this.editPassword 
        // this.editPasswordConfirm
        // this.editFullName
        // this.editPhone 
        // this.editEmail 
        // this.editRemark
        // this.editEnabled
    }

    this.onSaveAuth = function(){
        this.selectAuthData = []
        let authCheckBoxList = document.getElementsByName("authCheckBox")
        for(var index=0;index<authCheckBoxList.length;index++){
            if(authCheckBoxList[index].checked){
                this.selectAuthData.push(authCheckBoxList[index].value)
            }
        }
        console.log(this.selectAuthData)
    }


    this.users=[
        {id:1,username:"zhangxl1",fullname:"zhangxl1",mobile:"15505565626",createdTime:"2021-07-31 15:50:58",status:1}
    ]

    this.authData = [
        {id:1,checkState:1,name:"记录访问日志0",subName:"用户可以开启服务的访问日志0"},
        {id:2,checkState:0,name:"记录访问日志1",subName:"用户可以开启服务的访问日志1"},
        {id:3,checkState:0,name:"记录访问日志2",subName:"用户可以开启服务的访问日志2"},
        {id:4,checkState:1,name:"记录访问日志3",subName:"用户可以开启服务的访问日志3"},
        {id:5,checkState:1,name:"记录访问日志4",subName:"用户可以开启服务的访问日志4"},
        {id:6,checkState:0,name:"记录访问日志5",subName:"用户可以开启服务的访问日志5"},
        {id:7,checkState:1,name:"记录访问日志6",subName:"用户可以开启服务的访问日志6"},
    ]
})