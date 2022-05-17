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
RUN npm i -g npm yarn


# Install various developer tools
RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    echo


# Runtime
WORKDIR /home/developer
