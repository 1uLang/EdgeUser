Tea.context(function () {
    this.nTableState = 1;

    this.onChangeTableState = function (state) {
        if (this.nTableState != state) {
            this.nTableState = state;
        }
    };

    this.onGoBack = function () {
        window.location.href = "/hids/invade/abnormalProcess"
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
            this.$post("/hids/invade/abnormalProcess").params({
                Opt: "close",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.onCancelClose = function (item) {
        teaweb.confirm("确定取消关闭该事件吗？", function () {
            this.$post("/hids/invade/abnormalProcess").params({
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