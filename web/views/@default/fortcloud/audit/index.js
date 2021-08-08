Tea.context(function () {


    this.pageState = 1

    this.onChangeState = function (state) {
        if (this.pageState != state) {
            this.pageState = state
        }
    }

    this.getTimeLong = function (start, end) {
        //格式： 2021-07-27 17:50:16
        if (start == null || end == null) {
            return ""
        }
        let st = new Date(start)
        let et = new Date(end)

        if (et.getTime() === st.getTime()) {
            return ""
        } else {
            let sec = (et.getTime() - st.getTime()) / 1000
            if (sec > 60) {
                let m = (sec / 60)
                if (m > 60) {
                    return (m / 60).toFixed(1) + "时"
                } else {
                    return m.toFixed(1) + "分"
                }
            } else {
                return sec + ".0秒"
            }
        }
    }
    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            resultTime = time.substring(0, time.indexOf(" +"));
        }
        return resultTime;
    };


    //监控
    this.onDelete = function (id) {

        teaweb.confirm("确定要删除该会话吗？", function () {
            this.$post(".delete")
                .params({
                    Id: id
                })
                .refresh()
        })
    }

    //回放
    this.onReplay = function (item) {
        teaweb.confirm("确定要回放该会话吗？", function () {
            this.onTestReplay("http://192.168.137.8:8002/fortcloud/audit/repaly?id="+item.id)
        })

    }

    this.bShowAudioPlayBox = false
    this.onTestReplay = function (url) {
        this.bShowAudioPlayBox = true
        var RECORDING_URL = url;
        var display = document.getElementById('display');
        var tunnel = new Guacamole.StaticHTTPTunnel(RECORDING_URL);
        var recording = new Guacamole.SessionRecording(tunnel);
        var recordingDisplay = recording.getDisplay();
        display.appendChild(recordingDisplay.getElement());
        recording.connect();
        recording.onplay = () => {
            console.log("onPlayHandle")
        };
        recording.onpause = () => {
            console.log("onPauseHandle")
        };

    }
})