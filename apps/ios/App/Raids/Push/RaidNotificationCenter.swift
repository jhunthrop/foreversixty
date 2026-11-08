import Foundation
import Routing
import UserNotifications

/// The notifications' delegate: registers the categories, answers In, Late, Bench and Out in the
/// background and posts the answer in place, and lands every tap on its answer (PushLanding).
/// Permission and the device token are PushRegistrar's.
@MainActor
final class RaidNotificationCenter: NSObject, UNUserNotificationCenterDelegate {
    static let shared = RaidNotificationCenter()

    private var answerer: RaidPushAnswerer?
    /// The app's router; a tap that arrives before it is set (a cold launch) waits in `pending`.
    var onOpen: ((IntentDestination) -> Void)? {
        didSet {
            if let destination = pending, let onOpen { pending = nil; onOpen(destination) }
        }
    }
    private var pending: IntentDestination?

    /// Installs the delegate and the categories, once at launch.
    func install(answerer: RaidPushAnswerer) {
        self.answerer = answerer
        let center = UNUserNotificationCenter.current()
        center.delegate = self
        center.setNotificationCategories(Self.categories)
    }

    static var categories: Set<UNNotificationCategory> {
        let signIn = UNNotificationAction(identifier: RaidPushCategory.actionIn, title: Copy.pushActionIn, options: [])
        let late = UNNotificationAction(identifier: RaidPushCategory.actionLate, title: Copy.pushActionLate, options: [])
        let bench = UNNotificationAction(identifier: RaidPushCategory.actionBench, title: Copy.pushActionBench, options: [])
        let out = UNNotificationAction(identifier: RaidPushCategory.actionOut, title: Copy.pushActionOut, options: [])
        let askSpec = UNNotificationAction(identifier: RaidPushCategory.actionAskSpec, title: Copy.pushActionAskSpec, options: [.foreground])
        let choose = UNNotificationAction(identifier: RaidPushCategory.actionChoose, title: Copy.pushActionChoose, options: [.foreground])
        return [
            UNNotificationCategory(identifier: RaidPushCategory.signup, actions: [signIn, late, bench, out], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.askSpec, actions: [askSpec, out], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.choose, actions: [choose], intentIdentifiers: []),
        ]
    }

    nonisolated func userNotificationCenter(
        _ center: UNUserNotificationCenter,
        willPresent notification: UNNotification
    ) async -> UNNotificationPresentationOptions {
        [.banner, .list, .sound]
    }

    nonisolated func userNotificationCenter(_ center: UNUserNotificationCenter, didReceive response: UNNotificationResponse) async {
        let request = response.notification.request
        let content = request.content
        let payload = RaidPushPayload(userInfo: content.userInfo)
        let landing = PushLanding.destination(userInfo: content.userInfo)
        await handle(
            payload: payload, landing: landing, action: response.actionIdentifier,
            identifier: request.identifier, thread: content.threadIdentifier
        )
    }

    /// A background answer posts in place; a tap lands on the event's answer (flows spec F4).
    private func handle(payload: RaidPushPayload?, landing: IntentDestination?, action: String, identifier: String, thread: String) async {
        if let payload, let signup = RaidPushCategory.signup(forAction: action), let answerer {
            switch await answerer.answer(payload, signup) {
            case let .answered(confirmation): await confirmation.post(replacing: identifier, thread: thread)
            case let .open(link): open(.link(link))
            }
            return
        }
        if let payload, action == RaidPushCategory.actionAskSpec || action == RaidPushCategory.actionChoose {
            open(.link(payload.deepLink(askingSpec: action == RaidPushCategory.actionAskSpec)))
        } else if let landing {
            open(landing)
        } else if let payload {
            open(.link(payload.deepLink()))
        }
    }

    private func open(_ destination: IntentDestination) {
        if let onOpen { onOpen(destination) } else { pending = destination }
    }
}
