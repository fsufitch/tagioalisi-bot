package org.tagioalisi

import org.springframework.context.annotation.Configuration
import org.springframework.context.annotation.Profile
import org.springframework.context.annotation.PropertySource


@Profile("default")
@Configuration
@PropertySource("classpath:application.properties")
open class DefaultProfilePropertiesConfiguration()
