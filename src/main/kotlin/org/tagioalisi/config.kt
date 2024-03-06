package org.tagioalisi


import org.springframework.context.annotation.Configuration
import org.springframework.context.annotation.Profile
import org.springframework.context.annotation.PropertySource
import org.springframework.core.env.Environment
import org.springframework.stereotype.Component


fun <T> Environment.get(key: String, targetType: Class<T>, missingValue: T? = null): T =
        this.getProperty(key, targetType) ?: missingValue ?: throw PropertyUndefinedException(key)

@Component
class TagioalisiConfiguration(private val environment: Environment) {
    val debugMode = environment.get("tagioalisi.debug", Boolean::class.java, false)
    val cleanupCommands = environment.get("tagioalisi.clean-commands", Boolean::class.java, false)

    val discordBotToken = environment.get("discord.bot-token", String::class.java)
    val discordApplicationId = environment.get("discord.application-id", String::class.java)
}

@Profile("env")
@Configuration
@PropertySource("classpath:application.properties")
@PropertySource("classpath:env-profile.properties")
open class EnvironmentProfilePropertiesConfiguration()

@Profile("dev")
@Configuration
@PropertySource("classpath:application.properties")
@PropertySource("classpath:dev-profile.properties")
open class DevelopmentProfilePropertiesConfiguration()


class PropertyUndefinedException(
        @Suppress("MemberVisibilityCanBePrivate") val propertyKey: String,
) : Exception("Application environment does not contain property: $propertyKey")

