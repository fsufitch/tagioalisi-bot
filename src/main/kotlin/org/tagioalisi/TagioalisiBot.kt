package org.tagioalisi

import dev.kord.core.Kord
import kotlinx.coroutines.runBlocking
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.annotation.Bean
import org.springframework.stereotype.Component
import org.tagioalisi.feature.TagioalisiFeature

@Component
class TagioalisiBot(
        private val botConfig: TagioalisiBotConfiguration,
        private val discordConnectionConfiguration: DiscordConnectionConfiguration,
        private val kord: Kord,
) {
    @Autowired
    lateinit var features: Array<TagioalisiFeature>


    suspend fun start() {
        features.forEach { it.start() }
        kord.login()
    }

}

@Component
class KordProvider(private val discordConnectionConfiguration: DiscordConnectionConfiguration) {
    @Bean
    fun initializeKord(): Kord {
        return runBlocking { Kord(discordConnectionConfiguration.botToken) }
    }
}
