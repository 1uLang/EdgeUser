Tea.context(function () {

    this.links = []
    this.attacks = []

    this.address = ''
    this.nShowState = 1
    this.endTime = ""
    this.startTime = ""
    this.attackType = "0"
    this.status = "0"

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")

        let curSelectNode = localStorage.getItem("ddosSelectNodeId");
		if(curSelectNode){
			this.nodeId = curSelectNode
		}

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
        // this.getTraffic(this.nShowState)
    })
    this.getAttacks = function (state) {

        this.$get(".attacks").params({NodeId: this.nodeId, level: this.level}).success(resp => {
            if (resp.code === 200) {
                if (resp.data.attacks)
                    this.attacks = resp.data.attacks
                else
                    this.attacks = []
                this.nodeId = resp.data.nodeId
                this.nShowState = state
            }
        })
    }
    this.showHost = function () { //重新加载该页面
        let node = this.nodeId
        localStorage.setItem("ddosSelectNodeId", node);
        window.location.href = '/ddos/logs?nodeId=' + node
    }

    this.onChangeShowState = function (state) {
        this.level = 1
        if (this.nShowState != state) {
            if (state === 2) {
                this.getAttacks(state)
            } else {
                this.$get(".link").params({NodeId: this.nodeId, level: this.level}).success(resp => {
                    if (resp.code === 200) {
                        if (resp.data.links)
                            this.links = resp.data.links
                        else
                            this.links = []
                        this.level = resp.data.level
                        this.nShowState = state
                    }
                })
            }

        }
    }
    this.search = function () {
        if (this.nShowState == 1) {
            window.location.href = "/ddos/logs?NodeId=" + this.nodeId + "&Level=" + this.level
        } else if (this.nShowState == 3) {
            this.$get(".link").params({NodeId: this.nodeId , Level: this.level}).success(resp => {
                if (resp.code === 200) {
                    if (resp.data.links)
                        this.links = resp.data.links
                    else
                        this.links = []
                    this.level = resp.data.level
                }
            })
        } else if (this.nShowState == 2) {
            let start = document.getElementById("day-from-picker").value
            let end = document.getElementById("day-to-picker").value
            this.$get(".attacks").params({
                NodeId: this.nodeId,
                startTime: start,
                endTime: end,
                attackType: this.attackType,
                status: this.status,
                address: this.address,
            }).success(resp => {
                if (resp.code === 200) {
                    if (resp.data.attacks)
                        this.attacks = resp.data.attacks
                    else
                        this.attacks = []
                    this.nodeId = resp.data.nodeId
                    this.startTime = resp.data.startTime
                    this.endTime = resp.data.endTime
                    this.address = resp.data.address
                    this.attackType = resp.data.attackType
                    this.status = resp.data.status
                }
            })
        }
    }
    this.toShowStatus = function (st) {
        if (st === "2") {
            return "结束"
        } else {
            return "保护中"
        }
    };

    this.onDownLoad = function (id) {

    }

})