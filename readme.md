# Query node deps

This parses the package.json and yarn.lock of all packages at Convoy, recording their dependencies in a postgres database.


## Setup

First, get [golang](https://go.dev/). 

Run these commands:


```bash
export GITHUB_ACCESS_TOKEN=$(printf "host=github.com\nprotocol=https\n" | git credential-osxkeychain get | grep "password" | cut -d= -f2)
cd db-image && ./build-it.sh
./run-it.sh
cd ../
go run cmd/migrate/main.go
yarn
```

## Usage

```bash
go run cmd/parseall/main.go
```

This will attempt populate the db with the parsed lockfile and package.json info. Then you can connect to the dockerized postgres to query around.

```
psql -U postgres -d postgres -h localhost -p 5432
```

## Queries

```sql
-- find repos that depend directly on moment
select distinct fully_qualified_git_slug from dependencies where name='moment' and source='PACKAGE_JSON';
```

## Known Limitations

This will not properly parse yarn berry lockfiles (see yarn.go)