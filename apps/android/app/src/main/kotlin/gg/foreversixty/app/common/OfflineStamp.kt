package gg.foreversixty.app.common

import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.temporal.ChronoUnit
import java.util.Locale

/**
 * When a saved answer was fetched, in the offline strip's words (spec §5.0): `9:41 pm` today,
 * `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`. Shared by every screen that draws a saved answer.
 */
object OfflineStamp {
    private const val DAYS_PER_WEEK = 7L
    private val TIME: DateTimeFormatter = DateTimeFormatter.ofPattern("h:mm a", Locale.US)
    private val WEEKDAY: DateTimeFormatter = DateTimeFormatter.ofPattern("EEE", Locale.US)
    private val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)

    fun format(at: Instant, now: Instant, zone: ZoneId = ZoneId.systemDefault()): String {
        val date = at.atZone(zone)
        val time = date.format(TIME).replace("AM", "am").replace("PM", "pm")
        val days = ChronoUnit.DAYS.between(date.toLocalDate(), now.atZone(zone).toLocalDate())
        return when {
            days == 0L -> time
            days < DAYS_PER_WEEK -> "${date.format(WEEKDAY)} $time"
            else -> "${date.format(DAY_MONTH)}, $time"
        }
    }
}
