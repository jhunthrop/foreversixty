import AppIntents
import Foundation
import WidgetKit

/// The widgets' one setting (spec §5.6 Configuration): which character, `Current character` by default.
struct CharacterEntity: AppEntity {
    static let currentID = "current"

    let id: String
    let name: String

    // App Intents metadata takes literal keys; each is a key of the copy catalog (Localizable.xcstrings).
    static let typeDisplayRepresentation = TypeDisplayRepresentation(name: "widget.characterSetting")

    var displayRepresentation: DisplayRepresentation {
        DisplayRepresentation(title: LocalizedStringResource(stringLiteral: name))
    }

    static let defaultQuery = CharacterQuery()

    static var current: CharacterEntity { CharacterEntity(id: currentID, name: Copy.widgetCurrentCharacter) }

    /// The snapshot file this choice reads; nil is the current character.
    var characterKey: String? { id == Self.currentID ? nil : id }
}

/// The account's characters, as the app last wrote them to the App Group.
struct CharacterQuery: EntityQuery {
    func entities(for identifiers: [String]) async throws -> [CharacterEntity] {
        let all = try await suggestedEntities()
        return identifiers.compactMap { id in all.first { $0.id == id } }
    }

    func suggestedEntities() async throws -> [CharacterEntity] {
        [CharacterEntity.current] + WidgetSnapshotStore.shared().characters().map { CharacterEntity(id: $0.key, name: $0.name) }
    }

    func defaultResult() async -> CharacterEntity? { .current }
}

/// The Next raid and Next upgrade widgets' configuration.
struct CharacterConfigurationIntent: WidgetConfigurationIntent {
    static let title: LocalizedStringResource = "widget.characterSetting"

    @Parameter(title: "widget.characterSetting")
    var character: CharacterEntity?

    init() {}

    init(character: CharacterEntity?) {
        self.character = character
    }

    var characterKey: String? { character?.characterKey }
}

/// In, Late or Out pressed on the medium widget: answered without opening the app (spec §5.6).
struct RaidAnswerIntent: AppIntent {
    static let title: LocalizedStringResource = "widget.nextRaid"
    static let isDiscoverable = false

    @Parameter(title: "State") var state: String
    @Parameter(title: "Character") var character: String?

    init() {}

    init(state: String, character: String?) {
        self.state = state
        self.character = character
    }

    func perform() async throws -> some IntentResult {
        guard let calls = LiveWidgetCalls.fromBundle() else { return .result() }
        let signup = WidgetSignup(
            store: .shared(), character: character, calls: calls, now: { Date() },
            reload: { WidgetCenter.shared.reloadAllTimelines() }
        )
        await signup.answer(state)
        return .result()
    }
}

/// Late or Out pressed on the raid Live Activity. A LiveActivityIntent runs in the app's process,
/// which then redraws or ends the activity from the new snapshot (LiveSurfaceHooks).
struct RaidLiveAnswerIntent: LiveActivityIntent {
    static let title: LocalizedStringResource = "widget.nextRaid"
    static let isDiscoverable = false

    @Parameter(title: "State") var state: String

    init() {}

    init(state: String) {
        self.state = state
    }

    func perform() async throws -> some IntentResult {
        guard let calls = LiveWidgetCalls.fromBundle() else { return .result() }
        let signup = WidgetSignup(
            store: .shared(), character: nil, calls: calls, now: { Date() },
            reload: { WidgetCenter.shared.reloadAllTimelines() }
        )
        await signup.answer(state)
        await LiveSurfaceHooks.afterLiveAnswer()
        return .result()
    }
}

/// What the app adds to an intent run in its own process: the Live Activity redraw.
enum LiveSurfaceHooks {
    nonisolated(unsafe) static var onLiveAnswer: (@Sendable () async -> Void)?

    static func afterLiveAnswer() async {
        await onLiveAnswer?()
    }
}
