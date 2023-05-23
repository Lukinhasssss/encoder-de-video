FROM golang:1.20-alpine3.18

# Install Bento
WORKDIR /tmp/bento4

ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:$BENTO4_BIN"

ENV BENTO4_BASE_URL="https://www.bok.net/Bento4/source/" \
    BENTO4_VERSION="1-6-0-640" \
    BENTO4_VERSION_FILE="1-6-0.640" \
    BENTO4_CHECKSUM="af3d8dd9cf54e97d5c427e918f200866a24d30dc" \
    BENTO4_PATH="/opt/bento4" \
    BENTO4_TYPE="SRC"

# Download and unzip Bento4
RUN apk add --update --upgrade curl ffmpeg python3 unzip bash gcc g++ make cmake scons && \
    curl -O -s ${BENTO4_BASE_URL}/Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}.zip && \
    sha1sum -b Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}.zip && \
    mkdir -p ${BENTO4_PATH} && \
    unzip Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}.zip -d ${BENTO4_PATH} && \
    rm -rf Bento4-${BENTO4_TYPE}-${BENTO4_VERSION_FILE}.zip && \
    apk del unzip && \
    cd ${BENTO4_PATH} && \
    mkdir bin utils && \
    cd ./bin  && cmake -DCMAKE_BUILD_TYPE=Release .. && cmake --build . --config Release && cd .. && \
    cp -R ${BENTO4_PATH}/Source/Python/utils ${BENTO4_PATH} && \
    cp -a ${BENTO4_PATH}/Source/Python/wrappers/. ${BENTO4_PATH}/bin

WORKDIR /go/src

ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:$BENTO4_BIN"

# Vamos mudar para o endpoint correto. Usando top apenas para manter o processo rodando
ENTRYPOINT [ "top" ]