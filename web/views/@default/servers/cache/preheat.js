Tea.context(function () {
    this.success = NotifyReloadSuccess("预热成功")

    this.isRequesting = false
    this.before = function () {
        this.isRequesting = true
    }

    this.done = function () {
        this.isRequesting = false
    }
})