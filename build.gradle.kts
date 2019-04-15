plugins {
    application
    kotlin("jvm") version "1.3.30"
}

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    implementation("com.discord4j:discord4j-core:3.0.2")
}

application {
    mainClassName = "com.opensourcenerd.boarbot.BoarBotApplicationKt"
}