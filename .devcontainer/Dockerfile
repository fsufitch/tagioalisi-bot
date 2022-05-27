FROM debian:sid AS devcontianer

# Install stuff necessary for a reasonable CLI
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
        bash \
        bash-completion \
        build-essential \
        curl \
        file \
        git \
        gnupg \
        inetutils-ping \
        man-db \
        net-tools \
        openssh-client \
        python3 \
        protobuf-compiler \
        sudo \
        vim \
        xz-utils


# Set up the devcontainer user
RUN useradd -ms /bin/bash -G sudo developer
RUN echo "developer ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/developer
USER developer
RUN mkdir -p /home/developer/opt

# Install Go
ARG GO_DOWNLOAD=https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
WORKDIR /home/developer/opt
RUN curl -L -o go.tar.gz ${GO_DOWNLOAD} && \
    tar xfz go.tar.gz && \
    rm go.tar.gz
RUN mkdir -p /home/developer/opt/goenv
ENV GOPATH=/home/developer/opt/goenv
ENV PATH=${PATH}:/home/developer/opt/goenv/bin:/home/developer/opt/go/bin


# Install Node
ARG NODE_DOWNLOAD=https://nodejs.org/dist/v16.15.0/node-v16.15.0-linux-x64.tar.xz
WORKDIR /home/developer/opt
RUN curl -Lo node.tar.xz ${NODE_DOWNLOAD} && \
    tar xfJ node.tar.xz && \
    rm node.tar.xz && \
    mv node-* node
ENV PATH=${PATH}:/home/developer/opt/node/bin
RUN npm i -g npm && \
    corepack enable yarn && \
    yarn set version berry && \
    yarn set version latest


# Install various developer tools
RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    echo

# Install the client-side GRPC code generator (protoc-gen-grpc-web)
RUN mkdir -p /home/developer/opt/bin && \
    os=$(uname -s | tr '[:upper:]' '[:lower:]') && \
    arch=$(uname -m | tr '[:upper:]' '[:lower:]' ) && \
    curl -L \
       -o /home/developer/opt/bin/protoc-gen-grpc-web \
       https://github.com/grpc/grpc-web/releases/download/1.3.1/protoc-gen-grpc-web-1.3.1-${os}-${arch} \
       && \
    chmod a+x /home/developer/opt/bin/protoc-gen-grpc-web && \
    echo
ENV PATH=${PATH}:/home/developer/opt/bin


# Runtime
WORKDIR /home/developer
