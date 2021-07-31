Tea.context(function () {
    this.success = NotifyReloadSuccess("保存成功")
    this.pageState = 1

    this.fileDesc = ""

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

    this.onDownFile = function (name) {
        teaweb.confirm("确定下载该文件？",function() {
            this.$get("/databackup/download").params({
                name: name,
            }).success((res)=>{
                this.onDownLoadLocalFile(res.data.url,res.data.token)
            })
        })
    }

    this.onEdit = function (id) {
        teaweb.popup(Tea.url(".create"), {
			height: "300px",
            width:"460px",
			callback: function () {
				
			}
		})
    }

    this.onDelete = function (name) {
        teaweb.confirm("确定要删除该文件？",function() {
            this.$post("/databackup/delete").params({
                Opt: "delete",
                name: name,
            }).refresh()
        })
    }

    this.onDownLoadLocalFile=function(downUrl,headToken){
        // function sucHandle(res){
        //     console.log('sucHandle')
        //     console.log(res)
        // }
        // function failHandle(){
        //     console.log('failHandle')
        // }
        // Tea.openDownloadUrl("post",downUrl,headToken,sucHandle,failHandle)
        let xhr = new XMLHttpRequest()
        xhr.open("post",downUrl,true)
        xhr.setRequestHeader("Authorization",headToken)
        xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded")
        xhr.send()
        xhr.responseType="blob"
        xhr.onload=function(){
            if(this.status==200){
                let tempBlob = this.response
                let tempReader = new FileReader()
                reatempReaderder.readAsDataURL(tempBlob);
                tempReader.onload=function(e){
                    let link = document.createElement("a")
                    link.href = e.target.result
                    link.setAttribute("download","test")
                    link.click()
                    link=null
                }
            }
        }
    }

    // this.onuploadFile = function (file) {
    //     this.$post("/databackup").params({
    //         uploadFile: file
    //     }).success()
    // }

    this.tableData = [
        {id:1,value1:"数据库agent.exe",value2:"CloudShield用户手册V1.0",value3:"100K",value4:"2021-06-30T21:52:20.123"}
    ]
})