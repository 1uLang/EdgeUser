Tea.context(function () {
	this.sSelectValue = [1,2,3,4]
	this.bInputNumCount = "8"
	this.bInputOverDay = "90"
	this.bInputHistoryCount = ""
	this.bInputFailCount = "5"
	this.bInputFailWaitSec = "30"
	this.bInputOutWaitSec = "30"

	let that = this

	this.onCheckSelectValue = function (id) {
		for(var index = 0;index<that.sSelectValue.length;index++){
			if(that.sSelectValue[index] == id){
				return true
			}
		}
        return false
	}

	this.onAddConfigValue = function (id) {
        if(that.onCheckSelectValue(id)){
            that.sSelectValue = that.sSelectValue.filter((itemIndex) => {
                return itemIndex != id;
            });
        }else{
            that.sSelectValue.push(id);
        }
    }

	this.getShowSelectValueImage = function (id) {
        let bValue = that.onCheckSelectValue(id);

        if (bValue) {
            // return "/images/select_select.png";
			return "/images/select_select_grey.png";
        }
        return "/images/select_box.png";
    }

	this.passwordConfigData = [
		{id:1,value:"必须包含大写字符"},
		{id:2,value:"必须包含小写字符"},
		{id:3,value:"必须包含数字"},
		{id:4,value:"必须包含特殊字符"},
	]
})