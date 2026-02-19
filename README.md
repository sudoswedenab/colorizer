# Colorizer

This program adds colors to the terminal output of commands in a best
effort manner. Which is useful for more easily seeing at a glance which
test in a pipeline has failed and why.

Correct usage is as follows:

```sh

set -eo pipefail # Needed since piping output would otherwise swallow error exit code

go test ./... | colorizer

```
