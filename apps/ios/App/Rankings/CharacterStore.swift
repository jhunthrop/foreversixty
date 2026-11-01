import Foundation
import Models
import Networking
import Observation

/// What the character page draws (spec F11.4): the profile and the offline stamp.
struct CharacterContent: Equatable, Sendable {
    let profile: CharacterProfile
    let offlineSince: Date?
    /// Lane F15: set while the cached answer shows because Forever Sixty did not answer.
    var staleSince: Date?
}

enum CharacterPhase: Equatable, Sendable {
    case loading
    case loaded(CharacterContent)
    /// 404: no such character, or its account is anonymous (F11.7).
    case notFound
    case failed
}

/// Loads one character's page (cache first) and its class's items for the Paperdoll.
@MainActor
@Observable
final class CharacterStore {
    static let notFound = 404

    private(set) var phase: CharacterPhase = .loading
    /// The class items file by id; empty until read (the sockets draw plain until then).
    private(set) var items: [Int: Item] = [:]

    let ref: CharacterRef
    private let service: any RankingsService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date
    private let itemsFor: @Sendable (String) async -> [Int: Item]
    /// Lane F15: the connection, for the stale strip; the screen sets it, nil keeps v1's rule.
    var connectivity: (any Connectivity)?

    init(
        ref: CharacterRef, service: any RankingsService, cache: any RaidCache, now: @escaping @Sendable () -> Date,
        items: @escaping @Sendable (String) async -> [Int: Item]
    ) {
        self.ref = ref
        self.service = service
        self.cache = cache
        self.now = now
        itemsFor = items
    }

    private var cacheKey: String { "character-\(ref.key)" }


    private var content: CharacterContent? {
        if case let .loaded(content) = phase { return content }
        return nil
    }

    func load() async {
        if content == nil, let cached = cache.read(CharacterProfile.self, key: cacheKey) {
            phase = .loaded(CharacterContent(profile: cached.value, offlineSince: nil))
        }
        do {
            let profile = try await service.character(ref)
            cache.write(profile, key: cacheKey, at: now())
            phase = .loaded(CharacterContent(profile: profile, offlineSince: nil))
        } catch {
            if case let .server(status, _, _) = error, status == Self.notFound {
                phase = .notFound
            } else if let cached = cache.read(CharacterProfile.self, key: cacheKey), let since = CacheFreshness.staleSince(
                error, connectivity: connectivity, savedAt: cached.savedAt, limit: .pages, now: now()
            ) {
                phase = .loaded(CharacterContent(profile: cached.value, offlineSince: nil, staleSince: since))
            } else if SignupRefusal.isOffline(error), let cached = cache.read(CharacterProfile.self, key: cacheKey) {
                phase = .loaded(CharacterContent(profile: cached.value, offlineSince: cached.savedAt))
            } else if content == nil {
                phase = .failed
            }
        }
        await loadItems()
    }

    private func loadItems() async {
        guard items.isEmpty, let profile = content?.profile, !profile.gear.isEmpty,
              let className = profile.character.className?.lowercased()
        else { return }
        items = await itemsFor(className)
    }
}

/// The account's watch list (spec F11.4 Watching): which characters carry the gold star. One
/// optimistic toggle for the character page's star and a rankings row's long-press menu.
@MainActor
@Observable
final class WatchList {
    private(set) var keys: Set<String> = []
    /// The character whose last toggle failed and was put back (`char.watchFailed`).
    private(set) var failedKey: String?
    /// Bumps on every confirmed watch: the success haptic's trigger.
    private(set) var confirmed = 0

    let signedIn: Bool
    private let service: any RankingsService

    init(service: any RankingsService, signedIn: Bool, keys: Set<String> = []) {
        self.service = service
        self.signedIn = signedIn
        self.keys = keys
    }

    func isWatching(_ key: String) -> Bool { keys.contains(key) }

    func load() async {
        guard signedIn, let list = try? await service.watching() else { return }
        keys = Set(list.rows.map(\.key))
    }

    /// Flips at once; a failed call puts it back and names the character.
    func toggle(_ ref: CharacterRef) async {
        let watching = keys.contains(ref.key)
        failedKey = nil
        if watching { keys.remove(ref.key) } else { keys.insert(ref.key) }
        do {
            if watching {
                try await service.unwatch(ref)
            } else {
                try await service.watch(ref)
                confirmed += 1
            }
        } catch {
            if watching { keys.insert(ref.key) } else { keys.remove(ref.key) }
            failedKey = ref.key
        }
    }
}
