# ===== Image for building tagi web UI
FROM node:lts-alpine AS ui-builder

# Python needed to be sure node-sass works
RUN apk add python3 && \
    npm config set python /usr/bin/python3

# Use the latest yarn @2+
RUN corepack enable yarn && \
    yarn set version berry && \
    yarn set version latest

WORKDIR /web
COPY .yarnrc.yml package.json yarn.lock ./
RUN yarn install

COPY . .

ARG WEBPACK_MODE=development
ENV API_ENDPOINT=http://localhost:8081 \
    WEBPACK_MODE=${WEBPACK_MODE}
RUN npx webpack --mode ${WEBPACK_MODE}

RUN [ "echo", "this shouldn't be run"]


# ===== Image for running a super simple webserver with the UI
FROM nginx:alpine AS ui-server

COPY --from=ui-builder /web/dist /usr/share/nginx/html
COPY nginx-default.conf /etc/nginx/conf.d/default.conf