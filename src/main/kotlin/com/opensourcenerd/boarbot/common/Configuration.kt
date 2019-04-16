package com.opensourcenerd.boarbot.common

data class Configuration(
        val webEnabled: Boolean,
        val webPort: Int,
        val webSecret: String,
        val discordToken: String,
        private val blacklistBotModulesString: String
) {
    val blacklistBotModules = blacklistBotModulesString.split(",").toSet().filter { !it.isEmpty() }

    init {
        if (discordToken.isEmpty()) throw ConfigurationException("No discord token found")

        if (webEnabled) {
            if (webPort == 0) throw ConfigurationException("Web enabled, but no port found")
            if (webSecret.isEmpty()) throw ConfigurationException("Web enabled, but no secret found")
        }
    }

    class ConfigurationException(details: String): Exception(details)
}

