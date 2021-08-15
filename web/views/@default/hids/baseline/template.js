Tea.context(function () {

    this.onStartCheck = function () {
        this.$post(".check").params({
            MacCode: [this.macCode],
            serverIp: this.serverIp,
            templateId: this.sSelectValue,
        }).success(function (){
            teaweb.closePopup()
            parent.location.reload()
        }).error(function (){
            teaweb.warn("失败：该主机agent已暂停服务，命令无法执行！")
        })
    }

    //添加/删除元素
    this.onSetSelectValue = function (index) {
        if(this.sSelectValue != index)
            this.sSelectValue = index
        else{
            this.sSelectValue = 0
        }
    }

});
  