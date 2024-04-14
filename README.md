# Documentação do Projeto

Este projeto é um sistema de cadastro de produtos desenvolvido em Go utilizando o framework Echo e o ORM GORM.

## Estrutura do Projeto

O projeto está dividido em diferentes pacotes:

- **handler**: Contém os manipuladores HTTP para os endpoints da API.
- **model**: Define os modelos de dados da aplicação.
- **repository**: Implementa as operações de acesso ao banco de dados.
- **service**: Contém a lógica de negócios da aplicação.

## Pacotes

### `handler`

Este pacote contém o manipulador HTTP para lidar com as requisições da API.

- **UserHandler**: Manipulador para as operações relacionadas aos produtos.

### `model`

Neste pacote estão definidos os modelos de dados da aplicação.

- **Produtos**: Estrutura que representa um produto com os seguintes campos:
  - `ID`: Identificador único do produto.
  - `Nome`: Nome do produto.
  - `Descricao`: Descrição do produto.
  - `Preco`: Preço do produto.

### `repository`

O pacote repository implementa as operações de acesso ao banco de dados.

- **Repository**: Responsável por interagir com o banco de dados e executar as operações CRUD.

### `service`

Contém a lógica de negócios da aplicação.

- **UserService**: Contém os métodos para manipular os produtos.

## Configuração do Banco de Dados
O projeto utiliza o MySQL como banco de dados. A conexão é configurada no arquivo main.go.

#### Como Executar o Projeto
- Certifique-se de ter o Go instalado em sua máquina.
- Clone o repositório do projeto.
- Instale as dependências com o comando go mod tidy.
- Configure o banco de dados conforme necessário.
- Execute o comando go run main.go para iniciar o servidor.
- A API estará disponível em http://localhost:8080

## Endpoints da API

### `POST /create`

Cria um novo produto com os dados fornecidos no corpo da requisição.

#### Parâmetros da Requisição

- **Corpo da Requisição**:
  - `Nome`: Nome do produto (string).
  - `Descricao`: Descrição do produto (string).
  - `Preco`: Preço do produto (float64).

#### Exemplo de Requisição

```json
{
  "Nome": "Produto A",
  "Descricao": "Descrição do Produto A",
  "Preco": 29.99
} 

