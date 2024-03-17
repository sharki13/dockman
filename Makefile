.PHONY: cli
cli:
	go build -tags cli -o build/bin/

.PHONY: clean
clean:
	rm -f build/bin/*