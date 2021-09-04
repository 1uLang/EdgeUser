Tea.context(function () {


    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        // teaweb.datepicker("day-to-picker")
    })

    this.onSearch = function () {
        window.location.href = '/maltrail/index?nodeId=' + this.selectNode+"&dayFrom="+this.dayFrom
    }
    //获取当前选中的节点
    this.GetSelectNode = function (event) {
        this.selectNode = event.target.value; //获取option对应的value值
        // localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/maltrail/index?nodeId=' + node+"&dayFrom="+this.dayFrom

    }
    this.onDate = function () {
        let start = document.getElementById("day-from-picker").value
        // console.log(start)
        this.dayFrom = start
    }
})