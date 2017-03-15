#!/bin/bash
echo "Running Test : "$TestName
#Install and update depenencies
glide install
glide update

# Run Test
case $TestName in

    #Run Block Tests
    [Bb][Ll][Oo][Cc][Kk])
          go test -v integrationTests/*block*
          ;;
    *) echo "Invalid Test Flag used"
         ;;

esac


