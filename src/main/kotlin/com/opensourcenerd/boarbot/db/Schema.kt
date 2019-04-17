package com.opensourcenerd.boarbot.db

import org.jetbrains.exposed.sql.*

object KV : Table() {
    override val tableName = "kv"

    val key = varchar("key", 512).uniqueIndex()
    val value = varchar("value", 512)
    val timestamp = datetime("timestamp")
}

object RoleAcl: Table() {
    override val tableName = "role_acl"

    val rowId = integer("row_id").autoIncrement().primaryKey()
    val aclId = varchar("acl_id", 512)
    val roleId = varchar("role_id", 512)
    val details = varchar("details", 512)

    init {
        index(true , aclId, roleId)
    }
}

object UserAcl: Table() {
    override val tableName = "user_acl"

    val rowId = integer("row_id").autoIncrement().primaryKey()
    val aclId = varchar("acl_id", 512)
    val userId = varchar("user_id", 512)
    val details = varchar("details", 512)

    init {
        index(true , aclId, userId)
    }
}

object Memes: Table() {
    override val tableName = "memes"

    val id = integer("id").autoIncrement("memes_id_seq").primaryKey()
}

object MemeNames: Table() {
    override val tableName = "meme_names"

    val id = integer("id").autoIncrement("meme_names_id_seq").primaryKey()
    val name = varchar("name", 512).uniqueIndex()
    val timestamp = datetime("timestamp")
    val author = varchar("author", 512)
    val memeId = integer("meme_id").references(Memes.id, onDelete = ReferenceOption.CASCADE)
}

object MemeUrls: Table() {
    override val tableName = "meme_urls"

    val id = integer("id").autoIncrement("meme_urls_id_seq").primaryKey()
    val url = varchar("url", 512)
    val timestamp = datetime("timestamp")
    val author = varchar("author", 512)
    val memeId = integer("meme_id").references(Memes.id, onDelete = ReferenceOption.CASCADE)
}