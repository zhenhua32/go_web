FROM golang:1.13 as build

ENV GOPROXY="https://goproxy.io"
# https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
# build for static link
ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app
RUN make build

# production stage
FROM alpine as production

WORKDIR /app
COPY ./conf/ /app/conf
COPY --from=build /app/web /app
EXPOSE 8081
ENTRYPOINT ["/app/web"]
CMD [ "-c", "./conf/config_docker.yaml" ]
