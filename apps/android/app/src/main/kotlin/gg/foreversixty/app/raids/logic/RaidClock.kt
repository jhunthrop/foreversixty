package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.temporal.ChronoUnit
import java.util.Locale

/**
 * The raid's times as the viewer reads them (spec 5.4 F4.1 and F4.2; the site's raid-copy.ts):
 * clocks, the date block, the long date, the planned row's `{when}`, the countdown and the
 * guild-time line. Everything takes the zone and the instant, so a test pins both.
 */
class RaidClock(
    private val res: Resources,
    val zone: ZoneId = ZoneId.systemDefault(),
    private val locale: Locale = Locale.UK,
) {
    private fun format(pattern: String, at: Instant, zone: ZoneId = this.zone): String {
        val formatter = DateTimeFormatter.ofPattern(pattern, if ('z' in pattern) Locale.US else locale)
        return formatter.format(at.atZone(zone)).replace("AM", "am").replace("PM", "pm")
    }

    /** `7:45 pm`. */
    fun clock(at: Instant): String = format("h:mm a", at)

    /** `8:00 pm CST` in the guild's zone. */
    fun guildClock(at: Instant, guildZone: ZoneId): String = format("h:mm a zzz", at, guildZone)

    /** The date block: `THU`, `10`, `DEC`. */
    fun dateBlock(at: Instant): Triple<String, String, String> =
        Triple(format("EEE", at).uppercase(locale), format("d", at), format("MMM", at).uppercase(locale))

    /** `Thursday 10 December`. */
    fun longDate(at: Instant): String = format("EEEE d MMMM", at)

    /** `Thu 10 Dec`. */
    fun shortDate(at: Instant): String = format("EEE d MMM", at)

    fun weekday(at: Instant): String = format("EEEE", at)
    fun dayMonth(at: Instant): String = format("d MMMM", at)

    /** The planned row's `{when}`: `Tonight 8:00 pm`, `Tomorrow 8:00 pm`, `Saturday 7:00 pm`, `Thu 17 Dec, 8:00 pm`. */
    fun whenPhrase(at: Instant, now: Instant): String {
        val time = clock(at)
        return when (calendarDays(now, at)) {
            0L -> res.getString(R.string.raids_when_tonight, time)
            1L -> res.getString(R.string.raids_when_tomorrow, time)
            in 2L..WEEKDAY_WINDOW_DAYS -> res.getString(R.string.raids_when_weekday, weekday(at), time)
            else -> res.getString(R.string.raids_when_date, shortDate(at), time)
        }
    }

    /** True when the viewer's zone and the guild's differ at that instant. */
    fun offsetDiffers(at: Instant, guildZone: ZoneId): Boolean =
        zone.rules.getOffset(at) != guildZone.rules.getOffset(at)

    private fun calendarDays(from: Instant, to: Instant): Long =
        ChronoUnit.DAYS.between(from.atZone(zone).toLocalDate(), to.atZone(zone).toLocalDate())

    /** The site's countdown: `in 2 days 2 h`, `starting now`, `started 20 min ago`, `ended 1 h ago`. */
    fun countdown(start: Instant, end: Instant, now: Instant): String {
        val untilStart = Duration.between(now, start)
        return when {
            untilStart.abs() < Duration.ofMinutes(1) -> res.getString(R.string.raid_countdown_now)
            !untilStart.isNegative -> res.getString(R.string.raid_countdown_in, span(untilStart))
            now < end -> res.getString(R.string.raid_countdown_started, span(untilStart.negated()))
            else -> res.getString(R.string.raid_countdown_ended, span(Duration.between(end, now)))
        }
    }

    /** `2 days 2 h`, `5 h 20 min`, `12 min`. */
    fun span(duration: Duration): String {
        val (days, hours, minutes) = parts(duration)
        return when {
            days > 0 -> listOfNotNull(
                res.getQuantityString(R.plurals.raids_span_days, days, days),
                if (hours > 0) res.getString(R.string.raids_span_hours, hours) else null,
            ).joinToString(" ")
            hours > 0 -> listOfNotNull(
                res.getString(R.string.raids_span_hours, hours),
                if (minutes > 0) res.getString(R.string.raids_span_minutes, minutes) else null,
            ).joinToString(" ")
            else -> res.getString(R.string.raids_span_minutes, minutes)
        }
    }

    /** The same span in whole words, for TalkBack (7.1: `2 days 3 hours`, never `h`). */
    fun spanWords(duration: Duration): String {
        val (days, hours, minutes) = parts(duration)
        return when {
            days > 0 -> listOfNotNull(
                res.getQuantityString(R.plurals.raids_spanWords_days, days, days),
                if (hours > 0) res.getQuantityString(R.plurals.raids_spanWords_hours, hours, hours) else null,
            ).joinToString(" ")
            hours > 0 -> listOfNotNull(
                res.getQuantityString(R.plurals.raids_spanWords_hours, hours, hours),
                if (minutes > 0) res.getQuantityString(R.plurals.raids_spanWords_minutes, minutes, minutes) else null,
            ).joinToString(" ")
            else -> res.getQuantityString(R.plurals.raids_spanWords_minutes, minutes, minutes)
        }
    }

    private fun parts(duration: Duration): Triple<Int, Int, Int> {
        val total = duration.coerceAtLeast(Duration.ZERO)
        return Triple(total.toDays().toInt(), (total.toHours() % HOURS_PER_DAY).toInt(), (total.toMinutes() % MINUTES_PER_HOUR).toInt())
    }

    /** The offline strip's `{time}`: `9:41 pm` today, `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`. */
    fun stripTime(at: Instant, now: Instant): String = gg.foreversixty.app.common.OfflineStamp.format(at, now, zone)

    private companion object {
        /** A raid within this many calendar days reads by weekday. */
        const val WEEKDAY_WINDOW_DAYS = 6L
        const val HOURS_PER_DAY = 24L
        const val MINUTES_PER_HOUR = 60L
    }
}
