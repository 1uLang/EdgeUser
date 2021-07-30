Vue.component("sql-ver-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-assembly-ver","v-assembly-type"],
	data: function () {
		let assemblyVer = this.vAssemblyVer
		if (assemblyVer == null) {
			assemblyVer = ""
		}
		let assemblyType = this.vAssemblyType
		console.log(assemblyType)
        if (assemblyType == null) {
            assemblyType = -1
        }
		let assemblys = [
			{"name":"5.5"},{"name":"5.6"},
			{"name":"5.7"},
			{"name":"8.0"},
		]
		return {
			assemblys: assemblys,
			assemblyVer: assemblyVer,
		}
	},
	methods:{
      onResetData(){
       let assembly1 = [
        {"name":"5.5"},
        {"name":"5.6"},
        {"name":"5.7"},
        {"name":"8.0"},
       ]
       let assembly2 = [
        {"name":"2005"},
        {"name":"2008"},
        {"name":"2012"},
        {"name":"2014"},
        {"name":"2016"},
        {"name":"2017"},
        {"name":"2019"},
       ]
       this.assemblys = assembly1
       if(this.assemblyType == 2){
        this.assemblys = assembly2
       }
       this.assemblyVer = ""
      }
    },
	watch:{
		assemblyVer(newVal, oldVale) {
            if (newVal !== oldVale) {
                this.$emit("update:vAssemblyVer", newVal)
            }
        },
        vAssemblyType(newVal, oldVale) {
           this.assemblyType = newVal
           this.onResetData()
          }

	},
	template: `<div>
	<select name="assemblyVer" v-model="assemblyVer" style="width: 250px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option value="">请选择</option>
		<option v-for="assembly in assemblys" :value="assembly.name">{{assembly.name}}</option>
	</select>
</div>`
})