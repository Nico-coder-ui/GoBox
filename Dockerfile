FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    bash coreutils \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /sandbox

USER nobody

CMD ["/bin/bash"]
