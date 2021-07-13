Tea.context(function () {
    this.nTableState = 1;

    this.onChangeTableState = function (state) {
        if (this.nTableState != state) {
            this.nTableState = state;
        }
    };

    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }
    this.onGoBack = function () {
        window.location = "/hids/invade/abnormalAccount"
    };
    this.onDetail = function (item) {
        teaweb.popup(Tea.url(".detail?macCode=" + this.macCode +
            '&riskId=' + item.riskId +
            '&isProcess=' + (item.state != 0)
        ), {
            height: "30em",
        })
    };
    this.onClose = function (item) {
        teaweb.confirm("确定关闭该事件吗？", function () {
            this.$post("/hids/invade/abnormalAccount").params({
                Opt: "close",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.onCancelClose = function (item) {
        teaweb.confirm("确定取消关闭该事件吗？", function () {
            this.$post("/hids/invade/abnormalAccount").params({
                Opt: "cancel_close",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.getStateName = function (state) {
        if (state == 0){
            return "未处理"
        }else{
            return "已关闭"
        }
    }
})