Tea.context(function () {

    this.otpIsOn = false
    this.onListenEditCheckBoxOTP = function () {
        let enabled = document.getElementById("otpIsOn").checked
        if (enabled) {
            this.otpIsOn = true
        } else {
            this.otpIsOn = false
        }
    }
    this.submitBefore = function () {
        //获取表单提交按钮
        var btnSubmit = document.getElementById("submit-btn");
        //将表单提交按钮设置为不可用，可以避免用户再次点击提交按钮进行提交
        btnSubmit.disabled = "disabled";
    }
    this.submitFail = function (resp) {

        var error = resp.errors[0].messages[0]
        teaweb.warn(error)
        //获取表单提交按钮
        var btnSubmit = document.getElementById("submit-btn");
        //将表单提交按钮设置为不可用，可以避免用户再次点击提交按钮进行提交
        btnSubmit.disabled = false;
    }
    this.submitError = function (resp) {
        var btnSubmit = document.getElementById("submit-btn");
        //将表单提交按钮设置为不可用，可以避免用户再次点击提交按钮进行提交
        btnSubmit.disabled = false;
    }
})