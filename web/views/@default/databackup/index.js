Tea.context(function () {

    this.pageState = 1

    this.fileDesc = ""
    this.uploadFileName = "未选择任何文件"

    this.onChangeState=function (id) {
        if(this.pageState!=id){
            this.pageState = id
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

    this.onDownFile = function (id) {
        
    }

    this.onEdit = function (id) {
        teaweb.popup(Tea.url(".create"), {
			height: "300px",
            width:"460px",
			callback: function () {
				
			}
		})
    }

    this.onDelete = function (id) {
        
    }

    this.tableData = [
        {id:1,value1:"数据库agent.exe",value2:"CloudShield用户手册V1.0",value3:"100K",value4:"2021-06-30T21:52:20.123"}
    ]
})