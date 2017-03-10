FROM golang

CMD go get -v github.com/dangula/goDockerTest && echo $(pwd) && cd src/github.com/dangula/goDockerTest && echo $(pwd) && ./startTest.sh sleep 120
