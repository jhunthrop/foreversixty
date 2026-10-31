import Foundation
import Routing
import UserNotifications

/// The raid notifications' delegate: registers the categories, answers In, Late and Out in the
/// background and posts the answer in place, and hands taps to the app as deep links. Lane F7
/// owns permission and token registration; its own delegate work joins this one.
@MainActor
final class RaidNotificationCenter: NSObject, UNUserNotificationCenterDelegate {
    static let shared = RaidNotificationCenter()

    private var answerer: RaidPushAnswerer?
    /// The app's router; links that arrive before it is set wait in `pending`.
    var onOpen: ((DeepLink) -> Void)? {
        didSet {
            if let link = pending, let onOpen { pending = nil; onOpen(link) }
        }
    }
    private var pending: DeepLink?

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
        let out = UNNotificationAction(identifier: RaidPushCategory.actionOut, title: Copy.pushActionOut, options: [])
        let askSpec = UNNotificationAction(identifier: RaidPushCategory.actionAskSpec, title: Copy.pushActionAskSpec, options: [.foreground])
        let choose = UNNotificationAction(identifier: RaidPushCategory.actionChoose, title: Copy.pushActionChoose, options: [.foreground])
        return [
            UNNotificationCategory(identifier: RaidPushCategory.signup, actions: [signIn, late, out], intentIdentifiers: []),
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
        guard let payload = RaidPushPayload(userInfo: content.userInfo) else { return }
        let action = response.actionIdentifier
        let identifier = request.identifier
        let thread = content.threadIdentifier
        await handle(payload: payload, action: action, identifier: identifier, thread: thread)
    }

    private func handle(payload: RaidPushPayload, action: String, identifier: String, thread: String) async {
        if let state = RaidPushCategory.state(forAction: action), let answerer {
            switch await answerer.answer(payload, state: state) {
            case let .answered(answer): await post(answer, replacing: identifier, thread: thread)
            case let .open(link): open(link)
            }
            return
        }
        open(payload.deepLink(askingSpec: action == RaidPushCategory.actionAskSpec))
    }

    private func open(_ link: DeepLink) {
        if let onOpen { onOpen(link) } else { pending = link }
    }

    /// The answer replaces the question: the same identifier, the same thread, no actions.
    private func post(_ answer: RaidPushPayload.Answer, replacing identifier: String, thread: String) async {
        let content = UNMutableNotificationContent()
        content.title = answer.title
        content.body = answer.body
        content.threadIdentifier = thread
        let request = UNNotificationRequest(identifier: identifier, content: content, trigger: nil)
        try? await UNUserNotificationCenter.current().add(request)
    }
}
