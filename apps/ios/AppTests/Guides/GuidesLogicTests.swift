@testable import ForeverSixty
import Foundation
import Models
import Testing

@Suite("Guides logic")
struct GuidesLogicTests {
    // MARK: F9.1 the tab

    @Test func grid_isTheNineClassesInIndexOrder() throws {
        let cells = GuidesIndexLogic.cells(try GuidesTestData.index())
        #expect(cells.map(\.classSlug) == ["warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid"])
        #expect(cells.first?.name == "Warrior")
    }

    @Test func yourGuide_isTheCharactersSpec() throws {
        let index = try GuidesTestData.index()
        let mine = try #require(GuidesIndexLogic.yourGuide(index, character: GuidesTestData.lightningbolltz))
        #expect(mine.spec == "shaman-elemental")
        #expect(mine.title == "Elemental Shaman")
        #expect(GuideCopy.cardMeta(mine) == "DPS · Updated 1 Nov 2026")
        #expect(GuidesIndexLogic.yourGuide(index, character: nil) == nil)
    }

    @Test func yourGuide_absentForAnUnknownSpec() throws {
        let nobody = GuidesTestData.character(name: "Nospec", className: "shaman", spec: "Tidecaller", race: "orc")
        #expect(GuidesIndexLogic.yourGuide(try GuidesTestData.index(), character: nobody) == nil)
    }

    @Test func classPage_listsTheClassSpecsInTreeOrder() throws {
        let specs = GuidesIndexLogic.specs(of: "shaman", in: try GuidesTestData.index())
        #expect(specs.map(\.spec) == ["shaman-elemental", "shaman-enhancement", "shaman-restoration"])
        #expect(GuideCopy.rowMeta(specs[2]) == "Healer · Updated 1 Nov")
    }

    @Test func saved_footerNamesTheClassesAndTheSize() throws {
        let index = try GuidesTestData.index()
        let saved = GuidesTestData.saved()
        #expect(saved.isSaved("shaman-elemental"))
        #expect(!saved.isSaved("shaman-enhancement"))
        #expect(saved.footer(index) == "Saved on this phone: Warrior, Shaman · 32 KB. Guides update by themselves when the game data does.")
        #expect(SavedGuides(bytes: [:]).footer(index) == "Nothing is saved on this phone yet. Open a guide once and it stays.")
    }

    // MARK: F9.2 a guide

    @Test func sections_areTheGuidesOwnInReadingOrder() throws {
        let guide = try GuidesTestData.elemental()
        let sections = GuideSection.of(guide)
        #expect(sections == [.overview, .talents, .rotation, .stats, .gear, .races])
        #expect(GuideSection.a11y(.talents, in: sections) == "Talents, section 2 of 6")
    }

    @Test func head_namesTheRecommendedRaces() throws {
        let guide = try GuidesTestData.elemental()
        #expect(GuideCopy.title(guide) == "Elemental Shaman")
        #expect(GuideCopy.headMeta(guide) == "Updated 1 Nov 2026 · Recommended races: Dwarf, Orc")
        #expect(GuideCopy.shareURL(guide).absoluteString == "https://foreversixty.gg/guides/shaman/elemental")
    }

    @Test func statChips_areNumberedInOrder() throws {
        let chips = GuideCopy.statChips(try GuidesTestData.elemental())
        #expect(chips.first?.number == 1)
        #expect(chips.first?.stat == "Spell power")
        #expect(chips.first?.a11y == "Priority 1, Spell power")
        #expect(chips.count == 7)
    }

    @Test func talents_rankWordsAndBuildRows() throws {
        let guide = try GuidesTestData.elemental()
        let lavaBurst = try #require(guide.talents.key.first { $0.name == "Lava Burst" })
        #expect(GuideCopy.ranks(lavaBurst) == "1 rank")
        let fury = try #require(guide.talents.key.first { $0.name == "Elemental Fury" })
        #expect(GuideCopy.ranks(fury) == "5 ranks")
        let rows = GuideBuildRows.rows(guide.talents)
        #expect(rows.map(\.title) == ["Raid build", "Leveling build"])
        #expect(rows.map(\.totals) == ["31 / 10 / 10", "30 / 3 / 18"])
    }

    @Test func rotation_numbersEachListAndPrintsRanks() throws {
        let groups = RotationRows.groups(try GuidesTestData.elemental().rotation)
        #expect(groups.map(\.title) == ["Opener", "Priority", "Cooldowns"])
        let priority = try #require(groups.first { $0.title == "Priority" })
        #expect(priority.rows.map(\.number) == Array(1...8))
        #expect(priority.rows.first?.spell.name == "Lava Burst")
        #expect(priority.rows.first?.rankText == "Rank 4")
        let cooldowns = try #require(groups.last)
        #expect(cooldowns.rows.first?.rankText == nil)
        #expect(RotationRows.groups(nil).isEmpty)
    }

    @Test func gearRow_namesTheSpecAndPreset() throws {
        let guide = try GuidesTestData.elemental()
        #expect(GuideCopy.gearRow(guide) == "Best in slot for Elemental · 60 · Raid-ready, Phase 1")
    }

    // MARK: F9.4 tier lists

    @Test func tierChoice_defaultsToTheCharacter() {
        let specs = SpecCatalog.bundled()
        #expect(TierChoice.defaults(GuidesTestData.tank, specs: specs) == TierChoice(role: .tank, faction: .horde))
        #expect(TierChoice.defaults(nil, specs: specs) == TierChoice(role: .dps, faction: .horde))
        let alliance = GuidesTestData.character(name: "Lumen", className: "priest", spec: "Holy", race: "human", faction: "alliance")
        #expect(TierChoice.defaults(alliance, specs: specs) == TierChoice(role: .healer, faction: .alliance))
    }

    @Test func tierRows_keepTheFilesOrderAndMarkYours() throws {
        let list = TiersLogic.list(try GuidesTestData.tiers(), choice: TierChoice(role: .dps, faction: .horde))
        let rows = TiersLogic.rows(list, yourSpec: "mage-fire")
        #expect(rows.map(\.row.rank) == Array(1...rows.count))
        #expect(rows.first?.row.spec == "warrior-fury")
        #expect(rows.first?.gap == "Top")
        #expect(rows.filter(\.isYours).map(\.row.spec) == ["mage-fire"])
        #expect(rows.first?.letter == "S")
    }

    @Test func tierRow_tieMarksTheGap() throws {
        let list = TiersLogic.list(try GuidesTestData.tiers(), choice: TierChoice(role: .dps, faction: .horde))
        let rows = TiersLogic.rows(list, yourSpec: nil)
        let tied = try #require(rows.first { $0.row.tie != nil })
        #expect(tied.gap == "≈ \(tied.row.gapText)")
        #expect(tied.a11y.hasSuffix(", Within 1 percent of the next spec"))
    }

    @Test func tierRow_a11yReadsRankTierAndGap() throws {
        let list = TiersLogic.list(try GuidesTestData.tiers(), choice: TierChoice(role: .dps, faction: .horde))
        let rows = TiersLogic.rows(list, yourSpec: nil)
        #expect(rows[0].a11y == "Fury Warrior, rank 1 of \(rows.count), tier S, top spec, Troll")
        #expect(rows[1].a11y == "Arms Warrior, rank 2 of \(rows.count), tier S, 4.1 percent behind the top, Orc")
    }

    @Test func tankAndHealer_haveNoLettersAndShowTheMetric() throws {
        let file = try GuidesTestData.tiers()
        let tanks = TiersLogic.rows(TiersLogic.list(file, choice: TierChoice(role: .tank, faction: .horde)), yourSpec: nil)
        #expect(tanks.allSatisfy { $0.letter == nil })
        #expect(tanks.first?.metric == "365")
        #expect(tanks.first?.a11y == "Protection Warrior, rank 1 of 3, top spec, Tauren")
    }

    @Test func callout_isTheCharactersRowWithTheLastNote() throws {
        let list = TiersLogic.list(try GuidesTestData.tiers(), choice: TierChoice(role: .dps, faction: .horde))
        let callout = try #require(TiersLogic.callout(list, yourSpec: "warrior-fury"))
        #expect(callout.row.title == "Fury Warrior is 1st of 20 DPS specs.")
        #expect(callout.row.detail == "Level with the top spec.")
        #expect(callout.note == list.notes.last)
        #expect(TiersLogic.callout(list, yourSpec: nil) == nil)
        #expect(TiersLogic.callout(list, yourSpec: "priest-holy") == nil)
    }

    @Test func yourSpec_isTheCatalogKey() {
        #expect(TiersLogic.specKey(GuidesTestData.emberquill, specs: SpecCatalog.bundled()) == "mage-fire")
        #expect(TiersLogic.specKey(nil, specs: SpecCatalog.bundled()) == nil)
    }

    // MARK: F9.5 what changed

    @Test func keyDates_dropThePastSortAndRun() throws {
        let rows = KeyDatesLogic.rows(try GuidesTestData.dates(), now: GuidesTestData.now)
        #expect(rows.map(\.label) == ["Name reservation", "First raids"])
        #expect(rows[0].running == "Name reservation open through Nov 3")
        #expect(rows[1].running == nil)
        #expect(rows[1].key == "Dec 9")
    }

    @Test func changelog_rowsCarryDateAndTag() throws {
        let rows = ChangelogLogic.rows(try GuidesTestData.changelog())
        #expect(rows.first?.date == "10/28")
        #expect(rows.first?.title == "The log coach is here.")
        #expect(rows.first?.tag == "Site")
        #expect(ChangelogLogic.tag("game") == "Game")
        #expect(ChangelogLogic.tag("patch") == "Game")
    }

    // MARK: Sync

    @Test func newData_onlyForAnotherBuild() {
        #expect(NewDataCheck.state(shown: "a", latest: "b") == .ready("b"))
        #expect(NewDataCheck.state(shown: "a", latest: "a") == .none)
        #expect(NewDataCheck.state(shown: "a", latest: nil) == .none)
    }

    @Test func prefetch_isTheCharactersClassesSpecs() throws {
        let specs = GuidesSync.specs(in: try GuidesTestData.index(), classes: ["shaman", "warrior"])
        #expect(specs == ["warrior-arms", "warrior-fury", "warrior-protection", "shaman-elemental", "shaman-enhancement", "shaman-restoration"])
    }

    @Test func sync_runsHourly() {
        let now = GuidesTestData.now
        #expect(GuidesSync.isDue(last: nil, now: now))
        #expect(!GuidesSync.isDue(last: now.addingTimeInterval(-60), now: now))
        #expect(GuidesSync.isDue(last: now.addingTimeInterval(-3601), now: now))
    }
}
