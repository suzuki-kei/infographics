#!/bin/bash

readonly ROOT_DIR="$(cd "$(dirname "$0")"/.. && pwd)"
cd "${ROOT_DIR}"

for module_dir in $(find src/ -type f -name '*_test.go' | xargs dirname | sort -u)
do
    (cd "${module_dir}" && go test)
done

