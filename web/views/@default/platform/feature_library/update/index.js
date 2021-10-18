Tea.context(function () {

    this.getTime = function (time) {
        var d = new Date(time);
        return d.toLocaleDateString() + " " + d.toLocaleTimeString()
    }


    this.update = function (){
        this.status="更新中.."
        this.$get("platform/feature_library/status_update")
            .params({
            })
            .refresh()

    }

    this.onUpdate = function () {
        let open = !this.authUpdate
        this.$get("platform/feature_library/auth_update")
            .params({
                open:open,
            })
            .refresh()
    }
})