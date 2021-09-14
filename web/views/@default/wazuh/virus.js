Tea.context(function () {

    this.agentItem = {}
    this.$delay(function () {

        for (let idx = 0; idx < this.agents.length; idx++) {

            console.log(this.agent, this.agents[idx])
            if (this.agent === this.agents[idx].id) {
                this.agentItem = this.agents[idx]
                console.log(this.agentItem)
                break
            }
        }
    })
    this.search = function () {
        window.location = "/hids/virus?agent=" + this.agent
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