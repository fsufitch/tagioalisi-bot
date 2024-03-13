package org.tagioalisi.interest_role

import com.kotlindiscord.kord.extensions.extensions.ephemeralSlashCommand
import com.kotlindiscord.kord.extensions.extensions.event
import dev.kord.common.entity.Snowflake
import dev.kord.core.event.guild.GuildCreateEvent
import org.springframework.stereotype.Component
import org.tagioalisi.TagioalisiBotExtension


@Component
class InterestRoleExtension(
    private val join: SubCommandJoin,
    private val leave: SubCommandLeave,
    private val create: SubCommandCreate,
) : TagioalisiBotExtension() {
    override val name = "role-tag"


    override suspend fun setup() {
        event<GuildCreateEvent> {
            action {
                registerCommands(event.guild.id)
            }
        }
    }

    private suspend fun registerCommands(guildId: Snowflake) =
        ephemeralSlashCommand {
            name = "interest"
            description = "interact with special interest roles"
            guild(guildId)
            
            join.addSubCommand(this)
            leave.addSubCommand(this)

            create.addSubCommand(this)
        }
}
