package gg.foreversixty.app.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * One character linked to the account: the character card and the switcher's row.
 * Mirrors api/internal/auth/store.go `Character`. [key] is opaque; never parse it, use
 * [region], [ruleset] and [name] to build character routes.
 */
@Serializable
data class Character(
    val key: String,
    val region: String,
    val ruleset: String,
    val name: String,
    @SerialName("class") val characterClass: String? = null,
    val realm: String? = null,
    val level: Int? = null,
    val faction: String? = null,
    val race: String? = null,
    val gender: String? = null,
    @SerialName("item_level") val itemLevel: Int? = null,
    @SerialName("avatar_url") val avatarUrl: String? = null,
    @SerialName("render_url") val renderUrl: String? = null,
    /** "export" (addon or companion) or "bnet" (Battle.net import). */
    val source: String,
    val build: CharacterBuild? = null,
    val guild: CharacterGuild? = null,
    /** The primary talent tree; null while not known. */
    val spec: String? = null,
)

/** api/internal/auth/store.go `CharacterBuild`: the last simmable build and its sync facts. */
@Serializable
data class CharacterBuild(
    val source: String,
    @SerialName("captured_at") val capturedAt: String,
    /** Slot name to item id. */
    val gear: Map<String, Int> = emptyMap(),
    val talents: BuildTalents? = null,
    val level: Int? = null,
    @SerialName("data_build") val dataBuild: String? = null,
    @SerialName("median_sync_gap_sec") val medianSyncGapSec: Int? = null,
    @SerialName("sync_error") val syncError: String? = null,
)

/** api/internal/auth/store.go `BuildTalents`: per-tree rank strings and points, in tree order. */
@Serializable
data class BuildTalents(
    val trees: List<String> = emptyList(),
    val points: List<Int> = emptyList(),
)

/** api/internal/auth/store.go `CharacterGuild`: the character's current guild membership. */
@Serializable
data class CharacterGuild(
    val id: Long,
    val name: String,
    val rank: String,
    @SerialName("rank_index") val rankIndex: Int? = null,
    val verified: Boolean = false,
    val attendance: Attendance? = null,
)

/** Raids attended over raids counted. */
@Serializable
data class Attendance(
    val present: Int,
    val nights: Int,
)
