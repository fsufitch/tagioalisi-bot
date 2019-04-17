package com.opensourcenerd.boarbot.discord

import com.opensourcenerd.boarbot.common.ApplicationContext
import com.opensourcenerd.boarbot.common.DATABASE_MODULE_QNAME
import com.opensourcenerd.boarbot.common.DISCORD_MODULE_QNAME
import com.opensourcenerd.boarbot.db.BoarBotDatabase
import discord4j.core.DiscordClient
import discord4j.core.DiscordClientBuilder
import discord4j.core.`object`.entity.Channel
import discord4j.core.`object`.entity.MessageChannel
import discord4j.core.`object`.util.Snowflake
import discord4j.core.event.domain.lifecycle.ReadyEvent
import discord4j.core.event.domain.message.MessageCreateEvent
import mu.KotlinLogging

class BoarBotDiscord(private val context: ApplicationContext): ApplicationContext.AppModule{
    override val qname = DISCORD_MODULE_QNAME
    private val db = context.getAppModule(DATABASE_MODULE_QNAME) as BoarBotDatabase
    private val logger = KotlinLogging.logger {}
    private val botToken = context.configuration.discordToken

    private var client: DiscordClient = DiscordClientBuilder(botToken).build()

    init {
        client.eventDispatcher.on(ReadyEvent::class.java)
                .subscribe { ev -> logger.info("Logged in as ${ev.self.username}") }

        client.eventDispatcher.on(MessageCreateEvent::class.java)
                .map {ev -> ev.message}
                .filter { msg -> msg.content.map { it.contains("!ping") }.orElse(false)}
                .flatMap { msg -> msg.channel}
                .flatMap { channel -> channel.createMessage("Pong!") }
                .subscribe()
    }

    fun start() {
        client.login().block()
    }

    fun sendMessage(channelID: String, message: String) {
        client.getChannelById(Snowflake.of(channelID))
                .filter { ch -> ch.type in listOf(Channel.Type.DM, Channel.Type.GUILD_TEXT, Channel.Type.GROUP_DM)}
                .map {ch -> ch as MessageChannel }
                .flatMap { ch -> ch.createMessage(message)}
                .subscribe()
    }
}