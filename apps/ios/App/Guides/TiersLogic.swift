import Foundation
import Models

/// The tier list's role chips (spec F9.4).
enum TierRole: String, CaseIterable, Identifiable, Sendable {
    case dps, tank, healer

    var id: String { rawValue }
    var title: String { GuideRole.label(rawValue) }
}

/// The tier list's faction chip.
enum TierFaction: String, CaseIterable, Identifiable, Sendable {
    case alliance, horde

    var id: String { rawValue }

    var title: String {
        switch self {
        case .alliance: Copy.tiersFactionAlliance
        case .horde: Copy.tiersFactionHorde
        }
    }
}

/// Which list shows: a role and a faction.
struct TierChoice: Equatable, Sendable {
    let role: TierRole
    let faction: TierFaction

    static let fallback = TierChoice(role: .dps, faction: .horde)

    /// The current character's role and faction; with no character, DPS and Horde.
    static func defaults(_ character: AccountCharacter?, specs: SpecCatalog) -> TierChoice {
        let role = TiersLogic.specKey(character, specs: specs).flatMap { specs.spec($0) }.flatMap { TierRole(rawValue: $0.role) }
        let faction = character?.faction.flatMap(TierFaction.init(rawValue:))
        return TierChoice(role: role ?? fallback.role, faction: faction ?? fallback.faction)
    }

    func with(role: TierRole) -> TierChoice { TierChoice(role: role, faction: faction) }
    func with(faction: TierFaction) -> TierChoice { TierChoice(role: role, faction: faction) }
}

/// One tier row as drawn: the file's own words, never re-ranked (spec F9.4).
struct TierRowModel: Identifiable, Equatable, Sendable {
    let row: TierRow
    let count: Int
    let isYours: Bool

    var id: String { row.spec }
    /// `S` to `D` on a DPS list; nil on tank and healer lists.
    var letter: String? { row.tier }
    /// The file's gap, with `≈` before it on a tie.
    var gap: String { row.tie == nil ? row.gapText : "\(Copy.tiersTie) \(row.gapText)" }
    /// The tank and healer lists' right column.
    var metric: String { row.metricText }

    /// `Fury Warrior, rank 1 of 20, tier S, top spec, Troll` (spec §6).
    var a11y: String {
        let tier = letter.map { Copy.tiersA11yTier(letter: $0) } ?? ""
        let base = row.rank == 1
            ? Copy.tiersA11yTop(name: row.name, count: "\(count)", tier: tier, race: row.race)
            : Copy.tiersA11yBehind(
                name: row.name, rank: "\(row.rank)", count: "\(count)", tier: tier,
                gap: Copy.tiersA11yGap(n: Self.spokenPercent(row.gapPercent)), race: row.race
            )
        return row.tie == nil ? base : [base, Copy.tiersA11yTie].joined(separator: GuideCopy.listSeparator)
    }

    private static func spokenPercent(_ value: Double) -> String {
        value.formatted(.number.precision(.fractionLength(1)).locale(Locale(identifier: "en_US")))
    }
}

/// The callout over the list: the character's row and the list's note.
struct TierCallout: Equatable, Sendable {
    let row: TierRow
    /// The list's last `notes` line.
    let note: String?
}

enum TiersLogic {
    static func list(_ file: TiersFile, choice: TierChoice) -> TierList {
        let roles = choice.faction == .horde ? file.factions.horde : file.factions.alliance
        switch choice.role {
        case .dps: return roles.dps
        case .tank: return roles.tank
        case .healer: return roles.healer
        }
    }

    static func rows(_ list: TierList, yourSpec: String?) -> [TierRowModel] {
        list.rows.map { TierRowModel(row: $0, count: list.rows.count, isYours: $0.spec == yourSpec) }
    }

    static func callout(_ list: TierList, yourSpec: String?) -> TierCallout? {
        guard let yourSpec, let row = list.rows.first(where: { $0.spec == yourSpec }) else { return nil }
        return TierCallout(row: row, note: list.notes.last)
    }

    /// The character's guide key, `mage-fire`, from the bundled spec catalog.
    static func specKey(_ character: AccountCharacter?, specs: SpecCatalog) -> String? {
        guard let classSlug = character?.className, let name = character?.spec else { return nil }
        return specs.spec(classSlug: classSlug, named: name)?.spec
    }
}
