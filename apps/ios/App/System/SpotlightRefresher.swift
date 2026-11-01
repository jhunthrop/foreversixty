import Auth
import Foundation
import Models
import Networking

/// Rebuilds Spotlight from the data the app already loads (spec §5.7): on launch, when the
/// account or the current character changes, and on foreground. Behind the `system` switch; off
/// or signed out, the index is empty.
struct SpotlightRefresher: Sendable {
    let client: APIClient
    let calls: any IntentCalls
    let guides: any GuidesSource
    let specs: SpecCatalog
    let indexer: SpotlightIndexer
    let now: @Sendable () -> Date

    func refresh(me: Me?, currentKey: String?, enabled: Bool) async {
        SystemSwitch.record(enabled)
        guard enabled, let me else {
            await indexer.clear()
            return
        }
        let watched = (try? await client.send(.getWatching).rows) ?? []
        let current = me.characters.first { $0.key == currentKey } ?? me.mainCharacter
        let guildID = current?.guild?.id
        let planned = await guildID.asyncFlatMap { try? await calls.raids(guild: $0).planned } ?? []
        let catalog = SpotlightCatalog.build(
            own: me.characters, watched: watched, planned: planned, guildID: guildID,
            guides: await savedGuides(), answers: IntentAnswers(clock: WidgetClock(zone: .current)), now: now()
        )
        await indexer.index(catalog)
    }

    /// The guides whose file is on the phone for the active build.
    private func savedGuides() async -> [SpotlightCatalog.SavedGuide] {
        let build = await guides.activeBuild()
        let saved = guides.savedBytes(build: build, specs: specs.specs.map(\.spec))
        return specs.specs.compactMap { spec in
            guard saved.isSaved(spec.spec), let file = guides.cachedGuide(build: build, spec: spec.spec) else { return nil }
            return SpotlightCatalog.SavedGuide(spec: spec, savedAt: file.savedAt)
        }
    }
}

extension Optional where Wrapped: Sendable {
    func asyncFlatMap<T>(_ transform: (Wrapped) async -> T?) async -> T? {
        guard let self else { return nil }
        return await transform(self)
    }
}

extension SpotlightRefresher {
    @MainActor
    static func live(_ environment: AppEnvironment, session: SessionStore?) -> SpotlightRefresher {
        let guides = GuidesDependencies.live(environment, session: session)
        return SpotlightRefresher(
            client: environment.client, calls: LiveIntentCalls(client: environment.client), guides: guides.source,
            specs: guides.specs, indexer: SpotlightIndexer(), now: { Date() }
        )
    }
}
