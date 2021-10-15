Tea.context(function () {

    this.getTime = function (time) {
        var d = new Date(time);
        return d.toLocaleDateString() + " " + d.toLocaleTimeString()
    }
})