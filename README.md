 #Rook Test container
 
 This container is designe to run end to end tests on start up and report the results.
 The tests are designed to read and write data to specified folder. A folder
 that was privisioned  using Rook.
  
 You Should start the container/pod on a cluster running rook.
 
 #####  Instructions for Running Block Strorgae Tests :
  On a cluster running rook, do the following
  1. CreateStorage:  class so Rook can provision storage
     *  Export Rook Moitors First
      export MONS=$(kubectl -n rook get pod mon0 mon1 mon2 -o json|jq ".items[].status.podIP"|tr -d "\""|sed -e 's/$/:6790/'|paste -s -d, -)
     * Create Strorage Class : 
       sed 's#INSERT_HERE#'$MONS'#' rook-storageclass.yaml | kubectl create -f -
     
  
  2. Start Rook Test Container
     * kubectl create -f block_test.yaml 
  3. Look at Pod logs to see if tests pased or failed
     * kubectl logs block-test
 
