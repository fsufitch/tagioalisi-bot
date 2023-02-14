#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd );
cd $SCRIPT_DIR;

HTTPS_KEY=${SCRIPT_DIR}/keys/tagioalisi-dev.key
HTTPS_CERT=${SCRIPT_DIR}/keys/tagioalisi-dev.pem

grpcwebproxy \
    --allow_all_origins \
    --backend_addr localhost:8092 \
    --backend_tls=false \
    --server_http_debug_port 8093 \
    --server_http_tls_port 8082 \
    --server_tls_cert_file=${HTTPS_CERT} \
    --server_tls_key_file=${HTTPS_KEY} \
    --use_websockets