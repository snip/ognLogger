# ognLogger / OGN APRS messages logger written in Golang
Quick and maybe dirty [OGN](https://glidernet.org) APRS messages logger written in Golang

## Concept
ognLogger is running on a host as daemon mode.
It logs all OGN traffic present on APRS server and generate a file per day with this data.

## Usage
Download binary from https://github.com/snip/ognLogger/releases or build it yourself.
Then run `./ognLogger`

## Check activity
ognLogger will output to the terminal it is started in.

## Building
Golang installation on Ubuntu/Raspbian:
https://github.com/golang/go/wiki/Ubuntu

Normal build:
```
go get
go build
```
