module github.com/TeaOSLab/EdgeUser

go 1.15

replace github.com/TeaOSLab/EdgeCommon => ../EdgeCommon

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/TeaOSLab/EdgeCommon v0.0.0-00010101000000-000000000000
	github.com/cespare/xxhash v1.1.0
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/iwind/TeaGo v0.0.0-20210411134150-ddf57e240c2f
	github.com/miekg/dns v1.1.35
	github.com/shirou/gopsutil v3.21.5+incompatible
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	google.golang.org/grpc v1.38.0
)
