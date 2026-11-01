package gg.foreversixty.app.rankings.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.model.CharacterRank
import gg.foreversixty.app.model.RankedEncounter
import gg.foreversixty.app.model.RankingRow
import gg.foreversixty.app.model.RankingsPage
import gg.foreversixty.app.phrase
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.reports.logic.FightRows
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale
import kotlin.math.roundToInt

/** The answer card's meta: the sentence, and the character name it opens with (drawn in gold), if any. */
data class AnswerMeta(val text: String, val name: String?)

/** Every F11 sentence built from data (spec §5.3); the words themselves live in the copy catalog. */
class RankingsCopy(private val res: Resources, private val zone: ZoneId = ZoneId.systemDefault()) {
    // Filters

    fun rulesetLabel(ruleset: String): String = when (ruleset) {
        "" -> res.getString(R.string.rank_ruleset_all)
        "normal" -> res.getString(R.string.rank_ruleset_normal)
        "pvp" -> res.getString(R.string.rank_ruleset_pvp)
        "rp" -> res.getString(R.string.rank_ruleset_rp)
        "hardcore" -> res.getString(R.string.rank_ruleset_hardcore)
        else -> HomePresenter.words(ruleset)
    }

    /** A region chip: `All`, `US`. */
    fun regionChip(region: String): String = if (region.isEmpty()) res.getString(R.string.rank_region_all) else region.uppercase(Locale.ROOT)

    fun factionLabel(faction: String): String = when (faction) {
        "alliance" -> res.getString(R.string.rank_faction_alliance)
        "horde" -> res.getString(R.string.rank_faction_horde)
        else -> res.getString(R.string.rank_faction_both)
    }

    /** The API's phase id in words; an id the app does not know is shown as it is (web phases.ts). */
    fun phaseLabel(phase: String): String = when (phase) {
        "pre-beta" -> res.getString(R.string.rank_phase_preBeta)
        "beta" -> res.getString(R.string.rank_phase_beta)
        "launch" -> res.getString(R.string.rank_phase_launch)
        "raids-1" -> res.getString(R.string.rank_phase_raids1)
        else -> phase
    }

    fun measureLabel(measure: Measure): String = res.getString(measure.label)

    /** `Elemental Shaman`, `Shaman`, or empty for every class. */
    fun cohortName(cohort: Cohort?): String {
        cohort ?: return ""
        val className = className(cohort.classSlug)
        return listOf(cohort.specName, className).filter { it.isNotEmpty() }.joinToString(" ")
    }

    fun className(slug: String?): String = slug?.let { GameClass.of(it)?.label?.let(res::getString) ?: HomePresenter.words(it) }.orEmpty()

    /**
     * `{Encounter} · {Ruleset} · {Region} · {Phase} · {Measure}`, then a faction and `Today only`
     * when chosen; every part in words, never cut (the web's phone rule).
     */
    fun filterLine(filters: RankingsFilters, encounterNames: Map<String, String>): String = listOfNotNull(
        encounterNames[filters.encounter] ?: filters.encounter,
        rulesetLabel(filters.ruleset),
        if (filters.region.isEmpty()) res.getString(R.string.rank_line_allRegions) else filters.region.uppercase(Locale.ROOT),
        if (filters.phase.isEmpty()) null else phaseLabel(filters.phase),
        measureLabel(filters.measure),
        if (filters.faction.isEmpty()) null else factionLabel(filters.faction),
        if (filters.todayOnly) res.getString(R.string.rank_filter_today) else null,
    ).filter { it.isNotEmpty() }.joinToString(SEPARATOR)

    /** `Onyxia's Lair · 40 players · 38 ranked kills`. */
    fun encounterMeta(encounter: RankedEncounter): String {
        val kills = res.getQuantityString(R.plurals.rank_kills, encounter.kills, encounter.kills)
        val zone = encounter.zone
        val size = encounter.size
        return if (zone != null && size != null) res.getString(R.string.rank_encounterMeta, zone, size, kills) else kills
    }

    // Answer card

    fun answerUnit(total: Int, cohort: Cohort?, deduped: Boolean): String {
        val spec = cohort?.specName.orEmpty()
        val cls = className(cohort?.classSlug)
        val id = if (deduped) R.string.rank_unit else R.string.rank_unitKills
        return collapse(res.getString(id, total, spec, cls))
    }

    /** `Your best kill on Onyxia, Normal, Phase 1.`, or `{name} · best kill on …` when the row is not [currentName]'s. */
    fun answerMeta(row: RankingRow, filters: RankingsFilters, encounterNames: Map<String, String>, currentName: String?): AnswerMeta {
        val encounter = encounterNames[filters.encounter] ?: filters.encounter
        val ruleset = rulesetLabel(filters.ruleset)
        val phase = if (filters.phase.isEmpty()) "" else phaseLabel(filters.phase)
        if (row.player.name == currentName) return AnswerMeta(res.getString(R.string.rank_meta, encounter, ruleset, phase), null)
        return AnswerMeta(res.getString(R.string.rank_metaNamed, row.player.name, encounter, ruleset, phase), row.player.name)
    }

    fun unranked(filters: RankingsFilters, encounterNames: Map<String, String>): String =
        res.getString(R.string.rank_none, encounterNames[filters.encounter] ?: filters.encounter)

    fun jump(rank: Int): String = res.getString(R.string.rank_jump, Rankings.ordinal(rank))

    /** Home's Rankings row: `14th of 212 Elemental Shaman characters on Onyxia`, or `Not ranked yet.` */
    fun homeLine(page: RankingsPage, filters: RankingsFilters, encounterNames: Map<String, String>): String =
        homeLine(page.characterRank, page.total, filters, encounterNames)

    fun homeLine(placed: CharacterRank?, total: Int, filters: RankingsFilters, encounterNames: Map<String, String>): String {
        placed ?: return res.getString(R.string.home_rankings_none)
        val cohort = filters.cohort
        return collapse(
            res.getString(
                R.string.home_rankings_line, Rankings.ordinal(placed.rank), total, cohort?.specName.orEmpty(),
                className(cohort?.classSlug), encounterNames[filters.encounter] ?: filters.encounter,
            ),
        )
    }

    // Board rows

    fun unit(measure: Measure): String = res.getString(measure.unit)

    /** `Hordebreakers · 40 players · 3:13 · 2 d ago`. */
    fun rowMeta(row: RankingRow, now: Instant): String {
        val time = FightRows.clock(row.durationMs.toLong())
        val age = age(row.foughtAt, now)
        val guild = row.guild?.name
        return if (guild != null) res.getString(R.string.rank_rowMeta, guild, row.size, time, age) else res.getString(R.string.rank_rowMetaNoGuild, row.size, time, age)
    }

    /** `Rank 14 of 212, Lightningbolltz, 398 damage per second, Hordebreakers, your character`. */
    fun rowLabel(row: RankingRow, total: Int, measure: Measure, yours: Boolean): String {
        val value = when (measure) {
            Measure.Execution -> (row.value * PERCENT).roundToInt().toString()
            else -> row.value.roundToInt().toString()
        }
        val base = res.getString(R.string.rank_a11y_row, row.rank, total, row.player.name, value, res.getString(measure.spoken), row.guild?.name.orEmpty())
            .removeSuffix(", ")
        return if (yours) "$base, ${res.getString(R.string.rank_a11y_yours)}" else base
    }

    /** `3 d ago`: the one relative-age phrase. */
    fun age(iso: String, now: Instant): String =
        runCatching { Instant.parse(iso) }.getOrNull()?.let { RelativeAge.of(it, now).phrase(res) }.orEmpty()

    /** `17 Dec`. */
    fun shortDate(at: Instant): String = SHORT_DATE.format(at.atZone(zone))

    companion object {
        const val SEPARATOR = " · "
        private const val PERCENT = 100.0
        private val SHORT_DATE: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)

        /** A sentence with an empty part (no spec) keeps single spaces. */
        fun collapse(text: String): String = text.replace(Regex(" {2,}"), " ").trim()
    }
}
