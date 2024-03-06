package org.tagioalisi.feature

import dev.kord.common.entity.Snowflake
import dev.kord.core.Kord
import dev.kord.core.entity.application.GlobalApplicationCommand
import dev.kord.core.entity.application.GuildApplicationCommand
import kotlinx.coroutines.flow.filter
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Component
import org.tagioalisi.TagioalisiConfiguration

@Component
class CommandCleanup(private val kord: Kord, private val config: TagioalisiConfiguration) {
    private val log = LoggerFactory.getLogger(javaClass)

    @Suppress("MemberVisibilityCanBePrivate")
    suspend fun deleteGlobalApplicationCommands(filter: (GlobalApplicationCommand) -> Boolean = { true }) {
        kord.getGlobalApplicationCommands().filter { filter(it) }.collect {
                    log.info("deleting global application command ${it.name} (id=${it.id})")
                    it.delete()
                }
    }

    @Suppress("MemberVisibilityCanBePrivate")
    suspend fun deleteGuildApplicationCommands(guildId: Snowflake, filter: (GuildApplicationCommand) -> Boolean = { true }) {
        kord.getGuildApplicationCommands(guildId).filter { filter(it) }.collect {
                    log.info("deleting application command ${it.name} (id=${it.id} guild=${guildId})")
                    it.delete()
                }
    }

    suspend fun globalAutoClean() {
        if (!config.cleanupCommands) {
            log.info("Not cleaning up pre-existing global command configurations")
            return
        }
        deleteGlobalApplicationCommands()
    }

    suspend fun guildAutoClean(guildId: Snowflake) {
        if (!config.cleanupCommands) {
            log.info("Not cleaning up pre-existing command configurations (guild=${guildId})")
            return
        }
        deleteGuildApplicationCommands(guildId)
    }

}