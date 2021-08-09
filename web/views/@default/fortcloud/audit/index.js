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

    this.$delay(function () {
        console.log(window.location.origin)
        console.log(window.location.hostname)
        console.log(window.location.host)
        console.log(window.location.port)
    })
    //回放
    this.bShowAudioPlayBox = false
    this.bAudioDisabled = true
    this.selectReplayItemData = null
    this.curProgress = 0
    this.percentDuration = 0
    this.maxDuration = 0
    this.playDuration = 0
    this.playPosition = 0
    this.isPlaying = false
    this.recording = null

    this.getMaxDuration = function (start, end) {
        //格式： 2021-07-27 17:50:16
        if (start == null || end == null) {
            return ""
        }
        let st = new Date(start)
        let et = new Date(end)
        return et.getTime() - st.getTime()
    }
    this.zeroPad = function (num, minLength) {

        // Convert provided number to string
        var str = num.toString();

        // Add leading zeroes until string is long enough
        while (str.length < minLength)
            str = '0' + str;

        return str;

    };
    this.formatTime = function(millis) {

        // Calculate total number of whole seconds
        var totalSeconds = Math.floor(millis / 1000);

        // Split into seconds and minutes
        var seconds = totalSeconds % 60;
        var minutes = Math.floor(totalSeconds / 60);

        // Format seconds and minutes as MM:SS
        return this.zeroPad(minutes, 2) + ':' + this.zeroPad(seconds, 2);

    };

    this.onReplay = function (item) {
        teaweb.confirm("确定要回放该会话吗？", function () {
            this.bShowAudioPlayBox = true
            this.selectReplayItemData = item
            this.maxDuration = this.getMaxDuration(item.connectedTime,item.disconnectedTime)
            this.playDuration = this.formatTime(this.maxDuration)
            this.percentDuration = 0
            this.bAudioDisabled = true
        })
    }

    this.onCloseReplay = function () { 
        this.bShowAudioPlayBox = false
        this.recording = null
    }

    this.onPlayReplay = function () { 
        if(this.bAudioDisabled && !this.recording){
            let path = window.location.origin
            this.onLoadReplay(path + "/fortcloud/audit/replay?id="+this.selectReplayItemData.id)
        }else{
            if(this.recording){
                if(this.percentDuration>=this.maxDuration){
                    this.percentDuration = 0
                    this.recording.seek(0, () => {
                        this.recording.play();
                    });
                }else{
                    
                    this.recording.play();
                }
            }
            
        }
       
    }

    this.onPauseReplay = function(){
        if(this.recording){
            this.recording.pause();
        }
    }

    this.handleSliderChange = function (){
        if (this.recording) {
            // Request seek
            this.recording.seek(this.percentDuration, () => {
                console.log('complete');
            });
        }

    }


    this.onLoadReplay = function (url) {
        var RECORDING_URL = url;
        var display = document.getElementById('audio-display-box');
        var tunnel = new Guacamole.StaticHTTPTunnel(RECORDING_URL);
        this.recording = new Guacamole.SessionRecording(tunnel);
        var recordingDisplay = this.recording.getDisplay();
        display.appendChild(recordingDisplay.getElement());
        this.recording.connect();

        recordingDisplay.onresize = function (width, height) {
            console.log(width)
            // Do not scale if display has no width
            if (!width)
                return;

            // Scale display to fit width of container
            recordingDisplay.scale(display.offsetWidth / width,display.offsetHeight / height);
        };
        this.recording.onplay = () => {
            console.log("onPlayHandle")
            this.isPlaying = true
        };
        this.recording.onpause = () => {
            console.log("onPauseHandle")
            this.isPlaying = false
        };
        this.recording.onseek = (millis) => {
            console.log("onseek")
            this.playPosition=this.formatTime(millis)
            this.percentDuration = millis
        };

        this.recording.onprogress = (millis) => {
            if(millis>=this.maxDuration){
                this.recording.play()
                this.bAudioDisabled = false
            }
        };

    }
})