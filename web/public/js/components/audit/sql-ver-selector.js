Vue.component("sql-ver-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-assembly-ver","v-assembly-type","v-assembly-edit"],
	data: function () {
        let assemblyVer = this.vAssemblyVer
        if (assemblyVer == null) {
            assemblyVer = ""
        }
        let assemblyType = this.vAssemblyType
        // console.log(assemblyType)
        if (assemblyType == null) {
            assemblyType = -1
        }
        let isEdit = this.vAssemblyEdit

        let assembly0 = [
            {"name":"5.1"},{"name":"5.2"},{"name":"5.3"},{"name":"5.5"},{"name":"10.0"},
            {"name":"10.1"},{"name":"10.2"}, {"name":"10.3"}, {"name":"10.4"}, {"name":"10.5"},{"name":"10.6"},
        ]
        let assembly1 = [
            {"name":"5.5"},
            {"name":"5.6"},
            {"name":"5.7"},
            {"name":"8.0"},
            {"name":"8.1"},
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
        let assembly3 = [
            {"name":"9.5"},
            {"name":"9.6"},
            {"name":"10"},
            {"name":"11"},
            {"name":"12"},
            {"name":"13"},
            {"name":"14"},
        ]
        let assembly4 = [
            {"name":"2.1"},
            {"name":"2.4"},
            {"name":"3.0"},
            {"name":"4.0"},
            {"name":"4.2"},
        ]
        let assembly5 = [
            {"name":"10G"},
            {"name":"11G"},
            {"name":"12C"},
            {"name":"18C"},
            {"name":"19C"},
        ]
        switch(assemblyType){
            case 0:
                assemblys = assembly0
                break;
            case 1:
                assemblys = assembly1
                break;
            case 2:
                assemblys = assembly2
                break;
            case 3:
                assemblys = assembly3
                break;
            case 4:
                assemblys = assembly4
                break;
            case 5:
                assemblys = assembly5
                break;
            default:
                assemblys = assembly1
        }
		return {
			assemblys: assemblys,
			assemblyVer: assemblyVer,
            isEdit:isEdit,
		}
	},
	methods:{
      onResetData(){
          let assembly0 = [
              {"name":"5.1"},{"name":"5.2"},{"name":"5.3"},{"name":"5.5"},{"name":"10.0"},
              {"name":"10.1"},{"name":"10.2"}, {"name":"10.3"}, {"name":"10.4"}, {"name":"10.5"},{"name":"10.6"},
          ]
          let assembly1 = [
              {"name":"5.5"},
              {"name":"5.6"},
              {"name":"5.7"},
              {"name":"8.0"},
              {"name":"8.1"},
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
          let assembly3 = [
              {"name":"9.5"},
              {"name":"9.6"},
              {"name":"10"},
              {"name":"11"},
              {"name":"12"},
              {"name":"13"},
              {"name":"14"},
          ]
          let assembly4 = [
              {"name":"2.1"},
              {"name":"2.4"},
              {"name":"3.0"},
              {"name":"4.0"},
              {"name":"4.2"},
          ]
          let assembly5 = [
              {"name":"10G"},
              {"name":"11G"},
              {"name":"12C"},
              {"name":"18C"},
              {"name":"19C"},
          ]
          this.assemblys = assembly1

          switch(this.assemblyType){
              case "0":
                  this.assemblys = assembly0
                  break;
              case "1":
                  this.assemblys = assembly1
                  break;
              case "2":
                  this.assemblys = assembly2
                  break;
              case "3":
                  this.assemblys = assembly3
                  break;
              case "4":
                  this.assemblys = assembly4
                  break;
              case "5":
                  this.assemblys = assembly5
                  break;
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
		    if(newVal !== oldVale) {
                this.assemblyType = newVal
                this.onResetData()
            }

          }

	},
	template: `<div>
	<select name="assemblyVer" v-model="assemblyVer" :disabled="isEdit" style="width: 250px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option value="">请选择</option>
		<option v-for="assembly in assemblys" :value="assembly.name">{{assembly.name}}</option>
	</select>
</div>`
})