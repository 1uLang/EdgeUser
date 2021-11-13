Vue.component("report-restype-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-assembly-type"],
	data: function () {
		let assemblyType = this.vAssemblyType
		if (assemblyType == null) {
			assemblyType = 0
		}
		return {
			assemblys: [
				{"id":1,"name":"数据库"},
				{"id":2,"name":"主机"},
				{"id":3,"name":"应用"},
			],
			assemblyType: assemblyType,
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
	<select name="assemblyType" v-model="assemblyType" style="width: 100px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
<!--		<option value="0">全部类型</option>-->
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})