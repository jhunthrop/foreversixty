import AppIntents
import Foundation
import Observation
import Routing
import SwiftUI

/// Where an intent or a Spotlight tap sends the app; RootView reads `pending` and opens it the
/// way a deep link opens (spec §5.7: a tap opens the same screen as in-app).
enum IntentDestination: Equatable, Sendable {
    case link(DeepLink)
    /// A character: F11.4 when `rankings` is on, else Home on that character.
    case character(key: String)
    /// Rankings on Home, on its default board.
    case rankingsBoard
    case tab(AppTab)
}

@MainActor
@Observable
final class IntentRouter {
    static let shared = IntentRouter()

    var pending: IntentDestination?

    func open(_ destination: IntentDestination) { pending = destination }
}

/// The answers' shared plumbing: the live calls, the phone's zone and the `system` switch.
enum IntentRuntime {
    static var answers: IntentAnswers { IntentAnswers(clock: WidgetClock(zone: .current)) }

    static func queries() -> IntentQueries? {
        LiveIntentCalls.fromBundle().map { IntentQueries(calls: $0, answers: answers, enabled: { SystemSwitch.isOn() }) }
    }
}

struct NextRaidIntent: AppIntent {
    static let title: LocalizedStringResource = "intent.nextRaid.title"

    @Parameter(title: "intent.character") var character: CharacterEntity?

    func perform() async throws -> some IntentResult & ProvidesDialog {
        let text = await IntentRuntime.queries()?.nextRaid(character: character?.characterKey) ?? Copy.commonFailed
        return .result(dialog: IntentDialog(stringLiteral: text))
    }
}

struct WhoIsComingIntent: AppIntent {
    static let title: LocalizedStringResource = "intent.whoIsComing.title"

    @Parameter(title: "intent.character") var character: CharacterEntity?

    func perform() async throws -> some IntentResult & ProvidesDialog {
        let text = await IntentRuntime.queries()?.whoIsComing(character: character?.characterKey) ?? Copy.commonFailed
        return .result(dialog: IntentDialog(stringLiteral: text))
    }
}

/// `What's my next upgrade`: a snippet with the item icon in its quality border, the name in its
/// quality colour, the meta line and the gain (board F15.2).
struct NextUpgradeIntent: AppIntent {
    static let title: LocalizedStringResource = "intent.nextUpgrade.title"

    @Parameter(title: "intent.character") var character: CharacterEntity?

    func perform() async throws -> some IntentResult & ProvidesDialog & ShowsSnippetView {
        let reply = await IntentRuntime.queries()?.nextUpgrade(character: character?.characterKey) ?? .sentence(Copy.commonFailed)
        let dialog = IntentDialog(stringLiteral: reply.sentence)
        guard case let .upgrade(answer, build) = reply else {
            return .result(dialog: dialog, view: AnyView(EmptyView()))
        }
        let image = await answer.icon.asyncFlatMap { await RemoteItemIconSource.shared.load(icon: $0, build: build) }
        return .result(dialog: dialog, view: AnyView(UpgradeSnippetView(answer: answer, image: image)))
    }
}

/// `Open my best in slot list`: the BiS tab.
struct OpenBisIntent: AppIntent {
    static let title: LocalizedStringResource = "intent.openBis.title"
    static let supportedModes: IntentModes = .foreground(.immediate)

    @MainActor
    func perform() async throws -> some IntentResult {
        IntentRouter.shared.open(.tab(.bis))
        return .result()
    }
}

/// Every intent with its phrases, so they work with no setup (spec §5.7). Each phrase names the
/// app; `Tell Forever Sixty I'm late` names it in place rather than repeating it at the end.
struct ForeverShortcuts: AppShortcutsProvider {
    static var appShortcuts: [AppShortcut] {
        AppShortcut(
            intent: SignUpForRaidIntent(state: .in, askFirst: true),
            phrases: ["Sign me up for tonight's raid in \(.applicationName)"],
            shortTitle: "intent.state.in", systemImageName: "checkmark.circle"
        )
        AppShortcut(
            intent: SignUpForRaidIntent(state: .late, askFirst: true),
            phrases: ["Tell \(.applicationName) I'm late"],
            shortTitle: "intent.state.late", systemImageName: "clock"
        )
        AppShortcut(
            intent: SignUpForRaidIntent(state: .out, askFirst: true),
            phrases: ["Mark me out of the next raid in \(.applicationName)"],
            shortTitle: "intent.state.out", systemImageName: "xmark.circle"
        )
        AppShortcut(
            intent: NextUpgradeIntent(), phrases: ["What's my next upgrade in \(.applicationName)"],
            shortTitle: "intent.nextUpgrade.title", systemImageName: "arrow.up.circle"
        )
        AppShortcut(
            intent: NextRaidIntent(), phrases: ["When is the next raid in \(.applicationName)"],
            shortTitle: "intent.nextRaid.title", systemImageName: "calendar"
        )
        AppShortcut(
            intent: WhoIsComingIntent(), phrases: ["Who is coming to raid in \(.applicationName)"],
            shortTitle: "intent.whoIsComing.title", systemImageName: "person.3"
        )
        AppShortcut(
            intent: OpenBisIntent(), phrases: ["Open my best in slot list in \(.applicationName)"],
            shortTitle: "intent.openBis.title", systemImageName: "shield"
        )
    }
}
