import Foundation
import Models
import Networking
import Observation

/// The search index as the screen has it.
enum SearchIndexPhase: Sendable {
    case loading
    case ready(SearchCorpus)
    case failed

    var corpus: SearchCorpus? {
        if case let .ready(corpus) = self { return corpus }
        return nil
    }
}

/// The character group's own state: only it needs a connection (spec F12.6).
enum CharacterSearchPhase: Equatable, Sendable {
    case idle
    case loading
    case loaded(query: String, rows: [CharacterSearchRow])
    case offline
    case failed
}

/// F12's state: the query and scope, the index, the characters, Recent. Items, guides and talents
/// filter in memory on every keystroke; characters are asked after 250 ms and two letters.
@MainActor
@Observable
final class SearchStore {
    var query = ""
    var scope: SearchScope = .all
    private(set) var index: SearchIndexPhase = .loading
    private(set) var characters: CharacterSearchPhase = .idle
    private(set) var recent: [RecentSearch]
    private(set) var context: SearchContext

    private let dependencies: SearchDependencies

    init(dependencies: SearchDependencies) {
        self.dependencies = dependencies
        recent = dependencies.recent.load()
        context = dependencies.context()
    }

    /// The rows for the current query and scope.
    var results: SearchResults {
        SearchLogic.results(
            corpus: index.corpus, query: query, scope: scope, viewer: context.viewer, characters: characterRows
        )
    }

    /// The character rows that belong to the current query (a late answer for an older one is dropped).
    private var characterRows: [CharacterSearchRow] {
        guard case let .loaded(asked, rows) = characters, asked == SearchLogic.normalized(query) else { return [] }
        return rows
    }

    var hasQuery: Bool { !SearchLogic.normalized(query).isEmpty }

    /// Loads the active build's index: cache first, so it works offline once fetched.
    func loadIndex() async {
        index = .loading
        guard let build = await dependencies.build() else {
            index = .failed
            return
        }
        do {
            index = .ready(try await dependencies.index.corpus(build: build))
        } catch {
            index = .failed
        }
    }

    /// Run on every query change (the view's `.task(id:)` cancels the last one): waits 250 ms,
    /// then asks for characters when the query has two letters and the scope shows them.
    func searchCharacters() async {
        let asked = SearchLogic.normalized(query)
        guard SearchLogic.asksCharacters(asked, scope: scope) else {
            characters = .idle
            return
        }
        do {
            try await dependencies.sleep(SearchLogic.characterDebounce)
        } catch {
            return
        }
        guard !Task.isCancelled else { return }
        characters = .loading
        do {
            let rows = try await dependencies.characters.search(asked, limit: SearchLogic.characterAsk)
            guard !Task.isCancelled else { return }
            characters = .loaded(query: asked, rows: rows)
        } catch {
            guard !Task.isCancelled else { return }
            characters = SignupRefusal.isOffline(error) ? .offline : .failed
        }
    }

    /// `Show all {n}`: the group's scope, without its cap.
    func showAll(_ group: SearchGroupKind) {
        scope = SearchScope(group)
    }

    /// Opening a result puts it first in Recent.
    func opened(_ entry: RecentSearch) {
        recent = RecentSearches.adding(entry, to: recent)
        dependencies.recent.save(recent)
    }

    func clearRecent() {
        recent = []
        dependencies.recent.clear()
    }
}

extension RecentSearch {
    init(item: SearchItem) {
        self.init(target: .item(id: item.id), title: item.name, icon: item.icon, quality: item.quality, className: nil)
    }

    init(guide: SearchGuideHit) {
        self.init(
            target: .guide(className: guide.guide.className, spec: guide.guide.spec), title: guide.title,
            icon: guide.guide.icon, quality: nil, className: guide.guide.className
        )
    }

    init(talent: SearchTalent) {
        self.init(
            target: .talent(id: talent.id, className: talent.className), title: talent.name,
            icon: talent.icon, quality: nil, className: nil
        )
    }

    init(character: CharacterSearchRow) {
        self.init(target: .character(key: character.key), title: character.name, icon: nil, quality: nil, className: character.className)
    }
}
