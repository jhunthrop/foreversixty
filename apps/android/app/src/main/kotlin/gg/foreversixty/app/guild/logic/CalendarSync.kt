package gg.foreversixty.app.guild.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import java.time.Instant

/**
 * The device calendar, behind a thin adapter (Android: CalendarContract) so tests use a fake.
 * Every call blocks; the caller runs it off the main thread.
 */
interface CalendarWriter {
    fun hasAccess(): Boolean
    fun calendarExists(id: Long): Boolean
    fun createCalendar(name: String): Long
    fun deleteCalendar(id: Long)

    /** Updates [eventId] when it still exists, else inserts; returns the event's id. */
    fun upsert(calendarId: Long, eventId: Long?, event: CalendarEvent): Long
    fun delete(eventId: Long)
}

interface CalendarSettingsStore {
    fun read(): CalendarSettings
    fun write(settings: CalendarSettings)
}

/** What the switch and the line under it show. */
enum class CalendarStatus { Off, On, AccessOff, Removed, Failed }

/**
 * Keeps the app's one calendar in step with the planned raids (spec F16.4): [enable] after the
 * player flips the switch and grants access, [sync] on every Raids refresh and sign-up answer.
 */
class CalendarSync(
    private val writer: CalendarWriter,
    private val store: CalendarSettingsStore,
    private val planner: CalendarPlanner,
    private val calendarName: String = DEFAULT_NAME,
    private val now: () -> Instant,
) {
    var settings: CalendarSettings by mutableStateOf(store.read())
        private set

    /** The status of the last write; [status] folds in access and the switch. */
    private var lastWrite: CalendarStatus? by mutableStateOf(null)

    private fun save(next: CalendarSettings) {
        settings = next
        store.write(next)
    }

    /** The switch's state now, without writing. */
    fun status(): CalendarStatus {
        val access = writer.hasAccess()
        if (access && settings.accessDenied) save(settings.copy(accessDenied = false))
        return when {
            !access && (settings.enabled || settings.accessDenied) -> CalendarStatus.AccessOff
            lastWrite == CalendarStatus.Removed && !settings.enabled -> CalendarStatus.Removed
            !settings.enabled -> CalendarStatus.Off
            else -> lastWrite ?: CalendarStatus.On
        }
    }

    /** The player refused access after the flip. */
    fun denied() {
        save(settings.copy(enabled = false, accessDenied = true))
    }

    /** The switch turned on and access is granted: make the calendar and fill it. */
    fun enable(input: CalendarInput?): CalendarStatus {
        val id = runCatching { writer.createCalendar(calendarName) }.getOrNull() ?: return record(CalendarStatus.Failed)
        save(settings.copy(enabled = true, accessDenied = false, calendarId = id, events = emptyMap()))
        lastWrite = null
        return sync(input)
    }

    /** The switch turned off: the calendar and its events go. */
    fun disable() {
        settings.calendarId?.let { id -> runCatching { writer.deleteCalendar(id) } }
        save(CalendarSettings(which = settings.which, alert = settings.alert))
        lastWrite = null
    }

    fun setWhich(which: CalendarWhich, input: CalendarInput?): CalendarStatus {
        save(settings.copy(which = which))
        return sync(input)
    }

    fun setAlert(alert: CalendarAlert, input: CalendarInput?): CalendarStatus {
        save(settings.copy(alert = alert))
        return sync(input)
    }

    /** One pass: upsert what changed, delete what went; a deleted calendar turns the switch off. */
    fun sync(input: CalendarInput?): CalendarStatus {
        val current = settings
        if (!current.enabled) return status()
        if (!writer.hasAccess()) return CalendarStatus.AccessOff
        val calendarId = current.calendarId
        if (calendarId == null || !writer.calendarExists(calendarId)) {
            save(CalendarSettings(which = current.which, alert = current.alert))
            return record(CalendarStatus.Removed)
        }
        if (input == null) return record(CalendarStatus.On)
        return apply(calendarId, planner.plan(input, current, now()), current)
    }

    private fun apply(calendarId: Long, plan: CalendarPlan, current: CalendarSettings): CalendarStatus {
        val events = current.events.toMutableMap()
        val outcome = runCatching {
            plan.deletes.forEach { raid -> events.remove(raid)?.let { writer.delete(it.eventId) } }
            plan.upserts.forEach { upsert ->
                val id = writer.upsert(calendarId, upsert.eventId, upsert.event)
                events[upsert.event.raidId] = SyncedEvent(id, upsert.event.fingerprint)
            }
        }
        save(current.copy(events = events))
        return record(if (outcome.isSuccess) CalendarStatus.On else CalendarStatus.Failed)
    }

    private fun record(status: CalendarStatus): CalendarStatus {
        lastWrite = status
        return status
    }

    companion object {
        const val DEFAULT_NAME = "Forever Sixty raids"
    }
}
