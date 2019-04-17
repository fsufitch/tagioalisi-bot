import org.gradle.kotlin.dsl.support.unzipTo

val discord4jVersion  = project.property("discord4j.version") as String
val undertowVersion  = project.property("undertow.version") as String
val kotlinLoggingVersion = project.property("kotlin-logging.version") as String
var slf4jBridgeVersion = project.property("slf4j-log4j12.version") as String
val coroutinesVersion = project.property("kotlinx-coroutines.version") as String
val exposedVersion = project.property("exposed.version") as String
val postgresqlVersion = project.property("postgresql.version") as String
val klaxonVersion = project.property("klaxon.version") as String

plugins {
    application
    kotlin("jvm") version "1.3.30"
}

repositories {
    mavenCentral()
    maven("https://dl.bintray.com/kotlin/exposed")
    maven("https://dl.bintray.com/cbeust/maven")
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))
    implementation(kotlin("reflect"))
    implementation("com.discord4j", "discord4j-core", discord4jVersion)
    implementation("io.undertow", "undertow-core", undertowVersion)
    implementation("io.github.microutils", "kotlin-logging", kotlinLoggingVersion)
    implementation("org.slf4j", "slf4j-log4j12", slf4jBridgeVersion)
    implementation("org.jetbrains.kotlinx", "kotlinx-coroutines-core", coroutinesVersion)
    implementation("org.jetbrains.exposed", "exposed", exposedVersion)
    implementation("org.postgresql", "postgresql", postgresqlVersion)
    implementation("com.beust", "klaxon", klaxonVersion)
}

application {
    mainClassName = "com.opensourcenerd.boarbot.MainKt"
}



tasks {
    compileKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
    compileTestKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
}

tasks.register("stage") {
    dependsOn("build")
    doLast {
        mkdir("build/stage")
        unzipTo(File("build/stage/."), File("build/distributions/discord-boar-bot.zip"))
    }
}