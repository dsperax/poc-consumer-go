# poc-consumer-go

Golang consumer - Kafka to PG SQL database.

## Overview

(EN) The project consists of a consumer application in Golang that communicates a messaging system (kafka) with an SQL database (postgres). The purpose of the system is to replicate non-sensitive data to the database that the sales team has access to.

## How to use

(EN) Download the project. It is necessary to have docker and golang installed on the machine. It is recommended to have some IDE to access the database and kafka (DBeaver and offset explorer, for example);

In the project folder, to download the images and prepare the environment (kafka and postgres) run the command:

    docker compose up -d

In the ``poc-consumer-go\docs\postgres`` folder there is the postgres table creation script and in ``poc-consumer-go\docs\kafka`` some JSON models to generate the kafka message.

Finally, in the ``poc-consumer-go\app`` folder, run the following command to upload the consumer:

    go run .\main.go

Finally, just trigger the message in the pattern prepared via kafka and perform the select on the database to see the saved data. It is worth mentioning that the project has methods of insert, update (in case of change in the customer's data) and delete (in case the status changes to "I").

<hr>

## Visão geral

(PT-BR) O projeto consiste em uma aplicação de consumer em Golang que comunica um sistema de messageria (kafka) com um banco de dados SQL (postgres). O intuito do sistema é replicar dados não sensíveis para o banco em que o time de vendas tem acesso.

## Como usar

(PT-BR) Faça o download do projeto. É necessário possuir o docker e golang instalado na máquina. É recomendado possuir alguma IDE para acessar o banco de dados e o kafka (DBeaver e offset explorer, por exemplo);

Na pasta do projeto, para baixar as imagens e preparar o ambiente (kafka e postgres) execute o comando:

    docker compose up -d

Na pasta ``poc-consumer-go\docs\postgres`` tem o script de criação da tabela no postgres e em ``poc-consumer-go\docs\kafka`` alguns modelos de JSON para gerar a mensagem do kafka.

Por último, na pasta ``poc-consumer-go\app``, execute o seguinte comando para subir o consumer:

    go run .\main.go

Por ultimo, basta disparar a mensagem no padrão da preparada via kafka e realizar o select no banco para ver os dados salvos. Vale ressaltar que o projeto conta com métodos de insert, update (em caso de alteração nos dados do cliente) e delete (caso o status mude para "I").

<hr>

## Evidences

![General archtecture](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/Archtecture-diagram-call-center-system.drawio.png)
![compose exec](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/docker-compose-up.png)
![create table](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/create-table.png)
![kafka on](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/zookeeper-on.png)
![aplication up](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/aplication-up.png)
![message](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/message.png)
![message evidence](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/message-evidence.png)
![registro processado](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/registro-processado.png)
![insert db](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/update-db.png)
![update db](https://github.com/dsperax/poc-consumer-go/blob/main/documentation/insert-db.png)
