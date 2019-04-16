package com.opensourcenerd.boarbot.web

import com.opensourcenerd.boarbot.common.ApplicationContext
import com.opensourcenerd.boarbot.common.DISCORD_MODULE_QNAME
import com.opensourcenerd.boarbot.common.WEB_MODULE_QNAME
import com.opensourcenerd.boarbot.discord.BoarBotDiscord
import mu.KotlinLogging
import io.undertow.Undertow
import io.undertow.util.Headers


class BoarBotWeb(private val context: ApplicationContext): ApplicationContext.AppModule {
    override val qname = WEB_MODULE_QNAME
    private val logger = KotlinLogging.logger {}
    private val discordModule = context.getAppModule(DISCORD_MODULE_QNAME) as BoarBotDiscord

    fun start() {
        val server = Undertow.builder()
                .addHttpListener(context.configuration.webPort, "0.0.0.0")
                .setHandler { exchange ->
                    exchange.responseHeaders.put(Headers.CONTENT_TYPE, "text/plain")
                    exchange.responseSender.send("Hello World")

                    discordModule.sendMessage("327526752203177984", "web trigger!")
                }.build()
        server.start()
    }

}

