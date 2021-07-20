Tea.context(function () {


    this.pageState = 1

    this.onChangeState = function (state) { 
        if( this.pageState!= state) {
            this.pageState = state
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

    //中断
    this.onStop = function (id) { 

    }

    //监控
    this.onStart = function (id) { 

    }

    //回放
    this.onReplay = function (id) { 

    }

    this.tableData1=[
        {id:1,value1:"luobing(luobing)",value2:"等保云demo服务器",value4:"118.116.10.36",value5:"ssh",value6:"Web Terminal",value7:"1",value8:"2021-06-23T16:00:00.012",value9:"7.0秒"},
        {id:2,value1:"luobing(luobing)",value2:"等保云demo服务器",value4:"118.116.10.36",value5:"ssh",value6:"Web Terminal",value7:"1",value8:"2021-06-23T16:00:00.012",value9:"7.0秒"},
        {id:3,value1:"luobing(luobing)",value2:"等保云demo服务器",value4:"118.116.10.36",value5:"ssh",value6:"Web Terminal",value7:"1",value8:"2021-06-23T16:00:00.012",value9:"7.0秒"}
    ]
    this.tableData2=[
        {id:1,value1:"luobing(luobing)",value2:"等保云demo服务器",value4:"118.116.10.36",value5:"ssh",value6:"Web Terminal",value7:"1",value8:"2021-06-23T16:00:00.012",value9:"7.0秒"},
        {id:2,value1:"luobing(luobing)",value2:"等保云demo服务器",value4:"118.116.10.36",value5:"ssh",value6:"Web Terminal",value7:"1",value8:"2021-06-23T16:00:00.012",value9:"7.0秒"},
        {id:3,value1:"luobing(luobing)",value2:"等保云demo服务器",value4:"118.116.10.36",value5:"ssh",value6:"Web Terminal",value7:"1",value8:"2021-06-23T16:00:00.012",value9:"7.0秒"}
    ]
})