import Foundation
import Models
import Networking

/// One spec of the bundled catalog (data/curated/specs.json, copied into the app: the CDN does
/// not publish it). Only the fields the picker reads; the file is the source of truth.
struct SpecInfo: Decodable, Equatable, Sendable, Identifiable {
    /// The spec key, `shaman-elemental`: the BiS file's name.
    let spec: String
    let classSlug: String
    /// The spec's display name, `Elemental`.
    let name: String
    /// `dps`, `healer` or `tank`.
    let role: String
    /// The icon stem on the CDN (`/data/<build>/icons/<icon>.webp`).
    let icon: String

    var id: String { spec }
}

/// The specs of every class, in the catalog's order.
struct SpecCatalog: Equatable, Sendable {
    let specs: [SpecInfo]

    static let resourceName = "specs"

    /// The copy bundled with the app (project.yml copies data/curated/specs.json in). Empty only
    /// for a broken build, which the bundle test catches; never a crash on a user's phone.
    static func bundled(_ bundle: Bundle = .main) -> SpecCatalog {
        guard let url = bundle.url(forResource: resourceName, withExtension: "json"),
              let data = try? Data(contentsOf: url),
              let catalog = try? decode(data)
        else { return SpecCatalog(specs: []) }
        return catalog
    }

    static func decode(_ data: Data) throws -> SpecCatalog {
        try SpecCatalog(specs: WireFormat.decoder().decode([SpecInfo].self, from: data))
    }

    func specs(of classSlug: String) -> [SpecInfo] {
        specs.filter { $0.classSlug == classSlug }
    }

    func spec(_ key: String) -> SpecInfo? {
        specs.first { $0.spec == key }
    }

    /// The spec a character names by its class slug and tree name (`warrior`, `Fury`).
    func spec(classSlug: String, named name: String) -> SpecInfo? {
        specs.first { $0.classSlug == classSlug && $0.name.caseInsensitiveCompare(name) == .orderedSame }
    }
}

/// A level range of the picker: a band of the BiS file and its preset (spec F3.2).
struct BisRange: Hashable, Sendable, Codable {
    let band: Int
    let preset: String

    static let barePreset = "bare"
    static let raidPreset = "raid"
    static let firstBand = 20
    static let lastBand = 60
    static let bandStep = 10

    /// `20 to 29` through `60`, then `60 · Phase 1 raid`.
    static let all: [BisRange] = stride(from: firstBand, through: lastBand, by: bandStep)
        .map { BisRange(band: $0, preset: barePreset) } + [BisRange(band: lastBand, preset: raidPreset)]

    /// The site's band rule (`bandForLevel`): rounded down to the step and clamped to 20 to 60.
    /// A level-60 character reads the raid list, the site's headline preset.
    static func forLevel(_ level: Int) -> BisRange {
        let clamped = min(lastBand, max(firstBand, level))
        let band = clamped / bandStep * bandStep
        return BisRange(band: band, preset: band == lastBand ? raidPreset : barePreset)
    }
}

/// Which list the BiS tab shows.
struct BisSelection: Hashable, Sendable, Codable {
    let classSlug: String
    let spec: String
    let range: BisRange
    let faction: String

    static let factions = ["alliance", "horde"]

    /// The first visit, signed out: Warrior preselected (spec F3.2), the headline list.
    static func signedOutDefault(_ catalog: SpecCatalog) -> BisSelection {
        let warrior = "warrior"
        return BisSelection(
            classSlug: warrior,
            spec: catalog.specs(of: warrior).first?.spec ?? "warrior-arms",
            range: BisRange.all.last ?? .forLevel(BisRange.lastBand),
            faction: factions[0]
        )
    }

    /// The character's own list, when the catalog knows its class and spec.
    static func forCharacter(_ character: AccountCharacter, catalog: SpecCatalog) -> BisSelection? {
        guard let classSlug = character.className, let specName = character.spec,
              let spec = catalog.spec(classSlug: classSlug, named: specName),
              let level = character.level,
              let faction = character.faction, factions.contains(faction)
        else { return nil }
        return BisSelection(classSlug: classSlug, spec: spec.spec, range: .forLevel(level), faction: faction)
    }

    func withClass(_ slug: String, catalog: SpecCatalog) -> BisSelection {
        let first = catalog.specs(of: slug).first?.spec ?? spec
        return BisSelection(classSlug: slug, spec: first, range: range, faction: faction)
    }

    func withSpec(_ key: String) -> BisSelection {
        BisSelection(classSlug: classSlug, spec: key, range: range, faction: faction)
    }

    func withRange(_ range: BisRange) -> BisSelection {
        BisSelection(classSlug: classSlug, spec: spec, range: range, faction: faction)
    }

    func withFaction(_ faction: String) -> BisSelection {
        BisSelection(classSlug: classSlug, spec: spec, range: range, faction: faction)
    }
}

extension BisFile {
    /// The band the site shows for a range and faction (`presets.ts` `selectBand`): the asked
    /// preset, else the bare list. Nothing is re-ranked; the file's own entry is returned.
    func band(for range: BisRange, faction: String) -> BisBand? {
        let candidates = bands.filter { $0.band == range.band && $0.faction == faction }
        return candidates.first { $0.preset == range.preset }
            ?? candidates.first { $0.preset == BisRange.barePreset }
    }
}
