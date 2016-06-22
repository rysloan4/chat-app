## SetupDB

```
mysql --user=root mysql
create database chat;
CREATE USER 'chat'@'localhost' IDENTIFIED BY 'chat';
grant select,insert,update,delete,create,drop on chat.* to chat;
```

then

```
cd migrations
mysql chat < create_message_table.sql
mysql chat < create_user_table.sql
```

## Run Chat App Locally
```
make run-chat-dev
```

## Development

Test
```
make test
```
Lint
```
make lint
```
Format
```
make fmt
```
Build
```
make build
```

## Production
[https://nameless-savannah-87003.herokuapp.com/login](https://nameless-savannah-87003.herokuapp.com/login0) to login or create a username

Deploy

```
git push heroku master
```
Logs
```
heroku logs
```
Production DB
```
mysql heroku_7dda9dbd4cbc075 --host=us-cdbr-iron-east-04.cleardb.net --user=b3fd3325d24b40 --password=2761ce0f
```

## Next Steps

- Better test coverage and integration tests
- Make Hub an interface so implementation can be swapped out
- Move ui to standalone app using react (general design and ui improvements included)
