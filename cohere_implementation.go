package main

import (
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/generativeaiinference"
)

// CohereImplementation implementa a interface ModelImplementation para modelos Cohere
type CohereImplementation struct{}

// CreateChatRequest cria uma requisição de chat específica para modelos Cohere
func (c *CohereImplementation) CreateChatRequest(compartmentId, modelId, inputText string) generativeaiinference.ChatRequest {
	return generativeaiinference.ChatRequest{
		ChatDetails: generativeaiinference.ChatDetails{
			CompartmentId: common.String(compartmentId),
			ServingMode: generativeaiinference.OnDemandServingMode{
				ModelId: common.String(modelId),
			},
			ChatRequest: generativeaiinference.CohereChatRequest{
				Message:     common.String(inputText),
				MaxTokens:   common.Int(600),
				Temperature: common.Float64(0.1),
				TopP:        common.Float64(0.75),
				TopK:        common.Int(0),
				IsStream:    common.Bool(false),
			},
		},
	}
}

// CreateChatRequestWithContext cria uma requisição de chat com contexto histórico para modelos Cohere
func (c *CohereImplementation) CreateChatRequestWithContext(compartmentId, modelId, inputText string, context []Question) generativeaiinference.ChatRequest {
	// Para modelos Cohere, incluimos o contexto como parte da mensagem
	contextMessage := inputText

	// Adicionar contexto histórico limitado (últimas 3 interações para não exceder limites)
	maxContext := 3
	startIndex := 0
	if len(context) > maxContext {
		startIndex = len(context) - maxContext
	}

	if len(context) > 0 {
		contextMessage = "Contexto da conversa anterior:\n"
		for i, q := range context[startIndex:] {
			if q.Success {
				contextMessage += fmt.Sprintf("\nPergunta %d: %s\nResposta %d: %s\n", i+1, q.Text, i+1, q.Response)
			}
		}
		contextMessage += "\nPergunta atual: " + inputText
	}

	return generativeaiinference.ChatRequest{
		ChatDetails: generativeaiinference.ChatDetails{
			CompartmentId: common.String(compartmentId),
			ServingMode: generativeaiinference.OnDemandServingMode{
				ModelId: common.String(modelId),
			},
			ChatRequest: generativeaiinference.CohereChatRequest{
				Message:     common.String(contextMessage),
				MaxTokens:   common.Int(600),
				Temperature: common.Float64(0.1),
				TopP:        common.Float64(0.75),
				TopK:        common.Int(0),
				IsStream:    common.Bool(false),
			},
		},
	}
}

// ProcessResponse processa a resposta específica para modelos Cohere
func (c *CohereImplementation) ProcessResponse(response generativeaiinference.ChatResponse) (string, error) {
	if chatResponse, ok := response.ChatResult.ChatResponse.(generativeaiinference.CohereChatResponse); ok {
		if chatResponse.Text != nil {
			return *chatResponse.Text, nil
		}
		return "", fmt.Errorf("resposta vazia do modelo Cohere")
	}

	return "", fmt.Errorf("formato de resposta inesperado para Cohere: %T", response.ChatResult.ChatResponse)
}

// GetModelFamily retorna a família do modelo
func (c *CohereImplementation) GetModelFamily() string {
	return "cohere"
}
