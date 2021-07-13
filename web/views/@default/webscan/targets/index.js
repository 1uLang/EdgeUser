Tea.context(function () {
    this.checkValues = []; //选中的ID

    this.scanValues = []  //需要扫描的ID

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
                    this.targets = resp.data.targets
                }
            })
        }
    }
    this.onCreateLoopTimeOut = function () {
        this.onReleaseTimeOut()
        this.checkTimer = createTimer(this.onCallBack, {timeout: 30000});
        this.checkTimer.start();
    }
    this.onReleaseTimeOut = function () {
        if (this.checkTimer) {
            this.checkTimer.stop()
            this.checkTimer = null
        }
    }
    this.checkScans = function () {
        for (item of this.targets) {
            if (item.last_scan_session_status === "processing") {
                return true
            }
        }
        return false
    }

    this.onScan = function () {
        if (this.scanValues.length > 0) {
            let that = this
            let target_ids = JSON.parse(JSON.stringify(this.scanValues))
            teaweb.confirm("确定要扫描这个目标吗？", function () {
                that.$post("/webscan/scans/create")
                    .params({
                        TargetIds: target_ids
                    }).success(function () {
                    that.$delay(function () {   //延迟 500毫秒跳转
                        window.location = "/webscan/scans"
                    }, 500)
                })
            })
        }
    };
    this.onDelete = function () {
        if (this.checkValues.length > 0) {
            let that = this
            let target_ids = JSON.parse(JSON.stringify(this.checkValues))
            teaweb.confirm("确定要删除这个目标吗？", function () {
                that.$post(".delete")
                    .params({
                        TargetIds: target_ids
                    })
                    .refresh()
            })
        }
    };

    this.hanleOpen = function () {
        teaweb.popup(Tea.url(".create"), {
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload();
                });
            },
        });
    };
    this.save = function () {
        const tempAddress = console.log(document.getElementById("key").value);
        const tempDesc = console.log(document.getElementById("value").value);
    };

    this.handleOnCheck = function () {
        const scanBtn = document.getElementById("scan-btn");
        scanBtn.style.backgroundColor = "#14539A";
        scanBtn.style.cursor = "pointer";

        const delBtn = document.getElementById("del-btn");
        delBtn.style.backgroundColor = "#D9001B";
        delBtn.style.cursor = "pointer";
    };

    /* 数据模版
      var scanData = [
          {'id':1,'ip':'192.168.0.1','disc':'test','type':'web','loophole':{'red':1,'yellow':2,'blue':3,'green':4},'lastStatus':'已完成','lastTime':'2021-06-11 09:29:23'},
          {'id':2,'ip':'192.168.1.1','disc':'test','type':'web','loophole':{'red':1,'yellow':3,'blue':4,'green':2},'lastStatus':'未完成','lastTime':'2021-06-12 09:29:23'},
          {'id':3,'ip':'192.168.5.1','disc':'test','type':'web','loophole':{'red':3,'yellow':2,'blue':1,'green':4},'lastStatus':'进行中','lastTime':'2021-06-13 09:29:23'},
          {'id':4,'ip':'192.168.6.1','disc':'test','type':'web','loophole':{'red':6,'yellow':5,'blue':1,'green':4},'lastStatus':'进行中','lastTime':'2021-06-14 09:29:23'},
      ]
      */

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


    this.getItemInfo = function (id) {
        if (this.targets && this.targets.length > 0) {
            for (var i = 0; i < this.targets.length; i++) {
                if (this.targets[i].target_id == id) {
                    return this.targets[i]
                }
            }
        }
        return null
    }

    this.checkCanScan = function () {
        this.scanValues = []
        for (id of this.checkValues) {
            let itemInfo = this.getItemInfo(id)
            if (itemInfo && (itemInfo.last_scan_session_status == "aborted" ||
                itemInfo.last_scan_session_status == "completed" ||
                itemInfo.last_scan_session_status == null)) {
                this.scanValues.push(id)
            }
        }
    }

    this.updateBtnStatus = function () {
        this.checkCanScan()
        const scanBtn = document.getElementById("scan-btn");
        const delBtn = document.getElementById("del-btn");
        if (this.checkValues.length > 0) {

            delBtn.style.backgroundColor = "#D9001B";
            delBtn.style.cursor = "pointer";
        } else {

            delBtn.style.backgroundColor = "#AAAAAA";
            delBtn.style.cursor = null;
        }

        if (this.scanValues.length > 0) {
            scanBtn.style.backgroundColor = "#14539A";
            scanBtn.style.cursor = "pointer";
        } else {
            scanBtn.style.backgroundColor = "#AAAAAA";
            scanBtn.style.cursor = null;
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
        if (status) {
            switch (status) {
                case "aborted":
                    return "已停止";
                case "completed":
                    return "已完成";
                case "processing":
                    return "扫描中";
                default:
                    return status;
            }
        } else {
            return "未扫描"
        }

    };
});
