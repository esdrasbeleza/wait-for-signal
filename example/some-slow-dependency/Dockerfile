FROM alpine

RUN wget https://github.com/esdrasbeleza/wait-for-signal/raw/master/bin/send-signal-linux-amd64-0.1 -O /tmp/send-signal
RUN chmod +x /tmp/send-signal

COPY ./run.sh .
ENTRYPOINT ["./run.sh"]
