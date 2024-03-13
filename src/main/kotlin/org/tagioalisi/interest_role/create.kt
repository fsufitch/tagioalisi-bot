package org.tagioalisi.interest_role

import com.kotlindiscord.kord.extensions.checks.hasPermission
import com.kotlindiscord.kord.extensions.commands.Arguments
import com.kotlindiscord.kord.extensions.commands.application.slash.EphemeralSlashCommand
import com.kotlindiscord.kord.extensions.commands.application.slash.publicSubCommand
import com.kotlindiscord.kord.extensions.commands.converters.impl.string
import com.kotlindiscord.kord.extensions.utils.any
import dev.kord.common.entity.Permission
import dev.kord.core.behavior.createRole
import dev.kord.rest.builder.message.embed
import org.springframework.stereotype.Component

val CREATE_INTEREST_PERMISSION = Permission.ManageRoles

@Component
class SubCommandCreate(private val interests: InterestRolesService) {

    inner class CreateArguments : Arguments() {
        val interestName = string {
            name = "interest"
            description = "special interest group name"
        }
    }

    suspend fun addSubCommand(parent: EphemeralSlashCommand<*, *>) =
        parent.publicSubCommand (arguments = ::CreateArguments) {
            name = "create"
            description = "create a special interest role"
            requirePermission(CREATE_INTEREST_PERMISSION)

            check {
                hasPermission(CREATE_INTEREST_PERMISSION)
            }

            action {
                val interestName = arguments.interestName.parsed
                val roleName = interests.interestNameToRoleName(interestName)
                val guild = guild ?: throw IllegalStateException("could not find a guild for this action")

                if (guild.roles.any { it.name == roleName }) {
                    respond {
                        embed {
                            title = "Failed to create special interest role"
                            description = "Role already exists: $roleName"
                        }
                    }
                    return@action
                }


                val role = guild.createRole {
                    name = interests.interestNameToRoleName(interestName)
                    mentionable = true
                }

                respond {
                    content = "Created role: ${role.mention}"
                }
            }
        }
}