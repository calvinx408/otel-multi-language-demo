FROM golang as builder

COPY . /code
WORKDIR /code
ENV GOPROXY=https://proxy.golang.org
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./frontend ./cmd

FROM scratch

COPY --from=builder /code/frontend /
EXPOSE 8080
CMD ["/frontend"]
