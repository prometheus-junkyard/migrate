# Copyright 2015 The Prometheus Authors
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERSION  := 0.1.0
TARGET   := migrate

include Makefile.COMMON

GODEP := PATH=$(PATH):$(GOROOT)/bin GOROOT=$(GOROOT) GOPATH=$(GOPATH) $(GOPATH)/bin/godep

dependencies-stamp: $(GOCC) $(SRC) | $(SELFLINK)
	if [ -d "Godeps" ]; then $(GO) get github.com/tools/godep; else $(GO) get -d; fi
	touch $@

$(BINARY): $(GOCC) $(SRC) dependencies-stamp Makefile Makefile.COMMON
	if [ -d "Godeps" ]; then $(GODEP) go build $(GOFLAGS) -o $@; else $(GO) build $(GOFLAGS) -o $@; fi
