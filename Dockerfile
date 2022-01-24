#Stage-1
FROM golang:1.16 as builder
COPY ./*.go ./
RUN go build -o /sample-go-app ./*.go

#Stage-2
FROM gcr.io/distroless/base
EXPOSE 8080
COPY --from=builder /sample-go-app /.
ENTRYPOINT ["./sample-go-app"]