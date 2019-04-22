FROM gcr.io/distroless/base

COPY bin/rds-combined-ca-bundle.pem /bin/rds-combined-ca-bundle.pem
COPY bin/milmove /bin/milmove
COPY bin/chamber /bin/chamber

COPY config /config
COPY swagger/* /swagger/
COPY build /build

ENTRYPOINT ["/bin/milmove"]

CMD ["serve", "--debug-logging"]

EXPOSE 8080
