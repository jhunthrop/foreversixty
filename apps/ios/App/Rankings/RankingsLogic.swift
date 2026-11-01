import Foundation
import Models

/// The four measures of a board (spec F11.2), by the API's `metric` word.
enum RankMetric: String, CaseIterable, Codable, Equatable, Hashable, Sendable, Identifiable {
    case dps
    case hps
    case damageTaken = "damage_taken"
    case execution

    var id: String { rawValue }

    /// The measure a role is judged on (defaults.ts `ROLE_METRIC`): healing for a healer, damage
    /// taken for a tank, damage for everyone else.
    static func forRole(_ role: String?) -> RankMetric {
        switch role {
        case "healer": .hps
        case "tank": .damageTaken
        default: .dps
        }
    }
}

/// What a board shows (spec F11.2). An empty string is "all" for that filter.
struct RankingsFilters: Equatable, Hashable, Codable, Sendable {
    var encounter: String
    var metric: RankMetric
    var ruleset: String
    var region: String
    var phase: String
    var faction: String
    var todayOnly: Bool
    /// The cohort: the character's own class slug and spec name; empty for every class.
    var classSlug: String
    var spec: String
}

/// The filters a viewer starts on (defaults.ts `rankingsDefaults`).
struct RankingsDefaults: Equatable, Sendable {
    let classSlug: String
    let spec: String
    let ruleset: String
    let phase: String
    let metric: RankMetric
}

/// The signed-in viewer's current character, as the board reads it.
struct RankingsCharacter: Equatable, Hashable, Sendable {
    let key: String
    let name: String
    /// The API's class word (`Shaman` or `shaman`).
    let className: String?
    /// The primary tree's name (`Elemental`).
    let spec: String?
    let ruleset: String
}

/// The viewer: signed in or not, every character of the account (the `You` tag), the current
/// one (the board's aim and your row) and `entitlements.history` (Parse history).
struct RankingsAccount: Equatable, Sendable {
    let signedIn: Bool
    let characterKeys: [String]
    let current: RankingsCharacter?
    let history: Bool

    static let signedOut = RankingsAccount(signedIn: false, characterKeys: [], current: nil, history: false)

    static func of(_ me: Me?, current: AccountCharacter?) -> RankingsAccount {
        guard let me else { return .signedOut }
        return RankingsAccount(
            signedIn: true,
            characterKeys: me.characters.map(\.key),
            current: current.map {
                RankingsCharacter(key: $0.key, name: $0.name, className: $0.className, spec: $0.spec, ruleset: $0.ruleset)
            },
            history: me.entitlements.history
        )
    }
}

/// Where the board starts (web/src/lib/rankings/defaults.ts and phases.ts).
enum RankingsDefaulting {
    /// The phase a moment falls in: the last phase whose start is not after it; empty when none.
    static func phaseAt(_ date: Date, phases: [PhaseBoundary]) -> String {
        phases.filter { $0.start <= date }.max { $0.start < $1.start }?.name ?? ""
    }

    /// The viewer's own class, spec, ruleset and role's measure, and the current phase. Signed
    /// out (`character` nil) nothing is aimed and the measure is damage.
    static func defaults(_ character: RankingsCharacter?, specs: SpecCatalog, phases: [PhaseBoundary], now: Date) -> RankingsDefaults {
        let phase = phaseAt(now, phases: phases)
        guard let character else {
            return RankingsDefaults(classSlug: "", spec: "", ruleset: "", phase: phase, metric: .dps)
        }
        let classSlug = character.className?.lowercased() ?? ""
        let known = classSlug.isEmpty ? nil : specs.spec(classSlug: classSlug, named: character.spec ?? "")
        return RankingsDefaults(
            classSlug: classSlug,
            spec: classSlug.isEmpty ? "" : known?.name ?? character.spec ?? "",
            ruleset: character.ruleset,
            phase: phase,
            metric: RankMetric.forRole(known?.role)
        )
    }

    /// The encounter the character most recently killed and so is ranked on, else the first one
    /// listed, else none. Only a listed encounter is named, so the slug is one the API accepts.
    static func defaultEncounter(history: [CharacterFight], encounters: [RankedEncounter]) -> String {
        for fight in history where fight.kill != false {
            if let listed = encounters.first(where: { $0.id == Int64(fight.encounterId) }) { return listed.slug }
        }
        return encounters.first?.slug ?? ""
    }

    static func filters(_ defaults: RankingsDefaults, encounter: String) -> RankingsFilters {
        RankingsFilters(
            encounter: encounter, metric: defaults.metric, ruleset: defaults.ruleset, region: "", phase: defaults.phase,
            faction: "", todayOnly: false, classSlug: defaults.classSlug, spec: defaults.spec
        )
    }
}

/// One `GET /v1/rankings` read: one best kill per character, placing `character` when given.
struct RankingsQuery: Equatable, Hashable, Sendable {
    static let distinctCharacter = "character"
    static let today = "today"

    let encounter: String
    let metric: String
    let spec: String?
    let classSlug: String?
    let phase: String?
    let region: String?
    let ruleset: String?
    let faction: String?
    let since: String?
    let distinct: String
    let character: String?
    let page: Int

    static func board(_ filters: RankingsFilters, character: String?, page: Int) -> RankingsQuery {
        RankingsQuery(
            encounter: filters.encounter, metric: filters.metric.rawValue, spec: named(filters.spec),
            classSlug: named(filters.classSlug), phase: named(filters.phase), region: named(filters.region),
            ruleset: named(filters.ruleset), faction: named(filters.faction), since: filters.todayOnly ? today : nil,
            distinct: distinctCharacter, character: character, page: page
        )
    }

    private static func named(_ value: String) -> String? { value.isEmpty ? nil : value }
}

/// A rank's colour on the parse ladder (row-view.ts `rankColor`).
enum RankColor {
    /// The legend's top: the first rank of a board is its 100th percentile and still reads pink.
    static let topLegendPercentile = 99.0

    /// First is 100, last is just above 0.
    static func percentile(rank: Int, total: Int) -> Double {
        100 - Double(rank - 1) / Double(max(total, 1)) * 100
    }

    static func tier(rank: Int, total: Int) -> ParseTier {
        let capped = min(topLegendPercentile, max(0, percentile(rank: rank, total: total)))
        return ParseTier(Int(capped.rounded()))
    }
}

/// Where the viewer's own row goes (row-view.ts `ownRowPlacement`): marked where it already is,
/// pinned over the board when it is off the loaded rows, or nowhere. Carries the true rank.
enum OwnRow: Equatable, Sendable {
    case inPlace(Int)
    case pinned(Int)
    case none

    static func placement(rows: [RankingRow], characterRank: CharacterRank?) -> OwnRow {
        guard let placed = characterRank else { return .none }
        return rows.contains { $0.player.key == placed.row.player.key } ? .inPlace(placed.rank) : .pinned(placed.rank)
    }
}

/// The board's per-row rules.
enum RankingsBoard {
    /// A board spanning more than one class draws a crest per row; one class draws none.
    static func showsCrests(_ rows: [RankingRow]) -> Bool {
        Set(rows.map { $0.player.className?.lowercased() ?? "" }).count > 1
    }

    /// The `You` tag: a row of any of the account's characters.
    static func isYours(_ row: RankingRow, keys: Set<String>) -> Bool {
        keys.contains(row.player.key)
    }
}

/// Paging: 100 a page, merged by rank so a page read twice adds nothing.
enum RankingsPaging {
    static let perPage = 100

    static func merge(_ rows: [RankingRow], _ next: [RankingRow]) -> [RankingRow] {
        var seen = Set(rows.map(\.rank))
        return rows + next.filter { seen.insert($0.rank).inserted }
    }
}

/// A character named by its key (`us/normal/voltaris`): the path the character page and the
/// watch calls take, and the web link a share carries.
struct CharacterRef: Hashable, Identifiable, Sendable {
    static let site = "https://foreversixty.gg"

    let region: String
    let ruleset: String
    /// The key's last part, the name as the URL spells it.
    let slug: String
    let name: String

    var id: String { key }
    var key: String { "\(region)/\(ruleset)/\(slug)" }

    init(region: String, ruleset: String, slug: String, name: String) {
        self.region = region
        self.ruleset = ruleset
        self.slug = slug
        self.name = name
    }

    /// Nil for a key that is not `region/ruleset/name`.
    init?(key: String, name: String) {
        let parts = key.split(separator: "/").map(String.init)
        guard parts.count == 3, parts.allSatisfy({ !$0.isEmpty }) else { return nil }
        self.init(region: parts[0], ruleset: parts[1], slug: parts[2], name: name)
    }

    /// `https://foreversixty.gg/character/us/normal/voltaris`.
    var webURL: URL {
        URL(string: "\(Self.site)/character/\(key)") ?? URL(fileURLWithPath: "/")
    }

    /// The web character page's parse history (`#history`).
    var historyURL: URL {
        URL(string: "\(webURL.absoluteString)#history") ?? webURL
    }
}
