package gg.foreversixty.app.reports.logic

import android.content.res.Resources
import android.icu.text.MeasureFormat
import android.icu.text.MessageFormat
import android.icu.util.Measure
import android.icu.util.MeasureUnit
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.model.FightEntry
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.model.ReportRow
import gg.foreversixty.app.raids.logic.GameClass
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** A report's visibility: its glyph's word (spec F10.1). */
enum class Visibility(@param:StringRes val word: Int) {
    Public(R.string.reports_visibility_public),
    Unlisted(R.string.reports_visibility_unlisted),
    Guild(R.string.reports_visibility_guild),
    Private(R.string.reports_visibility_private);

    companion object {
        fun of(value: String): Visibility = when (value) {
            "public" -> Public
            "unlisted" -> Unlisted
            "guild" -> Guild
            else -> Private
        }
    }
}

/** Every F10 sentence built from data (spec §5.2); the words themselves live in the copy catalog. */
class ReportsCopy(private val res: Resources, private val zone: ZoneId = ZoneId.systemDefault()) {
    // List (F10.1)

    /** `15 Dec · 1 boss · 1 kill`; a live row only `Updated 20 s ago`; a failed row `29 Nov · could not be read`. */
    fun rowMeta(row: ReportRow, now: Instant): String {
        val date = Reports.instantOrNull(row.createdAt)?.let(::shortDate).orEmpty()
        return when (ReportState.of(row.status)) {
            ReportState.Failed -> res.getString(R.string.reports_unreadable, date)
            ReportState.Live -> updated(row, now) ?: res.getString(R.string.reports_rowMeta, date, bosses(row.bossCount), kills(row.killCount))
            ReportState.Processing, ReportState.Done -> res.getString(R.string.reports_rowMeta, date, bosses(row.bossCount), kills(row.killCount))
        }
    }

    fun bosses(n: Int): String = res.getQuantityString(R.plurals.reports_bosses, n, n)

    fun kills(n: Int): String = res.getQuantityString(R.plurals.reports_kills, n, n)

    private fun updated(row: ReportRow, now: Instant): String? =
        Reports.secondsSinceUpdate(row.heartbeatAt, row.lastUploadAt, now)?.let { res.getString(R.string.reports_updated, it.toInt()) }

    fun stateTag(state: ReportState): String? = when (state) {
        ReportState.Live -> res.getString(R.string.reports_state_live)
        ReportState.Processing -> res.getString(R.string.reports_state_processing)
        ReportState.Failed -> res.getString(R.string.reports_state_failed)
        ReportState.Done -> null
    }

    fun visibility(value: String): String = res.getString(Visibility.of(value).word)

    /** `Onyxia's Lair, Monday 14 December, 1 boss, 1 kill, Guild` and, live, `live, updated 20 seconds ago`. */
    fun rowLabel(row: ReportRow, now: Instant): String {
        val date = Reports.instantOrNull(row.createdAt)?.let { "${WEEKDAY.format(it.atZone(zone))} ${DAY_MONTH.format(it.atZone(zone))}" }.orEmpty()
        val base = res.getString(R.string.reports_a11y_row, Reports.name(row), date, bosses(row.bossCount), kills(row.killCount), visibility(row.visibility))
        val seconds = Reports.secondsSinceUpdate(row.heartbeatAt, row.lastUploadAt, now)
        return if (ReportState.of(row.status) == ReportState.Live && seconds != null) {
            "$base, ${res.getString(R.string.reports_a11y_live, seconds.toString())}"
        } else {
            base
        }
    }

    // Report (F10.2)

    /** `Monday 14 December 2026 · logged by Obnoxious Yell`. */
    fun reportMeta(report: Report): String {
        val at = Reports.instantOrNull(report.createdAt)?.atZone(zone)
        val weekday = at?.let(WEEKDAY::format).orEmpty()
        val date = at?.let(LONG_DATE::format).orEmpty()
        return res.getString(R.string.report_meta, weekday, date, report.loggingCharacter?.name.orEmpty())
    }

    fun bossPulls(n: Int): String = if (n == 0) res.getString(R.string.report_noBossPulls) else res.getQuantityString(R.plurals.report_bossPulls, n, n)

    fun players(n: Int): String = res.getQuantityString(R.plurals.report_players, n, n)

    fun deaths(n: Int): String = res.getQuantityString(R.plurals.report_deaths, n, n)

    fun pulls(n: Int): String = res.getQuantityString(R.plurals.report_pulls, n, n)

    /** `Onyxia · pull 1 of 3` for a boss pulled more than once, else the boss's name. */
    fun bossTitle(fight: FightEntry, pull: PullNumber?): String =
        pull?.let { res.getString(R.string.report_pullOf, fight.name, it.n, it.of) } ?: fight.name

    fun bossMeta(row: FightRow.Boss): String {
        val fight = row.fight
        val time = FightRows.clock(fight.durationMs)
        val wipe = row.wipePercent
        return when {
            row.inProgress -> res.getString(R.string.report_inProgress)
            fight.kill -> res.getString(R.string.report_killMeta, time, players(fight.players.size), deaths(fight.deaths ?: 0))
            wipe != null -> res.getString(R.string.report_wipeMeta, time, "$wipe%")
            else -> res.getString(R.string.report_wipeMetaPlain, time)
        }
    }

    fun trashTitle(zone: String): String = res.getString(R.string.report_trashTitle, zone)

    fun trashMeta(row: FightRow.Trash): String = res.getString(R.string.report_trashMeta, pulls(row.pulls), FightRows.clock(row.group.durationMs))

    /** A trash pull with no zone: `1 pull · 1:35`. */
    fun looseMeta(row: FightRow.LoosePull): String = res.getString(R.string.report_trashMeta, pulls(1), FightRows.clock(row.fight.durationMs))

    fun looseTitle(): String = res.getString(R.string.report_trash)

    fun parseMeta(row: ParseRow): String = res.getString(R.string.report_parseMeta, FightRows.clock(row.fight.durationMs), row.name, row.specName)

    fun unit(unit: String): String = res.getString(
        when (unit) {
            "hps" -> R.string.report_unit_hps
            "dtps" -> R.string.report_unit_dtps
            else -> R.string.report_unit_dps
        },
    )

    private fun perSecond(unit: String): String = res.getString(
        when (unit) {
            "hps" -> R.string.report_a11y_perSecond_hps
            "dtps" -> R.string.report_a11y_perSecond_dtps
            else -> R.string.report_a11y_perSecond_dps
        },
    )

    fun bossLabel(row: FightRow.Boss): String {
        val fight = row.fight
        val spoken = spokenDuration(fight.durationMs)
        val wipe = row.wipePercent
        return when {
            fight.kill -> res.getString(R.string.report_a11y_kill, bossTitle(fight, row.pull), spoken, players(fight.players.size), deaths(fight.deaths ?: 0))
            wipe != null -> res.getString(R.string.report_a11y_wipe, bossTitle(fight, row.pull), spoken, wipe)
            else -> res.getString(R.string.report_a11y_wipePlain, bossTitle(fight, row.pull), spoken)
        }
    }

    fun trashLabel(row: FightRow.Trash): String = res.getString(R.string.report_a11y_trash, row.group.zone, pulls(row.pulls), spokenDuration(row.group.durationMs))

    fun parseLabel(row: ParseRow): String {
        val fight = row.fight
        val spoken = spokenDuration(fight.durationMs)
        val figure = Parses.figure(fight.value)
        val percentile = fight.percentile
        return if (percentile == null) {
            res.getString(R.string.report_a11y_parseNoRank, fight.encounter, spoken, figure, perSecond(fight.unit))
        } else {
            res.getString(R.string.report_a11y_parse, fight.encounter, spoken, figure, perSecond(fight.unit), ordinal(percentile))
        }
    }

    // Meters (F10.3)

    fun metersTitle(metric: Metric): String = res.getString(
        when (metric) {
            Metric.Damage -> R.string.meters_title_damage
            Metric.Healing -> R.string.meters_title_healing
            Metric.Taken -> R.string.meters_title_taken
        },
    )

    fun segment(metric: Metric): String = res.getString(
        when (metric) {
            Metric.Damage -> R.string.meters_damage
            Metric.Healing -> R.string.meters_healing
            Metric.Taken -> R.string.meters_taken
        },
    )

    /**
     * `Kill · 4:28 · 40 players`, `Wipe · 1:12 · 40 players`, or for trash `6 pulls · 7:40`. The
     * players are the pull's roster, never the meter's rows (a segment leaves off players with none of it).
     */
    fun metersHead(target: MetersTarget): String = when (target) {
        is MetersTarget.Trash -> res.getString(R.string.meters_trashHead, pulls(target.group.fights.size), FightRows.clock(target.group.durationMs))
        is MetersTarget.Pull -> {
            val time = FightRows.clock(target.fight.durationMs)
            val players = players(target.fight.players.size)
            res.getString(if (target.fight.kill) R.string.meters_killHead else R.string.meters_wipeHead, time, players)
        }
    }

    fun meterLabel(row: MeterRow): String {
        val p = row.player
        val who = listOf(p.specName, GameClass.of(p.className)?.let { res.getString(it.label) }.orEmpty()).filter { it.isNotEmpty() }.joinToString(" ")
        val label = res.getString(R.string.meters_a11y_row, row.rank, p.name, who, Parses.figure(row.value))
        return if (p.mine) res.getString(R.string.meters_a11y_mine, label) else label
    }

    fun metersTitleFor(target: MetersTarget, metric: Metric): String = when (target) {
        is MetersTarget.Trash -> trashTitle(target.group.zone)
        is MetersTarget.Pull -> metersTitle(metric)
    }

    fun headName(target: MetersTarget): String? = (target as? MetersTarget.Pull)?.let { bossTitle(it.fight, it.pull) }

    // Coach (F10.4)

    /** `Reviewed 15 Dec · Based on: a; b.`: each label once, its own full stop dropped. */
    fun coachFooter(reviewedAt: String, labels: List<String>): String {
        val date = Reports.instantOrNull(reviewedAt)?.let(::shortDate).orEmpty()
        return res.getString(R.string.coach_footer, date, labels.map { it.trim().trimEnd('.') }.distinct().joinToString("; "))
    }

    // Shared

    fun shortDate(at: Instant): String = SHORT_DATE.format(at.atZone(zone))

    /** `3 minutes, 12 seconds`: the platform's own duration words. */
    fun spokenDuration(ms: Long): String {
        val total = (ms / MS_PER_SECOND).toInt()
        val minutes = total / SECONDS_PER_MINUTE
        val seconds = total % SECONDS_PER_MINUTE
        val parts = buildList {
            if (minutes > 0) add(Measure(minutes, MeasureUnit.MINUTE))
            if (seconds > 0 || minutes == 0) add(Measure(seconds, MeasureUnit.SECOND))
        }
        return MeasureFormat.getInstance(locale(), MeasureFormat.FormatWidth.WIDE).formatMeasures(*parts.toTypedArray())
    }

    /** `91st`: the platform's ordinal. */
    fun ordinal(n: Int): String = MessageFormat("{0,ordinal}", locale()).format(arrayOf<Any>(n))

    private fun locale(): Locale = res.configuration.locales[0] ?: Locale.US

    private companion object {
        const val MS_PER_SECOND = 1000L
        const val SECONDS_PER_MINUTE = 60
        val SHORT_DATE: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
        val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMMM", Locale.US)
        val LONG_DATE: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMMM yyyy", Locale.US)
        val WEEKDAY: DateTimeFormatter = DateTimeFormatter.ofPattern("EEEE", Locale.US)
    }
}
