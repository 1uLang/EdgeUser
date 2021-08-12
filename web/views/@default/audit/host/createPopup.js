Tea.context(function () {

    //这边数据是从服务器获取的 本地只做展示
    // this.name = ""
    // this.ip = ""
    // this.systemSelect = 1
    // this.openState = 1
    
    this.onSelectOpenState = function () {
        this.openState = this.openState===1 ? 0:1
    }

    this.onSelectSystem = function () {
        this.systemSelect = this.systemSelect===1 ? 0:1
    }

    this.onSave = function(){
        this.$post("/audit/host/createPopup")
            .params({
                name: this.name,
                ip: this.ip,
                system: this.systemSelect,
                status: this.openState,
                id:this.id,
            })
            .success(function (){
                this.success("保存成功", function () {
                    window.location = "/audit/host"
                })
            })
    }
})