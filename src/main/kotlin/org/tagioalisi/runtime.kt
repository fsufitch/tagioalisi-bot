package org.tagioalisi

import kotlinx.coroutines.runBlocking
import org.apache.commons.logging.LogFactory
import org.springframework.boot.CommandLineRunner
import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication

class Main {
    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            SpringApplication.run(TagioalisiRuntime::class.java, *args)
        }
    }
}


@SpringBootApplication
open class TagioalisiRuntime(
    private val debug: DebugScope,
    private val tagioalisiBot: TagioalisiBot,
    ) : CommandLineRunner {
    private val logger = LogFactory.getLog(javaClass)
    override fun run(vararg args: String?) {
        debug {
            logger.info("debug mode is enabled")
        }
        runBlocking {
            tagioalisiBot.start()
        }
    }
}