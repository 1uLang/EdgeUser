Tea.context(function () {

    this.$delay(function () {

        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")

    })

    this.onChangeCheckTime = function (index) {

        if (this.index != index) {
            this.index = index
        }
        this.onSearch()
    }

    this.onSearch = function () {
        window.location = "/fortcloud/command?index=" + this.index +
            "&asset=" + this.asset + "&input=" + this.input + "&username=" +
            this.username + "&riskLevel=" + this.riskLevel + "&dayFrom=" +
            this.dayFrom + "&dayTo=" + this.dayTo
    }
    this.onRefresh = function () {
        window.location = "/fortcloud/command"
    }
    this.onTimeChange = function () {

        let startTime = document.getElementById("day-from-picker").value
        let endTime = document.getElementById("day-to-picker").value

        if(this.dayFrom != startTime || this.dayTo != endTime) {
            this.dayFrom = startTime
            this.dayTo = endTime
            this.index = -1
            this.onSearch()
        }

    }

    this.onChangeTimeFormat = function (timestamp) {
        var date = new Date(timestamp * 1000);
        return date.format("yyyy-MM-dd hh:mm:ss");
    }

    this.onOpenItemDetail = function (id,table) {
        if(id && table){
            for(var index = 0;index<table.length;index++){
                if(table[index]._id == id){
                    table[index].bOpen = !table[index].bOpen
                    break
                }
            }
        }
    }
})