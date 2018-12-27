#!/bin/bash
# Continue when error
set +e
# Break when error
# set -e

# Specify your linter variable here
GometalinterVariable=(
   "vetshadow"
   "vet"
   "gochecknoglobals"
   "gochecknoinits"
   "gofmt"
   "goimports"
   "gosimple"
#    "lll"
   "misspell"
   "staticcheck"
   "test"
   "testify"
   "unused"
    ##########
    "deadcode"
    # "dupl"
    "errcheck"
    "goconst"
    "gocyclo"
    "golint"
    "gosec"
    "gotype"
    "gotypex"
    "ineffassign"
    "interfacer"
    "maligned"
    "megacheck"
    "nakedret"
    "safesql"
    "structcheck"
    "unconvert"
    "unparam"
    "varcheck"
)

# Specify your directory relative path here
Directory=(
    "helper"
    "domain"
    "validator"
    "validator/decimal"
    "validator/email"
    "validator/errors"
    "validator/numeric"
    "validator/text"
)

arrayGometalinterVariable=${#GometalinterVariable[@]}
arrayDirectory=${#Directory[@]}

for ((k=0; k<arrayDirectory; k++));
do
    for ((i=0; i<arrayGometalinterVariable; i++));
    do
        echo "Currently linter running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
        gometalinter -j 1 --disable-all --enable="${GometalinterVariable[$i]}" "${Directory[$k]}"/ 2>&1
        sleep 0.5
        wait
    done
done
