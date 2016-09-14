Rabot
=====

[![CircleCI](https://circleci.com/gh/ngs/rabot.svg?style=svg)](https://circleci.com/gh/ngs/rabot)
[![Docker Automated build](https://img.shields.io/docker/automated/atsnngs/rabot.svg?maxAge=2592000)](https://hub.docker.com/r/atsnngs/rabot/)

The Chatbot deals with [docker-radiko-recorder-s3]

Environment Variables
---------------------

```sh
RADIKO_LOGIN
RADIKO_PASSWORD
S3_BUCKET
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
SLACK_WEBHOOK_URL
SLACK_TOKEN
IMAGE_NAME (default: atsnngs/radiko-recorder-s3)
DEBUG
```

Run it
------

```sh
make
./rabot
```

or

```sh
docker pull atsnngs/rabot
docker run -v /var/run/docker.sock:/var/run/docker.sock --env-file=.envrc atsnngs/rabot
```

How it works
------------

```
@rabot start recording ALPHA-STATION for 1 min
@rabot list containers
@rabot remove container 452ca45d449a
```

[docker-radiko-recorder-s3]: https://github.com/ngs/docker-radiko-recorder-s3
