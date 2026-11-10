import Auth
import Foundation
import Models
import Networking
import Observation

/// `GET /v1/me/activity` and `POST /v1/me/activity/read`, behind a protocol for tests.
protocol ActivityService: Sendable {
    func page(group: String?, before: Int64?) async throws(APIError) -> ActivityPage
    func markRead(through: Int64) async throws(APIError) -> Int
}

struct LiveActivityService: ActivityService {
    let client: APIClient
    /// Rows per page.
    static let pageSize = 50

    func page(group: String?, before: Int64?) async throws(APIError) -> ActivityPage {
        try await client.send(.getActivity(group: group, before: before, limit: Self.pageSize))
    }

    func markRead(through: Int64) async throws(APIError) -> Int {
        try await client.send(.markActivityRead(ActivityReadRequest(through: through))).unread
    }
}

/// Activity's filters (flows spec 4.4): All, then the groups a player reads. Officer rows sit under
/// Raids for an officer.
enum ActivityFilter: String, CaseIterable, Identifiable {
    case all, raids, gear, logs, ladder, tiers

    var id: String { rawValue }

    var title: String {
        switch self {
        case .all: Copy.activityFilterAll
        case .raids: Copy.pushGroupRaids
        case .gear: Copy.pushGroupGear
        case .logs: Copy.pushGroupLogs
        case .ladder: Copy.pushGroupLadder
        case .tiers: Copy.pushGroupTiers
        }
    }

    /// The groups the filter shows.
    func admits(_ group: String) -> Bool {
        switch self {
        case .all: true
        case .raids: group == NotificationGroup.raids.rawValue || group == NotificationGroup.officer.rawValue
        default: group == rawValue
        }
    }
}

/// The day bands rows fall in, newest first.
enum ActivityDay: Int, CaseIterable, Comparable {
    case today, yesterday, thisWeek, earlier

    var title: String {
        switch self {
        case .today: Copy.activityToday
        case .yesterday: Copy.activityYesterday
        case .thisWeek: Copy.activityThisWeek
        case .earlier: Copy.activityEarlier
        }
    }

    static func < (lhs: ActivityDay, rhs: ActivityDay) -> Bool { lhs.rawValue < rhs.rawValue }

    static func of(_ date: Date, now: Date, calendar: Calendar) -> ActivityDay {
        if calendar.isDate(date, inSameDayAs: now) { return .today }
        if let yesterday = calendar.date(byAdding: .day, value: -1, to: now), calendar.isDate(date, inSameDayAs: yesterday) {
            return .yesterday
        }
        let days = calendar.dateComponents([.day], from: calendar.startOfDay(for: date), to: calendar.startOfDay(for: now)).day ?? 0
        return days < 7 ? .thisWeek : .earlier
    }
}

/// The list as drawn: rows grouped by day, each with its time in the band's words.
struct ActivitySections: Equatable {
    struct Row: Equatable, Identifiable {
        let row: ActivityRow
        let when: String

        var id: Int64 { row.id }
    }

    struct Section: Equatable, Identifiable {
        let day: ActivityDay
        let rows: [Row]

        var id: ActivityDay { day }
    }

    let sections: [Section]

    init(_ rows: [ActivityRow], filter: ActivityFilter, now: Date, calendar: Calendar = .current, locale: Locale = .current) {
        let shown = rows.filter { filter.admits($0.group) }.sorted { $0.occurredAt > $1.occurredAt }
        let byDay = Dictionary(grouping: shown) { ActivityDay.of($0.occurredAt, now: now, calendar: calendar) }
        sections = ActivityDay.allCases.compactMap { day in
            guard let rows = byDay[day] else { return nil }
            return Section(day: day, rows: rows.map { Row(row: $0, when: Self.when($0.occurredAt, day: day, now: now, calendar: calendar, locale: locale)) })
        }
    }

    /// Today `8 h ago`; yesterday the clock, `6:02 pm`; this week the weekday, `Tue`; earlier the date.
    static func when(_ date: Date, day: ActivityDay, now: Date, calendar: Calendar, locale: Locale) -> String {
        let style = Date.FormatStyle(locale: locale, calendar: calendar, timeZone: calendar.timeZone)
        switch day {
        case .today: return RelativeAge(from: date, to: now).phrase
        case .yesterday: return ClockText.time(date, calendar: calendar, locale: locale)
        case .thisWeek: return date.formatted(style.weekday(.abbreviated))
        case .earlier: return date.formatted(style.day().month(.abbreviated))
        }
    }
}

@MainActor
@Observable
final class ActivityModel {
    enum Phase: Equatable {
        case loading
        case loaded
        case failed
    }

    private(set) var phase: Phase = .loading
    private(set) var rows: [ActivityRow] = []
    private(set) var unread = 0
    private(set) var nextBefore: Int64?
    var filter: ActivityFilter = .all
    /// Hears the unread count, for the tab's badge.
    var onUnread: (Int) -> Void = { _ in }
    private let service: any ActivityService

    init(service: any ActivityService) {
        self.service = service
    }

    func load() async {
        do {
            let page = try await service.page(group: nil, before: nil)
            rows = page.rows
            nextBefore = page.nextBefore
            setUnread(page.unread)
            phase = .loaded
        } catch {
            if phase != .loaded { phase = .failed }
        }
    }

    /// The next page is out (Show older reads `…` and waits).
    private(set) var loadingMore = false

    /// Older rows exist past the loaded ones.
    var hasOlder: Bool { nextBefore != nil }

    /// The next page, from Show older (i1).
    func loadMore() async {
        guard let before = nextBefore, !loadingMore else { return }
        loadingMore = true
        defer { loadingMore = false }
        guard let page = try? await service.page(group: nil, before: before) else { return }
        rows += page.rows.filter { new in !rows.contains { $0.id == new.id } }
        nextBefore = page.nextBefore
        setUnread(page.unread)
    }

    /// Mark all read: every row through the newest, drawn at once.
    func markAllRead() async {
        guard let newest = rows.map(\.id).max() else { return }
        let before = (rows, unread)
        rows = rows.map(\.asRead)
        setUnread(0)
        do {
            setUnread(try await service.markRead(through: newest))
        } catch {
            rows = before.0
            setUnread(before.1)
        }
    }

    private func setUnread(_ count: Int) {
        unread = count
        onUnread(count)
    }
}

extension ActivityRow {
    var asRead: ActivityRow {
        ActivityRow(
            id: id, kind: kind, group: group, title: title, body: body, link: link, occurredAt: occurredAt, read: true, icon: icon,
            segments: segments, crest: crest, figure: figure
        )
    }
}
