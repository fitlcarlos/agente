package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// OCIConfig contém as configurações necessárias para autenticação no OCI
type OCIConfig struct {
	TenancyOCID string
	UserOCID    string
	KeyFile     string
	Fingerprint string
	Region      string
}

// LoadConfig carrega a configuração do arquivo .env
func LoadConfig() OCIConfig {
	// Verificar se o arquivo .env existe antes de tentar carregá-lo
	envFile := ".env"

	if _, err := os.Stat(envFile); err != nil {
		if os.IsNotExist(err) {
			log.Printf("⚠️  Arquivo %s não encontrado. Tentando carregar variáveis do ambiente do sistema...", envFile)
		} else {
			log.Printf("⚠️  Erro ao verificar arquivo %s: %v. Tentando carregar variáveis do ambiente do sistema...", envFile, err)
		}
	} else {
		// Arquivo existe, tentar carregar
		if err := godotenv.Load(envFile); err != nil {
			log.Printf("⚠️  Erro ao carregar arquivo %s: %v. Tentando carregar variáveis do ambiente do sistema...", envFile, err)
		} else {
			log.Printf("✅ Arquivo %s carregado com sucesso", envFile)
		}
	}

	cfg := OCIConfig{
		TenancyOCID: os.Getenv("OCI_TENANCY_ID"),
		UserOCID:    os.Getenv("OCI_USER_ID"),
		KeyFile:     os.Getenv("OCI_KEY_FILE"),
		Fingerprint: os.Getenv("OCI_FINGERPRINT"),
		Region:      os.Getenv("OCI_REGION"),
	}

	// Validar se todas as configurações necessárias estão presentes
	if err := cfg.Validate(); err != nil {
		log.Fatalf("❌ Erro na configuração: %v", err)
	}

	return cfg
}

// Validate verifica se todas as configurações obrigatórias estão presentes
func (c *OCIConfig) Validate() error {
	if c.TenancyOCID == "" {
		return fmt.Errorf("OCI_TENANCY_ID não encontrado. Verifique se a variável está definida no arquivo .env ou no ambiente do sistema")
	}
	if c.UserOCID == "" {
		return fmt.Errorf("OCI_USER_ID não encontrado. Verifique se a variável está definida no arquivo .env ou no ambiente do sistema")
	}
	if c.KeyFile == "" {
		return fmt.Errorf("OCI_KEY_FILE não encontrado. Verifique se a variável está definida no arquivo .env ou no ambiente do sistema")
	}
	if c.Fingerprint == "" {
		return fmt.Errorf("OCI_FINGERPRINT não encontrado. Verifique se a variável está definida no arquivo .env ou no ambiente do sistema")
	}
	if c.Region == "" {
		return fmt.Errorf("OCI_REGION não encontrado. Verifique se a variável está definida no arquivo .env ou no ambiente do sistema")
	}

	// Verificar se o arquivo de chave existe
	if _, err := os.Stat(c.KeyFile); os.IsNotExist(err) {
		return fmt.Errorf("arquivo de chave não encontrado: %s. Verifique se o caminho está correto", c.KeyFile)
	}

	return nil
}

// PrintConfig exibe a configuração atual (sem mostrar dados sensíveis)
func (c *OCIConfig) PrintConfig() {
	fmt.Println("📋 Configuração OCI carregada:")
	fmt.Printf("  • Tenancy ID: %s...%s\n", c.TenancyOCID[:20], c.TenancyOCID[len(c.TenancyOCID)-10:])
	fmt.Printf("  • User ID: %s...%s\n", c.UserOCID[:20], c.UserOCID[len(c.UserOCID)-10:])
	fmt.Printf("  • Key File: %s\n", c.KeyFile)
	fmt.Printf("  • Fingerprint: %s\n", c.Fingerprint)
	fmt.Printf("  • Region: %s\n", c.Region)
}
