Tea.context(function () {

    //这边数据是从服务器获取的 本地只做展示
    // this.name = ""
    // this.email = ""
    // this.format = 1
    // this.assets_type = 1
    // this.assets_id = 0
    // this.cycle=1
    // this.cycle_day = 1
    // this.send_time = "00:00"

    this.weekData = [
        {id: 1, name: "周一"}, {id: 2, name: "周二"}, {id: 3, name: "周三"}, {id: 4, name: "周四"}, {id: 5, name: "周五"}, {
            id: 6,
            name: "周六"
        }, {id: 7, name: "周日"},
    ]
    this.monthData = [
        {id: 1, name: "1号"}, {id: 2, name: "2号"}, {id: 3, name: "3号"}, {id: 4, name: "4号"}, {id: 5, name: "5号"}, {
            id: 6,
            name: "6号"
        }, {id: 7, name: "7号"},
        {id: 8, name: "8号"}, {id: 9, name: "9号"}, {id: 10, name: "10号"}, {id: 11, name: "11号"}, {
            id: 12,
            name: "12号"
        }, {id: 13, name: "13号"}, {id: 14, name: "14号"},
        {id: 15, name: "15号"}, {id: 16, name: "16号"}, {id: 17, name: "17号"}, {id: 18, name: "18号"}, {
            id: 19,
            name: "19号"
        }, {id: 20, name: "20号"}, {id: 21, name: "21号"},
        {id: 22, name: "22号"}, {id: 23, name: "23号"}, {id: 24, name: "24号"}, {id: 25, name: "25号"}, {
            id: 26,
            name: "26号"
        }, {id: 27, name: "27号"}, {id: 28, name: "28号"},
        {id: 29, name: "29号"}, {id: 30, name: "30号"}, {id: 31, name: "31号"}
    ]
    this.timeData = [
        {id: 1, name: "00:00"}, {id: 2, name: "01:00"}, {id: 3, name: "02:00"}, {id: 4, name: "03:00"}, {
            id: 5,
            name: "04:00"
        }, {id: 6, name: "05:00"}, {id: 7, name: "06:00"},
        {id: 8, name: "07:00"}, {id: 9, name: "08:00"}, {id: 10, name: "09:00"}, {id: 11, name: "10:00"}, {
            id: 12,
            name: "11:00"
        }, {id: 13, name: "12:00"}, {id: 14, name: "13:00"},
        {id: 15, name: "14:00"}, {id: 16, name: "15:00"}, {id: 17, name: "16:00"}, {id: 18, name: "17:00"}, {
            id: 19,
            name: "18:00"
        }, {id: 20, name: "19:00"}, {id: 21, name: "20:00"},
        {id: 22, name: "21:00"}, {id: 23, name: "22:00"}, {id: 24, name: "23:00"}
    ]

    this.onSelectTableType = function () {
        this.format = this.format === 1 ? 2 : 1
    }

    this.onSave = function () {
        this.$post("/audit/report/createPopup")
            .params({
                name: this.name,
                email: this.email,
                format: this.format,
                assetsType: this.assets_type,
                assetsId: this.assets_id,
                cycle: this.cycle,
                cycleDay: this.cycle_day,
                sendTime: this.send_time,
                id: this.id,
            })
            .success(function () {
                 this.success("保存成功", function () {
                     window.location = "/audit/report"
                 })
            })
    }
})