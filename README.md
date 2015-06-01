# Prometheus Migration Tool

A tool for upgrading Prometheus setups to a newer version.

Currently, it migrates ASCII protcol buffer configurations from pre-v0.14 setups to the
respective YAML equivalent.

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

Migration will not preserve comments. It is generally recommended for 
larger files that are tedious to translate by hand.

Reading the configuration documentation will provide you with further insight 
about new possibilities.
