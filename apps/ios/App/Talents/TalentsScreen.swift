import DesignSystem
import DesignTokens
import Networking
import SwiftUI

/// The Talents screen pushed from Home's Talents row (spec F6 §3): the model, the chosen build
/// and tree, and the rank sheet.
struct TalentsScreen: View {
    let character: HomeCharacter

    @State private var model: TalentsModel
    @State private var shown: TalentBuildKind = .yours
    @State private var tree = 0
    @State private var sheet: TalentSheetModel?
    @State private var settled = false
    @Environment(\.openURL) private var openURL

    /// [model] is Home's, already loading for the Talents row, so the push draws at once.
    init(character: HomeCharacter, model: TalentsModel) {
        self.character = character
        _model = State(initialValue: model)
    }

    var body: some View {
        TalentsView(
            phase: model.phase,
            build: model.build ?? "",
            offlineStamp: offlineStamp,
            shown: $shown,
            tree: $tree,
            selectedID: sheet?.id,
            openTalent: openTalent,
            retry: { Task { await model.load() } },
            openPlanner: { openURL(TalentLinks.planner(classSlug: character.className, race: character.race)) }
        )
        .navigationTitle(Copy.talentsTitle)
        .navigationBarTitleDisplayMode(.inline)
        .toolbarBackground(FSColor.bg, for: .navigationBar)
        .toolbarBackground(.visible, for: .navigationBar)
        .sheet(item: $sheet) { TalentRankSheet(model: $0, build: model.build ?? "") }
        .task {
            if model.phase.presenter == nil { await model.load() }
            settle()
        }
    }

    /// Opens on Your build when it can be drawn and on its fullest tree, once (spec F6 §5.3, §5.4).
    private func settle() {
        guard !settled, case let .loaded(presenter, _) = model.phase else { return }
        settled = true
        shown = presenter.initialBuild
        tree = presenter.openingTree
    }

    private func openTalent(_ id: Int) {
        guard case let .loaded(presenter, _) = model.phase else { return }
        sheet = presenter.sheet(shown, talentID: id)
    }

    private var offlineStamp: String? {
        guard case let .loaded(_, since?) = model.phase else { return nil }
        return OfflineStamp.format(since)
    }
}

/// The planner on the site, for the character's class and race (spec F6 §5.1 item 8).
enum TalentLinks {
    static func planner(classSlug: String?, race: String?) -> URL {
        var components = URLComponents(url: HomeLinks.planner, resolvingAgainstBaseURL: false)
        let items = [
            classSlug.map { URLQueryItem(name: "class", value: $0) },
            race.map { URLQueryItem(name: "race", value: raceSlug($0)) },
        ].compactMap(\.self)
        components?.queryItems = items.isEmpty ? nil : items
        return components?.url ?? HomeLinks.planner
    }

    /// "Night Elf" to "night-elf", the planner's race slug (`account-character.ts` `raceSlugOf`).
    static func raceSlug(_ race: String) -> String {
        race.trimmingCharacters(in: .whitespaces).lowercased().split(whereSeparator: \.isWhitespace).joined(separator: "-")
    }
}

extension LiveTalentsSource {
    /// The live source over the same CDN, offline cache and active build the BiS tab reads.
    @MainActor
    static func make(_ environment: AppEnvironment) -> LiveTalentsSource {
        let transport: any HTTPTransport = switch environment.configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled() ?? FixtureTransport(directory: Bundle.main.bundleURL)
        }
        let files = CachedStaticDataClient(client: StaticDataClient(transport: transport), cache: DiskStaticFileCache.applicationSupport())
        let bis = LiveBisDataSource(files: files, appConfig: environment.appConfig, character: { nil }, bundledBuild: BundledBuild.read())
        return LiveTalentsSource(files: files, bisSource: bis)
    }
}
