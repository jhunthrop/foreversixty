@testable import ForeverSixty
import Auth
import DesignSystem
import Foundation
import Models
import Networking
import Routing
import SwiftUI
import Testing
import UIKit

/// The flows mocks' Gear scenario (design/mocks/ws_app_flows.py `gear`, `DOLL`, `item_sheet`):
/// Thandrel's Fury Warrior 60 raid Alliance list (data/builds/1.60.1.70790/bis/warrior-fury.json,
/// every name checked against it), four slots worn. Don Julio's Band needs exalted in the loot
/// file (the mock said revered).
enum GearTestData {
    static let build = "1.60.1.70790"

    private static func slot(_ slot: String, _ label: String, _ name: String, _ icon: String, _ quality: Int, _ gain: String?) -> GearSlotRow {
        let look = HomeItemLook(id: abs(name.hashValue % 100_000), name: name, quality: quality, icon: icon, itemLevel: nil)
        let trail: GearSlotRow.Trail = gain.map { .gain(HomeGain(figure: $0, unit: "DPS")) } ?? .worn
        return GearSlotRow(slot: slot, label: label, pick: GearSlotRow.Pick(
            look: look, trail: trail, item: GearItemRef(itemID: look.id, slot: slot, gain: nil)
        ))
    }

    /// Greenhammer unlocks with the first raids, so the off hand is the Persuader he wears (round 5).
    static let slots = [
        slot("head", "Head", "Lionheart Helm", "inv_helmet_36", 4, "+14.5"),
        slot("neck", "Neck", "Skibi’s Pendant", "inv_jewelry_amulet_07", 3, nil),
        slot("shoulder", "Shoulder", "Theramore Pauldrons", "inv_shoulder_06", 4, nil),
        slot("back", "Back", "Howler’s Furs", "inv_misc_cape_10", 3, "+2.4"),
        slot("chest", "Chest", "Timbermaw Tunic", "inv_chest_fur", 4, "+5.3"),
        slot("wrist", "Wrist", "Forest Stalker’s Bracers", "inv_bracer_12", 4, "+0.1"),
        slot("hands", "Hands", "Theramore Gauntlets", "inv_gauntlets_29", 4, "+1.0"),
        slot("waist", "Waist", "Valiant’s Waistguard", "inv_belt_33", 4, nil),
        slot("legs", "Legs", "Sentinel’s Plate Legguards", "inv_pants_plate_20", 4, "+4.7"),
        slot("feet", "Feet", "Scalegut Treaders", "inv_boots_chain_11", 3, "+0.9"),
        slot("finger1", "Finger 1", "Don Julio’s Band", "inv_jewelry_ring_35", 4, "+6.8"),
        slot("finger2", "Finger 2", "Cutthroat’s Signet", "inv_jewelry_ring_01", 3, "+0.8"),
        slot("trinket1", "Trinket 1", "Hand of Justice", "inv_jewelry_talisman_01", 3, nil),
        slot("trinket2", "Trinket 2", "Molten Heart of the Mountain", "inv_misc_gem_ruby_01", 3, "+3.2"),
        slot("main_hand", "Main hand", "Ironwood Blade", "inv_sword_07", 4, "+5.9"),
        slot("off_hand", "Off hand", "Persuader", "inv_hammer_08", 4, nil),
        slot("ranged", "Ranged", "Blackcrow", "inv_weapon_crossbow_04", 3, "+1.9"),
    ]

    private static func place(_ title: String, _ sub: String, _ icon: String, _ quality: Int, _ figure: String) -> GearPlaceRow {
        GearPlaceRow(
            id: title, title: title, sub: sub, icon: icon, quality: quality, figure: figure, item: GearItemRef(itemID: 1, slot: "head", gain: nil)
        )
    }

    static let content = GearContent(
        title: "Fury Warrior", figure: "1,004", unit: "DPS", facts: "60 raid · Alliance · 5 of 17 slots match the list",
        summary: GearSummary(figures: "+47.5 DPS from 12 upgrades", synced: "synced 2 h ago"),
        tonight: [
            place("Blackrock Spire", "Blackcrow · Shadow Hunter Vosh’gajin, 8%", "inv_weapon_crossbow_04", 3, "1 upgrade · +1.9 DPS"),
            place("Blacksmithing", "Lionheart Helm, Ironwood Blade · ask a blacksmith", "inv_helmet_36", 4, "2 upgrades · +20.4 DPS"),
            place("Stormpike Guard", "Don Julio’s Band, needs revered", "inv_jewelry_ring_35", 4, "1 upgrade · +6.8 DPS"),
        ],
        slots: slots
    )

    /// Deathbringer from Onyxia (data/builds 1.60.1.70790, 8%) in the client's order, the round-5 sheet's list.
    static let tooltip = ItemTooltipModel(
        look: HomeItemLook(id: 17068, name: "Deathbringer", quality: 4, icon: "inv_axe_09", itemLevel: nil),
        headerMeta: "Epic One-Hand", known: true, binds: "Binds when picked up", unique: false, slotName: "One-Hand", typeName: "Axe",
        armor: nil, weaponDamage: "114 - 213 Damage", weaponSpeed: "Speed 2.90", weaponDPS: "(56.4 damage per second)", stats: [],
        equipLines: [], effect: "Chance on hit: Sends a shadowy bolt at the enemy causing 125 Shadow damage.", setName: nil,
        setBonuses: [], requiresLevel: "Requires Level 60", itemLevel: nil, source: nil, alternatives: [], pickFooter: nil
    )

    static let sheet = ItemSheetContent(
        tooltip: tooltip,
        gain: HomeGain(figure: "+12.6", unit: "DPS"),
        source: ItemSheetRow(id: "source", title: "Onyxia, Onyxia’s Lair, 8%", sub: nil, lead: .none),
        replaces: ItemSheetRow(id: "replaces", title: "Replaces Ironwood Blade", sub: "Equipped", lead: .item(icon: "inv_sword_07", quality: 4)),
        runnerUp: ItemSheetRow(
            id: "runner-up", title: "Vis’kag the Bloodletter", titleQuality: 4, sub: "Runner-up · \u{2212}1.8 DPS",
            lead: .item(icon: "inv_sword_18", quality: 4)
        ),
        guild: .loaded(rows: [
            ItemSheetRow(id: "g", title: "Grunvald", titleClass: "Warrior", sub: Copy.gearSheetNeedsPick(spec: "Protection Warrior"), lead: .crest(className: "warrior")),
            ItemSheetRow(id: "t", title: "Torvak", titleClass: "Shaman", sub: Copy.gearSheetHas(spec: "Enhancement Shaman"), lead: .crest(className: "shaman")),
        ], note: nil),
        raidZone: "Onyxia’s Lair",
        nextRaid: ItemSheetBuilder.nextRaid(nextOnyxia, guild: 2, zone: "Onyxia’s Lair", locale: Locale(identifier: "en_US"), timeZone: chicago)
    )

    /// Thandrel's Battle.net portrait in the paperdoll's medallion, as the mock draws it.
    static var portrait: UIImage? {
        UIImage(contentsOfFile: HomeFixtures.repoRoot.appending(path: "design/mocks/data/avatars/thandrel.jpg").path)
    }

    static let chicago = TimeZone(identifier: "America/Chicago")!
    /// Wednesday 11 November 2026, 8:00 pm in Chicago.
    static let nextOnyxia = ItemNextRaid(raidId: 31, zoneKey: "raid:onyxias-lair", startsAt: Date(timeIntervalSince1970: 1_794_448_800))
}

@Suite("Item sheet next raid")
struct ItemSheetNextRaidTests {
    @Test("the day and time follow the device: 12-hour in the US, 24-hour in the UK; the line opens that raid")
    func formats() throws {
        let us = try #require(ItemSheetBuilder.nextRaid(
            GearTestData.nextOnyxia, guild: 2, zone: "Onyxia’s Lair", locale: Locale(identifier: "en_US"), timeZone: GearTestData.chicago
        ))
        #expect(us.line == "Next raid: Onyxia’s Lair, Wed 8:00 PM")
        #expect(us.link == .raid(guild: 2, raid: 31))
        let uk = ItemSheetBuilder.nextRaid(
            GearTestData.nextOnyxia, guild: 2, zone: "Onyxia’s Lair", locale: Locale(identifier: "en_GB"), timeZone: GearTestData.chicago
        )
        #expect(uk?.line == "Next raid: Onyxia’s Lair, Wed 20:00")
    }

    @Test("no line without a planned raid or a raid source")
    func absent() {
        #expect(ItemSheetBuilder.nextRaid(nil, guild: 2, zone: "Onyxia’s Lair") == nil)
        #expect(ItemSheetBuilder.nextRaid(GearTestData.nextOnyxia, guild: 2, zone: nil) == nil)
    }
}

@MainActor
@Suite("Gear parity")
struct GearParityTests {
    static let device = CGSize(width: 393, height: 852)

    @Test("ios-gear")
    func gear() async throws {
        let view = GearListView(content: GearTestData.content, build: GearTestData.build, portrait: GearTestData.portrait)
            .environment(\.itemIconSource, RepoItemIconSource())
        try await ParityCheck.assert(view, screen: "gear", size: Self.device)
    }

    @Test("ios-gear-out")
    func signedOut() async throws {
        let view = GearClassesView(classes: try BisTestData.classes().map(\.slug), name: { $0.capitalized })
        try await ParityCheck.assert(view, screen: "gear-out", size: Self.device)
    }

    @Test("ios-item-sheet")
    func itemSheet() async throws {
        // The mock's sheet sits 120 under the status bar, over the dimmed screen.
        let view = VStack(spacing: 0) {
            Color.black.opacity(0.45).frame(height: 120)
            ItemSheetView(content: GearTestData.sheet, build: GearTestData.build)
                .clipShape(.rect(topLeadingRadius: 16, topTrailingRadius: 16))
        }
        .background(FSColor.bg)
        .parityContent()
        .environment(\.itemIconSource, RepoItemIconSource())
        try await ParityCheck.assert(view, screen: "item-sheet", size: Self.device)
    }
}

/// Where to go, Gear's words and the item sheet's rows.
@Suite("Gear")
struct GearLogicTests {
    // Head: a Blacksmithing pick with a Blackrock Spire runner-up 0.4 behind; finger: a Stormpike
    // Guard pick (exalted); main hand: a Blacksmithing pick; chest: a Leatherworking pick; wrist:
    // a Maraudon pick above a level 50's reach.
    static let slots = """
    [{"slot":"head","verified":true,"item_id":1,"item_name":"Lionheart Helm","source":"Blacksmithing","source_kind":"crafted",
      "alternatives":[{"item_id":2,"item_name":"Blackcrow","source":"Blackrock Spire","source_kind":"dungeon","dps_delta":-0.4}]},
     {"slot":"finger1","verified":true,"item_id":3,"item_name":"Don Julio's Band","source":"Stormpike Guard","source_kind":"rep"},
     {"slot":"main_hand","verified":true,"item_id":4,"item_name":"Ironwood Blade","source":"Blacksmithing","source_kind":"crafted"},
     {"slot":"chest","verified":true,"item_id":5,"item_name":"Timbermaw Tunic","source":"Leatherworking","source_kind":"crafted"},
     {"slot":"wrist","verified":true,"item_id":6,"item_name":"Far Bracers","source":"Maraudon","source_kind":"dungeon"}]
    """

    static let loot = """
    {"sources":[
      {"id":"dungeon:brs","kind":"dungeon","name":"Blackrock Spire","level_min":55,"bosses":[
        {"id":"b1","name":"Shadow Hunter Vosh'gajin","items":[2],"item_chances":{"2":8.2}}]},
      {"id":"dungeon:mara","kind":"dungeon","name":"Maraudon","level_min":56,"items":[6]},
      {"id":"crafted:bs","kind":"crafted","name":"Blacksmithing","profession":"blacksmithing","items":[1,4],"required_skill":{"4":245}},
      {"id":"crafted:lw","kind":"crafted","name":"Leatherworking","profession":"leatherworking","items":[5]},
      {"id":"rep:sp","kind":"rep","name":"Stormpike Guard","standing":"exalted","items":[3]}]}
    """

    private func whereToGo(gains: [String: Double], level: Int?) throws -> WhereToGo {
        let slots = try WireFormat.decoder().decode([BisSlot].self, from: Data(Self.slots.utf8))
        let loot = try WireFormat.decoder().decode(LootFile.self, from: Data(Self.loot.utf8))
        return WhereToGo(slots: slots, gains: gains, places: ItemPlaces(loot: loot, level: level))
    }

    @Test("round 7, row 3: the Gear sheet's Where line from the loaded loot file: boss, instance, chance; profession; reputation")
    func sheetWhere() throws {
        let loot = try WireFormat.decoder().decode(LootFile.self, from: Data(Self.loot.utf8))
        let places = ItemPlaces(loot: loot, level: nil)
        func row(_ item: Int, _ kind: String) -> ItemSheetRow? {
            ItemSheetBuilder.source(item, kind: kind, named: nil, places: places, fallback: nil)
        }
        // Round 8, row 1: the shared source words with the figure, no caption, no image.
        #expect(row(2, "dungeon")?.title == "Shadow Hunter Vosh'gajin, Blackrock Spire, 8%")
        #expect(row(1, "crafted")?.title == "Crafted: Blacksmithing")
        #expect(row(4, "crafted")?.title == "Crafted: Blacksmithing (245)", "the recipe's skill, WoW's Requires shape")
        #expect(row(3, "rep")?.title == "Stormpike Guard, exalted")
        #expect(row(2, "dungeon")?.sub == nil && row(2, "dungeon")?.lead == ItemSheetRow.Lead.none)
        let world = ItemSheetBuilder.source(99, kind: "world", named: nil, places: places, fallback: "World drop")
        #expect(world?.title == "World drop" && world?.lead == ItemSheetRow.Lead.none)
    }

    /// A runner-up row for a head pick (Lionheart Helm, crafted) whose first alternative is `alt`.
    private func runnerUpRow(_ alt: String) throws -> ItemSheetRow? {
        let slot = try WireFormat.decoder().decode(BisSlot.self, from: Data("""
        {"slot":"head","verified":true,"item_id":1,"item_name":"Lionheart Helm","source":"Blacksmithing","source_kind":"crafted",
         "alternatives":[\(alt)]}
        """.utf8))
        let places = ItemPlaces(loot: try WireFormat.decoder().decode(LootFile.self, from: Data(Self.loot.utf8)), level: nil)
        let listing = ItemSheetBuilder.Listing(slot: slot, isPick: true, kind: "crafted", source: "Blacksmithing")
        return ItemSheetBuilder.runnerUp(listing, band: nil, items: [:], places: places, pickPlace: places.place(of: 1, kind: "crafted"))
    }

    @Test("round 9: the runner-up names its source only when it differs from the pick's, in the source row's words")
    func runnerUpWords() throws {
        // The same crafter as the pick: no source, the delta alone.
        let crafted = #"{"item_id":4,"item_name":"Ironwood Blade","source":"Blacksmithing","source_kind":"crafted","dps_delta":-5.0}"#
        #expect(try runnerUpRow(crafted)?.sub == "Runner-up · \u{2212}5.0 DPS")
        // A boss drop: `Boss · 8%`.
        let boss = #"{"item_id":2,"item_name":"Blackcrow","source":"Blackrock Spire","source_kind":"dungeon","dps_delta":-0.4}"#
        #expect(try runnerUpRow(boss)?.sub == "Runner-up · Shadow Hunter Vosh'gajin · 8% · \u{2212}0.4 DPS")
        // PvP without the faction word, and `Same DPS` when nothing separates them.
        let pvp = #"{"item_id":9,"item_name":"Blood Guard's Plate Helm","source":"PvP rank 11 · Blood Guard · Horde","source_kind":"vendor","dps_delta":0}"#
        #expect(try runnerUpRow(pvp)?.sub == "Runner-up · PvP rank 11 · Blood Guard · Same DPS")
        // A crafted runner-up from another crafter, in the shared words.
        let other = #"{"item_id":5,"item_name":"Timbermaw Tunic","source":"Leatherworking","source_kind":"crafted","dps_delta":-1.2}"#
        #expect(try runnerUpRow(other)?.sub == "Runner-up · Crafted: Leatherworking · \u{2212}1.2 DPS")
        // Not in the loot file: the list's words decide, and the same crafter is still left out.
        let unplaced = #"{"item_id":77,"item_name":"Thorium Cestus","source":"Blacksmithing","source_kind":"crafted","dps_delta":-5.0}"#
        #expect(try runnerUpRow(unplaced)?.sub == "Runner-up · \u{2212}5.0 DPS")
    }

    @Test("round 9: a PvP source drops its faction word; anything else is left alone")
    func withoutFaction() {
        #expect(ItemSheetBuilder.withoutFaction("PvP rank 11 · Blood Guard · Horde") == "PvP rank 11 · Blood Guard")
        #expect(ItemSheetBuilder.withoutFaction("PvP rank 12 · Knight · Alliance") == "PvP rank 12 · Knight")
        #expect(ItemSheetBuilder.withoutFaction("World drop") == "World drop")
    }

    @Test("round 8: drop chances in the site's dropPercentLabel: whole percent, half up, <1% under one")
    func dropPercent() {
        #expect(ItemNumbers.percent(1.3) == "1%")
        #expect(ItemNumbers.percent(19.6) == "20%")
        #expect(ItemNumbers.percent(2.5) == "3%")
        #expect(ItemNumbers.percent(8.2) == "8%")
        #expect(ItemNumbers.percent(0.4) == "<1%")
        #expect(ItemNumbers.percent(1) == "1%")
    }

    @Test("every upgrade grouped by place; tonight is the best place of each kind, instance first")
    func groups() throws {
        let gains = ["head": 14.5, "finger1": 6.8, "main_hand": 5.9, "chest": 5.3, "wrist": 2.0]
        let places = try whereToGo(gains: gains, level: 52)
        #expect(places.groups.map(\.place.name) == ["Blackrock Spire", "Blacksmithing", "Leatherworking", "Stormpike Guard"])
        #expect(places.tonight.map(\.place.name) == ["Blackrock Spire", "Blacksmithing", "Stormpike Guard"])
        let smith = try #require(places.groups.first { $0.place.name == "Blacksmithing" })
        #expect(smith.entries.map(\.name) == ["Lionheart Helm", "Ironwood Blade"])
        let spire = try #require(places.groups.first)
        #expect(abs((spire.entries.first?.gain ?? 0) - 14.1) < 1e-9)
        #expect(spire.entries.first?.chance == 8.2)
        #expect(places.groups.last?.standing == "exalted")
    }

    @Test("a slot that does not gain is not an upgrade")
    func noGain() throws {
        #expect(try whereToGo(gains: [:], level: 60).groups.isEmpty)
        #expect(try whereToGo(gains: ["head": -1], level: 60).groups.isEmpty)
    }

    @Test("the place lines in the mock's words")
    func placeWords() throws {
        let places = try whereToGo(gains: ["head": 14.5, "finger1": 6.8, "main_hand": 5.9], level: 60)
        let words = places.tonight.map { GearPresenter.placeSub($0, names: $0.entries.map(\.name).joined(separator: ", ")) }
        #expect(words == [
            "Blackcrow · Shadow Hunter Vosh'gajin, 8%",
            "Lionheart Helm, Ironwood Blade · ask a blacksmith",
            "Don Julio's Band, needs exalted",
        ])
    }

    @Test("each place row adds up its upgrades and their gains")
    func placeFigure() throws {
        let places = try whereToGo(gains: ["head": 14.5, "finger1": 6.8, "main_hand": 5.9], level: 60)
        let smith = try #require(places.tonight.first { $0.place.name == "Blacksmithing" })
        #expect(GearPresenter.placeFigure(smith, unit: "DPS") == "2 upgrades · +20.4 DPS")
        let rep = try #require(places.tonight.first { $0.place.name == "Stormpike Guard" })
        #expect(GearPresenter.placeFigure(rep, unit: "DPS") == "1 upgrade · +6.8 DPS")
    }

    @Test("the header line: the gain from the upgrades, the gains not modelled, the synced age")
    func headerLine() throws {
        let home = try HomeFixtures.home()
        let synced = try #require(home.syncedAt)
        let line = GearPresenter.summary(home, unit: "DPS", now: synced.addingTimeInterval(2 * 3600))
        #expect(line?.line == "+102.5 DPS from 12 upgrades · 2 gains not modelled yet · synced 2 h ago")
        #expect(line?.stale == false)
    }

    @Test("no list for the spec and band: the list line with Change and one row, never a blank page")
    func notice() {
        let selection = BisSelection(classSlug: "hunter", spec: "hunter-beast-mastery", range: BisRange(band: 30, preset: "bare"), faction: "horde")
        let presenter = GearPresenter(context: BisContext(selection: selection, characterSelection: selection))
        let title = "Beast Mastery Hunter"
        #expect(presenter.notice(phase: .noList(offlineSince: nil), title: title) == GearNotice(
            title: title, facts: "30 to 39 · Horde", row: "No best in slot list published yet for this spec and level range."
        ))
        let failed = presenter.notice(phase: .failed, title: title)
        #expect(failed?.row == "We could not load this list")
        #expect(failed?.sub == Copy.homeErrorLine)
        #expect(failed?.retry == true)
        #expect(presenter.notice(phase: .notCached, title: title)?.retry == true)
        #expect(presenter.notice(phase: .loading, title: title) == nil)
    }

    @Test("row 10: the counter, the Equipped cells, the header's count and the gain cells agree with the answer")
    func countsAgree() throws {
        let home = try HomeFixtures.home("character.home.gorvash")
        let range = BisRange(band: 50, preset: "bare")
        let band = try #require(try BisTestData.bis().band(for: range, faction: "horde"))
        // A list whose picks are the answer's own: its upgrades' picks, then what is already best.
        let pick = { (id: Int, name: String) in
            BisRow.Content.pick(BisPick(
                item: BisItemRef(id: id, name: name, item: nil), slotSource: "", next: "", isUnmodelled: false, isWorn: false,
                runnersUp: [], source: ""
            ))
        }
        let rows = home.upgrades.map { BisRow(slot: $0.slot, slotLabel: $0.slotLabel, content: pick($0.pick.id, $0.pick.name ?? "")) }
            + home.alreadyBest.map { BisRow(slot: $0.slot, slotLabel: $0.slotLabel, content: pick($0.item.id, $0.item.name ?? "")) }
        let list = BisList(setDpsLine: nil, rows: rows, unit: .dps, specLabel: "Fury Warrior", rangeLabel: "50 to 59")
        let selection = BisSelection(classSlug: "warrior", spec: "warrior-fury", range: range, faction: "horde")
        let content = GearPresenter(context: BisContext(selection: selection, characterSelection: selection))
            .content(list: list, inputs: GearInputs(band: band, home: home, loot: nil, level: 58))
        let worn = content.slots.filter { $0.pick?.trail == .worn }.count
        let gains = content.slots.filter { if case .gain = $0.pick?.trail { true } else { false } }.count
        #expect(worn == home.alreadyBest.count)
        #expect(content.facts.hasSuffix(Copy.gearOnList(n: worn, total: rows.count)))
        #expect(gains == home.upgradeCount)
        #expect(content.summary?.line.contains(Copy.homeAnswerUpgrades(n: home.upgradeCount)) == true)
        #expect(content.slots.map(\.label).contains("Finger 2"), "the API's slot label")
    }

    @Test("row 10, the fixture app: Gorvash's 50 to 59 list reads 3 of 17, Head, Neck and Main hand Equipped")
    func gorvashFixture() throws {
        let home = try HomeFixtures.home("character.home.gorvash")
        let range = BisRange(band: 50, preset: "bare")
        let file = try BisTestData.bis()
        let band = try #require(file.band(for: range, faction: "horde"))
        let list = BisListBuilder(items: [], worn: [], timeZone: .gmt).list(file: file, band: band, specLabel: "Fury Warrior", rangeLabel: "50 to 59")
        let selection = BisSelection(classSlug: "warrior", spec: "warrior-fury", range: range, faction: "horde")
        let content = GearPresenter(context: BisContext(selection: selection, characterSelection: selection))
            .content(list: list, inputs: GearInputs(band: band, home: home, loot: nil, level: 58))
        let worn = content.slots.filter { $0.pick?.trail == .worn }.map(\.slot)
        #expect(worn == ["head", "neck", "main_hand"])
        #expect(content.facts.hasSuffix(Copy.gearOnList(n: 3, total: 17)))
        // Round 7, row 1: the header in the house format, whole from 100.
        #expect(content.figure == band.setDps.map(ReportsCopy.houseFigure))
        #expect(content.figure?.contains(".") == false)
    }

    @Test("round 4, row 3: a pick neither Equipped nor gaining reads Not modelled or No gain")
    func flatTrails() throws {
        let home = try HomeFixtures.home()
        let upgrade = try #require(home.upgrades.first)
        let gain = HomeGain(figure: "+4.2", unit: "DPS")
        #expect(GearPresenter.trail(upgrade, gain: gain, unit: "DPS") == .gain(gain))
        let unmodelled = try #require(home.upgrades.first { !$0.modelled }, "the fixture has gains not modelled")
        #expect(GearPresenter.trail(unmodelled, gain: nil, unit: "DPS") == .notModelled)
        #expect(GearPresenter.trail(nil, gain: nil, unit: "DPS") == .noGain, "round 8: `No gain`, not +0.0 DPS")
    }

    @Test("round 4, row 2: Gear's synced age takes the warning tone past a week, as the drawer's")
    func staleSummary() throws {
        let home = try HomeFixtures.home()
        let synced = try #require(home.syncedAt)
        #expect(GearPresenter.summary(home, unit: "DPS", now: synced.addingTimeInterval(SyncAge.staleAfter - 60))?.stale == false)
        let stale = try #require(GearPresenter.summary(home, unit: "DPS", now: synced.addingTimeInterval(SyncAge.staleAfter + 60)))
        #expect(stale.stale)
        #expect(stale.synced == Copy.gearHeaderSynced(age: RelativeAge.days(7).phrase))
    }

    @Test("the counter reads how many slots match the list")
    func counterWords() {
        #expect(Copy.gearOnList(n: 0, total: 17) == "0 of 17 slots match the list")
    }

    @Test("the summary names the band: 60 raid, or the band's range")
    func band() {
        #expect(GearPresenter.band(BisRange(band: 60, preset: "raid")) == "60 raid")
        #expect(GearPresenter.band(BisRange(band: 30, preset: "bare")) == "30 to 39")
    }

    @Test("who in the guild: needs (pick, then runner-up), then has, crests by class")
    func guild() throws {
        struct Wrapped: Decodable { let data: ItemNeeds }
        let answer = try BisTestData.fixture("guild.item-needs", as: Wrapped.self).data
        guard case let .loaded(rows, note) = ItemSheetBuilder.guildRows(answer) else { Issue.record("not loaded"); return }
        #expect(rows.map(\.title) == ["Kraggor", "Brightwen", "Ironhide"])
        #expect(rows.map(\.sub) == [
            "Fury Warrior · needs it, the pick on their list", "Combat Rogue · needs it, a runner-up on their list", "Protection Warrior · has it",
        ])
        #expect(rows.first?.lead == .crest(className: "warrior"))
        #expect(note == "3 guildmates have not shared their gear yet")
        let empty = ItemNeeds(itemId: 1, needs: [], has: [], unknownGear: 0)
        #expect(ItemSheetBuilder.guildRows(empty) == .loaded(rows: [], note: "Nobody in the guild has it on their list."))
    }

    @Test("the sheet's tooltip leaves the item level out")
    func itemLevel() throws {
        let request = ItemTooltipRequest.forHome(
            HomeItemLook(id: 252_605, name: "Black Dragonscale Helm", quality: 4, icon: nil, itemLevel: 61),
            home: try HomeFixtures.home(), character: HomeFixtures.velaren
        )
        let model = ItemTooltipModel.build(request, sources: try HomeFixtures.tooltipSources())
        #expect(model.itemLevel != nil)
        #expect(model.withoutItemLevel.itemLevel == nil)
        #expect(model.withoutItemLevel.binds == model.binds)
    }

    @Test("what it replaces: the worn item in the slot, never the item itself")
    func replaces() {
        let worn = Item(
            id: 9, name: "Ironwood Blade", icon: "inv_sword_07", slot: "one_hand", quality: 4, requiredLevel: 60, requiredLevelSource: "",
            itemLevel: 63, armor: 0, damageMin: 0, damageMax: 0, speed: 0, dps: 0, twoHand: false, effectText: "", unique: false,
            clientUnconfirmed: false, iconSource: ""
        )
        let row = ItemSheetBuilder.replaces(9, item: 17068, items: [9: worn])
        #expect(row?.title == "Replaces Ironwood Blade")
        #expect(row?.sub == Copy.gearWorn)
        #expect(ItemSheetBuilder.replaces(17068, item: 17068, items: [:]) == nil)
        #expect(ItemSheetBuilder.replaces(nil, item: 17068, items: [9: worn]) == nil)
    }
}
