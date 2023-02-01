
create a new db migration
```bash
$ migrate create -ext sql -dir db/migration -seq <name-file>
```

map all env variables on AWS secrets manager into app.env file
```bash
$ aws secretsmanager get-secret-value --secret-id simple_bank \
 --query SecretString --output text \
 | jq -r 'to_entries|map("(\.key)=\(.value)")|.[]' \
 > app.env
```