import Foundation
import Models

/// What Gear compares a list with: the band the list was built from, the character's home answer
/// (its gains, only for the character's own list), the loot file and the character's level.
struct GearInputs: Equatable, Sendable {
    let band: BisBand
    let home: CharacterHome?
    let loot: LootFile?
    let level: Int?
}

/// Words the Gear tab from a list and the character's answer (flows spec 4.3). Nothing is ranked
/// here: the picks and their order are the file's, the gains the server's, the places the loot
/// file's (`WhereToGo`).
struct GearPresenter {
    let context: BisContext

    func content(list: BisList, inputs: GearInputs) -> GearContent {
        let home = ownHome(inputs.home)
        let gains = home.map { BestRun.gains($0.upgrades) } ?? [:]
        let places = inputs.loot.map { ItemPlaces(loot: $0, level: inputs.level) }
        let unit = home?.unitWord ?? list.unit.word
        let refs = Self.itemRefs(list)
        let tonight = places.map { WhereToGo(slots: inputs.band.slots, gains: gains, places: $0).tonight } ?? []
        let slots = list.rows.map { slotRow($0, band: inputs.band, home: home, places: places, unit: unit) }
        return GearContent(
            summary: summary(list: list, slots: slots),
            tonight: tonight.compactMap { placeRow($0, refs: refs, unit: unit) },
            slots: slots
        )
    }

    /// The answer's gains hold for the list it was computed on: the character's own.
    private func ownHome(_ home: CharacterHome?) -> CharacterHome? {
        guard context.isYours, let home, home.spec == nil || home.spec == context.selection.spec else { return nil }
        return home
    }

    // MARK: The summary

    private func summary(list: BisList, slots: [GearSlotRow]) -> HomeLine {
        let selection = context.selection
        let picks = slots.compactMap { row -> GearSlotRow.Trail?? in
            if case let .pick(_, _, trail, _) = row.content { return .some(trail) }
            return nil
        }
        let worn = picks.filter { $0 == .worn }.count
        return HomeLine(
            title: Copy.gearList(spec: list.specLabel, band: Self.band(selection.range), faction: BisLabels.faction(selection.faction)),
            sub: Copy.gearOnList(n: worn, total: picks.count)
        )
    }

    /// `60 raid` for the raid list, the band's words otherwise.
    static func band(_ range: BisRange) -> String {
        range.preset == BisRange.raidPreset ? Copy.gearBandRaid(band: BisLabels.band(range)) : BisLabels.band(range)
    }

    // MARK: Every slot

    private func slotRow(_ row: BisRow, band: BisBand, home: CharacterHome?, places: ItemPlaces?, unit: String) -> GearSlotRow {
        guard let pick = row.pick else {
            if case let .empty(reason) = row.content { return GearSlotRow(slot: row.slot, content: .empty(reason)) }
            return GearSlotRow(slot: row.slot, content: .empty(Copy.bisNoPick(slot: row.slotLabel)))
        }
        let worn = pick.isWorn || home?.alreadyBest.contains { $0.slot == row.slot && $0.item.id == pick.item.id } == true
        let gain = worn ? nil : Self.gain(home?.upgrades.first { $0.slot == row.slot && $0.pick.id == pick.item.id }, unit: unit)
        let raw = band.slots.first { $0.slot == row.slot }
        let source = Self.sourceWords(raw?.source ?? "", kind: raw?.sourceKind, place: pick.item.id, places: places)
        let look = HomeItemLook(
            id: pick.item.id, name: pick.item.name, quality: pick.item.item?.quality, icon: pick.item.icon, itemLevel: pick.item.item?.itemLevel
        )
        return GearSlotRow(slot: row.slot, content: .pick(
            look: look,
            source: source.isEmpty ? row.slotLabel : Copy.bisSlotSource(slot: row.slotLabel, source: source),
            trail: worn ? .worn : gain.map { .gain($0) },
            item: GearItemRef(itemID: pick.item.id, slot: row.slot, gain: gain)
        ))
    }

    /// A measured gain above zero, the server's, as Home words it.
    static func gain(_ upgrade: HomeUpgrade?, unit: String) -> HomeGain? {
        guard let upgrade, upgrade.modelled, let gain = upgrade.gain, gain > 0 else { return nil }
        return HomeGain(figure: ItemNumbers.signedTenths(gain), unit: unit)
    }

    /// Gear's source words: a profession by its name, a vendor as `Pix Xizzix, vendor`, an
    /// instance with its chance; every other kind as the list words it.
    static func sourceWords(_ source: String, kind: String?, place item: Int, places: ItemPlaces?) -> String {
        guard !source.isEmpty else { return source }
        switch kind {
        case ItemPlaces.craftedKind: return source
        case "vendor" where !source.hasPrefix(BisListBuilder.pvpRankPrefix): return Copy.gearSourceVendor(name: source)
        case ItemPlaces.dungeonKind, ItemPlaces.raidKind:
            guard let chance = places?.place(of: item, kind: kind)?.chance else { return source }
            return Copy.gearSourceChance(source: source, chance: ItemNumbers.percent(chance))
        default: return BisListBuilder.sourceLine(source, kind: kind)
        }
    }

    // MARK: Where to go tonight

    private func placeRow(_ group: WhereToGo.Group, refs: [Int: BisItemRef], unit: String) -> GearPlaceRow? {
        guard let top = group.entries.first else { return nil }
        let names = group.entries.map(\.name).joined(separator: ", ")
        let ref = refs[top.itemID]
        return GearPlaceRow(
            id: group.id,
            title: group.place.name,
            sub: Self.placeSub(group, names: names),
            icon: ref?.icon,
            quality: ref?.item?.quality,
            item: GearItemRef(itemID: top.itemID, slot: top.slot, gain: HomeGain(figure: ItemNumbers.signedTenths(top.gain), unit: unit))
        )
    }

    static func placeSub(_ group: WhereToGo.Group, names: String) -> String {
        switch group.place.kind {
        case .instance:
            guard group.entries.count == 1, let top = group.entries.first, let boss = top.place.boss else { return names }
            guard let chance = top.chance else { return Copy.gearWhereBossOnly(items: names, boss: boss) }
            return Copy.gearWhereBoss(items: names, boss: boss, chance: ItemNumbers.percent(chance))
        case .profession:
            return Copy.gearWhereCrafter(items: names, crafter: crafter(group.place.profession))
        case .reputation:
            guard let standing = group.standing else { return names }
            let word = ItemVocabulary.standingName(standing)?.lowercased() ?? standing
            return Copy.gearWhereStanding(items: names, standing: word)
        }
    }

    /// Who to ask: `a blacksmith`, `an enchanter`.
    static func crafter(_ profession: String?) -> String {
        switch profession {
        case "alchemy": Copy.gearCrafterAlchemy
        case "blacksmithing": Copy.gearCrafterBlacksmithing
        case "enchanting": Copy.gearCrafterEnchanting
        case "engineering": Copy.gearCrafterEngineering
        case "leatherworking": Copy.gearCrafterLeatherworking
        case "tailoring": Copy.gearCrafterTailoring
        default: Copy.gearCrafterOther
        }
    }

    /// Every item the list draws, picks and runners-up, for the place rows' icons.
    static func itemRefs(_ list: BisList) -> [Int: BisItemRef] {
        let refs = list.rows.compactMap(\.pick).flatMap { [$0.item] + $0.runnersUp.map(\.item) }
        return Dictionary(refs.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
    }
}
