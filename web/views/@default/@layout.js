Tea.context(function () {
	this.moreOptionsVisible = false
	this.globalMessageBadge = 0

	this.curSelectCode = sessionStorage.getItem("leftSelectCode")? sessionStorage.getItem("leftSelectCode"):"dashboard"

	if (typeof this.leftMenuItemIsDisabled == "undefined") {
		this.leftMenuItemIsDisabled = false
	}
	this.$delay(function () {
		if (this.$refs.focus != null) {
			this.$refs.focus.focus()
		}

		let curSelectCode = sessionStorage.getItem("leftSelectCode")
        if(curSelectCode){
            this.onSetLeftTouchCode(curSelectCode)
        }

		// 检查消息
		this.checkMessages()
	})

	/**
	 * 左侧子菜单
	 */
	this.showSubMenu = function (menu) {
		if (menu.alwaysActive) {
			return
		}
		if (this.teaSubMenus.menus != null && this.teaSubMenus.menus.length > 0) {
			this.teaSubMenus.menus.$each(function (k, v) {
				if (menu.id == v.id) {
					return
				}
				v.isActive = false
			})
		}
		menu.isActive = !menu.isActive
	};

	/**
	 * 检查消息
	 */
	this.checkMessages = function () {
		this.$post("/messages/badge")
			.params({})
			.success(function (resp) {
				this.globalMessageBadge = resp.data.count
			})
			.done(function () {
				let delay = 6000
				if (this.globalMessageBadge > 0) {
					delay = 30000
				}
				this.$delay(function () {
					this.checkMessages()
				}, delay)
			})
	}

	/**
	 * 底部伸展框
	 */
	this.showQQGroupQrcode = function () {
		teaweb.popup("/about/qq", {
			width: "21em",
			height: "24em"
		})
	}

	/**
	 * 弹窗中默认成功回调
	 */
	if (window.IS_POPUP === true) {
		this.success = window.NotifyPopup
	}

	this.onChangeUrl = function (module,code) {
        let tempUrl = module.url

		//跳转到一级子菜单
		if(module.subItems && module.subItems.length>0){
			tempUrl = module.subItems[0].url
		}
        if(tempUrl){
            if(tempUrl.indexOf("nfw") != -1){
                let curSelectNode = localStorage.getItem("nfwSelectNodeId");
                if(curSelectNode){
                    tempUrl = tempUrl+"?nodeId="+curSelectNode
                }
            }else if(tempUrl.indexOf("ddos") != -1){
                let curSelectNode = localStorage.getItem("ddosSelectNodeId");
                if(curSelectNode){
                    tempUrl = tempUrl+"?nodeId="+curSelectNode
                }
            }else if (tempUrl === '/waf/alarm') {
				let curSelectNode = localStorage.getItem("nfwSelectNodeId");
				if(curSelectNode){
					tempUrl = tempUrl+"?nodeId="+curSelectNode
				}
			}else if(tempUrl.indexOf("hids") != -1){
				let agent = localStorage.getItem("hidsSelectAgentId");
				if(agent){
					tempUrl = tempUrl+"?agent="+agent
				}
			}
        }

        return tempUrl
    }

    this.onSetLeftTouchCode = function (code) {
        if(this.curSelectCode!=code){
            this.curSelectCode = code
        }
		// this.onOpenDialog()
        sessionStorage.setItem("leftSelectCode",this.curSelectCode)
    }

	this.onOpenDialog = function () {
        Tea.dialogBoxEnabled("block")
    }

});

window.NotifySuccess = function (message, url, params) {
	if (typeof (url) == "string" && url.length > 0) {
		if (url[0] != "/") {
			url = Tea.url(url, params);
		}
	}
	return function () {
		teaweb.success(message, function () {
			window.location = url;
		});
	};
};

window.NotifyReloadSuccess = function (message) {
	return function () {
		teaweb.success(message, function () {
			window.location.reload()
		})
	}
}

window.NotifyDelete = function (message, url, params) {
	teaweb.confirm(message, function () {
		Tea.Vue.$post(url)
			.params(params)
			.refresh();
	});
};

window.NotifyPopup = function (resp) {
	window.parent.teaweb.popupFinish(resp);
};

window.ChangePageSize = function (size) {
	let url = window.location.toString();
	if (url.indexOf("pageSize") > 0) {
		url = url.replace(/pageSize=\d+/g, "pageSize=" + size);
		url = url.replace(/page=\d+/g, "page=1");
	} else {
		if (url.indexOf("?") > 0) {
			url += "&pageSize=" + size+"&page=1";
		} else {
			url += "?pageSize=" + size+"&page=1";
		}
	}
	window.location = url;
};