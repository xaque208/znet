#
# Makefile Fragment for Compiling
#

GO         ?= go
BUILD_DIR  ?= ./bin/
PROJECT_MODULE ?= $(shell $(GO) list -m)
# $b replaced by the binary name in the compile loop, -s/w remove debug symbols
LDFLAGS    ?= "-s -w -X main.Version=$(PROJECT_VER) -X main.appName=$$b"
SRCDIR     ?= .
COMPILE_OS ?= freebsd linux

# Determine commands by looking into cmd/*
COMMANDS   ?= $(wildcard ${SRCDIR}/cmd/*)

# Determine binary names by stripping out the dir names
BINS       := $(foreach cmd,${COMMANDS},$(notdir ${cmd}))


compile-clean:
	@echo "=== $(PROJECT_NAME) === [ compile-clean    ]: removing binaries..."
	@rm -rfv $(BUILD_DIR)/*

compile: deps compile-only proto

proto:
	@echo "=== $(PROJECT_NAME) === [ proto compile    ]: compiling protobufs:"
	@protoc -I rpc/ rpc/rpc.proto --go_out=plugins=grpc:rpc

compile-all: deps-only
	@echo "=== $(PROJECT_NAME) === [ compile          ]: building commands:"
	@mkdir -p $(BUILD_DIR)/$(GOOS)
	@for b in $(BINS); do \
		for os in $(COMPILE_OS); do \
			echo "=== $(PROJECT_NAME) === [ compile          ]:     $(BUILD_DIR)$$os/$$b"; \
			BUILD_FILES=`find $(SRCDIR)/cmd/$$b -type f -name "*.go"` ; \
			GOOS=$$os $(GO) build -ldflags=$(LDFLAGS) -o $(BUILD_DIR)/$$os/$$b $$BUILD_FILES ; \
		done \
	done

compile-only: deps-only
	@echo "=== $(PROJECT_NAME) === [ compile          ]: building commands:"
	@mkdir -p $(BUILD_DIR)/$(GOOS)
	GOOS=$(GOOS) $(GO) build -ldflags=$(LDFLAGS) -o $(BUILD_DIR)/$(GOOS)/$(PROJECT_NAME) . ; 
	@# @for b in $(BINS); do \
	# 	echo "=== $(PROJECT_NAME) === [ compile          ]:     $(BUILD_DIR)$(GOOS)/$$b"; \
	# 	BUILD_FILES=`find $(SRCDIR)/cmd/$$b -type f -name "*.go"` ; \
		GOOS=$(GOOS) $(GO) build -ldflags=$(LDFLAGS) -o $(BUILD_DIR)/$(GOOS)/$$b $$BUILD_FILES ; \
	# done

# Override GOOS for these specific targets
compile-darwin: GOOS=darwin
compile-darwin: deps-only compile-only

compile-linux: GOOS=linux
compile-linux: deps-only compile-only

compile-windows: GOOS=windows
compile-windows: deps-only compile-only


.PHONY: clean-compile compile compile-darwin compile-linux compile-only compile-windows