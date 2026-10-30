package gg.foreversixty.app.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * GET /v1/characters/{region}/{ruleset}/{name}: the public character card with the Forever Score
 * and the gear panel (docs/contracts/app-v1.md "Characters"). Every score field is null for a
 * character with no rated fight: render "no logs yet", never 0. Picks, gaps and scores are the
 * API's rules; the app never re-derives them. Feature lanes add the fight lists when a screen
 * reads them.
 */
@Serializable
data class CharacterProfile(
    val character: CharacterHeader,
    val score: Double? = null,
    val tier: String? = null,
    @SerialName("as_of") val asOf: String? = null,
    val components: ScoreComponents? = null,
    val gear: List<GearRow> = emptyList(),
    @SerialName("gear_gap_total") val gearGapTotal: Double? = null,
    @SerialName("gear_spec") val gearSpec: String? = null,
    val supporter: Boolean = false,
)

/** Who the card is for. */
@Serializable
data class CharacterHeader(
    val key: String,
    val region: String,
    val ruleset: String,
    val name: String,
    @SerialName("class") val characterClass: String? = null,
    val level: Int? = null,
    val race: String? = null,
    val realm: String? = null,
    val guild: GuildLine? = null,
    @SerialName("avatar_url") val avatarUrl: String? = null,
    @SerialName("render_url") val renderUrl: String? = null,
)

@Serializable
data class GuildLine(val name: String, val rank: String)

/** The score's five parts. */
@Serializable
data class ScoreComponents(
    val gear: Double,
    val experience: Double,
    val parses: Double,
    val attendance: Double,
    val discipline: Double,
)

/** One slot of the gear panel: what is worn, the band's pick and the points it adds. */
@Serializable
data class GearRow(
    val slot: String,
    val worn: WornItem? = null,
    val pick: PickedItem? = null,
    val gap: Double? = null,
)

/** [name] and [itemLevel] are null when the data build does not know the item. */
@Serializable
data class WornItem(
    val id: Int,
    val name: String? = null,
    @SerialName("item_level") val itemLevel: Int? = null,
)

/** [chance] is the drop fraction, null for a source with no rate. */
@Serializable
data class PickedItem(
    val id: Int,
    val name: String,
    val source: String,
    val chance: Double? = null,
)
