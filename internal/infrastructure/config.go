package infrastructure

// OCIConfig contém as configurações necessárias para autenticação no OCI
type OCIConfig struct {
	TenancyOCID string
	UserOCID    string
	KeyFile     string
	Fingerprint string
	Region      string
}
