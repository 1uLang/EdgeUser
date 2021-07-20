Vue.component("command-level-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-assembly-id"],
	data: function () {
		let assemblyType = this.vAssemblyId
		if (assemblyType == null) {
			assemblyType = -1
		}
		return {
			assemblys: [],
			assemblyType: assemblyType,
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
	<select name="assemblyType" v-model="assemblyType" style="width: 220px;height: 26px;border: 1px solid #d7d7d7;">
		<option value="-1">linux</option>
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})