# tail2redis
Tailing a file to Redis server (PubSub)

## Example

```bash
$ ./tail2redis -h
NAME:
   tail2redis - Tailing a file to redis server (PubSub)

USAGE:
   tail2redis [global options] command [command options] [arguments...]

AUTHOR:
   Namhyeon Go <abuse@catswords.net>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value      Choose a file (default: "/var/log/messages")
   --host value, -H value      Redis host (default: "localhost")
   --port value, -P value      Redis host (default: 6379)
   --password value, -p value  Redis password
   --channel value, -c value   Redis PubSub channel (default: "tail2redis")
   --help, -h                  show help

$ ./tail2redis -H localhost -P 6379 -p [PASSWORD] -c [CHANNEL] -f /var/log/messages
```

## Contact me
* abuse@catswords.net
