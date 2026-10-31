@testable import ForeverSixty
import Foundation
import Models
import Testing

/// What the Talents screen says and draws (spec F6 §4 to §7). The same list runs on Android
/// (TalentsPresenterTest.kt).
@Suite("Talents presenter")
struct TalentsPresenterTests {
    private let shaman: TalentsFile
    private let warrior: TalentsFile

    init() throws {
        shaman = try TalentTestData.shaman()
        warrior = try TalentTestData.warrior()
    }

    private func presenter(
        yours: String? = TalentTestData.shamanYours,
        guide: TalentGuideBuild = .build(talents: TalentTestData.shamanGuide, band: 60),
        export: TalentExport? = nil,
        level: Int? = 60
    ) -> TalentsPresenter {
        let mine = export ?? yours.map { TalentExport.ranks(TalentRules.ranks(ofBand: $0)) } ?? .none
        return TalentsPresenter(file: shaman, guide: guide, export: mine, level: level, specName: "Elemental")
    }

    private func id(_ name: String, in file: TalentsFile? = nil) throws -> Int {
        try #require(TalentRules.trees(file ?? shaman).flatMap(\.talents).first { $0.name == name }).id
    }

    @Test("Differs: the count sentence, unspent and missing lines, tab totals and dots")
    func differs() {
        let model = presenter()
        #expect(model.status == TalentStatusLine(tag: "Differs", tagKind: .differs, sentence: "2 of your 48 points differ from the 60 guide build."))
        #expect(model.notes == ["3 points unspent.", "The guide build takes 5 points you have not."])
        #expect(model.tabs(.yours).map(\.points) == [28, 3, 17])
        #expect(model.tabs(.yours).map(\.differing) == [3, 0, 1])
        #expect(model.tabs(.guide).map(\.points) == [30, 3, 18])
        #expect(model.tabs(.guide).allSatisfy { $0.differing == 0 })
        #expect(model.showsSegments)
        #expect(model.initialBuild == .yours)
        #expect(model.openingTree == 0)
        #expect(model.legend(.yours) == Copy.talentsLegend)
        #expect(model.legend(.guide) == Copy.talentsTapHint)
    }

    @Test("Matches: the guide sentence, no notes, the tap hint")
    func matches() {
        let model = presenter(yours: TalentTestData.shamanGuide)
        #expect(model.status == TalentStatusLine(tag: "Matches", tagKind: .matches, sentence: "Your build is the Elemental 60 guide build."))
        #expect(model.notes.isEmpty)
        #expect(model.legend(.yours) == Copy.talentsTapHint)
    }

    @Test("A build inside the guide's reads as points still to take")
    func toTake() {
        let model = TalentsPresenter(
            file: warrior, guide: .build(talents: TalentTestData.band60, band: 60),
            export: .ranks(TalentRules.ranks(ofBand: TalentTestData.band50)), level: 60, specName: "Fury"
        )
        #expect(model.status.sentence == "10 of the 60 guide build's points still to take.")
        #expect(model.notes == ["10 points unspent."])
    }

    @Test("Level 58 holding the whole 41-point guide matches, with the points beyond it noted")
    func beyondTheBand() {
        let model = TalentsPresenter(
            file: warrior, guide: .build(talents: TalentTestData.band50, band: 50),
            export: .ranks(TalentTestData.yoursTrees), level: 58, specName: "Fury"
        )
        #expect(model.status == TalentStatusLine(
            tag: "Matches", tagKind: .matches, sentence: "Your build has every point of the Fury 50 to 59 guide build."
        ))
        #expect(model.notes == ["+8 beyond the guide"])
    }

    @Test("Stray points past the guide's total do not differ; the rest do")
    func strayPastTheBand() {
        let model = TalentsPresenter(
            file: warrior, guide: .build(talents: TalentTestData.band50, band: 50),
            export: .ranks(TalentRules.ranks(ofBand: TalentTestData.strayAndShort)), level: 58, specName: "Fury"
        )
        #expect(model.status.sentence == "5 of your 46 points differ from the 50 to 59 guide build.")
        #expect(model.notes == ["3 points unspent.", "The guide build takes 5 points you have not.", "+5 beyond the guide"])
    }

    @Test("No export: the guide alone, no segments, the addon line")
    func noExport() {
        let model = presenter(yours: nil)
        #expect(model.status == TalentStatusLine(tag: "Guide build", tagKind: .muted, sentence: Copy.homeNextTalentsNone))
        #expect(model.builds == [.guide])
        #expect(!model.showsSegments)
        #expect(model.initialBuild == .guide)
        #expect(model.notes.isEmpty)
    }

    @Test("No export and no guide: both lines and no tree")
    func nothingToDraw() {
        let model = presenter(yours: nil, guide: .none)
        #expect(model.status.tag == nil)
        #expect(model.notes == [Copy.talentsNoList])
        #expect(!model.hasTree)
    }

    @Test("A stale export draws the guide with the older-build line")
    func staleExport() {
        let model = presenter(export: .stale)
        #expect(model.status.sentence == Copy.talentsStaleExport)
        #expect(model.builds == [.guide])
    }

    @Test("No build: your build alone with the no-list line and the unspent line")
    func noBuild() {
        let model = presenter(guide: .none)
        #expect(model.status == TalentStatusLine(tag: "Your build", tagKind: .muted, sentence: Copy.talentsNoList))
        #expect(model.builds == [.yours])
        #expect(model.notes == ["3 points unspent."])
        #expect(model.cells(.yours, tree: 0).allSatisfy { $0.delta == nil })
    }

    @Test("No spec: your build alone with the spec line")
    func noSpec() {
        let model = presenter(guide: .noSpec)
        #expect(model.status.sentence == Copy.talentsPickSpec)
        #expect(!model.showsSegments)
    }

    @Test("The guide failed: your build with the failed line and no compare")
    func guideFailed() {
        let model = presenter(guide: .failed)
        #expect(model.status.isError)
        #expect(model.status.sentence == Copy.talentsGuideFailed)
        #expect(!model.showsSegments)
    }

    @Test("Cells on Your build carry the signed chip and the guide rank; the guide's carry none")
    func cellMarks() throws {
        let model = presenter()
        let cells = Dictionary(uniqueKeysWithValues: model.cells(.yours, tree: 0).map { ($0.talent.name, $0) })
        let reverberation = try #require(cells["Reverberation"])
        let callOfFlame = try #require(cells["Call of Flame"])
        #expect(reverberation.chip == "+2")
        #expect(callOfFlame.chip == "\u{2212}2")
        #expect(callOfFlame.accessibilityLabel == "Call of Flame, rank 1 of 3, guide rank 3")
        #expect(try #require(cells["Convection"]).accessibilityLabel == "Convection, rank 5 of 5")
        #expect(model.cells(.guide, tree: 0).allSatisfy { $0.chip == nil })
    }

    @Test("The sheet: rank, current and next text, and the guide line only where ranks differ")
    func sheet() throws {
        let model = presenter()
        let sheet = try #require(model.sheet(.yours, talentID: try id("Call of Flame")))
        #expect(sheet.rankLine == "Rank 1/3")
        #expect(sheet.current == sheet.talent.ranks[0].description)
        #expect(sheet.next == sheet.talent.ranks[1].description)
        #expect(sheet.guideLine == "Guide build: rank 3/3")
        #expect(!sheet.isOver)
        #expect(try #require(model.sheet(.yours, talentID: try id("Reverberation"))).isOver)
        #expect(try #require(model.sheet(.guide, talentID: try id("Call of Flame"))).guideLine == nil)
        let convection = try #require(model.sheet(.yours, talentID: try id("Convection")))
        #expect(convection.next == nil)
        #expect(convection.guideLine == nil)
    }

    @Test("The sheet names unmet requirements in the game's words")
    func requirements() throws {
        let model = presenter(guide: .none, export: .ranks([[], [], []]))
        let tree = try #require(TalentRules.trees(shaman).first)
        let gated = try #require(tree.talents.first { $0.tier >= 2 && $0.prereqTalentId != nil })
        let sheet = try #require(model.sheet(.yours, talentID: gated.id))
        #expect(sheet.requirements.first == "Requires \(gated.tier * 5) points in \(tree.name) Talents")
        #expect(sheet.requirements.count == 2)
        let open = try #require(tree.talents.first { $0.tier == 0 })
        #expect(try #require(model.sheet(.yours, talentID: open.id)).requirements.isEmpty)
    }

    @Test("Home's Talents row says what the Talents screen says")
    func homeRowFollowsTheScreen() throws {
        let status = presenter().status
        let home = HomePresenter(
            character: HomeFixtures.velaren, finishedSims: [], time: HomeTime(now: HomeFixtures.now), talentStatus: status
        )
        let row = try #require(home.model(try HomeFixtures.home()).next.first { $0.kind == .talents })
        #expect(row.line == "2 of your 48 points differ from the 60 guide build.")
        #expect(row.tag == .differs)
        let noExport = presenter(yours: nil).status
        let muted = HomePresenter(character: HomeFixtures.velaren, finishedSims: [], time: HomeTime(now: HomeFixtures.now), talentStatus: noExport)
        #expect(muted.nextRowsWithoutAnswer().first == HomeNextRow(kind: .talents, title: Copy.homeNextTalents, line: Copy.homeNextTalentsNone, tag: nil))
    }
}
