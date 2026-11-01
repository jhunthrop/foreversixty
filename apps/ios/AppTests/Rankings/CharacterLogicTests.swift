@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// The character page's rules (spec F11.4 to F11.7) and the shared Paperdoll model (F11.5, F2.7).
@Suite("Character page logic")
struct CharacterLogicTests {
    private let now = RankingsTestData.now
    private let clock = RankingsTestData.clock

    // MARK: head

    @Test func head_descriptorAndWhere() throws {
        let voltaris = try RankingsTestData.voltaris()
        #expect(CharacterPage.descriptor(voltaris, specs: RankingsTestData.specs) == "Orc Elemental Shaman · Level 60")
        #expect(CharacterPage.whereLine(voltaris.character) == "Hordebreakers · US Normal")
        let fresh = try RankingsTestData.freshling()
        #expect(CharacterPage.descriptor(fresh, specs: RankingsTestData.specs) == "Mage")
        #expect(CharacterPage.whereLine(fresh.character) == "US PvP")
    }

    @Test func yours_isAnyCharacterOfTheAccount() throws {
        let keys = Set(RankingsTestData.account().characterKeys)
        #expect(CharacterPage.isYours(try RankingsTestData.lightningbolltz().character, keys: keys))
        #expect(!CharacterPage.isYours(try RankingsTestData.voltaris().character, keys: keys))
    }

    // MARK: Forever Score

    @Test func scoreBars_areValueOverCap() throws {
        let bars = CharacterPage.bars(try RankingsTestData.voltaris())
        #expect(bars.map(\.label) == ["Gear", "Experience", "Parses", "Attendance", "Discipline"])
        #expect(bars.map(\.value) == [412, 620, 540, 180, 78])
        #expect(bars.map(\.cap) == [500, 700, 600, 200, 100])
        #expect(abs(bars[0].fraction - 0.824) < 0.0001)
        #expect(bars[4].fraction == 0.78)
        #expect(bars[0].a11y == "Gear 412 of 500.")
    }

    @Test func scoreBar_fractionIsClamped_andAZeroCapIsEmpty() {
        #expect(ScoreBar(label: "Gear", value: 600, cap: 500).fraction == 1)
        #expect(ScoreBar(label: "Gear", value: 10, cap: 0).fraction == 0)
        #expect(ScoreBar(label: "Gear", value: -5, cap: 100).fraction == 0)
    }

    @Test func score_wordsAndTier() throws {
        let voltaris = try RankingsTestData.voltaris()
        #expect(CharacterPage.score(voltaris) == "1,830")
        #expect(CharacterPage.tierName(voltaris.tier) == "Epic")
        #expect(CharacterPage.tierQuality(voltaris.tier) == 4)
        #expect(CharacterPage.asOf(voltaris.asOf) == "As of 17 Dec")
        #expect(CharacterPage.scoreA11y(voltaris) == "Forever Score 1,830, Epic.")
        #expect(CharacterPage.hasScore(voltaris))
        #expect(!CharacterPage.hasScore(try RankingsTestData.freshling()))
    }

    @Test func sparkline_oldestFirst_normalized() throws {
        let points = CharacterPage.sparkline(try RankingsTestData.voltaris().scoreHistory.reversed())
        #expect(points.first == 0)
        #expect(points.last == 1)
        #expect(points.count == 5)
        #expect(CharacterPage.sparkline([]) == [])
    }

    // MARK: bests, recent, builds

    @Test func bests() throws {
        let rows = CharacterPage.bests(try RankingsTestData.voltaris(), now: now)
        #expect(rows.map(\.encounter) == ["Onyxia", "Emperor Dagran Thaurissan"])
        #expect(rows[0].meta == "Normal · 2 d ago")
        #expect(rows[0].value == "462")
        #expect(rows[0].unit == "DPS")
        #expect(rows[0].percentile == 99)
        #expect(rows[1].percentile == 81)
    }

    @Test func recent_killAndWipe() throws {
        let rows = CharacterPage.recent(try RankingsTestData.voltaris(), now: now)
        #expect(rows.map(\.kill) == [true, false, true])
        #expect(rows[0].meta == "Kill · 2 d ago")
        #expect(rows[1].meta == "Wipe · 2 d ago")
        #expect(rows[0].reportID == "onyx0001kill")
    }

    @Test func builds() throws {
        let rows = CharacterPage.builds(try RankingsTestData.voltaris(), clock: clock)
        #expect(rows == [BuildLine(trees: "30 / 0 / 21", meta: "Elemental · first seen 10 Dec")])
    }

    @Test func gap() throws {
        #expect(CharacterPage.gap(try RankingsTestData.voltaris().gearGapTotal) == "3.0 points")
        #expect(CharacterPage.gap(try RankingsTestData.lightningbolltz().gearGapTotal) == "36.0 points")
        #expect(CharacterPage.gap(nil) == nil)
    }

    // MARK: Paperdoll

    @Test func paperdoll_character_fullSet_andGoldDots() throws {
        let slots = PaperdollModel.character(gear: try RankingsTestData.voltaris().gear, items: try RankingsTestData.items())
        let frame = PaperdollModel.frame(slots)
        #expect(frame.left.map(\.slot) == ["head", "neck", "shoulder", "back", "chest", "wrist"])
        #expect(frame.right.map(\.slot) == ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"])
        #expect(frame.weapons.map(\.slot) == ["main_hand", "off_hand", "ranged"])
        #expect(frame.all.count == 17)
        // Voltaris wears his own list: one dot, Trinket 2.
        #expect(slots.filter(\.better).map(\.slot) == ["trinket2"])
        guard case let .item(look) = slots[0].state else { Issue.record("head not an item"); return }
        #expect(look.name == "Living Crown")
        #expect(look.quality == 4)
        #expect(look.icon == "inv_helmet_32")
        #expect(look.itemLevel == 61)
        let behind = PaperdollModel.character(gear: try RankingsTestData.lightningbolltz().gear, items: try RankingsTestData.items())
        #expect(behind.filter(\.better).map(\.slot) == ["head", "shoulder", "chest", "waist", "legs", "trinket2"])
    }

    @Test func paperdoll_neverADot_whenThePickIsWorn() throws {
        var gear = try RankingsTestData.lightningbolltz().gear
        // Trinket 2's pick (11302) worn in trinket 1: the pick is already worn, either way round.
        gear[12] = CharacterGearSlot(slot: "trinket1", worn: CharacterGearSlotWorn(id: 11302, name: "Uther's Strength", itemLevel: 60), pick: nil, gap: nil)
        // Head's pick worn in head itself.
        gear[0] = CharacterGearSlot(slot: "head", worn: CharacterGearSlotWorn(id: 252561, name: "Living Crown", itemLevel: 61), pick: gear[0].pick, gap: gear[0].gap)
        let slots = PaperdollModel.character(gear: gear, items: try RankingsTestData.items())
        #expect(slots.filter(\.better).map(\.slot) == ["shoulder", "chest", "waist", "legs"])
    }

    @Test func paperdoll_twoHander_emptiesTheOffHandWithWords() throws {
        let staff = try #require(RankingsTestData.items()[2243])
        let items = try RankingsTestData.items().merging([2243: staff.asTwoHand()]) { _, new in new }
        let profile = try RankingsTestData.voltaris { data in
            var gear = data["gear"] as? [[String: Any]] ?? []
            gear = gear.filter { ($0["slot"] as? String) != "off_hand" }
            data["gear"] = gear
        }
        let off = try #require(PaperdollModel.character(gear: profile.gear, items: items).first { $0.slot == "off_hand" })
        #expect(off.state == .twoHandOff)
        #expect(PaperdollModel.name(off) == "Two-hand in main hand")
        #expect(PaperdollModel.a11y(off) == "Off hand, Two-hand in main hand.")
        #expect(!off.better)
    }

    @Test func paperdoll_unknownAndEmpty() throws {
        let profile = try RankingsTestData.voltaris { data in
            var gear = data["gear"] as? [[String: Any]] ?? []
            gear[1]["worn"] = ["id": 288_887, "name": NSNull(), "item_level": NSNull()]
            gear = gear.filter { ($0["slot"] as? String) != "ranged" }
            data["gear"] = gear
        }
        let slots = PaperdollModel.character(gear: profile.gear, items: try RankingsTestData.items())
        let neck = try #require(slots.first { $0.slot == "neck" })
        #expect(neck.state == .unknown(288_887))
        #expect(PaperdollModel.name(neck) == "Neck · Item 288887")
        #expect(!PaperdollModel.showsSlotLine(neck))
        #expect(PaperdollModel.a11y(neck) == "Neck · Item 288887.")
        #expect(PaperdollModel.showsSlotLine(slots[0]))
        #expect(slots.first { $0.slot == "ranged" }?.state == .empty)
        // An empty or unknown slot carries no dot without a pick that gains.
        #expect(!neck.better)
        let noGain = CharacterGearSlot(slot: "ranged", worn: nil, pick: CharacterGearSlotPick(id: 1, name: "X", source: "Y", chance: nil), gap: nil)
        #expect(PaperdollModel.character(gear: [noGain], items: [:]).first { $0.slot == "ranged" }?.better == false)
    }

    @Test func paperdoll_a11y() throws {
        let slots = PaperdollModel.character(gear: try RankingsTestData.voltaris().gear, items: try RankingsTestData.items())
        #expect(PaperdollModel.a11y(slots[0]) == "Head, Living Crown, Epic, item level 61.")
        #expect(PaperdollModel.a11y(slots[1]) == "Neck, Archlight Talisman, Rare, item level 60.")
        let trinket = try #require(slots.first { $0.slot == "trinket2" })
        #expect(PaperdollModel.a11y(trinket).hasSuffix(" A better item is on your best in slot list."))
        let empty = PaperdollSlot(slot: "ranged", state: .empty, better: false)
        #expect(PaperdollModel.a11y(empty) == "Ranged, empty.")
    }

    @Test func paperdoll_fromHome_dotsEveryUpgrade() throws {
        let home = try HomeFixtures.home()
        let sources = try HomeFixtures.tooltipSources()
        let rows = GearModel.rows(home: home, gear: [:], items: sources.items, offHandBlockedBy: "Arcanite Reaper")
        let slots = PaperdollModel.home(rows)
        #expect(slots.count == 17)
        // Only measured gains carry a dot: the unknown Wrist and the empty Trinket 2 do not.
        #expect(Set(slots.filter(\.better).map(\.slot)) == Set(rows.filter(\.gains).map(\.slot)))
        #expect(slots.first { $0.slot == "wrist" }?.better == false)
        #expect(slots.first { $0.slot == "trinket2" }?.better == false)
        #expect(slots.first { $0.slot == "wrist" }.map(PaperdollModel.name) == "Wrist · Item 999999")
        #expect(slots.first { $0.slot == "off_hand" }?.state == .twoHandOff)
    }
}

extension Item {
    /// The same item, two-handed: a test's two-hander from a one-hand fixture.
    func asTwoHand() throws -> Item {
        var object = try JSONSerialization.jsonObject(with: WireFormat.encoder().encode(self)) as? [String: Any] ?? [:]
        object["two_hand"] = true
        return try RankingsTestData.decode(Item.self, object)
    }
}
