# Gopportunities

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

gopportunities é um projeto feito na linguagem Go para cadastrar e listar oportunidades de emprego. O objetivo é proporcionar uma interface simples e eficiente para gerenciar vagas de emprego.

## Índice

- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [Funcionalidades](#funcionalidades)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Instalação

1. Clone o repositório:

    ```bash
    git clone https://github.com/mtheusvalle/gopportunities.git
    cd gopportunities
    ```

2. Instale as dependências:

    ```bash
    go mod tidy
    ```

3. Compile o projeto:

    ```bash
    go build
    ```

## Como Usar

Após compilar o projeto, você pode executar o binário gerado para iniciar a aplicação:

```bash
./gopportunities
```

A aplicação estará disponível em `http://localhost:8080`.

### Rotas Disponíveis

- `GET /api/v1/openings` - Lista todas as oportunidades de emprego
- `POST /api/v1/opening` - Cria uma nova oportunidade de emprego
- `GET /api/v1/opening?id=?` - Exibe detalhes de uma oportunidade específica
- `PUT /api/v1/opening?id=?` - Atualiza uma oportunidade existente
- `DELETE /api/v1/opening?id=?` - Remove uma oportunidade de emprego

## Funcionalidades

- **Cadastrar Vagas**: Adicione novas oportunidades de emprego.
- **Listar Vagas**: Veja uma lista de todas as vagas cadastradas.
- **Atualizar Vagas**: Edite as informações de vagas existentes.
- **Remover Vagas**: Exclua vagas que não estão mais disponíveis.

## Contribuição

Contribuições são bem-vindas! Se você deseja contribuir com este projeto, siga os passos abaixo:

1. Fork o repositório.
2. Crie uma branch com a nova funcionalidade: `git checkout -b minha-nova-funcionalidade`
3. Faça o commit das suas alterações: `git commit -m 'Adiciona nova funcionalidade'`
4. Envie para a branch principal: `git push origin minha-nova-funcionalidade`
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
