package gg.foreversixty.app.guides.logic

import gg.foreversixty.app.model.ChangelogEntry
import gg.foreversixty.app.model.KeyDate
import java.time.LocalDate

/** A key date still ahead or in progress, with its gold `running` line while it runs. */
data class KeyDateRow(val date: KeyDate, val running: String?)

/** F9.5's Key dates (the dates contract, app-v1.md): past dates omitted, sorted by `iso`. */
object KeyDates {
    fun upcoming(dates: List<KeyDate>, today: LocalDate): List<KeyDateRow> = dates
        .mapNotNull { date ->
            val start = GuideDates.parse(date.iso) ?: return@mapNotNull null
            val end = date.endIso?.let(GuideDates::parse) ?: start
            if (end.isBefore(today)) return@mapNotNull null
            val runs = !start.isAfter(today)
            KeyDateRow(date, date.running?.takeIf { runs })
        }
        .sortedBy { it.date.iso }
}

enum class ChangeTag { Site, Game }

/** F9.5's changelog rows: the file's order (newest first), tagged Site or Game. */
object Changes {
    private const val SITE_KIND = "site"

    fun tag(entry: ChangelogEntry): ChangeTag = if (entry.kind == SITE_KIND) ChangeTag.Site else ChangeTag.Game
}
