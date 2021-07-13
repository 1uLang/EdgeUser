Tea.context(function () {
    this.address = ''
    this.severity = ''
    this.detailInfo = null
    this.bShowDetail = false

    this.onCloseDetail = function () {
        this.bShowDetail = false
    };

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };
    this.onChangeSeverityFormat = function (severity) {
        var resultSeverity = severity;

        switch (severity) {
            case 3:
                return '高危'
            case 2:
                return '中危'
            case 1:
                return '低危'
            default:
                return '信息'
        }
        return resultSeverity;
    };

    this.getDetailInfo = function (vul) {
        this.detailInfo = null
        this.$get("/webscan/vulnerabilities/details").params({
            VulId: vul.vuln_id
        }).success(resp => {
            if (resp.code === 200) {
                this.detailInfo = resp.data.data
                this.detailInfo.affects_url = "URL:           " + this.detailInfo.affects_url
                this.bShowDetail = true
            } else {
                this.bShowDetail = false
            }
        })
    }

});
