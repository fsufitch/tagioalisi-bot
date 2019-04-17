package com.opensourcenerd.boarbot.db

import com.opensourcenerd.boarbot.common.ApplicationContext
import com.opensourcenerd.boarbot.common.DATABASE_MODULE_QNAME
import mu.KotlinLogging
import org.jetbrains.exposed.sql.Database

class BoarBotDatabase(val context: ApplicationContext): ApplicationContext.AppModule {
    override val qname = DATABASE_MODULE_QNAME
    val logger = KotlinLogging.logger {}

    val db = Database.connect(context.configuration.databaseUrl, "org.postgresql.Driver")

    init {
        logger.info("Database initialized")
    }
}