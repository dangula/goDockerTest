#!/bin/bash
echo $TestName
glide updates
go test -v $TestName
