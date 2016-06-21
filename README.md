## Install Dependencies
```
go get github.com/tools/godep
godep save
```

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
It was my intention to deploy the chat app to heroku, however it is unclear whether cleardb (heroku's mysql provider) supports go. Trying to work through the bug but may not have time.

