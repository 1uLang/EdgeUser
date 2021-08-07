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
})