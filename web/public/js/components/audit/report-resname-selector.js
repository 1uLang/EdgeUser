Vue.component("report-resname-selector", {
    mounted: function () {
        let that = this

        Tea.action("/audit/report/assets?assetsType="+that.assemblyType)
            .get()
            .success(function (resp) {
                that.assemblys = resp.data.assetsList
            })
    },
    props: ["v-assembly-id", "v-assembly-type"],
    data: function () {
        let assemblyId = this.vAssemblyId
        if (assemblyId == null) {
            assemblyId = 0
        }
        let assemblyType = this.vAssemblyType
        if (assemblyType == null) {
            assemblyType = 1
        }
        return {
            assemblys: [],
            assemblyType: assemblyType,
            assemblyId: assemblyId,
        }
    },
    watch: {
        assemblyId(newVal, oldVale) {
            if (newVal !== oldVale) {
                this.$emit("update:vAssemblyId", newVal)
            }
        },
		assemblyType(newVal, oldVale) {
			if (newVal !== oldVale) {
				this.$emit("update:vAssemblyType", newVal)
			}
		},
    },
    template: `<div>
	<select name="assemblyId" v-model="assemblyId" style="width: 220px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option value="0">全部资产</option>
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})