Tea.context(function () {
    this.timeSelectIndex = 0
    this.dayFrom = ""
    this.dayTo = ""

    this.userRes = -1
    this.userName = ""
    this.dangerLevel = -1
    this.inputCommand = ""

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
    })

    this.onChangeCheckTime = function (index) { 
        if(this.timeSelectIndex!=index){
            this.timeSelectIndex = index
        }
    }

    this.onTimeChange = function () { 

    }

    this.getDangerLevel = function (state) {
        switch (status) {
            case 1:
                return "危险"
            case 0:
                return "普通"
            default:
                return "未知"
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

    this.onOpenItemDetail = function (id,table) {
        if(id && table){
            for(var index = 0;index<table.length;index++){
                if(table[index].id == id){
                    table[index].bOpen = !table[index].bOpen
                    break
                }
            }
        }
    }
    this.tableData = [
        {id:1,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd [dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd ",value2:1,value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"2021-06-23T16:00:00.312",bOpen:false},
        {id:2,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",value2:1,value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"2021-06-23T16:00:00.312",bOpen:false},
        {id:3,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",value2:1,value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"2021-06-23T16:00:00.312",bOpen:false},
        {id:4,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",value2:1,value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"2021-06-23T16:00:00.312",bOpen:false},
        {id:5,value1:"[dby_web@Server-a4e6d510-e580-48a2-af92-48ce81ae860b ~]$ cd /etc/passwd",value2:1,value3:"luobing(luobing)",value4:"等保云demo服务器",value5:"2021-06-23T16:00:00.312",bOpen:false},
    ]
})