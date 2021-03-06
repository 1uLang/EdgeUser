Vue.component("health-check-config-box", {
	props: ["v-health-check-config"],
	data: function () {
		let healthCheckConfig = this.vHealthCheckConfig
		let urlProtocol = "http"
		let urlPort = ""
		let urlRequestURI = "/"

		if (healthCheckConfig == null) {
			healthCheckConfig = {
				isOn: false,
				url: "",
				interval: {count: 60, unit: "second"},
				statusCodes: [200],
				timeout: {count: 10, unit: "second"},
				countTries: 3,
				tryDelay: {count: 100, unit: "ms"},
				autoDown: true,
				countUp: 1,
				countDown: 1
			}
			let that = this
			setTimeout(function () {
				that.changeURL()
			}, 500)
		} else {
			try {
				let url = new URL(healthCheckConfig.url)
				urlProtocol = url.protocol.substring(0, url.protocol.length - 1)
				urlPort = url.port
				urlRequestURI = url.pathname
				if (url.search.length > 0) {
					urlRequestURI += url.search
				}
			} catch (e) {
			}

			if (healthCheckConfig.statusCodes == null) {
				healthCheckConfig.statusCodes = [200]
			}
			if (healthCheckConfig.interval == null) {
				healthCheckConfig.interval = {count: 60, unit: "second"}
			}
			if (healthCheckConfig.timeout == null) {
				healthCheckConfig.timeout = {count: 10, unit: "second"}
			}
			if (healthCheckConfig.tryDelay == null) {
				healthCheckConfig.tryDelay = {count: 100, unit: "ms"}
			}
			if (healthCheckConfig.countUp == null || healthCheckConfig.countUp < 1) {
				healthCheckConfig.countUp = 1
			}
			if (healthCheckConfig.countDown == null || healthCheckConfig.countDown < 1) {
				healthCheckConfig.countDown = 1
			}
		}
		console.log(healthCheckConfig.countUp, healthCheckConfig.countDown)
		return {
			healthCheck: healthCheckConfig,
			advancedVisible: false,
			urlProtocol: urlProtocol,
			urlPort: urlPort,
			urlRequestURI: urlRequestURI
		}
	},
	watch: {
		urlRequestURI: function () {
			if (this.urlRequestURI.length > 0 && this.urlRequestURI[0] != "/") {
				this.urlRequestURI = "/" + this.urlRequestURI
			}
			this.changeURL()
		},
		urlPort: function (v) {
			let port = parseInt(v)
			if (!isNaN(port)) {
				this.urlPort = port.toString()
			} else {
				this.urlPort = ""
			}
			this.changeURL()
		},
		urlProtocol: function () {
			this.changeURL()
		},
		"healthCheck.countTries": function (v) {
			let count = parseInt(v)
			if (!isNaN(count)) {
				this.healthCheck.countTries = count
			} else {
				this.healthCheck.countTries = 0
			}
		},
		"healthCheck.countUp": function (v) {
			let count = parseInt(v)
			if (!isNaN(count)) {
				this.healthCheck.countUp = count
			} else {
				this.healthCheck.countUp = 0
			}
		},
		"healthCheck.countDown": function (v) {
			let count = parseInt(v)
			if (!isNaN(count)) {
				this.healthCheck.countDown = count
			} else {
				this.healthCheck.countDown = 0
			}
		}
	},
	methods: {
		showAdvanced: function () {
			this.advancedVisible = !this.advancedVisible
		},
		changeURL: function () {
			this.healthCheck.url = this.urlProtocol + "://${host}" + ((this.urlPort.length > 0) ? ":" + this.urlPort : "") + this.urlRequestURI
		},
		changeStatus: function (values) {
			this.healthCheck.statusCodes = values.$map(function (k, v) {
				let status = parseInt(v)
				if (isNaN(status)) {
					return 0
				} else {
					return status
				}
			})
		}
	},
	template: `<div>
<input type="hidden" name="healthCheckJSON" :value="JSON.stringify(healthCheck)"/>
<table class="ui table definition selectable">
	<tbody>
		<tr>
			<td class="title">????????????</td>
			<td>
				<div class="ui checkbox">
					<input type="checkbox" value="1" v-model="healthCheck.isOn"/>
					<label></label>
				</div>
			</td>
		</tr>
	</tbody>
	<tbody v-show="healthCheck.isOn">
		<tr>
			<td>URL *</td>
			<td>
				<div class="ui fields inline">
					<div class="ui field">
						<select class="ui dropdown" v-model="urlProtocol">
							<option value="http">http://</option>
							<option value="https">https://</option>
						</select>
					</div>
					<div class="ui field">
						<var style="color:grey">\${host}</var>
					</div>
					<div class="ui field">:</div>
					<div class="ui field">
						<input type="text" maxlength="5" style="width:5.4em" placeholder="??????" v-model="urlPort"/>
					</div>
					<div class="ui field">
						<input type="text" v-model="urlRequestURI" placeholder="/" style="width:23em"/>
					</div>
				</div>
				<div class="ui divider"></div>
				<p class="comment" v-if="healthCheck.url.length > 0">????????????URL???<code-label>{{healthCheck.url}}</code-label>?????????\${host}????????????????????????</p>
			</td>
		</tr>
		<tr>
			<td>??????????????????</td>
			<td>
				<time-duration-box :v-value="healthCheck.interval"></time-duration-box>
			</td>
		</tr>
		<tr>
			<td>??????????????????</td>
			<td>
				<div class="ui checkbox">
					<input type="checkbox" value="1" v-model="healthCheck.autoDown"/>
					<label></label>
				</div>
				<p class="comment">????????????????????????????????????????????????????????????????????????/????????????????????????????????????DNS?????????</p>
			</td>
		</tr>
		<tr v-show="healthCheck.autoDown">
			<td>??????????????????</td>
			<td>
				<input type="text" v-model="healthCheck.countUp" style="width:5em" maxlength="6"/>
				<p class="comment">??????N???????????????????????????????????????</p>
			</td>
		</tr>
		<tr v-show="healthCheck.autoDown">
			<td>??????????????????</td>
			<td>
				<input type="text" v-model="healthCheck.countDown" style="width:5em" maxlength="6"/>
				<p class="comment">??????N?????????????????????????????????</p>
			</td>
		</tr>
	</tbody>	
	<tbody v-show="healthCheck.isOn">
		<tr>
			<td colspan="2"><more-options-angle @change="showAdvanced"></more-options-angle></td>
		</tr>
	</tbody>
	<tbody v-show="advancedVisible && healthCheck.isOn">
		<tr>
			<td>??????????????????</td>
			<td>
				<values-box :values="healthCheck.statusCodes" maxlength="3" @change="changeStatus"></values-box>
			</td>
		</tr>
		<tr>
			<td>????????????</td>
			<td>
				<time-duration-box :v-value="healthCheck.timeout"></time-duration-box>
			</td>	
		</tr>
		<tr>
			<td>??????????????????</td>
			<td>
				<input type="text" v-model="healthCheck.countTries" style="width: 5em" maxlength="2"/>
			</td>
		</tr>
		<tr>
			<td>??????????????????</td>
			<td>
				<time-duration-box :v-value="healthCheck.tryDelay"></time-duration-box>
			</td>
		</tr>
	</tbody>
</table>
<div class="margin"></div>
</div>`
})