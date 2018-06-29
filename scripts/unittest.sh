#!/usr/bin/env bash

if [ -z "${REPORTER}" ]
then
  export REPORTER=cc-test-reporter
fi

function unittest() {
    # unit tests all the packages
    for dir in `ls . | grep -v vendor`;
    do
        if [[ -d ${dir} && ! "${dir}" =~ ^(scripts)$ ]]
        then
            go test ./${dir} -race -v -coverprofile cover.${dir}.out
        fi

    done

    # unit test the main package
    go test . -race -v -coverprofile cover.out
}


function generate_coverage() {
    echo "Using $REPORTER to generate JSON files"
}


# Execute the scripts
case $1 in

run)
  unittest
;;

coverage)
  generate_report
;;

*)
  exit 1

esac
