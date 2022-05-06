#!/bin/bash

readonly ROOT_DIR="$(cd "$(dirname "$0")"/.. && pwd)"
cd "${ROOT_DIR}/src/main"

readonly TARGET_DIR="${ROOT_DIR}/target"
mkdir -p "${TARGET_DIR}"

go build -o "${TARGET_DIR}/infographics"

