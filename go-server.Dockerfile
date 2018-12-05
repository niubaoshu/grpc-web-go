FROM alpine:3.8

WORKDIR /helloworld

COPY ./server ./server
CMD ["./server"]