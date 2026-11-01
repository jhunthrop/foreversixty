import Foundation
import Models

/// `Which raids` (spec F16.4).
enum CalendarWhich: String, Codable, CaseIterable, Identifiable, Sendable {
    case all, mine

    var id: String { rawValue }

    var label: String {
        switch self {
        case .all: Copy.calendarWhichAll
        case .mine: Copy.calendarWhichMine
        }
    }
}

/// `Alert` (spec F16.4). The event starts at invites, so "1 hour before invites" is an hour
/// before the event's start.
enum CalendarAlert: String, Codable, CaseIterable, Identifiable, Sendable {
    case none, minutes15, hourBeforeInvites, day

    var id: String { rawValue }

    var label: String {
        switch self {
        case .none: Copy.calendarAlertNone
        case .minutes15: Copy.calendarAlertMinutes15
        case .hourBeforeInvites: Copy.calendarAlertHourBeforeInvites
        case .day: Copy.calendarAlertDay
        }
    }

    /// Seconds relative to the event's start; nil is no alarm.
    var offset: TimeInterval? {
        switch self {
        case .none: nil
        case .minutes15: -15 * 60
        case .hourBeforeInvites: -60 * 60
        case .day: -24 * 60 * 60
        }
    }
}

struct CalendarSettings: Codable, Equatable, Sendable {
    var enabled = false
    var which: CalendarWhich = .all
    var alert: CalendarAlert = .hourBeforeInvites
}

/// The app's record of an event it wrote: the event's id and the raid's end, so a past event stays.
struct CalendarMappingEntry: Codable, Equatable, Sendable {
    let eventID: String
    let end: Date
}

/// What the calendar is built from: the guild's planned raids as the Raids tab loaded them.
struct CalendarInput: Equatable, Sendable {
    let guildID: Int64
    let guildName: String
    let raids: [RaidView]
}

/// One event the calendar should hold for a raid.
struct CalendarEventDraft: Equatable, Sendable {
    let raidID: Int64
    let title: String
    let start: Date
    let end: Date
    let location: String
    let notes: String
    let url: URL
    let alarmOffset: TimeInterval?
}

struct CalendarUpsert: Equatable, Sendable {
    let draft: CalendarEventDraft
    /// The event already written for this raid, else nil (a new event).
    let eventID: String?
}

struct CalendarDiff: Equatable, Sendable {
    let upserts: [CalendarUpsert]
    /// Event ids to delete.
    let deletes: [String]
}

/// The pure part of calendar sync (spec F16.4): planned raids, settings and the existing
/// mapping in; the writes and deletes out.
enum CalendarPlan {
    /// The viewer's best answer, In over Late over Bench over Out.
    private static let order: [SignupState] = [.signedIn, .late, .bench, .out]

    static func best(of states: [String]) -> SignupState? {
        let given = Set(states.compactMap(SignupState.init(rawValue:)))
        return order.first { given.contains($0) }
    }

    static func state(of states: [String]) -> String {
        switch best(of: states) {
        case .signedIn: Copy.calendarEventIn
        case .late: Copy.calendarEventLate
        case .bench: Copy.calendarEventBench
        case .out: Copy.calendarEventOut
        case nil: Copy.calendarEventNotSigned
        }
    }

    static func draft(_ raid: RaidView, input: CalendarInput, alert: CalendarAlert, clock: RaidClock) -> CalendarEventDraft {
        let facts = RaidFacts(raid: raid, now: raid.startsAt)
        let states = raid.viewer.characters.compactMap(\.state)
        let pull = clock.clock(raid.startsAt)
        let line = raid.size.map { Copy.calendarEventNotes(time: pull, n: facts.signedUp, size: $0) }
            ?? Copy.calendarEventNotesNoSize(time: pull, n: facts.signedUp)
        let url = RaidLinks.web(guild: input.guildID, raid: raid.id)
        return CalendarEventDraft(
            raidID: raid.id,
            title: Copy.calendarEventTitle(zone: raid.zoneName, state: state(of: states)),
            start: raid.inviteAt ?? raid.startsAt,
            end: facts.endsAt,
            location: input.guildName.isEmpty ? raid.zoneName : Copy.calendarEventLocation(guild: input.guildName, zone: raid.zoneName),
            notes: "\(line)\n\(url.absoluteString)",
            url: url,
            alarmOffset: alert.offset
        )
    }

    static func diff(
        input: CalendarInput, settings: CalendarSettings, mapping: [Int64: CalendarMappingEntry], now: Date, clock: RaidClock
    ) -> CalendarDiff {
        let live = input.raids.filter { $0.status != "cancelled" && RaidFacts(raid: $0, now: now).endsAt > now }
        let wanted = live.filter { raid in
            settings.which == .all || best(of: raid.viewer.characters.compactMap(\.state)).map(\.isComing) == true
        }
        let wantedIDs = Set(wanted.map(\.id))
        let upserts = wanted.map { raid in
            CalendarUpsert(draft: draft(raid, input: input, alert: settings.alert, clock: clock), eventID: mapping[raid.id]?.eventID)
        }
        let deletes = mapping.compactMap { raidID, entry -> String? in
            guard !wantedIDs.contains(raidID) else { return nil }
            let cancelled = input.raids.contains { $0.id == raidID && $0.status == "cancelled" }
            // A past event stays; a cancelled, vanished or no-longer-wanted future one goes.
            if cancelled { return entry.eventID }
            return entry.end > now ? entry.eventID : nil
        }
        return CalendarDiff(upserts: upserts, deletes: deletes.sorted())
    }
}
