Tea.context(function () {
    this.curIndex = -1

    this.onGoBack = function () {
        window.location.href = "/hids/risk";
    }

    this.$delay(function () {
        console.log(this.errorMessage)
        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage)
        }
    })
    this.getCount = function (item) {
        return item.lowRiskCount + item.middleRiskCount + item.highRiskCount + item.criticalCount
    }
    this.onOpenDetail = function (item) {
        window.location.href = "/hids/risk/dangerAccountList?ip=" + item.serverIp + '&macCode=' + item.macCode +
            "&os=" + item.os.osType + "&lastUpdateTime=" + item.os.lastUpdateTime
    }
    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }

    this.enters = function (index) {
        this.curIndex = index;
    }

    this.leaver = function (index) {
        this.curIndex = -1;
    }

})
