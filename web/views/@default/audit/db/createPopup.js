Tea.context(function () {

    //这边数据是从服务器获取的 本地只做展示
    // this.name = ""
    // this.typeSelect = -1 //类型
    // this.verSelect = -1 //版本
    // this.ip = "1.1.1.1"
    // this.port = ""
    // this.systemSelect = 0
    // this.openState = 1
    
    this.onSelectOpenState = function () {
        this.openState = this.openState===1 ? 0:1
    }

    this.onSelectSystem = function () {
        this.systemSelect = this.systemSelect===1 ? 0:1
    }

    this.onSave = function(){
        this.$post("/audit/db/createPopup")
            .params({
                name: this.name,
                type: this.typeSelect,
                version: this.verSelect,
                ip: this.ip,
                port: this.port,
                system: this.systemSelect,
                status: this.openState,
                id:this.id,
            })
            .success(function (){
                this.success("保存成功", function () {
                    window.location = "/audit/db"
                })
            })
    }
})