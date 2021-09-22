Tea.context(function () {
	this.tempToken = "";
    this.tempCsrfToke = "";
	this.tempShowOTP = false;
	this.showPageState = 1;

	this.bShowDialog = false
	this.errorMsg = ""
	this.tipImage = ""
	this.callBackFunc = null

	this.$delay(function () {
		sessionStorage.setItem("leftSelectCode","dashboard")
		localStorage.removeItem("ddosSelectNodeId")
		localStorage.removeItem("nfwSelectNodeId")
		this.onGetRefreshToken()
		this.onGetToken()
	});


	this.onGetRefreshToken = function () {
		let that = this
        that.refreshCsrfToken();
        setInterval(function () {
			that.refreshCsrfToken();
        }, 10 * 60 * 1000);
      }

	this.refreshCsrfToken = function () {
		let that = this
        reqApi("get", "/csrf/token", null, null, 
		(res) => {
			that.tempCsrfToke = res.data.token;
        },
		(res)=>{
			that.onOpenErrorDialog(res.message)
		}
		);
    }

    this.onGetToken = function () {
		let that = this
        reqApi("get", "/?token=1", null, null, 
		(res) => {
			that.tempToken = res.data.token;
        },
		(res)=>{
			that.onOpenErrorDialog(res.message)
		}
		);
      }

	this.onRefreshShowOpt = function () {
		let that = this
		var tempUserName = document.getElementById("username").value;
        var tempFormData = new FormData();
        tempFormData.append("username", tempUserName);
        reqApi("post", "/checkOTP", tempFormData, null, (res) => {
          	that.tempShowOTP = res.data.requireOTP;
        });
      }

	this.onCheckDoLogin = function(frm, event) {
        var event = window.event ? window.event : event;
        if (event.keyCode == 13) {
          this.login();
        }
      }

	this.login = function () {
        var tempUserName = document.getElementById("username").value;
        var tempPassword = document.getElementById("password").value;
        tempPassword = md5(tempPassword.trim());
        var tempFormData = new FormData();
        tempFormData.append("username", tempUserName);
        tempFormData.append("password", tempPassword);
        tempFormData.append("token", this.tempToken);
        tempFormData.append("csrfToken", this.tempCsrfToke);
        if (this.tempShowOTP) {
          var tempOTPCode = document.getElementById("otpCode").value;
          tempFormData.append("otpCode", tempOTPCode);
        }

		let that = this
        reqApi(
          "post",
          "",
          tempFormData,
          null,
          (res) => {
            if (res.code != 200) {
              window.location.reload();
            }
            if (res.data.from != null && res.data.from.length > 0) {
              window.location = res.data.from;
            } else {
              window.location = "/dashboard";
            }
          },
          (res) => {
            if (res.data.from == "/updatePwd") { //如果是密码过期
				that.callBackFunc = function () {
					document.getElementById("password").value = "";
					that.showPageState = 2
					that.onGetRefreshToken();
					that.onGetToken();
					setTimeout(()=>{
						document.getElementById("resetPassword").value = "";
						document.getElementById("confirmPassword").value = "";
					},10)
				}
				that.onOpenErrorDialog(res.message)
              
            } else if (res.data.from == "/renewal") {//到期 续订

			  	that.callBackFunc = function () {
					that.showPageState = 3
                	document.getElementById("systemCode").value = res.data.systemCode;
			 	}
				that.onOpenErrorDialog(res.message)
            } else if (res.data.from == "/页面过期") {//刷新页面
              
			  	that.callBackFunc = function () {
					window.location.reload();
			 	}
				that.onOpenErrorDialog(res.message)
             
            } else {
				that.callBackFunc = function () {
					that.onGetRefreshToken();
                	that.onGetToken();
			 	}
				that.onOpenErrorDialog(res.message)
            }
          }
        );
      }

	this.onCheckDoReset = function (frm, event) {
        var event = window.event ? window.event : event;
        if (event.keyCode == 13) {
          this.onResetPwd();
        }
    }
    this.onCheckDoRenewal = function (frm, event) {
        var event = window.event ? window.event : event;
        if (event.keyCode == 13) {
          this.onRenewal();
        }
    }

	this.onResetPwd = function () {

        var tempPassword = document.getElementById("resetPassword").value;
        var tempConfirmPassword =
          document.getElementById("confirmPassword").value;
        tempPassword = tempPassword.trim();
        tempConfirmPassword = tempConfirmPassword.trim();
        var tempFormData = new FormData();
        tempFormData.append("password", tempPassword);
        tempFormData.append("confirmPassword", tempConfirmPassword);
        tempFormData.append("token", this.tempToken);
        tempFormData.append("csrfToken", this.tempCsrfToke);
		let that = this
        reqApi(
          "post",
          "/updatePwd",
          tempFormData,
          null,
          (res) => {
			that.callBackFunc = function () {
				
				that.showPageState = 1
				that.onGetRefreshToken();
				that.onGetToken();
			}
			that.onOpenSucDialog(res.message)
          },
          (res) => {
			that.callBackFunc = function () {
				document.getElementById("resetPassword").value = "";
				document.getElementById("confirmPassword").value = "";
				that.onGetRefreshToken();
				that.onGetToken();
			 }
			that.onOpenErrorDialog(res.message)
          }
        );
      }

	this.onRenewal = function() {
        var secret = document.getElementById("secret").value;
        var tempFormData = new FormData();
        tempFormData.append("secret", secret);
		let that = this
        reqApi(
          "post",
          "/renewal",
          tempFormData,
          null,
          (res) => {
			that.callBackFunc = function () {
				that.showPageState = 1
				that.onGetRefreshToken();
				that.onGetToken();
			}
			that.onOpenSucDialog(res.message)
          },
          (res) => {
			that.callBackFunc = function () {
				that.onGetRefreshToken();
				that.onGetToken();
			 }
			that.onOpenErrorDialog(res.message)
          }
        );
      }

	this.onOpenErrorDialog = function(errorMsg){
		if(errorMsg && errorMsg.length>0){
			this.errorMsg = errorMsg
		}else{
			this.errorMsg = "请求发生异常"
		}
		this.tipImage="/images/image_login_tip_warring.png"
		this.bShowDialog = true
	}
	this.onOpenSucDialog = function(tipMsg){
		if(tipMsg && tipMsg.length>0){
			this.errorMsg = tipMsg
		}else{
			this.errorMsg = "操作成功"
		}
		this.tipImage="/images/image_login_tip_suc.png"
		this.bShowDialog = true
		
		
	}
	this.onCloseDialog = function(){
		this.bShowDialog = false
		if(this.callBackFunc!=null){
			this.callBackFunc()
			this.callBackFunc = null
		}
	}

});