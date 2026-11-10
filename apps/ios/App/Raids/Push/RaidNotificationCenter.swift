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
    /// Declines an application from the applicant push; true when the API took it.
    private var decline: (@Sendable (_ guild: Int64, _ application: Int64) async -> Bool)?
    /// The app's router; a tap that arrives before it is set (a cold launch) waits in `pending`.
    var onOpen: ((IntentDestination) -> Void)? {
        didSet {
            if let destination = pending, let onOpen { pending = nil; onOpen(destination) }
        }
    }
    private var pending: IntentDestination?

    /// Installs the delegate and the categories, once at launch.
    func install(answerer: RaidPushAnswerer, decline: @escaping @Sendable (_ guild: Int64, _ application: Int64) async -> Bool) {
        self.answerer = answerer
        self.decline = decline
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
        let declineAction = UNNotificationAction(
            identifier: RaidPushCategory.actionDecline, title: Copy.pushActionDecline, options: [.destructive, .authenticationRequired]
        )
        return [
            UNNotificationCategory(identifier: RaidPushCategory.signup, actions: [signIn, late, bench, out], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.askSpec, actions: [askSpec, out], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.choose, actions: [choose], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.applicant, actions: [declineAction], intentIdentifiers: []),
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
        if response.actionIdentifier == RaidPushCategory.actionDecline {
            await declineApplicant(data: PushLanding.data(userInfo: content.userInfo), landing: landing)
            return
        }
        await handle(
            payload: payload, landing: landing, action: response.actionIdentifier,
            identifier: request.identifier, thread: content.threadIdentifier
        )
    }

    /// Decline from the applicant push; a refused or failed decline opens the review instead.
    private func declineApplicant(data: [String: String], landing: IntentDestination?) async {
        if let application = PushLanding.applicant(data: data), let decline,
           await decline(application.guild, application.application) {
            return
        }
        if let landing { open(landing) }
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
