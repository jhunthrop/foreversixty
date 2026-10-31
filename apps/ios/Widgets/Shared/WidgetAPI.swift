import Auth
import Foundation
import Models
import Networking

/// The two calls the widget surfaces make, in the app and in the extension: the snapshot, and
/// an answer to the next raid. Live goes through `APIClient`; tests pass a fake.
protocol WidgetCalls: Sendable {
    func snapshot(character: String?) async throws(APIError) -> Widget
    func answer(guild: Int64, raid: Int64, characterKey: String, state: String) async throws(APIError)
}

struct LiveWidgetCalls: WidgetCalls {
    let client: APIClient

    func snapshot(character: String?) async throws(APIError) -> Widget {
        try await client.send(.getWidget(character: character))
    }

    /// The sign-up PUT, keeping the character's note (the PUT writes note too), as the raid
    /// notifications' answers do.
    func answer(guild: Int64, raid: Int64, characterKey: String, state: String) async throws(APIError) {
        let view = try await client.send(.getGuildRaid(id: guild, raid: raid))
        let note = view.signups.first { $0.characterKey == characterKey }?.note ?? ""
        let parts = characterKey.split(separator: "/", omittingEmptySubsequences: false).map(String.init)
        guard parts.count == 3 else { throw .decoding("not a character key: \(characterKey)") }
        _ = try await client.send(.putRaidSignup(
            id: guild, raid: raid, region: parts[0], ruleset: parts[1], name: parts[2],
            SignupRequest(state: state, role: nil, note: note, spec: nil)
        ))
    }
}

extension LiveWidgetCalls {
    /// The extension's own client, from its Info.plist (the same FSAPI keys as the app) and the
    /// token the app shares through the App Group keychain. Fixture builds answer from fixtures.
    static func fromBundle(_ bundle: Bundle = .main) -> LiveWidgetCalls? {
        guard let configuration = try? APIConfiguration(infoDictionary: bundle.infoDictionary ?? [:]) else { return nil }
        let tokens: any TokenStore = switch configuration.mode {
        case .live: KeychainTokenStore(accessGroup: AppGroup.id)
        case .fixture: InMemoryTokenStore(token: "fsd_fixture")
        }
        let identity = ClientIdentity.bundled(bundle, installID: InstallID.current(in: AppGroup.defaults))
        return LiveWidgetCalls(client: .make(configuration, tokens: tokens, identity: identity, bundle: bundle))
    }
}

/// A widget press (spec §5.6): the answer draws at once, then the server's; a failure redraws
/// the previous answer with `Not saved` on the button for 10 seconds.
struct WidgetSignup: Sendable {
    let store: WidgetSnapshotStore
    /// The character the pressed widget is set to; nil is the current character.
    var character: String?
    let calls: any WidgetCalls
    let now: @Sendable () -> Date
    let reload: @Sendable () -> Void

    @discardableResult
    func answer(_ state: String) async -> Bool {
        let file = fileCharacter
        guard var snapshot = store.load(character: file), let raid = snapshot.widget.nextRaid,
              let character = snapshot.widget.character else { return false }
        snapshot.pending = PendingAnswer(raidID: raid.raidId, state: state)
        snapshot.notSavedUntil = nil
        snapshot.notSavedState = nil
        try? store.save(snapshot, character: file)
        reload()
        do throws(APIError) {
            try await calls.answer(guild: raid.guildId, raid: raid.raidId, characterKey: character.key, state: state)
        } catch {
            snapshot.pending = nil
            snapshot.notSavedUntil = now().addingTimeInterval(WidgetSnapshot.notSavedFor)
            snapshot.notSavedState = state
            try? store.save(snapshot, character: file)
            reload()
            return false
        }
        // The server's own answer (counts included); while it is out of reach the saved answer stands.
        if let fresh = try? await calls.snapshot(character: character.key) {
            try? store.save(WidgetSnapshot(widget: fresh, fetchedAt: now()), character: file)
        } else {
            snapshot.pending = nil
            try? store.save(WidgetSnapshot(widget: snapshot.widget.answered(state), fetchedAt: snapshot.fetchedAt), character: file)
        }
        reload()
        return true
    }

    /// The file the press reads: the character's own, unless it is the current character.
    private var fileCharacter: String? {
        guard let character, store.load()?.widget.character?.key != character else { return nil }
        return character
    }
}

extension Widget {
    /// The snapshot with the next raid's own answer replaced, for a saved press the server's
    /// fresh snapshot could not confirm yet.
    func answered(_ state: String) -> Widget {
        guard let raid = nextRaid else { return self }
        let my = WidgetMine(state: state, spec: raid.my.spec, role: raid.my.role)
        let updated = WidgetNextRaid(
            raidId: raid.raidId, guildId: raid.guildId, guildName: raid.guildName, zone: raid.zone, zoneName: raid.zoneName,
            startsAt: raid.startsAt, inviteAt: raid.inviteAt, pullAt: raid.pullAt, endsAt: raid.endsAt, size: raid.size,
            signedUp: raid.signedUp, status: raid.status, my: my, canSignUp: raid.canSignUp
        )
        return Widget(
            generatedAt: generatedAt, character: character, guild: guild, nextRaid: updated, lastRaid: lastRaid,
            upgrades: upgrades, live: live
        )
    }
}
