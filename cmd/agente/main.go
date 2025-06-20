package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/generativeaiinference"

	"agente/internal/domain"
	"agente/internal/infrastructure"
)

func main() {
	fmt.Println("🚀 Oracle AI Generative Agent")
	fmt.Println("=============================")

	// Carregar configuração OCI do arquivo .env
	cfg := infrastructure.LoadConfig()
	cfg.PrintConfig()
	fmt.Println()

	// Selecionar modelo interativamente
	selectedModel := domain.SelectModelInteractively()

	// Validar se o modelo é suportado
	if !domain.IsModelSupported(selectedModel) {
		log.Fatalf("Modelo não suportado: %s", selectedModel)
	}

	// Obter informações do modelo
	description, family, _ := domain.GetModelInfo(selectedModel)
	fmt.Printf("Usando modelo: %s (%s)\n", description, family)
	fmt.Printf("Família: %s\n\n", family)

	// Criar provider de configuração
	provider := createProvider(cfg)

	// Criar cliente OCI
	client, err := generativeaiinference.NewGenerativeAiInferenceClientWithConfigurationProvider(provider)
	if err != nil {
		log.Fatalf("Erro ao criar cliente: %v", err)
	}

	// Criar implementação específica do modelo
	modelImpl := domain.CreateModelImplementation(selectedModel)
	if modelImpl == nil {
		log.Fatalf("Implementação não encontrada para o modelo: %s", selectedModel)
	}

	// Criar sessão de chat
	session := domain.NewChatSession(selectedModel, description)

	// Iniciar sessão de múltiplas perguntas
	startChatSession(client, modelImpl, cfg, selectedModel, description, session)
}

func startChatSession(client generativeaiinference.GenerativeAiInferenceClient, modelImpl domain.ModelImplementation, cfg infrastructure.OCIConfig, selectedModel, description string, session *domain.ChatSession) {
	reader := bufio.NewReader(os.Stdin)

	// Exibir instruções
	printInstructions()

	for {
		// Solicitar pergunta
		fmt.Printf("\n📝 Pergunta %d: ", len(session.Questions)+1)
		inputText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Erro ao ler entrada: %v\n", err)
			continue
		}

		inputText = strings.TrimSpace(inputText)

		// Verificar comandos especiais
		if shouldExit(inputText) {
			fmt.Println("\n👋 Encerrando sessão...")
			session.ShowStats()
			fmt.Println("Até logo!")
			break
		}

		if shouldChangeModel(inputText) {
			fmt.Println("\n🔄 Funcionalidade de troca de modelo será implementada em versão futura.")
			fmt.Println("Por enquanto, reinicie o programa para trocar de modelo.")
			continue
		}

		if shouldShowHelp(inputText) {
			printInstructions()
			continue
		}

		if shouldShowHistory(inputText) {
			session.ShowHistory()
			continue
		}

		if shouldShowStats(inputText) {
			session.ShowStats()
			continue
		}

		if shouldClearScreen(inputText) {
			clearScreen()
			fmt.Printf("🤖 Sessão ativa com %s\n", description)
			fmt.Printf("📊 Perguntas feitas: %d\n", len(session.Questions))
			fmt.Printf("🧠 %s\n", session.GetContextStatus())
			continue
		}

		if shouldToggleContext(inputText) {
			session.ToggleContext()
			fmt.Printf("🔄 %s\n", session.GetContextStatus())
			continue
		}

		if shouldShowContextStatus(inputText) {
			fmt.Printf("📋 %s\n", session.GetContextStatus())
			if session.IsContextEnabled() && len(session.Questions) > 0 {
				fmt.Printf("💭 Perguntas no contexto: %d\n", len(session.Questions))
			}
			continue
		}

		if inputText == "" {
			fmt.Println("⚠️  Pergunta vazia. Digite sua pergunta ou 'ajuda' para ver os comandos.")
			continue
		}

		// Processar pergunta
		processQuestion(client, modelImpl, cfg, selectedModel, description, inputText, session)
	}
}

func processQuestion(client generativeaiinference.GenerativeAiInferenceClient, modelImpl domain.ModelImplementation, cfg infrastructure.OCIConfig, selectedModel, description, inputText string, session *domain.ChatSession) {
	questionNumber := len(session.Questions) + 1
	fmt.Printf("🤔 Processando pergunta %d...\n", questionNumber)

	startTime := time.Now()

	// Criar requisição usando a implementação específica com contexto
	var chatRequest generativeaiinference.ChatRequest
	if session.IsContextEnabled() && len(session.Questions) > 0 {
		// Usar contexto se está ativado e há perguntas anteriores
		chatRequest = modelImpl.CreateChatRequestWithContext(cfg.TenancyOCID, selectedModel, inputText, session.Questions)
		fmt.Printf("💭 Usando contexto de %d perguntas anteriores\n", len(session.Questions))
	} else {
		// Primeira pergunta ou contexto desativado
		chatRequest = modelImpl.CreateChatRequest(cfg.TenancyOCID, selectedModel, inputText)
		if len(session.Questions) == 0 {
			fmt.Println("🆕 Primeira pergunta da sessão")
		} else {
			fmt.Println("🧠 Contexto desativado - pergunta independente")
		}
	}

	// Fazer a requisição
	resp, err := client.Chat(context.Background(), chatRequest)
	processTime := time.Since(startTime)

	if err != nil {
		errorMsg := fmt.Sprintf("Erro ao processar pergunta: %v", err)
		fmt.Printf("❌ %s\n", errorMsg)
		fmt.Println("💡 Tente reformular sua pergunta ou verificar sua conexão.")

		// Adicionar ao histórico como erro
		session.AddQuestion(inputText, "", processTime, false, errorMsg)
		return
	}

	// Processar resposta usando a implementação específica
	response, err := modelImpl.ProcessResponse(resp)
	if err != nil {
		errorMsg := fmt.Sprintf("Erro ao processar resposta: %v", err)
		fmt.Printf("❌ %s\n", errorMsg)

		// Adicionar ao histórico como erro
		session.AddQuestion(inputText, "", processTime, false, errorMsg)
		return
	}

	// Adicionar ao histórico como sucesso
	session.AddQuestion(inputText, response, processTime, true, "")

	// Exibir resultado
	printResponse(description, response, questionNumber, processTime)
}

func printResponse(description, response string, questionNumber int, processTime time.Duration) {
	separator := strings.Repeat("=", 70)
	fmt.Printf("\n%s\n", separator)
	fmt.Printf("🤖 Resposta %d - %s:\n", questionNumber, description)
	fmt.Printf("⚡ Processado em: %v\n", processTime.Round(time.Millisecond))
	fmt.Printf("%s\n", separator)
	fmt.Println(response)
	fmt.Printf("%s\n", separator)
}

func printInstructions() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("📋 INSTRUÇÕES DE USO")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("• Digite suas perguntas normalmente")
	fmt.Println("• Comandos especiais:")
	fmt.Println("  - 'sair', 'exit', 'quit' → Encerrar sessão")
	fmt.Println("  - 'ajuda', 'help', '?' → Mostrar estas instruções")
	fmt.Println("  - 'historico', 'history' → Ver histórico de perguntas")
	fmt.Println("  - 'stats', 'estatisticas' → Ver estatísticas da sessão")
	fmt.Println("  - 'limpar', 'clear' → Limpar tela")
	fmt.Println("  - 'contexto', 'context' → Ativar/desativar contexto")
	fmt.Println("  - 'status', 'estado' → Ver status do contexto")
	fmt.Println("  - 'trocar', 'modelo' → Informações sobre troca de modelo")
	fmt.Println("• Pressione Enter após cada pergunta")
	fmt.Println("• Para perguntas longas, digite normalmente em uma linha")
	fmt.Println("• 🧠 Contexto: Quando ativado, o modelo lembra das perguntas anteriores")
	fmt.Println(strings.Repeat("=", 70))
}

func shouldExit(input string) bool {
	exitCommands := []string{"sair", "exit", "quit", "bye", "tchau", "fim"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range exitCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func shouldChangeModel(input string) bool {
	changeCommands := []string{"trocar", "modelo", "change", "switch"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range changeCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func shouldShowHelp(input string) bool {
	helpCommands := []string{"ajuda", "help", "?", "comandos", "instrucoes"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range helpCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func shouldShowHistory(input string) bool {
	historyCommands := []string{"historico", "history", "hist"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range historyCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func shouldShowStats(input string) bool {
	statsCommands := []string{"stats", "estatisticas", "estatística", "statistics"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range statsCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func shouldClearScreen(input string) bool {
	clearCommands := []string{"limpar", "clear", "cls"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range clearCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func shouldToggleContext(input string) bool {
	contextCommands := []string{"contexto", "context", "toggle", "alternar"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range contextCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func shouldShowContextStatus(input string) bool {
	statusCommands := []string{"status", "estado", "contexto?", "context?"}
	input = strings.ToLower(strings.TrimSpace(input))

	for _, cmd := range statusCommands {
		if input == cmd {
			return true
		}
	}
	return false
}

func clearScreen() {
	// Limpar tela (funciona no Windows e Unix)
	fmt.Print("\033[2J\033[H")
}

func createProvider(cfg infrastructure.OCIConfig) common.ConfigurationProvider {
	// Ler o conteúdo do arquivo PEM
	privateKeyContent, err := os.ReadFile(cfg.KeyFile)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo PEM: %v", err)
	}

	provider := common.NewRawConfigurationProvider(
		cfg.TenancyOCID,
		cfg.UserOCID,
		cfg.Region,
		cfg.Fingerprint,
		string(privateKeyContent),
		nil,
	)

	return provider
}
