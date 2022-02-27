FROM alpine:latest
MAINTAINER "SpotMax"
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && apk add -U --no-cache tzdata ca-certificates curl bash bash-completion shadow && \
    cp  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY main /spotmax/
ENV PATH="/spotmax:${PATH}"
WORKDIR /spotmax/
CMD ["./main"]
