import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// A talent result's sheet, by the talent it names.
struct SearchTalentRoute: Identifiable {
    let talent: SearchTalent
    var id: Int { talent.id }
}

/// What a keystroke or a chip changes for the character ask (`.task(id:)` restarts on it).
private struct CharacterAskKey: Equatable {
    let query: String
    let scope: SearchScope
}

/// F12, full screen over the tabs (spec §3.2, F12.1 to F12.6): the field with the keyboard up,
/// the chips, then the empty state or the results. Items push the item page on this stack.
struct SearchScreen: View {
    let navigator: SearchNavigator

    @State private var store: SearchStore
    @State private var path: [ItemPageRoute] = []
    @State private var talent: SearchTalentRoute?
    @FocusState private var focused: Bool
    @Environment(\.openURL) private var openURL

    init(navigator: SearchNavigator) {
        self.navigator = navigator
        _store = State(initialValue: SearchStore(dependencies: navigator.dependencies))
    }

    private var dependencies: SearchDependencies { navigator.dependencies }

    var body: some View {
        @Bindable var store = store
        NavigationStack(path: $path) {
            SearchChrome(query: $store.query, scope: $store.scope, focused: $focused, cancel: { navigator.isPresented = false }) {
                SearchContentView(
                    hasQuery: store.hasQuery, query: store.query, scope: store.scope, index: store.index,
                    results: store.results, characters: store.characters, recent: store.recent,
                    forCharacter: SearchForCharacter(store.context, specs: dependencies.specs),
                    viewer: store.context.viewer, build: build, actions: actions
                )
            }
            .navigationTitle(Copy.itemPageBack)
            .toolbar(.hidden, for: .navigationBar)
            .navigationDestination(for: ItemPageRoute.self) { route in
                ItemPageScreen(route: route, dependencies: dependencies, backTitle: Copy.itemPageBack) { navigator.openList($0) }
            }
        }
        .sheet(item: $talent) { route in
            SearchTalentSheet(talent: route.talent, build: build, index: dependencies.index) { openURL(SearchLinks.planner) }
        }
        .task { await store.loadIndex() }
        .task(id: CharacterAskKey(query: store.query, scope: store.scope)) { await store.searchCharacters() }
        .onAppear { focused = true }
    }

    private var build: String { store.index.corpus?.build ?? "" }

    private var actions: SearchActions {
        SearchActions(
            openItem: { item in
                store.opened(RecentSearch(item: item))
                path.append(ItemPageRoute(itemID: item.id))
            },
            openCharacter: { row in
                store.opened(RecentSearch(character: row))
                openURL(SearchLinks.character(key: row.key))
            },
            openGuide: { hit in
                store.opened(RecentSearch(guide: hit))
                openURL(SearchLinks.guide(className: hit.guide.className, spec: hit.guide.spec))
            },
            openTalent: { found in
                store.opened(RecentSearch(talent: found))
                talent = SearchTalentRoute(talent: found)
            },
            openRecent: open,
            showAll: { store.showAll($0) },
            clearRecent: { store.clearRecent() },
            retryIndex: { Task { await store.loadIndex() } },
            retryCharacters: { Task { await store.searchCharacters() } },
            openMyBis: {
                navigator.openList(store.context.character.flatMap { BisSelection.forCharacter($0, catalog: dependencies.specs) })
            },
            openURL: { openURL($0) }
        )
    }

    /// A Recent row opens what it remembers, and moves to the top.
    private func open(_ entry: RecentSearch) {
        store.opened(entry)
        switch entry.target {
        case let .item(id):
            path.append(ItemPageRoute(itemID: id))
        case let .character(key):
            openURL(SearchLinks.character(key: key))
        case let .guide(className, spec):
            openURL(SearchLinks.guide(className: className, spec: spec))
        case let .talent(id, _):
            if let found = store.index.corpus?.talent(id) { talent = SearchTalentRoute(talent: found) }
        }
    }
}
