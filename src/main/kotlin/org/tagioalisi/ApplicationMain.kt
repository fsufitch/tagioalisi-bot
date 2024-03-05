package org.tagioalisi

import kotlinx.coroutines.runBlocking
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.CommandLineRunner
import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication

@SpringBootApplication
open class ApplicationMain : CommandLineRunner {
    @Autowired private lateinit var bot: TagioalisiBot

    override fun run(vararg args: String?) = runBlocking {
        if (listOfNotNull(*args).isNotEmpty()) {
            throw Exception("bot is configured via environment, not args")
        }
        bot.start()
    }
    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            SpringApplication.run(ApplicationMain::class.java, *args)
        }
    }
}
