package org.tagioalisi

import org.springframework.beans.factory.annotation.Value
import org.springframework.stereotype.Component


@Component
class TagioalisiBotConfiguration(
        @Value("\${tagioalisi.debug:false}") val debugMode: Boolean,
)

@Component
class DiscordConnectionConfiguration(
        @Value("\${discord.bot-token}") val botToken: String,
        @Value("\${discord.application-id}") val applicationId: String,
)
