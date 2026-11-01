import Foundation
import Models
import Networking
import Observation

/// One region of the item page, which loads and fails on its own (spec §5.0).
enum ItemPageSection<Value: Equatable & Sendable>: Equatable, Sendable {
    case loading
    case loaded(Value)
    case failed
}

/// Everything the item page draws, pure so the snapshot tests render it directly.
struct ItemPageState: Equatable, Sendable {
    /// The head: the index's name, icon and quality, else the caller's look.
    var look: HomeItemLook
    var meta: String?
    /// False once the index is read and does not list the item (an old link).
    var isKnown = true
    var pick: YourPickCard?
    /// The character has a list, so the pick card's slot is held while the page loads.
    var reservesPick = false
    var tooltip: ItemPageSection<ItemTooltipModel> = .loading
    var sources: ItemPageSection<[ItemSourceRow]> = .loading
    var specs: ItemPageSection<[ItemSpecRow]> = .loading
    var build: String

    init(route: ItemPageRoute, build: String = "") {
        look = route.fallback ?? HomeItemLook(id: route.itemID, name: nil, quality: nil, icon: nil, itemLevel: nil)
        self.build = build
    }
}

/// Loads the item page: the index (head and shard), then the tooltip files and the item's uses
/// shard side by side. The character's own list decides the pick card.
@MainActor
@Observable
final class ItemPageStore {
    private(set) var state: ItemPageState
    /// Lane F15: a newer build found while the page is open; it waits for `Update`.
    private(set) var newData: NewDataState = .none
    /// The build `Update` moved the page to, ahead of the app's own read.
    private var pinnedBuild: String?
    private let route: ItemPageRoute
    private let dependencies: SearchDependencies

    init(route: ItemPageRoute, dependencies: SearchDependencies) {
        self.route = route
        self.dependencies = dependencies
        state = ItemPageState(route: route)
    }

    func load() async {
        state = ItemPageState(route: route, build: state.build)
        let read: String? = if let pinnedBuild { pinnedBuild } else { await dependencies.build() }
        guard let build = read else { return fail() }
        state.build = build
        guard let corpus = try? await dependencies.index.corpus(build: build) else { return fail() }
        guard let item = corpus.item(route.itemID) else {
            state.isKnown = false
            return
        }
        state.look = HomeItemLook(id: item.id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel)
        state.meta = ItemPageLogic.meta(item)
        async let tooltip: Void = loadTooltip(item, build: build)
        async let uses: Void = loadUses(item, corpus: corpus)
        _ = await (tooltip, uses)
    }

    /// On open and on foreground: a newer active build shows the strip; nothing else moves.
    func checkForNewData() async {
        guard !state.build.isEmpty else { return }
        newData = NewDataCheck.state(shown: state.build, latest: await dependencies.latestBuild())
    }

    /// `Update`: the page reads the new build.
    func applyNewData() async {
        guard case let .ready(next) = newData else { return }
        newData = .none
        pinnedBuild = next
        await load()
    }

    private func fail() {
        state.tooltip = .failed
        state.sources = .failed
        state.specs = .failed
    }

    /// The F2.6 panel from the items file, and the pick card from the character's own BiS list.
    private func loadTooltip(_ item: SearchItem, build: String) async {
        let character = dependencies.context().character
        let own = character.flatMap { BisSelection.forCharacter($0, catalog: dependencies.specs) }
        state.reservesPick = own != nil
        let request = ItemPageLogic.tooltipRequest(item: item, route: route, build: build, character: character, own: own)
        do {
            let sources = try await dependencies.catalog.sources(for: request)
            state.tooltip = .loaded(.build(request, sources: sources))
            state.reservesPick = false
            state.pick = ItemPageLogic.yourPick(
                itemID: item.id, band: ItemTooltipModel.band(for: request, in: sources.bis), selection: own,
                specName: own.flatMap { dependencies.specs.spec($0.spec)?.name }, className: character?.className
            )
        } catch {
            state.reservesPick = false
            state.tooltip = .failed
        }
    }

    private func loadUses(_ item: SearchItem, corpus: SearchCorpus) async {
        do {
            let uses = try await dependencies.index.uses(of: item.id, in: corpus)
            state.sources = .loaded(ItemPageLogic.sources(uses?.sources ?? []))
            state.specs = .loaded(ItemPageLogic.specs(uses?.specs ?? [], catalog: dependencies.specs))
        } catch {
            state.sources = .failed
            state.specs = .failed
        }
    }
}

extension ItemPageLogic {
    /// The tooltip's request: the items file of the character's class when it can use the item,
    /// else the caller's class or the first that can; the BiS file of the character's own list.
    static func tooltipRequest(
        item: SearchItem, route: ItemPageRoute, build: String, character: AccountCharacter?, own: BisSelection?
    ) -> ItemTooltipRequest {
        let mine = character?.className.flatMap { item.classes.contains($0) ? $0 : nil }
        let classSlug = mine ?? route.classSlug.flatMap { item.classes.contains($0) ? $0 : nil } ?? item.classes.first ?? ""
        return ItemTooltipRequest(
            itemID: item.id,
            fallback: HomeItemLook(id: item.id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel),
            build: build, classSlug: classSlug, spec: own?.spec, specName: nil,
            band: own?.range.band, bandLabel: nil, preset: own?.range.preset, faction: own?.faction,
            sourceLine: nil, wornID: nil
        )
    }
}
