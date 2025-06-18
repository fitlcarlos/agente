package main

import (
	"fmt"
	"strings"
	"time"
)

// ChatSession representa uma sessão de chat com histórico
type ChatSession struct {
	ModelID        string
	ModelName      string
	StartTime      time.Time
	Questions      []Question
	TotalTime      time.Duration
	ContextEnabled bool // Controla se o contexto deve ser mantido entre perguntas
}

// Question representa uma pergunta e sua resposta
type Question struct {
	ID          int
	Text        string
	Response    string
	Timestamp   time.Time
	ProcessTime time.Duration
	Success     bool
	Error       string
}

// NewChatSession cria uma nova sessão de chat
func NewChatSession(modelID, modelName string) *ChatSession {
	return &ChatSession{
		ModelID:        modelID,
		ModelName:      modelName,
		StartTime:      time.Now(),
		Questions:      make([]Question, 0),
		ContextEnabled: true, // Contexto ativado por padrão
	}
}

// AddQuestion adiciona uma pergunta ao histórico
func (cs *ChatSession) AddQuestion(text, response string, processTime time.Duration, success bool, errorMsg string) {
	question := Question{
		ID:          len(cs.Questions) + 1,
		Text:        text,
		Response:    response,
		Timestamp:   time.Now(),
		ProcessTime: processTime,
		Success:     success,
		Error:       errorMsg,
	}

	cs.Questions = append(cs.Questions, question)
}

// GetStats retorna estatísticas da sessão
func (cs *ChatSession) GetStats() SessionStats {
	totalQuestions := len(cs.Questions)
	successfulQuestions := 0
	totalProcessTime := time.Duration(0)

	for _, q := range cs.Questions {
		if q.Success {
			successfulQuestions++
		}
		totalProcessTime += q.ProcessTime
	}

	sessionDuration := time.Since(cs.StartTime)

	return SessionStats{
		TotalQuestions:      totalQuestions,
		SuccessfulQuestions: successfulQuestions,
		FailedQuestions:     totalQuestions - successfulQuestions,
		SessionDuration:     sessionDuration,
		AverageProcessTime:  calculateAverageTime(totalProcessTime, successfulQuestions),
		ModelUsed:           cs.ModelName,
	}
}

// ShowHistory exibe o histórico de perguntas
func (cs *ChatSession) ShowHistory() {
	if len(cs.Questions) == 0 {
		fmt.Println("📝 Nenhuma pergunta foi feita ainda nesta sessão.")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("📚 HISTÓRICO DA SESSÃO")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("🤖 Modelo: %s\n", cs.ModelName)
	fmt.Printf("⏰ Iniciado em: %s\n", cs.StartTime.Format("15:04:05"))
	fmt.Printf("📊 Total de perguntas: %d\n", len(cs.Questions))
	fmt.Println(strings.Repeat("-", 70))

	for _, q := range cs.Questions {
		status := "✅"
		if !q.Success {
			status = "❌"
		}

		fmt.Printf("\n%s Pergunta %d [%s]:\n", status, q.ID, q.Timestamp.Format("15:04:05"))
		fmt.Printf("❓ %s\n", q.Text)

		if q.Success {
			// Truncar resposta se muito longa
			response := q.Response
			if len(response) > 200 {
				response = response[:200] + "..."
			}
			fmt.Printf("💬 %s\n", response)
			fmt.Printf("⚡ Tempo de processamento: %v\n", q.ProcessTime.Round(time.Millisecond))
		} else {
			fmt.Printf("💥 Erro: %s\n", q.Error)
		}

		if q.ID < len(cs.Questions) {
			fmt.Println(strings.Repeat("-", 50))
		}
	}

	fmt.Println(strings.Repeat("=", 70))
}

// ShowStats exibe estatísticas da sessão
func (cs *ChatSession) ShowStats() {
	stats := cs.GetStats()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 ESTATÍSTICAS DA SESSÃO")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("🤖 Modelo utilizado: %s\n", stats.ModelUsed)
	fmt.Printf("⏰ Duração da sessão: %v\n", stats.SessionDuration.Round(time.Second))
	fmt.Printf("📝 Total de perguntas: %d\n", stats.TotalQuestions)
	fmt.Printf("✅ Perguntas bem-sucedidas: %d\n", stats.SuccessfulQuestions)
	fmt.Printf("❌ Perguntas com erro: %d\n", stats.FailedQuestions)

	if stats.SuccessfulQuestions > 0 {
		successRate := float64(stats.SuccessfulQuestions) / float64(stats.TotalQuestions) * 100
		fmt.Printf("📈 Taxa de sucesso: %.1f%%\n", successRate)
		fmt.Printf("⚡ Tempo médio por pergunta: %v\n", stats.AverageProcessTime.Round(time.Millisecond))
	}

	fmt.Println(strings.Repeat("=", 60))
}

// SessionStats contém estatísticas da sessão
type SessionStats struct {
	TotalQuestions      int
	SuccessfulQuestions int
	FailedQuestions     int
	SessionDuration     time.Duration
	AverageProcessTime  time.Duration
	ModelUsed           string
}

// calculateAverageTime calcula o tempo médio
func calculateAverageTime(total time.Duration, count int) time.Duration {
	if count == 0 {
		return 0
	}
	return total / time.Duration(count)
}

// GetLastQuestions retorna as últimas N perguntas
func (cs *ChatSession) GetLastQuestions(n int) []Question {
	if len(cs.Questions) == 0 {
		return []Question{}
	}

	start := len(cs.Questions) - n
	if start < 0 {
		start = 0
	}

	return cs.Questions[start:]
}

// ExportHistory exporta o histórico em formato texto
func (cs *ChatSession) ExportHistory() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("=== SESSÃO DE CHAT - %s ===\n", cs.StartTime.Format("02/01/2006 15:04:05")))
	builder.WriteString(fmt.Sprintf("Modelo: %s\n", cs.ModelName))
	builder.WriteString(fmt.Sprintf("Total de perguntas: %d\n\n", len(cs.Questions)))

	for _, q := range cs.Questions {
		builder.WriteString(fmt.Sprintf("PERGUNTA %d [%s]:\n", q.ID, q.Timestamp.Format("15:04:05")))
		builder.WriteString(fmt.Sprintf("%s\n\n", q.Text))

		if q.Success {
			builder.WriteString("RESPOSTA:\n")
			builder.WriteString(fmt.Sprintf("%s\n", q.Response))
			builder.WriteString(fmt.Sprintf("(Processado em %v)\n\n", q.ProcessTime.Round(time.Millisecond)))
		} else {
			builder.WriteString(fmt.Sprintf("ERRO: %s\n\n", q.Error))
		}

		builder.WriteString(strings.Repeat("-", 50) + "\n\n")
	}

	return builder.String()
}

// ToggleContext alterna o estado do contexto
func (cs *ChatSession) ToggleContext() {
	cs.ContextEnabled = !cs.ContextEnabled
}

// SetContext define o estado do contexto
func (cs *ChatSession) SetContext(enabled bool) {
	cs.ContextEnabled = enabled
}

// IsContextEnabled retorna se o contexto está ativado
func (cs *ChatSession) IsContextEnabled() bool {
	return cs.ContextEnabled
}

// GetContextStatus retorna uma string descrevendo o status do contexto
func (cs *ChatSession) GetContextStatus() string {
	if cs.ContextEnabled {
		return "🧠 Contexto: ATIVADO - O modelo lembrará das perguntas anteriores"
	}
	return "🧠 Contexto: DESATIVADO - Cada pergunta será independente"
}
