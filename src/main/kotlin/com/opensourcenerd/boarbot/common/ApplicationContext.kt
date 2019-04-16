package com.opensourcenerd.boarbot.common

import mu.KotlinLogging

class ApplicationContext(val configuration: Configuration) {
    private val logger = KotlinLogging.logger {}
    private val _appModules = mutableMapOf<String, AppModule>()
    private val _botModules = mutableMapOf<String, BotModule>()

    val appModules: Set<Any> get() = _appModules.values.toSet()
    val botModules: Set<Any> get() = _botModules.values.toSet()

    fun addAppModule(module: AppModule) {
        if (_appModules.containsKey(module.qname)) {
            throw BootstrapException("Module of this qname already registered: ${module.qname}")
        }

        _appModules[module.qname] = module
    }

    fun addBotModule(module: BotModule) {
        if (_botModules.containsKey(module.qname)) {
            throw BootstrapException("Bot module of this qname already registered: ${module.qname}")
        }

        _botModules[module.qname] = module
    }

    fun getAppModule(qname: String): AppModule? {
        return _appModules[qname]
    }


    fun getBotModule(qname: String): BotModule? {
        return _botModules[qname]
    }

    class BootstrapException(details: String): Exception(details)

    interface AppModule {
        val qname: String get
    }

    interface BotModule {
        val qname: String get
    }
}