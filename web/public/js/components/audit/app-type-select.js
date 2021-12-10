Vue.component("app-type-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-assembly-type","v-assembly-edit"],
	data: function () {
		let assemblyType = this.vAssemblyType
		if (assemblyType == null) {
			assemblyType = -1
		}
		let assemblys = [
			{"id":0,"name":"nginx"},
			{"id":1,"name":"iis"},
			{"id":2,"name":"tomcat"},
		]
		assemblyEdit = this.vAssemblyEdit
		return {
			assemblys: assemblys,
			assemblyType: assemblyType,
			assemblyEdit: assemblyEdit,
		}
	},
	watch:{
		assemblyType(newVal, oldVale) {
            if (newVal !== oldVale) {
                this.$emit("update:vAssemblyType", newVal)
            }
        },
	},
	template: `<div>
	<select name="assemblyType" v-model="assemblyType" :disabled="assemblyEdit" style="width: 250px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option value="-1">请选择</option>
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})