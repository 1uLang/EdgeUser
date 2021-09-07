Tea.context(function () {

    this.moreSet = false

    this.changeMoreSetEditing = function () {
        this.moreSet = !this.moreSet
    }
    this.radioSet = 1
    this.changeRadioEditing = function (radio) {
        this.radioSet = radio
    }
    this.osSet = 1
    this.changeOsEditing = function (os) {
        this.osSet = os
    }
})