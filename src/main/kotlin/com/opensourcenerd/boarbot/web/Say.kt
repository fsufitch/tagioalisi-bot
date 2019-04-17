package com.opensourcenerd.boarbot.web

import com.beust.klaxon.Klaxon
import io.undertow.server.HttpServerExchange
import io.undertow.util.Headers

data class SayPostBody(
        val channelId: String,
        val message: String
)

fun BoarBotWeb.say(exchange: HttpServerExchange) {
    exchange.startBlocking()
    val data = Klaxon().parse<SayPostBody>(exchange.inputStream)!!

    discordModule.sendMessage(data.channelId, data.message)

    exchange.statusCode = 200
    exchange.responseHeaders.put(Headers.CONTENT_TYPE, "text/plain")
    exchange.responseSender.send("Sent")
}
