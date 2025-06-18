package main

import (
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/generativeaiinference"
)

// MetaImplementation implementa a interface ModelImplementation para modelos Meta Llama
type MetaImplementation struct{}

// CreateChatRequest cria uma requisição de chat específica para modelos Meta Llama
func (m *MetaImplementation) CreateChatRequest(compartmentId, modelId, inputText string) generativeaiinference.ChatRequest {
	return generativeaiinference.ChatRequest{
		ChatDetails: generativeaiinference.ChatDetails{
			CompartmentId: common.String(compartmentId),
			ServingMode: generativeaiinference.OnDemandServingMode{
				ModelId: common.String(modelId),
			},
			ChatRequest: generativeaiinference.GenericChatRequest{
				Messages: []generativeaiinference.Message{
					generativeaiinference.UserMessage{
						Content: []generativeaiinference.ChatContent{
							generativeaiinference.TextContent{
								Text: common.String(inputText),
							},
						},
					},
				},
				MaxTokens:   common.Int(600),
				Temperature: common.Float64(0.1),
				TopP:        common.Float64(0.75),
				IsStream:    common.Bool(false),
			},
		},
	}
}

// CreateChatRequestWithContext cria uma requisição de chat com contexto histórico para modelos Meta Llama
func (m *MetaImplementation) CreateChatRequestWithContext(compartmentId, modelId, inputText string, context []Question) generativeaiinference.ChatRequest {
	var messages []generativeaiinference.Message

	// Adicionar contexto histórico limitado (últimas 5 interações para não exceder limites de token)
	maxContext := 5
	startIndex := 0
	if len(context) > maxContext {
		startIndex = len(context) - maxContext
	}

	for _, q := range context[startIndex:] {
		if q.Success {
			// Adicionar pergunta do usuário
			messages = append(messages, generativeaiinference.UserMessage{
				Content: []generativeaiinference.ChatContent{
					generativeaiinference.TextContent{
						Text: common.String(q.Text),
					},
				},
			})

			// Adicionar resposta do assistente
			messages = append(messages, generativeaiinference.AssistantMessage{
				Content: []generativeaiinference.ChatContent{
					generativeaiinference.TextContent{
						Text: common.String(q.Response),
					},
				},
			})
		}
	}

	// Adicionar a pergunta atual
	messages = append(messages, generativeaiinference.UserMessage{
		Content: []generativeaiinference.ChatContent{
			generativeaiinference.TextContent{
				Text: common.String(inputText),
			},
		},
	})

	return generativeaiinference.ChatRequest{
		ChatDetails: generativeaiinference.ChatDetails{
			CompartmentId: common.String(compartmentId),
			ServingMode: generativeaiinference.OnDemandServingMode{
				ModelId: common.String(modelId),
			},
			ChatRequest: generativeaiinference.GenericChatRequest{
				Messages:    messages,
				MaxTokens:   common.Int(600),
				Temperature: common.Float64(0.1),
				TopP:        common.Float64(0.75),
				IsStream:    common.Bool(false),
			},
		},
	}
}

// ProcessResponse processa a resposta específica para modelos Meta Llama
func (m *MetaImplementation) ProcessResponse(response generativeaiinference.ChatResponse) (string, error) {
	if chatResponse, ok := response.ChatResult.ChatResponse.(generativeaiinference.GenericChatResponse); ok {
		if len(chatResponse.Choices) > 0 {
			content := chatResponse.Choices[0].Message.GetContent()
			if len(content) > 0 {
				// Extrair o texto do primeiro conteúdo
				if textContent, ok := content[0].(generativeaiinference.TextContent); ok {
					if textContent.Text != nil {
						return *textContent.Text, nil
					}
				}
			}
		}
		return "", fmt.Errorf("nenhuma resposta recebida do modelo Meta Llama")
	}

	return "", fmt.Errorf("formato de resposta inesperado para Meta Llama: %T", response.ChatResult.ChatResponse)
}

// GetModelFamily retorna a família do modelo
func (m *MetaImplementation) GetModelFamily() string {
	return "meta"
}
