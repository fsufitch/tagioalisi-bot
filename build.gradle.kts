plugins {
    application
    kotlin("jvm") version "1.3.30"
}

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
}

application {
    mainClassName = "com.opensourcenerd.boarbot.BoarBotApplicationKt"
}