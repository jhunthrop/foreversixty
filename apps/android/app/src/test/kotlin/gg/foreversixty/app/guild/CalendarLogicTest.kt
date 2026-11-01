package gg.foreversixty.app.guild

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.guild.logic.CalendarAlert
import gg.foreversixty.app.guild.logic.CalendarInput
import gg.foreversixty.app.guild.logic.CalendarLinkModel
import gg.foreversixty.app.guild.logic.CalendarLinkNotice
import gg.foreversixty.app.guild.logic.CalendarLinkState
import gg.foreversixty.app.guild.logic.CalendarPlanner
import gg.foreversixty.app.guild.logic.CalendarSettings
import gg.foreversixty.app.guild.logic.CalendarStatus
import gg.foreversixty.app.guild.logic.CalendarSync
import gg.foreversixty.app.guild.logic.CalendarWhich
import gg.foreversixty.app.guild.logic.SyncedEvent
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.raids.resources
import java.io.IOException
import java.time.Instant
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class CalendarLogicTest {
    private val planner = CalendarPlanner(resources(), clock())
    private val planned: List<RaidView> = RaidFixture.load(GuildRaids.serializer(), "guild.raids").planned.orEmpty()
    private val ony = planned.first { it.id == 31L }
    private val barrow = planned.first { it.id == 32L }
    private val input = CalendarInput(guildId = 2, guildName = "OLYMPUS XXVII", planned = planned)

    private fun withState(raid: RaidView, state: String?) =
        raid.copy(viewer = raid.viewer.copy(characters = raid.viewer.characters.map { it.copy(state = state) }))

    // The event builder

    @Test
    fun `the title carries your sign-up, best of your characters`() {
        assertEquals("Onyxia's Lair · In", planner.event(ony, input, CalendarAlert.None).title)
        assertEquals("Barrow Deeps · Not signed up", planner.event(barrow, input, CalendarAlert.None).title)
        assertEquals("Onyxia's Lair · Late", planner.event(withState(ony, "late"), input, CalendarAlert.None).title)
        assertEquals("Onyxia's Lair · Bench", planner.event(withState(ony, "bench"), input, CalendarAlert.None).title)
        assertEquals("Onyxia's Lair · Out", planner.event(withState(ony, "out"), input, CalendarAlert.None).title)
        val two = ony.copy(
            viewer = ony.viewer.copy(
                characters = listOf(
                    ony.viewer.characters.first().copy(characterKey = "a", state = "out"),
                    ony.viewer.characters.first().copy(characterKey = "b", state = "late"),
                ),
            ),
        )
        assertEquals("Onyxia's Lair · Late", planner.event(two, input, CalendarAlert.None).title)
    }

    @Test
    fun `the event starts at invites, else the pull, and ends at the raid's end`() {
        val event = planner.event(ony, input, CalendarAlert.None)
        assertEquals(Instant.parse("2026-12-20T00:45:00Z"), event.startsAt)
        assertEquals(Instant.parse("2026-12-20T04:00:00Z"), event.endsAt)
        assertEquals(Instant.parse("2026-12-23T01:00:00Z"), planner.event(barrow, input, CalendarAlert.None).startsAt)
    }

    @Test
    fun `location and notes`() {
        val event = planner.event(ony, input, CalendarAlert.None)
        assertEquals("OLYMPUS XXVII · Onyxia's Lair", event.location)
        assertEquals("Pull 7:00 pm. 24 of 40 signed up.\nhttps://foreversixty.gg/app/raid/2/31", event.notes)
        assertTrue(planner.event(barrow, input, CalendarAlert.None).notes.startsWith("Pull 7:00 pm. 11 signed up."))
    }

    @Test
    fun `alarm offsets`() {
        assertNull(planner.event(ony, input, CalendarAlert.None).alarmMinutes)
        assertEquals(15, planner.event(ony, input, CalendarAlert.Minutes15).alarmMinutes)
        assertEquals(60, planner.event(ony, input, CalendarAlert.HourBeforeInvites).alarmMinutes)
        assertEquals(1440, planner.event(ony, input, CalendarAlert.Day).alarmMinutes)
    }

    // The diff

    @Test
    fun `all planned upserts every raid not yet past, the locked one included`() {
        val plan = planner.plan(input, CalendarSettings(enabled = true), RaidTime.now)
        assertEquals(listOf(31L, 32L, 33L), plan.upserts.map { it.event.raidId })
        assertTrue(plan.deletes.isEmpty())
    }

    @Test
    fun `an unchanged event is not written again, a changed one updates its event`() {
        val first = planner.plan(input, CalendarSettings(enabled = true), RaidTime.now)
        val events = first.upserts.associate { it.event.raidId to SyncedEvent(it.event.raidId * 10, it.event.fingerprint) }
        val settings = CalendarSettings(enabled = true, events = events)
        assertTrue(planner.plan(input, settings, RaidTime.now).upserts.isEmpty())

        val changed = input.copy(planned = planned.map { if (it.id == 31L) withState(it, "late") else it })
        val upserts = planner.plan(changed, settings, RaidTime.now).upserts
        assertEquals(listOf(31L), upserts.map { it.event.raidId })
        assertEquals(310L, upserts.single().eventId)
    }

    @Test
    fun `which raids mine keeps only In and Late, and removes the others' events`() {
        val events = mapOf(32L to SyncedEvent(320, "old"))
        val plan = planner.plan(input, CalendarSettings(enabled = true, which = CalendarWhich.Mine, events = events), RaidTime.now)
        assertEquals(listOf(31L, 33L), plan.upserts.map { it.event.raidId })
        assertEquals(listOf(32L), plan.deletes)
    }

    @Test
    fun `a cancelled raid's event is deleted and past events stay`() {
        val cancelled = input.copy(planned = planned.map { if (it.id == 32L) it.copy(status = "cancelled") else it })
        val events = mapOf(32L to SyncedEvent(320, "x"), 7L to SyncedEvent(70, "gone from the list"))
        val plan = planner.plan(cancelled, CalendarSettings(enabled = true, events = events), RaidTime.now)
        assertEquals(listOf(32L), plan.deletes)
        val past = planner.plan(input, CalendarSettings(enabled = true, events = mapOf(31L to SyncedEvent(310, "x"))), Instant.parse("2027-01-10T00:00:00Z"))
        assertTrue(past.upserts.isEmpty())
        assertTrue(past.deletes.isEmpty())
    }

    // The sync

    @Test
    fun `turning on makes the one calendar and fills it`() {
        val writer = FakeCalendarWriter()
        val store = MemoryCalendarSettings()
        val sync = CalendarSync(writer, store, planner) { RaidTime.now }
        assertEquals(CalendarStatus.On, sync.enable(input))
        assertEquals(listOf("Forever Sixty raids"), writer.calendars.values.toList())
        assertEquals(3, writer.events.size)
        assertEquals(setOf(31L, 32L, 33L), store.value.events.keys)
        assertTrue(store.value.enabled)
    }

    @Test
    fun `a refresh rewrites only what changed and deletes what mine drops`() {
        val writer = FakeCalendarWriter()
        val store = MemoryCalendarSettings()
        val sync = CalendarSync(writer, store, planner) { RaidTime.now }
        sync.enable(input)
        writer.log.clear()
        sync.sync(input)
        assertTrue(writer.log.isEmpty())
        sync.setWhich(CalendarWhich.Mine, input)
        assertEquals(1, writer.log.count { it.startsWith("delete:") })
        assertFalse(32L in store.value.events)
    }

    @Test
    fun `a calendar the player deleted turns the switch off with its line`() {
        val writer = FakeCalendarWriter()
        val store = MemoryCalendarSettings()
        val sync = CalendarSync(writer, store, planner) { RaidTime.now }
        sync.enable(input)
        writer.calendars.clear()
        assertEquals(CalendarStatus.Removed, sync.sync(input))
        assertFalse(store.value.enabled)
        assertTrue(store.value.events.isEmpty())
        assertNull(store.value.calendarId)
    }

    @Test
    fun `access off before the flip shows nothing, after a refusal shows the card`() {
        val writer = FakeCalendarWriter(access = false)
        val store = MemoryCalendarSettings()
        val sync = CalendarSync(writer, store, planner) { RaidTime.now }
        assertEquals(CalendarStatus.Off, sync.status())
        sync.denied()
        assertEquals(CalendarStatus.AccessOff, sync.status())
        writer.access = true
        assertEquals(CalendarStatus.Off, sync.status())
        assertFalse(store.value.accessDenied)
    }

    @Test
    fun `turning off removes the calendar`() {
        val writer = FakeCalendarWriter()
        val store = MemoryCalendarSettings()
        val sync = CalendarSync(writer, store, planner) { RaidTime.now }
        sync.enable(input)
        sync.disable()
        assertTrue(writer.calendars.isEmpty())
        assertEquals(CalendarSettings(), store.value)
    }

    @Test
    fun `a refused write says the calendar did not update`() {
        val writer = FakeCalendarWriter()
        val sync = CalendarSync(writer, MemoryCalendarSettings(), planner) { RaidTime.now }
        writer.failWrites = true
        assertEquals(CalendarStatus.Failed, sync.enable(input))
    }

    @Test
    fun `nothing syncs while the switch is off`() {
        val writer = FakeCalendarWriter()
        val sync = CalendarSync(writer, MemoryCalendarSettings(), planner) { RaidTime.now }
        assertEquals(CalendarStatus.Off, sync.sync(input))
        assertTrue(writer.log.isEmpty())
    }

    // The calendar link

    @Test
    fun `a url in the answer is shown and stored for the account`() = runTest {
        val secrets = MemorySecretStore()
        val model = CalendarLinkModel(FakeGuildService(), secrets, "78")
        model.load()
        assertEquals(CalendarLinkState.Shown(GuildFixture.feed().url!!), model.state)
        assertEquals(GuildFixture.feed().url, secrets.values["calendar-link-78"])
    }

    @Test
    fun `a null url uses the stored one, or says it was shown elsewhere`() = runTest {
        val service = FakeGuildService().apply { feedAnswer = { ok(GuildFixture.feed(shown = false)) } }
        val secrets = MemorySecretStore().apply { values["calendar-link-78"] = "https://stored.ics" }
        val stored = CalendarLinkModel(service, secrets, "78")
        stored.load()
        assertEquals(CalendarLinkState.Shown("https://stored.ics"), stored.state)

        val elsewhere = CalendarLinkModel(service, MemorySecretStore(), "78")
        elsewhere.load()
        assertEquals(CalendarLinkState.Elsewhere, elsewhere.state)
    }

    @Test
    fun `reset stores the new link`() = runTest {
        val secrets = MemorySecretStore()
        val model = CalendarLinkModel(FakeGuildService(), secrets, "78")
        model.load()
        model.reset()
        assertEquals(CalendarLinkState.Shown("https://api.foreversixty.gg/v1/calendar/new.ics"), model.state)
        assertEquals("https://api.foreversixty.gg/v1/calendar/new.ics", secrets.values["calendar-link-78"])
    }

    @Test
    fun `a failed reset keeps the link and says so`() = runTest {
        val service = FakeGuildService().apply { resetAnswer = { fail(ApiError.Network(IOException("offline"))) } }
        val model = CalendarLinkModel(service, MemorySecretStore(), "78")
        model.load()
        model.reset()
        assertEquals(CalendarLinkState.Shown(GuildFixture.feed().url!!), model.state)
        assertEquals(CalendarLinkNotice.ResetFailed, model.notice)
    }

    @Test
    fun `offline shows the stored link with the buttons off, and an error without one`() = runTest {
        val service = FakeGuildService().apply { feedAnswer = { fail(ApiError.Network(IOException("offline"))) } }
        val secrets = MemorySecretStore().apply { values["calendar-link-78"] = "https://stored.ics" }
        val stored = CalendarLinkModel(service, secrets, "78")
        stored.load()
        assertEquals(CalendarLinkState.Shown("https://stored.ics"), stored.state)
        assertTrue(stored.offline)

        val none = CalendarLinkModel(service, MemorySecretStore(), "78")
        none.load()
        assertEquals(CalendarLinkState.Failed, none.state)
        assertEquals(CalendarLinkNotice.LoadFailed, none.notice)
    }

    @Test
    fun `copy says Link copied`() = runTest {
        val model = CalendarLinkModel(FakeGuildService(), MemorySecretStore(), "78")
        model.load()
        assertEquals(GuildFixture.feed().url, model.copy())
        assertEquals(CalendarLinkNotice.Copied, model.notice)
    }
}
