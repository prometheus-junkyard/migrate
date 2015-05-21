# Prometheus Migration Tool

A tool for migrating Prometheus configuration files across versions.

Install:
```
$ go get github.com/tools/godep
$ go get github.com/prometheus/migrate
$ cd $GOPATH/src/github.com/prometheus/migrate
$ godep go install
```

Usage:
```
migrate -out=new_conf.yml old_conf.conf
```

Migration will not preserve your comments. It is generally recommended for larger files that are tedious to migrate by hand.

Reading the configuration documentation will provide you with further insight about new possibilities.
