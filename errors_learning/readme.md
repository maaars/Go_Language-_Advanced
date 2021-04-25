init DB
```cassandraql
cd model/

reform-db --db-driver --db-source "root:@tcp(127.0.0.1:3306)/demo" init
cd ..
reform model/
```