Tea.context(function () {

    this.os = ""
    this.version = ""
    this.architecture = ""
    this.command = ""

    this.osSelectIndex = 0
    this.versionSelectIndex = 0
    this.architectureSelectIndex = 0

    this.OSList = [
        {id: 0, name: "请选择"},
        {id: 1, name: "Red Hat / CentOS"},
        {id: 2, name: "Debian / Ubuntu"},
        {id: 3, name: "Windows"},
        {id: 4, name: "MacOS"},
    ]
    this.versionList = [
        {id: 0, name: "请选择"},
        {id: 1, name: "CentOS5"},
        {id: 2, name: "CentOS6 or higher"},
    ]

    this.architectureList = [
        {id: 0, name: "请选择"},
        {id: 1, name: "i386"},
        {id: 2, name: "x86_64"},
        {id: 3, name: "armhf"},
        {id: 4, name: "aarch64"},
    ]

    this.commandList = [
        "sudo WAZUH_MANAGER='ADDR' WAZUH_AGENT_GROUP='GROUP' yum install URL/UFILE",//centos
        "curl -so FILE URL/UFILE && sudo WAZUH_MANAGER='ADDR' WAZUH_AGENT_GROUP='GROUP' dpkg -i ./FILE",//ubuntu
        "Invoke-WebRequest -Uri URL/UFILE -OutFile FILE; ./FILE /q WAZUH_MANAGER='ADDR' WAZUH_REGISTRATION_SERVER='ADDR' WAZUH_AGENT_GROUP='GROUP' ",//windows
        "curl -so FILE URL/UFILE && sudo launchctl setenv WAZUH_MANAGER 'ADDR' WAZUH_AGENT_GROUP 'GROUP' && sudo installer -pkg ./FILE -target /",//macos
    ]

    this.showArchitecture = function () {
        if (this.version === 'Red Hat / CentOS' || this.os === 'Debian / Ubuntu')
            return true
        return false
    }
    this.changeOS = function () {
        this.version = ''
        this.architecture = ''
        this.command = ''
    }

    //获取命令
    this.onGetItemCommand = function (id, table) {
        if (id && table && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].id == id) {
                    return table[index].command
                }
            }
        }
        return ""
    }

    //获取名称
    this.onGetItemName = function (id, table) {
        if (id && table && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].id == id) {
                    return table[index].name
                }
            }
        }
        return ""
    }

    this.createCommand = function (update) {

        this.command = ""
        let file = ""
        if (update === 1 || this.osSelectIndex == 0 || this.osSelectIndex == 3 || this.osSelectIndex == 4) {
            this.versionSelectIndex = 0
            this.architectureSelectIndex = 0
        } else if (this.osSelectIndex == 2) {
            this.versionSelectIndex = 0
        }

        if (this.osSelectIndex == 3 || this.osSelectIndex == 4) {
            this.command = this.commandList[this.osSelectIndex - 1]
            file = this.commands[this.osSelectIndex]
        } else {
            if (this.osSelectIndex == 1 && this.versionSelectIndex != 0 && this.architectureSelectIndex != 0) {
                this.command = this.commandList[this.osSelectIndex - 1]
                file = this.commands[this.osSelectIndex][this.versionSelectIndex][this.architectureSelectIndex]
            }
            if (this.osSelectIndex == 2 && this.architectureSelectIndex != 0) {
                this.command = this.commandList[this.osSelectIndex - 1]
                file = this.commands[this.osSelectIndex][this.architectureSelectIndex]
            }
        }
        //替换command
        this.command = this.command.replaceAll("ADDR", this.server)
        this.command = this.command.replaceAll("'GROUP'", "'" + this.group + "'")
        let url = window.location
        this.command = this.command.replaceAll("URL", url.origin)

        this.command = this.command.replaceAll("UFILE", "file/" + file)
        this.command = this.command.replaceAll("FILE", file)

    }
})