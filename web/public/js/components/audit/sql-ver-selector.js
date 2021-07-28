Vue.component("sql-ver-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-assembly-ver"],
	data: function () {
		let assemblyVer = this.vAssemblyVer
		if (assemblyVer == null) {
			assemblyVer = -1
		}
		let assemblys = [
			{"name":"5.5"},
			{"name":"5.7"},
			{"name":"8.0"},
		]
		return {
			assemblys: assemblys,
			assemblyVer: assemblyVer,
		}
	},
	watch:{
		assemblyVer(newVal, oldVale) {
            if (newVal !== oldVale) {
                this.$emit("update:vAssemblyVer", newVal)
            }
        },
	},
	template: `<div>
	<select name="assemblyVer" v-model="assemblyVer" style="width: 250px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option value="-1">请选择</option>
		<option v-for="assembly in assemblys" :value="assembly.name">{{assembly.name}}</option>
	</select>
</div>`
})