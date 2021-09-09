package consul

type ServiceInstance interface {
	GetInstanceId() string
	GetServiceId() string
	GetHost() string
	GetPort() int
	IsSecure() bool
	GetMetadata() map[string]string
}

type DefaultServiceInstance struct {
	InstanceId string
	ServiceId  string
	Host       string
	Port       int
	Secure     bool
	Metadata   map[string]string
}

func NewDefaultServiceInstance(serviceId string) {

}
