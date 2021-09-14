Tea.context(function () {

    this.hanleOpen = function () {
        teaweb.popup(Tea.url("/hids/agents/create"), {height:'23em',width:'50em'});
    };
    this.onDelete = function (agent){

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

});
