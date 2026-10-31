import Auth
import Foundation
import Models
import Networking

/// The character search F12 asks of the API (`GET /v1/characters/search`). Tests pass a fake.
protocol CharacterSearchService: Sendable {
    func search(_ query: String, limit: Int) async throws(APIError) -> [CharacterSearchRow]
}

struct LiveCharacterSearchService: CharacterSearchService {
    let client: APIClient

    func search(_ query: String, limit: Int) async throws(APIError) -> [CharacterSearchRow] {
        try await client.send(.searchCharacters(q: query, limit: limit)).rows
    }
}

/// Who is searching, read from the session when the search opens.
struct SearchContext: Equatable, Sendable {
    var character: AccountCharacter?
    var ownKeys: Set<String> = []

    var viewer: SearchViewer {
        SearchViewer(classSlug: character?.className, region: character?.region, ruleset: character?.ruleset, ownKeys: ownKeys)
    }

    static let signedOut = SearchContext()
}

/// What F12 and the item page need from the app: the shared static data, the character search,
/// Recent's storage, the build in use and who is searching.
struct SearchDependencies {
    let index: SearchIndexReader
    let catalog: ItemCatalog
    let specs: SpecCatalog
    let characters: any CharacterSearchService
    let recent: any RecentSearchStorage
    let build: @Sendable () async -> String?
    let context: @MainActor () -> SearchContext
    /// The 250 ms wait before characters are asked; tests pass one that returns at once.
    var sleep: @Sendable (Duration) async throws -> Void = { try await Task.sleep(for: $0) }

    @MainActor
    static func live(_ environment: AppEnvironment) -> SearchDependencies {
        let session = environment.session
        let appConfig = environment.appConfig
        return SearchDependencies(
            index: environment.searchIndex,
            catalog: ItemCatalog(store: environment.staticData),
            specs: .bundled(),
            characters: LiveCharacterSearchService(client: environment.client),
            recent: UserDefaultsRecentSearches(),
            build: {
                if let build = appConfig.current?.activeBuild { return build }
                await appConfig.refresh()
                return appConfig.current?.activeBuild ?? BundledBuild.read()
            },
            context: {
                SearchContext(
                    character: session.currentCharacter,
                    ownKeys: Set(session.phase.me?.characters.map(\.key) ?? [])
                )
            }
        )
    }
}
