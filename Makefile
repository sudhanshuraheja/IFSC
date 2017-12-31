.PHONY: all
all: build test

ALL_PACKAGES=$(go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(go list ./... | grep -v "vendor" | grep -v "featuretest")

APP_EXECUTABLE="out/ifsc"

setup:
	brew install dep

init:
	dep init

add:
	dep ensure -add $(repo)

update:
	dep ensure

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

restart:
	mkdir -p out/
	go install $(ALL_PACKAGES)
	ifsc start

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

build: update fmt vet lint compile

run: fmt vet lint install

install:
	go install $(ALL_PACKAGES)

test:
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -p=1 -race

test_ci:
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -p=1 -race

test-cover-html:
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o out/coverage.html

test-open-html:
	open out/coverage.html

test-coverage: test-cover-html test-open-html

copy-config:
	cp application.toml.sample application.toml

copy-config-ci:
	cp application.toml.sample application.toml.ci

copy-configs: copy-config copy-config-ci

build-update-banks:
	go install
	ifsc updateBanks