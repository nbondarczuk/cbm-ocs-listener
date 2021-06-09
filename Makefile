TARGET = cbm-ocs-listener
VERSION = 0.0.0
BUILD = $(shell date +"%F_%T_%Z")
LEVEL = "unknown"
IMAGE_BASE = registry1.corpo.t-mobile.pl/cbm-ocs-listener/centos-oracle-tibco-base
IMAGE = registry1.corpo.t-mobile.pl/cbm-ocs-listener/cbm-ocs-listener:$(VERSION)
DOCKERFILE = Dockerfile.centos
DOCKERFILE_BASE = Dockerfile-base.centos
GOOS = linux
GOARCH = amd64

DEPS = "gopkg.in/goracle.v2" \
"gopkg.in/gorp.v2" 

LDFLAGS = "-X main.version=$(VERSION) -X main.build=$(BUILD) -X main.level=$(LEVEL)"
STATIC_BUILD_PREFIX = "CGO_ENABLED=1 GOOS=linux GOARCH=amd64"
STATIC_LDFLAGS = "-w -extldflags '-static -I$(GOPATH)/src/gopkg.in/goracle.v2/odpi/include -I$(GOPATH)/src/gopkg.in/goracle.v2/odpi/src -I$(GOPATH)/src/gopkg.in/goracle.v2/odpi/embed -ldl' "

all: build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=$(LDFLAGS) -o $(TARGET) listener.go 

build-static:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -tags netgo -cflags=$(CFLAGS) -ldflags=$(STATIC_LDFLAGS) -o $(TARGET) listener.go -L /opt/tibco/ems/8.5/lib -ltibems64

run:
	go run listener.go

docker: build
	echo $(VERSION) > VERSION
	cp config/config.json.T17 config.json
	docker build -f $(DOCKERFILE) -t $(IMAGE) .

docker-base:
	docker build -f $(DOCKERFILE_BASE) -t $(IMAGE_BASE) .

docker-run:
	docker run -it --rm --network host --name $(TARGET) $(IMAGE)

docker-run-detached:
	docker run -d --network host --name $(TARGET) $(IMAGE)

docker-push:
	docker push $(IMAGE)

docker-base-push:
	docker push $(IMAGE_BASE)

deps:
	go get $(DEPS)

sql: sql-xe

sql-xe:
	(cd sql; bash install-xe.sh)

clean:
	go clean
	rm -f *~ */*~ $(TARGET)
