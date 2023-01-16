# A URL shortener made using Go-Fiber and Postgres and Gorm

Starting a docker with Postgres

```bash
docker run --name <NAME> -d -p <PORT_TO ASSIGN>:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=<PASSWORD> postgres
```

To check if it is running:

    docker ps

To get into the psql terminal

    docker exec -it <NAME> /bin/sh

inside psql terminal:

psql -U admin
\dt
select \* from golies;
