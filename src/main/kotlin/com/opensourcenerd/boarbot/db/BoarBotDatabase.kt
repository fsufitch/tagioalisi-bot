package com.opensourcenerd.boarbot.db

import com.opensourcenerd.boarbot.common.ApplicationContext
import com.opensourcenerd.boarbot.common.DATABASE_MODULE_QNAME
import mu.KotlinLogging

class BoarBotDatabase(private val context: ApplicationContext): ApplicationContext.AppModule {
    override val qname = DATABASE_MODULE_QNAME
    private val logger = KotlinLogging.logger {}

    fun hello() {
        logger.info {"hello database"}
    }
}