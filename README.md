```
mysql --user=root mysql
create database chat;
CREATE USER 'chat'@'localhost' IDENTIFIED BY 'chat';
grant select,insert,update,delete,create,drop on chat.* to chat;
```
