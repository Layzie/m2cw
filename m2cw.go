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
	arg := c.Args().First()
	cmd := exec.Command("markdown2confluence", arg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(strings.TrimSuffix(arg, filepath.Ext(arg))+".wiki", out.Bytes(), 0644)
}

func main() {
	app := cli.NewApp()
	app.Name = "m2cw"
	app.Usage = "Convert markdown file to confluence style when save md file."
	app.Author = "Layzie <HIRAKI Satoru>"
	app.Email = "saruko313@gmail.com"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		arg := c.Args().First()

		fmt.Println("Start watching md file. <C-c> makes stop the command.")

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.Watch("./")
		if err != nil {
			log.Fatal(err)
		}

		defer watcher.Close()

		for {
			select {
			case ev := <-watcher.Event:
				if ev.Name == c.Args().First() {
					md2conf(c)
					log.Println("convert md to wiki ", ev.Name+" -> "+strings.TrimSuffix(arg, filepath.Ext(arg))+".wiki")
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}

	app.Run(os.Args)
}
