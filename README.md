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

when you haven't setup the database or tables yet
```shell
$ populator populate --config ./examples/from_create_table.yaml
```

when you've already setup the database and tables
```shell
$ populator populate --config ./examples/only_populate_seed.yaml
```
