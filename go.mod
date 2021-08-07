module github.com/TeaOSLab/EdgeUser

go 1.15

replace github.com/1uLang/zhiannet-api => ../zhiannet-api

replace github.com/TeaOSLab/EdgeCommon => ../EdgeCommon

require (
	github.com/1uLang/zhiannet-api v0.0.0-00010101000000-000000000000
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/TeaOSLab/EdgeCommon v0.0.0-00010101000000-000000000000
	github.com/cespare/xxhash v1.1.0
	github.com/dlclark/regexp2 v1.4.0
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/iwind/TeaGo v0.0.0-20210720011303-fc255c995afa
	github.com/miekg/dns v1.1.35
	github.com/shirou/gopsutil v3.21.5+incompatible
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/tidwall/gjson v1.8.0
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	github.com/xlzd/gotp v0.0.0-20181030022105-c8557ba2c119
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	google.golang.org/grpc v1.38.0
)
