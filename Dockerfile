FROM alpine
ADD go-category /go-category
ENTRYPOINT [ "/go-category" ]
