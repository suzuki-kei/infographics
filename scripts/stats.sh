#!/bin/bash

readonly ROOT_DIR="$(cd "$(dirname "$0")"/.. && pwd)"
cd "${ROOT_DIR}"

echo '# Annotaion Comments'
(
    find src/ -type f -name '*.go' | xargs grep -P 'TODO|NOTE|FIXME|XXX'
) | sed 's/^/    /'

echo ''

echo '# Line of Codes'
(
    find src/ -type f -name '*.go' | xargs wc -l
) | sed 's/^/    /'

