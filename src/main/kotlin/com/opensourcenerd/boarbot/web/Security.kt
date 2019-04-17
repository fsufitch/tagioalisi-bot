package com.opensourcenerd.boarbot.web

import io.undertow.server.HttpServerExchange

fun BoarBotWeb.authorized(wrapped: (HttpServerExchange)->Unit): (HttpServerExchange)->Unit {
    return fun(ex: HttpServerExchange) {
        if (ex.requestHeaders["Authorization"]?.isEmpty() != false ||
                ex.requestHeaders["Authorization"].first.isEmpty()) return unauthorized(ex)

        val authString = ex.requestHeaders["Authorization"].first
        if (!authString.startsWith("Bearer ")) return forbidden(ex)

        val token = authString.substring(7)

        if (token != context.configuration.webSecret) return forbidden(ex)

        return wrapped(ex)
    }
}
