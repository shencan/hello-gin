FROM alpine:latest
MAINTAINER "SpotMax"
RUN set -eux && \
    cp  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY main /spotmax/
ENV PATH="/spotmax:${PATH}"
WORKDIR /spotmax/
CMD ["./main"]
