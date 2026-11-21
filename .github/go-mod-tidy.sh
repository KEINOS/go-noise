#!/bin/sh
# Upgrade the root module and sync the example with the updated root.
# The tidy service runs this script with the minimum supported Go toolchain.
set -eu

cd "$(dirname "$0")/.."
export GOTOOLCHAIN=local GOWORK=off

# Use a fixed directory so two updates cannot overwrite each other's backups.
backup_dir=.go-mod-update
mkdir "$backup_dir"
restore=false

finish() {
    exit_code=$?
    trap - 0
    trap '' HUP INT TERM
    if [ "$restore" = true ]; then
        rollback_failed=false
        cp -p "$backup_dir/root.mod" go.mod || rollback_failed=true
        cp -p "$backup_dir/root.sum" go.sum || rollback_failed=true
        cp -p "$backup_dir/example.mod" _example/2d/go.mod || rollback_failed=true
        cp -p "$backup_dir/example.sum" _example/2d/go.sum || rollback_failed=true

        if [ "$rollback_failed" = true ]; then
            printf 'Could not restore module files. Backups are in %s.\n' "$backup_dir" >&2
            exit 1
        fi
    fi

    rm -f "$backup_dir/root.mod" "$backup_dir/root.sum" \
        "$backup_dir/example.mod" "$backup_dir/example.sum"
    rmdir "$backup_dir"
    exit "$exit_code"
}

trap finish 0
trap 'exit 1' HUP INT TERM

cp -p go.mod "$backup_dir/root.mod"
cp -p go.sum "$backup_dir/root.sum"
cp -p _example/2d/go.mod "$backup_dir/example.mod"
cp -p _example/2d/go.sum "$backup_dir/example.sum"
restore=true

# Keep each module's Go and toolchain directives. Tidy rejects a dependency
# if it needs a higher Go version than the module declares.
tidy() {
    minimum_go=$(awk '$1 == "go" { print $2 }' "$1")
    toolchain=$(awk '$1 == "toolchain" { print $2 }' "$1")

    go mod tidy "-go=$minimum_go"
    go mod edit "-toolchain=${toolchain:-none}"
}

printf 'Updating dependencies with %s\n' "$(go version)"
go get -t -u ./... ./_example/3d
tidy "$backup_dir/root.mod"
go test -mod=readonly ./... ./_example/3d

(
    cd _example/2d
    # Sync the local root replacement. Gonum Plot stays on its Go 1.22 release.
    tidy "../../$backup_dir/example.mod"
    go test -mod=readonly .
)

restore=false
printf 'Module updates passed validation.\n'
