# Dockerfile-tor
FROM alpine:latest

RUN apk --no-cache add tor

# Use the tor user
USER tor

COPY deploy/tor/torrc /etc/tor/torrc

CMD ["tor"]