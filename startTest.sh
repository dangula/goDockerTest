#!/bin/bash
echo "Running Test : "$TestName
#Install and update depenencies
glide install
glide update

# Run Test
case $TestName in

    #Run Block Tests
    [Bb][Ll][Oo][Cc][Kk])
          go test -v integrationTests/blockTests/*block*
          ;;
    [Ff][Ii][Ll][Ee])
          go test -v integrationTests/fileTests/*file*
          ;;
    [Oo][Bb][Jj][Ee][Cc][Tt])
          go test -v integrationTests/objectTests/*object*
          ;;
    *) echo "Invalid Test Flag used - only block,file or object allowed"
         ;;
esac


