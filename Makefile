.PHONY: all clean
all: test README.md

.PHONY: default
default: test

.PHONY: test
test: check

# Installs pre-commit hooks
.PHONY: install-hooks
install-hooks:
	pre-commit install --install-hooks

# Runs pre-commit checks on files
.PHONY: check
check:
	pre-commit run --all-files


README.md: ./add-docs-to-readme.sh *.go go.mod
	./add-docs-to-readme.sh
