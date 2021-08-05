Tea.context(function () {
    this.uploadFileSuccess = NotifyReloadSuccess("保存成功")
    this.pageState = 1

    this.fileDesc = ""

    this.bShowDialog = false
    this.nDialogTxt="正在上传中..."
    this.sTxtList=["正在上传中...","正在下载中..."]

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

    this.onDownFile = function (fileName,fileType) {
        
        let that =this
        teaweb.confirm("确定下载该文件？",function() {
            that.nDialogTxt = this.sTxtList[1] 
            that.onShowLoading()
            that.$get("/databackup/download").params({
                name: fileName,
            }).success((res)=>{
                that.onDownloadFlie(res,fileType)
            }).fail((res)=>{
                that.onHideLoading()
                teaweb.warn(res.message)
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


    this.onDownloadFlie = function(res,fileType){
        try{
            let that = this
            var bstr = atob(res.data.body)
            let n = bstr.length
            let u8arr =new Uint8Array(n)
            while (n--) {
                u8arr[n] = bstr.charCodeAt(n);
            }

            const blob = new Blob([u8arr], { type:fileType});
            const reader = new FileReader();
            reader.readAsDataURL(blob);
            reader.onload = (e) => {
                const a = document.createElement('a');
                a.download = res.data.fileName;
                a.href = e.target.result;
                document.body.appendChild(a);
                a.click();
                document.body.removeChild(a);
                that.onHideLoading()
                that.uploadFileSuccess()
            }
        }catch(e){
            this.onHideLoading()
        }
        
    }

    this.onuploadFile = function () {
        this.nDialogTxt = this.sTxtList[0]
        let that = this
        var uploadFile = document.getElementById("uploadFile");
        if(uploadFile.value==""){
            teaweb.warn("请选择上传文件")
            return
        }
        this.onShowLoading()
        var fm = document.getElementById('formData');
        var fd = new FormData(fm);

        this.$post("/databackup").params(fd)
        .success(()=>{
            this.onHideLoading()
            that.uploadFileSuccess()
            return true
        }).done(()=>{
            this.onHideLoading()
        })
    }
    this.onShowLoading =  function() {
        this.bShowDialog = true
    }
     
     
    this.onHideLoading = function () {
        this.bShowDialog = false
    }

    this.tableData = [
        {id:1,value1:"数据库agent.exe",value2:"CloudShield用户手册V1.0",value3:"100K",value4:"2021-06-30T21:52:20.123"}
    ]
})