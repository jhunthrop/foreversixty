package gg.foreversixty.app.rankings.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.model.RankedEncounter
import gg.foreversixty.app.model.RankingRow
import gg.foreversixty.app.phrase
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.reports.logic.FightRows
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale
import kotlin.math.roundToInt

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

    /** `Onyxia's Lair · 40 players · 38 ranked kills`. */
    fun encounterMeta(encounter: RankedEncounter): String {
        val kills = res.getQuantityString(R.plurals.rank_kills, encounter.kills, encounter.kills)
        val zone = encounter.zone
        val size = encounter.size
        return if (zone != null && size != null) res.getString(R.string.rank_encounterMeta, zone, size, kills) else kills
    }

    // The page (design/mocks/web web-rankings*)

    /** The head's line: where you stand signed in, how the board ranks signed out. */
    fun line(signedIn: Boolean): String = res.getString(if (signedIn) R.string.rank_flows_lineIn else R.string.rank_flows_lineOut)

    /** `Fury Warriors`, or `Every class and spec`. */
    fun cohortTitle(cohort: Cohort?): String =
        cohortName(cohort).ifEmpty { null }?.let { res.getString(R.string.rank_flows_cohortPlural, it) } ?: res.getString(R.string.rank_flows_everyClass)

    /** `Normal ruleset`, or `All rulesets`. */
    fun rulesetScope(ruleset: String): String =
        if (ruleset.isEmpty()) res.getString(R.string.rank_flows_allRulesets) else res.getString(R.string.rank_flows_rulesetScope, rulesetLabel(ruleset))

    /** The head's facts: `Rankings · Fury Warriors · Normal ruleset`, then `Updated 13 Dec`. */
    fun facts(filters: RankingsFilters, updatedAt: String): List<String> = listOfNotNull(
        res.getString(R.string.rank_flows_facts, cohortTitle(filters.cohort), rulesetScope(filters.ruleset)),
        runCatching { Instant.parse(updatedAt) }.getOrNull()?.let { res.getString(R.string.rank_flows_updated, shortDate(it)) },
    )

    /** `of 212 Fury Warriors` once each row is a character's best kill, else `of 212 Fury Warrior kills`. */
    fun answerUnit(total: Int, cohort: Cohort?, deduped: Boolean): String {
        val name = cohortName(cohort).ifEmpty { res.getString(R.string.rank_flows_everyClassField) }
        return res.getString(if (deduped) R.string.rank_flows_answerUnit else R.string.rank_flows_answerUnitKills, total, name)
    }

    /** `On Onyxia, Normal ruleset`. */
    fun answerNote(filters: RankingsFilters, encounterNames: Map<String, String>): String =
        res.getString(R.string.rank_flows_answerNote, encounterName(filters, encounterNames), rulesetScope(filters.ruleset))

    fun encounterName(filters: RankingsFilters, encounterNames: Map<String, String>): String = encounterNames[filters.encounter] ?: filters.encounter

    fun unranked(filters: RankingsFilters, encounterNames: Map<String, String>): String =
        res.getString(R.string.rank_none, encounterName(filters, encounterNames))

    /** `212 ranked, each character's best kill`, or `212 ranked kills`. */
    fun sectionMeta(total: Int, deduped: Boolean): String =
        res.getString(if (deduped) R.string.rank_flows_sectionMeta else R.string.rank_flows_sectionMetaKills, total)

    // Board rows

    fun unit(measure: Measure): String = res.getString(measure.unit)

    /** `Olympus XXVII · Fury Warrior · 40-man`: the guild when filed to one, the spec and class, the raid size. */
    fun rowSub(row: RankingRow): String {
        val who = listOfNotNull(row.player.spec, className(row.player.className?.lowercase(Locale.ROOT))).filter { it.isNotEmpty() }.joinToString(" ")
        return listOfNotNull(row.guild?.name, who.ifEmpty { null }, res.getString(R.string.rank_flows_size, row.size)).joinToString(SEPARATOR)
    }

    /** `90%` executed, when the kill has an execution score. */
    fun executed(row: RankingRow): String? = row.executionScore?.let { "${(it * PERCENT).roundToInt()}$PERCENT_SIGN" }

    /** The kill's length, `2:41`. */
    fun length(row: RankingRow): String = FightRows.clock(row.durationMs.toLong())

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
        private const val PERCENT_SIGN = "%"
        private val SHORT_DATE: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)

        /** A sentence with an empty part (no spec) keeps single spaces. */
        fun collapse(text: String): String = text.replace(Regex(" {2,}"), " ").trim()
    }
}
