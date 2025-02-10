.DEFAULT_GOAL := generate

generate:
	go generate ./...
	mv gen/grammar/* gen
	rm -fr gen/grammar
.PHONY:generate
