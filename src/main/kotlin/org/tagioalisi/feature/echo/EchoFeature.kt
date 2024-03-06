package org.tagioalisi.feature.echo

import dev.kord.core.Kord
import dev.kord.core.behavior.interaction.response.respond
import dev.kord.core.behavior.interaction.suggestString
import dev.kord.rest.builder.interaction.string
import org.apache.commons.logging.LogFactory
import org.springframework.stereotype.Component
import org.tagioalisi.feature.SlashCommand
import org.tagioalisi.feature.SlashCommandContext
import org.tagioalisi.feature.TagioalisiFeature

@Component
class EchoFeature(kord: Kord) : TagioalisiFeature(kord) {
    private val log = LogFactory.getLog(javaClass)

    override val slashCommands = listOf(SlashCommand().apply {
        name = "echo2"
        context = SlashCommandContext.GUILD
        description = "Echo the received message back"
        ephemeral = true
        builder = {
            string("message", "message to echo") {
                required = true
                autocomplete = true
                minLength = 1
            }
        }

        interact = { interaction, response ->
            val message = interaction.command.strings["message"]!!
            response.respond {
                content = message
            }
        }

        autocomplete = { interaction ->
            val partialMessage = interaction.command.strings["message"]!!
            if ("hello world".startsWith(partialMessage.lowercase())) {
                interaction.suggestString { choice("hello world", "hello world") }
            }
        }
    })
}
