@testable import ForeverSixty
import Foundation
import Models
import Testing

/// The planner's rules on the app (spec F6 §4; web/src/lib/planner band-compare, connectors,
/// rules, styles), over the real talent files. The same list runs on Android (TalentRulesTest.kt).
@Suite("Talent rules")
struct TalentRulesTests {
    private let warrior: TalentsFile
    private let shaman: TalentsFile
    private let arms: TalentTree

    init() throws {
        warrior = try TalentTestData.warrior()
        shaman = try TalentTestData.shaman()
        arms = try #require(TalentRules.trees(warrior).first)
    }

    @Test("A tree string reads one base-36 digit per talent, a bad digit as 0")
    func decodesTreeDigits() {
        #expect(TalentRules.ranks(ofTree: "503a") == [5, 0, 3, 10])
        #expect(TalentRules.ranks(ofTree: "5?1") == [5, 0, 1])
    }

    @Test("A band string splits into three trees, empty trees kept")
    func decodesBandString() {
        #expect(TalentRules.ranks(ofBand: "32-05-") == [[3, 2], [0, 5], []])
    }

    @Test("Ranks clamp to each talent's maximum; a short string and a missing tree read as 0")
    func clampsAndPadsRanks() throws {
        let ranks = TalentRules.rankByID(warrior, treeRanks: [[9]])
        let heroicStrike = try #require(arms.talents.first)
        #expect(ranks[heroicStrike.id] == heroicStrike.maxRank)
        #expect(ranks.values.reduce(0, +) == heroicStrike.maxRank)
        #expect(ranks.count == TalentRules.trees(warrior).flatMap(\.talents).count)
    }

    @Test("The spec's vector: diffCount 2, missingCount 5, four dashed cells, totals 28 3 17")
    func specVector() {
        let guide = TalentRules.rankByID(shaman, treeRanks: TalentRules.ranks(ofBand: TalentTestData.shamanGuide))
        let yours = TalentRules.rankByID(shaman, treeRanks: TalentRules.ranks(ofBand: TalentTestData.shamanYours))
        let diff = TalentRules.compare(shaman, yours: yours, guide: guide)
        #expect(diff.diffCount == 2)
        #expect(diff.missingCount == 5)
        let names = Set(TalentRules.trees(shaman).flatMap(\.talents).filter { diff.differing.contains($0.id) }.map(\.name))
        #expect(names == ["Reverberation", "Call of Flame", "Elemental Warding", "Tidal Focus"])
        #expect(diff.status == .differs(2))
        #expect(TalentRules.pointsPerTree(shaman, ranks: yours) == [28, 3, 17])
        #expect(TalentRules.pointsAvailable(atLevel: 60) - yours.values.reduce(0, +) == 3)
    }

    @Test("The same build matches, with nothing marked, totals 30 3 18")
    func sameBuildMatches() {
        let guide = TalentRules.rankByID(shaman, treeRanks: TalentRules.ranks(ofBand: TalentTestData.shamanGuide))
        let diff = TalentRules.compare(shaman, yours: guide, guide: guide)
        #expect(diff == TalentDiff(diffCount: 0, missingCount: 0, extraCount: 0, differing: []))
        #expect(diff.status == .matches)
        #expect(TalentRules.pointsPerTree(shaman, ranks: guide) == [30, 3, 18])
    }

    @Test("A build inside the guide's and short of it reads as points still to take")
    func toTake() {
        let guide = TalentRules.rankByID(warrior, treeRanks: TalentRules.ranks(ofBand: TalentTestData.band60))
        let yours = TalentRules.rankByID(warrior, treeRanks: TalentRules.ranks(ofBand: TalentTestData.band50))
        let diff = TalentRules.compare(warrior, yours: yours, guide: guide)
        #expect(diff.diffCount == 0)
        #expect(diff.extraCount == 0)
        #expect(diff.status == .toTake(10))
    }

    @Test("Level 58: 49 points holding the whole 41-point guide match, 8 beyond it")
    func beyondTheGuideMatches() {
        let guide = TalentRules.rankByID(warrior, treeRanks: TalentRules.ranks(ofBand: TalentTestData.band50))
        let yours = TalentRules.rankByID(warrior, treeRanks: TalentTestData.yoursTrees)
        let diff = TalentRules.compare(warrior, yours: yours, guide: guide)
        #expect(diff.diffCount == 0)
        #expect(diff.missingCount == 0)
        #expect(diff.extraCount == 8)
        #expect(diff.status == .matches)
    }

    @Test("Stray points differ only up to the guide's total: 10 stray − (46 − 41) = 5")
    func strayCappedAtGuideTotal() {
        let guide = TalentRules.rankByID(warrior, treeRanks: TalentRules.ranks(ofBand: TalentTestData.band50))
        let yours = TalentRules.rankByID(warrior, treeRanks: TalentRules.ranks(ofBand: TalentTestData.strayAndShort))
        let diff = TalentRules.compare(warrior, yours: yours, guide: guide)
        #expect(diff.diffCount == 5)
        #expect(diff.missingCount == 5)
        #expect(diff.extraCount == 5)
        #expect(diff.status == .differs(5))
    }

    @Test("Cells read maxed, filled, open and locked")
    func cellStates() throws {
        let ranks = TalentRules.rankByID(warrior, treeRanks: TalentTestData.yoursTrees)
        let byName = Dictionary(uniqueKeysWithValues: arms.talents.map { ($0.name, $0) })
        let state = { (name: String) throws -> TalentCellState in
            TalentRules.cellState(try #require(byName[name]), in: arms, ranks: ranks)
        }
        #expect(try state("Improved Heroic Strike") == .maxed)
        #expect(try state("Deflection") == .filled)
        #expect(try state("Improved Tactical Mastery") == .open)
        #expect(try state("Anger Management") == .locked) // prerequisite unmet
        #expect(try state("Bloodthrill") == .locked) // tier 4 needs 20 points
    }

    @Test("Prerequisite lines follow connectorsFor, gold once met")
    func connectors() {
        let ranks = TalentRules.rankByID(warrior, treeRanks: TalentTestData.yoursTrees)
        let lines = TalentRules.connectors(arms, ranks: ranks)
        #expect(lines.count == 4)
        #expect(lines.allSatisfy { $0.shape == .down })
        #expect(lines.filter(\.met).map(\.toID).sorted() == [105_947, 105_950])
    }

    @Test("A line is down a column, across a row, or down then across")
    func connectorShapes() {
        let line = { (from: TalentCell, to: TalentCell) in
            TalentConnector(from: from, to: to, fromID: 1, toID: 2, met: false).shape
        }
        #expect(line(TalentCell(tier: 0, column: 1), TalentCell(tier: 2, column: 1)) == .down)
        #expect(line(TalentCell(tier: 1, column: 0), TalentCell(tier: 1, column: 1)) == .across)
        #expect(line(TalentCell(tier: 1, column: 0), TalentCell(tier: 2, column: 2)) == .elbow)
    }

    @Test("Points available follow the level, capped at 51")
    func pointsAvailable() {
        #expect(TalentRules.pointsAvailable(atLevel: 9) == 0)
        #expect(TalentRules.pointsAvailable(atLevel: 58) == 49)
        #expect(TalentRules.pointsAvailable(atLevel: 70) == 51)
    }

    @Test("The grid is as large as the tree's last tier and column")
    func gridSize() {
        let size = TalentRules.gridSize(arms)
        #expect(size.tiers == 7)
        #expect(size.columns == 4)
    }

    @Test("The tree with the most points opens, ties to the first")
    func openingTree() {
        #expect(TalentRules.openingTree([17, 32, 0]) == 1)
        #expect(TalentRules.openingTree([10, 10, 0]) == 0)
        #expect(TalentRules.openingTree([0, 0, 0]) == 0)
    }

    @Test("The band follows the level: 60 from 60, the decade from 20, none under 20")
    func bandForLevel() {
        #expect(TalentGuide.band(forLevel: 19) == nil)
        #expect(TalentGuide.band(forLevel: 20) == 20)
        #expect(TalentGuide.band(forLevel: 59) == 50)
        #expect(TalentGuide.band(forLevel: 60) == 60)
        #expect(TalentGuide.band(forLevel: nil) == nil)
    }

    @Test("The entry is the raid preset when the band has one, else bare, for the faction")
    func entryChoice() throws {
        let file = try TalentTestData.fury()
        #expect(TalentGuide.entry(in: file, band: 60, faction: "horde")?.preset == "raid")
        #expect(TalentGuide.entry(in: file, band: 50, faction: "horde")?.preset == "bare")
        #expect(TalentGuide.entry(in: file, band: 40, faction: "horde") == nil)
        #expect(TalentGuide.entry(in: file, band: 60, faction: "alliance") == nil)
    }

    @Test("The band label is 60, or the span below it")
    func bandLabel() {
        #expect(TalentGuide.label(band: 60) == "60")
        #expect(TalentGuide.label(band: 50) == "50 to 59")
    }

    @Test("An export reads on another build only when the layouts match")
    func staleLayout() {
        #expect(TalentGuide.sameLayout(warrior, warrior))
        #expect(!TalentGuide.sameLayout(warrior, shaman))
    }
}
