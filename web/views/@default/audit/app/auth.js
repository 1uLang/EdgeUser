Tea.context(function () {
    // this.authValue = ""
    this.onSaveAuth = function () {
        this.$post(".auth")
            .params({
                email: this.authValue,
                id:this.id,
            }).success(function (){
            this.success("保存成功", function () {
                // window.location = "/audit/db"
            })
        })

        // teaweb.closePopup()
    }
})