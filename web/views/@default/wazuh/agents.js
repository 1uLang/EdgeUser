Tea.context(function () {

    this.hanleOpen = function () {
        teaweb.popup(Tea.url("/hids/agents/create"), {height: '23em', width: '50em'});
    };
    this.onDelete = function (agent) {

        teaweb.confirm("确定要删除所选资产吗？", function () {
            this.$post("/hids/agents/delete")
                .params({
                    agent: agent,
                }).success(function () {
                window.location.reload()
            })
        })

    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.lastIndexOf("Z"));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.dbClickId = -1
    this.newremake = ""
    this.newremake2 = ""
    this.onDoubleClick = function (item) {
        if (this.dbClickId === item.id)
            return
        if (this.dbClickId != -1) {    //保存
            this.onSaveRemark()
        }
        this.dbClickId = item.id
        this.newremake = item.remake
        this.newremake2 = item.remake
    }

    this.onCheckKeyDown = function (frm, event) {

        var event = window.event ? window.event : event;
        if (event.keyCode == 13) {
            this.onSaveRemark();
        }
    }

    this.onSaveRemark = function () {

        if(this.newremake !== "" && this.newremake !==  this.newremake2){
            this.$post("/hids/agents/update")
                .params({
                    agent: this.dbClickId,
                    remake: this.newremake,
                }).success(function () {
                window.location.reload()
            })
        }

        this.dbClickId = -1
    }

    this.onCheck = function (agent) {

        teaweb.confirm("确定要立即体检所选资产吗？", function () {
            this.$post("/hids/agents/check")
                .params({
                    agent: agent,
                }).success(function () {
                window.location.reload()
            })
        })
    }
});
