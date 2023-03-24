FROM scratch AS final

COPY ./bin/ports /ports

WORKDIR /

ENTRYPOINT ["/ports"]

EXPOSE 8080