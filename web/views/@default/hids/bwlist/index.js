Tea.context(function () {


    this.$delay(function () {

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })

    this.onAddNameList = function () {
        teaweb.popup(Tea.url(".createPopup"), {
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload();
                });
            },
        });
    }

    this.onDelete = function (item) {
        teaweb.confirm("确定删除该黑白名单吗？", function () {
            this.$post(".del").params({
                Id: item.id,
            }).refresh()
        })
    }

    this.showHost = function () { //重新加载该页面
        window.location.href = '/hids/bwlist'
    }
    this.toShowFlags = function (black) {
        if (black) {
            return "黑名单"
        } else {
            return "白名单"
        }
    }
})