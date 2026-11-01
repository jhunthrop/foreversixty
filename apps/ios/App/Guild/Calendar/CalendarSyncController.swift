import Foundation
import Observation

/// Calendar access as the system reports it.
enum CalendarAccess: Equatable, Sendable {
    /// Never asked, or only write access so far (v1's Add to calendar): asking upgrades it.
    case notDetermined
    case granted
    case denied
}

/// The system calendar behind a thin seam: EventKit live, a fake in tests.
@MainActor
protocol CalendarStoreAdapter: AnyObject {
    func access() -> CalendarAccess
    func requestAccess() async -> Bool
    func calendarExists(_ id: String) -> Bool
    func createCalendar(named name: String) throws -> String
    func removeCalendar(_ id: String) throws
    /// Writes the event (new when `eventID` is nil or gone) and returns its id.
    func save(_ draft: CalendarEventDraft, eventID: String?, calendarID: String) throws -> String
    func remove(eventID: String) throws
}

/// What sync keeps between launches: the settings, the app's calendar and the raid-to-event map.
struct CalendarSyncRecord: Codable, Equatable, Sendable {
    var settings = CalendarSettings()
    var calendarID: String?
    var mapping: [Int64: CalendarMappingEntry] = [:]
}

protocol CalendarSyncStore: Sendable {
    func load() -> CalendarSyncRecord
    func save(_ record: CalendarSyncRecord)
}

struct UserDefaultsCalendarSyncStore: CalendarSyncStore {
    static let key = "calendar.sync"
    nonisolated(unsafe) let defaults: UserDefaults

    init(defaults: UserDefaults = .standard) { self.defaults = defaults }

    func load() -> CalendarSyncRecord {
        defaults.data(forKey: Self.key).flatMap { try? JSONDecoder().decode(CalendarSyncRecord.self, from: $0) } ?? CalendarSyncRecord()
    }

    func save(_ record: CalendarSyncRecord) {
        guard let data = try? JSONEncoder().encode(record) else { return }
        defaults.set(data, forKey: Self.key)
    }
}

/// What the Calendar screen's first part shows (spec F16.4, F16.5).
enum CalendarSyncPhase: Equatable, Sendable {
    case off
    case on
    /// Access denied: the access-off card; the switch draws off and disabled.
    case accessOff
    /// The player deleted the calendar: the switch is off with `calendar.removed`.
    case removed
}

/// Keeps the app's calendar, `Forever Sixty raids`, in step with the planned raids. It never
/// asks for access before the player turns the switch on.
@MainActor
@Observable
final class CalendarSyncController {
    private(set) var record: CalendarSyncRecord
    private(set) var phase: CalendarSyncPhase
    /// The last sync failed: `calendar.syncFailed` under the switch.
    private(set) var syncFailed = false

    private let adapter: any CalendarStoreAdapter
    private let store: any CalendarSyncStore
    private let now: @Sendable () -> Date
    private let clock: RaidClock
    private let loadInput: () async -> CalendarInput?
    private var input: CalendarInput?

    init(
        adapter: any CalendarStoreAdapter, store: any CalendarSyncStore, now: @escaping @Sendable () -> Date, clock: RaidClock,
        loadInput: @escaping () async -> CalendarInput?
    ) {
        self.adapter = adapter
        self.store = store
        self.now = now
        self.clock = clock
        self.loadInput = loadInput
        let record = store.load()
        self.record = record
        self.phase = Self.phase(for: record, access: adapter.access(), adapter: adapter)
    }

    private static func phase(for record: CalendarSyncRecord, access: CalendarAccess, adapter: any CalendarStoreAdapter) -> CalendarSyncPhase {
        if access == .denied { return .accessOff }
        guard record.settings.enabled else { return .off }
        if access == .granted, let id = record.calendarID, !adapter.calendarExists(id) { return .removed }
        return .on
    }

    /// Re-reads access when the screen shows (the player may have changed it in Settings).
    func refresh() {
        let access = adapter.access()
        if access == .denied {
            phase = .accessOff
        } else if phase == .accessOff {
            phase = record.settings.enabled ? .on : .off
        }
    }

    func setEnabled(_ on: Bool) async {
        on ? await turnOn() : turnOff()
    }

    private func turnOn() async {
        switch adapter.access() {
        case .denied:
            phase = .accessOff
            return
        case .notDetermined:
            guard await adapter.requestAccess() else { phase = .accessOff; return }
        case .granted:
            break
        }
        do {
            let calendarID = try adapter.createCalendar(named: Copy.calendarName)
            save(CalendarSyncRecord(settings: CalendarSettings(enabled: true, which: record.settings.which, alert: record.settings.alert), calendarID: calendarID))
            phase = .on
            await sync()
        } catch {
            syncFailed = true
        }
    }

    private func turnOff() {
        if let id = record.calendarID { try? adapter.removeCalendar(id) }
        save(CalendarSyncRecord(settings: CalendarSettings(enabled: false, which: record.settings.which, alert: record.settings.alert)))
        phase = adapter.access() == .denied ? .accessOff : .off
        syncFailed = false
    }

    func setWhich(_ which: CalendarWhich) async {
        var settings = record.settings
        settings.which = which
        save(CalendarSyncRecord(settings: settings, calendarID: record.calendarID, mapping: record.mapping))
        await sync()
    }

    func setAlert(_ alert: CalendarAlert) async {
        var settings = record.settings
        settings.alert = alert
        save(CalendarSyncRecord(settings: settings, calendarID: record.calendarID, mapping: record.mapping))
        await sync()
    }

    /// A Raids refresh or a sign-up answer: the new planned raids, synced when the switch is on.
    func update(_ input: CalendarInput) async {
        self.input = input
        await sync()
    }

    private func sync() async {
        guard record.settings.enabled else { return }
        guard adapter.access() == .granted else { phase = .accessOff; return }
        guard let calendarID = record.calendarID, adapter.calendarExists(calendarID) else {
            save(CalendarSyncRecord(settings: CalendarSettings(enabled: false, which: record.settings.which, alert: record.settings.alert)))
            phase = .removed
            return
        }
        if input == nil { input = await loadInput() }
        guard let input else { return }
        let diff = CalendarPlan.diff(input: input, settings: record.settings, mapping: record.mapping, now: now(), clock: clock)
        apply(diff, calendarID: calendarID)
    }

    private func apply(_ diff: CalendarDiff, calendarID: String) {
        var mapping = record.mapping
        var failed = false
        for eventID in diff.deletes {
            do {
                try adapter.remove(eventID: eventID)
                mapping = mapping.filter { $0.value.eventID != eventID }
            } catch {
                failed = true
            }
        }
        for upsert in diff.upserts {
            do {
                let id = try adapter.save(upsert.draft, eventID: upsert.eventID, calendarID: calendarID)
                mapping[upsert.draft.raidID] = CalendarMappingEntry(eventID: id, end: upsert.draft.end)
            } catch {
                failed = true
            }
        }
        syncFailed = failed
        save(CalendarSyncRecord(settings: record.settings, calendarID: calendarID, mapping: mapping))
    }

    private func save(_ record: CalendarSyncRecord) {
        self.record = record
        store.save(record)
    }
}
