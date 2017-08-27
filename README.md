api
===

[![CircleCI](https://circleci.com/gh/torinos-io/api.svg?style=svg&circle-token=37af875aec12df40b622bff3a74de69e94416d44)](https://circleci.com/gh/torinos-io/api)

The core API server


Setup
-----

Install Docker, docker-compose and https://github.com/creasty/rid.

```sh-session
$ rid compose pull
```

```
$ rid script/migrate
$ rid compose exec torinos_db psql -t maindb -f data/seeds/users.sql
$ rid compose exec torinos_db psql -t maindb -f data/seeds/projects.sql
```


Run
---

```sh-session
# install deps
$ rid glide i

# run server
$ rid make run

# run test
$ rid make test
```
