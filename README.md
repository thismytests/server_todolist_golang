<p align="center">
  <a href="http://nestjs.com/" target="blank"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png" width="320" alt="Nest Logo" /></a>
</p>

## Description

Golang prototype api

## Dependencies



* Installed [docker](https://docs.docker.com/compose/install/)

* Installed [docker-compose](https://docs.docker.com/compose/install/)

* Installed [db connector](https://dbeaver.io/download/)

* Installed [postman](https://www.postman.com/downloads/)


## Check dependencies
* docker
  ``` 
  docker -v

* docker-compose
  ``` 
  docker-compose -v


# Start APP and DB in container
```bash
  cd docker-compose
  docker-compose  --env-file ../env/.env.dev  -f docker-compose.yml down
  docker-compose  --env-file ../env/.env.dev  -f docker-compose.yml up
```

# Init DB

#### Using backup

* Make import to the db from poweshell for windows in the root folder
  cat [DumpName].sql | docker exec -i docker-compose_db_1 /usr/bin/psql -h localhost -d eat_beat -U user password=password


*       cat prod_db_dump.sql | docker exec -i docker-compose_db_1 /usr/bin/psql -h localhost -d eat_beat -U user password=password



# Show all dabases
docker exec -i docker-compose_db_1 /usr/bin/psql -h localhost -d  golang-prototype -U user
SELECT datname FROM pg_database;



# Tools
## DBeaver
* **Open DBeaver**

* **In the toolbar** `File>Import DBeaver>Project`

* Press Button `Next`

* Import `File` tools/DBeaver/go-prototype.dbp

* Choose checkbox at the `Projects`

* Press button `Finish`

* **In the toolbar** `Window>Projects`

* Choose project in the left sidebar


## Postman

* **Open DBeaver**

* **In the toolbar** `File>Import>File`

* Press button **Upload Files** and choose file tools/postman/**go-sever-prototype.postman_collection.json**



* **In the toolbar** `File>Import>File`

* Press button **Upload Files** and choose file tools/postman/**go-sever-prototype.postman_environment.json**