module github.com/onosproject/onos-config

go 1.16

require (
	github.com/Pallinder/go-randomdata v1.2.0
	github.com/atomix/atomix-go-client v0.6.2
	github.com/atomix/atomix-go-framework v0.10.1
	github.com/atomix/go-client v0.4.1
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/onosproject/config-models/modelplugin/devicesim-1.0.0 v0.8.8
	github.com/onosproject/config-models/modelplugin/testdevice-1.0.0 v0.8.8
	github.com/onosproject/config-models/modelplugin/testdevice-2.0.0 v0.8.8
	github.com/onosproject/helmit v0.6.18
	github.com/onosproject/onos-api/go v0.7.110
	github.com/onosproject/onos-config-model v1.0.1
	github.com/onosproject/onos-lib-go v0.8.0
	github.com/onosproject/onos-test v0.6.5
	github.com/openconfig/gnmi v0.0.0-20200617225440-d2b4e6a45802
	github.com/openconfig/goyang v0.3.1
	github.com/openconfig/ygot v0.12.4
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible
	k8s.io/client-go v0.21.0
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200229013735-71373c6105e3
