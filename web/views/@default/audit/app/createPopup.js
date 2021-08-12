Tea.context(function () {

    //这边数据是从服务器获取的 本地只做展示
    // this.name = ""
    // this.typeSelect = -1
    // this.ip = ""
    // this.openState = 1
    
    this.onSelectOpenState = function () {
        this.openState = this.openState===1 ? 0:1
    }

    this.onSave = function(){
        this.$post(".createPopup")
            .params({
                name: this.name,
                type: this.typeSelect,
                ip: this.ip,
                status: this.openState,
                id:this.id,
            })
            .success(function (){
                this.success("保存成功", function () {
                    window.location = "/audit/app"
                })
            })
    }
})