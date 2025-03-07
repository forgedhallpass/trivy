package types

type DockerOption struct {
	// Auth
	UserName    string
	Password    string
	Credentials []Credential
	// RegistryToken is a bearer token to be sent to a registry
	RegistryToken string

	// ECR
	AwsAccessKey    string
	AwsSecretKey    string
	AwsSessionToken string
	AwsRegion       string

	// GCP
	GcpCredPath string

	// SSL/TLS
	InsecureSkipTLSVerify bool
	NonSSL                bool

	// Architecture
	Platform string
}

type Credential struct {
	UserName string
	Password string
}
