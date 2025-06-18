package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Função para listar modelos disponíveis
func ListAvailableModels() {
	fmt.Println("\n=== MODELOS DISPONÍVEIS ===")
	fmt.Println()

	// Agrupar por família
	cohereModels := make([]string, 0)
	metaModels := make([]string, 0)

	for modelId, description := range SupportedModels {
		if isCohere(modelId) {
			cohereModels = append(cohereModels, fmt.Sprintf("  %s - %s", modelId, description))
		} else if isMetaLlama(modelId) {
			metaModels = append(metaModels, fmt.Sprintf("  %s - %s", modelId, description))
		}
	}

	fmt.Println("🤖 Modelos Cohere:")
	for _, model := range cohereModels {
		fmt.Println(model)
	}

	fmt.Println("\n🦙 Modelos Meta Llama:")
	for _, model := range metaModels {
		fmt.Println(model)
	}

	fmt.Println()
}

// Função para selecionar modelo interativamente
func SelectModelInteractively() string {
	ListAvailableModels()

	fmt.Println("Escolha um modelo:")
	fmt.Println("1. Cohere Command A (Março 2025)")
	fmt.Println("2. Cohere Command R (Agosto 2024)")
	fmt.Println("3. Cohere Command R Plus (Agosto 2024)")
	fmt.Println("4. Meta Llama 3.3 70B Instruct")
	fmt.Println("5. Meta Llama 3.1 70B Instruct")
	fmt.Println("6. Meta Llama 3.1 8B Instruct")
	fmt.Println("7. Meta Llama 2 70B Chat")

	fmt.Print("\nDigite o número do modelo (1-7): ")

	var choice string
	fmt.Scanln(&choice)

	choiceNum, err := strconv.Atoi(strings.TrimSpace(choice))
	if err != nil || choiceNum < 1 || choiceNum > 7 {
		fmt.Println("Escolha inválida. Usando modelo padrão: Meta Llama 3.3 70B")
		return ModelMetaLlama33_70B
	}

	models := []string{
		ModelCohereCommandA03,
		ModelCohereCommandR08,
		ModelCohereCommandRPlus08,
		ModelMetaLlama33_70B,
		ModelMetaLlama31_70B,
		ModelMetaLlama31_8B,
		ModelMetaLlama2_70B,
	}

	selectedModel := models[choiceNum-1]
	fmt.Printf("Modelo selecionado: %s (%s)\n\n", selectedModel, SupportedModels[selectedModel])

	return selectedModel
}

// Função para validar se o modelo é suportado
func IsModelSupported(modelId string) bool {
	_, exists := SupportedModels[modelId]
	return exists
}

// Função para obter informações do modelo
func GetModelInfo(modelId string) (string, string, bool) {
	description, exists := SupportedModels[modelId]
	if !exists {
		return "", "", false
	}

	family := GetModelFamily(modelId)
	return description, family, true
}
