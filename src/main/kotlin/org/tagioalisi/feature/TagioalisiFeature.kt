package org.tagioalisi.feature

import dev.kord.core.Kord
import kotlinx.coroutines.runBlocking
import org.apache.commons.logging.LogFactory
import kotlin.concurrent.thread


abstract class TagioalisiFeature(@Suppress("MemberVisibilityCanBePrivate") protected val kord: Kord) {
    abstract val slashCommands: Collection<SlashCommand>
    private var started = false
    suspend fun start() {
        if (started) {
            throw Exception("tried to start feature multiple times")
        }

        slashCommands.forEach { cmd ->
            cmd.register(kord)
            cmd.listen(kord)
        }
    }
}
