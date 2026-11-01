@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@MainActor
@Suite("Search store")
struct SearchStoreTests {
    private func dependencies(
        characters: FakeCharacterSearch = FakeCharacterSearch(.success(SearchTestData.eleCharacters)),
        recent: MemoryRecentSearches = MemoryRecentSearches(),
        transport: CountingTransport = CountingTransport(),
        build: String? = SearchTestData.build
    ) -> SearchDependencies {
        let store = StaticDataStore(client: StaticDataClient(transport: transport), directory: nil)
        return SearchDependencies(
            index: SearchIndexReader(store: store),
            catalog: ItemCatalog(store: store),
            specs: (try? BisTestData.catalog()) ?? SpecCatalog(specs: []),
            characters: characters,
            recent: recent,
            build: { build },
            context: { SearchTestData.context },
            sleep: { _ in }
        )
    }

    @Test("the index loads from the fixture; items filter on every keystroke")
    func loadsAndFilters() async {
        let store = SearchStore(dependencies: dependencies())
        await store.loadIndex()
        #expect(store.index.corpus != nil)
        store.query = "spritec"
        #expect(store.results.items.rows.map(\.id) == [SearchTestData.cape])
        store.query = "convection"
        #expect(store.results.talents.rows.first?.name == "Convection")
    }

    @Test("no build, or a failed index, is the failed state")
    func indexFails() async {
        let offline = CountingTransport()
        offline.goOffline()
        let store = SearchStore(dependencies: dependencies(transport: offline))
        await store.loadIndex()
        guard case .failed = store.index else { Issue.record("expected failed"); return }
        let noBuild = SearchStore(dependencies: dependencies(build: nil))
        await noBuild.loadIndex()
        guard case .failed = noBuild.index else { Issue.record("expected failed"); return }
    }

    @Test("characters wait for two letters; then the answer belongs to its query")
    func characterAsk() async {
        let fake = FakeCharacterSearch(.success(SearchTestData.eleCharacters))
        let store = SearchStore(dependencies: dependencies(characters: fake))
        store.query = "e"
        await store.searchCharacters()
        #expect(fake.asked.isEmpty)
        #expect(store.characters == .idle)
        store.query = "Ele"
        await store.searchCharacters()
        #expect(fake.asked == ["ele"])
        #expect(store.results.characters.rows.count == 3)
        store.query = "elem"
        #expect(store.results.characters.rows.isEmpty)
    }

    @Test("a cancelled wait never asks (the debounce)")
    func debounceCancels() async {
        let fake = FakeCharacterSearch(.success([]))
        var deps = dependencies(characters: fake)
        deps.sleep = { _ in throw CancellationError() }
        let store = SearchStore(dependencies: deps)
        store.query = "ele"
        await store.searchCharacters()
        #expect(fake.asked.isEmpty)
    }

    @Test("offline: characters say so; any other failure is the failed line")
    func characterFailures() async {
        let offline = SearchStore(dependencies: dependencies(characters: FakeCharacterSearch(.failure(.transport("offline")))))
        offline.query = "xqz"
        await offline.searchCharacters()
        #expect(offline.characters == .offline)
        let failed = SearchStore(dependencies: dependencies(characters: FakeCharacterSearch(.failure(.unexpectedStatus(500)))))
        failed.query = "xqz"
        await failed.searchCharacters()
        #expect(failed.characters == .failed)
    }

    @Test("opening adds to Recent and saves it; Clear empties it; Show all switches the scope")
    func recentAndShowAll() {
        let recent = MemoryRecentSearches()
        let store = SearchStore(dependencies: dependencies(recent: recent))
        store.opened(SearchTestData.recents[1])
        store.opened(SearchTestData.recents[0])
        #expect(recent.load().map(\.title) == ["Spritecaster Cape", "Elemental Shaman"])
        store.clearRecent()
        #expect(store.recent.isEmpty && recent.load().isEmpty)
        store.showAll(.talents)
        #expect(store.scope == .talents)
    }
}
