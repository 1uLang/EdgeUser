Tea.context(function () {
    this.uploadFileSuccess = NotifyReloadSuccess("保存成功")
    this.pageState = 1

    this.fileDesc = ""

    this.bShowDialog = false
    this.nDialogTxt="正在上传中..."
    this.sTxtList=["正在上传中...","正在下载中...","正在创建中..."]

    this.bShowView = false
    this.bEditFileBagName = false
    this.nUploadPath = ""

    this.$delay(function () {
        if(this.title && this.title.length > 0){
            this.onCreateTitle(this.title)
        }else{
            this.nUploadPath = ""
        }
    })

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

    this.onDownFile = function (fileName,url,fileType) {
        
        let that =this
        teaweb.confirm("确定下载该文件？",function() {
            that.nDialogTxt = this.sTxtList[1] 
            that.onShowLoading()
            that.$get("/databackup/download").params({
                name: fileName,
                fp: url
            }).timeout(120).success((res)=>{
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

    this.onDelete = function (url) {
        teaweb.confirm("确定要删除该文件？",function() {
            this.$post("/databackup/delete").params({
                Opt: "delete",
                fp: url,
            }).refresh()
        })
    }

    this.onCreateFileBag = function(){
        this.nDialogTxt = this.sTxtList[2] 
        console.log(this.nUploadPath)
        var tempFileBagName = document.getElementById("fileBagName").value
        this.$post("/databackup/dir").params({
            purl: this.nUploadPath,
	        name:tempFileBagName
        })
        .timeout(120)
        .success(()=>{
            this.onHideLoading()
            this.uploadFileSuccess()
            return true
        }).done(()=>{
            this.onHideLoading()
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
        var fd = new FormData();
        fd.append("uploadFile",uploadFile.files[0]);
        fd.append("dirpath",this.nUploadPath);

        this.$post("/databackup").params(fd)
        .timeout(120)
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
    
    this.onCreateTitle = function(titleData){
        var parentNode = document.getElementById("parentNode")
        var titleNode = document.getElementById("titleNode")
        if(titleNode){
            parentNode.removeChild(titleNode);
        }
        
        var newTitleNode = document.createElement('div')
        newTitleNode.setAttribute("id","titleNode");
        for(var index = 0;index<titleData.length;index++){
            // var newSpan = document.createElement('span')
            // var newText = document.createTextNode(titleData[index].Name+" >");
            // newSpan.setAttribute("url",titleData[index].Url);
            // newSpan.appendChild(newText);
            // newSpan.style.cursor = 'pointer'
            // newSpan.style.marginRight = '2px'
            // newSpan.style.color = '#276ac6'
            // newSpan.onclick = function(e){
            //     console.log(e.target)
            //     var openUrl = e.target.getAttribute("url")
            //     window.location = "/databackup?Dirpath=/"+openUrl
            // }
            // newTitleNode.appendChild(newSpan)
            var newA = document.createElement('a')
            var newText = document.createTextNode(titleData[index].Name+" >");
            newA.appendChild(newText);
            newA.style.cursor = 'pointer'
            newA.style.marginRight = '2px'
            newA.href = "/databackup?Dirpath=/" +unescape(titleData[index].Url)
            newTitleNode.appendChild(newA)
        }
        parentNode.appendChild(newTitleNode)
        this.nUploadPath = titleData[titleData.length-1].Url
    }

    this.onOpenHome = function(){
        window.location = "/databackup?Dirpath="
    }
    this.onOpenFile = function(item){
        // window.location = item.url
        window.location = "/databackup?Dirpath="+item.url
        // if(item.childFile && item.childFile.length>0){
        //     var childFileData = this.getItemInfo(item.id,tableData)
        //     if(childFileData){
        //         this.list = this.deepClone(childFileData)
        //         this.titleData = [
        //             {
        //                 titleName:"数据库agent >",
        //                 titleUrl:"",
        //             },
        //         ]
        //         this.onCreateTitle()
        //     }
        // }else{
        //     this.onDownFile(item.name,item.url,item.content_type)
        // }
        
    }

    this.getItemInfo = function (id,tableData) {
        if(id && tableData){
            for (var index=0;index<tableData.length;index++){
                if(tableData[index].id ==id){
                    return tableData[index].childFile
                }
            }
        }
        
        return null
    }
 
    this.deepClone = function(source){
        if (!source && typeof source !== 'object') {
            throw new Error('error arguments', 'deepClone')
        }
        const targetObj = source.constructor === Array ? [] : {}
        Object.keys(source).forEach((keys) => {
        if (source[keys] && typeof source[keys] === 'object') {
            targetObj[keys] = this.deepClone(source[keys])
        } else {
            targetObj[keys] = source[keys]
        }
        })
        return targetObj
    }

    this.onOpenView = function(){
        this.bShowView = !this.bShowView
    }

    this.onCloseView = function(){
        this.bShowView = false
        this.bEditFileBagName = false
    }

    this.onEditFileBagName = function(){
        this.bEditFileBagName = true
    }

    // this.list = []

    // this.dataList = [
    //     {   
    //         id:1,
    //         name:"数据库agent",
    //         content_type:"CloudShield用户手册V1.0",
    //         used_bytes:"100K",
    //         value4:"2021-06-30T21:52:20.123",
    //         childFile:[
    //             {
    //                 id:11,
    //                 name:"数据库agent2.exe",
    //                 content_type:"CloudShield用户手册V2.0",
    //                 used_bytes:"12K",
    //                 value4:"2021-06-30T21:52:20.123",
    //                 childFile:{
                        
    //                 }
            
    //             }
    //         ]
    //     }
    // ]

    // this.titleData = [
        // {
        //     titleName:"一级菜单 >",
        //     titleUrl:"",
        // },
        // {
        //     titleName:"二级菜单 >",
        //     titleUrl:"",
        // },
        // {
        //     titleName:"三级菜单 >",
        //     titleUrl:"",
        // }
    // ]
    
})