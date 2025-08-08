docker image ls

docker image pull ..:..

docker image rm ..:..

# melihat semua container
docker container ls -a

# membuat kontainer
docker container create --name contohredis redis:latest

# menghapus container
docker container rm contohredis

# menjalankan container
docker container start namecontainer/idcontainer

# menghentikan container
docker container stop namecontainer/idcontainer

# melihat log container
docker container logs namecontainer/idcontainer

# excute container
docker container exec -it namecontainer/idcontainer /bin/bash

# melakukan Port Forwarding
docker container create --name contohredis -p 8080:80 nginx:latest