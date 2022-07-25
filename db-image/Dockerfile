# syntax=docker/dockerfile:1

FROM postgres:14-bullseye

ENV POSTGRES_PASSWORD="password"
ARG PG_SEMVER_COMMIT=d69c87e1f074f54f07d9279586e12b43b72b9649

RUN echo 'apt::install-recommends "false";' >> /etc/apt/apt.conf.d/01-no-install-recommends && \
    apt-get update && \
    apt-get install -f -y --no-install-recommends \
        ca-certificates \
        software-properties-common \
        build-essential \
        pkg-config \
        git-core \
        postgresql-server-dev-$PG_MAJOR && \
    git clone https://github.com/theory/pg-semver.git && \
    cd pg-semver && \
    git checkout $PG_SEMVER_COMMIT && \
    gmake && \
    gmake install

