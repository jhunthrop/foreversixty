package gg.foreversixty.app.guild.platform

import android.Manifest
import android.content.ContentResolver
import android.content.ContentValues
import android.content.Context
import android.content.SharedPreferences
import android.content.pm.PackageManager
import android.net.Uri
import android.provider.CalendarContract
import androidx.core.content.ContextCompat
import androidx.core.content.edit
import gg.foreversixty.app.guild.logic.CalendarEvent
import gg.foreversixty.app.guild.logic.CalendarSettings
import gg.foreversixty.app.guild.logic.CalendarSettingsStore
import gg.foreversixty.app.guild.logic.CalendarWriter
import gg.foreversixty.app.model.ApiJson
import java.time.ZoneOffset
import kotlinx.serialization.SerializationException

/** The runtime permissions the switch asks for, after the flip (spec F16.4). */
val CALENDAR_PERMISSIONS = arrayOf(Manifest.permission.READ_CALENDAR, Manifest.permission.WRITE_CALENDAR)

/**
 * The app's calendar through CalendarContract: one local-account calendar the app owns, its
 * events, and each event's one reminder. Every call blocks on the provider.
 */
class AndroidCalendarWriter(private val context: Context) : CalendarWriter {
    private val resolver: ContentResolver get() = context.contentResolver

    override fun hasAccess(): Boolean = CALENDAR_PERMISSIONS.all {
        ContextCompat.checkSelfPermission(context, it) == PackageManager.PERMISSION_GRANTED
    }

    override fun calendarExists(id: Long): Boolean =
        resolver.query(
            CalendarContract.Calendars.CONTENT_URI,
            arrayOf(CalendarContract.Calendars._ID),
            "${CalendarContract.Calendars._ID} = ? AND ${CalendarContract.Calendars.DELETED} = 0",
            arrayOf(id.toString()),
            null,
        )?.use { it.count > 0 } == true

    override fun createCalendar(name: String): Long {
        val values = ContentValues().apply {
            put(CalendarContract.Calendars.ACCOUNT_NAME, ACCOUNT_NAME)
            put(CalendarContract.Calendars.ACCOUNT_TYPE, CalendarContract.ACCOUNT_TYPE_LOCAL)
            put(CalendarContract.Calendars.NAME, name)
            put(CalendarContract.Calendars.CALENDAR_DISPLAY_NAME, name)
            put(CalendarContract.Calendars.CALENDAR_COLOR, CALENDAR_COLOR)
            put(CalendarContract.Calendars.CALENDAR_ACCESS_LEVEL, CalendarContract.Calendars.CAL_ACCESS_OWNER)
            put(CalendarContract.Calendars.OWNER_ACCOUNT, ACCOUNT_NAME)
            put(CalendarContract.Calendars.VISIBLE, 1)
            put(CalendarContract.Calendars.SYNC_EVENTS, 1)
        }
        val uri = requireNotNull(resolver.insert(asSyncAdapter(CalendarContract.Calendars.CONTENT_URI), values)) { "calendar not created" }
        return uri.lastPathSegment?.toLongOrNull() ?: error("calendar id missing")
    }

    override fun deleteCalendar(id: Long) {
        resolver.delete(asSyncAdapter(CalendarContract.Calendars.CONTENT_URI), "${CalendarContract.Calendars._ID} = ?", arrayOf(id.toString()))
    }

    override fun upsert(calendarId: Long, eventId: Long?, event: CalendarEvent): Long {
        val values = ContentValues().apply {
            put(CalendarContract.Events.CALENDAR_ID, calendarId)
            put(CalendarContract.Events.TITLE, event.title)
            put(CalendarContract.Events.DTSTART, event.startsAt.toEpochMilli())
            put(CalendarContract.Events.DTEND, event.endsAt.toEpochMilli())
            put(CalendarContract.Events.EVENT_TIMEZONE, ZoneOffset.UTC.id)
            put(CalendarContract.Events.EVENT_LOCATION, event.location)
            put(CalendarContract.Events.DESCRIPTION, event.notes)
            put(CalendarContract.Events.HAS_ALARM, if (event.alarmMinutes != null) 1 else 0)
        }
        // An event the player deleted by hand is inserted again.
        val id = eventId?.takeIf { resolver.update(eventUri(it), values, null, null) > 0 } ?: insert(values)
        writeReminder(id, event.alarmMinutes)
        return id
    }

    private fun insert(values: ContentValues): Long {
        val uri = requireNotNull(resolver.insert(CalendarContract.Events.CONTENT_URI, values)) { "event not created" }
        return uri.lastPathSegment?.toLongOrNull() ?: error("event id missing")
    }

    override fun delete(eventId: Long) {
        resolver.delete(eventUri(eventId), null, null)
    }

    private fun writeReminder(eventId: Long, minutes: Int?) {
        resolver.delete(CalendarContract.Reminders.CONTENT_URI, "${CalendarContract.Reminders.EVENT_ID} = ?", arrayOf(eventId.toString()))
        if (minutes == null) return
        resolver.insert(
            CalendarContract.Reminders.CONTENT_URI,
            ContentValues().apply {
                put(CalendarContract.Reminders.EVENT_ID, eventId)
                put(CalendarContract.Reminders.MINUTES, minutes)
                put(CalendarContract.Reminders.METHOD, CalendarContract.Reminders.METHOD_ALERT)
            },
        )
    }

    private fun eventUri(id: Long): Uri = Uri.withAppendedPath(CalendarContract.Events.CONTENT_URI, id.toString())

    /** Only a sync adapter may create a calendar; a local account is its own adapter. */
    private fun asSyncAdapter(uri: Uri): Uri = uri.buildUpon()
        .appendQueryParameter(CalendarContract.CALLER_IS_SYNCADAPTER, "true")
        .appendQueryParameter(CalendarContract.Calendars.ACCOUNT_NAME, ACCOUNT_NAME)
        .appendQueryParameter(CalendarContract.Calendars.ACCOUNT_TYPE, CalendarContract.ACCOUNT_TYPE_LOCAL)
        .build()

    private companion object {
        const val ACCOUNT_NAME = "Forever Sixty"

        /** The site's gold, `#e5b955`, as the calendar's colour. */
        const val CALENDAR_COLOR = 0xFFE5B955.toInt()
    }
}

/** The sync settings and the raid-to-event mapping as one JSON value in the app's preferences. */
class PrefsCalendarSettingsStore(private val prefs: SharedPreferences) : CalendarSettingsStore {
    override fun read(): CalendarSettings = prefs.getString(KEY, null)?.let {
        try {
            ApiJson.decodeFromString(CalendarSettings.serializer(), it)
        } catch (_: SerializationException) {
            null
        } catch (_: IllegalArgumentException) {
            null
        }
    } ?: CalendarSettings()

    override fun write(settings: CalendarSettings) {
        prefs.edit { putString(KEY, ApiJson.encodeToString(CalendarSettings.serializer(), settings)) }
    }

    private companion object {
        const val KEY = "calendar_settings"
    }
}
