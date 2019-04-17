package com.opensourcenerd.boarbot.db

import org.jetbrains.exposed.sql.insert
import org.jetbrains.exposed.sql.select
import org.jetbrains.exposed.sql.transactions.transaction
import org.jetbrains.exposed.sql.update
import org.joda.time.DateTime
import org.joda.time.DateTimeZone

fun BoarBotDatabase.kvGet(key: String): String? {
    var result: String? = null
    transaction(db) {
        val queryIterator = KV.select { KV.key eq key }.limit(1).iterator()
        result = if (queryIterator.hasNext()) queryIterator.next()[KV.value] else null
    }
    return result
}

fun BoarBotDatabase.kvSet(key: String, value: String) {
    val isUpdate = kvGet(key) != null

    if (!isUpdate) transaction(db) {
        KV.insert {
            it[KV.key] = key
            it[KV.value] = value
            it[KV.timestamp] = DateTime.now(DateTimeZone.UTC)
        }
    } else transaction(db) {
        KV.update({ KV.key eq key }) {
            it[KV.value] = value
            it[KV.timestamp] = DateTime.now(DateTimeZone.UTC)
        }
    }
}