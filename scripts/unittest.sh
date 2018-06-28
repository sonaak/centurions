#!/usr/bin/env bash

for dir in `ls . | grep -v vendor`;
do
  if [[ -d ${dir} && ! "${dir}" =~ ^(scripts)$ ]]
  then
    go test ./${dir} -race -v -coverprofile cover.${dir}.out
  fi

done
go test . -race -v -coverprofile cover.out
