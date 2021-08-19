Tea.context(function () {
    // this.keyword = ""
    //directionState 方向 1 进 2 出
    this.getDirection=function (state) { 
        switch(state){
            case "in":
                return "进"
            case "out":
                return "出"
            default:
                return "未知"
        }
    }

    this.getAddr = function (ip,port) {
        if(ip != null){
            return ip+":"+port
        }
        return ""
    }

    //
    this.getNumber = function (num) {
        if(num < 1024){
            return num
        }
        return Math.floor(num /1024) + " K"
    }

    this.onSearch = function () {
        window.location.href = '/nfw/conversation?nodeId=' + this.selectNode+"&keyword="+this.keyword
    }
    //获取当前选中的节点
    this.GetSelectNode = function (event) {
        this.selectNode = event.target.value; //获取option对应的value值
        // localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/nfw/conversation?nodeId=' + node+"&keyword="+this.keyword

    }
})