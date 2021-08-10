Tea.context(function () {

    this.onOpenCommand = function () {
        teaweb.popup(Tea.url(".install"), {
            height: "350px",
        });
    }
    this.onOpenCreate = function () {
        teaweb.popup(Tea.url(".create"), {
            height: "180px",
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload();
                });
            },
        });

    }

    this.$delay(function () {

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }

        let that = this
        that.onCreateLoopTimeOut()
        window.addEventListener('beforeunload', function () {
            that.onReleaseTimeOut()
        })
    })

    this.onCallBack = function () {
        if (this.checkScans()) {
            this.$post(".").success(resp => {
                if (resp.code === 200) {
                    this.agents = resp.data.agents
                }
            })
        }
    }
    this.onCreateLoopTimeOut = function () {
        this.onReleaseTimeOut()
        this.checkTimer = createTimer(this.onCallBack, {timeout: 60000});
        this.checkTimer.start();
    }
    this.onReleaseTimeOut = function () {
        if (this.checkTimer) {
            this.checkTimer.stop()
            this.checkTimer = null
        }
    }
    this.checkScans = function () {
        for (item of this.agents) {
            if (item.agentState == '1' || item.agentState == '3' || item.agentState == '5') {
                return true
            }
        }
        return false
    }
    this.onStartConfig = function (item) {
        teaweb.confirm("确定要启动吗？", function () {
            this.$post(".disport")
                .params({
                    MacCode: item.macCode,
                    Opt: 'enable',
                }).success(function () {
                window.location.reload()
            })
        })

    }

    this.onStopConfig = function (item) {
        teaweb.confirm("确定要停用吗？", function () {
            this.$post(".disport")
                .params({
                    MacCode: item.macCode,
                    Opt: 'disable',
                }).success(function () {
                window.location.reload()
            })
        })

    }

    this.onDelete = function (item) {
        teaweb.confirm("确定要删除吗？", function () {
            this.$post(".delete")
                .params({
                    id: item.id,
                }).success(function () {
                window.location.reload()
            })
        })

    }
    this.onUninstall = function (item) {
        teaweb.confirm("确定要卸载吗？", function () {
            this.$post(".disport")
                .params({
                    MacCode: item.macCode,
                    Opt: 'delete',
                }).success(function () {
                window.location.reload()
            })
        })
    }

    this.getStateName = function (state) {
        if (state == '1')
            return "启用中"
        else if (state == '2')
            return "已启用"
        else if (state == '3')
            return "停用中"
        else if (state == '4')
            return "已停用"
        else if (state == '5')
            return "卸载中"
        else if (state == '6')
            return "已卸载"
    }
});
  