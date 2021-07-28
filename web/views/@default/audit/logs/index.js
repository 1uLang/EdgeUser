Tea.context(function () {
    this.pageState = 1
    this.objName = "数据库"

    this.timeSelectIndex = "5min"
    this.dayFrom = ""
    this.dayTo = ""

    this.sqlIndex = -1
    this.clientIp = ""
    this.sqlName = ""
    this.dangerLevelIndex = -1
    this.searchKey = ""

    this.bShowSqlName = false
    this.sSelectSqlName = []  //选中 的审计ID

    this.bShowDangerLevel = false
    this.sSelectDangerLevel = []  //风险等级

    //list 数据
    this.dbLogList = []
    this.hostLogList = []
    this.appLogList = []

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

    this.onChangeCheckTime = function (index) {
        if (this.timeSelectIndex != index) {
            this.timeSelectIndex = index
        }
    }

    this.onTimeChange = function () {

    }

    this.onChangeState = function (id) {
        if (this.pageState != id) {
            this.pageState = id
        }
        switch (this.pageState) {
            case 1:
                this.objName = "数据库"
                break;
            case 2:
                this.objName = "主机"
                break;
            case 3:
                this.objName = "应用"
                break;
        }

        this.sSelectSqlName = []
    }

    this.getDangerLevel = function (status) {
        // switch (status) {
        //     case 1:
        //         return "高"
        //     case 0:
        //         return "--"
        //     default:
        //         return "--"
        // }
        if (status) {
            return "高"
        } else {
            return "--"
        }
    }

    this.getDangerType = function (status) {
        // switch (status) {
        //     case 1:
        //         return "SQL注入"
        //     case 0:
        //         return "--"
        //     default:
        //         return "--"
        // }
        if (status) {
            return "sql注入"
        } else {
            return "--"
        }
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.getLogType = function (status) {
        switch (status) {
            case 1:
                return "安全策略"
            case 2:
                return "文件修改"
            case 3:
                return "安全审计"
            case 4:
                return "其他"
            case 0:
                return "系统配置"
            default:
                return "系统配置"
        }
    }

    this.onCloseDropMenu = function () {
        if (this.bShowSqlName) {
            this.bShowSqlName = false
        }
        if (this.bShowDangerLevel) {
            this.bShowDangerLevel = false
        }
    }

    this.onShowSqlName = function () {
        this.bShowSqlName = !this.bShowSqlName
    }

    this.onShowDangerLevel = function () {
        this.bShowDangerLevel = !this.bShowDangerLevel
    }

    this.onCheckSelectValue = function (id, table) {
        if (id && table) {
            for (var index = 0; index < table.length; index++) {
                if (table[index] == id) {
                    return true
                }
            }
        }
        return false
    }

    this.getItemInfo = function (id, table) {
        if (id && table) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].audit_id == id) {
                    return table[index].name
                }
            }
        }
        return ""
    }
    this.onAddSqlNameValue = function (id) {
        if (this.onCheckSelectValue(id, this.sSelectSqlName)) {
            this.sSelectSqlName = this.sSelectSqlName.filter((itemIndex) => {
                return itemIndex != id;
            });
        } else {
            this.sSelectSqlName.push(id);
        }
    }

    this.onAddDangerLevelValue = function (id) {
        if (this.onCheckSelectValue(id, this.sSelectDangerLevel)) {
            this.sSelectDangerLevel = this.sSelectDangerLevel.filter((itemIndex) => {
                return itemIndex != id;
            });
        } else {
            this.sSelectDangerLevel.push(id);
        }
    }

    this.getShowSelectValueImage = function (id, table) {
        let bValue = this.onCheckSelectValue(id, table);

        if (bValue) {
            return "/images/select_select.png";
        }
        return "/images/select_box.png";
    }

    //搜索
    this.onSearch = function (exp) {
        this.$post(".").params({
            timeType: this.timeSelectIndex,
            startTime: this.dayFrom,
            endTime: this.dayTo,
            auditId: this.sSelectSqlName,
            cIp: this.clientIp,
            user: this.sqlName,
            risk: this.sSelectDangerLevel,
            message: this.searchKey,
            logType: this.pageState,
            export: exp,//导出
            pageNum: 1,
            pageSize: 100,
        }).success(resp => {
            if (resp.code === 200 && resp.data.list) {
                if(exp == "false"){
                    switch (this.pageState) {
                        case 1:
                            this.dbLogList = resp.data.list
                            break;
                        case 2:
                            this.hostLogList = resp.data.list
                            break;
                        case 3:
                            this.appLogList = resp.data.list
                            break;
                        default:
                            this.list = []
                    }
                }else{
                    //文件下载路径
                    let filepath = resp.data.filepath

                }



                // this.level = resp.data.level
            }
        })
    }

    this.tableData1 = [
        {
            id: 1,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "192.11.11.1",
            value4: "root",
            value5: "mysql",
            value6: "SELECT * FROM `edgeLogins` WHERE `",
            value7: 1,
            value8: 0,
            value9: "2021-06-19T16:38:00.123"
        },
        {
            id: 2,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "192.11.11.1",
            value4: "root",
            value5: "mysql",
            value6: "SELECT * FROM `edgeLogins` WHERE `",
            value7: 0,
            value8: 1,
            value9: "2021-06-19T16:38:00.123"
        },
        {
            id: 3,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "192.11.11.1",
            value4: "root",
            value5: "mysql",
            value6: "SELECT * FROM `edgeLogins` WHERE `",
            value7: 1,
            value8: 1,
            value9: "2021-06-19T16:38:00.123"
        },
        {
            id: 4,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "192.11.11.1",
            value4: "root",
            value5: "mysql",
            value6: "SELECT * FROM `edgeLogins` WHERE `",
            value7: 0,
            value8: 0,
            value9: "2021-06-19T16:38:00.123"
        },
        {
            id: 5,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "192.11.11.1",
            value4: "root",
            value5: "mysql",
            value6: "SELECT * FROM `edgeLogins` WHERE `",
            value7: 0,
            value8: 0,
            value9: "2021-06-19T16:38:00.123"
        },
    ]

    this.tableData2 = [
        {
            id: 1,
            value1: "192.168.12.12",
            value2: "Windows",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0 auid=4294967295",
            value4: 0,
            value5: "2021-06-19T16:38:00.123"
        },
        {
            id: 2,
            value1: "192.168.12.12",
            value2: "Linux",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0 auid=4294967295",
            value4: 1,
            value5: "2021-06-19T16:38:00.123"
        },
        {
            id: 3,
            value1: "192.168.12.12",
            value2: "Windows",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0 auid=4294967295",
            value4: 2,
            value5: "2021-06-19T16:38:00.123"
        },
        {
            id: 4,
            value1: "192.168.12.12",
            value2: "Linux",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0 auid=4294967295",
            value4: 3,
            value5: "2021-06-19T16:38:00.123"
        },
    ]

    this.tableData3 = [
        {
            id: 1,
            value1: "192.168.12.12",
            value2: "Nginx",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0",
            value4: "GET",
            value5: "200",
            value6: "2021-06-19T16:38:00.123"
        },
        {
            id: 2,
            value1: "192.168.12.12",
            value2: "Nginx",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0",
            value4: "POST",
            value5: "400",
            value6: "2021-06-19T16:38:00.123"
        },
        {
            id: 3,
            value1: "192.168.12.12",
            value2: "Nginx",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0",
            value4: "DELETE",
            value5: "500",
            value6: "2021-06-19T16:38:00.123"
        },
        {
            id: 4,
            value1: "192.168.12.12",
            value2: "Nginx",
            value3: "type=CRYPTO_KEY_USER msg=audit(1624267796.088:1806632): pid=21873 uid=0",
            value4: "PUT",
            value5: "404",
            value6: "2021-06-19T16:38:00.123"
        },
    ]

    this.sqlNameData = [
        {id: 1, name: "115.169.23.236", value: "115.169.23.236"},
        {id: 2, name: "115.169.23.237", value: "115.169.23.236"},
        {id: 3, name: "115.169.23.238", value: "115.169.23.236"},
        {id: 4, name: "115.169.23.238", value: "115.169.23.236"},
    ]

    this.dangerLevelData = [
        // {id:1,name:"低",value:"低"},
        // {id:2,name:"中",value:"中"},
        {id: 3, name: "高", value: "高"},
        // {id:4,name:"严重",value:"严重"},
        // {id:5,name:"正常",value:"正常"},
    ]
})