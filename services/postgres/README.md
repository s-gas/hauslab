# postgres

![Postgres](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Docker container running PostgreSQL database.

## Configuration

The database is configured with the following environment variables:
- `POSTGRES_PASSWORD_FILE`: path to the Docker secret file containing the password
- `POSTGRES_USER`: default superuser
- `POSTGRES_DB`: defaul database

## Connecting

From inside the postgres container run:

```bash
psql -U $POSTGRES_USER -d $POSTGRES_DB
```

From another container in the same network run:

```bash
psql -h postgres -U $POSTGRES_USER -d $POSTGRES_DB
```
