FROM golang

RUN curl https://glide.sh/get | sh

CMD go get -v github.com/dangula/goDockerTest && cd src/github.com/dangula/goDockerTest && chmod +x startTest.sh && ./startTest.sh && sleep 180
