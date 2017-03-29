#!/bin/bash


isPodUp(){
  	x=1
	while [ $x -le 15 ]
	do
	  if [ "$#" -eq 3 ]; then
	  	declare -i y=$(kubectl get pods -n $3 2>/dev/null  | grep $2 | wc -l )
	  else
	  	declare -i y=$(kubectl get pods 2>/dev/null  | grep $1 | wc -l )
	  fi
	  if [ $y -ge $2 ]
		then
			break
	 fi
	  x=$(( $x + 1 ))
	  sleep 10
	done

	if [ $x -gt 15 ]; then
		retval=1
	else
		retval=0
	fi

	return $retval

}

isPodRunning(){
  	x=1
	while [ $x -le 15 ]
	do
	  if [ "$#" -eq 3 ]; then
	  	declare -i y=$(kubectl get pods -n $3 2>/dev/null | grep $1 | awk '{print $3}' | grep Running | wc -l)
	  else
	  	declare -i y=$(kubectl get pod 2>/dev/null  | grep $1 | awk '{print $3}' | grep Running | wc -l)
	  fi
	  if [ $y -ge $2 ]
		then
			break
	 fi
	  x=$(( $x + 1 ))
	  sleep 10
	done

	if [ $x -gt 15 ]; then
		retval=1
	else
		retval=0
	fi

	return $retval
}


checkTestResult(){
	pod=$1
	x=1
	while [ $x -le 8 ]
	do
	  res=$(kubectl logs $pod|tail -2| head -1)
	  if [ "$res" == "PASS" ]
		then
			return 0
	 fi
	  x=$(( $x + 1 ))
	  sleep 20
	done
	return 1
}

checkTestResultInNameSpace(){
	pod=$1
	x=1
	while [ $x -le 8 ]
	do
	  res=$(kubectl logs $pod -n $2|tail -2| head -1)
	  if [ "$res" == "PASS" ]
		then
			return 0
	 fi
	  x=$(( $x + 1 ))
	  sleep 20
	done
	return 1
}




isTestPodUp(){
	TestPodName=$1
	if [ "$#" -eq 2 ]; then
		isPodUp $TestPodName 1 $2
	else
		isPodUp $TestPodName 1
	fi
	iretval=$?
	if [ "$retval" == 0 ]
	then
	     echo "$TestPodName Pod Started"
	else
	     echo "$TestPodName Pod Not Started"
	     return 1
	fi

	if [ "$#" -eq 2 ]; then
		isPodRunning $TestPodName 1 $2
	else
		isPodRunning $TestPodName 1
	fi
	iretval=$?
	if [ "$retval" == 0 ]
	then
	     echo "$TestPodName Running"
	else
	     echo "$TestPodName Not Running"
	     return 1
	fi


	return 0
}




##
##

case $1 in

    #Run Block Tests
    [Bb][Ll][Oo][Cc][Kk])
                echo "Create rook-storageclass"
                export MONS=$(kubectl -n rook get pod mon0 mon1 mon2 -o json|jq ".items[].status.podIP"|tr -d "\""|sed -e 's/$/:6790/'|paste -s -d, -)
                sed 's#INSERT_HERE#'$MONS'#' rook/rook-storageclass.yaml | kubectl create -f -
                sleep 10
		echo "Starting Block Test Pod"
		sed 's#INSERT_TEST_TYPE#'block'#' rook/block_test.yaml | kubectl create -f -
                testPod=$?
		if [ $testPod -ne 0 ]; then
			echo "Could'nt start testPod"
			exit 1
		fi
		isTestPodUp block-test
		TestPodUp=$?
		if [ $TestPodUp -ne 0 ]; then
			echo "Couldn't start Pod"
			exit 1
		fi
		checkTestResult block-test
		testRes=$?
		if [ $testRes == 0 ]
		then
			echo "Block End 2 End Test Passed"
		else
			echo "Block End 2 End Test Failed"
			kubectl delete -f rook/block_test.yaml
		    kubectl delete -f rook/rook-storageclass.yaml
			exit 1
		fi
		kubectl delete -f rook/block_test.yaml
		kubectl delete -f rook/rook-storageclass.yaml
		exit 0
         ;;
    #Run Object Test
    [Oo][Bb][Jj][Ee][Cc][Tt])
		echo "start Object Test Pod"
		kubectl exec -it rook-client -n rook -- rook object create
	        kubectl exec -it rook-client -n rook -- rook object user create rook-user "A rook rgw User"
                eval $(kubectl exec -it rook-client -n rook -- rook object connection rook-user --format env-var)
		echo $AWS_HOST
                echo $AWS_ENDPOINT
                echo $AWS_ACCESS_KEY_ID
		echo $AWS_SECRET_ACCESS_KEY
  		sed 's#INSERT_TEST_TYPE#'object'#;s#AWS_ENDPOINT_VALUE#'$AWS_ENDPOINT'#;s#AWS_KEY_VALUE#'$AWS_ACCESS_KEY_ID'#;s#AWS_SECRET_VALUE#'$AWS_SECRET_ACCESS_KEY'#' rook/object_test.yaml | kubectl create -f -
		testPod=$?
		if [ $testPod -ne 0 ]; then
			echo "Could'nt start testPod"
			exit 1
		fi
		isTestPodUp object-test rook
		TestPodUp=$?
		if [ $TestPodUp -ne 0 ]; then
			echo "Couldn't start Pod"
			exit 1
		fi
		checkTestResultInNameSpace object-test rook
		testRes=$?
		if [ $testRes == 0 ]
		then
			echo "Object End 2 End Test Passed"
		else
			echo "Object End 2 End Test Failed"
	    	kubectl delete -f rook/object_test.yaml
			exit 1
		fi
		kubectl delete -f rook/object_test.yaml
		exit 0
         ;;
    #Run File Test
    [Ff][Ii][Ll][Ee])
		echo "Start File Test Pod"
		kubectl exec -it rook-client -n rook -- rook filesystem create --name testfs
		export CEPH_MON0=$(kubectl -n rook get pod mon0 -o json|jq ".status.podIP"|tr -d "\""|sed -e 's/$/:6790/')
		export CEPH_MON1=$(kubectl -n rook get pod mon1 -o json|jq ".status.podIP"|tr -d "\""|sed -e 's/$/:6790/')
		export CEPH_MON2=$(kubectl -n rook get pod mon2 -o json|jq ".status.podIP"|tr -d "\""|sed -e 's/$/:6790/')
		sed 's#INSERT_TEST_TYPE#'file'#;s#CEPH_MON0#'$CEPH_MON0'#;s#CEPH_MON1#'$CEPH_MON1'#;s#CEPH_MON2#'$CEPH_MON2'#' rook/file_test.yaml | kubectl create -f -
                testPod=$?
		if [ $testPod -ne 0 ]; then
			echo "Could'nt start testPod"
			exit 1
		fi
		isTestPodUp file-test rook
		TestPodUp=$?
		if [ $TestPodUp -ne 0 ]; then
			echo "Couldn't start Pod"
			exit 1
		fi
		checkTestResultInNameSpace file-test rook
		testRes=$?
		if [ $testRes == 0 ]
		then
			echo "File End 2 End Test Passed"
		else
			echo "File End 2 End Test Failed"
			kubectl delete -f rook/file_test.yaml
			exit 1
		fi
		kubectl delete -f rook/file_test.yaml
		exit 0
         ;;
    *) echo "Invalid Test Flag used - only block,object or file allowed"
		exit 1
         ;;

esac