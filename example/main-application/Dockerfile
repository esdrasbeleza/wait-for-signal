FROM golang:latest AS builder

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /hello-world .

FROM alpine
EXPOSE 80

COPY --from=builder /hello-world .
COPY ./run.sh .

RUN wget https://github.com/esdrasbeleza/wait-for-signal/raw/master/bin/wait-for-signal-linux-amd64-0.1 -O /tmp/wait-for-signal
RUN chmod +x /tmp/wait-for-signal

ENTRYPOINT ["./run.sh"]
