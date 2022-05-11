# ===== Image for building tagi web UI
FROM node:lts-alpine AS ui-builder

# Python needed to be sure node-sass works
RUN apk add python3
RUN npm config set python /usr/bin/python3

WORKDIR /web
COPY package.json package-lock.json ./
RUN npm install

COPY . .

ENV API_ENDPOINT=http://localhost:8081 \
    WEBPACK_MODE=development
RUN npx webpack --mode $WEBPACK_MODE

RUN [ "echo", "this shouldn't be run"]

# ===== Simple webserver to run the UI
FROM nginx:alpine as ui-server

COPY --from=ui-builder /web/dist /usr/share/nginx/html
