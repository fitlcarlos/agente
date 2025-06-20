# Oracle AI Generative Agent 🤖

Este projeto é um agente de IA generativa modular que utiliza o Oracle Cloud Infrastructure (OCI) Generative AI service para interagir com diferentes modelos de linguagem de forma dinâmica, organizando sessões interativas com múltiplas perguntas e contexto entre conversas.

## 📋 Descrição

O Oracle AI Generative Agent é uma aplicação Go que se conecta ao serviço de IA Generativa da Oracle Cloud com suporte a múltiplos modelos de linguagem. O projeto utiliza uma arquitetura modular que permite fácil adição de novos modelos, mantém histórico de conversas com contexto e oferece estatísticas detalhadas da sessão.

## 🚀 Funcionalidades

- ✅ **Múltiplas Perguntas**: Sessão interativa contínua com o modelo
- ✅ **Contexto Inteligente**: Sistema de contexto que preserva o histórico da conversa
- ✅ **Configuração via .env**: Sistema robusto de configuração com fallback
- ✅ **Histórico de Conversas**: Registro completo de perguntas e respostas
- ✅ **Estatísticas da Sessão**: Métricas de performance e uso em tempo real
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
│       ├── main.go                    # Aplicação principal
│       ├── .env                       # Configurações OCI (não commitado)
│       ├── agente.exe                # Executável compilado
│       └── *.pem                     # Chave privada OCI
├── internal/
│   ├── domain/                       # Lógica de negócio e domínio
│   │   ├── models.go                 # Constantes e interfaces dos modelos
│   │   ├── chat_session.go           # Sistema de sessões e histórico
│   │   ├── utils.go                  # Utilitários e funções auxiliares
│   │   ├── cohere_implementation.go  # Implementação específica Cohere
│   │   └── meta_implementation.go    # Implementação específica Meta Llama
│   └── infrastructure/               # Configurações e infraestrutura
│       └── config.go                 # Sistema de configuração com .env
├── go.mod                           # Dependências Go
├── go.sum                           # Lock das dependências
└── README.md                        # Esta documentação
```

## 🛠 Pré-requisitos

- **Go 1.19+** instalado
- **Conta Oracle Cloud** ativa
- **Chave API OCI** configurada
- **Arquivo PEM** da chave privada

## ⚙️ Configuração

### 1. Configurar Credenciais OCI

Crie um arquivo `.env` no diretório `cmd/agente/` com as seguintes variáveis:

```bash
# Configuração Oracle Cloud Infrastructure (OCI)
# ID do seu tenancy OCI
OCI_TENANCY_ID=ocid1.tenancy.oc1..aaaaaaaa...

# ID do seu usuário OCI  
OCI_USER_ID=ocid1.user.oc1..aaaaaaaa...

# Caminho para o arquivo de chave privada (.pem)
OCI_KEY_FILE=sua-chave-privada.pem

# Fingerprint da sua chave pública
OCI_FINGERPRINT=xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx

# Região OCI (ex: sa-saopaulo-1, us-ashburn-1, etc.)
OCI_REGION=sa-saopaulo-1
```

### 2. Sistema de Configuração Robusto

O sistema de configuração implementa as seguintes funcionalidades:

#### ✅ **Verificação Inteligente**
- Verifica se o arquivo `.env` existe antes de carregá-lo
- Fallback para variáveis de ambiente do sistema
- Mensagens informativas sobre o status do carregamento

#### ✅ **Tratamento de Erros Gracioso**
- Não encerra o programa abruptamente se o `.env` não existir
- Tenta carregar do ambiente do sistema como alternativa
- Fornece feedback claro sobre problemas de configuração

#### ✅ **Validação Completa**
- Verifica se todas as variáveis obrigatórias estão presentes
- Valida se o arquivo de chave privada existe
- Mensagens de erro específicas e orientativas

### 3. Arquivo de Chave Privada

Coloque seu arquivo `.pem` no diretório `cmd/agente/` e configure o caminho no `.env`:

```bash
OCI_KEY_FILE=minha-chave-privada.pem
```

## 🔧 Compilação

```bash
# Navegar para o diretório do projeto
cd /caminho/para/AgenteAI-Git

# Baixar dependências
go mod download

# Compilar o projeto
go build -o cmd/agente/agente.exe ./cmd/agente
```

## 🚀 Execução

```bash
# Navegar para o diretório do executável
cd cmd/agente

# Executar o programa
./agente.exe
```

### 📋 Fluxo de Inicialização

1. **Carregamento de Configuração**: Sistema verifica `.env` e carrega configurações
2. **Validação**: Verifica credenciais e arquivos necessários
3. **Seleção de Modelo**: Interface interativa para escolha do modelo
4. **Início da Sessão**: Sistema pronto para receber perguntas

## 📖 Como Usar

### 🎯 Fluxo Principal

1. **Execute o programa**: `./agente.exe`
2. **Configuração automática**: Sistema carrega configurações do `.env`
3. **Escolha um modelo**: Selecione entre os modelos disponíveis
4. **Inicie a sessão**: Faça múltiplas perguntas com contexto
5. **Use comandos especiais**: Controle avançado da sessão
6. **Veja estatísticas**: Métricas automáticas ao final

### 🧠 Sistema de Contexto

O sistema de contexto é uma funcionalidade avançada que permite:

- **Contexto Automático**: Por padrão, o modelo lembra das perguntas anteriores
- **Controle Manual**: Use `contexto` para ativar/desativar
- **Status Visual**: Feedback sobre o estado do contexto
- ⚡ Melhores respostas com histórico de conversas

#### Comandos de Contexto:
- `contexto` ou `context` → Alternar contexto (ativado/desativado)
- `status` ou `estado` → Ver status atual do contexto

### 🎮 Comandos Especiais

Durante a sessão, você pode usar os seguintes comandos:

| Comando | Aliases | Função |
|---------|---------|---------|
| `sair` | `exit`, `quit`, `tchau`, `fim` | Encerrar sessão com estatísticas |
| `ajuda` | `help`, `?`, `comandos` | Mostrar instruções completas |
| `historico` | `history`, `hist` | Ver histórico completo de perguntas |
| `stats` | `estatisticas`, `statistics` | Ver estatísticas da sessão atual |
| `limpar` | `clear`, `cls` | Limpar tela mantendo contexto |
| `contexto` | `context`, `toggle` | Ativar/desativar contexto |
| `status` | `estado`, `contexto?` | Ver status do contexto atual |
| `trocar` | `modelo`, `change` | Informações sobre troca de modelo |

### 📊 Funcionalidades da Sessão

#### **Sistema de Contexto Inteligente** 🧠
- 💭 Preserva contexto entre perguntas na mesma sessão
- 🔄 Controle manual para ativar/desativar contexto
- 📊 Indicadores visuais do status do contexto
- ⚡ Melhores respostas com histórico de conversas

#### **Histórico de Conversas** 📝
- 📝 Registro completo de todas as perguntas e respostas
- ⏰ Timestamp de cada interação
- ⚡ Tempo de processamento individual
- ✅ Status de sucesso/erro para cada pergunta

#### **Estatísticas em Tempo Real** 📊
- 📈 Taxa de sucesso das perguntas (%)
- ⏱️ Tempo médio de processamento
- 📊 Total de perguntas feitas
- 🕐 Duração da sessão
- 🤖 Modelo utilizado na sessão

### 💡 Exemplo de Uso Completo

```bash
🚀 Oracle AI Generative Agent
=============================
✅ Arquivo .env carregado com sucesso
📋 Configuração OCI carregada:
  • Tenancy ID: ocid1.tenancy.oc1..a...udhfgets
  • User ID: ocid1.user.oc1..aaaa...d8fge5gwf6
  • Key File: minha-chave.pem
  • Fingerprint: xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx
  • Region: sa-saopaulo-1

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
  - 'contexto', 'context' → Ativar/desativar contexto
  - 'status', 'estado' → Ver status do contexto
  - 'trocar', 'modelo' → Informações sobre troca de modelo
• Pressione Enter após cada pergunta
• Para perguntas longas, digite normalmente em uma linha
• 🧠 Contexto: Quando ativado, o modelo lembra das perguntas anteriores
======================================================================

📝 Pergunta 1: Qual é o maior planeta do sistema solar?
🤔 Processando pergunta 1...
🆕 Primeira pergunta da sessão

======================================================================
🤖 Resposta 1 - Meta Llama 3.3 70B Instruct:
⚡ Processado em: 1.234s
======================================================================
O maior planeta do sistema solar é Júpiter.
======================================================================

📝 Pergunta 2: E qual é o menor?
🤔 Processando pergunta 2...
💭 Usando contexto de 1 perguntas anteriores

======================================================================
🤖 Resposta 2 - Meta Llama 3.3 70B Instruct:
⚡ Processado em: 987ms
======================================================================
O menor planeta do sistema solar é Mercúrio.
======================================================================

📝 Pergunta 3: status
📋 Contexto ativado - o modelo lembra das perguntas anteriores
💭 Perguntas no contexto: 2

📝 Pergunta 4: stats
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

📝 Pergunta 5: sair

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
    github.com/joho/godotenv v1.4.0
)
```

## 🏗️ Arquitetura e Padrões

O projeto segue a **estrutura padrão do Go** para melhor organização e manutenibilidade:

### 📁 Organização dos Diretórios

- **`cmd/agente/`** - Contém o ponto de entrada da aplicação (main.go)
- **`internal/domain/`** - Lógica de negócio, modelos e implementações específicas
- **`internal/infrastructure/`** - Configurações, integrações externas e infraestrutura

### ✨ Benefícios desta Estrutura

- **🔒 Encapsulamento**: Código em `internal/` não pode ser importado por outros projetos
- **📦 Modularidade**: Separação clara entre domínio e infraestrutura
- **🔧 Manutenibilidade**: Facilita mudanças e adição de novas funcionalidades
- **📚 Padrão Go**: Segue as convenções estabelecidas pela comunidade Go
- **🎯 Testabilidade**: Estrutura favorece criação de testes unitários

### 🔗 Dependências entre Módulos

```
cmd/agente → internal/domain + internal/infrastructure
internal/domain ← internal/infrastructure
```

## 🔍 Resolução de Problemas

### ❌ Problemas de Configuração

#### Arquivo .env não encontrado
```bash
⚠️  Arquivo .env não encontrado. Tentando carregar variáveis do ambiente do sistema...
❌ Erro na configuração: OCI_TENANCY_ID não encontrado...
```
**Solução**: Crie o arquivo `.env` no diretório `cmd/agente/` com as configurações necessárias.

#### Erro ao carregar .env
```bash
⚠️  Erro ao carregar arquivo .env: ... Tentando carregar variáveis do ambiente do sistema...
```
**Solução**: Verifique a sintaxe do arquivo `.env` e se as permissões estão corretas.

### ❌ Problemas de Autenticação

#### Arquivo PEM não encontrado
```bash
❌ arquivo de chave não encontrado: minha-chave.pem. Verifique se o caminho está correto
```
**Soluções**:
- Verifique se o arquivo `.pem` existe no caminho especificado
- Confirme se o caminho no `.env` está correto
- Verifique as permissões do arquivo

#### Credenciais inválidas
- Confirme se as credenciais OCI estão corretas
- Verifique se o fingerprint corresponde à chave
- Teste a conectividade com a Oracle Cloud

### ❌ Problemas de Modelo

#### Modelo não encontrado
- Verifique se o modelo está disponível na sua região
- Confirme se você tem acesso ao modelo selecionado
- Tente usar um modelo diferente

### ❌ Problemas de Compilação
```bash
# Limpar cache e dependências
go clean -modcache
go mod download
go mod tidy

# Recompilar
go build -o cmd/agente/agente.exe ./cmd/agente
```

### ❌ Sessão Travada
- Use `Ctrl+C` para forçar saída
- Digite `sair` para encerramento normal
- Verifique conexão de rede se perguntas não processam

## 🚀 Melhorias Futuras

### ✅ Implementado
- [x] **Múltiplas perguntas**: Sistema de sessão contínua
- [x] **Histórico de conversas**: Registro completo de interações
- [x] **Estatísticas de sessão**: Métricas em tempo real
- [x] **Configuração via .env**: Sistema robusto de configuração
- [x] **Sistema de contexto**: Contexto entre perguntas
- [x] **Controle de contexto**: Ativar/desativar contexto manualmente

### 🔄 Em Desenvolvimento
- [ ] **Configuração avançada**: Mais opções de personalização via .env
- [ ] **Cache de respostas**: Sistema de cache para otimização
- [ ] **Troca de modelo em tempo real**: Mudar modelo durante a sessão
- [ ] **Persistência de sessão**: Salvar e restaurar sessões

### 🎯 Roadmap Futuro
- [ ] **Exportar histórico**: Salvar conversas em diferentes formatos
- [ ] **Interface web**: Criar interface web para facilitar uso
- [ ] **Logs detalhados**: Sistema de logging mais robusto
- [ ] **Testes automatizados**: Cobertura de testes unitários

### 🔧 Adicionando Novos Modelos

Para adicionar suporte a novos modelos:

1. **Adicione a constante** em `internal/domain/models.go`
2. **Atualize o mapa** `SupportedModels`
3. **Crie implementação específica** se necessário
4. **Atualize as funções** de verificação de família
5. **Teste** o novo modelo

## 📚 Referências

- **Documentação Oracle OCI**: https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/
- **SDK Go Oracle**: https://github.com/oracle/oci-go-sdk
- **Generative AI**: https://docs.oracle.com/iaas/Content/generative-ai/home.htm
- **Godotenv**: https://github.com/joho/godotenv

---