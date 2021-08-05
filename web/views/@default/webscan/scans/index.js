Tea.context(function () {
    this.checkValues = []; //选中的ID
    this.stopValues = []; //选中的ID 停止扫描
    this.createValues = []; //选中的ID 生成报表

    this.checkTargetValues = []; //选中的targetID
    this.stopTargetValues = []; //选中的targetID
    this.createTargetValues = []; //选中的targetID

    this.nShowState = 1   //三个界面的状态控制 1 2 3
    this.vulnerabilities = []
    this.statistics = {}
    this.severity = ""
    this.scanSeverity = 0
    this.checkPer = "12%"
    this.Address = ""
    this.scans_vulns = []
    this.scanId = ""
    this.scanSessionId = ""
    this.scanAddr = ""
    this.bLoopholeDetail = false    //漏洞详情是否显示
    this.hostVulFlag = false //主机漏洞扫描
    this.bShowDetail = false

    this.showDetailItem = null

    this.onCloseDetail = function () {
        this.bShowDetail = false
    };

    this.$delay(function () {
        //开启监听
        let that = this
        that.onCreateLoopTimeOut()
        window.addEventListener('beforeunload', function () {
            that.onReleaseTimeOut()
        })
    })

    this.checkTimer = null

    this.onCallBack = function () {
        if (this.checkScans()) {
            this.$post(".").success(resp => {
                if (resp.code === 200) {
                    this.scans = resp.data.scans
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
        if (this.nShowState !== 1)
            return false

        for (item of this.scans) {
            if (item.current_session.status === "processing" || item.current_session.status === "running" ||
                item.current_session.status === "queued" ||
                item.current_session.status === "aborting") {
                return true
            }
        }
        return false
    }

    this.onStopScan = function (item) {

        let curValue = []
        if(item.owner){
            curValue = [item.target_id]
        }else{
            curValue = [item.scan_id]
        }

        let that = this
        let scan_ids = JSON.parse(JSON.stringify(curValue))
        teaweb.confirm("确定要停止这个扫描吗？", function () {
            that.$post(".stop")
                .params({
                    ScanIds: scan_ids
                }).success(function () {
                window.location.reload()
            })
        })
    };
    this.onCreateReport = function (item) {

        let curValue = [item.scan_id]
        let curTargetValue = [item.target_id]
        let that = this
        let scan_ids = JSON.parse(JSON.stringify(curValue))
        let tarId = JSON.parse(JSON.stringify(curTargetValue))
        teaweb.confirm("确定要生成这个扫描的报表吗？", function () {
            that.$post("/webscan/reports/create")
                .params({
                    Ids: scan_ids,
                    TarIds: tarId,
                }).success(function () {
                window.location.href = "/webscan/reports"
            })
        })
    };

    this.onDelete = function () {
        if (this.checkValues.length > 0) {
            let that = this
            let scan_ids = JSON.parse(JSON.stringify(this.checkValues))
            let ids = JSON.parse(JSON.stringify(this.checkTargetValues))
            teaweb.confirm("确定要删除这个扫描吗？", function () {
                that.$post(".delete")
                    .params({
                        ScanIds: scan_ids,
                        ids: ids,
                    }).success(function () {
                    window.location.reload()
                })
            })
        }
    };

    this.clickCheckbox = function () {
        var checkDomArr = document.querySelectorAll(
            ".multi-table tbody input[type=checkbox]:checked"
        );
        this.checkValues = [];
        this.checkTargetValues = [];
        for (var i = 0, len = checkDomArr.length; i < len; i++) {
            this.checkValues.push(checkDomArr[i].value);
            let tar = checkDomArr[i].getAttribute("data")
            this.checkTargetValues.push(tar);
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
            this.checkTargetValues = [];
            for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
                var checkStatus = allCheckDomArr[i].checked;
                if (checkStatus) allCheckDomArr[i].checked = false;
            }
        } else {
            this.checkValues = [];
            this.checkTargetValues = [];
            for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
                var checkStatus = allCheckDomArr[i].checked;
                if (!checkStatus) allCheckDomArr[i].checked = true;
                this.checkValues.push(allCheckDomArr[i].value);
                let tar = allCheckDomArr[i].getAttribute("data")
                this.checkTargetValues.push(tar);
            }
        }
        this.updateBtnStatus();
    };


    this.getItemInfo = function (id) {
        if (this.scans && this.scans.length > 0) {
            for (var i = 0; i < this.scans.length; i++) {
                if (this.scans[i].scan_id == id) {
                    return this.scans[i]
                }
            }
        }
        return null
    }
    this.updateBtnStatus = function () {

        // const stopBtn = document.getElementById("stop-btn");
        // const createBtn = document.getElementById("create-btn");
        const delBtn = document.getElementById("del-btn");

        this.stopValues = []
        this.stopTargetValues = []
        this.createValues = []
        this.createTargetValues = []

        for (let idx = 0; this.checkValues.length > idx; idx++) {
            let id = this.checkValues[idx]
            let tid = this.checkTargetValues[idx]
            let itemInfo = this.getItemInfo(id)
            if (itemInfo && (itemInfo.current_session.status == "processing" || itemInfo.current_session.status == "running")) {
                this.stopValues.push(id)
                this.stopTargetValues.push(tid)
            } else if (itemInfo && itemInfo.current_session.status == "completed") {
                this.createValues.push(id)
                this.createTargetValues.push(tid)
            }
        }
        if (this.checkValues.length > 0) {
            delBtn.style.backgroundColor = "#D9001B";
            delBtn.style.cursor = "pointer";
        } else {
            delBtn.style.backgroundColor = "#AAAAAA";
            delBtn.style.cursor = null;
        }
        // if (this.stopValues.length > 0) {
        //     stopBtn.style.backgroundColor = "#14539A";
        //     stopBtn.style.cursor = "pointer";
        // } else {
        //     stopBtn.style.backgroundColor = "#AAAAAA";
        //     stopBtn.style.cursor = null;
        // }
        // if (this.createValues.length > 0) {
        //     createBtn.style.backgroundColor = "#14539A";
        //     createBtn.style.cursor = "pointer";
        // } else {
        //     createBtn.style.backgroundColor = "#AAAAAA";
        //     createBtn.style.cursor = null;
        // }
    };

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.onChangeState = function (state) {
        if (state === 3) { //漏洞详情
            this.nShowState = 3;
            this.severity = ""
            this.vulnerabilities = this.scans_vulns
        } else if (state === 2) {
            if (this.showDetailItem)
                this.onShowDetail(this.showDetailItem)
        } else {
            window.location.reload()
        }
    }
    //主机漏洞列表
    this.onHostShowDetail = function (item) {
        this.showDetailItem = item
        this.scanId = item.scan_id
        this.scanSessionId = item.target_id
        this.scanAddr = item.target.address
        this.hostVulFlag = true
        this.$get(".vulnerabilities").params({
            scanId: this.scanId,
            scanSessionId: item.target_id,
        }).success(resp => {
            this.scans_vulns = []
            if (resp.code === 200) {
                this.scans_vulns = resp.data.data
                this.vulnerabilities = this.scans_vulns
            }
            this.nShowState = 3
        })
    }

    this.onShowDetail = function (item) {
        this.showDetailItem = item
        this.scanId = item.scan_id
        this.scanSessionId = item.current_session.scan_session_id
        //获取漏洞详情报表
        this.$get(".statistics").params({
            ScanId: item.scan_id,
            ScanSessionId: item.current_session.scan_session_id
        }).success(resp => {
            if (resp.code === 200) {
                this.scanSeverity = resp.data.severity
                // this.statistics.status = this.onChangeStatusFormat(resp.data.statistics.status)
                // this.statistics.severity_counts = item.current_session.severity_counts
                // this.statistics.event_level = resp.data.statistics.scanning_app.wvs.event_level
                // this.statistics.host = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].host
                // this.statistics.os = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.os
                // this.statistics.responsive = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.responsive ? '是' : '否'
                // this.statistics.server = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.server
                // this.statistics.technologies = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].target_info.technologies
                // this.statistics.request_count = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].web_scan_status.request_count
                // this.statistics.avg_response_time = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].web_scan_status.avg_response_time
                // this.statistics.locations = resp.data.statistics.scanning_app.wvs.hosts[item.target_id].web_scan_status.locations
                this.statistics.vulns = resp.data.statistics.scanning_app.wvs.main.vulns
                this.scans_vulns = this.statistics.vulns
                this.vulnerabilities = this.scans_vulns
                // this.statistics.duration = resp.data.statistics.scanning_app.wvs.main.duration
                // this.statistics.progress = resp.data.statistics.scanning_app.wvs.main.progress
                // this.statistics.messages = resp.data.statistics.scanning_app.wvs.main.messages
                // this.nShowState = 2
                this.nShowState = 3
            }
        }).done(function () {
            if (this.statistics.progress != 100) {
                this.$delay(function () {
                    if (this.nShowState === 2)
                        this.onShowDetail(item)
                }, 5000)
            }
        })
    }

    this.refreshProgress = function () {
        var maxCount = 100
        var tempCount = 25
        var curPer = Math.floor(maxCount / tempCount)
        checkPer = curPer + "%"
        document.getElementById("barContent").style.width = checkPer;
    }
    this.search = function () {
        let vulns = []
        if (this.severity == "")
            this.vulnerabilities = this.scans_vulns
        else {
            for (let idx = 0; idx < this.scans_vulns.length; idx++) {
                let vul = this.scans_vulns[idx]
                if (vul.severity == this.severity)
                    vulns.push(vul)
            }
            this.vulnerabilities = vulns
        }
    }

    this.onChangeStatusFormat = function (status) {
        var resultStatus = status;
        if (status) {
            switch (status) {
                case "aborted":
                    return "已中止";
                case "canceled":
                    return "已中止";
                case "completed":
                    return "已完成";
                case "processing":
                    return "正在进行";
                case "running":
                    return "正在进行";
                case "queued":
                    return "队列中";
                case "stopping":
                    return "停止中";
                case "aborting":
                    return "停止中";
            }
        }
        return resultStatus;
    };
    this.onChangeTimeFormat2 = function (time) {
        if (time) {
            let m = parseInt(time / 60)
            let s = time % 60
            return m + 'm' + s + 's'
        }
        return '0s'
    };

    this.curIndex = -1

    this.mouseEnter = function (index) {
        this.curIndex = index;
    }

    this.mouseLeave = function (index) {
        this.curIndex = -1;
    }

    this.onChangeSeverityFormat = function (severity) {
        var resultSeverity = severity;

        switch (severity) {
            case 3:
                return '高危'
            case 2:
                return '中危'
            case 1:
                return '低危'
            case 0:
                return '信息'
            default:
                return "危机"
        }
        return resultSeverity;
    };
    this.getDetailInfo = function (vul) {
        this.detailInfo = null
        if (!this.hostVulFlag) {
            this.$get(".vulnerabilities").params({
                vulId: vul.vuln_id,
                scanId: this.scanId,
                scanSessionId: this.scanSessionId,
            }).success(resp => {
                if (resp.code === 200) {
                    this.detailInfo = resp.data.data
                    this.detailInfo.affects_url = "URL:           " + this.detailInfo.affects_url
                    this.bShowDetail = true
                } else {
                    this.bShowDetail = false
                }
            })
        } else {
            this.$get("/webscan/vulnerabilities/details").params({
                vulId: vul.plugin_id,
                scanId: this.scanId,
                scanSessionId: this.scanSessionId,
            }).success(resp => {
                if (resp.code === 200) {
                    this.detailInfo = resp.data.data
                    this.detailInfo.affects_url = "URL:           " + this.scanAddr
                    this.bShowDetail = true
                } else {
                    this.bShowDetail = false
                }
            })
        }
    }
})
;
