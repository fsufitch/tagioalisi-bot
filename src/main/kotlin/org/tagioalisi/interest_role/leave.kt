package org.tagioalisi.interest_role

import com.kotlindiscord.kord.extensions.checks.memberFor
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
class SubCommandLeave(private val interests: InterestRolesService) {

    inner class LeaveArguments : Arguments() {
        val interestName = string {
            name = "interest"
            description = "special interest group name"

            autoComplete { event ->
                val interests = memberFor(event)?.let { interests.memberInterests(it) } ?: emptyFlow()
                suggestString {
                    interests.filter { it.lowercase().startsWith(focusedOption.value.lowercase()) }
                        .collect { choice(it, it) }
                }
            }
        }
    }

    suspend fun addSubCommand(parent: EphemeralSlashCommand<*, *>) =
        parent.ephemeralSubCommand(arguments = ::LeaveArguments) {
            name = "leave"
            description = "leave a special interest role"

            action {
                val member = this.member
                requireNotNull(member) { "did not find member in command context" }
                val interestName = arguments.interestName.parsed

                try {
                    performLeave(member, interestName)
                } catch (exc: IllegalArgumentException) {
                    respond {
                        embed {
                            title = "Failed to leave interest group '${interestName}'"
                            description = exc.localizedMessage
                        }
                    }
                    return@action
                }

                respond {
                    content = "${member.mention} left interest group '${interestName}'"
                }
            }
        }

    private suspend fun performLeave(member: MemberBehavior, interestName: String) {
        val role = interests.guildInterestRole(member.guild, interestName)
        requireNotNull(role) { "no such role: ${interests.interestNameToRoleName(interestName)}" }
        require(
            member.asMemberOrNull()?.roles?.filter { it.id != role.id }?.firstOrNull() !== null
        ) { "${member.mention} is already not assigned role: ${interests.interestNameToRoleName(interestName)}" }
        member.removeRole(role.id)
    }
}