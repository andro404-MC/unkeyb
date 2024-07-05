# Unkeyb ![pass](https://github.com/andro404-MC/unkeyb/actions/workflows/test.yml/badge.svg) ![GitHub License](https://img.shields.io/github/license/andro404-MC/unkeyb)

A simple TUI keyboard typing speed test built using Go and the bubbletea framework

[preview.webm](https://github.com/andro404-MC/unkeyb/assets/94703538/d897f056-8a95-46af-a7ab-34f2d410ab38)

> [!NOTE]
> currently supporting english.
> Currently supporting US and GB layouts.

## Requirement :

Nothing unless :

`go` : if you are going to build from source.

## Build :

> You need a to have `GOPATH` added to `PATH`

```
$ git clone https://github.com/andro404-MC/gokeyb
$ cd unkeyb

// Run
$ go run .

// Install
$ go install .
```

## Usage :

To run :

```
$ unkeyb
```

set layout :

```
$ unkeyb -k gb
```

set language :

```
$ unkeyb -l en
```

show help :

```
$ unkeyb -h
Usage of unkeyb:
  -k string
    	layout (us,gb) (default "us")
  -l string
    	Language (en) (default "en")
```
