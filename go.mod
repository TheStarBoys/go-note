module github.com/TheStarBoys/go-note

go 1.12

require (
	github.com/Shopify/sarama v1.26.4
	github.com/astaxie/beego v1.12.2
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/tools v0.0.0-20200117065230-39095c1d176c
	google.golang.org/genproto v0.0.0-20200711021454-869866162049 // indirect
	google.golang.org/grpc v1.30.0 // indirect
	gopkg.in/olivere/elastic.v2 v2.0.61 // indirect
)

replace gopkg.in/fsnotify.v1 => github.com/fsnotify/fsnotify v1.4.9

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
