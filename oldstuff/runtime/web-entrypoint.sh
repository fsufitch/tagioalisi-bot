#!/bin/sh

set -x;

# Create a config using environment vars to substitute stuff in it, since Nginx can't do that natively
cat ./tagioalisi.nginx_template.conf \
    | sed "s#%%BOT_EXTERNAL_BASE_URL%%#${BOT_EXTERNAL_BASE_URL}#" \
    | sed "s#%%BOT_EXTERNAL_GRPC_URL%%#${BOT_EXTERNAL_GRPC_URL}#" \
    | sed "s#%%NGINX_PORT%%#${NGINX_PORT}#" \
    | tee /etc/nginx/conf.d/tagioalisi.conf

# Command from original nginx:alpine image
exec nginx -g 'daemon off;'