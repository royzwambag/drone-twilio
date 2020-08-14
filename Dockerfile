FROM alpine
ADD drone-twilio /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/drone-twilio
