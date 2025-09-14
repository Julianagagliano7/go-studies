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

### context.WithCancel()
Recebe um sinal para parar imediatamente a execução 

```bash
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go printUntilCancel(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func printUntilCancel(ctx context.Context) {
	count := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel signal received, exiting")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("printing until cancel, number = %d \n", count)
			count += 1
		}
	}
}
```

### context.WithDeadline()
Passo um horário exato em que desejo interromper a execução

```bash
func main() {
	ctx, _ := context.WithDeadline(
		context.Background(),
		time.Now().Add(5*time.Second),
	)

	printUntilCancel(ctx)
}

func printUntilCancel(ctx context.Context) {
	count := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time to stop operation has come, exiting")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("printing until cancel, number = %d \n", count)
			count += 1
		}
	}
}
```

### context.WithTimeout()
Passo daqui a quanto tempo a execução deve parar 

```bash
func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel() //não faz diferença, é só para não dar erro no cancel

	printUntilCancel(ctx)
}

func printUntilCancel(ctx context.Context) {
	count := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout reached, exiting")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("printing until cancel, number = %d \n", count)
			count += 1
		}
	}
}
```
### context.WithValue()
Armazena valores (chave-valor) em memória que podem ser passados para outros processos. Podem ser float, string, ...

```bash

func main() {
	ctx := context.WithValue(
		context.Background(),
		"testKey", "testValue",
	)

	printUntilCancel(ctx)
}

func printUntilCancel(ctx context.Context) {
	fmt.Println(ctx.Value("testKey"))

	//alterando o valor do contexto criando outro contexto e passando o existente 
	ctx2 := context.WithValue(ctx, "testKey", "juliana")
	fmt.Println(ctx2.Value("testKey"))
}
```

### Closure Function
Função anônima 

```bash
func main(){
	multiplica := func(x int) int {
		return x * 2
	}

	resultado := multplica(5)
	fmt.Println(resultado)  //output: 10
}
```
### Goroutine
Ordem dos processos não é garantida, pode ser que um processo mais demorado termine antes de um mais rápido. 
Olhar branch go-routine-2 para mais detalhes.
