@testable import ForeverSixty
import DesignSystem
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// F12's boards on iPhone 17 (402 x 874) at the default size and at AX5 (spec §12), from the
/// repo's build 70658 and the contract fixtures. Recorded on the shared iOS 26.5 simulator
/// (`SNAPSHOT_TESTING_RECORD=all` to re-record); the PNGs double as the review captures in
/// design/reviews/2026-10-31-mobile-wave3-search/build/ios.
@MainActor
@Suite("Search snapshots")
struct SearchSnapshotTests {
    static let phone = CGSize(width: 402, height: 874)
    static let safeArea = UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0)
    private let build = SearchTestData.build

    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: Self.safeArea, size: CGSize(width: Self.phone.width, height: height), traits: traits)
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    // MARK: Search

    private func search(
        query: String, scope: SearchScope = .all, characters: CharacterSearchPhase? = nil, recent: [RecentSearch] = []
    ) throws -> some View {
        let corpus = try SearchTestData.corpus()
        let context = SearchTestData.context
        let rows = SearchTestData.eleCharacters
        let phase = characters ?? .loaded(query: SearchLogic.normalized(query), rows: rows)
        let asked: [CharacterSearchRow] = if case let .loaded(_, found) = phase { found } else { [] }
        let results = SearchLogic.results(corpus: corpus, query: query, scope: scope, viewer: context.viewer, characters: asked)
        let forCharacter = SearchForCharacter(context, specs: try BisTestData.catalog())
        return SearchHost(query: query, scope: scope) {
            SearchContentView(
                hasQuery: !SearchLogic.normalized(query).isEmpty, query: query, scope: scope, index: .ready(corpus),
                results: results, characters: phase, recent: recent,
                forCharacter: forCharacter,
                viewer: context.viewer, build: build
            )
        }
    }

    @Test func f12_1_empty() throws {
        snap(try search(query: "", recent: SearchTestData.recents), "f12.1-empty-ios")
    }

    @Test func f12_2_all() throws {
        snap(try search(query: "ele"), "f12.2-all-ios", height: 1500)
    }

    @Test func f12_2_all_ax5() throws {
        snap(try search(query: "ele"), "f12.2-all-ax5-ios", height: 3600, large: true)
    }

    @Test func f12_3_items() throws {
        snap(try search(query: "devout", scope: .items), "f12.3-items-ios", height: 1100)
    }

    @Test func f12_6_noResults() throws {
        snap(try search(query: "xqz", characters: .loaded(query: "xqz", rows: [])), "f12.6-no-results-ios")
    }

    @Test func f12_6_offline() throws {
        snap(try search(query: "xqz", characters: .offline), "f12.6-offline-ios")
    }

    /// The account's own character carries `You`; nobody else does.
    @Test func f12_2_characters() throws {
        let rows = SearchTestData.lightningCharacters
        snap(try search(query: "lightning", scope: .characters, characters: .loaded(query: "lightning", rows: rows)), "f12.2-characters-ios")
    }

    @Test func f12_indexFailed() throws {
        let context = SearchTestData.context
        let view = SearchHost(query: "ele", scope: .all) {
            SearchContentView(
                hasQuery: true, query: "ele", scope: .all, index: .failed, results: .empty, characters: .idle, recent: [],
                forCharacter: nil, viewer: context.viewer, build: SearchTestData.build
            )
        }
        snap(view, "f12.index-failed-ios")
    }

    // MARK: The item page

    private let raid = BisSelection(classSlug: "shaman", spec: "shaman-elemental", range: BisRange(band: 60, preset: "raid"), faction: "horde")

    /// A page as it loads for Lightningbolltz (60, Phase 1 raid, Horde) from the repo's build 70658.
    private func itemState(_ id: Int, uses: ItemUses) throws -> ItemPageState {
        let corpus = try SearchTestData.corpus()
        let item = try #require(corpus.item(id))
        var state = ItemPageState(route: ItemPageRoute(itemID: item.id), build: build)
        state.look = HomeItemLook(id: item.id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel)
        state.meta = ItemPageLogic.meta(item)
        let request = ItemPageLogic.tooltipRequest(
            item: item, route: ItemPageRoute(itemID: item.id), build: build, character: SearchTestData.lightningbolltz, own: raid
        )
        let sources = try SearchTestData.tooltipSources(classSlug: request.classSlug)
        state.tooltip = .loaded(.build(request, sources: sources))
        state.pick = ItemPageLogic.yourPick(
            itemID: item.id, band: ItemTooltipModel.band(for: request, in: sources.bis), selection: raid,
            specName: "Elemental", className: "shaman"
        )
        state.sources = .loaded(ItemPageLogic.sources(uses.sources))
        state.specs = .loaded(ItemPageLogic.specs(uses.specs, catalog: try BisTestData.catalog()))
        return state
    }

    private func capeState() throws -> ItemPageState {
        try itemState(SearchTestData.cape, uses: SearchTestData.capeUses())
    }

    private func page(_ state: ItemPageState, anchor: ItemPageAnchor = .item, specsExpanded: Bool = false) -> some View {
        Pushed(back: Copy.itemPageBack) {
            ItemPageView(state: state, initialAnchor: anchor, specsExpanded: specsExpanded)
                .itemPageBar(backTitle: Copy.itemPageBack)
        }
    }

    @Test func f12_4_pick() throws {
        snap(page(try capeState()), "f12.4-pick-ios", height: 1400)
    }

    @Test func f12_4_pick_ax5() throws {
        snap(page(try capeState()), "f12.4-pick-ax5-ios", height: 3800, large: true)
    }

    @Test func f12_4_runnerUp() throws {
        snap(page(try itemState(SearchTestData.runnerUp, uses: SearchTestData.uses(SearchTestData.runnerUp))), "f12.4-runner-up-ios", height: 1400)
    }

    @Test func f12_4_notOnList() throws {
        snap(page(try itemState(SearchTestData.notListed, uses: SearchTestData.uses(SearchTestData.notListed))), "f12.4-not-on-list-ios", height: 1400)
    }

    @Test func f12_4_setPiece() throws {
        snap(page(try itemState(SearchTestData.setPiece, uses: SearchTestData.uses(SearchTestData.setPiece))), "f12.4-set-piece-ios", height: 1300)
    }

    @Test func f12_4_vendor() throws {
        snap(page(try itemState(SearchTestData.vendorItem, uses: SearchTestData.uses(SearchTestData.vendorItem))), "f12.4-vendor-ios", height: 1200)
    }

    @Test func f12_5_bisFor() throws {
        snap(page(try capeState(), anchor: .bisFor), "f12.5-bis-for-ios")
    }

    @Test func f12_5_bisFor_expanded() throws {
        snap(page(try capeState(), specsExpanded: true), "f12.5-bis-for-expanded-ios", height: 1900)
    }

    @Test func f12_4_loading() throws {
        var state = try capeState()
        state.pick = nil
        state.reservesPick = true
        state.tooltip = .loading
        state.sources = .loading
        state.specs = .loading
        snap(page(state), "f12.4-loading-ios", height: 1400)
    }

    @Test func f12_4_usesFailed() throws {
        var state = try capeState()
        state.sources = .failed
        state.specs = .failed
        snap(page(state), "f12.4-uses-failed-ios", height: 1200)
    }

    @Test func f12_unknownItem() {
        var state = ItemPageState(route: ItemPageRoute(itemID: 999_999), build: build)
        state.isKnown = false
        snap(page(state), "f12.4-unknown-ios")
    }

    @Test func f2_6_moreAbout() throws {
        let corpus = try SearchTestData.corpus()
        let item = try #require(corpus.item(SearchTestData.cape))
        let request = ItemPageLogic.tooltipRequest(item: item, route: ItemPageRoute(itemID: item.id), build: build, character: SearchTestData.lightningbolltz, own: raid)
        let model = ItemTooltipModel.build(request, sources: try SearchTestData.tooltipSources())
        snap(ItemTooltipContent(request: request, phase: .loaded(model), moreAbout: {}), "f2.6-more-about-ios", height: 1100)
    }
}

/// The field and chips over a body, with a focus state of its own.
private struct SearchHost<Content: View>: View {
    @State var query: String
    @State var scope: SearchScope
    @ViewBuilder let content: () -> Content
    @FocusState private var focused: Bool

    var body: some View {
        SearchChrome(query: $query, scope: $scope, focused: $focused, cancel: {}, content: content)
    }
}

/// A screen as pushed, so the back control shows (root titled `back`).
private struct Pushed<Content: View>: View {
    let back: String
    @ViewBuilder let content: () -> Content

    var body: some View {
        NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(back).navigationDestination(for: Int.self) { _ in content() }
        }
    }
}
