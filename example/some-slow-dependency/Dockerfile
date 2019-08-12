FROM alpine

RUN wget https://github.com/esdrasbeleza/wait-for-signal/releases/download/v0.1/send-signal-linux-amd64-v0.1 -O /tmp/send-signal
RUN chmod +x /tmp/send-signal

COPY ./run.sh .
ENTRYPOINT ["./run.sh"]
