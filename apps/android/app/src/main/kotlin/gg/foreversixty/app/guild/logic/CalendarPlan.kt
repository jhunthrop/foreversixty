package gg.foreversixty.app.guild.logic

import android.content.res.Resources
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.RaidChip
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.RaidLinks
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Instant
import kotlinx.serialization.Serializable

/** The `Which raids` row. */
enum class CalendarWhich(@param:StringRes val label: Int) {
    All(R.string.calendar_which_all),
    Mine(R.string.calendar_which_mine),
}

/** The `Alert` row: minutes before the event's start (the invite time, else the pull). */
enum class CalendarAlert(val minutesBefore: Int?, @param:StringRes val label: Int) {
    None(null, R.string.calendar_alert_none),
    Minutes15(15, R.string.calendar_alert_minutes15),
    HourBeforeInvites(60, R.string.calendar_alert_hourBeforeInvites),
    Day(24 * 60, R.string.calendar_alert_day),
}

/** One raid's event as written, and what it looked like (so an unchanged event is not rewritten). */
@Serializable
data class SyncedEvent(val eventId: Long, val fingerprint: String)

/** The calendar sync's settings and its raid-to-event mapping, persisted on the device. */
@Serializable
data class CalendarSettings(
    val enabled: Boolean = false,
    val which: CalendarWhich = CalendarWhich.All,
    val alert: CalendarAlert = CalendarAlert.HourBeforeInvites,
    val calendarId: Long? = null,
    /** Raid id to its event. */
    val events: Map<Long, SyncedEvent> = emptyMap(),
    /** The player refused calendar access after flipping the switch. */
    val accessDenied: Boolean = false,
)

/** One guild's planned raids, as the Raids tab last loaded them. */
data class CalendarInput(val guildId: Long, val guildName: String, val planned: List<RaidView>)

/** One event in the app's calendar (spec F16.4). */
data class CalendarEvent(
    val raidId: Long,
    val title: String,
    val startsAt: Instant,
    val endsAt: Instant,
    val location: String,
    val notes: String,
    val alarmMinutes: Int?,
) {
    val fingerprint: String get() = listOf(title, startsAt, endsAt, location, notes, alarmMinutes).joinToString("|")
}

/** An event to write; [eventId] is the existing event to update, null to insert. */
data class CalendarUpsert(val event: CalendarEvent, val eventId: Long?)

/** What one sync writes: upserts, and the raid ids whose events go. */
data class CalendarPlan(val upserts: List<CalendarUpsert>, val deletes: List<Long>)

/** Builds each raid's event and diffs the planned raids against what the calendar holds. */
class CalendarPlanner(private val res: Resources, private val clock: RaidClock) {
    fun event(raid: RaidView, input: CalendarInput, alert: CalendarAlert): CalendarEvent {
        val facts = RaidFacts(raid, ANY_INSTANT)
        val zone = raid.zoneName
        val pull = clock.clock(facts.startsAt)
        val signed = facts.signedUp
        val counts = raid.size?.let { res.getString(R.string.calendar_event_notes, pull, signed, it) }
            ?: res.getString(R.string.calendar_event_notesNoSize, pull, signed)
        return CalendarEvent(
            raidId = raid.id,
            title = res.getString(R.string.calendar_event_title, zone, res.getString(stateWord(bestState(raid)))),
            startsAt = facts.inviteAt ?: facts.startsAt,
            endsAt = facts.endsAt,
            location = res.getString(R.string.calendar_event_location, input.guildName, zone),
            notes = "$counts\n${RaidLinks.share(input.guildId, raid.id)}",
            alarmMinutes = alert.minutesBefore,
        )
    }

    fun plan(input: CalendarInput, settings: CalendarSettings, now: Instant): CalendarPlan {
        val upserts = mutableListOf<CalendarUpsert>()
        val deletes = mutableListOf<Long>()
        input.planned.forEach { raid ->
            val mapped = settings.events[raid.id]
            val chip = RaidFacts(raid, now).chip
            val wanted = settings.which == CalendarWhich.All || bestState(raid)?.isComing == true
            when {
                chip == RaidChip.Cancelled -> if (mapped != null) deletes += raid.id
                chip == RaidChip.Ended || chip == RaidChip.Done -> Unit
                !wanted -> if (mapped != null) deletes += raid.id
                else -> {
                    val event = event(raid, input, settings.alert)
                    if (mapped?.fingerprint != event.fingerprint) upserts += CalendarUpsert(event, mapped?.eventId)
                }
            }
        }
        return CalendarPlan(upserts, deletes)
    }

    /** The best answer of the viewer's characters: In, then Late, Bench, Out. */
    private fun bestState(raid: RaidView): SignupState? =
        raid.viewer.characters.mapNotNull { SignupState.of(it.state) }.minByOrNull { it.ordinal }

    @StringRes
    private fun stateWord(state: SignupState?): Int = when (state) {
        SignupState.In -> R.string.calendar_event_in
        SignupState.Late -> R.string.calendar_event_late
        SignupState.Bench -> R.string.calendar_event_bench
        SignupState.Out -> R.string.calendar_event_out
        null -> R.string.calendar_event_notSigned
    }

    private companion object {
        /** The event's own times do not depend on the instant; any instant reads them. */
        val ANY_INSTANT: Instant = Instant.EPOCH
    }
}
