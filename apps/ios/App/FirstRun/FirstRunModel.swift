import Auth
import Foundation
import Models
import Networking
import Observation

/// The first run after sign-in (flows spec F8): Pick your main, then, for a guild member, the
/// notification ask at its moment of value (E6). Never at launch, never twice.
@MainActor
@Observable
final class FirstRunModel {
    enum Step: Equatable, Identifiable {
        case pickMain
        case permission(guild: String)

        var id: String {
            switch self {
            case .pickMain: "pickMain"
            case let .permission(guild): "permission-\(guild)"
            }
        }
    }

    /// The characters Pick your main offers: level 10 and up, the highest first.
    static let minimumLevel = 10

    struct Dependencies {
        let setMain: @MainActor (String) async throws -> Void
        let permission: any NotificationPermission
        let asked: any FirstRunAskStore
    }

    private(set) var step: Step?
    /// The row the button continues with; the highest level until a tap.
    var selection: String?
    private(set) var saving = false
    private(set) var failed = false
    private let deps: Dependencies

    init(dependencies: Dependencies) {
        deps = dependencies
    }

    static func candidates(_ me: Me) -> [AccountCharacter] {
        me.characters
            .filter { ($0.level ?? 0) >= minimumLevel }
            .enumerated()
            .sorted { ($0.element.level ?? 0, -$0.offset) > ($1.element.level ?? 0, -$1.offset) }
            .map(\.element)
    }

    /// The account loaded or changed: asks for a main when it has none.
    func update(me: Me?) {
        guard let me else { return }
        if me.mainCharacterKey == nil, let first = Self.candidates(me).first {
            if step != .pickMain { selection = first.key }
            step = .pickMain
        } else if step == .pickMain {
            step = nil
        }
    }

    func continueWithMain(_ me: Me?) async {
        guard let key = selection, let character = me?.characters.first(where: { $0.key == key }) else { return }
        saving = true
        failed = false
        defer { saving = false }
        do {
            try await deps.setMain(key)
        } catch {
            failed = true
            return
        }
        step = await permissionStep(for: character)
    }

    /// The ask follows the main into a guild, once per install, while the system has not decided.
    private func permissionStep(for character: AccountCharacter) async -> Step? {
        guard let guild = character.guild?.name, !deps.asked.wasAsked(), await deps.permission.isUndetermined() else {
            return nil
        }
        return .permission(guild: guild)
    }

    func turnOnNotifications() async {
        deps.asked.markAsked()
        step = nil
        await deps.permission.request()
    }

    func notNow() {
        deps.asked.markAsked()
        step = nil
    }
}

/// Whether the first-run ask was shown on this install (UserDefaults, gone with the app).
protocol FirstRunAskStore: Sendable {
    func wasAsked() -> Bool
    func markAsked()
}

struct DefaultsFirstRunAskStore: FirstRunAskStore {
    static let key = "gg.foreversixty.first-run.permission-asked"

    func wasAsked() -> Bool { UserDefaults.standard.bool(forKey: Self.key) }
    func markAsked() { UserDefaults.standard.set(true, forKey: Self.key) }
}
