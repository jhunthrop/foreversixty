@testable import ForeverSixty
import DesignSystem
import Foundation
import Models
import Networking
import SwiftUI
import Testing

/// The flows mocks' Gear scenario (design/mocks/ws_app_flows.py `gear`, `DOLL`, `item_sheet`):
/// Thandrel's Fury Warrior 60 raid Alliance list (data/builds/1.60.1.70790/bis/warrior-fury.json,
/// every name checked against it), four slots worn. Don Julio's Band needs exalted in the loot
/// file (the mock said revered).
enum GearTestData {
    static let build = "1.60.1.70790"

    private static func slot(
        _ slot: String, _ label: String, _ name: String, _ icon: String, _ quality: Int, _ source: String, _ gain: String?
    ) -> GearSlotRow {
        let look = HomeItemLook(id: abs(name.hashValue % 100_000), name: name, quality: quality, icon: icon, itemLevel: nil)
        let trail: GearSlotRow.Trail = gain.map { .gain(HomeGain(figure: $0, unit: "DPS")) } ?? .worn
        return GearSlotRow(slot: slot, content: .pick(
            look: look, source: "\(label) · \(source)", trail: trail, item: GearItemRef(itemID: look.id, slot: slot, gain: nil)
        ))
    }

    static let slots = [
        slot("head", "Head", "Lionheart Helm", "inv_helmet_36", 4, "Blacksmithing", "+14.5"),
        slot("neck", "Neck", "Skibi’s Pendant", "inv_jewelry_amulet_07", 3, "World drop", nil),
        slot("shoulder", "Shoulder", "Theramore Pauldrons", "inv_shoulder_06", 4, "Theramore Expeditionary Force", nil),
        slot("back", "Back", "Howler’s Furs", "inv_misc_cape_10", 3, "Pix Xizzix, vendor", "+2.4"),
        slot("chest", "Chest", "Timbermaw Tunic", "inv_chest_fur", 4, "Leatherworking", "+5.3"),
        slot("wrist", "Wrist", "Forest Stalker’s Bracers", "inv_bracer_12", 4, "Silverwing Sentinels", "+0.1"),
        slot("hands", "Hands", "Theramore Gauntlets", "inv_gauntlets_29", 4, "Theramore Expeditionary Force", "+1.0"),
        slot("waist", "Waist", "Valiant’s Waistguard", "inv_belt_33", 4, "Pix Xizzix, vendor", nil),
        slot("legs", "Legs", "Sentinel’s Plate Legguards", "inv_pants_plate_20", 4, "Illiyana Moonblaze, vendor", "+4.7"),
        slot("feet", "Feet", "Scalegut Treaders", "inv_boots_chain_11", 3, "Leatherworking", "+0.9"),
        slot("finger1", "Finger", "Don Julio’s Band", "inv_jewelry_ring_35", 4, "Stormpike Guard", "+6.8"),
        slot("finger2", "Finger", "Cutthroat’s Signet", "inv_jewelry_ring_01", 3, "Pix Xizzix, vendor", "+0.8"),
        slot("trinket1", "Trinket", "Hand of Justice", "inv_jewelry_talisman_01", 3, "Blackrock Depths: Emperor Dagran Thaurissan", nil),
        slot("trinket2", "Trinket", "Molten Heart of the Mountain", "inv_misc_gem_ruby_01", 3, "Enchanting", "+3.2"),
        slot("main_hand", "Main hand", "Ironwood Blade", "inv_sword_07", 4, "Blacksmithing", "+5.9"),
        slot("off_hand", "Off hand", "Greenhammer", "inv_mace_1h_raidhyjalc60_d_01", 4, "Blacksmithing", "+9.7"),
        slot("ranged", "Ranged", "Blackcrow", "inv_weapon_crossbow_04", 3, "Blackrock Spire: Shadow Hunter Vosh’gajin, 8%", "+1.9"),
    ]

    private static func place(_ title: String, _ sub: String, _ icon: String, _ quality: Int) -> GearPlaceRow {
        GearPlaceRow(id: title, title: title, sub: sub, icon: icon, quality: quality, item: GearItemRef(itemID: 1, slot: "head", gain: nil))
    }

    static let content = GearContent(
        summary: HomeLine(title: "Fury Warrior · 60 raid · Alliance", sub: "4 of 17 slots on the list"),
        tonight: [
            place("Blackrock Spire", "Blackcrow · Shadow Hunter Vosh’gajin, 8%", "inv_weapon_crossbow_04", 3),
            place("Blacksmithing", "Lionheart Helm, Greenhammer, Ironwood Blade · ask a blacksmith", "inv_helmet_36", 4),
            place("Stormpike Guard", "Don Julio’s Band, needs exalted", "inv_jewelry_ring_35", 4),
        ],
        slots: slots
    )

    static let sheet = ItemSheetContent(
        look: HomeItemLook(id: 17068, name: "Deathbringer", quality: 4, icon: "inv_axe_09", itemLevel: 75),
        slotLabel: "Main hand",
        gain: HomeGain(figure: "+12.6", unit: "DPS"),
        lines: [
            ItemSheetLine(text: "Binds when picked up"),
            ItemSheetLine(text: "One-Hand · Axe"),
            ItemSheetLine(text: "114 - 213 Damage · Speed 2.90"),
            ItemSheetLine(text: "(56.4 damage per second)"),
            ItemSheetLine(text: "Chance on hit: Sends a shadowy bolt at the enemy causing 125 Shadow damage.", tone: .green),
            ItemSheetLine(text: "Requires Level 60"),
        ],
        source: ItemSheetRow(
            id: "source", title: "Onyxia, Onyxia’s Lair", sub: "8% drop chance · raid",
            lead: .item(icon: "inv_axe_09", quality: 4)
        ),
        runnerUp: ItemSheetRow(
            id: "runner-up", title: "Vis’kag the Bloodletter", titleQuality: 4, sub: "Onyxia · 8% · \u{2212}1.8 DPS",
            lead: .item(icon: "inv_sword_18", quality: 4)
        ),
        guild: .loaded(rows: [
            ItemSheetRow(id: "g", title: "Grunvald", titleClass: "Warrior", sub: "Protection Warrior · needs it", lead: .crest(className: "warrior")),
            ItemSheetRow(id: "t", title: "Torvak", titleClass: "Shaman", sub: "Enhancement Shaman · has it", lead: .crest(className: "shaman")),
        ], note: nil)
    )
}

@MainActor
@Suite("Gear parity")
struct GearParityTests {
    static let device = CGSize(width: 393, height: 852)

    @Test("ios-gear")
    func gear() async throws {
        let view = GearListView(content: GearTestData.content, build: GearTestData.build)
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
                .clipShape(.rect(topLeadingRadius: 12, topTrailingRadius: 12))
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
      {"id":"crafted:bs","kind":"crafted","name":"Blacksmithing","profession":"blacksmithing","items":[1,4]},
      {"id":"crafted:lw","kind":"crafted","name":"Leatherworking","profession":"leatherworking","items":[5]},
      {"id":"rep:sp","kind":"rep","name":"Stormpike Guard","standing":"exalted","items":[3]}]}
    """

    private func whereToGo(gains: [String: Double], level: Int?) throws -> WhereToGo {
        let slots = try WireFormat.decoder().decode([BisSlot].self, from: Data(Self.slots.utf8))
        let loot = try WireFormat.decoder().decode(LootFile.self, from: Data(Self.loot.utf8))
        return WhereToGo(slots: slots, gains: gains, places: ItemPlaces(loot: loot, level: level))
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

    @Test("Gear's source words: a profession bare, a vendor after the name, an instance with its chance")
    func sourceWords() throws {
        let loot = try WireFormat.decoder().decode(LootFile.self, from: Data(Self.loot.utf8))
        let places = ItemPlaces(loot: loot, level: 60)
        #expect(GearPresenter.sourceWords("Blacksmithing", kind: "crafted", place: 1, places: places) == "Blacksmithing")
        #expect(GearPresenter.sourceWords("Pix Xizzix", kind: "vendor", place: 9, places: places) == "Pix Xizzix, vendor")
        #expect(GearPresenter.sourceWords("Blackrock Spire: Shadow Hunter Vosh'gajin", kind: "dungeon", place: 2, places: places)
            == "Blackrock Spire: Shadow Hunter Vosh'gajin, 8%")
        #expect(GearPresenter.sourceWords("World drop", kind: "world_drop", place: 9, places: places) == "World drop")
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
            "Fury Warrior · needs it", "Combat Rogue · needs it, a runner-up on the list", "Protection Warrior · has it",
        ])
        #expect(rows.first?.lead == .crest(className: "warrior"))
        #expect(note == "3 more have no gear on file.")
        let empty = ItemNeeds(itemId: 1, needs: [], has: [], unknownGear: 0)
        #expect(ItemSheetBuilder.guildRows(empty) == .loaded(rows: [], note: "Nobody in the guild needs it or has it."))
    }

    @Test("the sheet's tooltip lines leave the item level out")
    func lines() throws {
        let request = ItemTooltipRequest.forHome(
            HomeItemLook(id: 252_605, name: "Black Dragonscale Helm", quality: 4, icon: nil, itemLevel: 61),
            home: try HomeFixtures.home(), character: HomeFixtures.velaren
        )
        let model = ItemTooltipModel.build(request, sources: try HomeFixtures.tooltipSources())
        let lines = ItemSheetBuilder.lines(model).map(\.text)
        #expect(!lines.isEmpty)
        #expect(!lines.contains { $0.hasPrefix("Item Level") })
        #expect(!lines.contains(model.look.name ?? ""))
    }
}
