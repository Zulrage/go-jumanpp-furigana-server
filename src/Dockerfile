FROM golang:1.8-alpine

WORKDIR /usr/local/src/
# JUMANPP ßß

ARG JUMANPP_VERSION=1.01

RUN apk --no-cache add curl which xz bzip2 gcc bash g++ boost-dev make ;\
    curl -o /usr/local/src/jumanpp-${JUMANPP_VERSION}.tar.xz "http://lotus.kuee.kyoto-u.ac.jp/nl-resource/jumanpp/jumanpp-${JUMANPP_VERSION}.tar.xz" ;\
    tar Jxf jumanpp-${JUMANPP_VERSION}.tar.xz ;\
    cd jumanpp-${JUMANPP_VERSION} ;\
    ./configure ; make ; make install ;\
    cd ../ ; rm -rf jumanpp-${JUMANPP_VERSION} ;\
    rm -rf jumanpp-${JUMANPP_VERSION}.tar.xz ;\
    apk del curl which xz bzip2 gcc g++ make

CMD ["jumanpp"]
