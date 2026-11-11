import Auth
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
    /// The synced stamp's clock.
    var now = AppClock.now()

    func content(list: BisList, inputs: GearInputs) -> GearContent {
        let home = ownHome(inputs.home)
        let gains = home.map { BestRun.gains($0.upgrades) } ?? [:]
        let places = inputs.loot.map { ItemPlaces(loot: $0, level: inputs.level) }
        let unit = home?.unitWord ?? list.unit.word
        let refs = Self.itemRefs(list)
        let tonight = places.map { WhereToGo(slots: inputs.band.slots, gains: gains, places: $0).tonight } ?? []
        let slots = list.rows.map { slotRow($0, home: home, unit: unit) }
        return GearContent(
            title: list.specLabel,
            // The header in the house format (round 7, row 1): `182`; a tank band keeps none.
            figure: list.setFigure.flatMap { _ in inputs.band.setDps.map(ReportsCopy.houseFigure) },
            unit: unit,
            facts: facts(slots: slots),
            summary: home.flatMap { Self.summary($0, unit: unit, now: now) },
            tonight: tonight.compactMap { placeRow($0, refs: refs, unit: unit) },
            slots: slots
        )
    }

    /// The answer's gains hold for the list it was computed on: the character's own.
    private func ownHome(_ home: CharacterHome?) -> CharacterHome? {
        guard context.isYours, let home, home.spec == nil || home.spec == context.selection.spec else { return nil }
        return home
    }

    // MARK: The line naming the list

    /// `60 raid · Alliance · 4 of 17 slots on the list`: the band, the faction, how many picks are worn.
    private func facts(slots: [GearSlotRow]) -> String {
        let picks = slots.compactMap(\.pick)
        let worn = picks.filter { $0.trail == .worn }.count
        return facts(more: [Copy.gearOnList(n: worn, total: picks.count)])
    }

    /// The band and the faction, then `more`.
    private func facts(more: [String] = []) -> String {
        let selection = context.selection
        return ([Self.band(selection.range), BisLabels.faction(selection.faction)] + more).joined(separator: Self.separator)
    }

    /// Gear's notice for a phase with no list to draw, under the spec's `title`; nil while the list
    /// loads or once it has one.
    func notice(phase: BisPhase, title: String) -> GearNotice? {
        switch phase {
        case .noList: GearNotice(title: title, facts: facts(), row: Copy.homeAnswerNoList)
        case .failed: GearNotice(title: title, facts: facts(), row: Copy.gearErrorTitle, sub: Copy.homeErrorLine, retry: true)
        case .notCached: GearNotice(title: title, facts: facts(), row: Copy.gearErrorTitle, sub: Copy.bisNotCached, retry: true)
        case .loading, .loaded: nil
        }
    }

    static let separator = " · "

    /// `+157.6 DPS from 13 upgrades · 2 gains not modelled yet · synced 2 h ago`: what the answer adds
    /// up to, what it could not measure, and how old the gear it read is.
    static func summary(_ home: CharacterHome, unit: String, now: Date) -> GearSummary? {
        var line = ""
        if home.upgradeCount > 0, let total = home.totalGain, total > 0 {
            line = Copy.gearHeaderGain(
                gain: ItemNumbers.signedTenths(total), unit: unit, upgrades: Copy.homeAnswerUpgrades(n: home.upgradeCount)
            )
        }
        if home.notModelledCount > 0 {
            // The key carries its own ` · ` lead; alone it starts the line.
            let words = Copy.homeAnswerNotModelled(n: home.notModelledCount)
            line += line.isEmpty ? String(words.drop { $0 == " " || $0 == "·" }) : words
        }
        let summary = GearSummary(
            figures: line.isEmpty ? nil : line,
            synced: home.syncedAt.map { Copy.gearHeaderSynced(age: RelativeAge(from: $0, to: now).phrase) },
            stale: home.syncedAt.map { SyncAge.isStale($0, now: now) } ?? false
        )
        return summary.line.isEmpty ? nil : summary
    }

    /// `60 raid` for the raid list, the band's words otherwise.
    static func band(_ range: BisRange) -> String {
        range.preset == BisRange.raidPreset ? Copy.gearBandRaid(band: BisLabels.band(range)) : BisLabels.band(range)
    }

    // MARK: Every slot

    private func slotRow(_ row: BisRow, home: CharacterHome?, unit: String) -> GearSlotRow {
        guard let pick = row.pick else { return GearSlotRow(slot: row.slot, label: row.slotLabel, pick: nil) }
        // Row 10, as Android: the answer's `already_best` decides by slot (the server read the gear against
        // this list); without an answer, the list's own worn mark.
        let worn = home.map { $0.alreadyBest.contains { $0.slot == row.slot } } ?? pick.isWorn
        let upgrade = home?.upgrades.first { $0.slot == row.slot && $0.pick.id == pick.item.id }
        let gain = worn ? nil : Self.gain(upgrade, unit: unit)
        let look = HomeItemLook(
            id: pick.item.id, name: pick.item.name, quality: pick.item.item?.quality, icon: pick.item.icon, itemLevel: pick.item.item?.itemLevel
        )
        return GearSlotRow(slot: row.slot, label: row.slotLabel, pick: GearSlotRow.Pick(
            look: look,
            trail: worn ? .worn : home.map { _ in Self.trail(upgrade, gain: gain, unit: unit) },
            item: GearItemRef(itemID: pick.item.id, slot: row.slot, gain: gain)
        ))
    }

    /// Round 4, row 3: with the character's answer every pick that is not Equipped says what it gives:
    /// its gain, `Not modelled` when the sim could not measure it, else `No gain` (round 8, row 3).
    static func trail(_ upgrade: HomeUpgrade?, gain: HomeGain?, unit: String) -> GearSlotRow.Trail {
        if let gain { return .gain(gain) }
        if upgrade?.modelled == false { return .notModelled }
        return .noGain
    }

    /// A measured gain above zero, the server's, as Home words it.
    static func gain(_ upgrade: HomeUpgrade?, unit: String) -> HomeGain? {
        guard let upgrade, upgrade.modelled, let gain = upgrade.gain, gain > 0 else { return nil }
        return HomeGain(figure: ItemNumbers.signedTenths(gain), unit: unit)
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
            figure: Self.placeFigure(group, unit: unit),
            item: GearItemRef(itemID: top.itemID, slot: top.slot, gain: HomeGain(figure: ItemNumbers.signedTenths(top.gain), unit: unit))
        )
    }

    /// `2 upgrades · +20.4 DPS`: the place's upgrades and their gains together.
    static func placeFigure(_ group: WhereToGo.Group, unit: String) -> String? {
        guard !group.entries.isEmpty else { return nil }
        let total = group.entries.reduce(0) { $0 + $1.gain }
        return Copy.gearWhereFigure(
            upgrades: Copy.homeAnswerUpgrades(n: group.entries.count), gain: ItemNumbers.signedTenths(total), unit: unit
        )
    }

    static func placeSub(_ group: WhereToGo.Group, names: String) -> String {
        switch group.place.kind {
        case .instance:
            guard group.entries.count == 1, let top = group.entries.first, let boss = top.place.boss else { return names }
            guard let chance = top.chance else { return Copy.gearWhereBoss(items: names, boss: boss) }
            return Copy.gearWhereDrop(items: names, boss: boss, chance: ItemNumbers.percent(chance))
        case .profession:
            return Copy.gearWhereAsk(items: names, crafter: crafter(group.place.profession))
        case .reputation:
            guard let standing = group.standing else { return names }
            let word = ItemVocabulary.standingName(standing)?.lowercased() ?? standing
            return Copy.gearWhereRep(items: names, standing: word)
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
