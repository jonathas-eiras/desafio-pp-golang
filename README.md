## Desafio-picpay

## Projeto desenvolvido como parte do processo de seleção para vaga de Quality Engineer.

## Projeto

### Implementação de scripts para testes e2e na API Users.


### - O projeto está dividido em duas partes: 
 - Branch e2e-golang com scripts em Golang;
 - Branch e2e-cypress com scripts em Cypress .

## **E2E - Cypress**

**OBS:** Scripts referentes aos testes com Cypress estão localizados na branch ***e2e-cypress***

## - Pré-requisitos

### Ferramentas:
- [Cypress](https://docs.cypress.io/guides/getting-started/installing-cypress)
- [VSCode](https://code.visualstudio.com/download)

### Dependências:
Para execução dos testes, é necessário a instalação das seguintes dependências:

- [Node](https://nodejs.org/en/download/)
- [NPM](https://www.npmjs.com/get-npm)

## - Executando os testes
- Executar o seguinte comando no terminal, dentro da pasta do projeto

`npx cypress run --spec "cypress/integration/services/Users/tests/e2e-test.spec.js"`

## **E2E - Golang**

**OBS:** Scripts referentes aos testes com Golang estão localizados na branch ***e2e-golang***

## - Pré-requisitos

### Ferramentas:
- [Golang](https://golang.org/doc/install)
- [VSCode](https://code.visualstudio.com/download)

## - Executando os testes
- Executar o seguinte comando no terminal, dentro da pasta do projeto

`go test -v tests/users/users_test.go`



