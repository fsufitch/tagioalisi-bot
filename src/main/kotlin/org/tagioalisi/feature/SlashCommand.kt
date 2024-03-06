package org.tagioalisi.feature

import dev.kord.common.entity.InteractionType
import dev.kord.core.Kord
import dev.kord.core.behavior.interaction.respondPublic
import dev.kord.core.behavior.interaction.response.DeferredMessageInteractionResponseBehavior
import dev.kord.core.entity.application.ApplicationCommand
import dev.kord.core.entity.interaction.AutoCompleteInteraction
import dev.kord.core.entity.interaction.ChatInputCommandInteraction
import dev.kord.core.event.guild.GuildCreateEvent
import dev.kord.core.event.interaction.AutoCompleteInteractionCreateEvent
import dev.kord.core.event.interaction.ChatInputCommandInteractionCreateEvent
import dev.kord.core.on
import dev.kord.rest.builder.interaction.ChatInputCreateBuilder
import kotlinx.coroutines.Job
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableSharedFlow
import org.apache.commons.logging.LogFactory


open class SlashCommand {
    private val log = LogFactory.getLog(javaClass)

    lateinit var name: String
    lateinit var description: String
    lateinit var builder: ChatInputCreateBuilder.() -> Unit
    var context: SlashCommandContext = SlashCommandContext.GUILD
    var ephemeral: Boolean = false

    var interact: suspend (ChatInputCommandInteraction, DeferredMessageInteractionResponseBehavior) -> Unit =
        { _, _ -> }
    var autocomplete: suspend (AutoCompleteInteraction) -> Unit = { }

    suspend fun register(kord: Kord): Flow<ApplicationCommand> {
        val outputFlow = MutableSharedFlow<ApplicationCommand>(replay = 5)
        if (context.isGlobal) {
            log.info("register global command '$name'")
            val cmd = kord.createGlobalChatInputCommand(name, description, builder)
            outputFlow.emit(cmd)
        }

        if (context.isGuild) {
            kord.on<GuildCreateEvent> {
                log.info("register guild command '$name' (guild=${guild.id})")
                val cmd = kord.createGuildChatInputCommand(guild.id, name, description, builder)
                outputFlow.emit(cmd)
            }
        }
        return outputFlow
    }

    suspend fun listen(kord: Kord): Job {
        kord.on<AutoCompleteInteractionCreateEvent> {
            autocomplete(interaction)
        }

        return kord.on<ChatInputCommandInteractionCreateEvent> {
            log.info("Received event: $this")
            if (interaction.invokedCommandName != name) {
                return@on
            }

            if (interaction.type != InteractionType.ApplicationCommand) {
                throw Exception("unsupported interaction type: ${interaction.type::type.name}")
            }

            try {
                val response =
                    if (ephemeral) interaction.deferEphemeralResponse() else interaction.deferPublicResponse()
                interact(interaction, response)
            } catch (exc: Exception) {
                interaction.respondPublic {
                    content = "⚠ Unexpected exception executing `${name}` command: $exc"
                }
            } finally {
                val responseMessage = interaction.getOriginalInteractionResponseOrNull()
                if (responseMessage == null) {
                    interaction.respondPublic {
                        content = "⚠ Command `${name}` ran successfully but did not produce a response. This is a bug."
                    }
                }
            }
        }
    }
}


enum class SlashCommandContext {
    GUILD, GLOBAL, GUILD_GLOBAL;

    val isGuild: Boolean by lazy { this == GUILD || this == GUILD_GLOBAL }
    val isGlobal: Boolean by lazy { this == GLOBAL || this == GUILD_GLOBAL }
}