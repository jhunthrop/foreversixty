package gg.foreversixty.app.reports.logic

import gg.foreversixty.app.model.FightEntry
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.MeterPlayer
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.model.ReportMe
import gg.foreversixty.app.model.ReportRow
import gg.foreversixty.app.model.TrashGroup
import java.time.Duration
import java.time.Instant
import kotlin.math.roundToInt
import kotlin.time.Duration.Companion.seconds

/** A report row's tag (spec F10.1): live, processing, failed, or none for a finished report. */
enum class ReportState {
    Live, Processing, Failed, Done;

    companion object {
        fun of(status: String): ReportState = when (status) {
            "live" -> Live
            "processing" -> Processing
            "failed" -> Failed
            else -> Done
        }
    }
}

/** Who reads the reports: the session's account, or nobody (a deep-linked public report). */
data class ReportsViewer(
    val signedIn: Boolean,
    /** `entitlements.log_coach`: the coach card's one gate. */
    val logCoach: Boolean,
    /** The account's characters, key to class slug, for the logging character's crest. */
    val classes: Map<String, String>,
    /** The account's current character's class, when the logging character is not one of them. */
    val fallbackClass: String?,
) {
    companion object {
        val SignedOut = ReportsViewer(signedIn = false, logCoach = false, classes = emptyMap(), fallbackClass = null)

        /** The signed-in account ([me]), with [currentKey] its current character. */
        fun of(me: Me, currentKey: String?): ReportsViewer {
            val classes = me.characters.mapNotNull { c -> c.className?.let { c.key to it } }.toMap()
            return ReportsViewer(true, me.entitlements.logCoach, classes, currentKey?.let { classes[it] })
        }
    }
}

/** The list's two sections (spec F10.1), each newest first as the server sends them. */
data class ReportSections(val thisWeek: List<ReportRow>, val earlier: List<ReportRow>)

/** The rules under Your reports (spec F10): pure data in, data out. Every sentence is in `ReportsCopy`. */
object Reports {
    /** `This week` is the last seven days of now. */
    val WEEK: Duration = Duration.ofDays(7)

    /** A live report re-reads `GET /v1/reports/{id}` this often while its screen is open. */
    val LIVE_POLL = 20.seconds

    fun sections(rows: List<ReportRow>, now: Instant): ReportSections {
        val since = now.minus(WEEK)
        val (week, earlier) = rows.partition { row -> instantOrNull(row.createdAt)?.let { !it.isBefore(since) } ?: false }
        return ReportSections(week, earlier)
    }

    /** The logging character's class: theirs when the account holds them, else the current character's, else none. */
    fun crestClass(row: ReportRow, viewer: ReportsViewer): String? =
        row.loggingCharacter?.key?.let { viewer.classes[it] } ?: viewer.fallbackClass

    /**
     * Home's `Your last raid`: the newest finished report with a boss pull (never a live dungeon
     * run); null when there is none.
     */
    fun lastRaid(rows: List<ReportRow>): ReportRow? = rows.firstOrNull { ReportState.of(it.status) == ReportState.Done && it.bossCount > 0 }

    /** `title` when the report has one, the zone otherwise. */
    fun name(row: ReportRow): String = row.title.ifEmpty { row.zone }

    /** Whole seconds since the companion last spoke (heartbeat, else the last upload), never negative. */
    fun secondsSinceUpdate(heartbeatAt: String?, lastUploadAt: String?, now: Instant): Long? {
        val at = (heartbeatAt ?: lastUploadAt)?.let(::instantOrNull) ?: return null
        return maxOf(0, now.epochSecond - at.epochSecond)
    }

    /** A live report polls; nothing else does. */
    fun polls(report: Report): Boolean = ReportState.of(report.status) == ReportState.Live

    fun instantOrNull(text: String): Instant? = runCatching { Instant.parse(text) }.getOrNull()
}

/** `pull 1 of 3`: a boss pull's place among that boss's pulls in the report (the web's pullNumbers). */
data class PullNumber(val n: Int, val of: Int)

/** One row of the Fights card (spec F10.2). */
sealed interface FightRow {
    /** A boss pull: its own row, in fight order; [pull] numbers a boss pulled more than once. */
    data class Boss(val fight: FightEntry, val pull: PullNumber? = null) : FightRow {
        val inProgress: Boolean get() = fight.inProgress == true

        /** The boss's health at a wipe, rounded; null on a kill, or when the log never showed the boss. */
        val wipePercent: Int? get() = if (fight.kill) null else fight.bossHealthPct?.takeIf { it >= 0 }?.roundToInt()
    }

    /** A zone's trash as one row, after that zone's boss pulls; the server's `trash_groups`. */
    data class Trash(val group: TrashGroup) : FightRow {
        val pulls: Int get() = group.fights.size
    }

    /** A trash pull with no zone keeps its own row. */
    data class LoosePull(val fight: FightEntry) : FightRow
}

object FightRows {
    private const val ENCOUNTER = "encounter"

    fun isBoss(fight: FightEntry): Boolean = fight.kind == ENCOUNTER

    /** `{n} boss pulls` counts boss pulls only; trash is never counted. */
    fun bossPulls(report: Report): Int = report.fights.count(::isBoss)

    /**
     * Boss pulls in fight order; each trash group after its zone's last boss pull (where its first
     * pull fell when the zone has none); ungrouped trash in place. Groups are never computed here.
     */
    fun of(report: Report): List<FightRow> {
        val groups = report.trashGroups.orEmpty()
        val grouped = groups.flatMap { it.fights }.toSet()
        val rows = mutableListOf<FightRow>()
        val anchors = mutableMapOf<Int, Int>()
        val pulls = pullNumbers(report)
        for (fight in report.fights) {
            when {
                isBoss(fight) -> rows += FightRow.Boss(fight, pulls[fight.index])
                fight.index in grouped -> groups.firstOrNull { it.key == fight.index }?.let { anchors[it.key] = rows.size }
                else -> rows += FightRow.LoosePull(fight)
            }
        }
        // Insert from the last position back (the night's later group first on a tie), so earlier positions stay valid.
        val placed = groups.mapIndexed { order, group -> Triple(group, position(rows, group, anchors[group.key] ?: rows.size), order) }
        placed.sortedWith(compareByDescending<Triple<TrashGroup, Int, Int>> { it.second }.thenByDescending { it.third })
            .forEach { (group, at, _) -> rows.add(at, FightRow.Trash(group)) }
        return rows
    }

    /** Each boss pull's number among the same boss's pulls, in fight order; only bosses pulled more than once. */
    fun pullNumbers(report: Report): Map<Int, PullNumber> {
        val bosses = report.fights.filter(::isBoss).groupBy { it.encounterId ?: it.name.hashCode().toLong() }
        return bosses.values.filter { it.size > 1 }.flatMap { pulls ->
            pulls.mapIndexed { i, fight -> fight.index to PullNumber(i + 1, pulls.size) }
        }.toMap()
    }

    private fun position(rows: List<FightRow>, group: TrashGroup, anchor: Int): Int {
        val lastBoss = rows.indexOfLast { it is FightRow.Boss && it.fight.zone == group.zone }
        return if (lastBoss >= 0) lastBoss + 1 else anchor
    }

    /** `m:ss` from milliseconds, as the web's fight rows print it. */
    fun clock(ms: Long): String {
        val total = (ms / MS_PER_SECOND).toInt()
        return "%d:%02d".format(total / SECONDS_PER_MINUTE, total % SECONDS_PER_MINUTE)
    }

    private const val MS_PER_SECOND = 1000L
    private const val SECONDS_PER_MINUTE = 60
}

/** The parse ladder (spec §4 Parse number), each tier from its floor up. */
enum class ParseTier(val floor: Int) {
    Grey(0), Green(25), Blue(50), Purple(75), Orange(95), Pink(99);

    companion object {
        fun of(percentile: Int): ParseTier = entries.last { percentile >= it.floor }
    }
}

/** One character fight in `Your parses`: one 72 dp row. */
data class ParseRow(
    val characterKey: String,
    val name: String,
    val className: String,
    val specName: String,
    val fight: gg.foreversixty.app.model.ReportMeFight,
)

object Parses {
    /** Every character fight in the server's order; empty means `No kills of yours in this report.` */
    fun rows(me: ReportMe): List<ParseRow> = me.characters.flatMap { c ->
        c.fights.map { ParseRow(c.key, c.name, c.className, c.specName, it) }
    }

    /** The figure's integer, as `712` from 712.4. */
    fun figure(value: Double): String = value.roundToInt().toString()

    /** The first character with fights: the coach card reads their night. */
    fun coachPlayer(me: ReportMe): String? = me.characters.firstOrNull { it.fights.isNotEmpty() }?.key
}

/** The fight sheet's three segments (spec F10.3) and what each one reads. */
enum class Metric(val query: String, val unit: String) {
    Damage("damage", "dps"), Healing("healing", "hps"), Taken("taken", "dtps");

    fun value(player: MeterPlayer): Double = when (this) {
        Damage -> player.dps
        Healing -> player.hps
        Taken -> player.dtps
    }
}

/** What the fight sheet opens on: a boss pull, or a zone's trash. */
sealed interface MetersTarget {
    val index: Int
    val scope: String?

    data class Pull(val fight: FightEntry, val pull: PullNumber? = null) : MetersTarget {
        override val index: Int get() = fight.index
        override val scope: String? get() = null
    }

    data class Trash(val group: TrashGroup) : MetersTarget {
        override val index: Int get() = group.key
        override val scope: String? get() = TRASH_SCOPE
    }

    private companion object {
        const val TRASH_SCOPE = "trash"
    }
}

/** One meter row: its rank, the player, the figure and the bar's length against the top row. */
data class MeterRow(val rank: Int, val player: MeterPlayer, val value: Double, val fraction: Float)

object Meters {
    /** Rows before `Show all {n}` (spec §9): 11 on a phone, 4 at the largest text. */
    const val ROWS = 11
    const val ROWS_LARGE = 4

    /** From this font scale the meter rows put the value under the name and show four. */
    const val LARGE_SCALE = 1.5f

    /** The loading state draws this many skeleton rows. */
    const val SKELETON_ROWS = 8

    fun cut(fontScale: Float): Int = if (fontScale >= LARGE_SCALE) ROWS_LARGE else ROWS

    /** The server's order, each bar the row's figure over the top figure. */
    fun rows(players: List<MeterPlayer>, metric: Metric): List<MeterRow> {
        val top = players.maxOfOrNull(metric::value)?.takeIf { it > 0 }
        return players.mapIndexed { i, p ->
            val value = metric.value(p)
            MeterRow(i + 1, p, value, if (top == null) 0f else (value / top).toFloat().coerceIn(0f, 1f))
        }
    }

    /** The first [cut] rows, and your rows after them when they fall below the cut; every row after Show all. */
    fun visible(rows: List<MeterRow>, cut: Int, showAll: Boolean): List<MeterRow> {
        if (showAll) return rows
        val top = rows.take(cut)
        return top + rows.drop(cut).filter { it.player.mine }
    }

    /** `Show all {n}` shows while rows are hidden. */
    fun hasMore(rows: List<MeterRow>, cut: Int, showAll: Boolean): Boolean = !showAll && visible(rows, cut, false).size < rows.size
}
