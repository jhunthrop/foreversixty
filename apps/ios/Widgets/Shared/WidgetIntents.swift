import AppIntents
import Foundation
import UIKit
import WidgetKit

/// The widgets' one setting (spec §5.6 Configuration): which character, `Current character` by default.
struct CharacterEntity: AppEntity {
    static let currentID = "current"

    let id: String
    let name: String
    /// Lane F15: the Spotlight line, `Elemental Shaman · 60 · Olympus XXVII`.
    var subtitle: String?
    /// Lane F15: the class slug whose ringed crest is the entry's thumbnail (one crest language).
    var crest: String?

    // App Intents metadata takes literal keys; each is a key of the copy catalog (Localizable.xcstrings).
    static let typeDisplayRepresentation = TypeDisplayRepresentation(name: "widget.characterSetting")

    var displayRepresentation: DisplayRepresentation {
        DisplayRepresentation(
            title: LocalizedStringResource(stringLiteral: name),
            subtitle: subtitle.map { LocalizedStringResource(stringLiteral: $0) },
            image: BrandImages.crest(crest)?.pngData().map { DisplayRepresentation.Image(data: $0) }
        )
    }

    static let defaultQuery = CharacterQuery()

    static var current: CharacterEntity { CharacterEntity(id: currentID, name: Copy.widgetCurrentCharacter) }

    /// The snapshot file this choice reads; nil is the current character.
    var characterKey: String? { id == Self.currentID ? nil : id }
}

/// The account's characters, as the app last wrote them to the App Group.
struct CharacterQuery: EntityQuery {
    func entities(for identifiers: [String]) async throws -> [CharacterEntity] {
        // Lane F15: a Spotlight tap on a watched character resolves through the indexed list too.
        let indexed = WidgetSnapshotStore.shared().spotlightCharacters().map { CharacterEntity(id: $0.key, name: $0.name) }
        let all = try await suggestedEntities() + indexed
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
