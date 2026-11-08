import UIKit
import UserNotifications
import UserNotificationsUI

/// The raid notification's content extension (flows spec F1, mock ios-notifications): the system
/// draws the notification; after Late this adds `How late?` and swaps the buttons for the 15, 30
/// and 60 min chips. Late is saved the moment it is pressed, so closing without a chip leaves a
/// plain Late; a chip saves its minutes and posts the server's confirmation. In, Bench and Out
/// go on to the app, which answers in the background.
final class NotificationViewController: UIViewController, @MainActor UNNotificationContentExtension {
    private static let questionHeight: CGFloat = 36
    private static let inset: CGFloat = 16

    private let question = UILabel()
    private var payload: RaidPushPayload?
    private var lateSaved: Task<RaidPushOutcome?, Never>?

    override func viewDidLoad() {
        super.viewDidLoad()
        question.text = Copy.pushLateQuestion
        question.font = UIFont.preferredFont(forTextStyle: .subheadline)
        question.adjustsFontForContentSizeCategory = true
        question.isHidden = true
        question.translatesAutoresizingMaskIntoConstraints = false
        view.addSubview(question)
        NSLayoutConstraint.activate([
            question.leadingAnchor.constraint(equalTo: view.leadingAnchor, constant: Self.inset),
            question.trailingAnchor.constraint(equalTo: view.trailingAnchor, constant: -Self.inset),
            question.centerYAnchor.constraint(equalTo: view.centerYAnchor),
        ])
        preferredContentSize = CGSize(width: view.bounds.width, height: 0)
    }

    func didReceive(_ notification: UNNotification) {
        payload = RaidPushPayload(userInfo: notification.request.content.userInfo)
    }

    func didReceive(
        _ response: UNNotificationResponse,
        completionHandler completion: @escaping (UNNotificationContentExtensionResponseOption) -> Void
    ) {
        let action = response.actionIdentifier
        guard let payload, let signup = RaidPushCategory.signup(forAction: action) else {
            completion(.dismissAndForwardAction)
            return
        }
        if action == RaidPushCategory.actionLate {
            showChips()
            let answerer = Self.answerer()
            lateSaved = Task { await answerer?.answer(payload, signup) }
            completion(.doNotDismiss)
            return
        }
        guard signup.lateMinutes != nil, let answerer = Self.answerer() else {
            completion(.dismissAndForwardAction)
            return
        }
        let identifier = response.notification.request.identifier
        let thread = response.notification.request.content.threadIdentifier
        let lateSaved = lateSaved
        Task {
            // The plain Late goes first, so the minutes are the last word.
            _ = await lateSaved?.value
            if case let .answered(confirmation) = await answerer.answer(payload, signup) {
                await confirmation.post(replacing: identifier, thread: thread)
            }
            completion(.dismiss)
        }
    }

    private func showChips() {
        question.isHidden = false
        preferredContentSize = CGSize(width: view.bounds.width, height: Self.questionHeight)
        extensionContext?.notificationActions = RaidPushCategory.lateMinutes.map { minutes in
            UNNotificationAction(identifier: RaidPushCategory.actionLate(minutes: minutes), title: Copy.pushLateMinutes(n: minutes), options: [])
        }
    }

    private static func answerer() -> RaidPushAnswerer? {
        LiveIntentCalls.fromBundle().map { RaidPushAnswerer(calls: $0) }
    }
}
