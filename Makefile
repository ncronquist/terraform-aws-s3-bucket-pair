check: fmt-check
	cd test && go test -v -timeout 30m

fmt:
	terraform fmt -recursive
	@goimports -w ./test

fmt-check:
	terraform fmt -check -recursive
	test -z $$(goimports -l ./test)

# Alias for check
test: check
