# Currency upvote

## Bem-vindo(a)!
- Projeto proposto com o intuito de contabilizar "Upvotes" e "Downvotes" à cripto moedas.

## Importando a aplicação
Clone o repositório em sua máquina local utilizando:
```
git clone https://github.com/IIGabriel/Upvote-crypto-currency.git
```
Para baixar as dependencias do projeto execute no terminal:
```
go mod tidy
```

## Configurações do projeto
É necessário a conexão com um banco de dados **PortgresSQL** para a persistência de dados e funcionamento da aplicação.

- Arquivo ```.env```

Esse arquivo contém os campos personalizados da aplicação onde deve ser definido:
```
- listen_port = "<Porta da aplicação>"
- psql_settings = "host=<Suas definicoes> user=<Suas definicoes> password=<Suas definicoes> dbname=<Suas definicoes> port=<Suas definicoes> sslmode=disable"
- coingecko_token = "<Token de acesso resgatado em coingecko >"
- main_currency = "<SIMBOLO PADRAO DA MOEDA DE COMPARACAO>"
```
- O Token referente a coingecko_token deve ser adquirido em: https://rapidapi.com/coingecko/api/coingecko

## Como utilizar

- Rota utilizada para criar um Upvote (POST)
```
/upvote/:currency_name
```

- Rota utilizada para criar um Downvote (POST)
```
/downvote/:currency_name
```

Exemplo:
```
localhost:7777/upvote/klever
```

### Rota que retorna as informações da moeda (GET)
```
/currency/:currency_name
```
Essa rota retornará as informações da moeda juntamente com os seus valores e votos

- Query Params personalizáveis:
```
initial_date=YYYY-MM-DD
final_date=YYYY-MM-DD
```
São utilizados para definir o range de dias dos preços da moeda.
Por **Padrão**, caso o 'initial_date' não tenha sido declarado é selecionado o range de 1 ano antes da data final.

- Caso não deseje as informações dos precos da moeda e busque uma requisição mais rápida utilize:
```
omit_price=true
```

- Para selecionar a moeda de comparação utilize:
```
compare_currency=<SIMBOLO DA MOEDA DE COMPARACAO>
```

**Exemplo** com as querys:
```
localhost:7777/upvote/klever?initial_date=2022-09-29&final_date=2022-10-01&compare_currency=BTC&omit_price=true
```

### Rota para criação de uma nova moeda (POST)
```
/currency
```

Header necessário para obter permissão de utilização da rota:
```
Key: Permission_token
Value: <Token de acesso resgatado em coingecko >
```

Deve ser passado no body da requisição um JSON contendo as informações da moeda,
sendo todos os campos obrigatórios
Exemplo:
```
{
    "id": "testeId",
    "name": "testeName",
    "symbol": "TNM"
}
```

### Rota para atualizar moeda ja existente (PUT)
```
/currency/:currency_name
```

Header necessário para obter permissão de utilização da rota:
```
Key: Permission_token
Value: <Token de acesso resgatado em coingecko >
```
Deve ser passado no body da requisição um JSON contendo as novas informações da moeda
Exemplo:
```
{
    "id": "testeId",
    "name": "testeName",
    "symbol": "TNM"
}
```
### Rota para deletar uma moeda ja existente (DELETE)
```
/currency/:currency_name
```

Header necessário para obter permissão de utilização da rota:
```
Key: Permission_token
Value: <Token de acesso resgatado em coingecko >
```

## API externa utilizada
https://rapidapi.com/coingecko/api/coingecko

## Testes
- Para rodar os testes utilize o seguinte comando na raiz do seu projeto
```
go test ./tests
```
