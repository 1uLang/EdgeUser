Tea.context(function () {
    this.nTableState = 1;

    this.onChangeTableState = function (state) {
        if (this.nTableState != state) {
            this.nTableState = state;
        }
    };

    this.getWeakTypeName = function (tp) {
        switch (tp) {
            case '0201':
                return "操作系统弱口令";
            case '0202':
                return "数据库弱口令";
            case '0203':
                return "应用软件弱口令";
            default:
                return "未知";
        }
    }


    this.getStatusName = function (status) {
        switch (status) {
            case 1:
                return "已修复";
            case 2:
                return "待修复";
            case 3:
                return "修复中";
            case 4:
                return "修复失败";
            case 5:
                return "已忽略";
            default:
                return "未知";
        }
    };

    this.onGoBack = function () {
        window.location.href = "/hids/risk/weak";
    };

    this.parseServerLocalIp = function (ip){
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }
    this.getDangerName = function (status) {
        switch (status) {
            case 1:
                return "低危";
            case 2:
                return "中危";
            case 3:
                return "高危";
            case 4:
                return "危机";
            default:
                return "未评级";
        }
    };
    this.onDetail = function (item) {
        teaweb.popup(Tea.url(".weakDetail?macCode=" + this.macCode +
            '&riskId=' + item.riskId +
            '&state=' + this.nTableState
        ),{
            height: "50em",
            width: "800px",
        })
    };
    this.onIgnore = function (item) {
        teaweb.confirm("确定忽略该弱口令吗？", function () {
            this.$post(".weak").params({
                Opt: "ignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.cancelOnIgnore = function (item) {
        teaweb.confirm("确定取消忽略该弱口令吗？", function () {
            this.$post(".weak").params({
                Opt: "cancel_ignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };

});
  