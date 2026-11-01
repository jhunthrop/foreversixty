@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@Suite("Calendar events")
struct CalendarPlanTests {
    let now = RaidTime.now

    private func plan(
        _ raids: [RaidView]? = nil,
        settings: CalendarSettings = CalendarSettings(enabled: true),
        mapping: [Int64: CalendarMappingEntry] = [:]
    ) throws -> CalendarDiff {
        let input = CalendarInput(guildID: 2, guildName: "OLYMPUS XXVII", raids: try raids ?? GuildFixture.planned())
        return CalendarPlan.diff(input: input, settings: settings, mapping: mapping, now: now, clock: RaidTime.clock)
    }

    private func draft(_ id: Int64, settings: CalendarSettings = CalendarSettings(enabled: true)) throws -> CalendarEventDraft {
        try #require(try plan(settings: settings).upserts.first { $0.draft.raidID == id }).draft
    }

    @Test("the title is the zone and your best answer, In over Late over Bench over Out")
    func title() throws {
        #expect(try draft(31).title == "Onyxia's Lair · In")
        #expect(try draft(32).title == "Barrow Deeps · Not signed up")
        #expect(CalendarPlan.state(of: ["out", "bench"]) == Copy.calendarEventBench)
        #expect(CalendarPlan.state(of: ["late", "in"]) == Copy.calendarEventIn)
        #expect(CalendarPlan.state(of: ["out"]) == Copy.calendarEventOut)
        #expect(CalendarPlan.state(of: []) == Copy.calendarEventNotSigned)
    }

    @Test("the event starts at invites, else at the pull, and ends at the raid's end")
    func times() throws {
        let ony = try draft(31)
        #expect(ony.start == RaidTime.date("2026-12-20T00:45:00Z"))
        #expect(ony.end == RaidTime.date("2026-12-20T04:00:00Z"))
        #expect(try draft(32).start == RaidTime.date("2026-12-23T01:00:00Z"))
    }

    @Test("location, notes and the raid's link")
    func details() throws {
        let ony = try draft(31)
        #expect(ony.location == "OLYMPUS XXVII · Onyxia's Lair")
        #expect(ony.notes == "Pull 7:00 pm. 24 of 40 signed up.\nhttps://foreversixty.gg/app/raid/2/31")
        #expect(ony.url.absoluteString == "https://foreversixty.gg/app/raid/2/31")
        #expect(try draft(32).notes.hasPrefix("Pull 7:00 pm. 11 signed up."))
    }

    @Test("alarm offsets: none, 15 minutes, 1 hour before invites, 1 day")
    func alarms() throws {
        #expect(CalendarAlert.none.offset == nil)
        #expect(CalendarAlert.minutes15.offset == -900)
        #expect(CalendarAlert.hourBeforeInvites.offset == -3600)
        #expect(CalendarAlert.day.offset == -86400)
        #expect(CalendarSettings().alert == .hourBeforeInvites)
        #expect(try draft(31).alarmOffset == -3600)
        #expect(try draft(31, settings: CalendarSettings(enabled: true, alert: .none)).alarmOffset == nil)
    }

    @Test("an upsert keeps the event it already has")
    func upsertKeepsEvent() throws {
        let entry = CalendarMappingEntry(eventID: "ev-31", end: RaidTime.date("2026-12-20T04:00:00Z"))
        let diff = try plan(mapping: [31: entry])
        #expect(diff.upserts.first { $0.draft.raidID == 31 }?.eventID == "ev-31")
        #expect(diff.upserts.first { $0.draft.raidID == 32 }?.eventID == nil)
    }

    @Test("Which raids = mine keeps In and Late only and removes the others' events")
    func whichMine() throws {
        let mine = CalendarSettings(enabled: true, which: .mine)
        let mapping: [Int64: CalendarMappingEntry] = [32: CalendarMappingEntry(eventID: "ev-32", end: RaidTime.date("2026-12-23T04:00:00Z"))]
        let diff = try plan(settings: mine, mapping: mapping)
        #expect(Set(diff.upserts.map(\.draft.raidID)) == [31, 33])
        #expect(diff.deletes == ["ev-32"])
    }

    @Test("a cancelled raid's event is deleted; a vanished future raid too; past events stay")
    func deletes() throws {
        let cancelled = try RaidFixture.raid("guild.raid.cancelled")
        let mapping: [Int64: CalendarMappingEntry] = [
            31: CalendarMappingEntry(eventID: "ev-31", end: RaidTime.date("2026-12-20T04:00:00Z")),
            98: CalendarMappingEntry(eventID: "ev-98", end: RaidTime.date("2026-12-30T04:00:00Z")),
            99: CalendarMappingEntry(eventID: "ev-99", end: RaidTime.date("2026-12-10T04:00:00Z")),
        ]
        let diff = try plan([cancelled], mapping: mapping)
        #expect(diff.upserts.isEmpty)
        #expect(Set(diff.deletes) == ["ev-31", "ev-98"])
    }
}

@MainActor
@Suite("Calendar sync")
struct CalendarSyncTests {
    private func controller(
        _ adapter: FakeCalendarAdapter, store: MemoryCalendarSyncStore = MemoryCalendarSyncStore()
    ) -> CalendarSyncController {
        CalendarSyncController(
            adapter: adapter, store: store, now: { RaidTime.now }, clock: RaidTime.clock,
            loadInput: { CalendarInput(guildID: 2, guildName: "OLYMPUS XXVII", raids: (try? GuildFixture.planned()) ?? []) }
        )
    }

    @Test("access is never asked before the switch turns on")
    func noEarlyAsk() async throws {
        let adapter = FakeCalendarAdapter()
        let sync = controller(adapter)
        #expect(sync.phase == .off)
        await sync.update(CalendarInput(guildID: 2, guildName: "G", raids: try GuildFixture.planned()))
        #expect(adapter.requests == 0)
        #expect(adapter.events.isEmpty)
    }

    @Test("turning it on asks, makes the calendar and writes one event per planned raid")
    func turnOn() async throws {
        let adapter = FakeCalendarAdapter()
        let store = MemoryCalendarSyncStore()
        let sync = controller(adapter, store: store)
        await sync.setEnabled(true)
        #expect(adapter.requests == 1)
        #expect(sync.phase == .on)
        #expect(adapter.events.count == 3)
        #expect(store.record.settings.enabled)
        #expect(store.record.mapping.count == 3)
        #expect(store.record.calendarID != nil)
    }

    @Test("denied: the access-off card, the switch off")
    func denied() async throws {
        let adapter = FakeCalendarAdapter()
        adapter.grants = false
        let sync = controller(adapter)
        await sync.setEnabled(true)
        #expect(sync.phase == .accessOff)
        #expect(!sync.record.settings.enabled)
        #expect(adapter.events.isEmpty)

        let refused = FakeCalendarAdapter()
        refused.accessAnswer = .denied
        #expect(controller(refused).phase == .accessOff)
    }

    @Test("a deleted calendar turns the switch off with the removed line")
    func removed() async throws {
        let adapter = FakeCalendarAdapter()
        let store = MemoryCalendarSyncStore()
        let sync = controller(adapter, store: store)
        await sync.setEnabled(true)
        adapter.deleteCalendarOutsideTheApp()
        await sync.update(CalendarInput(guildID: 2, guildName: "G", raids: try GuildFixture.planned()))
        #expect(sync.phase == .removed)
        #expect(!store.record.settings.enabled)
        #expect(store.record.mapping.isEmpty)
        #expect(controller(adapter, store: store).phase == .off)
    }

    @Test("Which raids = mine resyncs and removes the event of the raid you are not In for")
    func whichResync() async throws {
        let adapter = FakeCalendarAdapter()
        let sync = controller(adapter)
        await sync.setEnabled(true)
        await sync.setWhich(.mine)
        #expect(adapter.events.count == 2)
        #expect(adapter.removed == ["event-32"])
        #expect(sync.record.settings.which == .mine)
    }

    @Test("turning it off removes the app's calendar")
    func turnOff() async throws {
        let adapter = FakeCalendarAdapter()
        let sync = controller(adapter)
        await sync.setEnabled(true)
        await sync.setEnabled(false)
        #expect(sync.phase == .off)
        #expect(adapter.calendars.isEmpty)
        #expect(sync.record.mapping.isEmpty)
    }
}

@MainActor
@Suite("Calendar link")
struct CalendarLinkTests {
    private func model(_ service: FakeGuildExtrasService, store: MemoryCalendarLinkStore) -> CalendarLinkModel {
        CalendarLinkModel(service: service, store: store, account: "1812")
    }

    @Test("a url shown once is the link, and it goes into the secure store")
    func shown() async throws {
        let store = MemoryCalendarLinkStore()
        let link = model(FakeGuildExtrasService(), store: store)
        await link.load()
        let url = try GuildFixture.feed().url
        #expect(link.phase == .ready(try #require(url)))
        #expect(store.saved["1812"] == url)
    }

    @Test("a null url uses the stored link; none stored says it was shown elsewhere")
    func storedOrElsewhere() async throws {
        let service = FakeGuildExtrasService()
        service.feedAnswer = { () throws(APIError) in try RaidFixture.api(CalendarFeed.self, "calendar.feed.shown") }
        let store = MemoryCalendarLinkStore()
        store.saved = ["1812": "https://api.foreversixty.gg/v1/calendar/old.ics"]
        let stored = model(service, store: store)
        await stored.load()
        #expect(stored.phase == .ready("https://api.foreversixty.gg/v1/calendar/old.ics"))

        let other = MemoryCalendarLinkStore()
        other.saved = ["99": "https://api.foreversixty.gg/v1/calendar/someone-else.ics"]
        let elsewhere = model(service, store: other)
        await elsewhere.load()
        #expect(elsewhere.phase == .elsewhere)
    }

    @Test("reset issues a new link and stores it; a failed reset keeps the old one and says so")
    func reset() async throws {
        let service = FakeGuildExtrasService()
        service.feedAnswer = { () throws(APIError) in try RaidFixture.api(CalendarFeed.self, "calendar.feed.shown") }
        let store = MemoryCalendarLinkStore()
        let link = model(service, store: store)
        await link.load()
        #expect(link.phase == .elsewhere)
        await link.reset()
        let url = try #require(try GuildFixture.feed().url)
        #expect(link.phase == .ready(url))
        #expect(store.saved["1812"] == url)
        #expect(service.resets == 1)

        service.resetAnswer = { () throws(APIError) in throw .transport("down") }
        await link.reset()
        #expect(link.phase == .ready(url))
        #expect(link.message == Copy.calendarResetError)
    }

    @Test("offline keeps the stored link with the buttons off; an error with nothing stored is the error")
    func failures() async throws {
        let service = FakeGuildExtrasService()
        service.feedAnswer = { () throws(APIError) in throw .transport("down") }
        let store = MemoryCalendarLinkStore()
        store.saved = ["1812": "https://api.foreversixty.gg/v1/calendar/old.ics"]
        let offline = model(service, store: store)
        await offline.load()
        #expect(offline.phase == .ready("https://api.foreversixty.gg/v1/calendar/old.ics"))
        #expect(offline.offline)

        let failing = FakeGuildExtrasService()
        failing.feedAnswer = { () throws(APIError) in throw .server(status: 500, body: ErrorBody(code: "internal", message: "x"), requestID: nil) }
        let broken = model(failing, store: MemoryCalendarLinkStore())
        await broken.load()
        #expect(broken.phase == .failed)
        #expect(broken.message == Copy.calendarLinkError)
    }

    @Test("the webcal link swaps the scheme only")
    func webcal() throws {
        #expect(CalendarLinkModel.webcal("https://api.foreversixty.gg/v1/calendar/x.ics")?.absoluteString == "webcal://api.foreversixty.gg/v1/calendar/x.ics")
    }
}
