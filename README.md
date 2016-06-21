```
mysql --user=root mysql
create database chat;
CREATE USER 'chat'@'localhost' IDENTIFIED BY 'chat';
grant select,insert,update,delete,create,drop on chat.* to chat;
```

mysql heroku_7dda9dbd4cbc075 --host=us-cdbr-iron-east-04.cleardb.net --user=b3fd3325d24b40 --password=2761ce0f
