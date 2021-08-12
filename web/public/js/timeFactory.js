function createTimer(callBack, options)
{
    /* external settable */
    var timeout = 1000;
    var timeCallBack = function(){
        console.log(timerId.toString() + "-please add custom fnAlarm")
    };
    
    /* inner maintain variable */
    var timerId;
    
    /* set external variable */
    if ( callBack )
    {
        timeCallBack = callBack;
    }

    if ( options && options.timeout )
    {
        timeout = options.timeout;
    }
    
    return {
        start: function(){
            timerId = setInterval(timeCallBack,timeout);
        },
        stop: function(){
            clearInterval(timerId)
            timerId = null
        },
    };
}