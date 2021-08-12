Vue.component("report-timeloop-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-cycle"],
	data: function () {
		let cycle = this.vCycle
		if (cycle == null) {
			cycle = 1
		}
		return {
			assemblys: [{id:1,name:"每天(日报)"},{id:2,name:"每周(周报)"},{id:3,name:"每月(月报)"}],
			cycle: cycle,

		}
	},
	watch:{
		cycle(newVal, oldVale) {
            if (newVal !== oldVale) {
                this.$emit("update:vCycle", newVal)
            }
        },
	},
	template: `<div>
	<select name="cycle" v-model="cycle" style="width: 320px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})