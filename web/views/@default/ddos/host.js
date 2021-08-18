Tea.context(function () {

    this.$delay(function () {
        let curSelectNode = localStorage.getItem("ddosSelectNodeId");
		if(curSelectNode){
			this.nodeId = curSelectNode
		}

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })

    this.level = "1"//防护策略等级
    this.ignore = false //ip直通
    this.dropNode = ''
    this.src_ip = ""
    this.shieldList = []
    this.linkList = []
    this.nShowState = 1 //当前显示的页面

    this.searchAddress = ''//当前策略配置 地址

    this.tableState = 1 //只有在nShowState==2时生效 屏蔽和连接列表的切换

    this.onAddNodeIP = function () {
        let node = this.getNodeId()
        teaweb.popup(Tea.url(".createPopup?nodeId=" + node),
            {
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload();
                });
            },
                height: "23em"
        });
    }
    this.onUpdateConfig = function (item){
        console.log(item)
        teaweb.popup(Tea.url(".updatePopup?nodeId=" + item.node_id+"&addr="+item.addr+"&remark="+item.remark+"&hostId="+item.host_id),
            {
                callback: function () {
                    teaweb.success("修改成功", function () {
                        teaweb.reload();
                    });
                },
                height: "23em"
            });
    }
    this.showHost = function () { //重新加载该页面
        let node = ''
        if (this.nodeId === '') {    //重新加载该页面
            node = document.getElementById('selectBox').value
            this.nodeId = node
        } else {
            node = this.nodeId
        }
        localStorage.setItem("ddosSelectNodeId", node);
        window.location.href = '/ddos/host?nodeId=' + node
    }
    this.setHost = function (notice) {
        if (notice !== true && notice !== false)
            return;
        let ignore = document.getElementById('btn-switch-ignore').checked
        let that = this
        let msg = !this.ignore ? '开启' : '关闭'
        let node = this.nodeId

        teaweb.confirm("确定要" + msg + "IP直通吗？", function () {
            that.$post(".set").params({
                Addr: that.searchAddress,
                ignore: ignore,
                NodeId: node,
                set: that.level,
            }).success(resp => {
                if (resp.code === 200)
                    document.getElementById("btn-switch-ignore").checked = ignore
            })
        })

    }
    //分页切换
    this.changeState = function (state) {
        if (this.nShowState != state) {
            this.nShowState = state
        }
    }
    this.getNodeId = function () {
        let node = this.nodeId
        return node
    }
    this.shieldSearchList = function (state) {
        let searchIp = this.src_ip === '' ? this.searchAddress : this.searchAddress + '-' + this.src_ip
        let node = this.getNodeId()
        //屏蔽列表
        this.$get(".shield").params({Addr: searchIp, NodeId: node,_:new Date().getMilliseconds()}).success(resp => {
            if (resp.code === 200) {
                if (resp.data.shield)
                    this.shieldList = resp.data.shield
                else
                    this.shieldList = []
                this.tableState = state
            }
        })

    };
    this.onOpenConfig = function (addr) {
        this.searchAddress = addr

        let node = this.getNodeId()
        let that = this
        //ip直通 防护策略
        this.$get(".set").params({Addr: addr, NodeId: node}).success(resp => {
            if (resp.code === 200) {
                that.ignore = resp.data.ignore
                that.level = resp.data.level
                if (parseInt(that.level) > 3)
                    that.level = "3"
                //屏蔽列表
                that.$get(".shield").params({Addr: addr, NodeId: node}).success(resp => {
                    if (resp.code === 200) {
                        if (resp.data.shield)
                            that.shieldList = resp.data.shield
                        else
                            that.shieldList = []
                        that.changeState(2)
                    }
                })
            }
        })
    }
    
    //删除
    this.onDeleteConfig = function (host_id) {

        teaweb.confirm("确定删除该高防ip吗？", function () {
            this.$post(".delete").params({
                hostIds: [host_id],
            }).refresh()
        })
    }

    //配置里面的列表切换
    this.changeListState = function (state) {
        if (this.tableState !== state) {
            if (state === 1) { //屏蔽列表
                this.shieldSearchList(state)
            } else {//连接列表
                this.linkSearchList(state)
            }
        }
    }

    //连接列表查询
    this.linkSearchList = function (state) {
        let searchIp = this.src_ip === '' ? this.searchAddress : this.searchAddress + '-' + this.src_ip
        let node = this.getNodeId()
        //屏蔽列表
        this.$get(".link").params({Addr: searchIp, NodeId: node}).success(resp => {
            if (resp.code === 200) {
                this.linkList = resp.data.list
                this.tableState = state
            }
        })
    }
    //导出
    this.onExport = function () {
        var inputValue = ""
        if (this.tableState == 1) {
            inputValue = document.getElementById("linkHostInput").value
        } else {
            inputValue = document.getElementById("unlinkHostInput").value
        }
        //todo
    }
    //全部释放 如果传入id 则单独释放 否则释放全部
    this.onRelease = function (item) {
        let adds = []
        let node = this.getNodeId()
        let msg = ''
        if (item === "all") {//全部释放
            adds[0] = this.searchAddress
            msg = "全部"
        } else {
            adds[0] = item.LocalAddress + '-' + item.RemoteAddress
        }
        let that = this
        //ip直通 防护策略
        teaweb.confirm("确定要" + msg + "释放吗？", function () {
            this.$post(".shield").params({Addr: adds, NodeId: node}).success(resp => {
                if (resp.code === 200) {
                    teaweb.success("释放成功", function () {
                        that.$delay(function () {
                            setTimeout(function() {
                                that.shieldSearchList(1)
                            }, 1000);  //1秒后将会调用执行remind()函数

                        },1000)
                    })
                }
            })
        })
    }

    //策略切换回调
    this.onChangeHandle = function () {
        let curLevel = document.getElementById('ddosLevel').value
        let addr = this.searchAddress
        let node = this.nodeId
        let ignore = this.ignore
        if (this.level === curLevel)
            return

        document.getElementById('ddosLevel').value = this.level
        teaweb.confirm("确定更改防护策略？", function () {
            this.$post(".set")
                .params({
                    Addr: addr,
                    ignore: ignore,
                    set: curLevel,
                    NodeId: node,
                }).success(resp => {
                if (resp.code === 200) {
                    this.level = curLevel
                }
            })
        })

    }
})