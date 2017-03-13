#!/bin/bash
echo $TestName
go install
go test -v $TestName
