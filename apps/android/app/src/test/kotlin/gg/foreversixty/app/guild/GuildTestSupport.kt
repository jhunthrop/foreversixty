package gg.foreversixty.app.guild

import gg.foreversixty.app.auth.SecretStore
import gg.foreversixty.app.guild.logic.CalendarEvent
import gg.foreversixty.app.guild.logic.CalendarSettings
import gg.foreversixty.app.guild.logic.CalendarSettingsStore
import gg.foreversixty.app.guild.logic.CalendarWriter
import gg.foreversixty.app.guild.logic.GuildService
import gg.foreversixty.app.model.CalendarFeed
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.ok

/** The F16 fixtures (docs/contracts/fixtures/app). */
object GuildFixture {
    fun home(): GuildHome = RaidFixture.load(GuildHome.serializer(), "guild.home.roster")
    fun history(page: Int = 1): GuildRaids =
        RaidFixture.load(GuildRaids.serializer(), if (page == 1) "guild.raids.history" else "guild.raids.history-2")
    fun feed(shown: Boolean = true): CalendarFeed =
        RaidFixture.load(CalendarFeed.serializer(), if (shown) "calendar.feed" else "calendar.feed.shown")

    /** The account's characters (me.json: Obnoxious Yell, Yellfury), as the Raids tab reads them. */
    val mine: Set<String> get() = RaidFixture.load(Me.serializer(), "me").characters.map { it.key }.toSet()
}

/** A scripted guild service; every call is recorded. */
class FakeGuildService : GuildService {
    var homeAnswer: () -> ApiResult<GuildHome> = { ok(GuildFixture.home()) }
    var historyAnswer: (String?) -> ApiResult<GuildRaids> = { cursor -> ok(GuildFixture.history(if (cursor == null) 1 else 2)) }
    var feedAnswer: () -> ApiResult<CalendarFeed> = { ok(GuildFixture.feed()) }
    var resetAnswer: () -> ApiResult<CalendarFeed> = {
        ok(CalendarFeed(url = "https://api.foreversixty.gg/v1/calendar/new.ics", createdAt = "2026-12-17T22:00:00Z"))
    }
    val calls = mutableListOf<String>()

    override suspend fun home(guild: Long): ApiResult<GuildHome> {
        calls += "home"
        return homeAnswer()
    }

    override suspend fun history(guild: Long, cursor: String?): ApiResult<GuildRaids> {
        calls += "history:${cursor ?: "-"}"
        return historyAnswer(cursor)
    }

    override suspend fun calendarFeed(): ApiResult<CalendarFeed> {
        calls += "feed"
        return feedAnswer()
    }

    override suspend fun resetCalendarFeed(): ApiResult<CalendarFeed> {
        calls += "reset"
        return resetAnswer()
    }
}

class MemorySecretStore : SecretStore {
    val values = mutableMapOf<String, String>()
    override suspend fun read(key: String): String? = values[key]
    override suspend fun write(key: String, value: String) {
        values[key] = value
    }
    override suspend fun clear() = values.clear()
}

class MemoryCalendarSettings(var value: CalendarSettings = CalendarSettings()) : CalendarSettingsStore {
    override fun read(): CalendarSettings = value
    override fun write(settings: CalendarSettings) {
        value = settings
    }
}

/** A calendar provider in memory: calendars by id, events by id, and a log of every write. */
class FakeCalendarWriter(var access: Boolean = true) : CalendarWriter {
    val calendars = mutableMapOf<Long, String>()
    val events = mutableMapOf<Long, Pair<Long, CalendarEvent>>()
    val log = mutableListOf<String>()
    var failWrites = false
    private var nextId = 100L

    override fun hasAccess(): Boolean = access

    override fun calendarExists(id: Long): Boolean = id in calendars

    override fun createCalendar(name: String): Long {
        val id = nextId++
        calendars[id] = name
        log += "create:$name"
        return id
    }

    override fun deleteCalendar(id: Long) {
        calendars.remove(id)
        events.entries.removeAll { it.value.first == id }
        log += "deleteCalendar"
    }

    override fun upsert(calendarId: Long, eventId: Long?, event: CalendarEvent): Long {
        check(!failWrites) { "provider refused" }
        val id = eventId?.takeIf { it in events } ?: nextId++
        events[id] = calendarId to event
        log += "upsert:${event.raidId}"
        return id
    }

    override fun delete(eventId: Long) {
        events.remove(eventId)
        log += "delete:$eventId"
    }
}
