.PHONY: all
all: build test

ALL_PACKAGES=$(go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(go list ./... | grep -v "vendor" | grep -v "featuretest")

APP_EXECUTABLE="out/ifsc"

setup_mac:
	brew install dep

setup_linux:
	go get -u github.com/golang/dep/cmd/dep

init:
	dep init

add:
	dep ensure -add $(repo)

update:
	dep ensure

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test:
	go test ./... -p=1 -race

testp:
	go test -covermode=count -v

coverage:
	./test.sh
	go tool cover -html=coverage.txt -o coverage.html

build: update fmt vet lint compile test

build_ci: clean setup_linux update fmt vet lint copy-config compile coverage

copy-config:
	cp application.toml.sample application.toml

install:
	go install

clean:
	rm -rf application.toml
	rm -rf coverage.html
	rm -rf coverage.txt
	rm -rf out/ifsc
	rm -rf out/coverage.html
	rm -rf coverage-all.out

perftest:
	vegeta attack -targets=perftest/target.txt -rate=100 -duration=10s | vegeta report -reporter=plot -output=perftest/report.html
	open perftest/report.html