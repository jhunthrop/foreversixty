@testable import ForeverSixty
import Foundation
import Models
import Networking

/// The F16 fixtures (docs/contracts/fixtures/app) and fakes for the guild extras' stores.
enum GuildFixture {
    static func home() throws -> GuildHome { try RaidFixture.load(GuildHome.self, "guild.home.roster") }
    static func historyPage1() throws -> GuildRaids { try RaidFixture.load(GuildRaids.self, "guild.raids.history") }
    static func historyPage2() throws -> GuildRaids { try RaidFixture.load(GuildRaids.self, "guild.raids.history-2") }
    static func planned() throws -> [RaidView] { try RaidFixture.load(GuildRaids.self, "guild.raids").planned ?? [] }
    static func feed(_ name: String = "calendar.feed") throws -> CalendarFeed { try RaidFixture.load(CalendarFeed.self, name) }

    /// The roster fixture's viewer, me.json's Obnoxious Yell.
    static let viewerKey = "us/pvp/obnoxious-yell"

    /// "Your characters": every character on the account (me.json), as the app reads them.
    static func yourKeys() throws -> Set<String> { Set(try RaidFixture.load(Me.self, "me").characters.map(\.key)) }
}

/// A scripted guild-extras service; each call answers from its closure and is counted.
final class FakeGuildExtrasService: GuildExtrasService, @unchecked Sendable {
    var homeAnswer: () throws(APIError) -> GuildHome = { () throws(APIError) in try RaidFixture.api(GuildHome.self, "guild.home.roster") }
    var historyAnswer: (String?) throws(APIError) -> GuildRaids = { cursor throws(APIError) in
        try RaidFixture.api(GuildRaids.self, cursor == nil ? "guild.raids.history" : "guild.raids.history-2")
    }
    var feedAnswer: () throws(APIError) -> CalendarFeed = { () throws(APIError) in try RaidFixture.api(CalendarFeed.self, "calendar.feed") }
    var resetAnswer: () throws(APIError) -> CalendarFeed = { () throws(APIError) in try RaidFixture.api(CalendarFeed.self, "calendar.feed") }
    private(set) var cursors: [String?] = []
    private(set) var resets = 0

    func home(guild: Int64) async throws(APIError) -> GuildHome { try homeAnswer() }

    func history(guild: Int64, cursor: String?) async throws(APIError) -> GuildRaids {
        cursors.append(cursor)
        return try historyAnswer(cursor)
    }

    func calendarFeed() async throws(APIError) -> CalendarFeed { try feedAnswer() }

    func resetCalendarFeed() async throws(APIError) -> CalendarFeed {
        resets += 1
        return try resetAnswer()
    }
}

/// The calendar link's secure store, in memory.
final class MemoryCalendarLinkStore: CalendarLinkStore, @unchecked Sendable {
    var saved: [String: String] = [:]
    func load(account: String) -> String? { saved[account] }
    func save(_ url: String, account: String) { saved = [account: url] }
    func clear() { saved = [:] }
}

/// The EventKit stand-in: records every write and answers access as told.
@MainActor
final class FakeCalendarAdapter: CalendarStoreAdapter {
    var accessAnswer: CalendarAccess = .notDetermined
    var grants = true
    private(set) var requests = 0
    private(set) var calendars: Set<String> = []
    private(set) var events: [String: CalendarEventDraft] = [:]
    private(set) var removed: [String] = []
    private var nextID = 0

    func access() -> CalendarAccess { accessAnswer }

    func requestAccess() async -> Bool {
        requests += 1
        accessAnswer = grants ? .granted : .denied
        return grants
    }

    func calendarExists(_ id: String) -> Bool { calendars.contains(id) }

    func createCalendar(named name: String) throws -> String {
        nextID += 1
        let id = "cal-\(nextID)"
        calendars.insert(id)
        return id
    }

    func removeCalendar(_ id: String) throws { calendars.remove(id) }

    /// The player deleted the calendar in the Calendar app.
    func deleteCalendarOutsideTheApp() { calendars.removeAll() }

    func save(_ draft: CalendarEventDraft, eventID: String?, calendarID: String) throws -> String {
        let id = eventID ?? "event-\(draft.raidID)"
        events[id] = draft
        return id
    }

    func remove(eventID: String) throws {
        events[eventID] = nil
        removed.append(eventID)
    }
}

final class MemoryCalendarSyncStore: CalendarSyncStore, @unchecked Sendable {
    var record = CalendarSyncRecord()
    func load() -> CalendarSyncRecord { record }
    func save(_ record: CalendarSyncRecord) { self.record = record }
}
