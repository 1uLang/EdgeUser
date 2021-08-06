Tea.context(function () {
	this.confirmPassword = ""
	this.confirmPasswordMd5 = ""
	this.password = ""
	this.passwordMd5 = ""


	this.token = ""

	this.isSubmitting = false

	this.$delay(function () {

	});

	this.changePassword = function () {
		// this.passwordMd5 = md5(this.password.trim());
		this.passwordMd5 = this.password.trim()
	};
	this.changePassword1 = function () {
		// this.confirmPasswordMd5 = md5(this.confirmPassword.trim());
		this.confirmPasswordMd5 = this.confirmPassword.trim()
	};
	this.submitBefore = function () {
		this.isSubmitting = true;
	};

	this.submitDone = function () {
		this.isSubmitting = false;
	};

	this.submitSuccess = function () {
		setTimeout(function() {
			window.location = "/";
		}, 1000);  //1秒后将会调用执行函数

	};
});