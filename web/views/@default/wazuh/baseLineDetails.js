Tea.context(function () {

    this.getShowSelectValueImage = function (bOpen) {
        if (bOpen) {
            return "/images/image-grey-open.png"
        }
        return "/images/image-grey-close.png"
    }

    this.$delay(function () {

        for (let idx = 0; idx < this.details.length; idx++) {
            this.details[idx].bOpen = false
        }
    })

    this.onEnabledDetail = function (item, idx) {

        // this.details[idx].bOpen = !this.details[idx].bOpen
        item.bOpen = !item.bOpen
    }
})