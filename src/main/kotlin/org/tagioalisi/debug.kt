package org.tagioalisi

import org.springframework.beans.factory.annotation.Value
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

const val VAL_DEBUG_MODE = "\${DEBUG_MODE:false}"

interface DebugScope {
    operator fun invoke(block: () -> Unit)
}

@Configuration
open class DebugScopeProvider {
    @Bean
    open fun getDebugScope(@Value(VAL_DEBUG_MODE) debug: Boolean): DebugScope = if (debug) object : DebugScope {
        override fun invoke(block: () -> Unit) {}
    } else object : DebugScope {
        override fun invoke(block: () -> Unit) = block()
    }
}