FROM alpine:3.4
MAINTAINER Makoto Tajitsu

RUN apk -U add curl && apk -U add bash && rm -rf /var/cache/apk/*
RUN curl -sL http://install.freedom-man.com/spm | bash
ENTRYPOINT ["spm"]
