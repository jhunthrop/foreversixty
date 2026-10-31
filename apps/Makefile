# The apps' one release command (docs/runbooks/mobile-apps.md "Cutting a release").
#   make -C apps release                                  dry run: the plan, nothing changes
#   make -C apps release ARGS="--bump minor --execute"    bump, build, sign, upload
# ARGS takes tools/release-apps.sh's flags: --bump build|patch|minor|major,
# --platform both|android|ios, --execute.
ARGS ?=

.PHONY: release release-test

release:
	../tools/release-apps.sh $(ARGS)

release-test:
	bash ../tools/tests/test_release_apps.sh
