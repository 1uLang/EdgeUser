Tea.context(function () {
    let that = this
    if (this.featureIsOn) {
        this.accessLogs.forEach(function (accessLog) {
            if (typeof (that.regions[accessLog.remoteAddr]) == "string") {
                accessLog.region = that.regions[accessLog.remoteAddr]
            } else {
                accessLog.region = ""
            }
        })
    }
})