Tea.context(function () {

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