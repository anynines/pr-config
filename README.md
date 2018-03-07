# pr-config

Web service to store and retrieve JSON information for (GitHub) Pull Requests.

# API

## GET /v1/org/repo/pr-id

```
curl -u admin:password -X GET localhost:4455/v1/anynines/cool-project/13
```

```json
{
  "test-backup": true,
  "cache": false
}
```

## POST /v1/org/repo/pr-id

```
curl -u admin:password -X POST -H "Content-Type: application/json" localhost:4455/v1/anynines/cool-project/13 -d '{ "test-backup": true, "cache": false }'
```
