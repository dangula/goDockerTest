 # Rook Test container
 This container is designe to run end to end tests on a Rook.
 The tests are designed to test specific storage volumes that were provisioned by rook are presnet and accessable(read and write data). The storage volume can be a bock storage, file storage or an object store. 
  
 You Should start the container/pod on a cluster running rook.
 ### Pre-requiste
  Running k8s cluster with Rook running. Tests are designed to run as pods in a k8s cluster with rook installed.
 
 #####  Instructions for Tests :
   1. check out this project on host running k8s cluster with Rook. 
   2. Use the start up script- TestPodsSetUp.sh  provided to run test - the script does the necesary set up and teardown for test. 
   3. The start up script takes a single paramater - the type of tests to run  - allowed values are block,file or object.
   4. The start up script does the following - 
      i. Does necessary set for storage privisioning - eg. create filesystem for mounting file system volume or creating storageClass for block volumes
      ii. Start up a test appropriate test pod and replace environment values
      iii. The test pods come up and run appropritae tests
      iv. The script makes sure the test pod is up and running and polls for resutls
      v. Script returns success exit code if test pass , if test fail or something goes wrong with test setup the script returns unsuccessful exit code. 
      
   
 
