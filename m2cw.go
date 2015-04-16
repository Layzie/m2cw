package main

import (
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func md2conf(c *cli.Context) {
	cmd := exec.Command("markdown2confluence", c.Args().First())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(strings.TrimSuffix(c.Args().First(), filepath.Ext(c.Args().First()))+".wiki", out.Bytes(), 0644)
}

func main() {
	app := cli.NewApp()
	app.Name = "m2cw"
	app.Usage = "Convert markdown file to confluence style when save md file."
	app.Author = "Layzie <HIRAKI Satoru>"
	app.Email = "saruko313@gmail.com"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		fmt.Println("Start watching md file. <C-c> makes stop the command.")
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		done := make(chan bool)

		// Process events
		go func() {
			for {
				select {
				case ev := <-watcher.Event:
					// log.Println("event:", ev)
					if ev.Name == c.Args().First() {
						md2conf(c)
						log.Println("convert md to wiki ", ev.Name+" -> "+strings.TrimSuffix(c.Args().First(), filepath.Ext(c.Args().First()))+".wiki")
					}
				case err := <-watcher.Error:
					log.Println("error:", err)
				}
			}
		}()

		err = watcher.Watch("./")
		if err != nil {
			log.Fatal(err)
		}

		<-done

		watcher.Close()
	}

	app.Run(os.Args)
}
