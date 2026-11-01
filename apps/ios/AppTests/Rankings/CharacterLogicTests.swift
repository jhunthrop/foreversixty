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
        #expect(rows[0].meta == "Normal · 1 d ago")
        #expect(rows[0].value == "462")
        #expect(rows[0].unit == "DPS")
        #expect(rows[0].percentile == 99)
        #expect(rows[1].percentile == 81)
    }

    @Test func recent_killAndWipe() throws {
        let rows = CharacterPage.recent(try RankingsTestData.voltaris(), now: now)
        #expect(rows.map(\.kill) == [true, false, true])
        #expect(rows[0].meta == "Kill · 1 d ago")
        #expect(rows[1].meta == "Wipe · 1 d ago")
        #expect(rows[0].reportID == "onyx0001kill")
    }

    @Test func builds() throws {
        let rows = CharacterPage.builds(try RankingsTestData.voltaris(), clock: clock)
        #expect(rows == [BuildLine(trees: "30 / 0 / 21", meta: "Elemental · first seen 10 Dec")])
    }

    @Test func gap() throws {
        #expect(CharacterPage.gap(try RankingsTestData.voltaris().gearGapTotal) == "22.7 points")
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
        #expect(slots.filter(\.better).map(\.slot) == ["head", "chest", "waist", "trinket2"])
        guard case let .item(look) = slots[0].state else { Issue.record("head not an item"); return }
        #expect(look.name == "Blue Dragonscale Helm")
        #expect(look.quality == 4)
        #expect(look.icon == "inv_helmet_09")
        #expect(look.itemLevel == 61)
    }

    @Test func paperdoll_neverADot_whenThePickIsWorn() throws {
        let profile = try RankingsTestData.voltaris { data in
            var gear = data["gear"] as? [[String: Any]] ?? []
            // Trinket 2's pick (11302) worn in trinket 1: the pick is already worn, either way round.
            if var worn = gear[12]["worn"] as? [String: Any] { worn["id"] = 11302; gear[12]["worn"] = worn }
            // Head's pick worn in head itself.
            if var worn = gear[0]["worn"] as? [String: Any] { worn["id"] = 252561; gear[0]["worn"] = worn }
            data["gear"] = gear
        }
        let slots = PaperdollModel.character(gear: profile.gear, items: try RankingsTestData.items())
        #expect(slots.filter(\.better).map(\.slot) == ["chest", "waist"])
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
        #expect(off.state == .twoHandOff("none: Hand of Edward the Odd is two-handed."))
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
        #expect(slots.first { $0.slot == "neck" }?.state == .unknown(288_887))
        #expect(slots.first { $0.slot == "ranged" }?.state == .empty)
    }

    @Test func paperdoll_a11y() throws {
        let slots = PaperdollModel.character(gear: try RankingsTestData.voltaris().gear, items: try RankingsTestData.items())
        #expect(PaperdollModel.a11y(slots[0]) == "Head, Blue Dragonscale Helm, Epic, item level 61. A better item is on your best in slot list.")
        #expect(PaperdollModel.a11y(slots[1]) == "Neck, Archlight Talisman, Rare, item level 60.")
        let empty = PaperdollSlot(slot: "ranged", state: .empty, better: false)
        #expect(PaperdollModel.a11y(empty) == "Ranged, empty.")
    }

    @Test func paperdoll_fromHome_dotsEveryUpgrade() throws {
        let home = try HomeFixtures.home()
        let sources = try HomeFixtures.tooltipSources()
        let rows = GearModel.rows(home: home, gear: [:], items: sources.items, offHandBlockedBy: "Arcanite Reaper")
        let slots = PaperdollModel.home(rows)
        #expect(slots.count == 17)
        #expect(Set(slots.filter(\.better).map(\.slot)) == Set(rows.filter { $0.pick != nil }.map(\.slot)))
        let off = try #require(slots.first { $0.slot == "off_hand" })
        if rows.first(where: { $0.slot == "off_hand" })?.blockedByTwoHand == true {
            #expect(off.state == .twoHandOff("none: Arcanite Reaper is two-handed."))
        }
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
