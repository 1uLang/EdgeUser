Tea.context(function () {
    // this.dayFrom = ""
    // this.dayTo = ""
    // this.keyword = ""
    // this.page = ""

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

    this.onSearch = function(){
        this.dayFrom = document.getElementById("day-from-picker").value
        this.dayTo = document.getElementById("day-to-picker").value
        console.log(this.dayFrom)
        console.log(this.dayTo)
        console.log(this.keyword)
        //req
    }

    this.showMore = function (log) {
        log.moreVisible = !log.moreVisible
    }


    this.exportExcel = function () {
        let that = this
        teaweb.confirm("确定要将当前列表导出到Excel吗？", function () {
            window.location = "/platform/logs/exportExcel?dayFrom=" + that.dayFrom
                + "&dayTo=" + that.dayTo + "&keyword=" + that.keyword
        })
    }

    // this.logs = [
    //     {id:1,level:"error",createdTime:"2021-08-03 16:21:06",userName:"user01",userId:"015455",ip:"192.168.1.1",region:"this is region",action:"this is action",description:"this is description"},
    // ]
})