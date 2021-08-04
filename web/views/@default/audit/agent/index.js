Tea.context(function () {
    this.uploadFileSuccess = NotifyReloadSuccess("保存成功")
    this.pageState = 1

    this.fileDesc = ""

    this.bShowDialog = false
    this.nDialogTxt="正在上传中。。。"
    this.sTxtList=["正在上传中。。。","正在下载中。。。"]
    
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

    this.onDownFile = function (id,fileType,fileName) {
        teaweb.confirm("确定下载该文件？",function() {
            this.$get("/audit/agent/download").params({
                id: id,
            }).success((res)=>{
                this.onDownloadFlie(res,fileType,fileName)
            })
        })
    }

    this.onEdit = function (id) {
        let that =this
        teaweb.popup(Tea.url(".create?id="+id), {
			height: "200px",
            width:"460px",
			callback: function () {
                that.uploadFileSuccess()
				// window.location.reload()
			}
		})
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该文件？",function() {
            this.$post("/audit/agent/delete").params({
                Opt: "delete",
                id: id,
            }).refresh()
        })
    }


    this.onDownloadFlie = function(res,fileType,fileName){
        this.onShowLoading()
        try{
            let that = this
            this.nDialogTxt = this.sTxtList[1] 
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
              a.download = fileName;
              a.href = e.target.result;
              document.body.appendChild(a);
              a.click();
              document.body.removeChild(a);
              this.onHideLoading()
              that.uploadFileSuccess()
            }
        }catch(e){
            this.onHideLoading()
        }
        
    }

    this.onuploadFile = function () {
        this.nDialogTxt = this.sTxtList[0] 
        let that = this
        var uploadFile = document.getElementById("uploadFile");
        if(uploadFile.value==""){
            teaweb.warn("请选择上传文件")
            return
        }
        this.onShowLoading()
        console.log(uploadFile.files[0].type)
        var fm = document.getElementById('formData');
        var fd = new FormData(fm);
        fd.append('format',uploadFile.files[0].type);
        this.$post("/audit/agent").params(fd)
        .success(()=>{
            this.onHideLoading()
            that.uploadFileSuccess()
            return true
        })
        .done(()=>{
            this.onHideLoading()
        })
    }

    this.onShowLoading =  function() {
        this.bShowDialog = true
    }
     
     
    this.onHideLoading = function () {
        this.bShowDialog = false
    }

    // this.list = [
    //     {id:1,name:"数据库agent.exe",content_type:"CloudShield用户手册V1.0",used_bytes:"100K",last_modified:"2021-06-30 21:52:20"}
    // ]
})