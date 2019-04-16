package com.opensourcenerd.boarbot

import com.opensourcenerd.boarbot.common.ApplicationContext
import com.opensourcenerd.boarbot.common.Configuration
import com.opensourcenerd.boarbot.common.DISCORD_MODULE_QNAME
import com.opensourcenerd.boarbot.common.WEB_MODULE_QNAME
import com.opensourcenerd.boarbot.db.BoarBotDatabase
import com.opensourcenerd.boarbot.discord.BoarBotDiscord
import com.opensourcenerd.boarbot.web.BoarBotWeb
import mu.KotlinLogging
import kotlinx.coroutines.*
import kotlinx.coroutines.channels.Channel
import java.lang.RuntimeException

val defaultConfiguration = Configuration(
        webPort = System.getenv("PORT")?.toIntOrNull() ?: 0,
        webEnabled = System.getenv("WEB_ENABLED")?.toBoolean() ?: false,
        webSecret = System.getenv("WEB_SECRET") ?: "",
        discordToken = System.getenv("DISCORD_TOKEN") ?: "",
        blacklistBotModulesString = System.getenv("BLACKLIST_MODULES") ?: ""
)

val logger = KotlinLogging.logger {}

fun main() {
    runBlocking {
        val appContext = ApplicationContext(defaultConfiguration)

        appContext.addAppModule(BoarBotDatabase(appContext))
        appContext.addAppModule(BoarBotDiscord(appContext))
        if (appContext.configuration.webEnabled) {
            appContext.addAppModule(BoarBotWeb(appContext))
        }

        val errChan = Channel<Exception>()


        if (appContext.configuration.webEnabled) {
            launch {
                try {
                    (appContext.getAppModule(WEB_MODULE_QNAME) as BoarBotWeb).start()
                } catch (e: Exception) {
                    errChan.send(e)
                }
                errChan.send(RuntimeException("Web runtime quit prematurely"))
            }
        }

        launch {
            try {
                (appContext.getAppModule(DISCORD_MODULE_QNAME) as BoarBotDiscord).start()
            } catch (e: Exception) {
                errChan.send(e)
            }
            errChan.send(RuntimeException("Discord runtime quit prematurely"))
        }

        throw errChan.receive()
    }
}
