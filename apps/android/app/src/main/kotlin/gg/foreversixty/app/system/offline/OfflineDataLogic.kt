package gg.foreversixty.app.system.offline

import gg.foreversixty.app.network.SavedStaticFile
import gg.foreversixty.app.network.StaticDataClient
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.temporal.ChronoUnit
import java.util.Locale

/** The rows of You › Offline data (spec F15.4), in the board's order. Item icons come from the image cache. */
enum class OfflineCategory { Guides, Bis, Talents, Search, Icons }

/** One row: the category, the classes it holds (class slugs, sorted), its bytes and its newest save. */
data class OfflineRow(val category: OfflineCategory, val classes: List<String>, val bytes: Long, val updated: Instant?)

/** What the screen draws: the rows that hold something, and their total. Empty means the empty card. */
data class OfflineReport(val rows: List<OfflineRow>) {
    val totalBytes: Long get() = rows.sumOf { it.bytes }
    val isEmpty: Boolean get() = rows.isEmpty()
}

/** The size of the image cache: its bytes and newest file. */
data class IconCacheSize(val bytes: Long, val updated: Instant?)

/** Pure rules for Offline data: which file is which row, the classes, sizes and times (tenet 7: no ids, hashes or counts). */
object OfflineData {
    private const val KB = 1_000.0
    private const val MB = 1_000_000.0
    private const val WHOLE_MB_FROM = 10.0
    private const val DAYS_PER_WEEK = 7L

    /** The row a saved `/data/...` file belongs to; null for files no row shows (the active build pointer). */
    fun category(path: String): OfflineCategory? {
        val rest = path.removePrefix(StaticDataClient.DATA_PREFIX)
        val parts = rest.split('/')
        if (parts.size == 1) {
            return when (parts.single()) {
                "changelog.json", "dates.json" -> OfflineCategory.Guides
                else -> null
            }
        }
        val inBuild = parts.drop(1)
        return when (inBuild.first()) {
            "guides", "tiers.json" -> OfflineCategory.Guides
            "talents" -> OfflineCategory.Talents
            "search.json" -> OfflineCategory.Search
            "bis", "items", "item-uses", "loot.json", "sets.json", "classes.json" -> OfflineCategory.Bis
            else -> null
        }
    }

    /**
     * The class a file is for: `guides/shaman-elemental.json` and `bis/shaman-elemental.json` name a
     * spec (class first), `talents/shaman.json` a class. Indexes, item files and the rest name none.
     */
    fun classOf(path: String): String? {
        val parts = path.removePrefix(StaticDataClient.DATA_PREFIX).split('/')
        if (parts.size != 3) return null
        val stem = parts[2].removeSuffix(".json")
        return when (parts[1]) {
            "guides" -> stem.takeUnless { it == "index" }?.substringBefore('-')
            "bis" -> stem.substringBefore('-')
            "talents" -> stem
            else -> null
        }
    }

    /** The report from the store's files and the icon cache, in category order, empty rows left out. */
    fun report(files: List<SavedStaticFile>, icons: IconCacheSize): OfflineReport {
        val grouped = files.mapNotNull { file -> category(file.path)?.let { it to file } }.groupBy({ it.first }, { it.second })
        val rows = OfflineCategory.entries.mapNotNull { category ->
            if (category == OfflineCategory.Icons) {
                return@mapNotNull icons.takeIf { it.bytes > 0 }?.let { OfflineRow(category, emptyList(), it.bytes, it.updated) }
            }
            val inRow = grouped[category].orEmpty().takeIf { it.isNotEmpty() } ?: return@mapNotNull null
            OfflineRow(
                category = category,
                classes = inRow.mapNotNull { classOf(it.path) }.distinct().sorted(),
                bytes = inRow.sumOf { it.bytes },
                updated = Instant.ofEpochMilli(inRow.maxOf { it.savedAtMillis }),
            )
        }
        return OfflineReport(rows)
    }

    /** `410 KB`, `2.4 MB` (one decimal under 10 MB), `41 MB` (whole above). Decimal units, as iOS. */
    fun size(bytes: Long): String = when {
        bytes < MB -> "${Math.round(bytes / KB).coerceAtLeast(if (bytes > 0) 1 else 0)} KB"
        bytes / MB < WHOLE_MB_FROM -> String.format(Locale.US, "%.1f MB", bytes / MB)
        else -> "${Math.round(bytes / MB)} MB"
    }

    /** The row's `Updated {when}`: `today`, `yesterday`, a weekday this week, else `10 Dec`. */
    fun updated(at: Instant, now: Instant, zone: ZoneId, today: String, yesterday: String): String {
        val days = ChronoUnit.DAYS.between(at.atZone(zone).toLocalDate(), now.atZone(zone).toLocalDate())
        return when {
            days <= 0L -> today
            days == 1L -> yesterday
            days < DAYS_PER_WEEK -> at.atZone(zone).format(WEEKDAY)
            else -> at.atZone(zone).format(DAY_MONTH)
        }
    }

    /** The footer's time: `9:41 pm`. */
    fun time(at: Instant, zone: ZoneId): String =
        at.atZone(zone).format(TIME).replace("AM", "am").replace("PM", "pm")

    private val TIME: DateTimeFormatter = DateTimeFormatter.ofPattern("h:mm a", Locale.US)
    private val WEEKDAY: DateTimeFormatter = DateTimeFormatter.ofPattern("EEE", Locale.US)
    private val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
}
