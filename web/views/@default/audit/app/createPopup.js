Tea.context(function () {

    //这边数据是从服务器获取的 本地只做展示
    this.name = ""
    this.typeSelect = -1
    this.ip = ""
    this.openState = 1
    
    this.onSelectOpenState = function () {
        this.openState = this.openState===1 ? 2:1
    }

})