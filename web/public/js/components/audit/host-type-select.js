Vue.component("host-type-selector", {
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
		// assemblyType(newVal, oldVale) {
        //     if (newVal !== oldVale) {
        //         this.$emit("update:vAssemblyId", newVal)
        //     }
        // },
	},
	template: `<div>
	<select name="assemblyType" v-model="assemblyType" style="width: 250px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option value="-1">请选择</option>
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})