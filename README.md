# pr-config

Web service to store & retrieve JSON information for (GitHub) pull requests.

```shell
# setup redis
docker run -d -p 6379:6379 redis redis-server --requirepass password

# setup env variables
export PORT=4455 # default value
export HTTP_USERNAME=admin
export HTTP_PASSWORD=password
export BACKEND_URL=redis://admin:password@localhost

# start pr-config
go run main.go
```

# Usage

```shell
# set JSON data for anynines/project PR #1
$ curl -v -u admin:password -X POST localhost:4455/v1/anynines/project/1 -d '{"unit-test": "true"}'
# get JSON data for anynines/project PR #1
$ curl -v -u admin:password -X GET localhost:4455/v1/anynines/project/1
{"unit-test": "true"}
```

# API

## GET /v1/:org/:repo/:pr-id

```
curl -u admin:password -X GET localhost:4455/v1/anynines/cool-project/13
```

```json
{
  "test-backup": true,
  "cache": false
}
```

## POST /v1/:org/:repo/:pr-id

```
curl -u admin:password -X POST -H "Content-Type: application/json" localhost:4455/v1/anynines/cool-project/13 -d '{ "test-backup": true, "cache": false }'
```
