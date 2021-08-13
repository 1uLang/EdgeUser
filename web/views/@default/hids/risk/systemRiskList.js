Tea.context(function () {
    this.nTableState = 1;

    this.onChangeTableState = function (state) {
        this.nTableState = state
    };


    this.getStatusName = function (status) {
        switch (status) {
            case 0:
                return "待处理";
            case 2:
                return "已隔离";
            case 3:
                return "已信任";
            case 4:
                return "已忽略";
            case 5:
                return "已修复";
            case 6:
                return "不适用";
            case 7:
                return "已关闭";
            case 11:
                return "等待隔离";
            case 21:
                return "等待信任";
            case 31:
                return "等待删除";
            case 51:
                return "等待修复";
            case 101:
                return "隔离中";
            case 201:
                return "信任中";
            case 301:
                return "删除中";
            case 501:
                return "修复中";
            case -1:
                return "隔离失败";
            case -2:
                return "信任失败";
            case -3:
                return "删除失败";
            case -5:
                return "修复失败";
            case -10:
                return "取消隔离失败";
            case -20:
                return "取消信任失败";
            default:
               return  "未知";
        }
    };

    this.getDangerName = function (status) {
        switch (status) {
            case 1:
                return "低危";
            case 2:
                return "中危";
            case 3:
                return "高危";
            case 4:
                return "危急";
            default:
                return "未评级";
        }
    };

    this.onGoBack = function () {
        window.location.href = "/hids/risk/systemRisk";
    };

    this.onDetail = function (item) {
        teaweb.popup(Tea.url(".riskDetail?macCode=" + this.macCode +
            '&riskId=' + item.riskId +
            '&state=' + this.nTableState +
            '&detailName=' + item.detailName +
            '&os=' + this.os
        ),{
            height: "50em",
            width: "800px",
        })
    };
    this.onFix = function (item) {
        teaweb.confirm("确定修复该系统漏洞吗？", function () {
            this.$post(".systemRisk").params({
                Opt: "repair",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.onIgnore = function (item) {
        teaweb.confirm("确定忽略该系统漏洞吗？", function () {
            this.$post(".systemRisk").params({
                Opt: "ignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.cancelOnIgnore = function (item) {
        teaweb.confirm("确定取消忽略该系统漏洞吗？", function () {
            this.$post(".systemRisk").params({
                Opt: "cancel_ignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
});