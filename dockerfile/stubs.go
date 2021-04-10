package dockerfile

var (
	stub = `FROM golang:1.16-alpine

WORKDIR /src

ADD . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -o /service .


FROM scratch

COPY --from=0 /service service
COPY --from=gomicro/probe /probe probe

EXPOSE 4567

CMD ["/service"]
`
)
