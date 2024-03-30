
SHELL = /bin/bash

.DEFAULT_GOAL = help

.PHONY: help
help: # 全てのルールを表示する.
	$(foreach file, $(MAKEFILE_LIST), $(call list-rules, $(file)))

# 引数で指定された Makefile に定義されるルールを表示する.
define list-rules
	@grep -P '^[0-9a-zA-Z_-]+:' $(1) \
		| xargs -r -L1 printf '    %-25s %s %s %s %s %s %s %s %s %s %s %s\n'

endef

