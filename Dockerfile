FROM golang:1.15

ENV GO111MODULE=on
RUN go get -u github.com/deukyunlee/protocol-camp
COPY ./ ./src
WORKDIR src
RUN go build -o protocol-camp
# 포트는 443 https 오픈
EXPOSE 443

# 실행할 애플리케이션
CMD ["./42report_calender_server"]