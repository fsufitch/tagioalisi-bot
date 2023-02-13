###########
# BASE DEVELOPMENT IMAGE
##########
FROM fedora:37 AS devcontainer_base

# Install stuff necessary for a reasonable CLI
RUN dnf install -y \
        bash-completion \
        curl \
        file \
        git \
        gzip \
        gnupg2 \
        iputils \
        man-db \
        net-tools \
        openssh-clients \
        python3 \
        protobuf-compiler \
        sudo \
        tar \
        vim \
        xz

# Set up the devcontainer user
RUN useradd -ms /bin/bash developer
RUN echo "developer ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/developer

###########
# BOT DEVCONTAINER
##########

# Install Go
ARG GO_DOWNLOAD=https://go.dev/dl/go1.20.linux-amd64.tar.gz
WORKDIR /opt
RUN curl -L -o go.tar.gz ${GO_DOWNLOAD} && \
    tar xfz go.tar.gz && \
    rm go.tar.gz
ENV PATH=${PATH}:/opt/go/bin

##########
# WEB UI DEVCONTAINER
##########

# Install Node
ARG NODE_DOWNLOAD=https://nodejs.org/dist/v19.6.0/node-v19.6.0-linux-x64.tar.xz
WORKDIR /opt
RUN curl -Lo node.tar.xz ${NODE_DOWNLOAD} && \
    tar xfJ node.tar.xz && \
    rm node.tar.xz && \
    mv node-* node
ENV PATH=${PATH}:/opt/node/bin
RUN npm i -g npm

# Install the client-side GRPC code generator (protoc-gen-grpc-web)
# RUN mkdir -p /home/developer/opt/bin && \
#     os=$(uname -s | tr '[:upper:]' '[:lower:]') && \
#     arch=$(uname -m | tr '[:upper:]' '[:lower:]' ) && \
#     curl -L \
#        -o /home/developer/opt/bin/protoc-gen-grpc-web \
#        https://github.com/grpc/grpc-web/releases/download/1.3.1/protoc-gen-grpc-web-1.3.1-${os}-${arch} \
#        && \
#     chmod a+x /home/developer/opt/bin/protoc-gen-grpc-web && \
#     echo
# ENV PATH=${PATH}:/home/developer/opt/bin


# Runtime
WORKDIR /home/developer
