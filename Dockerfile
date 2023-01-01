FROM ubuntu:18.04

LABEL maintainer="nautilis"

RUN mkdir -p /root/app/netease_music_bot
RUN apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends ca-certificates curl

WORKDIR /root/app/netease_music_bot
COPY netease_music_bot .
COPY run.sh .
COPY monitor.sh .
COPY conf ./conf

CMD ["sh", "./run.sh"]
