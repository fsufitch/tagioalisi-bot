package org.tagioalisi

import org.springframework.boot.CommandLineRunner
import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication

@SpringBootApplication
open class TagioalisiMain : CommandLineRunner {
    companion object {
        @JvmStatic
        fun main(args: Array<String>) {

            SpringApplication.run(TagioalisiMain::class.java, *args)
        }
    }

    override fun run(vararg args: String?) {

        println("hello ${listOf(*args).joinToString(separator = ", ") }")
    }
}