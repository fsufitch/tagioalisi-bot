import org.springframework.boot.gradle.tasks.bundling.BootJar

group = "org.tagioalisi"

version = "0.0.1-SNAPSHOT"

plugins {
    java
    `java-library`
    kotlin("jvm") version "1.9.22"
    id("org.springframework.boot") version "3.2.3"
}

repositories { mavenCentral() }

dependencies {
    implementation("org.springframework.boot", "spring-boot-starter-parent", "3.2.3")
    implementation("org.springframework.boot", "spring-boot-starter", "3.2.3")
    implementation("org.jetbrains.kotlinx", "kotlinx-coroutines-core", "1.8.0")
    implementation("dev.kord", "kord-core", "0.13.1")
}

kotlin { jvmToolchain(21) }

tasks {
    named<BootJar>("bootJar") {
        archiveClassifier.set("boot")
        mainClass.set("org.tagioalisi.ApplicationMain")
    }
}
