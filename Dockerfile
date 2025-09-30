FROM ubuntu:24.04
ARG TARGETPLATFORM

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY $TARGETPLATFORM/halsecur /halsecur
RUN chmod +x /halsecur

ENTRYPOINT ["/halsecur"]