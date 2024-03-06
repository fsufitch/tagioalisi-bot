package org.tagioalisi

import dev.kord.core.Kord
import kotlinx.coroutines.runBlocking
import org.apache.commons.logging.LogFactory
import org.springframework.context.annotation.Bean
import org.springframework.stereotype.Component
import org.tagioalisi.feature.CommandCleanup
import org.tagioalisi.feature.TagioalisiFeature

// See: https://discordapi.com/permissions.html#378228952256
const val JOIN_PERMISSIONS = 378228952256

@Component
class TagioalisiBot(
        private val kord: Kord,
        private val features: Collection<TagioalisiFeature>,
        private val cmdCleanup: CommandCleanup,
        configuration: TagioalisiConfiguration,
) {
    private val log = LogFactory.getLog(javaClass)

    private val inviteUrl = "https://discord.com/oauth2/authorize?client_id=${configuration.discordApplicationId}&scope=bot&permissions=${JOIN_PERMISSIONS}"

    suspend fun start() {
        cmdCleanup.globalAutoClean()
        features.forEach { it.start() }
        log.info("Starting bot listen loop... Invite URL: $inviteUrl")
        kord.login()
    }
}

@Component
class KordProvider(private val configuration: TagioalisiConfiguration) {
    private val botToken = configuration.discordBotToken

    @Bean
    fun initializeKord(): Kord {
        return runBlocking { Kord(botToken) }
    }
}
