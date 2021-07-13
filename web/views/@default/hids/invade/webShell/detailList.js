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
        window.location.href = "/hids/invade/webshell"
    };
    this.onDetail = function (item) {
        teaweb.popup(Tea.url(".detail?macCode=" + this.macCode +
            '&riskId=' + item.riskId +
            '&isProcess=' + (item.state != 0)
        ), {
            height: "30em",
        })
    };
    this.onTrust = function (item) {
        teaweb.confirm("确定信任该事件吗？", function () {
            this.$post("/hids/invade/webshell").params({
                Opt: "add_trust",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.onIsolate = function (item) {
        teaweb.confirm("确定隔离该事件吗？", function () {
            this.$post("/hids/invade/webshell").params({
                Opt: "isolate",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.onRevert = function (item) {
        teaweb.confirm("确定恢复该事件吗？", function () {
            this.$post("/hids/invade/webshell").params({
                Opt: "revert",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.onDelete = function (item) {
        teaweb.confirm("确定删除该事件吗？", function () {
            this.$post("/hids/invade/webshell").params({
                Opt: "delete",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.onCancelTrust = function (item) {
        teaweb.confirm("确定取消信任该事件吗？", function () {
            this.$post("/hids/invade/webshell").params({
                Opt: "cancel_trust",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.getStateName = function (state) {
        if (state == 0){
            return "未处理"
        }else if(state == 1){
            return "隔离"
        }else if(state == 2){
            return "信任"
        }else if(state == 3){
            return "删除"
        }else if(state == 101){
            return "隔离中"
        }else if(state == 2){
            return "信任中"
        }else if(state == 3){
            return "删除中"
        }else if(state == -1){
            return "隔离失败"
        }else if(state == -2){
            return "信任失败"
        }else if(state == -3){
            return "删除失败"
        }else{
            return "未知"+state
        }
    }
})