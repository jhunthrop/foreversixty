package gg.foreversixty.app.rankings.logic

import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.CharacterFight
import gg.foreversixty.app.model.CharacterRank
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.PhaseBoundary
import gg.foreversixty.app.model.RankedEncounter
import gg.foreversixty.app.model.RankingRow
import gg.foreversixty.app.reports.logic.ParseTier
import java.time.Instant
import kotlin.math.roundToInt

/** What a board is measured on (spec F11.2 `Measure`): the API's metric and the unit word under a value. */
enum class Measure(val query: String, @param:StringRes val label: Int, @param:StringRes val unit: Int, @param:StringRes val spoken: Int) {
    Damage("dps", R.string.rank_measure_damage, R.string.rank_unit_dps, R.string.rank_a11y_perSecond_dps),
    Healing("hps", R.string.rank_measure_healing, R.string.rank_unit_hps, R.string.rank_a11y_perSecond_hps),
    Taken("damage_taken", R.string.rank_measure_taken, R.string.rank_unit_dtps, R.string.rank_a11y_perSecond_dtps),
    Execution("execution", R.string.rank_measure_execution, R.string.rank_unit_execution, R.string.rank_a11y_executed),
    ;

    companion object {
        /** The metric a role is judged on: damage dealt, healing done, damage taken by a tank (web defaults.ts). */
        fun forRole(role: String?): Measure = when (role) {
            "healer" -> Healing
            "tank" -> Taken
            else -> Damage
        }

        fun of(metric: String): Measure? = entries.firstOrNull { it.query == metric }
    }
}

/** A class and spec the board is about (`shaman`, `Elemental`). */
data class Cohort(val classSlug: String, val specName: String)

/** A board's filters (spec F11.2). An empty string is "all" for ruleset, region, phase and faction. */
data class RankingsFilters(
    /** The encounter's slug; empty until the encounters are read. */
    val encounter: String,
    val measure: Measure,
    val ruleset: String,
    val region: String,
    val phase: String,
    val faction: String,
    val todayOnly: Boolean,
    /** The class and spec; null is every class (signed out, before a pick). */
    val cohort: Cohort?,
)

/** One `GET /v1/rankings` call, every "all" left out. */
data class RankingsQuery(
    val encounter: String,
    val metric: String,
    val classSlug: String?,
    val spec: String?,
    val phase: String?,
    val region: String?,
    val ruleset: String?,
    val faction: String?,
    val since: String?,
    val distinct: String,
    val character: String?,
    val page: Int,
)

/** A character of the signed-in account, as the board reads it. */
data class ViewerCharacter(
    val key: String,
    val name: String,
    val region: String,
    val ruleset: String,
    /** The class slug (`shaman`). */
    val classSlug: String?,
    /** The primary tree's name (`Elemental`). */
    val specName: String?,
) {
    companion object {
        fun of(character: AccountCharacter) = ViewerCharacter(
            character.key, character.name, character.region, character.ruleset, character.className?.lowercase(), character.spec,
        )
    }
}

/** Who looks at the rankings: the signed-in account and its current character, or nobody. */
data class RankingsViewer(
    val signedIn: Boolean,
    val current: ViewerCharacter?,
    /** Every character of the account: each one's row wears the `You` tag. */
    val accountKeys: Set<String>,
    /** `entitlements.history`: your own character page shows `Parse history`. */
    val history: Boolean,
) {
    companion object {
        val SignedOut = RankingsViewer(signedIn = false, current = null, accountKeys = emptySet(), history = false)

        fun of(me: Me, currentKey: String?): RankingsViewer = RankingsViewer(
            signedIn = true,
            current = me.characters.firstOrNull { it.key == currentKey }?.let(ViewerCharacter::of),
            accountKeys = me.characters.map { it.key }.toSet(),
            history = me.entitlements.history,
        )
    }
}

/** Where the viewer's own row goes (web row-view.ts ownRowPlacement). */
sealed interface OwnRow {
    data object None : OwnRow

    /** On a loaded page: marked where it is; [index] is its place in the loaded rows. */
    data class InPlace(val index: Int, val rank: Int) : OwnRow

    /** Off the loaded pages: pinned first under `Your best kill`, with its true rank. */
    data class Pinned(val rank: CharacterRank) : OwnRow
}

/** The rules under Rankings (spec F11), ported from web/src/lib/rankings: data in, data out. */
object Rankings {
    const val PER_PAGE = 100

    /** The top of the legend: a board's first rank is its 100th percentile and still reads pink. */
    private const val TOP_LEGEND_PERCENTILE = 99.0
    private const val HUNDRED = 100.0
    private const val TODAY = "today"
    private const val DISTINCT = "character"

    val rulesets = listOf("normal", "pvp", "rp", "hardcore", "")
    val regions = listOf("", "us", "eu", "kr", "tw", "cn")
    val factions = listOf("", "alliance", "horde")

    fun ordinal(n: Int): String = gg.foreversixty.app.raids.ui.ordinal(n)

    /** A rank's percentile within its board: first is 100, last is just above 0. */
    fun rankPercentile(rank: Int, total: Int): Double = HUNDRED - (rank - 1).toDouble() / maxOf(total, 1) * HUNDRED

    /** The rank numeral's parse colour, capped at 99 so the first rank is pink, never gold. */
    fun rankTier(rank: Int, total: Int): ParseTier =
        ParseTier.of(rankPercentile(rank, total).coerceIn(0.0, TOP_LEGEND_PERCENTILE).roundToInt())

    /** A board spanning more than one class shows a crest before each name; one class shows none. */
    fun showsCrests(rows: List<RankingRow>): Boolean = rows.map { it.player.className?.lowercase() }.toSet().size > 1

    fun placement(rows: List<RankingRow>, characterRank: CharacterRank?): OwnRow {
        characterRank ?: return OwnRow.None
        val index = rows.indexOfFirst { it.player.key == characterRank.row.player.key }
        return if (index >= 0) OwnRow.InPlace(index, characterRank.rank) else OwnRow.Pinned(characterRank)
    }

    fun isYours(row: RankingRow, viewer: RankingsViewer): Boolean = row.player.key in viewer.accountKeys

    /** The loaded rows and a next page, one row per rank (a page read twice adds nothing). */
    fun merge(rows: List<RankingRow>, next: List<RankingRow>): List<RankingRow> =
        (rows + next).distinctBy { it.rank }.sortedBy { it.rank }

    /** The phase a moment falls in: the last phase whose start is not after it (web phases.ts). */
    fun phaseAt(phases: List<PhaseBoundary>, at: Instant): String? =
        phases.filter { start(it)?.isAfter(at) == false }.maxByOrNull { start(it) ?: Instant.MIN }?.name

    /** The phase chips: every phase begun by [now], newest first. */
    fun phaseChoices(phases: List<PhaseBoundary>, now: Instant): List<String> =
        phases.filter { start(it)?.isAfter(now) == false }.sortedByDescending { start(it) }.map { it.name }

    private fun start(phase: PhaseBoundary): Instant? = runCatching { Instant.parse(phase.start) }.getOrNull()

    /**
     * The encounter a board opens on: the one the character most recently killed, else the first
     * the API lists, else none. Only a listed encounter is named (web defaults.ts pickDefaultEncounter).
     */
    fun defaultEncounter(history: List<CharacterFight>, encounters: List<RankedEncounter>): String? {
        val killed = history.filter { it.kill != false }.sortedByDescending { it.foughtAt }
        return killed.firstNotNullOfOrNull { fight -> encounters.firstOrNull { it.id == fight.encounterId.toLong() }?.slug }
            ?: encounters.firstOrNull()?.slug
    }

    /** The character's own cohort: its class and its spec's catalogue name. */
    fun cohort(character: ViewerCharacter, specs: SpecCatalog): Cohort? {
        val slug = character.classSlug ?: return null
        val name = character.specName?.let { specs.spec(slug, it)?.name ?: it }.orEmpty()
        return Cohort(slug, name)
    }

    /** Every default for the viewer (web defaults.ts rankingsDefaults): signed out is every class, every ruleset, damage. */
    fun defaults(viewer: RankingsViewer, specs: SpecCatalog, phases: List<PhaseBoundary>, now: Instant, encounter: String): RankingsFilters {
        val character = viewer.current
        val cohort = character?.let { cohort(it, specs) }
        val role = cohort?.let { c -> specs.spec(c.classSlug, c.specName)?.role }
        return RankingsFilters(
            encounter = encounter,
            measure = Measure.forRole(role),
            ruleset = character?.ruleset.orEmpty(),
            region = "",
            phase = phaseAt(phases, now).orEmpty(),
            faction = "",
            todayOnly = false,
            cohort = cohort,
        )
    }

    fun query(filters: RankingsFilters, characterKey: String?, page: Int) = RankingsQuery(
        encounter = filters.encounter,
        metric = filters.measure.query,
        classSlug = filters.cohort?.classSlug?.ifEmpty { null },
        spec = filters.cohort?.specName?.ifEmpty { null },
        phase = filters.phase.ifEmpty { null },
        region = filters.region.ifEmpty { null },
        ruleset = filters.ruleset.ifEmpty { null },
        faction = filters.faction.ifEmpty { null },
        since = if (filters.todayOnly) TODAY else null,
        distinct = DISTINCT,
        character = characterKey,
        page = page,
    )

    /** The value as the board prints it: whole per-second figures; execution as a percent. */
    fun valueText(row: RankingRow, measure: Measure): String = when (measure) {
        Measure.Execution -> "${(row.value * HUNDRED).roundToInt()}%"
        else -> row.value.roundToInt().toString()
    }
}
