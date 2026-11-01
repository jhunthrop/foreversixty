@testable import ForeverSixty
import Foundation
import Models
import Testing

/// The planner's golden vectors (spec F13 "Rules live in one place"): apps/shared/fixtures/
/// planner-rules.json, generated from web/src/lib/planner, read in place. Every vector is
/// replayed; any disagreement fails CI.
@Suite("Planner golden vectors")
struct PlannerGoldenTests {
    struct Golden: Decodable {
        struct Constants: Decodable {
            let maxPoints: Int
            let pointsPerTier: Int
            let firstPointLevel: Int
        }

        struct GoldenTalent: Decodable {
            let id: Int
            let name: String
            let tier: Int
            let column: Int
            let maxRank: Int
            let prereqTalentId: Int?
            let prereqRank: Int?
        }

        struct GoldenTree: Decodable {
            let id: Int
            let name: String
            let position: Int
            let talents: [GoldenTalent]
        }

        struct GoldenClass: Decodable {
            let classSlug: String
            let trees: [GoldenTree]
        }

        struct Decision: Decodable {
            let classSlug: String
            let level: Int
            let order: [Int]
            let talent: Int
            let ok: Bool
            let reason: String?
        }

        struct FieldError: Decodable, Equatable {
            let field: String
            let message: String
        }

        struct Validate: Decodable {
            let classSlug: String
            let order: [Int]
            let errors: [FieldError]
        }

        struct Split: Decodable {
            let classSlug: String
            let order: [Int]
            let totals: [Int]
        }

        struct Available: Decodable {
            let level: Int
            let points: Int
        }

        struct FS1: Decodable {
            let code: String
            let classSlug: String
            let raceSlug: String
            let treeRanks: [[Int]]
            let order: [Int]
            let dropped: [Int]
        }

        let build: String
        let constants: Constants
        let noneLeft: String
        let classes: [GoldenClass]
        let pointsAvailable: [Available]
        let add: [Decision]
        let remove: [Decision]
        let validate: [Validate]
        let split: [Split]
        let fs1: [FS1]
    }

    static let golden: Golden = {
        let url = BisTestData.repo.appending(path: "apps/shared/fixtures/planner-rules.json")
        do {
            let decoder = JSONDecoder()
            decoder.keyDecodingStrategy = .convertFromSnakeCase
            return try decoder.decode(Golden.self, from: Data(contentsOf: url))
        } catch {
            fatalError("planner-rules.json did not load: \(error)")
        }
    }()

    static let indexes: [String: PlannerIndex] = Dictionary(uniqueKeysWithValues: golden.classes.map { ($0.classSlug, PlannerIndex(file($0))) })

    /// A talent file from the vectors' slim trees: no icons or rank text, which no rule reads.
    static func file(_ golden: Golden.GoldenClass) -> TalentsFile {
        TalentsFile(build: Self.golden.build, classId: 0, classSlug: golden.classSlug, trees: golden.trees.map { tree in
            TalentTree(id: tree.id, name: tree.name, position: tree.position, background: "", icon: "", talents: tree.talents.map {
                Talent(
                    id: $0.id, name: $0.name, icon: "", maxRank: $0.maxRank, tier: $0.tier, column: $0.column,
                    prereqTalentId: $0.prereqTalentId, prereqRank: $0.prereqRank, spellId: 0
                )
            })
        })
    }

    private func index(_ slug: String) throws -> PlannerIndex { try #require(Self.indexes[slug]) }

    @Test("the constants and the budget sentence match the web's")
    func constants() {
        #expect(Self.golden.constants.maxPoints == TalentRules.maxPoints)
        #expect(Self.golden.constants.pointsPerTier == TalentRules.pointsPerTier)
        #expect(Self.golden.constants.firstPointLevel == PlannerRules.firstPointLevel)
        #expect(Self.golden.noneLeft == Copy.planRefuseNoneLeft)
    }

    @Test("every points-available vector")
    func pointsAvailable() {
        #expect(!Self.golden.pointsAvailable.isEmpty)
        for vector in Self.golden.pointsAvailable {
            #expect(PlannerRules.pointsAvailable(atLevel: vector.level) == vector.points, "level \(vector.level)")
        }
    }

    @Test("every add vector: the budget, then canAddPoint, word for word")
    func add() throws {
        #expect(!Self.golden.add.isEmpty)
        for (number, vector) in Self.golden.add.enumerated() {
            let decision = PlannerRules.add(try index(vector.classSlug), order: vector.order, talent: vector.talent, level: vector.level)
            #expect((decision == .allowed) == vector.ok, "add #\(number)")
            #expect(decision.reason == vector.reason, "add #\(number)")
        }
    }

    @Test("every remove vector: canRemovePoint, word for word")
    func remove() throws {
        #expect(!Self.golden.remove.isEmpty)
        for (number, vector) in Self.golden.remove.enumerated() {
            let decision = PlannerRules.canRemovePoint(try index(vector.classSlug), order: vector.order, talent: vector.talent)
            #expect((decision == .allowed) == vector.ok, "remove #\(number)")
            #expect(decision.reason == vector.reason, "remove #\(number)")
        }
    }

    @Test("every validate vector: the field errors, in order")
    func validate() throws {
        #expect(!Self.golden.validate.isEmpty)
        for (number, vector) in Self.golden.validate.enumerated() {
            let errors = PlannerRules.validate(try index(vector.classSlug), order: vector.order)
            #expect(errors.map { Golden.FieldError(field: $0.field, message: $0.message) } == vector.errors, "validate #\(number)")
        }
    }

    @Test("every split vector")
    func split() throws {
        #expect(!Self.golden.split.isEmpty)
        for (number, vector) in Self.golden.split.enumerated() {
            #expect(PlannerRules.pointsPerTree(try index(vector.classSlug), order: vector.order) == vector.totals, "split #\(number)")
        }
    }

    @Test("every FS1 vector: decode, then one legal order from the ranks")
    func fs1() throws {
        #expect(!Self.golden.fs1.isEmpty)
        for vector in Self.golden.fs1 {
            let decoded = try #require(FS1Talents(code: vector.code))
            #expect(decoded.classSlug == vector.classSlug)
            #expect(decoded.raceSlug == vector.raceSlug)
            #expect(decoded.treeRanks == vector.treeRanks)
            guard let index = Self.indexes[vector.classSlug] else { continue }
            let result = PlannerRules.orderFromRanks(index, treeRanks: decoded.treeRanks)
            #expect(result.order == vector.order, "\(vector.code)")
            #expect(result.dropped == vector.dropped, "\(vector.code)")
        }
    }

    @Test("a refusal of each kind is covered")
    func coverage() {
        let reasons = (Self.golden.add + Self.golden.remove).compactMap(\.reason)
        #expect(reasons.contains { $0.hasPrefix("Tier ") })
        #expect(reasons.contains { $0.contains(" first") && !$0.hasPrefix("Tier ") })
        #expect(reasons.contains(Self.golden.noneLeft))
        #expect(reasons.contains { $0.hasSuffix("has no points to remove") })
    }

    @Test("FS1 refuses what decodeFS1 refuses")
    func fs1Refusals() {
        #expect(FS1Talents(code: "FS2:1:shaman:orc:5/0/0:") == nil)
        #expect(FS1Talents(code: "FS1:1:shaman:orc:5/0") == nil)
        #expect(FS1Talents(code: "FS1:1:shaman:orc:5/0:") == nil)
        #expect(FS1Talents(code: "FS1:1:shaman:orc:5!/0/0:") == nil)
        #expect(FS1Talents(code: "FS1:1:shaman:orc:5/0/0:|bags=1")?.treeRanks == [[5], [0], [0]])
    }
}
