FROM debian:jessie
RUN groupadd -r paymentservice && useradd --no-log-init -r -g paymentservice paymentservice

FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /etc/group /etc/group
COPY --chown=paymentservice:paymentservice ./bin/api /paymentservice

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181

USER paymentservice

CMD [ "/paymentservice" ]