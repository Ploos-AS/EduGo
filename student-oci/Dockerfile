FROM golang:1.25-bookworm

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
       ca-certificates \
       curl \
       git \
       make \
       gcc \
       g++ \
       libc6-dev \
       bash \
       less \
       vim-tiny \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

CMD ["bash"]
