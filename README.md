# m2cw

## What's this?

This is golang port of [Layzie/md2conf-watcher](https://github.com/Layzie/md2conf-watcher).

Main purpose is my study of golang.

## Installation

```sh
$ gem install markdown2confluence # First time only.
$ go get github.com/Layzie/m2cw
```

## Usage

```sh
$ cd go/to/markdown
$ m2cw example.md
Start watching md file. <C-c> makes stop the command.
# if you edit example.md then convert to example.wiki on same directory
2015/04/16 19:06:12 convert md to wiki  example.md -> example.wiki
2015/04/16 19:07:23 convert md to wiki  example.md -> example.wiki
```

## Licence

MIT
