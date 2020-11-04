# Emoji mood

Measure the mood of your team, juding by emoji usage on Slack!

![](https://s3.amazonaws.com/f.cl.ly/items/0E3W453j2I44451b441x/Screen%20Shot%202016-05-31%20at%2015.01.18.png?v=7d9a7302)

This is a fork of [SlackMood](https://github.com/YoSmudge/SlackMood) (MIT licensed).

## Building

    go build -mod=vendor

## Running

First create `config.yml`, containing your Slack bot token and a path to the BoltDB file:

```
slack_token: "abcd"
db_path: "./db.bolt"
```

Then run:

    ./emojimood --config config.yml --bind :3044

## General info

* Version: 0.0.1
* License: MIT
