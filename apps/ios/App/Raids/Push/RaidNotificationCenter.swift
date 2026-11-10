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
    /// Accept, Decline and Mention from the officer pushes.
    private var officer = OfficerPushActions()
    /// The app's router; a tap that arrives before it is set (a cold launch) waits in `pending`.
    var onOpen: ((IntentDestination) -> Void)? {
        didSet {
            if let destination = pending, let onOpen { pending = nil; onOpen(destination) }
        }
    }
    private var pending: IntentDestination?

    /// Installs the delegate and the categories, once at launch.
    func install(answerer: RaidPushAnswerer, officer: OfficerPushActions) {
        self.answerer = answerer
        self.officer = officer
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
        let acceptAction = UNNotificationAction(
            identifier: RaidPushCategory.actionAccept, title: Copy.pushActionAccept, options: [.authenticationRequired]
        )
        let declineAction = UNNotificationAction(
            identifier: RaidPushCategory.actionDecline, title: Copy.pushActionDecline, options: [.destructive, .authenticationRequired]
        )
        // A category's titles are fixed at launch, so the role the server names (`Mention healers`)
        // reads in the body; the action says what it does.
        let mention = UNNotificationAction(
            identifier: RaidPushCategory.actionMention, title: Copy.pushActionMention, options: [.authenticationRequired]
        )
        return [
            UNNotificationCategory(identifier: RaidPushCategory.signup, actions: [signIn, late, bench, out], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.askSpec, actions: [askSpec, out], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.choose, actions: [choose], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.applicant, actions: [acceptAction, declineAction], intentIdentifiers: []),
            UNNotificationCategory(identifier: RaidPushCategory.short, actions: [mention], intentIdentifiers: []),
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
        if let answer = OfficerPushAnswer(action: response.actionIdentifier, userInfo: content.userInfo) {
            await answerOfficer(answer, landing: landing)
            return
        }
        await handle(
            payload: payload, landing: landing, action: response.actionIdentifier,
            identifier: request.identifier, thread: content.threadIdentifier
        )
    }

    /// Accept or Decline from the applicant push, Mention from the short raid's: done in place; a
    /// refused, failed or unreadable answer opens the push's landing instead.
    private func answerOfficer(_ answer: OfficerPushAnswer, landing: IntentDestination?) async {
        let done = switch answer {
        case let .accept(guild, application): await officer.accept(guild, application)
        case let .decline(guild, application): await officer.decline(guild, application)
        case let .mention(request): await officer.mention(request.guild, request.raid, request.role)
        case .unreadable: false
        }
        if !done, let landing { open(landing) }
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
