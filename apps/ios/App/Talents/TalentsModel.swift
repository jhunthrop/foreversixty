import Foundation
import Models
import Networking

/// The character the Talents screen is about: Home's current character (spec F6 §3).
struct TalentCharacter: Equatable, Sendable {
    let classSlug: String
    /// The primary tree's name as the API gives it (`Fury`); nil when unknown.
    let specName: String?
    let level: Int?
    let faction: String?
    let race: String?
    /// The export's per-tree rank strings; nil without an export.
    let trees: [String]?
    let dataBuild: String?

    init(_ character: HomeCharacter) {
        classSlug = character.className ?? ""
        specName = character.specName.flatMap { $0.isEmpty ? nil : $0 }
        level = character.level
        faction = character.faction
        race = character.race
        // An export whose points are all 0 holds no ranks (the site's `exportCodeFromAccount`).
        let points = character.talents?.points ?? []
        trees = points.contains { $0 > 0 } ? character.talents?.trees : nil
        dataBuild = character.dataBuild
    }
}

/// The screen's states (spec F6 §6).
enum TalentsPhase: Equatable, Sendable {
    case loading
    /// What to draw; `offlineSince` is when the oldest saved copy shown was fetched.
    case loaded(TalentsPresenter, offlineSince: Date?)
    /// The talent file is neither saved nor fetchable.
    case failed
}

/// Loads the files behind the Talents screen and turns them into a presenter.
@MainActor
@Observable
final class TalentsModel {
    private(set) var phase: TalentsPhase = .loading
    private(set) var build: String?

    private let source: any TalentsSource
    private let character: TalentCharacter
    private let catalog: SpecCatalog

    init(source: any TalentsSource, character: TalentCharacter, catalog: SpecCatalog) {
        self.source = source
        self.character = character
        self.catalog = catalog
    }

    /// The first load, and Try again: a refresh never puts skeletons over content.
    func load() async {
        if case .failed = phase { phase = .loading }
        let build = await source.activeBuild()
        self.build = build
        let file: CachedFile<TalentsFile>
        do {
            file = try await source.talents(build: build, classSlug: character.classSlug)
        } catch {
            phase = .failed
            return
        }
        let guideResult = await guideBuild(build: build)
        let exportResult = await export(active: file.value)
        let saved = [file.savedAtIfCached, guideResult.savedAt, exportResult.savedAt].compactMap(\.self)
        let presenter = TalentsPresenter(
            file: file.value, guide: guideResult.value, export: exportResult.value,
            level: character.level, specName: character.specName ?? ""
        )
        phase = .loaded(presenter, offlineSince: saved.min())
    }

    /// The guide build of the character's band (spec F6 §4.2).
    private func guideBuild(build: String) async -> Saved<TalentGuideBuild> {
        guard let specName = character.specName,
              let spec = catalog.spec(classSlug: character.classSlug, named: specName)
        else { return Saved(.noSpec) }
        guard let band = TalentGuide.band(forLevel: character.level) else { return Saved(.none) }
        do {
            let file = try await source.bis(build: build, spec: spec.spec)
            let entry = TalentGuide.entry(in: file.value, band: band, faction: character.faction)
            return Saved(entry.map { .build(talents: $0.talents, band: band) } ?? .none, savedAt: file.savedAtIfCached)
        } catch {
            return Saved(error.isMissingFile ? .none : .failed)
        }
    }

    /// The export, read on the active trees only when its own build lays them out the same.
    private func export(active: TalentsFile) async -> Saved<TalentExport> {
        guard let trees = character.trees else { return Saved(.none) }
        let ranks = TalentExport.ranks(trees.map { TalentRules.ranks(ofTree: $0) })
        guard let dataBuild = character.dataBuild, dataBuild != active.build else { return Saved(ranks) }
        guard let older = try? await source.talents(build: dataBuild, classSlug: character.classSlug) else {
            return Saved(.stale)
        }
        return Saved(TalentGuide.sameLayout(older.value, active) ? ranks : .stale, savedAt: older.savedAtIfCached)
    }
}

/// A value and, when it came from the offline copy, when that copy was fetched.
private struct Saved<Value: Sendable>: Sendable {
    let value: Value
    var savedAt: Date?

    init(_ value: Value, savedAt: Date? = nil) {
        self.value = value
        self.savedAt = savedAt
    }
}

extension CachedFile {
    /// When the saved copy was fetched, or nil for a fresh answer.
    var savedAtIfCached: Date? { isFromCache ? savedAt : nil }
}

extension APIError {
    /// A file the CDN does not have (404), or fixture mode has no fixture for: an empty state.
    var isMissingFile: Bool {
        switch self {
        case .unexpectedStatus(404), .missingFixture: true
        case let .server(status, _, _): status == 404
        default: false
        }
    }
}
