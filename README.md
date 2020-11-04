# Happy Team

Measure the mood of your team, judging by emoji usage on Slack.

This is a fork of [SlackMood](https://github.com/YoSmudge/SlackMood) (MIT licensed, no development since 2016).

THIS IS A WORK IN PROGRESS since many of the Slack API calls have been deprecated since 2016.

Pull requests are welcome.

[![Build Status](https://travis-ci.com/xyproto/happyteam.svg?branch=master)](https://travis-ci.com/xyproto/happyteam)

![](https://s3.amazonaws.com/f.cl.ly/items/0E3W453j2I44451b441x/Screen%20Shot%202016-05-31%20at%2015.01.18.png?v=7d9a7302)

## Building

    cd cmd/server && go build -mod=vendor

## Running

First create `config/config.yml`, containing your Slack bot token and a path to the BoltDB file:

```
slack_token: "abcd"
db_path: "db/db.bolt"
rank_file: "config/rank.csv"
```

Use your own Slack token, for a Slack bot. Creating a Slack token may be tricky, and may involve creating both an app and a bot user and also giving it permissions. Slack have made it as complicated as possible.

Then run:

    ./server -b :8000

## General info

* Version: 1.0.0
* License: MIT
