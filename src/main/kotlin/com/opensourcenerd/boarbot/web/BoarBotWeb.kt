package com.opensourcenerd.boarbot.web

import com.opensourcenerd.boarbot.common.ApplicationContext
import com.opensourcenerd.boarbot.common.DISCORD_MODULE_QNAME
import com.opensourcenerd.boarbot.common.WEB_MODULE_QNAME
import com.opensourcenerd.boarbot.discord.BoarBotDiscord
import io.undertow.Undertow
import io.undertow.server.RoutingHandler
import mu.KotlinLogging


class BoarBotWeb(val context: ApplicationContext): ApplicationContext.AppModule {
    override val qname = WEB_MODULE_QNAME
    val logger = KotlinLogging.logger {}
    val discordModule = context.getAppModule(DISCORD_MODULE_QNAME) as BoarBotDiscord

    fun start() {
        val server = Undertow.builder()
                .addHttpListener(context.configuration.webPort, "0.0.0.0")
                .setHandler { it.dispatch(routes) }
                .build()
        server.start()
    }

    private val routes: RoutingHandler = RoutingHandler()
            .post("/say", authorized { say(it)  })
            .setFallbackHandler { notFound(it) }

}

