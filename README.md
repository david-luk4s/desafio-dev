# Desafio programação - para vaga desenvolvedor

### Pré-requisitos

Instale [Docker Compose plugin](https://docs.docker.com/compose/install/) ou [Docker Compose standalone](https://docs.docker.com/compose/install/other/)

execute o programa com Compose plugin

```
docker-compose up
```


ou execute o programa com Compose standalone

```
docker compose up
```

### Teste Unitarios
Com Docker Compose plugin:
```
docker-compose exec -it app go test -v
```

Com Docker Compose standalone:
```
docker compose exec -it app go test -v
```

### Uso via navegador de sua preferência
```
http://127.0.0.1:8080/
```


### Uso via Curl
- endpoint upload
```
curl --location --request POST 'http://127.0.0.1:8080/api/upload' \
--form 'myfile=@"/home/david-luk4s/Downloads/CNAB.txt"'
```

- endpoint de listagem
```
curl --location --request GET 'http://127.0.0.1:8080/api/list'
```

## Author

* **David Lucas** - *Linkedin* - [david-luk4s](https://www.linkedin.com/in/david-lucas-souz4/)