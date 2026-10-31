import EventKit
import EventKitUI
import Models
import SwiftUI

/// The calendar event for a raid (spec F4.5): title, start, duration, `{Guild} · {Zone}`, the link.
struct CalendarEvent: Equatable {
    let title: String
    let start: Date
    let end: Date
    let location: String
    let url: URL

    init(raid: RaidView, guildName: String, guild: Int64) {
        title = raid.title
        start = raid.startsAt
        end = raid.startsAt.addingTimeInterval(TimeInterval(raid.durationMinutes * 60))
        location = guildName.isEmpty ? raid.zoneName : Copy.raidsCalendarLocation(guild: guildName, zone: raid.zoneName)
        url = RaidLinks.web(guild: guild, raid: raid.id)
    }
}

/// `EKEventEditViewController` with write-only access (iOS 17+): the system asks on first use.
struct CalendarSheet: UIViewControllerRepresentable {
    let event: CalendarEvent

    func makeUIViewController(context: Context) -> EKEventEditViewController {
        let store = EKEventStore()
        let controller = EKEventEditViewController()
        controller.eventStore = store
        let draft = EKEvent(eventStore: store)
        draft.title = event.title
        draft.startDate = event.start
        draft.endDate = event.end
        draft.location = event.location
        // The link has its own URL field here; repeating it in notes showed it twice.
        // (Android's insert intent has no URL field, so its description carries the link.)
        draft.url = event.url
        controller.event = draft
        controller.editViewDelegate = context.coordinator
        return controller
    }

    func updateUIViewController(_ controller: EKEventEditViewController, context: Context) {}

    func makeCoordinator() -> Coordinator { Coordinator() }

    @MainActor
    final class Coordinator: NSObject, @MainActor EKEventEditViewDelegate {
        func eventEditViewController(_ controller: EKEventEditViewController, didCompleteWith action: EKEventEditViewAction) {
            controller.dismiss(animated: true)
        }
    }
}
