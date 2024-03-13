package org.tagioalisi

import com.kotlindiscord.kord.extensions.ExtensibleBot
import com.kotlindiscord.kord.extensions.extensions.Extension
import dev.kord.core.event.guild.GuildCreateEvent
import kotlinx.coroutines.flow.buffer
import kotlinx.coroutines.runBlocking
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Value
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.context.annotation.Profile


const val VAL_BOT_TOKEN = "\${DISCORD_BOT_TOKEN:notset}"

typealias TagioalisiBot = ExtensibleBot
typealias TagioalisiBotExtension = Extension

@Configuration
open class TagioalisiBotProvider(
    @Value(VAL_BOT_TOKEN) private val token: String,
    private val tagiExtensions: Collection<TagioalisiBotExtension>,
) {
    private val logger = LoggerFactory.getLogger(javaClass)

    @Profile("!clear")
    @Bean
    open fun getBot(): TagioalisiBot = runBlocking {
        ExtensibleBot(token) {
            val bot = this  // To clarify scope
            tagiExtensions.forEach { ext -> bot.extensions { add { ext } } }
        }
    }

    @Profile("clear")
    @Bean
    open fun getClearingBot(): TagioalisiBot = runBlocking {
        println("clearing mode")
        val bot = ExtensibleBot(token) {

        }

        bot.on<GuildCreateEvent> {
            this.kord.getGlobalApplicationCommands().buffer().collect {
                logger.info("delete global application command ${it.id} (${it.name})")
                it.delete()
            }
            guild.getApplicationCommands().buffer().collect {
                logger.info("delete global application command ${it.id} (${it.name}) from guild ${guild.id} ($guild.name)")
                it.delete()
            }
        }

        bot
    }
}
