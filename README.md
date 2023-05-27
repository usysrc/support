# Support

a support ticket API

## How to start

Start the containers

```
docker compose up
```

Make sure goose can connect to the db.

```
goose postgres "host=db user=support password=support dbname=supportdb sslmode=disable" status
```

Then run the db migration

```
goose postgres "host=localhost user=support password=support dbname=supportdb sslmode=disable" up
```

Try to get the healthcheck endpoint

```
curl localhost:8080/healthcheck
```
