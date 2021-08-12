Vue.component("assert-proto-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-proto"],
	data: function () {
		let proto = this.vProto
		if (proto == null) {
			proto = "ssh"
		}
		return {
			proto: proto,
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
	<select name="assemblyType" v-model="proto" style="width: 110px;height: 30px;border: 1px solid #d7d7d7;">
		<option value="ssh">ssh</option>
		<option value="rdp">rdp</option>
		<option value="telnet">telnet</option>
		<option value="vnc">vnc</option>
	</select>
</div>`
})