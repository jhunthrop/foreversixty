import Foundation
import Models

/// One icon of `The gear it ran`, in the character frame's slot order.
struct SimGearCell: Equatable, Sendable, Identifiable {
    let slot: String
    let itemID: Int
    /// The top answer's swap on a Top Gear or Droptimizer result (gold tick).
    let changed: Bool

    var id: String { slot }
}

/// What F5.3 draws for one result (spec §5.5 F5.3). Reads the result and, when it has one, the
/// list row for its headline; nothing here re-derives a number the engine wrote.
struct SimResultModel: Equatable, Sendable {
    enum Answer: Equatable, Sendable {
        /// A plain run: the mean to one decimal, `DPS` beside it.
        case figure(String)
        /// Every other kind: the row's headline sentence.
        case sentence(String)
    }

    let title: String
    let kind: String
    let specName: String
    let className: String
    let classSlug: String
    let level: Int?
    let answer: Answer
    let errorLine: String
    let gear: [SimGearCell]
    /// `sims.changed` under the grid, when a slot carries the corner.
    let showsChangedLegend: Bool
    let breakdownURL: URL

    init(id: String, result: SimResult, row: SimRow?, specs: SpecCatalog) {
        let specKey = result.request?.spec ?? row?.spec ?? ""
        let spec = specs.spec(specKey)
        let character = result.request?.character
        let slug = character?.className ?? spec?.classSlug ?? ""
        specName = spec?.name ?? specKey
        classSlug = slug
        className = GameClass(rawValue: slug).map(\.name) ?? slug.capitalized
        level = character?.level
        let named = result.title ?? row?.title ?? ""
        title = named.isEmpty ? specName : named
        let kindKey = row?.kind ?? Self.kind(of: result)
        kind = Self.kindLabel(kindKey)
        let figure = Self.figure(result.dps.mean)
        if kindKey != "run" || result.healing != nil || result.tank != nil, let headline = row?.headline, !headline.isEmpty {
            answer = .sentence(headline)
        } else {
            answer = .figure(figure)
        }
        errorLine = Copy.simsError(
            error: Self.figure(result.dps.error), n: result.iterationsRun.formatted(), version: result.engineVersion
        )
        gear = Self.gear(result, kind: kindKey)
        showsChangedLegend = gear.contains(where: \.changed)
        breakdownURL = StaticSite.url.appending(path: "sim/\(id)")
    }

    var meta: String {
        let who = "\(specName) \(className)"
        guard let level else { return who }
        return "\(who) · \(Copy.simsLevel(n: level))"
    }

    static func figure(_ value: Double?) -> String {
        (value ?? 0).formatted(.number.precision(.fractionLength(1)).grouping(.automatic))
    }

    /// A result names its kind by its blocks, as the site's `kindOf`.
    static func kind(of result: SimResult) -> String {
        if result.weights != nil { return "weights" }
        if let mode = result.request?.bulk?.mode { return mode }
        return "run"
    }

    static func kindLabel(_ kind: String) -> String {
        switch kind {
        case "gear": Copy.simsKindGear
        case "talents": Copy.simsKindTalents
        case "drops": Copy.simsKindDrops
        case "weights": Copy.simsKindWeights
        default: Copy.simsKindRun
        }
    }

    /// The gear the run wore, in frame order. On a Top Gear result the slots its best set changes
    /// carry the gold corner (spec F5.3 item 6); every other kind marks nothing.
    static func gear(_ result: SimResult, kind: String) -> [SimGearCell] {
        let base = Dictionary((result.request?.character.gear ?? []).map { ($0.slot, $0.itemId) }, uniquingKeysWith: { first, _ in first })
        let changed = kind == "gear" ? Set((result.combos?.first?.substitutions ?? []).compactMap(\.slot)) : []
        return ItemVocabulary.frameSlots.compactMap { slot in
            base[slot].map { SimGearCell(slot: slot, itemID: $0, changed: changed.contains(slot)) }
        }
    }
}
