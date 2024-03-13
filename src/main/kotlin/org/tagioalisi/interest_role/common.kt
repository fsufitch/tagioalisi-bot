package org.tagioalisi.interest_role

import dev.kord.core.behavior.GuildBehavior
import dev.kord.core.behavior.MemberBehavior
import dev.kord.core.entity.Role
import kotlinx.coroutines.flow.emptyFlow
import kotlinx.coroutines.flow.filter
import kotlinx.coroutines.flow.firstOrNull
import kotlinx.coroutines.flow.map
import org.springframework.beans.factory.annotation.Value
import org.springframework.stereotype.Component

const val DEFAULT_INTEREST_PREFIX = "g-"
const val VAL_INTEREST_PREFIX = "\${TAGIOALISI_INTEREST_PREFIX:$DEFAULT_INTEREST_PREFIX}"

@Component
class InterestRolesService(
    @Value(VAL_INTEREST_PREFIX) val prefix: String,
) {
    fun roleNameToInterestName(roleName: String) =
        if (roleName.startsWith(prefix)) roleName.drop(prefix.length) else null

    fun interestNameToRoleName(interestName: String) = "${prefix}${interestName}"

    fun guildInterests(guild: GuildBehavior) =
        guild.roles.map { roleNameToInterestName(it.name) ?: "" }.filter { it.isNotBlank() }

    suspend fun memberInterests(member: MemberBehavior) =
        member.asMemberOrNull()?.roles?.map { roleNameToInterestName(it.name) ?: "" }?.filter { it.isNotBlank() }
            ?: emptyFlow()

    suspend fun guildInterestRole(guild: GuildBehavior, interestName: String): Role? {
        val roleName = interestNameToRoleName(interestName)
        return guild.roles.filter { it -> it.name == roleName }.firstOrNull()
    }
}