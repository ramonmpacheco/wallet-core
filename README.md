# WALLET-CORE

### PARA RECRIAR A IMAGEM
>```docker-compose up --build```

ControlCenter
> http://localhost:9021

### PARA RODAR NO MODO PROD, USE O COMANDO:
> o comando abaixo subirá todo o ambiente, pronto para usar localmente.

> ```CTX=prod docker-compose up --build```

<br />

### HOW TO DO A CLEAN RESTART OF A DOCKER INSTANCE

1. Stop the container(s) using the following command:
> ```docker-compose down```

2. Delete all containers using the following command:
> ```docker rm -f $(docker ps -a -q)```

3. Delete all volumes using the following command:
> ```docker volume rm $(docker volume ls -q)```