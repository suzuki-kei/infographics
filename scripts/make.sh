#!/usr/bin/env bash

set -eu -o posix -o pipefail

declare ROOT_DIR
ROOT_DIR="$(cd -- "$(dirname -- "$0")/.." && pwd)"
declare -gr ROOT_DIR

declare SOURCE_DIR
SOURCE_DIR="${ROOT_DIR}/src"
declare -gr SOURCE_DIR

declare TARGET_DIR
TARGET_DIR="${ROOT_DIR}/target"
declare -gr TARGET_DIR

function main
{
    declare -r command="${1:-}"

    case "${command}" in
        build | clean | run | stats | test)
            shift
            make_${command} "$@"
            ;;
        *)
            echo "invlid options -- ${command}" >&2
            exit 1
            ;;
    esac
}

function make_build
{
    mkdir -p "${TARGET_DIR}"
    go build -C "${SOURCE_DIR}" -o "${TARGET_DIR}/infographics"
}

function make_clean
{
    rm -rf "${TARGET_DIR}"
}

function make_run
{
    go run -C "${SOURCE_DIR}" . "$@"
}

function make_stats
{
    echo '# Annotaion Comments'
    (
        cd "${SOURCE_DIR}"
        find . -type f -name '*.go' | xargs grep -P 'TODO|NOTE|FIXME|XXX'
    ) | sed 's/^/    /'

    echo ''

    echo '# Line of Codes'
    (
        cd "${SOURCE_DIR}"
        find . -type f -name '*.go' | xargs wc -l
    ) | sed 's/^/    /'
}

function make_test
{
    for package_dir in $(find "${SOURCE_DIR}" -type f -name '*_test.go' | xargs dirname | sort -u)
    do
        (cd "${package_dir}" && go test)
    done
}

if [[ "$0" == "${BASH_SOURCE[0]}" ]]; then
    main "$@"
fi

