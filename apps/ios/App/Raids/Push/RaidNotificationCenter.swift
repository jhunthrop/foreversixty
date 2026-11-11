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
    /// Posts an answer in place of its notification (tests record it without the system center).
    private var post: @MainActor (_ answer: RaidPushOutcome.Confirmation, _ reply: PushReply) async -> Void = { answer, reply in
        await answer.post(replacing: reply.identifier, thread: reply.thread, tapThrough: reply.tapThrough)
    }
    /// The app's router; a tap that arrives before it is set (a cold launch) waits in `pending`.
    var onOpen: ((IntentDestination) -> Void)? {
        didSet {
            if let destination = pending, let onOpen { pending = nil; onOpen(destination) }
        }
    }
    private var pending: IntentDestination?

    /// Installs the delegate and the categories, once at launch.
    func install(answerer: RaidPushAnswerer, officer: OfficerPushActions) {
        use(answerer: answerer, officer: officer)
        let center = UNUserNotificationCenter.current()
        center.delegate = self
        center.setNotificationCategories(Self.categories)
    }

    /// The answers the actions run (tests set them, and `post`, without the system center).
    func use(
        answerer: RaidPushAnswerer?, officer: OfficerPushActions,
        post: (@MainActor (RaidPushOutcome.Confirmation, PushReply) async -> Void)? = nil
    ) {
        self.answerer = answerer
        self.officer = officer
        if let post { self.post = post }
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

    /// The completion-handler form, never the `async` one: the `async` form's thunk calls UIKit's
    /// completion from the cooperative pool, and UIKit asserts off the main thread while it takes
    /// the background snapshot (a crash on every action answered from the lock screen, round 4).
    nonisolated func userNotificationCenter(
        _ center: UNUserNotificationCenter,
        didReceive response: UNNotificationResponse,
        withCompletionHandler completionHandler: @escaping () -> Void
    ) {
        let reply = PushReply(response)
        let done = MainThreadCompletion(completionHandler)
        Task { @MainActor in
            await self.respond(reply)
            done.run()
        }
    }

    /// Answers an action or a tap: an officer's in place, a sign-up's in the background, else the landing.
    func respond(_ reply: PushReply) async {
        if let answer = reply.officer {
            await post(await answerOfficer(answer, reply: reply), reply)
            return
        }
        await handle(reply)
    }

    /// Accept or Decline from the applicant push, Mention from the short raid's: done in place, and
    /// what it did (or the failure line) replaces the push; a tap on that still lands (round 5, row 3).
    private func answerOfficer(_ answer: OfficerPushAnswer, reply: PushReply) async -> RaidPushOutcome.Confirmation {
        switch answer {
        case let .accept(guild, application):
            OfficerPushResult.applicant(done: await officer.accept(guild, application), accepted: true, applicant: reply.applicant)
        case let .decline(guild, application):
            OfficerPushResult.applicant(done: await officer.decline(guild, application), accepted: false, applicant: reply.applicant)
        case let .mention(request):
            OfficerPushResult.mention(
                await officer.mention(request.guild, request.raid, request.role), request: request, title: reply.title
            )
        case .unreadable: OfficerPushResult.unreadable(action: reply.action, title: reply.title, applicant: reply.applicant)
        }
    }

    /// A background answer posts in place; a tap lands on the event's answer (flows spec F4).
    private func handle(_ reply: PushReply) async {
        let (payload, landing, action) = (reply.payload, reply.landing, reply.action)
        if let payload, let signup = RaidPushCategory.signup(forAction: action), let answerer {
            switch await answerer.answer(payload, signup) {
            case let .answered(confirmation): await post(confirmation, reply)
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

/// What a notification response carries, read before it crosses to the main actor.
struct PushReply: Sendable {
    let action: String
    let identifier: String
    let thread: String
    /// The delivered push's title (the raid, the applicant), under an officer answer's result.
    let title: String
    let payload: RaidPushPayload?
    let landing: IntentDestination?
    let officer: OfficerPushAnswer?
    /// The push's `kind`, `link` and `data` strings, kept on the answer that replaces it so a tap on
    /// that still lands where the push did (never its actions).
    let tapThrough: [String: PushTapValue]
    /// Who an applicant push is about, for its answer's words.
    let applicant: ApplicantPushName

    init(action: String, identifier: String, thread: String, title: String = "", userInfo: [AnyHashable: Any]) {
        self.action = action
        self.identifier = identifier
        self.thread = thread
        self.title = title
        payload = RaidPushPayload(userInfo: userInfo)
        landing = PushLanding.destination(userInfo: userInfo)
        officer = OfficerPushAnswer(action: action, userInfo: userInfo)
        tapThrough = PushTapValue.tapThrough(userInfo)
        applicant = ApplicantPushName(data: PushLanding.data(userInfo: userInfo), pushTitle: title)
    }

    init(_ response: UNNotificationResponse) {
        let content = response.notification.request.content
        self.init(
            action: response.actionIdentifier, identifier: response.notification.request.identifier,
            thread: content.threadIdentifier, title: content.title, userInfo: content.userInfo
        )
    }
}
