#!/bin/bash

readonly ROOT_DIR="$(cd "$(dirname "$0")"/.. && pwd)"
cd "${ROOT_DIR}"

function generate_go_mod_files
{
    for module_dir in $(find src/ -mindepth 1 -maxdepth 1 -type d | sort)
    do
        generate_go_mod_file $module_dir
    done
}

function generate_go_mod_file
{
    local -r target_module_dir=$1
    (
        cd $target_module_dir
        rm -f go.mod
        local -r target_module_name=$(basename $target_module_dir)
        go mod init $target_module_name
    )

    for module_dir in $(find src/ -mindepth 1 -maxdepth 1 -type d | sort)
    do
        (
            cd $target_module_dir
            local -r module_name=$(basename $module_dir)
            go mod edit -require=$module_dir@v0.0.0
            go mod edit -replace=$module_dir=../$module_name
        )
    done
}

generate_go_mod_files

