Tea.context(function () {
	this.deleteServer = function (serverId) {
		let that = this
		teaweb.confirm("确定要删除此服务吗？", function () {
			that.$post(".delete")
				.params({
					serverId: serverId
				})
				.refresh()
		})
	}

	this.updateServerOn = function (serverId) {
		let that = this
		teaweb.confirm("确定要启用此服务吗？", function () {
			that.$post(".updateOn")
				.params({
					serverId: serverId,
					isOn: true
				})
				.refresh()
		})
	}

	this.updateServerOff = function (serverId) {
		let that = this
		teaweb.confirm("确定要停用此服务吗？", function () {
			that.$post(".updateOn")
				.params({
					serverId: serverId,
					isOn: false
				})
				.refresh()
		})
	}
})