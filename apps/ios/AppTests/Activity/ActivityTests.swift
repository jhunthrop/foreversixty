@testable import ForeverSixty
import Foundation
import Models
import Networking
import SwiftUI
import Testing

/// The mock's Activity (design/mocks/ws_app_flows.py ACTIVITY), at the parity clock in Chicago.
enum ActivityTestData {
    static let zone = TimeZone(identifier: "America/Chicago")!
    static var calendar: Calendar {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = zone
        return calendar
    }

    /// Fri 6 Nov 2026, 12:00 pm CST.
    static let now = Date(timeIntervalSince1970: 1_793_988_000)

    static func row(
        _ id: Int64, _ kind: String, _ group: String, _ title: String, _ body: String, icon: String, hoursAgo: Double, read: Bool,
        figure: Double? = nil
    ) -> ActivityRow {
        ActivityRow(
            id: id, kind: kind, group: group, title: title, body: body, link: "https://foreversixty.gg/",
            occurredAt: now.addingTimeInterval(-hoursAgo * 3600), read: read, icon: icon, figure: figure
        )
    }

    static let onyxia = "inv_misc_head_dragon_black"
    static let book = "inv_misc_book_09"

    static let rows = [
        row(8, "guildsim.passed", "ladder", "Kaelra passed you in the guild sims", "You are #7 of 14", icon: "ability_warrior_innerrage", hoursAgo: 8, read: false, figure: 1196.2),
        row(7, "loot.awarded", "gear", "You got Deathbringer", "Onyxia · main hand, +12.6 DPS", icon: "inv_axe_09", hoursAgo: 9, read: false),
        row(6, "report.ready", "logs", "Onyxia’s Lair report ready", "87th percentile, your best on Onyxia so far · the coach has 2 notes", icon: onyxia, hoursAgo: 9.2, read: false),
        row(5, "raid.signup_nudge", "raids", "You’re In for Onyxia", "As Thandrel, Damage · 8:00 pm CST", icon: onyxia, hoursAgo: 18, read: true),
        row(4, "sim.finished", "gear", "Sim done: 949 DPS ±6", "+21 vs your last", icon: "inv_misc_pocketwatch_01", hoursAgo: 21.5, read: true),
        row(3, "raid.short", "officer", "Healers short for Onyxia tonight", "3 of 8 missing", icon: onyxia, hoursAgo: 23, read: true),
        row(2, "tier.moved", "tiers", "Fury Warrior moved to S", "Raid tiers, 60", icon: book, hoursAgo: 72, read: true),
        row(1, "bis.changed", "gear", "Fury Warrior BiS changed", "A new pick for Hands: Theramore Gauntlets", icon: book, hoursAgo: 74, read: true),
    ]
}

final class FakeActivityService: ActivityService, @unchecked Sendable {
    var pages: [Int64?: ActivityPage] = [nil: ActivityPage(rows: ActivityTestData.rows, unread: 3, nextBefore: nil)]
    var readThrough: [Int64] = []
    var readFails = false

    func page(group: String?, before: Int64?) async throws(APIError) -> ActivityPage {
        guard let page = pages[before] else { throw .transport("offline") }
        return page
    }

    func markRead(through: Int64) async throws(APIError) -> Int {
        if readFails { throw .transport("offline") }
        readThrough.append(through)
        return 0
    }
}

@Suite("Activity")
struct ActivitySectionsTests {
    private func sections(_ filter: ActivityFilter = .all) -> ActivitySections {
        ActivitySections(
            ActivityTestData.rows, filter: filter, now: ActivityTestData.now, calendar: ActivityTestData.calendar,
            locale: Locale(identifier: "en_US")
        )
    }

    @Test("rows fall into Today, Yesterday and This week, newest first, each with its time")
    func days() {
        let sections = sections()
        #expect(sections.sections.map(\.day) == [.today, .yesterday, .thisWeek])
        #expect(sections.sections[0].rows.map(\.row.id) == [8, 7, 6])
        #expect(sections.sections[0].rows[0].when == "8 h ago")
        #expect(sections.sections[1].rows[0].when == "6:00 pm")
        #expect(sections.sections[2].rows[0].when == "Tue")
    }

    @Test("a filter keeps its group; Raids keeps the officer's rows too")
    func filters() {
        #expect(sections(.gear).sections.flatMap(\.rows).map(\.row.id) == [7, 4, 1])
        #expect(sections(.raids).sections.flatMap(\.rows).map(\.row.id) == [5, 3])
        #expect(sections(.ladder).sections.flatMap(\.rows).map(\.row.id) == [8])
    }
}

@MainActor
@Suite("Activity model")
struct ActivityModelTests {
    @Test("loading reads the rows and hands the unread count to the badge")
    func load() async {
        var badge = -1
        let model = ActivityModel(service: FakeActivityService())
        model.onUnread = { badge = $0 }
        await model.load()
        #expect(model.phase == .loaded)
        #expect(model.rows.count == 8)
        #expect(badge == 3)
    }

    @Test("Mark all read reads through the newest row at once; a failure puts the dots back")
    func markAllRead() async {
        let service = FakeActivityService()
        let model = ActivityModel(service: service)
        await model.load()
        await model.markAllRead()
        #expect(service.readThrough == [8])
        #expect(model.unread == 0)
        #expect(model.rows.allSatisfy { $0.read })

        let failing = FakeActivityService()
        failing.readFails = true
        let other = ActivityModel(service: failing)
        await other.load()
        await other.markAllRead()
        #expect(other.unread == 3)
        #expect(other.rows.filter { !$0.read }.count == 3)
    }

    @Test("the next page joins the list without repeats")
    func more() async {
        let service = FakeActivityService()
        service.pages[nil] = ActivityPage(rows: Array(ActivityTestData.rows.prefix(4)), unread: 3, nextBefore: 5)
        service.pages[5] = ActivityPage(rows: Array(ActivityTestData.rows.dropFirst(3)), unread: 3, nextBefore: nil)
        let model = ActivityModel(service: service)
        await model.load()
        #expect(model.hasOlder, "i1: Show older is offered")
        await model.loadMore()
        #expect(model.rows.map(\.id) == [8, 7, 6, 5, 4, 3, 2, 1])
        #expect(model.nextBefore == nil)
        #expect(!model.hasOlder, "the last page hides Show older")
        #expect(!model.loadingMore)
    }

    @Test("round 5, i1: the shell reads the Activity badge before the tab opens; signed out or offline, none")
    func badge() async {
        let service = FakeActivityService()
        #expect(await ActivityBadgeCount.load(signedIn: true, service: service) == 3)
        #expect(await ActivityBadgeCount.load(signedIn: false, service: service) == 0)
        service.pages = [:]
        #expect(await ActivityBadgeCount.load(signedIn: true, service: service) == 0)
    }
}

@MainActor
@Suite("Activity parity")
struct ActivityParityTests {
    static let device = CGSize(width: 393, height: 852)

    private func view(_ rows: [ActivityRow]) -> some View {
        ActivityView(
            phase: .loaded,
            sections: ActivitySections(rows, filter: .all, now: ActivityTestData.now, calendar: ActivityTestData.calendar),
            filter: .all, build: "1.60.1.70790", older: rows.isEmpty ? nil : .ready
        )
        .environment(\.itemIconSource, RepoItemIconSource())
    }

    @Test("ios-activity")
    func list() async throws {
        try await ParityCheck.assert(view(ActivityTestData.rows), screen: "activity", size: Self.device)
    }

    @Test("ios-activity-empty")
    func empty() async throws {
        try await ParityCheck.assert(view([]), screen: "activity-empty", size: Self.device)
    }

    @Test("ios-activity-out")
    func signedOut() async throws {
        try await ParityCheck.assert(ActivitySignedOutView(signingIn: false) {}, screen: "activity-out", size: Self.device)
    }
}
