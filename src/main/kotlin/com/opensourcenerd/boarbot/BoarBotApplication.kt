package com.opensourcenerd.boarbot

import discord4j.core.DiscordClientBuilder
import discord4j.core.event.domain.lifecycle.ReadyEvent
import discord4j.core.event.domain.message.MessageCreateEvent


val DISCORD_BOT_TOKEN = System.getenv("DISCORD_BOT_TOKEN") ?: ""

fun main(args: Array<String>) {
    val client = DiscordClientBuilder(DISCORD_BOT_TOKEN).build()

    client.eventDispatcher.on(ReadyEvent::class.java)
            .subscribe { ev -> System.out.println("Logged in as " + ev.self.username) }

    client.eventDispatcher.on(MessageCreateEvent::class.java)
            .map {ev -> ev.message}
            .filter { msg -> msg.content.map { it.contains("!ping") }.orElse(false)}
            .flatMap { msg -> msg.channel}
            .flatMap { channel -> channel.createMessage("Pong!") }
            .subscribe()

    client.login().block()
}
