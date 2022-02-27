FROM alpine:latest
MAINTAINER "SpotMax"
COPY main /spotmax/
ENV PATH="/spotmax:${PATH}"
WORKDIR /spotmax/
CMD ["./main"]
