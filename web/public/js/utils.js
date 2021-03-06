window.teaweb = {
	set: function (key, value) {
		localStorage.setItem(key, JSON.stringify(value));
	},
	get: function (key) {
		var item = localStorage.getItem(key);
		if (item == null || item.length == 0) {
			return null;
		}

		return JSON.parse(item);
	},
	getString: function (key) {
		var value = this.get(key);
		if (typeof (value) == "string") {
			return value;
		}
		return "";
	},
	getBool: function (key) {
		return Boolean(this.get(key));
	},
	remove: function (key) {
		localStorage.removeItem(key)
	},
	match: function (source, keyword) {
		if (source == null) {
			return false;
		}
		if (keyword == null) {
			return true;
		}
		source = source.trim();
		keyword = keyword.trim();
		if (keyword.length == 0) {
			return true;
		}
		if (source.length == 0) {
			return false;
		}
		var pieces = keyword.split(/\s+/);
		for (var i = 0; i < pieces.length; i++) {
			var pattern = pieces[i];
			pattern = pattern.replace(/(\+|\*|\?|[|]|{|}|\||\\|\(|\)|\.)/g, "\\$1");
			var reg = new RegExp(pattern, "i");
			if (!reg.test(source)) {
				return false;
			}
		}
		return true;
	},

	datepicker: function (element, callback) {
		if (typeof (element) == "string") {
			element = document.getElementById(element);
		}
		var year = new Date().getFullYear();
		var picker = new Pikaday({
			field: element,
			firstDay: 1,
			minDate: new Date(year - 1, 0, 1),
			maxDate: new Date(year + 10, 11, 31),
			yearRange: [year - 1, year + 10],
			format: "YYYY-MM-DD",
			i18n: {
				previousMonth: '上月',
				nextMonth: '下月',
				months: ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月'],
				weekdays: ['周日', '周一', '周二', '周三', '周四', '周五', '周六'],
				weekdaysShort: ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
			},
			theme: 'triangle-theme',
			onSelect: function () {
				if (typeof (callback) == "function") {
					callback.call(Tea.Vue, picker.toString());
				}
			}
		});
	},

	formatBytes: function (bytes) {
		bytes = Math.ceil(bytes);
		if (bytes < 1024) {
			return bytes + " bytes";
		}
		if (bytes < 1024 * 1024) {
			return (Math.ceil(bytes * 100 / 1024) / 100) + " k";
		}
		return (Math.ceil(bytes * 100 / 1024 / 1024) / 100) + " m";
	},
	formatNumber: function (x) {
		return x.toString().replace(/\B(?<!\.\d*)(?=(\d{3})+(?!\d))/g, ", ");
	},
	popup: function (url, options) {
		if (options == null) {
			options = {};
		}
		var width = "40em";
		var height = "20em";
		window.POPUP_CALLBACK = function () {
			Swal.close();
		};

		if (options["width"] != null) {
			width = options["width"];
		}
		if (options["height"] != null) {
			height = options["height"];
		}
		if (typeof (options["callback"]) == "function") {
			window.POPUP_CALLBACK = function () {
				Swal.close();
				options["callback"].apply(Tea.Vue, arguments);
			};
		}


        Swal.fire({
            html: '<iframe src="' + url + '#popup-' + width + '" style="border:0; width: 100%; height:' + height + '"></iframe>',
            width: width,
            padding: "0.5em",
            showConfirmButton: false,
            showCloseButton: true,
            focusConfirm: false,
            onClose: function (popup) {
                if (typeof (options["onClose"]) == "function") {
                    options["onClose"].apply(Tea.Vue, arguments)
                }
            }
        });
    },
    closePopup: function () {
		if (this.isPopup()) {
			window.parent.Swal.close();
		}
	},
    success: function (message, callback) {
        var width = "20em";
        if (message.length > 30) {
            width = "30em";
        }
		Swal.fire({
			html: '<iframe src="' + url + '#popup-' + width + '" style="border:0; width: 100%; height:' + height + '"></iframe>',
			width: width,
			padding: "0.5em",
			showConfirmButton: false,
			showCloseButton: true,
			focusConfirm: false,
			onClose: function (popup) {
				if (typeof (options["onClose"]) == "function") {
					options["onClose"].apply(Tea.Vue, arguments)
				}
			}
		});
	},
	popupFinish: function () {
		if (window.POPUP_CALLBACK != null) {
			window.POPUP_CALLBACK.apply(window, arguments);
		}
	},
	popupTip: function (html) {
		Swal.fire({
			html: '<i class="icon question circle"></i><span style="line-height: 1.7">' + html + "</span>",
			width: "30em",
			padding: "5em",
			showConfirmButton: false,
			showCloseButton: true,
			focusConfirm: false
		});
	},
	isPopup: function () {
		var hash = window.location.hash;
		return hash != null && hash.startsWith("#popup");
	},
	Swal: function () {
		return this.isPopup() ? window.parent.Swal : window.Swal;
	},
	success: function (message, callback) {
		var width = "20em";
		if (message.length > 30) {
			width = "30em";
		}
		let config = {
			confirmButtonText: "确定",
			buttonsStyling: false,
			icon: "success",
			customClass: {
				closeButton: "ui button",
				cancelButton: "ui button",
				confirmButton: "ui button primary"
			},
			width: width,
			onAfterClose: function () {
				if (typeof (callback) == "function") {
					setTimeout(function () {
						callback();
					});
				} else if (typeof (callback) == "string") {
					window.location = callback
				}
			}
		}

		if (message.startsWith("html:")) {
			config.html = message.substring(5)
		} else {
			config.text = message
		}

		Swal.fire(config);
	},
	successToast: function (message, timeout) {
		if (timeout == null) {
			timeout = 2000
		}
		var width = "20em";
		if (message.length > 30) {
			width = "30em";
		}
		Swal.fire({
			text: message,
			icon: "success",
			width: width,
			timer: timeout,
			showConfirmButton: false
		});
	},
	warn: function (message, callback) {
		var width = "20em";
		if (message.length > 30) {
			width = "30em";
		}
		Swal.fire({
			text: message,
			confirmButtonText: "确定",
			buttonsStyling: false,
			customClass: {
				closeButton: "ui button",
				cancelButton: "ui button",
				confirmButton: "ui button primary"
			},
			icon: "warning",
			width: width,
			onAfterClose: function () {
				if (typeof (callback) == "function") {
					setTimeout(function () {
						callback();
					});
				}
			}
		});
	},
	confirm: function (message, callback) {
		let width = "20em";
		if (message.length > 30) {
			width = "30em";
		}
		let config = {
			confirmButtonText: "确定",
			cancelButtonText: "取消",
			showCancelButton: true,
			showCloseButton: false,
			buttonsStyling: false,
			customClass: {
				closeButton: "ui button",
				cancelButton: "ui button",
				confirmButton: "ui button primary"
			},
			icon: "warning",
			width: width,
			preConfirm: function () {
				if (typeof (callback) == "function") {
					callback.call(Tea.Vue);
				}
			}
		}
		if (message.startsWith("html:")) {
			config.html = message.substring(5)
		} else {
			config.text = message
		}
		Swal.fire(config);
	},
	reload: function () {
		window.location.reload()
	}
};

String.prototype.quoteIP = function () {
	let ip = this.toString()
	if (ip.length == 0) {
		return ""
	}
	if (ip.indexOf(":") < 0) {
		return ip
	}
	if (ip.substring(0, 1) == "[") {
		return ip
	}
	return "[" + ip + "]"
}