//axios简易封装
function reqApi(method,url,data,param,sucHandle,failHandle,doneHandle){
    let config = {
        method: method,
        url: url,
        data: data,
        param:param,
        header:{
            "X-Requested-With": "XMLHttpRequest"
        }
    }
    axios(config).then(res=>{
        let response = res.data
        if(response.code==200){
            if(typeof(sucHandle) === "function"){
                sucHandle(response)
            }
        }else{
            if(failHandle && typeof(failHandle) === "function"){
                failHandle(response)
            }else{
                teaweb.warn(response.message)
            }
        }
    })
    .catch(error=>{
        console.log(error)
    })
    .then(()=>{
        if(doneHandle && typeof(doneHandle) === "function"){
            doneHandle()
        }
    })
}