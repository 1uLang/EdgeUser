Tea.context(function () {
    this.success = NotifyReloadSuccess("保存成功")

    this.onListenEditCheckBox = function () {
        this.enable = !this.enable;
    }

})