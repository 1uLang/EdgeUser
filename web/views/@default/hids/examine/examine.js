Tea.context(function () {

    this.webSearchKey = ""  //网页后门
    this.searchPath = ""    //病毒木马

    this.bTimeOutTip = false
    this.bShowScanPath = false

    let that = this

    this.sBottomSelectItem = [
        {id: "13", value: "反弹shell"},
        {id: "14", value: "异常账号"},
        {id: "15", value: "系统命令篡改"},
        {id: "16", value: "异常进程"},
        {id: "17", value: "日志异常删除"},
    ]

    this.sSelectCheckValue = ["01", "02", "03", "04", "13", "14", "15"]

    this.pCheckDetailData = [
        {
            checkName: "漏洞风险检查项：",
            checkValue: [
                {id: "01", value: "系统漏洞"},
                {id: "02", value: "弱口令"},
                {id: "03", value: "风险账号"},
                {id: "04", value: "配置缺陷"},
            ]
        },
        {
            checkName: "入侵威胁检查项：",
            checkValue: this.sBottomSelectItem
        }
    ]

    //检测是否包含元素
    this.checkSelectValue = function (index, selectValue) {
        if (selectValue) {
            for (var i = 0; i < selectValue.length; i++) {
                if (selectValue[i] == index) {
                    return true;
                }
            }
        }

        return false;
    }

    this.onSelectCheckValue = function (index) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(index, this.sSelectCheckValue);
        }
        if (bValue) {
            this.sSelectCheckValue = this.sSelectCheckValue.filter((itemIndex) => {
                return itemIndex != index;
            });
        } else {
            this.sSelectCheckValue.push(index);
        }
    }

    this.getShowSelectValueImage = function (id) {
        let bValue = false;
        if (that.checkSelectValue) {
            bValue = that.checkSelectValue(id, that.sSelectCheckValue);
        }
        if (bValue) {
            return "/images/select_select.png";
        }
        return "/images/select_box.png";
    }

    this.onCheckSelectItem = function (id) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(id, this.sSelectCheckValue);
        }
        return bValue
    }

    this.onStartCheck = function () {
        if (this.sSelectCheckValue.length == 0) {
            teaweb.warn("请选择体检项目")
            return
        }
        teaweb.confirm("确定立即体检吗？", function () {
            this.$post(".scans").params({
                Opt: 'now',
                serverIp: this.serverIp,
                VirusPath: this.searchPath,
                WebShellPath: this.webSearchKey,
                MacCode: [this.macCode],
                ScanItems: this.sSelectCheckValue.join(","),
            }).success(function () {
                teaweb.closePopup()
                parent.location.reload()
            }).error(function () {
                teaweb.warn("失败：该主机agent已暂停服务，命令无法执行！")
            })
        })
    }
    //检测是否显示扫描路径的输入框和提示框
    this.onCheckSelectValue = function () {

        var selextBox = document.getElementsByName("customScan")
        if (selextBox) {
            for (var item of selextBox) {
                if (item.checked) {
                    if (item.value == 2) {
                        this.bTimeOutTip = true
                        this.bShowScanPath = false
                    } else if (item.value == 3) {
                        this.bTimeOutTip = false
                        this.bShowScanPath = true
                    } else {
                        this.bTimeOutTip = false
                        this.bShowScanPath = false
                    }
                }
            }
        }
    }

})