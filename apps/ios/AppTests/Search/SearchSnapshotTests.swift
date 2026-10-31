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

    // MARK: The item page

    private func capeState(viewer: BisSelection?, className: String = "shaman") throws -> ItemPageState {
        let corpus = try SearchTestData.corpus()
        let item = try #require(corpus.item(SearchTestData.cape))
        var state = ItemPageState(route: ItemPageRoute(itemID: item.id), build: build)
        state.look = HomeItemLook(id: item.id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel)
        state.meta = ItemPageLogic.meta(item)
        let character = viewer == nil ? SearchTestData.fury : SearchTestData.lightningbolltz
        let request = ItemPageLogic.tooltipRequest(item: item, route: ItemPageRoute(itemID: item.id), build: build, character: character, own: viewer)
        let sources = try SearchTestData.tooltipSources(classSlug: request.classSlug)
        state.tooltip = .loaded(.build(request, sources: sources))
        if let viewer {
            let band = sources.bis?.bands.first { $0.band == viewer.range.band && $0.preset == viewer.range.preset && $0.faction == viewer.faction }
            state.pick = ItemPageLogic.yourPick(itemID: item.id, band: band, selection: viewer, specName: "Elemental", className: className)
        }
        let uses = try SearchTestData.capeUses()
        state.sources = .loaded(ItemPageLogic.sources(uses.sources))
        state.specs = .loaded(ItemPageLogic.specs(uses.specs, catalog: try BisTestData.catalog()))
        return state
    }

    private let raid = BisSelection(classSlug: "shaman", spec: "shaman-elemental", range: BisRange(band: 60, preset: "raid"), faction: "horde")

    private func page(_ state: ItemPageState, specsExpanded: Bool = false) -> some View {
        Pushed(back: Copy.itemPageBack) {
            ItemPageView(state: state, specsExpanded: specsExpanded)
                .navigationTitle("")
                .navigationBarTitleDisplayMode(.inline)
        }
    }

    @Test func f12_4_pick() throws {
        snap(page(try capeState(viewer: raid)), "f12.4-pick-ios", height: 1400)
    }

    @Test func f12_4_pick_ax5() throws {
        snap(page(try capeState(viewer: raid)), "f12.4-pick-ax5-ios", height: 3800, large: true)
    }

    @Test func f12_4_runnerUp() throws {
        snap(page(try capeState(viewer: raid.withRange(BisRange(band: 60, preset: "bare")))), "f12.4-runner-up-ios", height: 1400)
    }

    @Test func f12_4_notOnList() throws {
        snap(page(try capeState(viewer: nil)), "f12.4-not-on-list-ios", height: 1400)
    }

    @Test func f12_5_bisFor() throws {
        // Every spec that lists the Cape, after `5 more specs` (F12.5's section in full).
        snap(page(try capeState(viewer: raid), specsExpanded: true), "f12.5-bis-for-ios", height: 1900)
    }

    @Test func f12_4_loading() throws {
        var state = try capeState(viewer: raid)
        state.pick = nil
        state.tooltip = .loading
        state.sources = .loading
        state.specs = .loading
        snap(page(state), "f12.4-loading-ios")
    }

    @Test func f12_4_usesFailed() throws {
        var state = try capeState(viewer: raid)
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
