package com.opensourcenerd.boarbot.web

import io.undertow.server.HttpServerExchange
import io.undertow.util.Headers

fun BoarBotWeb.notFound(exchange: HttpServerExchange) {
    exchange.statusCode = 404
    exchange.responseHeaders.put(Headers.CONTENT_TYPE, "text/plain")
    exchange.responseSender.send("Not found")
}

fun BoarBotWeb.unauthorized(exchange: HttpServerExchange){
    logger.warn { "Unauthorized: ${exchange.requestMethod} ${exchange.requestPath} by ${exchange.sourceAddress}" }

    exchange.statusCode = 401
    exchange.responseHeaders.put(Headers.CONTENT_TYPE, "text/plain")
    exchange.responseSender.send("401 Unauthorized")
}

fun BoarBotWeb.forbidden(exchange: HttpServerExchange){
    logger.warn { "Forbidden: ${exchange.requestMethod} ${exchange.requestPath} by ${exchange.sourceAddress}" }

    exchange.statusCode = 403
    exchange.responseHeaders.put(Headers.CONTENT_TYPE, "text/plain")
    exchange.responseSender.send("403 Forbidden")
}