import EventKit
import Foundation

/// The system calendar through EventKit with full access (iOS 17+, `NSCalendarsFullAccessUsageDescription`).
/// The app's calendar lives on the local source, or the default calendar's source when there is
/// no local one (an iCloud-only device).
@MainActor
final class EventKitCalendarAdapter: CalendarStoreAdapter {
    private let store = EKEventStore()

    enum Failure: Error { case noSource, notFound }

    func access() -> CalendarAccess {
        switch EKEventStore.authorizationStatus(for: .event) {
        case .fullAccess: .granted
        case .denied, .restricted: .denied
        default: .notDetermined
        }
    }

    func requestAccess() async -> Bool {
        (try? await store.requestFullAccessToEvents()) ?? false
    }

    func calendarExists(_ id: String) -> Bool {
        store.calendar(withIdentifier: id) != nil
    }

    func createCalendar(named name: String) throws -> String {
        guard let source = store.sources.first(where: { $0.sourceType == .local }) ?? store.defaultCalendarForNewEvents?.source else {
            throw Failure.noSource
        }
        let calendar = EKCalendar(for: .event, eventStore: store)
        calendar.title = name
        calendar.source = source
        calendar.cgColor = CGColor(red: 0.898, green: 0.725, blue: 0.333, alpha: 1)
        try store.saveCalendar(calendar, commit: true)
        return calendar.calendarIdentifier
    }

    func removeCalendar(_ id: String) throws {
        guard let calendar = store.calendar(withIdentifier: id) else { return }
        try store.removeCalendar(calendar, commit: true)
    }

    func save(_ draft: CalendarEventDraft, eventID: String?, calendarID: String) throws -> String {
        guard let calendar = store.calendar(withIdentifier: calendarID) else { throw Failure.notFound }
        let event = eventID.flatMap { store.event(withIdentifier: $0) } ?? EKEvent(eventStore: store)
        event.calendar = calendar
        event.title = draft.title
        event.startDate = draft.start
        event.endDate = draft.end
        event.location = draft.location
        event.notes = draft.notes
        event.url = draft.url
        event.alarms = draft.alarmOffset.map { [EKAlarm(relativeOffset: $0)] }
        try store.save(event, span: .thisEvent, commit: true)
        return event.eventIdentifier
    }

    func remove(eventID: String) throws {
        guard let event = store.event(withIdentifier: eventID) else { return }
        try store.remove(event, span: .thisEvent, commit: true)
    }
}
