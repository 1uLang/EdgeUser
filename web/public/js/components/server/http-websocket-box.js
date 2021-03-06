Vue.component("http-websocket-box", {
	props: ["v-websocket-ref", "v-websocket-config", "v-is-location"],
	data: function () {
		let websocketRef = this.vWebsocketRef
		if (websocketRef == null) {
			websocketRef = {
				isPrior: false,
				isOn: false,
				websocketId: 0
			}
		}

		let websocketConfig = this.vWebsocketConfig
		if (websocketConfig == null) {
			websocketConfig = {
				id: 0,
				isOn: false,
				handshakeTimeout: {
					count: 30,
					unit: "second"
				},
				allowAllOrigins: true,
				allowedOrigins: [],
				requestSameOrigin: true,
				requestOrigin: ""
			}
		} else {
			if (websocketConfig.handshakeTimeout == null) {
				websocketConfig.handshakeTimeout = {
					count: 30,
					unit: "second",
				}
			}
			if (websocketConfig.allowedOrigins == null) {
				websocketConfig.allowedOrigins = []
			}
		}

		return {
			websocketRef: websocketRef,
			websocketConfig: websocketConfig,
			handshakeTimeoutCountString: websocketConfig.handshakeTimeout.count.toString(),
			advancedVisible: false
		}
	},
	watch: {
		handshakeTimeoutCountString: function (v) {
			let count = parseInt(v)
			if (!isNaN(count) && count >= 0) {
				this.websocketConfig.handshakeTimeout.count = count
			} else {
				this.websocketConfig.handshakeTimeout.count = 0
			}
		}
	},
	methods: {
		isOn: function () {
			return (!this.vIsLocation || this.websocketRef.isPrior) && this.websocketRef.isOn
		},
		changeAdvancedVisible: function (v) {
			this.advancedVisible = v
		},
		createOrigin: function () {
			let that = this
			teaweb.popup("/servers/server/settings/websocket/createOrigin", {
				height: "12.5em",
				callback: function (resp) {
					that.websocketConfig.allowedOrigins.push(resp.data.origin)
				}
			})
		},
		removeOrigin: function (index) {
			this.websocketConfig.allowedOrigins.$remove(index)
		}
	},
	template: `<div>
	<input type="hidden" name="websocketRefJSON" :value="JSON.stringify(websocketRef)"/>
	<input type="hidden" name="websocketJSON" :value="JSON.stringify(websocketConfig)"/>
	<table class="ui table definition selectable">
		<prior-checkbox :v-config="websocketRef" v-if="vIsLocation"></prior-checkbox>
		<tbody v-show="(!this.vIsLocation || this.websocketRef.isPrior)">
			<tr>
				<td class="title">??????????????????</td>
				<td>
					<div class="ui checkbox">
						<input type="checkbox" v-model="websocketRef.isOn"/>
						<label></label>
					</div>
				</td>
			</tr>
		</tbody>
		<tbody v-show="isOn()">
			<tr>
				<td class="color-border">?????????????????????<em>???Origin???</em></td>
				<td>
					<div class="ui checkbox">
						<input type="checkbox" v-model="websocketConfig.allowAllOrigins"/>
						<label></label>
					</div>
					<p class="comment">???????????????????????????????????????</p>
				</td>
			</tr>
		</tbody>
		<tbody v-show="isOn() && !websocketConfig.allowAllOrigins">
			<tr>
				<td class="color-border">????????????????????????<em>???Origin???</em></td>
				<td>
					<div v-if="websocketConfig.allowedOrigins.length > 0">
						<div class="ui label tiny" v-for="(origin, index) in websocketConfig.allowedOrigins">
							{{origin}} <a href="" title="??????" @click.prevent="removeOrigin(index)"><i class="icon remove"></i></a>
						</div>
						<div class="ui divider"></div>
					</div>
					<button style="background-color: #1b6aff;" class="ui button tiny" type="button" @click.prevent="createOrigin()">+</button>
					<p class="comment">??????????????????????????????????????????Websocket?????????</p>
				</td>
			</tr>
		</tbody>
		<more-options-tbody @change="changeAdvancedVisible" v-show="isOn()"></more-options-tbody>
		<tbody v-show="isOn() && advancedVisible">
			<tr>
				<td class="color-border">???????????????????????????</td>
				<td>
					<div class="ui checkbox">
						<input type="checkbox" v-model="websocketConfig.requestSameOrigin"/>
						<label></label>
					</div>
					<p class="comment">???????????????????????????????????????<span class="ui label tiny">Origin</span>????????????????????????</p>
				</td>
			</tr>
		</tbody>
		<tbody v-show="isOn() && advancedVisible && !websocketConfig.requestSameOrigin">
			<tr>
				<td class="color-border">????????????????????????</td>
				<td>
					<input type="text" v-model="websocketConfig.requestOrigin" maxlength="200"/>
					<p class="comment">????????????????????????<span class="ui label tiny">Origin</span>????????????</p>
				</td>
			</tr>
		</tbody>
		<!-- TODO ???????????????????????? -->
		<tbody v-show="isOn() && false">
			<tr>
				<td>??????????????????<em>???Handshake???</em></td>
				<td>
					<div class="ui fields inline">
						<div class="ui field">
							<input type="text" maxlength="10" v-model="handshakeTimeoutCountString" style="width:6em"/>
						</div>
						<div class="ui field">
							???
						</div>
					</div>
					<p class="comment">0????????????????????????????????????</p>
				</td>
			</tr>
		</tbody>
	</table>
	<div class="margin"></div>
</div>`
})