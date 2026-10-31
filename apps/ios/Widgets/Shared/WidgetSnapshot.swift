import Foundation
import Models
import Networking

/// The App Group both the app and the widget extension belong to (project.yml entitlements).
enum AppGroup {
    static let id = "group.gg.foreversixty.app"

    /// The shared container; nil when the process lacks the entitlement (a unit-test host
    /// built without signing), where the caller falls back to its own Application Support.
    static var containerURL: URL? {
        FileManager.default.containerURL(forSecurityApplicationGroupIdentifier: id)
    }

    /// Settings both processes read (the Live surfaces switches).
    static var defaults: UserDefaults { UserDefaults(suiteName: id) ?? .standard }
}

/// An answer pressed on a widget, drawn at once while the server is asked (spec §5.6).
struct PendingAnswer: Codable, Equatable, Sendable {
    let raidID: Int64
    let state: String
}

/// What every surface draws from: `GET /v1/me/widget`'s answer, when the app last fetched it,
/// and the answer a widget press is waiting on or just failed to save. One file in the App
/// Group, written by the app and the sign-up intent, read by every widget and Live Activity.
struct WidgetSnapshot: Codable, Equatable, Sendable {
    let widget: Widget
    let fetchedAt: Date
    var pending: PendingAnswer?
    /// `Not saved` shows on the pressed button until then (10 seconds after a failure).
    var notSavedUntil: Date?
    var notSavedState: String?

    /// How long `Not saved` stays (spec §5.6).
    static let notSavedFor: TimeInterval = 10
}

/// One of the account's characters, for the widgets' Character setting.
struct WidgetCharacterChoice: Codable, Equatable, Sendable, Identifiable {
    let key: String
    let name: String
    var id: String { key }
}

/// Reads and writes the snapshots: one for the current character (the app's own), and one per
/// character a widget is set to. Signed out is "no current file".
struct WidgetSnapshotStore: Sendable {
    let directory: URL

    static let fileName = "widget-snapshot.json"
    static let charactersFile = "characters.json"

    /// The App Group's store, else this process's Application Support (tests, previews).
    static func shared() -> WidgetSnapshotStore {
        let base = AppGroup.containerURL
            ?? FileManager.default.urls(for: .applicationSupportDirectory, in: .userDomainMask)[0]
        return WidgetSnapshotStore(directory: base.appending(path: "widgets", directoryHint: .isDirectory))
    }

    /// nil is the current character's file.
    func fileURL(character: String? = nil) -> URL {
        guard let character else { return directory.appending(path: Self.fileName) }
        let safe = character.map { $0.isLetter || $0.isNumber || $0 == "-" ? $0 : "_" }
        return directory.appending(path: "widget-snapshot-\(String(safe)).json")
    }

    func load(character: String? = nil) -> WidgetSnapshot? {
        guard let data = try? Data(contentsOf: fileURL(character: character)) else { return nil }
        return try? WireFormat.decoder().decode(WidgetSnapshot.self, from: data)
    }

    func save(_ snapshot: WidgetSnapshot, character: String? = nil) throws {
        try FileManager.default.createDirectory(at: directory, withIntermediateDirectories: true)
        try WireFormat.encoder().encode(snapshot).write(to: fileURL(character: character), options: .atomic)
    }

    /// The snapshot a widget set to `character` draws: its own file, else the current one when
    /// that is the same character, else nil (signed out when there is no current file either).
    func snapshot(for character: String?) -> WidgetSnapshot? {
        let current = load()
        guard let character, current?.widget.character?.key != character else { return current }
        return current == nil ? nil : load(character: character)
    }

    func saveCharacters(_ characters: [WidgetCharacterChoice]) {
        try? FileManager.default.createDirectory(at: directory, withIntermediateDirectories: true)
        try? JSONEncoder().encode(characters).write(to: directory.appending(path: Self.charactersFile), options: .atomic)
    }

    func characters() -> [WidgetCharacterChoice] {
        guard let data = try? Data(contentsOf: directory.appending(path: Self.charactersFile)) else { return [] }
        return (try? JSONDecoder().decode([WidgetCharacterChoice].self, from: data)) ?? []
    }

    /// Sign-out: every surface returns to its signed-out words.
    func clear() {
        try? FileManager.default.removeItem(at: directory)
    }

    /// Where the app leaves item icons for the extension, which cannot fetch while it draws.
    func iconURL(_ icon: String) -> URL {
        directory.appending(path: "icons/\(icon).png")
    }
}
