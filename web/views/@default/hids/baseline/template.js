Tea.context(function () {
    this.curIndex = -1

    this.onChangeCheckState = function (state) {
        window.location = "/hids/baseline?State="+state+'&ResultState='+this.ResultState
     }

    this.onChangeResultState = function(state){
        window.location = "/hids/baseline?State="+this.State+'&ResultState='+state
    }
    this.parseServerLocalIp = function (ip){
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }

    this.getStateName = function (status) {
        switch(status){
            case 0:
                return "未检查"
            case 1:
                return "检查中"
            case 2:
                return "已完成"
            case 3:
                return "检查失败"
            default:
                return "未知"
        }
     }

     this.enters = function (index) {
        this.curIndex = index;
      }
    
      this.leaver = function (index) {
        this.curIndex = -1;
      }
    
    this.getProgressPer = function (curValue,maxValue) { 
        if(curValue && maxValue && maxValue >= curValue){
            return parseInt(curValue/maxValue * 100)+"%"
        }
        return "0%"
     }

    //合规基线
     this.onOpenCheck = function (item) {
        //打开合规基线弹窗
         teaweb.popup(Tea.url(".template?macCode="+item.macCode), {
             height: "30em",
         })
      }

    this.onStartCheck = function () {
        this.$post(".check").params({
            MacCode: [this.macCode],
            serverIp: this.serverIp,
            templateId: this.sSelectValue,
        }).success(function (){
            teaweb.closePopup()
            window.location.reload()
        }).error(function (){
            teaweb.warn("失败：该主机agent已暂停服务，命令无法执行！")
        })
    }

    this.onOpenDetail = function (item) {
        window.location = "/hids/baseline/detail?macCode="+item.macCode+'&pageSize='+item.totalItemCount
    }

    //添加/删除元素
    this.onSetSelectValue = function (index) {
        if(this.sSelectValue != index)
            this.sSelectValue = index
        else{
            this.sSelectValue = 0
        }
    }
    this.getShowSelectImage = function (id) {

        if (this.sSelectValue == id) {
          return "/images/select_select.png";
        }
        return "/images/select_box.png";
      }
});
  