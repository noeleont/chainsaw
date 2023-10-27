# Release notes

Release notes for `TODO`.

<!--
## :star: Examples :star:

## :boat: Tutorials :boat:

## :wrench: Fixes :wrench:
-->

## :dizzy: New features :dizzy:

- Added `Command` and `Script` operations
- Added support to continue on error in `Delete`, `Apply`, `Assert` and `Error` operations
- Added support to specifify if a test should run concurrently or not at the `Test` level
- Added support for overriding the namespace used at the `Test` level
- Added `Event` and `Pod` logs collectors support
- Added `chainsaw kuttl migrate` command to migrate KUTTL tests to chainsaw
- Docker image is now released under `ghcr.io/kyverno/chainsaw`

## :sparkles: UI changes :sparkles:

- Improved logging of running tests

## :books: Docs :books:

- Improved CLI docs

## :guitar: Misc :guitar:

- Switched to subtests for better summary report from the `testing` package