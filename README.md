# Malgo
[![Build Status](https://travis-ci.com/th3-z/malgo.svg?branch=master)](https://travis-ci.com/th3-z/malgo) [![Go Report Card](https://goreportcard.com/badge/github.com/th3-z/malgo)](https://goreportcard.com/report/github.com/th3-z/malgo) [![GitHub license](https://img.shields.io/github/license/th3-z/malgo)](https://github.com/th3-z/malgo/blob/master/LICENSE) [![Codecov](https://img.shields.io/codecov/c/github/th3-z/malgo.svg?style=flat)](https://codecov.io/gh/th3-z/malgo)

A MyAnimeList SQL migration utility and Go library.

## Cli exporter usage

* Export your list from MyAnimeList
* Run `malgo -o anime.sqlite yourAnimeList.xml`
* Receieve SQLite database `./anime.sqlite`

## Databases

Malgo has so far been tested with the following databases. Let me know if you have success with any of the other database drivers.

* Sqlite3

## Golang examples

The golang module provides structs and functions for reading and editing the migrated database.

#### Migrate directly from an exported anime list and retrieve series.
```
malgo.storage.CreateSchema(someDb)
malgo.MigrateFile(someDb, "sample.xml")

user := malgo.models.SearchUser(someDb, "th3-z")
for _, review := range user.Reviews {
    print(review.Series.Name)
    print("\n")
}
```

#### Datatables example
A more complex example can be seen on [my public website](https://github.com/th3-z/beta-th3-z-xyz)'s repo. It produces the following table.

![Image of example usage](https://files.th3-z.xyz/standing/malgo.png)


## Building

Run `make build` to create a release for the CLI tool in `./bin`.

### Requirements

* Golang 1.13.x
* make
* gcc

