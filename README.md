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
Como se fosse um "npm install". Adiciona/remove automaticamente pacotes que não estão sendo usados ou que estão sem "import"
```bash
go mod tidy
```

### Executar o arquivo principal 
```bash
go run main.go
```

### Gerar o binário final (output)
Se eu acessar o arquivo somente via terminal, já executa, como se fosse um .exe no windows
```bash
go build -o <nome_arquivo>
```

### Expressão curta de If usando variável disponível somente naquele escopo
Variável err só pode ser acessada dentro do If pois só existe naquele escopo, fora dela, é undefined

```bash

func main() {
	if err := thisIsAnError(); err != nil {
		fmt.Println(err.Error())
	}
}

func thisIsAnError() error {
	return errors.New("isto é um erro")
}

```