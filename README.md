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

### Declaração curta de variável no If (somente disponível naquele escopo) 
Variável err só pode ser acessada dentro do If, pois foi declarada naquele escopo, fora daquela condição, é undefined
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
### Uso do Range 
Acesso ao valor dos elementos com mais facilidade 

1) Usando Slice 
```bash

nums := []string{"pedro", "caio", "juliana"}

for key, value := range nums {
    fmt.Println(key, value)
}
```

2) Usando Map
```bash

users := map[string]string{
    "nome":      "João",
    "documento": "CPF",
}

for key, value := range users {
    fmt.Println(key, value)
}
```

### Uso do package Context 
