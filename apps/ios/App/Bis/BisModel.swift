import Foundation
import Models
import Networking
import Observation

/// What the list region shows (spec F3.1 and its states).
enum BisPhase: Equatable, Sendable {
    /// First load with nothing saved: 17 skeleton rows.
    case loading
    /// The list; `offlineSince` is when the saved copy was fetched, shown in the offline strip.
    case loaded(BisList, offlineSince: Date?)
    /// The file has no list for this range and faction.
    case noList(offlineSince: Date?)
    /// Offline, and this list was never fetched (F3.3).
    case notCached
    /// The CDN answered with an error.
    case failed
}

/// The picker's choices and the character the tab belongs to.
struct BisContext: Equatable, Sendable {
    var selection: BisSelection
    /// The current character's own list; nil signed out or when its spec is unknown.
    var characterSelection: BisSelection?
    var characterName: String?
    var classes: [ClassInfo] = []
    /// When the first raids open; nil when the dates file has no raids entry.
    var raidsOpening: RaidsOpening?

    /// The gold `Yours` tag: the list is the character's spec, range and faction.
    var isYours: Bool { characterSelection == selection }

    /// `Back to {name}'s list` shows only when the picks differ from the character's.
    var backTarget: (name: String, selection: BisSelection)? {
        guard let own = characterSelection, let characterName, own != selection else { return nil }
        return (characterName, own)
    }

    func className(_ slug: String) -> String {
        classes.first { $0.slug == slug }?.name ?? slug.capitalized
    }

    /// A range's chip and pill words, with `opens Dec 9` on the raid list before raids open.
    func rangeLabel(_ range: BisRange, now: Date = .now) -> String {
        BisLabels.range(range, opening: raidsOpening, now: now)
    }
}

/// Remembers the last list and whether the signed-out picker has opened once.
protocol BisPreferences: Sendable {
    var lastSelection: BisSelection? { get nonmutating set }
    var hasOpenedPicker: Bool { get nonmutating set }
}

struct UserDefaultsBisPreferences: BisPreferences {
    private static let selectionKey = "gg.foreversixty.bis.selection"
    private static let pickerKey = "gg.foreversixty.bis.pickerOpened"
    // UserDefaults is documented thread-safe.
    private nonisolated(unsafe) let defaults: UserDefaults

    init(defaults: UserDefaults = .standard) {
        self.defaults = defaults
    }

    var lastSelection: BisSelection? {
        get { defaults.data(forKey: Self.selectionKey).flatMap { try? JSONDecoder().decode(BisSelection.self, from: $0) } }
        nonmutating set { defaults.set(newValue.flatMap { try? JSONEncoder().encode($0) }, forKey: Self.selectionKey) }
    }

    var hasOpenedPicker: Bool {
        get { defaults.bool(forKey: Self.pickerKey) }
        nonmutating set { defaults.set(newValue, forKey: Self.pickerKey) }
    }
}

/// The BiS tab's state: the selection, the published list for it, and its offline copy.
@MainActor
@Observable
final class BisModel {
    private(set) var phase: BisPhase = .loading
    private(set) var context: BisContext
    var isPickerPresented = false
    var searchText = ""

    let catalog: SpecCatalog
    private let source: any BisDataSource
    private let preferences: any BisPreferences
    private(set) var build: String?
    private var worn: Set<Int> = []
    private var generation = 0

    init(source: any BisDataSource, catalog: SpecCatalog, preferences: any BisPreferences) {
        self.source = source
        self.catalog = catalog
        self.preferences = preferences
        context = BisContext(selection: preferences.lastSelection ?? .signedOutDefault(catalog))
    }

    /// The first load: the character, then the saved list at once, then the network.
    func start() async {
        let character = await source.currentCharacter()
        let own = character.flatMap { BisSelection.forCharacter($0, catalog: catalog) }
        worn = Set(character?.build?.gear.map { Array($0.values) } ?? [])
        context.characterSelection = own
        context.characterName = own == nil ? nil : character?.name
        if let own {
            context.selection = own
        } else if character == nil, !preferences.hasOpenedPicker {
            preferences.hasOpenedPicker = true
            isPickerPresented = true
        }
        let build = await source.activeBuild()
        self.build = build
        context.raidsOpening = await RaidsOpening(source.keyDates())
        drawSaved(build: build)
        await load()
    }

    /// Lane F15: a newer build found while the list is open; `Update` applies it.
    private(set) var newData: NewDataState = .none
    /// Lane F15: the CDN's active build now; nil keeps the strip away (tests, offline).
    var latestBuild: (@Sendable () async -> String?)?

    /// On foreground: a newer active build shows the strip and changes nothing else.
    func checkForNewData() async {
        guard let build, let latestBuild else { return }
        newData = NewDataCheck.state(shown: build, latest: await latestBuild())
    }

    /// `Update`: the list of the new build replaces the open one.
    func applyNewData() async {
        guard case let .ready(next) = newData else { return }
        newData = .none
        build = next
        drawSaved(build: next)
        await load()
    }

    /// Pull to refresh and `Try again`: refetch the current list, keeping what is on screen.
    func refresh() async {
        await load()
    }

    /// A picker choice: applies at once and the list redraws under the sheet.
    func select(_ selection: BisSelection) async {
        guard selection != context.selection else { return }
        context.selection = selection
        preferences.lastSelection = selection
        if let build { drawSaved(build: build) }
        await load()
    }

    private func drawSaved(build: String) {
        let selection = context.selection
        guard let file = source.cachedBis(build: build, spec: selection.spec) else {
            phase = .loading
            return
        }
        if let classes = source.cachedClasses(build: build) { context.classes = classes.value }
        let items = source.cachedItems(build: build, classSlug: selection.classSlug)?.value.items ?? []
        phase = present(file.value, items: items, offlineSince: nil)
    }

    private func load() async {
        generation += 1
        let mine = generation
        let selection = context.selection
        let build: String
        if let known = self.build {
            build = known
        } else {
            build = await source.activeBuild()
            self.build = build
        }
        async let bis = attempt { () async throws(APIError) in try await self.source.bis(build: build, spec: selection.spec) }
        async let items = attempt { () async throws(APIError) in
            try await self.source.items(build: build, classSlug: selection.classSlug)
        }
        async let classes = attempt { () async throws(APIError) in try await self.source.classes(build: build) }
        let (bisResult, itemsResult, classesResult) = await (bis, items, classes)
        guard mine == generation else { return }
        if case let .success(answer) = classesResult { context.classes = answer.value }
        phase = phase(bis: bisResult, items: itemsResult)
    }

    private func phase(bis: Result<CachedFile<BisFile>, APIError>, items: Result<CachedFile<ItemsFile>, APIError>) -> BisPhase {
        switch bis {
        case let .failure(error):
            if case .loaded = phase { return phase }
            if case .transport = error { return .notCached }
            return .failed
        case let .success(file):
            let itemsFile = try? items.get()
            let itemsSavedAt = itemsFile?.isFromCache == true ? itemsFile?.savedAt : nil
            let offlineSince = file.isFromCache ? [file.savedAt, itemsSavedAt].compactMap(\.self).min() : nil
            return present(file.value, items: itemsFile?.value.items ?? [], offlineSince: offlineSince)
        }
    }

    private func present(_ file: BisFile, items: [Item], offlineSince: Date?) -> BisPhase {
        let selection = context.selection
        guard let band = file.band(for: selection.range, faction: selection.faction) else {
            return .noList(offlineSince: offlineSince)
        }
        let list = BisListBuilder(items: items, worn: worn).list(
            file: file,
            band: band,
            specLabel: specLabel(selection),
            rangeLabel: BisLabels.range(selection.range)
        )
        return .loaded(list, offlineSince: offlineSince)
    }

    func specLabel(_ selection: BisSelection) -> String {
        let spec = catalog.spec(selection.spec)?.name ?? selection.spec
        return Copy.bisSpecChip(spec: spec, className: context.className(selection.classSlug))
    }

    private func attempt<T: Sendable>(
        _ operation: @Sendable () async throws(APIError) -> T
    ) async -> Result<T, APIError> {
        do {
            return try await .success(operation())
        } catch {
            return .failure(error)
        }
    }
}

/// The words of the chips and pills.
enum BisLabels {
    /// The range's words; the raid list reads `60 · Phase 1 raid · opens Dec 9` until the first
    /// raids open, with the date from the data.
    static func range(_ range: BisRange, opening: RaidsOpening? = nil, now: Date = .now) -> String {
        if range.preset == BisRange.raidPreset {
            if let opening, opening.isUpcoming(at: now) { return Copy.bisRangeP1Opens(date: opening.label) }
            return Copy.bisRangeP1
        }
        return band(range)
    }

    /// The band's own words, whatever the preset: `50 to 59`, or `60` for the last band.
    static func band(_ range: BisRange) -> String {
        if range.band == BisRange.lastBand { return Copy.bisRangeTop(n: range.band) }
        return Copy.bisRangeSpan(from: range.band, to: range.band + BisRange.bandStep - 1)
    }

    /// Home's band head: `60 · raid` for the raid list, the band's words otherwise.
    static func head(_ range: BisRange) -> String {
        range.preset == BisRange.raidPreset ? Copy.homeUpgradesBandRaid(band: band(range)) : band(range)
    }

    static func faction(_ slug: String) -> String {
        slug == "horde" ? Copy.bisFactionHorde : Copy.bisFactionAlliance
    }

    static func role(_ role: String) -> String {
        switch role {
        case "healer": Copy.bisRoleHealer
        case "tank": Copy.bisRoleTank
        default: Copy.bisRoleDps
        }
    }
}
