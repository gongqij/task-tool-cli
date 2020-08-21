GO=go

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

VERSION := v1.0.0
BUILD := `git rev-parse --short HEAD`
TARGETS := task-tool-cli
project=task-tool-cli
LDFLAGS += -X "$(project)/version.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(project)/version.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(project)/version.Version=$(VERSION)"
LDFLAGS += -X "$(project)/version.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"

all: build

build: $(TARGETS)

$(TARGETS): $(SRC)
	$(GO) build -ldflags '$(LDFLAGS)' $(project)

.PHONY: all clean build  bash_completion

clean:
	rm -f $(TARGETS)

###zsh_completion:
###	$(call completion,"zsh",$(zshCompletionPath)/_task-tool-cli)

bash_completion:
	$(call completion,"bash","/etc/bash_completion.d/task-tool-cli")

############### functions ###############Day
define completion
	@echo "$(1) completion $(2)..."
    @./task-tool-cli completion $(1) > $(2)
    @echo "done"
endef