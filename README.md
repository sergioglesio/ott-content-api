# OTT Content Metadata API

## Descrição

O **OTT Content Metadata API** é uma API desenvolvida em Go, projetada para gerenciar e fornecer informações sobre conteúdos de vídeo, como filmes e séries, em plataformas de streaming. Este projeto permite a criação, leitura, atualização e exclusão (CRUD) de metadados relacionados a vídeos, facilitando a integração de dados em aplicações de streaming.

## Funcionalidades

- **CRUD de Vídeos**: Gerencie vídeos com operações de criar, ler, atualizar e excluir.
- **Estrutura de Dados**: Cada vídeo possui campos essenciais:
  - Título
  - Descrição
  - Ano de Lançamento
  - Diretor
  - Estrelas
  - Gênero
  - URL da Miniatura
  - URL do Vídeo
  - Duração
  - Avaliação
- **Banco de Dados**: Armazena dados de vídeos em um banco de dados MySQL.
- **Serviço Nginx**: Utiliza Nginx como servidor proxy reverso.
- **Containerização**: Utiliza Docker e Docker Compose para facilitar a implantação e o gerenciamento.

## Tecnologias Usadas

- **Linguagem**: Go (Golang)
- **Servidor Web**: Nginx
- **Banco de Dados**: MySQL
- **Containerização**: Docker, Docker Compose

## Estrutura do Projeto

```
ott-content-api/
├── cmd/
│   └── main.go           # Ponto de entrada da aplicação
├── init-scripts/         # Scripts de inicialização do banco de dados
├── nginx/
│   └── default.conf      # Configuração do Nginx
├── repositories/         # Repositórios de dados
├── .env                  # Variáveis de ambiente
├── docker-compose.yml    # Configuração do Docker Compose
├── Dockerfile            # Dockerfile da aplicação
└── Dockerfile.nginx      # Dockerfile do Nginx
```

## Como Executar o Projeto

### Pré-requisitos

- [Docker](https://www.docker.com/get-started) e [Docker Compose](https://docs.docker.com/compose/install/) instalados.

### Passos

1. Clone o repositório:
   ```bash
   git clone https://github.com/sergioglesio/ott-content-api.git
   cd ott-content-api
   ```

2. Crie um arquivo `.env` com as seguintes variáveis:
   ```plaintext
   MYSQL_ROOT_PASSWORD=XXXXXXXXXXXXX
   MYSQL_DATABASE=content
   MYSQL_USER=XXXXXXXXXXXXX
   MYSQL_PASSWORD=XXXXXXXXXXXXX
   ```

3. Execute os contêineres:
   ```bash
   docker-compose up --build
   ```

4. Acesse a API:
   - **Endpoints disponíveis**:
     - `POST /videos`: Criar um novo vídeo.
     - `GET /videos`: Obter todos os vídeos.
     - `GET /videos/{id}`: Obter um vídeo específico por ID.
     - `PUT /videos/{id}`: Atualizar um vídeo específico por ID.
     - `DELETE /videos/{id}`: Deletar um vídeo específico por ID.

## Contribuições e contato

Contribuições são bem-vindas! Sinta-se à vontade para abrir um issue ou enviar um pull request.

- **Nome:** Sérgio Glésio
- **E-mail:** sergioglesiojunior@outlook.com.br
- **GitHub:** [sergioglesio](https://github.com/sergioglesio)
- **LinkedIn:** [Sérgio Glésio](https://www.linkedin.com/in/sergioglesio/)

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
