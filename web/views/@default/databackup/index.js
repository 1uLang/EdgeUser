Tea.context(function () {

    this.pageState = 1

    this.onChangeState=function (id) {
        if(this.pageState!=id){
            this.pageState = id
        }
    }

})