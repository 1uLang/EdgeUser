Vue.component("cache-cond-box", {
	data: function () {
		return {
			conds: [],
			addingExt: false,
			addingPath: false,

			extDuration: null,
			pathDuration: null
		}
	},
	methods: {
		addExt: function () {
			this.addingExt = !this.addingExt
			this.addingPath = false

			if (this.addingExt) {
				let that = this
				setTimeout(function () {
					if (that.$refs.extInput != null) {
						that.$refs.extInput.focus()
					}
				})
			}
		},
		changeExtDuration: function (duration) {
			this.extDuration = duration
		},
		confirmExt: function () {
			let value = this.$refs.extInput.value
			if (value.length == 0) {
				return
			}

			let exts = []
			let pieces = value.split(/[,，]/)
			pieces.forEach(function (v) {
				v = v.trim()
				v = v.replace(/\s+/, "")
				if (v.length > 0) {
					if (v[0] != ".") {
						v = "." + v
					}
					exts.push(v)
				}
			})

			this.conds.push({
				type: "url-extension",
				value: JSON.stringify(exts),
				duration: this.extDuration
			})
			this.$refs.extInput.value = ""
			this.cancel()
		},
		addPath: function () {
			this.addingExt = false
			this.addingPath = !this.addingPath

			if (this.addingPath) {
				let that = this
				setTimeout(function () {
					if (that.$refs.pathInput != null) {
						that.$refs.pathInput.focus()
					}
				})
			}
		},
		changePathDuration: function (duration) {
			this.pathDuration = duration
		},
		confirmPath: function () {
			let value = this.$refs.pathInput.value
			if (value.length == 0) {
				return
			}

			if (value[0] != "/") {
				value = "/" + value
			}

			this.conds.push({
				type: "url-prefix",
				value: value,
				duration: this.pathDuration
			})
			this.$refs.pathInput.value = ""
			this.cancel()
		},
		remove: function (index) {
			this.conds.$remove(index)
		},
		cancel: function () {
			this.addingExt = false
			this.addingPath = false
		}
	},
	template: `<div>
	<input type="hidden" name="cacheCondsJSON" :value="JSON.stringify(conds)"/>
	<div v-if="conds.length > 0">
		<div v-for="(cond, index) in conds" class="ui label basic" style="margin-top: 0.2em; margin-bottom: 0.2em">
			<span v-if="cond.type == 'url-extension'">扩展名</span>
			<span v-if="cond.type == 'url-prefix'">路径</span>：{{cond.value}} &nbsp; <span class="grey small">(<time-duration-text :v-value="cond.duration"></time-duration-text>)</span> &nbsp;
			<a href="" title="删除" @click.prevent="remove(index)"><i class="icon remove"></i></a>
		</div>
		<div class="ui divider"></div>
	</div>
	
	<!-- 添加扩展名 -->
	<div v-if="addingExt">
		<div class="ui fields inline">
			<div class="ui field">
				<input type="text" placeholder="扩展名，比如.png, .gif，英文逗号分割" style="width:20em" ref="extInput" @keyup.enter="confirmExt" @keypress.enter.prevent="1"/>
			</div>
			<div class="ui field">
				<time-duration-box placeholder="缓存时长" :v-unit="'day'" :v-count="1" @change="changeExtDuration"></time-duration-box>
			</div>
			<div class="ui field">
				<a href="" class="ui button tiny" @click.prevent="confirmExt">确定</a> &nbsp; <a href="" title="取消" @click.prevent="cancel()"><i class="icon remove small"></i></a>
			</div>
		</div>
	</div>
	
	<!-- 添加路径 -->
	<div v-if="addingPath">
		<div class="ui fields inline">
			<div class="ui field">
				<input type="text" placeholder="路径，以/开头" style="width:20em" ref="pathInput" @keyup.enter="confirmPath" @keypress.enter.prevent="1"/>
			</div>
			<div class="ui field">
				<time-duration-box placeholder="缓存时长" :v-unit="'day'" :v-count="1" @change="changePathDuration"></time-duration-box>
			</div>
			<div class="ui field">
				<a href="" class="ui button tiny" @click.prevent="confirmPath">确定</a> &nbsp; <a href="" title="取消" @click.prevent="cancel()"><i class="icon remove small"></i></a>
			</div>
		</div>
	</div>
	
	<div style="margin-top: 1em">
		<button type="button" class="ui button tiny" @click.prevent="addExt">+缓存扩展名</button> &nbsp; 
		<button type="button" class="ui button tiny" @click.prevent="addPath">+缓存路径</button>
	</div>
</div>`
})