@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@Suite("Search logic")
struct SearchLogicTests {
    private func item(_ id: Int, _ name: String, quality: Int = 3, level: Int = 50, classes: [String] = ["mage"]) -> SearchItem {
        SearchItem(id: id, name: name, icon: "x", quality: quality, slot: "back", itemLevel: level, requiredLevel: 45, classes: classes)
    }

    private func corpus(items: [SearchItem], talents: [SearchTalent] = [], guides: [SearchGuide] = []) -> SearchCorpus {
        SearchCorpus(SearchIndex(build: "b", itemUsesShards: 64, items: items, talents: talents, guides: guides))
    }

    @Test("match order: exact, prefix, word prefix, contains; nothing for no match")
    func matchOrder() {
        let match = { (name: String, query: String) in SearchLogic.match(SearchEntry(0, name: name), query: query) }
        #expect(match("Cape", "cape") == .exact)
        #expect(match("Cape of Fire", "cape") == .prefix)
        #expect(match("Spritecaster Cape", "cape") == .wordPrefix)
        #expect(match("Escape Artist", "cape") == .contains)
        #expect(match("Robe", "cape") == nil)
        #expect(match("Hand-of-Ragnaros", "rag") == .wordPrefix)
    }

    @Test("items rank by match, then class-usable, then quality, then item level")
    func itemRanking() {
        let items = [
            item(1, "Arcane Cape", quality: 4, level: 60, classes: ["mage"]),
            item(2, "Cape of Storms", quality: 2, level: 40, classes: ["shaman"]),
            item(3, "Cape of Ash", quality: 3, level: 55, classes: ["mage"]),
            item(4, "Cape of Ice", quality: 3, level: 58, classes: ["mage"]),
            item(5, "Cape", quality: 1, level: 10, classes: ["mage"]),
        ]
        let viewer = SearchViewer(classSlug: "shaman")
        let ids = SearchLogic.items(corpus(items: items).items, query: "cape", viewer: viewer).map(\.id)
        // Exact first; prefix matches with the shaman's usable one first, then quality and level; word prefix last.
        #expect(ids == [5, 2, 4, 3, 1])
    }

    @Test("All caps each group at five with the total; a scope lifts the cap and hides the rest")
    func capsAndScope() {
        let items = (1...8).map { item($0, "Cape \($0)") }
        let guides = [SearchGuide(className: "mage", spec: "arcane", specName: "Arcane", role: "dps", icon: "x")]
        let corpus = corpus(items: items, guides: guides)
        let all = SearchLogic.results(corpus: corpus, query: "cape", scope: .all, viewer: .anonymous, characters: [])
        #expect(all.items.rows.count == 5)
        #expect(all.items.total == 8)
        #expect(all.items.hasMore)
        let scoped = SearchLogic.results(corpus: corpus, query: "Cape", scope: .items, viewer: .anonymous, characters: [])
        #expect(scoped.items.rows.count == 8)
        #expect(!scoped.items.hasMore)
        let guidesOnly = SearchLogic.results(corpus: corpus, query: "cape", scope: .guides, viewer: .anonymous, characters: [])
        #expect(guidesOnly.items.rows.isEmpty && guidesOnly.isEmpty)
        #expect(SearchLogic.results(corpus: corpus, query: "  ", scope: .all, viewer: .anonymous, characters: []) == .empty)
    }

    @Test("a guide matches by `{Spec} {Class}`; the viewer's class comes first")
    func guides() {
        let guides = [
            SearchGuide(className: "mage", spec: "frost", specName: "Frost", role: "dps", icon: "x"),
            SearchGuide(className: "shaman", spec: "elemental", specName: "Elemental", role: "dps", icon: "x"),
            SearchGuide(className: "shaman", spec: "enhancement", specName: "Enhancement", role: "dps", icon: "x"),
        ]
        let hits = SearchLogic.guides(corpus(items: [], guides: guides).guides, query: "shaman", viewer: .anonymous)
        #expect(hits.map(\.title) == ["Elemental Shaman", "Enhancement Shaman"])
        let mine = SearchLogic.guides(corpus(items: [], guides: guides).guides, query: "e", viewer: SearchViewer(classSlug: "shaman"))
        #expect(mine.first?.guide.className == "shaman")
    }

    @Test("characters: the viewer's region and ruleset first, then by match, at most 8")
    func characters() {
        let rows = [
            SearchTestData.row("Elegy", className: "priest", level: 58, region: "eu", ruleset: "pvp"),
            SearchTestData.row("Elementalx", className: "shaman", level: 60),
            SearchTestData.row("Ele", className: "mage", level: 60, region: "eu", ruleset: "pvp"),
        ] + (1...9).map { SearchTestData.row("Elex\($0)", className: nil, level: nil) }
        let viewer = SearchViewer(classSlug: "shaman", region: "us", ruleset: "normal")
        let ranked = SearchLogic.characters(rows, query: "ele", viewer: viewer)
        #expect(ranked.count == SearchLogic.characterLimit)
        #expect(ranked.allSatisfy { $0.region == "us" })
        #expect(ranked.first?.name == "Elementalx")
    }

    @Test("characters are asked from two letters, only in a scope that shows them")
    func characterGate() {
        #expect(!SearchLogic.asksCharacters("e", scope: .all))
        #expect(!SearchLogic.asksCharacters(" e ", scope: .all))
        #expect(SearchLogic.asksCharacters("el", scope: .all))
        #expect(SearchLogic.asksCharacters("el", scope: .characters))
        #expect(!SearchLogic.asksCharacters("el", scope: .items))
    }

    @Test("Recent: newest first, each once, at most 20; clear empties it")
    func recent() {
        let entry = { (id: Int) in RecentSearch(target: .item(id: id), title: "\(id)", icon: nil, quality: nil, className: nil) }
        var list: [RecentSearch] = []
        for id in 1...25 { list = RecentSearches.adding(entry(id), to: list) }
        #expect(list.count == RecentSearches.cap)
        #expect(list.first?.title == "25")
        list = RecentSearches.adding(entry(10), to: list)
        #expect(list.first?.title == "10")
        #expect(list.filter { $0.target == .item(id: 10) }.count == 1)

        let defaults = UserDefaults(suiteName: "search-recent-\(UUID().uuidString)") ?? .standard
        let storage = UserDefaultsRecentSearches(defaults: defaults)
        storage.save(list)
        #expect(storage.load() == list)
        storage.clear()
        #expect(storage.load().isEmpty)
    }

    @Test("the real 70658 index: `ele` finds all four kinds and keystrokes stay under the budget")
    func realIndex() throws {
        let corpus = try SearchTestData.corpus()
        let viewer = SearchTestData.context.viewer
        let results = SearchLogic.results(corpus: corpus, query: "ele", scope: .all, viewer: viewer, characters: SearchTestData.eleCharacters)
        #expect(results.guides.rows.first?.title == "Elemental Shaman")
        #expect(!results.talents.rows.isEmpty && !results.items.rows.isEmpty && results.characters.rows.count == 3)
        let clock = ContinuousClock()
        let elapsed = clock.measure {
            for query in ["d", "de", "dev", "devo", "devou", "devout"] {
                _ = SearchLogic.results(corpus: corpus, query: query, scope: .all, viewer: viewer, characters: [])
            }
        }
        // Six keystrokes on the simulator, well inside 6 x 50 ms (spec §11).
        #expect(elapsed < .milliseconds(300))
    }

    @Test("an item's shard is id % item_uses_shards; the Cape is in shard 39")
    func shard() throws {
        #expect(try SearchTestData.corpus().shard(of: SearchTestData.cape) == 39)
    }
}
