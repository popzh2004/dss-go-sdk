package dss

// MqProperties database connection parameters
type MqProperties struct {
	// IP
	IP string `yaml:"ip"`

	// Port
	Port int `yaml:"port"`

	// VHost
	VHost string `yaml:"vhost"`

	// UserName
	UserName string `yaml:"username"`

	// Password
	Password string `yaml:"password"`

	//集群
	Cluster string `yaml:"cluster"`

	//主题
	Topic string `yaml:"topic"`
}
