Tea.context(function () {
    this.curIndex = -1

    this.onGoBack = function () {
        window.location.href = "/hids/risk";
    }

    this.$delay(function () {

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })
    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }
    this.getCount = function (item) {
        return item.lowRiskCount + item.middleRiskCount + item.highRiskCount + item.criticalCount
    }
    this.onOpenDetail = function (item) {

        window.location.href = "/hids/risk/weakList?ip=" + item.serverIp + '&macCode=' + item.macCode +
            "&os=" + item.os.osType + "&lastUpdateTime=" + item.os.lastUpdateTime
    }

    this.enters = function (index) {
        this.curIndex = index;
    }

    this.leaver = function (index) {
        this.curIndex = -1;
    }
})