###########
# DEVCONTAINER
##########
FROM fedora:37 AS devcontainer

# Install stuff necessary for a reasonable CLI
COPY devcontainer-packages.txt /opt
RUN dnf install -y $(cat /opt/devcontainer-packages.txt) && \
    rm -rf /var/cache/dnf && \
    npm i -g node npm

# Set up the devcontainer user
RUN useradd -ms /bin/bash developer
RUN echo "developer ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/developer
WORKDIR /home/developer
USER developer

COPY ./.bashrc.d /home/developer/.bashrc.d
COPY ./.devcontainer-install-go-tools.sh /opt/go-tools.sh

RUN bash /opt/go-tools.sh

ENV PATH=/home/developer/go/bin:${PATH}

CMD ["echo", "devcontainer should have its command overridden by the IDE"]


###########
# DISCORDBOT BUILDER
##########
FROM devcontainer AS discordbot-builder

COPY --chown=developer discordbot /home/developer/tagioalisi-bot/discordbot
COPY --chown=developer proto /home/developer/tagioalisi-bot/proto
WORKDIR /home/developer/tagioalisi-bot/discordbot
RUN ./build.sh
RUN sudo mkdir -p /dist/bin && \
    sudo cp -r bin /dist/bin


##########
# DISCORDBOT RUNTIME
# This image is the ready-to-run container for launching the bot
##########
FROM alpine AS discordbot-runtime
WORKDIR /opt/tagioalisi-bot
RUN apk add gcompat

COPY --from=discordbot-builder /dist/bin/* .
COPY certs /certs

CMD ./tagi-migrate && ./tagi-bot


###########
# WEBAPP BUILDER
##########
FROM devcontainer AS webapp-builder

COPY --chown=developer ./webapp /home/developer/tagioalisi-bot/webapp
COPY --chown=developer ./proto /home/developer/tagioalisi-bot/proto
WORKDIR /home/developer/tagioalisi-bot/webapp
RUN npm install
RUN npm run proto && npm run build
RUN sudo mkdir -p /dist/webapp && \
    sudo cp -r dist/* /dist/webapp

##########
# WEBAPP RUNTIME
# This image runs a basic webserver exposing the web app
##########
FROM alpine:3 AS webapp-runtime

RUN apk add nodejs npm && \
    npm i -g http-server


COPY --from=webapp-builder /dist/webapp /webapp
COPY certs /certs
WORKDIR /webapp

# Needed for single page apps
RUN ln -s index.html 404.html

ENV WEBAPP_PORT=8080
ENV WEBAPP_TLS_CERT=/certs/default.pem
ENV WEBAPP_TLS_KEY=/certs/default.key

# See: https://www.npmjs.com/package/http-server
CMD http-server \
    --port ${WEBAPP_PORT} \
    -S -C ${WEBAPP_TLS_CERT} -K ${WEBAPP_TLS_KEY} \
    -d false && \
    -i false && \
    -e '' && \
    -c-1 && \
    .

##########
# GRPCWEBPROXY
# This image runs a proxy that converts a pure-gRPC endpoint to a gRPC-web endpoint
# See: https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy
##########

FROM golang:alpine AS grpcwebproxy

RUN go install -v github.com/improbable-eng/grpc-web/go/grpcwebproxy@latest

COPY ./grpcwebproxy /opt/grpcwebproxy
COPY certs /certs

ENV PROXY_BACKEND=localhost:9000
ENV PROXY_BACKEND_TLS=false

ENV PROXY_TLS_PORT=9001
ENV PROXY_TLS_CERT=/certs/default.pem
ENV PROXY_TLS_KEY=/certs/default.key

ENV PROXY_DEBUG_PORT=9002

CMD grpcwebproxy \
    --allow_all_origins \
    --use_websockets \
    --backend_addr ${PROXY_BACKEND} \
    --backend_tls=${PROXY_BACKEND_TLS} \
    --server_http_tls_port ${PROXY_TLS_PORT} \
    --server_tls_cert_file=${PROXY_TLS_CERT} \
    --server_tls_key_file=${PROXY_TLS_KEY}

##########
# DATABASE
##########

FROM postgres:alpine AS database

COPY certs /certs

VOLUME [ "/var/tagioalisi/db/pgdata" ]
ENV PGDATA=/var/tagioalisi/db/pgdata


##########
# STUB FOR DEFAULT BUILDS
##########

FROM scratch AS xxx_use_targets_please