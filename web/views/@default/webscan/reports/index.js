Tea.context(function () {
        this.checkValues = []; //选中的ID


        this.checkTimer = null

        this.$delay(function () {
            //开启监听
            let that = this
            that.onCreateLoopTimeOut()
            window.addEventListener('beforeunload', function () {
                that.onReleaseTimeOut()
            })
        })
        this.onCallBack = function () {
            if (this.checkScans()) {
                this.$post(".").success(resp => {
                    if (resp.code === 200) {
                        this.reports = resp.data.reports
                    }
                })
            }
        }
        this.onCreateLoopTimeOut = function () {
            this.onReleaseTimeOut()
            this.checkTimer = createTimer(this.onCallBack, {timeout: 10000});
            this.checkTimer.start();
        }
        this.onReleaseTimeOut = function () {
            if (this.checkTimer) {
                this.checkTimer.stop()
                this.checkTimer = null
            }
        }
        this.checkScans = function () {
            for (item of this.reports) {
                if (item.status === "processing" ||
                    item.status === "queued") {
                    return true
                }
            }
            return false
        }

        this.clickCheckbox = function () {
            var checkDomArr = document.querySelectorAll(
                ".multi-table tbody input[type=checkbox]:checked"
            );
            this.checkValues = [];
            for (var i = 0, len = checkDomArr.length; i < len; i++) {
                this.checkValues.push(checkDomArr[i].value);
            }
            var allCheckDomArr = document.querySelectorAll(
                ".multi-table tbody input[type=checkbox]"
            );
            var allCheckbox = document.getElementById("js-all-checkbox");
            for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
                if (!allCheckDomArr[i].checked) {
                    if (allCheckbox.checked) allCheckbox.checked = false;
                    break;
                } else if (i === len - 1) {
                    document.getElementById("js-all-checkbox").checked = true;
                    break;
                }
            }
            this.updateBtnStatus();
        };
        this.checkAll = function () {
            var curClickBox = document.getElementById("js-all-checkbox")
            var allCheckDomArr = document.querySelectorAll(
                ".multi-table tbody input[type=checkbox]"
            );
            if (!curClickBox.checked) {
                // 点击的时候, 状态已经修改, 所以没选中的时候状态时true
                this.checkValues = [];
                for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
                    var checkStatus = allCheckDomArr[i].checked;
                    if (checkStatus) allCheckDomArr[i].checked = false;
                }
            } else {
                this.checkValues = [];
                for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
                    var checkStatus = allCheckDomArr[i].checked;
                    if (!checkStatus) allCheckDomArr[i].checked = true;
                    this.checkValues.push(allCheckDomArr[i].value);
                }
            }
            this.updateBtnStatus();
        };
        this.updateBtnStatus = function () {
            const delBtn = document.getElementById("del-btn");
            if (this.checkValues.length > 0) {
                delBtn.style.backgroundColor = "#D9001B";
                delBtn.style.cursor = "pointer";
            } else {
                delBtn.style.backgroundColor = "#AAAAAA";
                delBtn.style.cursor = null;
            }
        };
        this.parseDesc = function (desc) {
            return desc.split(";", 1)[0]
        }
        //btn删除
        this.onDelete = function () {
            if (this.checkValues.length > 0) {
                let that = this
                let report_ids = JSON.parse(JSON.stringify(this.checkValues))
                teaweb.confirm("确定要删除这个报表吗？", function () {
                    that.$post(".delete")
                        .params({
                            ReportIds: report_ids
                        })
                        .refresh()
                })
            }
        };

        this.onChangeTimeFormat = function (time) {
            var resultTime = "";
            if (time) {
                var tempTime = time.substring(0, time.indexOf("."));
                resultTime = tempTime.replace("T", " ");
            }
            return resultTime;
        };

        this.onChangeStatusFormat = function (status) {
            var resultStatus = status;
            if (status) {
                switch (status) {
                    case "aborted":
                        return "已中止";
                    case "queued":
                        return "生成中";
                    case "completed":
                        return "已完成";
                    case "progressing":
                        return "正在进行";
                }
            }
            return resultStatus;
        };

        this.downloadFile = function (path,pdf) {

            window.open("reports/download?path=" + path+"&PDF="+pdf)
        }
        //主机漏洞扫描报表
        this.exportFile = function (item, format) {

            window.open("reports/export?id=" + item.scan_id + "&historyId=" + item.history_id + "&format=" + format)
        }
    }
);
