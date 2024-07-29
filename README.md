# <center>GO API REST</center>

## **Introdução**

Está é uma pequena API para criação e gerenciamento de viagens usando GO e o Framework GIN para a requisição HTTP e PostgreSQL para o banco de dados.

## **Features**

- Criar, atualizar, e deletar viagens;
- Convidar participants via email;
- Administrar atividades de viagens.

## **Iniciando**

### *Pré-requisitos*

Antes de você começar a usar a API, tenha certeza de ter os requistos em sua máquina.

- Go 1.22 (ou versão mais nova em sua máquina);
- PostgreSQL DataBase
- DBeaver (opcional, ou qualquer gerenciador de DataBases)
- Docker & Docker Compose
- Postman (opcional, ou qualquer app de requisição HTTP)

### *Instalação*

1. **Clone o Repositório**
    ```sh
        git clone https://github.com/Turgho/go-api-rest.git
        cd go-api-rest
    ```

2. **Iniciar a API com Docker**
    ```sh
        docker build -t go-api-rest .
        docker compose up -d
    ```

## *API Endpoints*

O servidor local será criado na porta `http://localhost:5050/`.

### Viagens

- `POST - /api/`
        Cria uma viagem.
    ```json
        //Request   
        {
        "destination": "New York",
        "start_date": "2024-10-10",
        "end_date": "2024-10-20",
        "owner_name": "John Doe",
        "owner_email": "john.doe@example.com"
        }

        // Response
        {
            "id": "uuid"
        }
    ```

- `GET - /api/trips/:id`
        Retorna uma viagem já criada.
    ```json
        // Response
        {
            "id": "uuid",
            "destination": "New York",
            "start_date": "2024-10-10",
            "end_date": "2024-10-20",
            "owner_name": "John Doe",
            "owner_email": "john.doe@example.com",
            "status": 0
        }
    ```
- `PUT - api/trips/:id`
        Confirma uma viagem.
    ```json
        // Response
        {
            "id": "uuid",
            "status": 1
        }
    ```
    
### Participantes

- `GET - /api/trips/:id/invite`
        Convida um participante
    ```json
        // Request
        {
            "names": [
                "João",
                "Marcos"
            ],
            "emails": [
                "joao@gmail.com",
                "marcos@gmail.com"
            ]
        }

        // Response
        {
            "message": "Participantes convidados"
        }
    ```

- `GET - /api/participants/:tripID`
        Retorna participantes de uma viagem
    ```json
        // Response
        [
            {
                "id": "uuid",
                "name": "João",
                "email": "joao@gmail.com",
                "is_confirmed": 0
            },
            {
                "id": "uuid",
                "name": "Marcos",
                "email": "marcos@gmail.com",
                "is_confirmed": 0
            }
        ]
    ```

- `PUT - /api/participants/:tripID/:participantID`
        Confirma um participante de uma viagem
    ```json
        // Response
        {
            "message": "Participante(s) confirmado."
        }
    ```

### Links

- `POST - /api/trips/:id/links`
    Cadastra links de uma viagem
    ```json
        // Request
        {
            "urls": [
                "www.hotel-1.com",
                "www.hotel-2.com"
            ],
            "titles": [
                "Hotel 1",
                "Hotel 2"
            ]
        }

        // Response
        {
            "message": "Links registrados."
        }
    ```

- `GET - /api/trips/:id/links`
        Retorna os links de uma viagem
    ```json
        // Response
        [
            {
                "id": "uuid",
                "trip_id": "uuid",
                "link": "www.hotel-1.com",
                "title": "Hotel 1"
            },
            {
                "id": "uuid",
                "trip_id": "uuid",
                "link": "www.hotel-2.com",
                "title": "Hotel 2"
            }
        ]
    ```

### Atividades

- `POST - /api/trips/:id/activities`
        Cria atividades de uma viagem
    ```json
        // Resquest
        {
            "title": "Comprar itens",
            "occurs_at": "2024-10-11 13:30:00"
        }

        // Response
        {
            "activity_id": "uuid"
        }
    ```

- `GET - /api/trips/:id/activities`
        Retorna atividades de uma viagem
    ```json
        // Response
        [
            {
                "id": "uuid",
                "title": "Comprar itens",
                "occurs_at": "2024-10-11 13:30:00"
            }
        ]
    ```

## *Observações*

Está é uma API apenas para propósitos de estudo, e ainda está em desenvolvimento.

## *License*

Este `README.md` fornece uma breve visão geral da API, incluindo seus recursos, etapas de instalação e um resumo dos principais endpoints. Ajuste os espaços reservados (como URL do repositório e credenciais do banco de dados) de acordo com os detalhes reais do seu projeto.
