package org.tagioalisi

import com.kotlindiscord.kord.extensions.commands.Arguments
import com.kotlindiscord.kord.extensions.commands.converters.impl.string
import com.kotlindiscord.kord.extensions.extensions.ephemeralSlashCommand
import com.kotlindiscord.kord.extensions.extensions.event
import com.kotlindiscord.kord.extensions.utils.suggestStringCollection
import dev.kord.common.entity.Snowflake
import dev.kord.core.event.guild.GuildCreateEvent
import org.slf4j.LoggerFactory
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration


@Configuration
open class EchoExtensionProvider {
    @Bean
    open fun getEchoExtension() = EchoExtension()
}

class EchoExtension : TagioalisiBotExtension() {
    private val logger = LoggerFactory.getLogger(javaClass)
    override val name: String = "echo"

    private suspend fun setupEchoCommand(guildId: Snowflake) =
        ephemeralSlashCommand(arguments = ::EchoArguments) {
            name = "echo"
            description = "respond with whatever text was received"
            guild(guildId)

            action {
                respond {
                    content = "Echo! ${arguments.echoMessageText}"
                }
            }
        }

    override suspend fun setup() {
        event<GuildCreateEvent> {
            action {
                event.guild.getApplicationCommands().collect { cmd ->
                    cmd.delete()
                }
                logger.info("setting up echo for guild ${event.guild.id} ('${event.guild.name}')")
                setupEchoCommand(event.guild.id)
            }
        }
    }

    inner class EchoArguments : Arguments() {
        val echoMessageText by string {
            name = "message"
            description = "the message to echo"
            autoComplete {
                val possibleOptions = listOf("Hello world!", "Test message please ignore", "Test message please do not ignore")
                val validOptions = possibleOptions
                    .filter { it.lowercase().startsWith(this.focusedOption.value.lowercase()) }
                suggestStringCollection(validOptions)
            }
        }
    }
}
