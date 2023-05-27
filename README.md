# Support

ticketing API

## How to run the API

1. Start the containers

```
docker compose up
```

2. Migrate the database

   Make sure [goose](https://github.com/pressly/goose) can connect to the db.

   ```
   goose postgres "host=db user=support password=support dbname=supportdb sslmode=disable" status
   ```

   Then run the db migration

   ```
   goose postgres "host=localhost user=support password=support dbname=supportdb sslmode=disable" up
   ```

3. Test the API

   Try to get the healthcheck endpoint

   ```
   curl localhost:8080/healthcheck
   ```
