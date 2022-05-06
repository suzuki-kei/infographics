#!/bin/bash

readonly ROOT_DIR="$(cd "$(dirname "$0")"/.. && pwd)"
cd "${ROOT_DIR}/src/main"
go run . "$@"

