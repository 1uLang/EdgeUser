Vue.component("assert-system-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-platform"],
	data: function () {
		let platform = this.vPlatform
		if (platform == null) {
			platform = "Linux"
		}
		return {
			assemblys: [],
			platform: platform,
		}
	},
	watch:{
		// assemblyType:function (){
		// 	if (Tea.Vue != null) {
		// 		Tea.Vue.showAPIAUTHVisible = this.assemblyType
		// 	}
		// }
	},
	template: `<div>
	<select name="platform" v-model="platform" style="width: 790px;height: 30px;border: 1px solid #d7d7d7;">
		<option value="Linux">Linux</option>
		<option value="Unix">Unix</option>
		<option value="MacOS">MacOS</option>
		<option value="BSD">BSD</option>
		<option value="Windows">Windows</option>
		<option value="Windows2016">Windows2016</option>
		<option value="Other">Other</option>
	</select>
</div>`
})