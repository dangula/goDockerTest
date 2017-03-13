#!/bin/bash
echo $TestName
glide update
go test -v $TestName
