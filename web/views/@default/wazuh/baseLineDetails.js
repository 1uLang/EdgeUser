Tea.context(function () {

    this.$delay(function () {
        if (this.errorMsg && this.errorMsg != "") {
            teaweb.warn(this.errorMsg)
        }
    })
    this.getShowSelectValueImage = function (bOpen) {
        if (bOpen) {
            return "/images/image-grey-open.png"
        }
        return "/images/image-grey-close.png"
    }


    this.onEnabledDetail = function (item) {
        item.bOpen = !item.bOpen
    }
})