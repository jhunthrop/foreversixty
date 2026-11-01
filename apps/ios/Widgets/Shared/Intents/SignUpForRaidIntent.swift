import AppIntents
import Foundation
import WidgetKit

/// The three answers a player gives (Bench is an officer's move, never spoken).
enum RaidAnswerState: String, AppEnum {
    case `in`, late, out

    // App Intents metadata takes literal keys; each is a key of the copy catalog (Localizable.xcstrings).
    static let typeDisplayRepresentation = TypeDisplayRepresentation(name: "intent.signup.title")
    static let caseDisplayRepresentations: [RaidAnswerState: DisplayRepresentation] = [
        .in: DisplayRepresentation(title: "intent.state.in"),
        .late: DisplayRepresentation(title: "intent.state.late"),
        .out: DisplayRepresentation(title: "intent.state.out"),
    ]
}

/// Sign up for a raid (spec §5.7): one intent for Siri, Shortcuts, the Action button and the
/// Control. Compiled into the app and the widget extension, so the Control runs it too.
struct SignUpForRaidIntent: AppIntent {
    static let title: LocalizedStringResource = "intent.signup.title"

    @Parameter(title: "intent.signup.title", default: .in) var state: RaidAnswerState
    @Parameter(title: "intent.character") var character: CharacterEntity?
    /// Set only on the App Shortcut phrases: run by voice, the sign-up asks once first. Shortcuts,
    /// the Action button, a widget and the Control leave it off and never ask.
    @Parameter(title: "intent.confirmAction", default: false) var askFirst: Bool

    static var parameterSummary: some ParameterSummary {
        Summary("\(\.$state)") { \.$character }
    }

    init() {}

    init(state: RaidAnswerState, askFirst: Bool = false) {
        self.state = state
        self.askFirst = askFirst
    }

    func perform() async throws -> some IntentResult & ProvidesDialog {
        guard let calls = LiveIntentCalls.fromBundle() else { return .result(dialog: IntentDialog(stringLiteral: Copy.commonFailed)) }
        let signup = IntentSignup(
            calls: calls, answers: IntentAnswers(clock: WidgetClock(zone: .current)),
            now: { Date() }, enabled: { SystemSwitch.isOn() }
        )
        let reply = try await signup.run(state: state.rawValue, character: character?.characterKey, askFirst: askFirst, asker: SiriAsker(intent: self))
        if reply.saved { WidgetCenter.shared.reloadAllTimelines() }
        return .result(dialog: IntentDialog(stringLiteral: reply.sentence))
    }
}

/// Siri's questions for the sign-up: a choice list (`Which raid?`, `Which spec is ...?`) and one
/// confirmation with `Sign up` and `Cancel`.
struct SiriAsker: IntentAsker {
    let intent: SignUpForRaidIntent

    func choose(_ question: String, options: [String]) async throws -> Int? {
        let choices = options.map { IntentChoiceOption(title: LocalizedStringResource(stringLiteral: $0)) }
        let picked = try await intent.requestChoice(
            between: choices + [.cancel], dialog: IntentDialog(stringLiteral: question)
        )
        return choices.firstIndex(of: picked)
    }

    func confirm(_ question: String) async throws -> Bool {
        let action = ConfirmationActionName.custom(
            acceptLabel: LocalizedStringResource(stringLiteral: Copy.intentConfirmAction), acceptAlternatives: [],
            denyLabel: LocalizedStringResource(stringLiteral: Copy.commonCancel), denyAlternatives: []
        )
        try await intent.requestConfirmation(actionName: action, dialog: IntentDialog(stringLiteral: question))
        return true
    }
}
