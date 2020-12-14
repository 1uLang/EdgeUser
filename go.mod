module github.com/TeaOSLab/EdgeUser

go 1.15

replace github.com/TeaOSLab/EdgeCommon => ../EdgeCommon

require (
	github.com/TeaOSLab/EdgeCommon v0.0.0-00010101000000-000000000000
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/iwind/TeaGo v0.0.0-20201209122854-4c8b1780a42b
	google.golang.org/grpc v1.32.0
)
