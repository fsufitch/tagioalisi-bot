#!/bin/bash

# Create a config using environment vars to substitute stuff in it, since Nginx can't do that natively
cat ./tagioalisi.nginx_template.conf \
    | envsubst \
    > /etc/nginx/conf.d/tagioalisi.conf

# Command from original nginx:alpine image
exec nginx -g daemon off