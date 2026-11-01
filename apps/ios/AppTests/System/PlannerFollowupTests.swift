@testable import ForeverSixty
import Foundation
import Models
import Networking
import Routing
import Testing

@Suite("Planner follow-ups: Open in planner and Compare with mine")
struct PlannerFollowupTests {
    @Test("Open in planner opens the class's planner, class only")
    func plannerLink() throws {
        let profile = try RankingsTestData.lightningbolltz()
        #expect(CharacterPage.plannerLink(profile) == .planner(className: "shaman"))
        #expect(CharacterPage.plannerLink(profile)?.url.absoluteString.hasSuffix("/app/planner/shaman") == true)
    }

    @Test("the compare's words: the guide build of a band, or a shared build")
    func words() {
        let guide = TalentCompareWords.guide(spec: "Elemental", band: "60")
        #expect(guide.side == "Guide build")
        #expect(guide.toTake(3) == "3 of the 60 guide build's points still to take.")
        let shared = TalentCompareWords.sharedBuild
        #expect(shared.side == "Shared build")
        #expect(shared.match(beyond: false) == "Your build is this shared build.")
        #expect(shared.differ(n: 2, of: 51) == "2 of your 51 points differ from the shared build.")
        #expect(shared.toTake(3) == "3 of the shared build's points still to take.")
        #expect(shared.beyond(2) == "+2 beyond the shared build")
        #expect(shared.otherRank(1, max: 5) == "Shared build: rank 1/5")
    }

    @Test("shown with an export of the build's class; hidden without")
    func shown() throws {
        let file = try PlannerTestData.shaman()
        let order = PlannerTestData.elementalOrder(PlannerIndex(file))
        let shaman = try SystemStory.exporter(dropping: 0, file: file, order: order)
        let warrior = try SystemStory.ownCharacters()[1]
        #expect(SharedBuildCompare.yourCharacter(classSlug: "shaman", current: warrior, characters: [warrior, shaman]) == shaman)
        #expect(SharedBuildCompare.yourCharacter(classSlug: "shaman", current: nil, characters: [warrior]) == nil)
        #expect(SharedBuildCompare.yourCharacter(classSlug: "mage", current: shaman, characters: [shaman]) == nil)
    }

    @Test("the compare reads the shared build in the guide's place, with its words")
    func compare() throws {
        let file = try PlannerTestData.shaman()
        let index = PlannerIndex(file)
        let order = PlannerTestData.elementalOrder(index)
        let summary = PlannerBuildSummary(
            build: PlannerTestData.build(order: order), classSlug: "shaman", className: "Shaman",
            raceName: "Orc", raceSlug: "orc", totals: PlannerRules.pointsPerTree(index, order: order)
        )
        let content = SharedBuildContent(summary: summary, file: file)
        let behind = SharedBuildCompare.presenter(
            content: content, file: file, character: try SystemStory.exporter(dropping: 3, file: file, order: order)
        )
        #expect(behind.words == .sharedBuild)
        #expect(behind.status.sentence == "3 of the shared build's points still to take.")
        let same = SharedBuildCompare.presenter(
            content: content, file: file, character: try SystemStory.exporter(dropping: 0, file: file, order: order)
        )
        #expect(same.status.sentence == "Your build is this shared build.")
    }
}
