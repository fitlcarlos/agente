package main

import (
	"github.com/oracle/oci-go-sdk/v65/generativeaiinference"
)

// Constantes dos modelos suportados
const (
	// Modelos Cohere
	ModelCohereCommandA03     = "cohere.command-a-03-2025"
	ModelCohereCommandR08     = "cohere.command-r-08-2024"
	ModelCohereCommandRPlus08 = "cohere.command-r-plus-08-2024"

	// Modelos Meta Llama
	ModelMetaLlama33_70B = "meta.llama-3.3-70b-instruct"
	ModelMetaLlama31_70B = "meta.llama-3.1-70b-instruct"
	ModelMetaLlama31_8B  = "meta.llama-3.1-8b-instruct"
	ModelMetaLlama2_70B  = "meta.llama-2-70b-chat"
)

// Mapa de modelos suportados com suas descrições
var SupportedModels = map[string]string{
	ModelCohereCommandA03:     "Cohere Command A (Março 2025)",
	ModelCohereCommandR08:     "Cohere Command R (Agosto 2024)",
	ModelCohereCommandRPlus08: "Cohere Command R Plus (Agosto 2024)",
	ModelMetaLlama33_70B:      "Meta Llama 3.3 70B Instruct",
	ModelMetaLlama31_70B:      "Meta Llama 3.1 70B Instruct",
	ModelMetaLlama31_8B:       "Meta Llama 3.1 8B Instruct",
	ModelMetaLlama2_70B:       "Meta Llama 2 70B Chat",
}

// Interface para implementações de modelos
type ModelImplementation interface {
	CreateChatRequest(compartmentId, modelId, inputText string) generativeaiinference.ChatRequest
	CreateChatRequestWithContext(compartmentId, modelId, inputText string, context []Question) generativeaiinference.ChatRequest
	ProcessResponse(response generativeaiinference.ChatResponse) (string, error)
	GetModelFamily() string
}

// Função para determinar a família do modelo
func GetModelFamily(modelId string) string {
	switch {
	case isCohere(modelId):
		return "cohere"
	case isMetaLlama(modelId):
		return "meta"
	default:
		return "unknown"
	}
}

// Função para verificar se é modelo Cohere
func isCohere(modelId string) bool {
	cohereModels := []string{
		ModelCohereCommandA03,
		ModelCohereCommandR08,
		ModelCohereCommandRPlus08,
	}

	for _, model := range cohereModels {
		if modelId == model {
			return true
		}
	}
	return false
}

// Função para verificar se é modelo Meta Llama
func isMetaLlama(modelId string) bool {
	metaModels := []string{
		ModelMetaLlama33_70B,
		ModelMetaLlama31_70B,
		ModelMetaLlama31_8B,
		ModelMetaLlama2_70B,
	}

	for _, model := range metaModels {
		if modelId == model {
			return true
		}
	}
	return false
}

// Factory para criar implementação do modelo
func CreateModelImplementation(modelId string) ModelImplementation {
	family := GetModelFamily(modelId)

	switch family {
	case "cohere":
		return &CohereImplementation{}
	case "meta":
		return &MetaImplementation{}
	default:
		return nil
	}
}
