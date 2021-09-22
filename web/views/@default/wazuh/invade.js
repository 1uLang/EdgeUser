Tea.context(function () {

    this.agentItem = {}
    this.$delay(function () {

        let agent = localStorage.getItem("hidsSelectAgentId");
        if(agent){
            this.agent = agent
        }
        if (this.errorMsg && this.errorMsg != "") {
            teaweb.warn(this.errorMsg)
        }
        for (let idx = 0; idx < this.agents.length; idx++) {

            if (this.agent === this.agents[idx].id) {
                this.agentItem = this.agents[idx]
                break
            }
        }

    })
    this.search = function () {
        localStorage.setItem("hidsSelectAgentId", this.agent);
        window.location = "/hids/invades?agent=" + this.agent
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.lastIndexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };
})