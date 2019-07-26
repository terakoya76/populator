# popuolator

[![CircleCI](https://circleci.com/gh/terakoya76/populator/tree/master.svg?style=svg)](https://circleci.com/gh/terakoya76/populator/tree/master)

The populator is a tool making your database populated via yaml config inputs.

The only thing you have to do is, describing database schema and number of records you want into yaml file, then run this tool designating that yaml.

This is useful when you want to inspect SQL query itself. You don't need to write up any stored procedures populating your database schema.

On the other hand, this tool not enough for inspecting Application Servers' performance, cuz this tool doesn't care about application layer's validation at all.

## Install
```shell
$ go get github.com/terakoya76/populator
```

## How to use
You need to prepare config file, config file is like below.

```yaml
driver: mysql
database:
  host: 127.0.0.1
  user: root
  password: root
  name: testdb
  port: 3306
tables:
- name: table_a
  columns:
    - name: col_1
      type: bigint
    - name: col_2
      type: varchar
  charset: utf8mb4
  record: 100000
```

Then, run this command.

```shell
$ populator -c ./path/to/configfile.yaml
```

Everything is done.

```shell
mysql> use testdb;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> show tables;
+------------------+
| Tables_in_testdb |
+------------------+
| table_a          |
+------------------+
1 rows in set (0.01 sec)

mysql> show create table table_a;
+---------+-----------------------------------------------------------------------------------------------------------------------------------------+
| Table   | Create Table                                                                                                                            |
+---------+-----------------------------------------------------------------------------------------------------------------------------------------+
| table_a | CREATE TABLE `table_a` (
  `col_1` bigint(20) DEFAULT NULL,
  `col_2` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 |
+---------+-----------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)

mysql> select count(*) from table_a;
+----------+
| count(*) |
+----------+
|   100000 |
+----------+
1 row in set (0.03 sec)

mysql> select * from table_a limit 5;
+----------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| col_1                | col_2                                                                                                                                                                                                                                                           |
+----------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
|  7889812248084406655 | RmYktDtHZczbJeuxpKcHCwVfLnrTHgKcOrJVOwmVwLwwlTFVKGxOQvcBbPKQplHnUJtWzxYphwWxydtBQzRdUhbEJxyurXQCCgzmEWKdRLbWqLsZXapbfiugWMCbXzXwkIKciRASZSEmRUccIeUGOHoFRcMZDIDtFSNnqyWxuCxgYjpFkVwxGDQnqiXZeArHNDfPmgzcbXDjqzTvpfFtXtDJZYludhfmoaEYliAafbEKNXKltmqEVZwaBfrECuk |
|  7149465096876260652 | YThJJGJhZroiXcGOsxTMDKQFnJWOaLDwiYDQuhmQXvpvKHFPoAOFIFRFAIGvIivwoGJIvSvHdipBoPhRyegrvxGfGQeJqIpZScoiKDsBTLnbvjjytQvaPwqCiMRcekKJrkjdbrELhCnPpzcwUOWgMLfVHhElLpOfzPeEJGDbiuAHkuVJNwdUILSMaMFXBzJKFTDjixkhpkBTSaPYjrbypAEVfZWqXJrUVBJzptjjFUbTWRjDJlwqCGQTbzGziwP |
|  2666746673038440181 | IiHvgrTezldsYKTLJOObfVhTMosVUJpTOlvgyoxOjnilJOxaRiUbHylMZABduWEBOfausQRxqjFnsJvYuGWoRiHDkCuwvytHIpJKcoAEHridFPZbrwKdPoFKMslDxPAZgwRmKSTYEXCzZMqYXawXustnPJyoulDPInlebKtJWxHEmffOuBlxfharCDxxMObaZaRELuvFWUihAoBZGQDvglOBihSQgkpIlrKYZrwrrQIkZjnOBhqFUcdNEpJopwb |
|  2153189013947154756 | tUmRasqhpHGFyKQWoLpAorIdvLuRDnQsVMnVYSCccNdqjGUeLEQHpmIamunxXOTQUqmTMzPwMxdoNiUfJWomLNaTKDHgkzgfVhCPQhYRszoUjQIsmSlAHtjsvrTybXXwZQJhxdaRZcKiedaNDWwdhkBwZFCUnaWHFvosbnBCiXfvdJzNZIAKYqbNFNhTdPhcClGPySDEvYyGcZJvRPnMfChoQYsRbFLiJJUZRsiGPojKCWdwrdvVMQvVuBXQWAc |
| -6885074001682187751 | QOiVkxndzVbQbxtITxWvSWNWQYppdvlfoirzvBbWnLjPMwRQceJQdJqsRoucehDvaPRIUYBLilwayRpgoWcMFYnNMLSAmQcEWUFXPelWctrZNiMLWmxWJTmbftXjUETNsizPnCwAySzNTOfiSxhkjcAkcwGMFJXFPgeMEsCbmPlhduuDLewZwxdWwkDGCjOoBCGkuXsNAkWKWYDMxaMmaEVHYnnYbXfmdrXFdYWQhPeZZuLEydesWZRWdAmfjEq |
+----------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
5 rows in set (0.01 sec)
```

## Config
This is full type of config file. You can add tables, columns, indexes as many as you want.

```yaml
driver: mysql
database:
  host: 127.0.0.1
  user: root
  password: root
  name: testdb
  port: 3306
tables:
- name: table_a
  columns:
    - name: col_1
      type: bigint
      order: 20
      precision:
      unsigned: true
      notNUll: true
      default:
      primary: true
      autoIncrement: true
      values:
    - name: col_2
      type: varchar
      order: 50
      precision:
      unsigned: false
      notNUll: false
      default:
      primary: false
      autoIncrement: false
      values:
  indexes:
    - name: index_1_on_table_a
      uniq: true
      columns:
        - col_2
  charset: utf8mb4
  record: 100000
```

### Driver
Choose rdb driver for database you want to populate.

```
driver: mysql
```

Currently only supports below RDBMS

- MySQL

### Database
```yaml
database:
  host: 127.0.0.1
  user: root
  password: root
  name: testdb
  port: 3306
```

### Tables
If the table w/ designated name is not existed, this tool create table firstly. This is executed by `CREATE TABLE IF NOT EXIST` statment, so charset is required.

Record is a number that how many records you want to insert to this table.

```yaml
tables:
- name: table_a
  columns:
    - name: col_1
      type: bigint
      order: 20
      precision:
      unsigned: true
      notNUll: true
      default:
      primary: true
      autoIncrement: true
      values:
    - name: col_2
      type: varchar
      order: 50
      precision:
      unsigned: false
      notNUll: false
      default:
      primary: false
      autoIncrement: false
      values:
  indexes:
    - name: index_1_on_table_a
      primary: false
      uniq: true
      columns:
        - col_2
  charset: utf8mb4
  record: 100000
```

### Columns
Column represents what kind of columns should be held by the table. This only works when table is not existed.

If you want to enable option like unsigned, notNull and so on, you need to set true flag. If non-support option is set to true like varchar w/ unsigned, that option is just ignored.

```yaml
  columns:
    - name: col_1
      type: bigint
      order: 20
      precision:
      unsigned: true
      notNUll: true
      default:
      primary: true
      autoIncrement: true
      values:
```

Some of options are not required. This is a minimal description.

```yaml
  columns:
    - name: col_1
      type: bigint
```

When you don't give order/precision to data-type rquiring them like int(x), varchar(x), the default values are used. The default valeus are basically based on rdb default (int's default order is 11). Also we support unofficial default order/precision for rest of data-types like varchar, float and so on. So if your focus is not on the order/precision, you don't need to describe them.

Also we support concrete value constraint by values. From below declaration, col_1 would be populated w/ only "YES", "NO".

```yaml
  columns:
    - name: col_1
      type: varchar
      values:
        - "YES"
        - "NO"
```

### Indexes
Index represents what kind of indexes should be held by the table. This only works when table is not existed.

```yaml
  indexes:
    - name: index_1_on_table_a
      primary: false
      uniq: true
      columns:
        - col_2
```

Some of options are not required. This is a minimal description.

```yaml
  indexes:
    - columns:
        - col_2
```

When you want to use covering index, you can describe like below.

```yaml
  indexes:
    - columns:
        - col_1
        - col_2
```

### Examples
There're sample config files, you can try it.

when you haven't setup the database or tables yet
```shell
$ populator -c ./examples/from_create_table.yaml
```

when you've already setup the database and tables
```shell
$ populator -c ./examples/only_populate_seed.yaml
```

when you want to re-create table schema w/ given declarations
```shell
$ populator -rc ./examples/from_create_table.yaml
```
