FROM golang

CMD go get -v github.com/dangula/goDockerTest && echo $(pwd) && cd src/github.com/dangula/goDockerTest && echo $(pwd) && chmod +x startTest.sh && ./startTest.sh && sleep 120
