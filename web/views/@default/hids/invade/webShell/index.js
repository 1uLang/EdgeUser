Tea.context(function () {
    this.curIndex = -1
    this.onGoBack = function () {
        window.location = "/hids/invade"
    }
    this.onOpenDetail = function (item) {
        window.location = "/hids/invade/abnormalProcess/detailList?macCode="+item.macCode+"&ip="+item.serverIp
    }
    this.enters = function (index) {
        // this.curIndex = index;
    }
    this.$delay(function () {

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })

    this.leaver = function () {
        this.curIndex = -1;
    }

    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }
    this.tableTitle = [
        {name: "主机IP", width: "834px"},
        {name: "网页后门数", width: "200px"},
        {name: "详情", width: "90px"}
    ]

})