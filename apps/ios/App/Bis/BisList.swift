import Foundation
import Models

/// The measure a band's figures are in: a healer band's are healing per second, a tank band's
/// its tank score (the site's `rateUnitForRole`).
enum BisUnit: Equatable, Sendable {
    case dps, hps, points

    init(role: String?) {
        switch role {
        case "healer": self = .hps
        case "tank": self = .points
        default: self = .dps
        }
    }

    var word: String {
        switch self {
        case .dps: Copy.bisUnitDps
        case .hps: Copy.bisUnitHps
        case .points: Copy.bisUnitPoints
        }
    }
}

/// An item as a row draws it: the file's id and name, and the build's item facts when known.
struct BisItemRef: Equatable, Sendable, Identifiable {
    let id: Int
    let name: String
    /// Nil when the class's items file is not on the phone; the row still names the item.
    let item: Item?

    var quality: Int { item?.quality ?? 1 }
    var icon: String? { item?.icon }
}

/// A runner-up and its gap to the pick, as published.
struct BisRunnerUp: Equatable, Sendable, Identifiable {
    let item: BisItemRef
    let source: String
    let dpsDelta: Double
    let isWorn: Bool

    var id: Int { item.id }
}

/// A slot's pick, ready to draw.
struct BisPick: Equatable, Sendable {
    let item: BisItemRef
    /// `Head · Crafted: Leatherworking`.
    let slotSource: String
    /// `Next: Warlord's Mail Helm, −1.3 DPS`, or `Nothing else ranked`.
    let next: String
    let isUnmodelled: Bool
    let isWorn: Bool
    let runnersUp: [BisRunnerUp]
    let source: String
}

/// One of the 17 rows, in the file's (the character frame's) order.
struct BisRow: Equatable, Sendable, Identifiable {
    enum Content: Equatable, Sendable {
        case pick(BisPick)
        /// `Off hand · none: {main hand} is two-handed.` or `{Slot} · no pick for this range.`
        case empty(String)
    }

    let slot: String
    let slotLabel: String
    let content: Content

    var id: String { slot }

    var pick: BisPick? {
        if case let .pick(pick) = content { return pick }
        return nil
    }
}

/// The list for one band, ready to draw. Every figure and order is the published file's.
struct BisList: Equatable, Sendable {
    let setDpsLine: String?
    let rows: [BisRow]
    let unit: BisUnit
    /// `Elemental Shaman`, `60 · Phase 1 raid` and the preset's label, for the tooltip footer.
    let specLabel: String
    let rangeLabel: String
    let presetLabel: String

    /// The rows whose pick or runner-up names contain `query` (spec F3.1 search).
    func filtered(by query: String) -> [BisRow] {
        let needle = query.trimmingCharacters(in: .whitespaces)
        guard !needle.isEmpty else { return rows }
        return rows.filter { row in
            guard let pick = row.pick else { return false }
            return ([pick.item.name] + pick.runnersUp.map(\.item.name))
                .contains { $0.localizedCaseInsensitiveContains(needle) }
        }
    }
}

/// Turns a published band into rows. Nothing is re-ranked or recomputed: the order, the picks,
/// the runners-up and every delta are the file's; this only words them (spec §5.3).
struct BisListBuilder {
    let items: [Int: Item]
    /// Item ids the current character wears, any slot (a ring may sit in either finger).
    let worn: Set<Int>
    let timeZone: TimeZone

    init(items: [Item], worn: Set<Int>, timeZone: TimeZone = .current) {
        self.items = Dictionary(items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
        self.worn = worn
        self.timeZone = timeZone
    }

    /// `29 Oct 2026`, the site's Updated date.
    func updated(_ date: Date) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.timeZone = timeZone
        formatter.dateFormat = "d MMM yyyy"
        return formatter.string(from: date)
    }

    static let twoHandEquipped = "two_hand_equipped"

    func list(file: BisFile, band: BisBand, specLabel: String, rangeLabel: String) -> BisList {
        let unit = BisUnit(role: band.role)
        let mainHand = band.slots.first { $0.slot == "main_hand" }?.itemName
        return BisList(
            setDpsLine: band.setDps.map {
                Copy.bisSetDps(dps: Self.figure($0), unit: unit.word, date: updated(file.generatedAt))
            },
            rows: band.slots.map { row($0, unit: unit, mainHand: mainHand) },
            unit: unit,
            specLabel: specLabel,
            rangeLabel: rangeLabel,
            presetLabel: file.presets[band.preset]?.label ?? band.preset
        )
    }

    func row(_ slot: BisSlot, unit: BisUnit, mainHand: String?) -> BisRow {
        let label = ItemWords.slot(slot.slot)
        guard let id = slot.itemId, id > 0, let name = slot.itemName else {
            let reason = if slot.emptyReason == Self.twoHandEquipped, let mainHand {
                Copy.bisOffHandTwoHand(item: mainHand)
            } else {
                Copy.bisNoPick(slot: label)
            }
            return BisRow(slot: slot.slot, slotLabel: label, content: .empty(reason))
        }
        let source = Self.sourceLine(slot.source ?? "", kind: slot.sourceKind)
        let runnersUp = (slot.alternatives ?? []).map {
            BisRunnerUp(
                item: ref($0.itemId, $0.itemName),
                source: Self.sourceLine($0.source, kind: $0.sourceKind),
                dpsDelta: $0.dpsDelta,
                isWorn: worn.contains($0.itemId)
            )
        }
        let pick = BisPick(
            item: ref(id, name),
            slotSource: source.isEmpty ? label : Copy.bisSlotSource(slot: label, source: source),
            next: Self.nextLine(runnersUp.first, unit: unit),
            isUnmodelled: slot.effectUnmodelled ?? false,
            isWorn: worn.contains(id),
            runnersUp: runnersUp,
            source: source
        )
        return BisRow(slot: slot.slot, slotLabel: label, content: .pick(pick))
    }

    private func ref(_ id: Int, _ name: String) -> BisItemRef {
        BisItemRef(id: id, name: name, item: items[id])
    }

    /// The source in the site's words: a quest, a recipe and a vendor say what they are; every
    /// other kind already reads in the game's words (`Stratholme: Balzaphon`, a PvP rank).
    static func sourceLine(_ source: String, kind: String?) -> String {
        guard !source.isEmpty else { return source }
        switch kind {
        case "quest": return Copy.bisSourceQuest(name: source)
        case "crafted": return Copy.bisSourceCrafted(name: source)
        case "vendor" where !source.hasPrefix(pvpRankPrefix): return Copy.bisSourceVendor(name: source)
        default: return source
        }
    }

    private static let pvpRankPrefix = "PvP rank"

    static func nextLine(_ runnerUp: BisRunnerUp?, unit: BisUnit) -> String {
        guard let runnerUp else { return Copy.tooltipNothingElse }
        guard let delta = signedDelta(runnerUp.dpsDelta) else {
            return Copy.bisNextSame(item: runnerUp.item.name, unit: unit.word)
        }
        return Copy.bisNext(item: runnerUp.item.name, delta: delta, unit: unit.word)
    }

    /// Below 0.05 reads `same` (the site's `alternativeGapLabel`); otherwise a signed figure to
    /// one decimal with a true minus sign.
    static func signedDelta(_ delta: Double) -> String? {
        guard abs(delta) >= sameThreshold else { return nil }
        return (delta < 0 ? "\u{2212}" : "+") + figure(abs(delta))
    }

    static let sameThreshold = 0.05

    static func figure(_ value: Double) -> String {
        value.formatted(
            .number.precision(.fractionLength(1)).rounded(rule: .toNearestOrAwayFromZero).grouping(.never)
                .locale(Locale(identifier: "en_US_POSIX"))
        )
    }
}
