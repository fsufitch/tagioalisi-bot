package org.tagioalisi.config

import java.util.Properties
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.annotation.Profile
import org.springframework.context.annotation.PropertySource
import org.springframework.stereotype.Component

typealias ProfileProperties = Properties

abstract class ProfilePropertiesProvider {
    @Autowired private lateinit var properties: Properties

    fun provide(): ProfileProperties {
        return properties
    }
}

@Profile("env")
@PropertySource("classpath:env-profile.properties")
@Component
class DefaultProfilePropertiesProvider : ProfilePropertiesProvider()

@Profile("dev")
@PropertySource("classpath:dev-profile.properties")
@Component
class DefaultPropetiesProvider : ProfilePropertiesProvider()
