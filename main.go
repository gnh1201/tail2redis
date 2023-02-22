//
// tail2redis
// Go Namhyeon <abuse@catswords.net>
//

package main

import (
    "fmt"
    "log"
    "os"
    "context"

    "github.com/urfave/cli/v2"
    "github.com/nxadm/tail"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
    app := &cli.App{
        Name:  "tail2redis",
        Usage: "Tailing a file to redis server (PubSub)",
        Compiled: time.Now(),
        Authors: []*cli.Author{
            &cli.Author{
                Name:  "Namhyeon Go",
                Email: "abuse@catswords.net",
            },
        },
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "file",
                Aliases: []string{"f"},
                Value:   "/var/log/messages",
                Usage:   "Choose a file",
            },
            &cli.StringFlag{
                Name:    "host",
                Aliases: []string{"H"},
                Value: "localhost",
                Usage:   "Redis host",
            },
            &cli.IntFlag{
                Name:    "port",
                Aliases: []string{"P"},
                Value: 6379,
                Usage:   "Redis host",
            },
            &cli.StringFlag{
                Name:    "password",
                Value:   "",
                Aliases: []string{"p"},
                Usage:   "Redis password",
            },
            &cli.StringFlag{
                Name:    "channel",
                Aliases: []string{"c"},
                Value: "tail2redis",
                Usage:   "Redis PubSub channel",
            },
        },
        Action: func(c *cli.Context) error {
            file := c.String("file")
            host := c.String("host")
            port := c.Int("port")
            password := c.String("password")
            channel := c.String("channel")

            // Connect to Redis server
            rdb := redis.NewClient(&redis.Options{
                Addr:     fmt.Sprintf("%s:%d", host, port),
                Password: password, // no password set
                DB:       0,  // use default DB
            })


            // Tailing a file
            t, err := tail.TailFile(file, tail.Config{Follow: true, ReOpen: true})
            if err != nil {
                fmt.Println(err)
            }

            for line := range t.Lines {
                err2 := rdb.Publish(ctx, channel, line.Text).Err()
                if err2 != nil {
                    fmt.Println(err2)
                }
            }

            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }

}
