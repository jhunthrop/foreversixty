package gg.foreversixty.app.widgets.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.Widget
import gg.foreversixty.app.model.WidgetNextRaid
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Duration
import java.time.Instant
import java.time.LocalDate
import java.time.ZoneId

/**
 * The widget's four lines (mocks `widgets-guild`, `-noguild`, `-open`): the label ([key]), the big
 * figure, the name under it and the muted line. [lock] is the one-line form the quick-settings
 * tile shows (`33d 8h · Onyxia's Lair`). [counts] is the medium widget's `24 in, 0 late`.
 */
data class RaidCard(
    val key: String,
    val big: String,
    val name: String,
    val line: String?,
    val lock: String,
    val cancelled: Boolean = false,
    val counts: String? = null,
)

/** The answer as the card says it: `You are In`, `Not signed up`, `Choose a spec`, `Locked`. */
data class CardAnswer(val state: SignupState?, val asksSpec: Boolean, val locked: Boolean)

/**
 * Derives [RaidCard] from the snapshot. Before the first raids a guildless player counts down to
 * them (`First raids`, the day, the zones); after, any player with no raid planned sees
 * `Raids open now`. A snapshot an older build cached has no first raids fields and gets no card.
 */
class RaidCards(private val res: Resources, private val zone: ZoneId, private val clock: RaidClock) {
    fun noGuild(widget: Widget, now: Instant): RaidCard? {
        val opens = firstRaidsStart(widget)
        if (opens != null && now < opens) {
            return card(res.getString(R.string.widget_card_firstRaids), span(Duration.between(now, opens)), clock.shortDate(opens), zones(widget))
        }
        return open(widget, res.getString(R.string.widget_noGuild))
    }

    fun noRaid(widget: Widget, guildName: String, now: Instant): RaidCard? {
        val opens = firstRaidsStart(widget)
        if (opens != null && now < opens) {
            val day = res.getString(R.string.widget_card_firstRaidsOn, clock.shortDate(opens))
            return card(guildName, span(Duration.between(now, opens)), day, res.getString(R.string.widget_card_noRaid))
        }
        return open(widget, res.getString(R.string.widget_card_noRaid))
    }

    fun raid(raid: WidgetNextRaid, countdown: Countdown?, answer: CardAnswer): RaidCard {
        val start = Instant.parse(raid.startsAt)
        val cancelled = raid.status == WidgetFigures.STATUS_CANCELLED
        val big = when {
            cancelled -> res.getString(R.string.widget_cancelled)
            countdown is Countdown.Left -> span(countdown)
            else -> res.getString(R.string.widget_card_now)
        }
        val whenText = res.getString(R.string.raids_when_date, clock.shortDate(start), clock.guildClock(start, zone))
        val line = if (cancelled) whenText else res.getString(R.string.widget_card_line, whenText, answerWords(answer))
        val counts = raid.inCount?.let { res.getString(R.string.widget_card_counts, it, raid.late ?: 0) }
        return card(raid.guildName, big, raid.zoneName, line, cancelled, counts)
    }

    private fun answerWords(answer: CardAnswer): String = when {
        answer.asksSpec -> res.getString(R.string.widget_chooseSpec)
        answer.state != null -> res.getString(R.string.raid_night_youAre, res.getString(answer.state.label))
        answer.locked -> res.getString(R.string.widget_locked)
        else -> res.getString(R.string.widget_notSignedUp)
    }

    private fun open(widget: Widget, line: String): RaidCard? {
        val zones = zones(widget) ?: return null
        return card(res.getString(R.string.widget_card_raidsOpen), res.getString(R.string.widget_card_open), zones, line)
    }

    private fun card(key: String, big: String, name: String, line: String?, cancelled: Boolean = false, counts: String? = null) =
        RaidCard(key, big, name, line, res.getString(R.string.widget_card_lock, big, name), cancelled, counts)

    /** The first raids open at the start of their day in the player's zone (the API sends the date). */
    private fun firstRaidsStart(widget: Widget): Instant? =
        widget.firstRaidsAt?.let { LocalDate.parse(it).atStartOfDay(zone).toInstant() }

    private fun zones(widget: Widget): String? = widget.firstRaidsZones?.takeIf { it.isNotEmpty() }?.joinToString(", ")

    private fun span(duration: Duration): String {
        val total = WidgetFigures.ceilMinutes(duration)
        return span(Countdown.Left((total / MINUTES_PER_DAY).toInt(), (total % MINUTES_PER_DAY / MINUTES_PER_HOUR).toInt(), (total % MINUTES_PER_HOUR).toInt()))
    }

    /** `33d 8h`, `2d`, `5h 20m`, `5h`, `12m`. */
    private fun span(left: Countdown.Left): String = when {
        left.days >= 1 && left.hours > 0 -> res.getString(R.string.widget_card_daysHours, left.days, left.hours)
        left.days >= 1 -> res.getString(R.string.widget_short_days, left.days)
        left.hours >= 1 && left.minutes > 0 -> res.getString(R.string.widget_card_hoursMinutes, left.hours, left.minutes)
        left.hours >= 1 -> res.getString(R.string.widget_short_hours, left.hours)
        else -> res.getString(R.string.widget_short_minutes, left.minutes)
    }

    private companion object {
        const val MINUTES_PER_DAY = 1440L
        const val MINUTES_PER_HOUR = 60L
    }
}
