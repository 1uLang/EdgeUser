Tea.context(function () {
    this.type = "file"

    this.success = NotifyReloadSuccess("刷新成功")

    this.isRequesting = false
    this.before = function () {
        this.isRequesting = true
    }

    this.done = function () {
        this.isRequesting = false
    }
})