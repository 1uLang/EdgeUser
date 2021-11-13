Tea.context(function () {

    this.getTime = function (time) {
        var d = new Date(time);
        return d.toLocaleDateString() + " " + d.toLocaleTimeString()
    }

    //获取当前选中的节点
    this.GetSelectNode = function (event) {
        this.selectNode = event.target.value; //获取option对应的value值
        localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/nfw/virus?nodeId=' + node
    }
})