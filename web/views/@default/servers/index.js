Tea.context(function () {
	this.$delay(function () {
		this.loadStatus()
	})

	this.loadStatus = function () {
		let serverIds = this.servers.map(function (v) {
			return v.id
		})
		this.$post(".status")
			.params({
				serverIds: serverIds
			})
			.success(function (resp) {
				let status = resp.data.status
				this.servers.forEach(function (server) {
					if (typeof status[server.id] === "object") {
						server.status = status[server.id]
					}
				})
			})
	}

	this.deleteServer = function (serverId) {
		let that = this
		teaweb.confirm("确定要删除此域名吗？", function () {
			that.$post(".delete")
				.params({
					serverId: serverId
				})
				.refresh()
		})
	}

	this.updateServerOn = function (serverId) {
		let that = this
		teaweb.confirm("确定要启用此域名吗？", function () {
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
		teaweb.confirm("确定要停用此域名吗？", function () {
			that.$post(".updateOn")
				.params({
					serverId: serverId,
					isOn: false
				})
				.refresh()
		})
	}
})