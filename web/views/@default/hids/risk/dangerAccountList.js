Tea.context(function () {
    this.nTableState = 1;

    this.onChangeTableState = function (state) {
        if (this.nTableState != state) {
            this.nTableState = state;
        }
    };
    this.getDangerAccountTypeName = function (tp) {
        if ('04011')
            return "操作系统配置缺陷";
        else if ('0301')
            return "可疑高权限账号";
        else if ('0302')
            return " WEB 空密码账号";
        else
            return "未知";
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

    this.onGoBack = function () {
        window.location.href = "/hids/risk/dangerAccount";
    };

    this.onDetail = function (item) {
        teaweb.popup(Tea.url(".dangerAccountDetail?macCode=" + this.macCode +
            '&riskId=' + item.riskId +
            '&state=' + this.nTableState
        ), {
            height: "350px",
            width: "800px",
        })
    };
    this.onIgnore = function (item) {
        teaweb.confirm("确定忽略该风险账号吗？", function () {
            this.$post(".dangerAccount").params({
                Opt: "ignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.cancelOnIgnore = function (item) {
        teaweb.confirm("确定取消忽略该风险账号吗？", function () {
            this.$post(".dangerAccount").params({
                Opt: "cancel_ignore",
                MacCode: this.macCode,
                ItemIds: [item.itemId],
                RiskIds: [item.riskId],
            }).refresh()
        })
    };
    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }
})
;