Tea.context(function () {
    this.success = NotifySuccess("保存成功", "/databackup?Dirpath="+this.url )

    this.passwordEditing = false

    this.changePasswordEditing = function () {
        this.passwordEditing = !this.passwordEditing
    }
})