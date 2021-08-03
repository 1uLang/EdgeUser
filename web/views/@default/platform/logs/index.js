Tea.context(function () {
    this.dayFrom = ""
    this.dayTo = ""
    this.keyword = ""
    this.page = ""

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

    //test
    // this.onSearch = function(){

    //     let tempPassword = this.getPassword()
    //     let curParam = {username:"admin",password:tempPassword}
    //     var xhr = new XMLHttpRequest();
    //     xhr.open("post","http://182.150.0.115:8092/core/auth/login/",true);//true为异步，false为同步
    //     xhr.onreadystatechange=function(){
    //         if (xhr.readystate == 4){
    //             console.log(xhr)
    //             // if(xhr.status == 200){
    //             //     window.open("http://182.150.0.115:8092/luna/?login_to=e5db747e-5eb9-453e-a861-7f047e641cd4")
    //             // }
    //         }
    //     }
    //     xhr.send(JSON.stringify(curParam));

    // }
    // this.encryptLoginPassword = function(password,rsaPublicKey){
    //     var jsencrypt = new JSEncrypt.JSEncrypt(); //加密对象
    //     jsencrypt.setPublicKey(rsaPublicKey); // 设置密钥
    //     return jsencrypt.encrypt(password); //加密

    // }

    // this.getPassword = function() {
    //     //公钥加密
    //     var rsaPublicKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC8i4VAYKaT4QYK7+AnmVGH1dpC\nfZYYB7YoqeyyJ6doXwXEucqkFM5IZio9ElgHZAPjGBTUBVrvhpZFhmMnvdtnM84O\nKm6Ptkf+DfFtWRCxc+QKOm80rgyOjd22W4kkvWFTiF0xtfYfMk3ay+bvNf5xgoFQ\ngfECFLN+mgC27iir9QIDAQAB\n-----END PUBLIC KEY-----"
    //     var passwordEncrypted = this.encryptLoginPassword("21ops.com.C", rsaPublicKey)
    //     return passwordEncrypted
    // }

    this.exportExcel = function () {
        let that = this
        teaweb.confirm("确定要将当前列表导出到Excel吗？", function () {

        })
    }

    this.logs = [
        {id:1,level:"error",createdTime:"2021-08-03 16:21:06",userName:"user01",userId:"015455",ip:"192.168.1.1",region:"this is region",action:"this is action",description:"this is description"},
    ]
})