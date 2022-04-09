COVERAGE_DIR ?= .coverage

test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	go test -v -race -covermode atomic -coverprofile $(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m ./...
