# UNT

Descrição do Projeto

## Requerimentos

- [x] Go Versão 1.18 superior. https://go.dev/dl/

- [x] Docker. https://www.docker.com/

- [x] Postman https://www.postman.com/ ou programa de prefencia para testar os ends-points.



### Infraestutura / Banco de Dados

```
docker-compose up -d

```

### Rodando o projeto


Compilando o projeto

```

go build

```

Executando a aplicação

```

./utest (linux / Mac(os)) 
test.exe (Windows)

```

Testes e cobertura

```

 go test ./... -coverprofile=coverage.out

```

### Exemplos das requisições e respostas

http://localhost:3000/api/swagger/


#### Log da aplicação

O log será salvo automaticamente na pasta com o nome info.log. Estruturado em duas categorias (Info / Error).

#### Info: Descrição detalhada sobre um ponto relevante.
#### Error: Erro ou indisponibilidade de atendimento da requisição.

#### Exemplo:

2022-08-29T20:11:34.344-0300	INFO	shared/log.go:28	Dados importados   <br />
2022-08-29T20:11:34.345-0300	INFO	shared/log.go:28	Servidor disponivel na porta 3000  <br />
2022-08-29T20:11:55.773-0300	ERROR	shared/log.go:33	registro não encontrado  <br />