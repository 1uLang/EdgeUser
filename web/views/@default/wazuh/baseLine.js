Tea.context(function () {
    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.lastIndexOf("Z"));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };
    
    this.onOpenProblemDetail = function(item){
        window.location = "/hids/baseLineDetails?agent="+item.agent_id+"&policy="+item.policy_id
    }
})