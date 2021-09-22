Tea.context(function () {

    this.agentItem = {}
    this.pageState = 1
    this.$delay(function () {

        if (this.event !== "") {
            this.pageState = 2
        }
        let agent = localStorage.getItem("hidsSelectAgentId");
        if (agent) {
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
    this.onChangeEvent = function (event) {
        switch (event) {
            case 'deleted':
                return "文件删除";
            case 'modified':
                return "文件修改";
            default :
                return event;
        }
    }
    this.search = function () {
        localStorage.setItem("hidsSelectAgentId", this.agent);
        window.location = "/hids/syscheck?agent=" + this.agent + "&event=" + this.event
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.lastIndexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };
    this.onChangeSize = function (size) {

        let level = ['Bytes', 'KB', 'MB', 'GB']
        let l = 0
        for (; ;) {
            if (size < 1024 || l > level.length)
                break
            l++
            size /= size
        }
        return size + level[l]
    }
    this.onChangeTimeFormat2 = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.lastIndexOf("Z"));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.onChangeState = function (id) {
        if (this.pageState != id) {
            let url = "/hids/syscheck?agent=" + this.agent
            if (id === 2)
                url += "&event=1"
            window.location = url
        }
    }
    this.onDetails = function (item) {
        window.location = "/hids/syscheck?agent=" + this.agent + "&path=" + item.file + '&event=1'
    }
})