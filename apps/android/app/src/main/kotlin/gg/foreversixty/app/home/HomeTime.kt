package gg.foreversixty.app.home

import android.content.res.Resources
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.phrase
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.temporal.ChronoUnit
import java.util.Locale

/**
 * The two time phrases Home prints: the sync age (`2 h ago`) and the offline stamp (`9:41 pm`
 * today, `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`; spec §5.0).
 */
class HomeTime(private val now: Instant, private val zone: ZoneId = ZoneId.systemDefault()) {
    fun age(since: Instant, res: Resources): String = RelativeAge.of(since, now).phrase(res)

    fun stamp(at: Instant): String {
        val date = at.atZone(zone)
        val today = now.atZone(zone).toLocalDate()
        val time = date.format(TIME).replace("AM", "am").replace("PM", "pm")
        val days = ChronoUnit.DAYS.between(date.toLocalDate(), today)
        return when {
            days == 0L -> time
            days < DAYS_PER_WEEK -> "${date.format(WEEKDAY)} $time"
            else -> "${date.format(DAY_MONTH)}, $time"
        }
    }

    private companion object {
        const val DAYS_PER_WEEK = 7L
        val TIME: DateTimeFormatter = DateTimeFormatter.ofPattern("h:mm a", Locale.US)
        val WEEKDAY: DateTimeFormatter = DateTimeFormatter.ofPattern("EEE", Locale.US)
        val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
    }
}

/** The API's RFC 3339 timestamps; null when absent or unreadable. */
fun parseInstant(text: String?): Instant? = text?.let { runCatching { Instant.parse(it) }.getOrNull() }
