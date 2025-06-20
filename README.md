# Oracle AI Generative Agent

Este projeto é um agente de IA generativa modular que utiliza o Oracle Cloud Infrastructure (OCI) Generative AI service para interagir com diferentes modelos de linguagem de forma dinâmica, organizando sessões interativas com múltiplas perguntas.

## 📋 Descrição

O Oracle AI Generative Agent é uma aplicação Go que se conecta ao serviço de IA Generativa da Oracle Cloud com suporte a múltiplos modelos de linguagem. O projeto utiliza uma arquitetura modular que permite fácil adição de novos modelos, mantém histórico de conversas e oferece estatísticas detalhadas da sessão.

## 🚀 Funcionalidades

- ✅ **Múltiplas Perguntas**: Sessão interativa contínua com o modelo
- ✅ **Histórico de Conversas**: Registro completo de perguntas e respostas
- ✅ **Estatísticas da Sessão**: Métricas de performance e uso
- ✅ **Seleção Dinâmica de Modelos**: Escolha interativa entre diferentes modelos
- ✅ **Arquitetura Modular**: Cada família de modelo tem sua própria implementação
- ✅ **Suporte Multi-Modelo**: Cohere e Meta Llama
- ✅ **Interface Unificada**: Mesma API para todos os modelos
- ✅ **Comandos Especiais**: Controle avançado da sessão
- ✅ **Validação de Modelos**: Verificação automática de compatibilidade
- ✅ **Autenticação Segura**: Via chave privada PEM
- ✅ **Região Configurável**: Suporte à região sa-saopaulo-1

## 🤖 Modelos Suportados

### Família Cohere:
- `cohere.command-a-03-2025` - Cohere Command A (Março 2025)
- `cohere.command-r-08-2024` - Cohere Command R (Agosto 2024)
- `cohere.command-r-plus-08-2024` - Cohere Command R Plus (Agosto 2024)

### Família Meta Llama:
- `meta.llama-3.3-70b-instruct` - Meta Llama 3.3 70B Instruct
- `meta.llama-3.1-70b-instruct` - Meta Llama 3.1 70B Instruct
- `meta.llama-3.1-8b-instruct` - Meta Llama 3.1 8B Instruct
- `meta.llama-2-70b-chat` - Meta Llama 2 70B Chat

## 📦 Estrutura do Projeto

```
agente/
├── cmd/
│   └── agente/
│       └── main.go              # Aplicação principal com sistema de múltiplas perguntas
├── internal/
│   ├── domain/                  # Lógica de negócio e domínio
│   │   ├── models.go           # Constantes e interfaces dos modelos
│   │   ├── chat_session.go     # Sistema de sessões e histórico
│   │   ├── utils.go            # Utilitários e funções auxiliares
│   │   ├── cohere_implementation.go  # Implementação específica Cohere
│   │   └── meta_implementation.go    # Implementação específica Meta Llama
│   └── infrastructure/          # Configurações e infraestrutura
│       └── config.go           # Configurações OCI
├── go.mod                      # Dependências Go
├── go.sum                      # Lock das dependências
├── agente.exe                  # Executável compilado
├── *.pem                       # Chave privada OCI
└── README.md                   # Documentação
```

## 🛠 Pré-requisitos

- **Go 1.19+** instalado
- **Conta Oracle Cloud** ativa
- **Chave API OCI** configurada
- **Arquivo PEM** da chave privada na raiz do projeto

## ⚙️ Configuração

### 1. Credenciais OCI

Certifique-se de ter as seguintes informações da sua conta Oracle Cloud:

- **Tenancy OCID**: Identificador do tenant
- **User OCID**: Identificador do usuário
- **Fingerprint**: Impressão digital da chave API
- **Region**: Região (ex: sa-saopaulo-1)
- **Arquivo PEM**: Chave privada no formato PEM

### 2. Configurar Credenciais

Edite o arquivo `main.go` e atualize a estrutura `OCIConfig` com suas credenciais:

```go
cfg := OCIConfig{
    TenancyOCID: "seu-tenancy-ocid",
    UserOCID:    "seu-user-ocid", 
    KeyFile:     "seu-arquivo.pem",
    Fingerprint: "sua-fingerprint",
    Region:      "sa-saopaulo-1",
}
```

## 🔧 Compilação

```bash
# Baixar dependências
go mod download

# Compilar o projeto (usa cmd/agente como entry point)
go build -o agente.exe ./cmd/agente
```

## 🚀 Execução

### Execução Interativa
```bash
# Executar com seleção interativa de modelo
./agente.exe
```

## 📖 Como Usar

### 🎯 Fluxo Principal

1. **Execute o programa**: `./agente.exe`
2. **Escolha um modelo**: Selecione entre os 7 modelos disponíveis
3. **Inicie a sessão**: Faça múltiplas perguntas na mesma sessão
4. **Use comandos especiais**: Controle avançado da sessão
5. **Veja estatísticas**: Métricas automáticas ao final

### 🎮 Comandos Especiais

Durante a sessão, você pode usar os seguintes comandos:

- **`sair`**, **`exit`**, **`quit`** → Encerrar sessão (mostra estatísticas)
- **`ajuda`**, **`help`**, **`?`** → Mostrar instruções
- **`historico`**, **`history`** → Ver histórico completo de perguntas
- **`stats`**, **`estatisticas`** → Ver estatísticas da sessão atual
- **`limpar`**, **`clear`** → Limpar tela
- **`trocar`**, **`modelo`** → Informações sobre troca de modelo

### 📊 Funcionalidades da Sessão

#### **Histórico de Conversas**
- 📝 Registro completo de todas as perguntas e respostas
- ⏰ Timestamp de cada interação
- ⚡ Tempo de processamento individual
- ✅ Status de sucesso/erro para cada pergunta

#### **Estatísticas em Tempo Real**
- 📈 Taxa de sucesso das perguntas
- ⏱️ Tempo médio de processamento
- 📊 Total de perguntas feitas
- 🕐 Duração da sessão

### Exemplo de Uso Completo:

```
🚀 Oracle AI Generative Agent
=============================

=== MODELOS DISPONÍVEIS ===

🤖 Modelos Cohere:
  cohere.command-a-03-2025 - Cohere Command A (Março 2025)
  cohere.command-r-08-2024 - Cohere Command R (Agosto 2024)
  cohere.command-r-plus-08-2024 - Cohere Command R Plus (Agosto 2024)

🦙 Modelos Meta Llama:
  meta.llama-3.3-70b-instruct - Meta Llama 3.3 70B Instruct
  meta.llama-3.1-70b-instruct - Meta Llama 3.1 70B Instruct
  meta.llama-3.1-8b-instruct - Meta Llama 3.1 8B Instruct
  meta.llama-2-70b-chat - Meta Llama 2 70B Chat

Escolha um modelo:
1. Cohere Command A (Março 2025)
2. Cohere Command R (Agosto 2024)
3. Cohere Command R Plus (Agosto 2024)
4. Meta Llama 3.3 70B Instruct
5. Meta Llama 3.1 70B Instruct
6. Meta Llama 3.1 8B Instruct
7. Meta Llama 2 70B Chat

Digite o número do modelo (1-7): 4

Modelo selecionado: meta.llama-3.3-70b-instruct (Meta Llama 3.3 70B Instruct)
Usando modelo: Meta Llama 3.3 70B Instruct (meta)
Família: meta

======================================================================
📋 INSTRUÇÕES DE USO
======================================================================
• Digite suas perguntas normalmente
• Comandos especiais:
  - 'sair', 'exit', 'quit' → Encerrar sessão
  - 'ajuda', 'help', '?' → Mostrar estas instruções
  - 'historico', 'history' → Ver histórico de perguntas
  - 'stats', 'estatisticas' → Ver estatísticas da sessão
  - 'limpar', 'clear' → Limpar tela
  - 'trocar', 'modelo' → Informações sobre troca de modelo
• Pressione Enter após cada pergunta
• Para perguntas longas, digite normalmente em uma linha
======================================================================

📝 Pergunta 1: Qual é o maior planeta do sistema solar?
🤔 Processando pergunta 1...

======================================================================
🤖 Resposta 1 - Meta Llama 3.3 70B Instruct:
⚡ Processado em: 1.234s
======================================================================
O maior planeta do sistema solar é Júpiter.
======================================================================

📝 Pergunta 2: E qual é o menor?
🤔 Processando pergunta 2...

======================================================================
🤖 Resposta 2 - Meta Llama 3.3 70B Instruct:
⚡ Processado em: 987ms
======================================================================
O menor planeta do sistema solar é Mercúrio.
======================================================================

📝 Pergunta 3: stats

============================================================
📊 ESTATÍSTICAS DA SESSÃO
============================================================
🤖 Modelo utilizado: Meta Llama 3.3 70B Instruct
⏰ Duração da sessão: 45s
📝 Total de perguntas: 2
✅ Perguntas bem-sucedidas: 2
❌ Perguntas com erro: 0
📈 Taxa de sucesso: 100.0%
⚡ Tempo médio por pergunta: 1.110s
============================================================

📝 Pergunta 3: sair

👋 Encerrando sessão...
============================================================
📊 ESTATÍSTICAS DA SESSÃO
============================================================
🤖 Modelo utilizado: Meta Llama 3.3 70B Instruct
⏰ Duração da sessão: 1m2s
📝 Total de perguntas: 2
✅ Perguntas bem-sucedidas: 2
❌ Perguntas com erro: 0
📈 Taxa de sucesso: 100.0%
⚡ Tempo médio por pergunta: 1.110s
============================================================
Até logo!
```

## 📚 Dependências

```go
require (
    github.com/oracle/oci-go-sdk/v65 v65.93.2
)
```

## 🏗️ Arquitetura e Padrões

O projeto segue a **estrutura padrão do Go** para melhor organização e manutenibilidade:

### 📁 Organização dos Diretórios

- **`cmd/agente/`** - Contém o ponto de entrada da aplicação (main.go)
- **`internal/domain/`** - Lógica de negócio, modelos e implementações específicas
- **`internal/infrastructure/`** - Configurações, integrações externas e infraestrutura
- **Raiz do projeto** - Arquivos de configuração (go.mod, README.md, etc.)

### ✨ Benefícios desta Estrutura

- **🔒 Encapsulamento**: Código em `internal/` não pode ser importado por outros projetos
- **📦 Modularidade**: SeparaçÃo clara entre domínio e infraestrutura
- **🔧 Manutenibilidade**: Facilita mudanças e adição de novas funcionalidades
- **📚 Padrão Go**: Segue as convenções estabelecidas pela comunidade Go
- **🎯 Testabilidade**: Estrutura favorece criação de testes unitários

### 🔗 Dependências entre Módulos

```
cmd/agente → internal/domain + internal/infrastructure
internal/domain ← internal/infrastructure
```

## 🔍 Resolução de Problemas

### Erro de Autenticação PEM
- Verifique se o arquivo .pem está na raiz do projeto
- Confirme se as credenciais estão corretas no código
- Teste a conectividade com a Oracle Cloud

### Modelo Não Encontrado
- Verifique se o modelo está disponível na sua região
- Confirme se você tem acesso ao modelo selecionado
- Tente usar um modelo diferente

### Erro de Compilação
```bash
go mod tidy
go build -o agente.exe
```

### Sessão Travada
- Use `Ctrl+C` para forçar saída
- Digite `sair` para encerramento normal
- Verifique conexão de rede se perguntas não processam

## 🚀 Melhorias Futuras

- [x] **Múltiplas perguntas**: ✅ Implementado
- [x] **Histórico de conversas**: ✅ Implementado
- [x] **Estatísticas de sessão**: ✅ Implementado
- [ ] **Configuração via arquivo**: Carregar credenciais de arquivo config
- [ ] **Cache de respostas**: Sistema de cache para otimização
- [ ] **Troca de modelo em tempo real**: Mudar modelo durante a sessão
- [x] **Contexto entre perguntas**: ✅ Manter contexto da conversa
- [ ] **Exportar histórico**: Salvar conversas em arquivo
- [ ] **Interface web**: Criar interface web para facilitar uso
- [ ] **Logs detalhados**: Sistema de logging mais robusto
- [ ] **Testes automatizados**: Cobertura de testes unitários


### Adicionando Novos Modelos

1. Adicione a constante em `models.go`
2. Atualize o mapa `SupportedModels`
3. Crie implementação específica se necessário
4. Atualize as funções de verificação de família

## 📚 Referência

- **Documentação Oracle OCI**: https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/
- **SDK Go Oracle**: https://github.com/oracle/oci-go-sdk
- **Generative AI**: https://docs.oracle.com/iaas/Content/generative-ai/home.htm

---