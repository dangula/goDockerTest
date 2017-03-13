#!/bin/bash
echo "Running Test : "$TestName
#update depenencies
glide update
#run test
go test -v $TestName
