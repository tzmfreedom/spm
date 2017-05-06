FROM alpine:3.4
MAINTAINER Makoto Tajitsu

RUN apk -U add curl bash && \
   rm -rf /var/cache/apk/* && \
   curl -sL http://install.freedom-man.com/spm | bash
ENTRYPOINT ["spm"]
