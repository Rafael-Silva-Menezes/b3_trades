# Aplicação B3 Trades

## Visão Geral

A aplicação B3 Trades foi desenvolvida para processar arquivos de dados de negociações financeiras da B3 (Bolsa de Valores do Brasil) 
e inseri-los em um banco de dados PostgreSQL. E disponobiliza-los via api, permitindo a consulta dos seguintes dados:
```json
{
   "ticker": "PETR4" ,
   "max_range_value": 0, 
   "max_daily_volume": 0, 
   "QuantidadeNegociada": 1000 
}
```

## Funcionalidades

A aplicação realiza as seguintes tarefas:

- Lê arquivos de dados de negociações financeiras de um diretório especificado.
- Faz o parsing e processamento dos dados em objetos estruturados do tipo `Trade`.
- Insere lotes de dados do tipo `Trade` no banco de dados PostgreSQL.
- Suporta concorrência para otimizar o processamento de arquivos usando goroutines.
- Limpa a tabela do banco de dados antes de inserir novos dados.

```
type Trade struct {
	ID              string    // ID único do trade
	Ticker          string    // Código do instrumento
	TradePrice      float64   // Preço da negociação
	TradedQuantity  int       // Quantidade negociada
	ClosingTime     string    // Horário de fechamento da negociação (formato string)
	TradeDate       time.Time // Data da negociação
}
```

## Como Funciona

1. **Configuração Inicial**:
    - Clone o repositório e navegue até o diretório do projeto.
    - Verifique se o Go está instalado (`go version` deve retornar uma versão válida).
    - Instale o PostgreSQL e Docker, se ainda não estiverem instalados.

2. **Preparação dos Arquivos de Dados**:
    - Baixe os arquivos de dados de negociações financeiras do site oficial da [B3](https://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/cotacoes/).
    - Salve os arquivos baixados em um diretório de sua escolha.
    - Lembre-se de anotar o caminho do diretório para colocar no DIRECTORY_PATH

3. **Configuração do Ambiente**:
    - Crie um arquivo `.env` na raíz do projeto, para guardar as informações do banco de dados
      ```dotenv
      # Configurações do PostgreSQL
      POSTGRES_USER=seu_usuario_postgres
      POSTGRES_PASSWORD=sua_senha_postgres
      POSTGRES_DB=seu_banco_postgres
      POSTGRES_HOST=db
      POSTGRES_PORT=5432
      ```
    - Crie um arquivo `.env` no b3-insert e preencha da seguinte forma.
    - Substitua `/caminho/para/seus/arquivos/b3` pelo caminho real onde você salvou os arquivos da B3.
   ```dotenv
      # Caminho do diretório dos arquivos
      DIRECTORY_PATH=/caminho/para/seus/arquivos/b3
      # Número máximo de workers (processamento concorrente)
      MAX_WORKERS=20
      # URL do Banco de Dados PostgreSQL
      DATABASE_URL=postgres://user:pass@host:$5432/postgres?sslmode=disable
   ```
 - Crie um arquivo `.env` no b3-api e preencha da seguinte forma.
   ```dotenv
      # Configuração da Api
      APP_PORT=8080
      # URL do Banco de Dados PostgreSQL
      DATABASE_URL=postgres://user:pass@host:$5432/postgres?sslmode=disable
   ```

4. **Configuração do Docker**:
    - Certifique-se de que o Docker está em execução no seu sistema.
    - Use o arquivo `docker-compose.yml` fornecido para inicializar o banco de dados PostgreSQL.
      ```bash
      docker compose up
      ```
5. **Execução do B3-Insert**:
    - Compile e execute a aplicação `b3-insert`:
      ```bash
      go build -o b3-insert ./cmd
      ./b3-insert
      ```
    - Para executar sem compilar, basta o seguinte comando:
   ```bash
      go run ./cmd/main.go
      ```
    - A aplicação começará a processar os arquivos do diretório especificado, inserindo os dados no banco de dados PostgreSQL.


6. **Execução do B3-Api**:
  - Compile e execute a aplicação `b3-api`:
      ```bash
      go build -o b3-api ./cmd
      ./b3-insert
      ```
- Para executar sem compilar, basta o seguinte comando:
   ```bash
      go run ./cmd/main.go
   ```
- Uma vez que `b3-api` esteja em execução (via Docker), você pode acessá-lo através de `http://localhost:8080` (considerando a configuração padrão).

- Exemplo de requisição curl:
```bash
curl -X GET \                              
  'http://localhost:8000/api/aggregated-data/PETR4?date=2024-06-28' \
  -H 'Content-Type: application/json'
```

7. **Parando a Aplicação**:
    - Para parar o banco de dados use o comando `docker compose down`
    - Para parar a aplicação, utilize `Ctrl + C` no terminal onde ela está sendo executada.

## Notas Adicionais

- Certifique-se de ter permissões suficientes e espaço em disco para as operações de banco de dados e processamento de dados.
- Ajuste o valor de `MAX_WORKERS` no arquivo `.env` de acordo com as capacidades do seu sistema e requisitos de desempenho.
