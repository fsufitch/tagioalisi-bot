package org.tagioalisi.interest_role

import com.kotlindiscord.kord.extensions.checks.guildFor
import com.kotlindiscord.kord.extensions.commands.Arguments
import com.kotlindiscord.kord.extensions.commands.application.slash.EphemeralSlashCommand
import com.kotlindiscord.kord.extensions.commands.application.slash.ephemeralSubCommand
import com.kotlindiscord.kord.extensions.commands.converters.impl.string
import dev.kord.core.behavior.MemberBehavior
import dev.kord.core.behavior.interaction.suggestString
import dev.kord.rest.builder.message.embed
import kotlinx.coroutines.flow.emptyFlow
import kotlinx.coroutines.flow.filter
import kotlinx.coroutines.flow.firstOrNull
import org.springframework.stereotype.Component


@Component
class SubCommandJoin(private val interests: InterestRolesService) {

    inner class JoinArguments : Arguments() {
        val interestName = string {
            name = "interest"
            description = "special interest group name"

            autoComplete { event ->
                val interests = guildFor(event)?.let { interests.guildInterests(it) } ?: emptyFlow()
                suggestString {
                    interests.filter { it.lowercase().startsWith(focusedOption.value.lowercase()) }
                        .collect { choice(it, it) }
                }
            }
        }
    }

    suspend fun addSubCommand(parent: EphemeralSlashCommand<*, *>) =
        parent.ephemeralSubCommand(arguments = ::JoinArguments) {
            name = "join"
            description = "join a special interest role"

            action {
                val member = this.member
                requireNotNull(member) { "did not find member in command context" }
                val interestName = arguments.interestName.parsed

                try {
                    performJoin(member, interestName)
                } catch (exc: IllegalArgumentException) {
                    respond {
                        embed {
                            title = "Failed to join interest group '${interestName}'"
                            description = exc.localizedMessage
                        }
                    }
                    return@action
                }

                respond {
                    content = "${member.mention} joined interest group '${interestName}'"
                }
            }
        }

    private suspend fun performJoin(member: MemberBehavior, interestName: String) {
        val role = interests.guildInterestRole(member.guild, interestName)
        requireNotNull(role) { "no such role: ${interests.interestNameToRoleName(interestName)}" }
        require(
            member.asMemberOrNull()?.roles?.filter { it.id != role.id }?.firstOrNull() === null
        ) { "${member.mention} already has the role: ${interests.interestNameToRoleName(interestName)}" }
        member.addRole(role.id)
    }
}