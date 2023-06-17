# include .env
# ifneq (,$(wildcard ./.env))
#     include .env
#     export
# endif

# SHELL=bash
.PHONY: exportenv unsetenv devEnv

# define setup_env
# 	$(eval ENV_FILE := $(1).env)
# 	@echo " - setup env $(ENV_FILE)"
# 	$(eval include $(1).env)
# 	$(eval export)
# endef

# devEnv:
# 	$(call setup_env, dev)

exportenv:
    # export $(grep -v '^\#' .env | xargs -d '\n')
    # export $(grep -v '^\#' .env | xargs -d '\n')
	export $(grep -v '^\#' .env | xargs) \
	$(info Environment variables exported)

unsetenv:
	unset $(grep -v '^\#' .env | sed -E '/AF|PO/p' | xargs)