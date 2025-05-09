# go-studies

Repositório para estudos e experimentações com a linguagem Go (Golang)

## Comandos iniciais 

### Iniciar projeto
Cria um go.mod que é como se fosse um package.json
```bash
go mod init <caminho_repo_git>

go mod init github.com/Julianagagliano7/go-studies
```

### Baixar dependências 
Como se fosse um "npm install"
Adiciona/remove automaticamente pacotes que não estão sendo usados ou que estão sem "import"
```bash
go mod tidy
```

### Executar o arquivo principal 
```bash
go run main.go
```

### Gerar o binário final (output)
```bash
go build -o <nome_arquivo>
```