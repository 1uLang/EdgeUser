Vue.component("report-resname-selector", {
    mounted: function () {
        let that = this

        Tea.action("/audit/report/assets?assetsType="+that.assemblyType)
            .get()
            .success(function (resp) {
                that.assemblys1 = resp.data.dbAssetsList
                that.assemblys2 = resp.data.hostAssetsList
                that.assemblys3 = resp.data.appAssetsList
                console.log(that.assemblys1)
                 that.onResetData()
            })
    },
    props: ["v-assembly-id", "v-assembly-type"],
    data: function () {
        let assemblyId = this.vAssemblyId
        if (assemblyId == null) {
            assemblyId = -1
        }
        let assemblyType = this.vAssemblyType
        if (assemblyType == null) {
            assemblyType = 1
        }
        return {
            assemblys: [],
            assemblys1: [],
            assemblys2: [],
            assemblys3: [],
            assemblyType: assemblyType,
            assemblyId: assemblyId,
        }
    },
    methods:{
      onResetData(){
//        console.log(this.assemblyType)
//        console.log(this.assemblys1)
       this.assemblys = this.assemblys1
       switch(this.assemblyType){
            case 1:
            this.assemblys = this.assemblys1
            break;
            case 2:
            this.assemblys = this.assemblys2
            break;
            case 3:
            this.assemblys = this.assemblys3
            break;
       }
//       console.log(this.assemblys)
      }

     },

    watch: {
        assemblyId(newVal, oldVale) {
            if (newVal !== oldVale) {
                this.$emit("update:vAssemblyId", newVal)
            }
        },
		vAssemblyType(newVal, oldVale) {
           this.assemblyType = newVal
           this.onResetData()
           this.assemblyId = -1
        }
    },
    template: `<div>
	<select name="assemblyId" v-model="assemblyId" style="width: 220px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option value="-1">请选择资产</option>
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})