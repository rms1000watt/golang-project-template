package serve

type Config struct {
	Host string
	Port int

	CertsPath string
	CertName  string
	KeyName   string
}
