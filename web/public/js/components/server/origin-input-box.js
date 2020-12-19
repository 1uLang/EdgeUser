Vue.component("origin-input-box", {
	data: function () {
		return {
			origins: [],
			isAdding: false
		}
	},
	methods: {
		add: function () {
			this.origins.push({
				id: "",
				host: "",
				isPrimary: true,
				isPrimaryValue: 1,
				scheme: "http"
			})
		},
		confirm: function () {
		},
		cancel: function () {
		},
		remove: function (index) {
			this.origins.$remove(index)
		},
		changePrimary: function (origin) {
			origin.isPrimary = origin.isPrimaryValue == 1
		}
	},
	template: `<div>
	<input type="hidden" name="originsJSON" :value="JSON.stringify(origins)"/>
	<div>
		<div class="ui fields inline">
			<div class="ui field" style="padding-left: 0.1em; width:7.5em; color: grey">源站协议</div>
			<div class="ui field" style="width:21em; color: grey">源站地址</div>
			<div class="ui field" style="color: grey">优先级 &nbsp;<tip-icon content="优先级：优先使用主源站，如果主源站无法连接时才会连接备用源站"></tip-icon></div>
		</div>
		<div class="ui divider"></div>
		<div v-for="(origin, index) in origins">
			<div class="ui fields inline" style="margin-top: 0.6em">
				<div class="ui field">
					<select class="ui dropdown auto-width" v-model="origin.scheme">
						<option value="http">http://</option>
						<option value="https">https://</option>
					</select>
				</div>
				<div class="ui field">
					<input type="text" placeholder="源站地址" v-model="origin.host" style="width:20em"/>
				</div>
				<div class="ui field">
					<select class="ui dropdown auto-width small" v-model="origin.isPrimaryValue" @change="changePrimary(origin)">
						<option value="1">主</option>
						<option value="0">备</option>
					</select>
				</div>
				<div class="ui field">
					<a href="" title="删除" @click.prevent="remove(index)"><i class="icon remove icon small"></i></a>
				</div>
			</div>
		</div>
	</div>
	<div style="margin-top: 1em">
		<button class="ui button tiny" type="button" @click.prevent="add()">+</button>
	</div>
</div>`
})