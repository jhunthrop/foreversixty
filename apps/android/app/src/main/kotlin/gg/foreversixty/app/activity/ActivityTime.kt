package gg.foreversixty.app.activity

import android.content.res.Resources
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.phrase
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/**
 * A feed row's when, the one formatter Activity and Home's New for you read (round 3 row 3), so one row
 * never reads two ways: Today `8 h ago`, Yesterday `6:02 pm`, This week `Tue`, Earlier `3 Nov`.
 */
object ActivityTime {
    private val CLOCK = DateTimeFormatter.ofPattern("h:mm a", Locale.UK)
    private val WEEKDAY = DateTimeFormatter.ofPattern("EEE", Locale.UK)
    private val DATE = DateTimeFormatter.ofPattern("d MMM", Locale.UK)

    fun label(res: Resources, at: Instant, now: Instant, zone: ZoneId): String =
        label(res, at, Activity.dayOf(at.atZone(zone).toLocalDate(), now.atZone(zone).toLocalDate()), now, zone)

    fun label(res: Resources, at: Instant, day: ActivityDay, now: Instant, zone: ZoneId): String = when (day) {
        ActivityDay.Today -> RelativeAge.of(at, now).phrase(res)
        ActivityDay.Yesterday -> CLOCK.withZone(zone).format(at).lowercase(Locale.UK)
        ActivityDay.ThisWeek -> WEEKDAY.withZone(zone).format(at)
        ActivityDay.Earlier -> DATE.withZone(zone).format(at)
    }
}
