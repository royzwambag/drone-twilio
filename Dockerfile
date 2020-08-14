FROM alpine
ADD twilio /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/twilio
