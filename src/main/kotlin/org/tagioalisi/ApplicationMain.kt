package org.tagioalisi


import kotlinx.coroutines.runBlocking
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.CommandLineRunner
import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.context.annotation.Bean

@SpringBootApplication
open class ApplicationMain : CommandLineRunner {

    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            SpringApplication.run(ApplicationMain::class.java, *args)
        }
    }

    @Autowired
    private lateinit var bot: TagioalisiBot

    override fun run(vararg args: String?) = runBlocking {
        if (listOfNotNull(*args).isNotEmpty()) {
            throw Exception("bot is configured via environment, not args")
        }
        bot.start()
    }
}