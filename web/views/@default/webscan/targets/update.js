Tea.context(function () {
    this.success = NotifySuccess("保存成功", "/webscan/targets")
    this.moreSet = this.username !== ""
    this.changeMoreSetEditing = function () {
        this.moreSet = !this.moreSet
    }
    this.radioSet = 1
    this.changeRadioEditing = function (radio) {
        this.radioSet = radio
    }
    this.changeOsEditing = function (os) {
        this.osSet = os
    }
})