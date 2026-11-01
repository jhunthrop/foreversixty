import Foundation
import Models

/// A class's talents indexed for the planner's rules (`rules.ts` `indexTalents`): the trees left
/// to right, each talent by id and the tree it sits in.
struct PlannerIndex: Equatable, Sendable {
    let file: TalentsFile
    let trees: [TalentTree]
    let byID: [Int: Talent]
    /// Talent id to its tree's index in `trees`.
    let treeOf: [Int: Int]

    init(_ file: TalentsFile) {
        self.file = file
        trees = TalentRules.trees(file)
        var byID: [Int: Talent] = [:]
        var treeOf: [Int: Int] = [:]
        for (index, tree) in trees.enumerated() {
            for talent in tree.talents {
                byID[talent.id] = talent
                treeOf[talent.id] = index
            }
        }
        self.byID = byID
        self.treeOf = treeOf
    }

    /// Every talent, tree by tree in list order: the strip's first talent and the keyboard's walk.
    var allTalents: [Talent] { trees.flatMap(\.talents) }
}

/// One refused index of a point order, keyed the way the API's 400 keys it (`point_order[7]`).
struct PlannerFieldError: Equatable, Sendable {
    let field: String
    let message: String
}

/// The answer to a tap or a long press: allowed, or refused with the sentence the strip shows.
enum PlannerDecision: Equatable, Sendable {
    case allowed
    case refused(String)

    var reason: String? {
        if case let .refused(reason) = self { return reason }
        return nil
    }
}

/// The site planner's rules (web/src/lib/planner `rules.ts`, `derive.ts`, `fs1.ts`) ported
/// natively. Every function cites its web twin; `apps/shared/fixtures/planner-rules.json` holds
/// the golden vectors the tests replay, so the app and the site refuse the same moves with the
/// same words. The words come from the `plan.refuse.*` copy keys.
enum PlannerRules {
    /// The level of the first talent point (`types.ts` `FIRST_POINT_LEVEL`).
    static let firstPointLevel = TalentRules.baseLevel + 1
    /// The highest level the Level sheet offers.
    static let maxLevel = 60

    // MARK: Points

    /// `derive.ts` `pointsAvailableAt`.
    static func pointsAvailable(atLevel level: Int) -> Int { TalentRules.pointsAvailable(atLevel: level) }

    /// The lowest level a build of `spent` points fits (the Level sheet's floor, spec F13.1).
    static func minimumLevel(spent: Int) -> Int { max(firstPointLevel, spent + TalentRules.baseLevel) }

    /// `derive.ts` `ranksByTalent`.
    static func ranks(_ order: [Int]) -> [Int: Int] {
        order.reduce(into: [:]) { ranks, id in ranks[id, default: 0] += 1 }
    }

    /// `derive.ts` `pointsPerTree`: points in each tree, left to right; unknown ids count nowhere.
    static func pointsPerTree(_ index: PlannerIndex, order: [Int]) -> [Int] {
        var counts = Array(repeating: 0, count: index.trees.count)
        for id in order {
            if let tree = index.treeOf[id] { counts[tree] += 1 }
        }
        return counts
    }

    // MARK: Validation

    /// `rules.ts` `validateOrder`: rules 2 to 5 point by point.
    static func validate(_ index: PlannerIndex, order: [Int]) -> [PlannerFieldError] {
        violations(index, order: order).map { PlannerFieldError(field: $0.field, message: $0.message) }
    }

    /// One broken rule at one index; `rule` names the rule itself (`tier <treeId> <need>` for a
    /// tier lock, the sentence for the others), so a removal compares rules, not sentences.
    private struct Violation {
        let field: String
        let message: String
        let rule: String
    }

    /// `rules.ts` `violations`. Both counters advance even on a refused point; only an unknown
    /// talent skips them, exactly as the web and the API replay it.
    private static func violations(_ index: PlannerIndex, order: [Int]) -> [Violation] {
        var errors: [Violation] = []
        var pointsInTree: [Int: Int] = [:]
        var ranks: [Int: Int] = [:]
        for (position, id) in order.enumerated() {
            let field = "point_order[\(position)]"
            guard let talent = index.byID[id], let treeIndex = index.treeOf[id] else {
                let message = Copy.planRefuseUnknownTalent(id: String(id))
                errors.append(Violation(field: field, message: message, rule: message))
                continue
            }
            let tree = index.trees[treeIndex]
            let inTree = pointsInTree[treeIndex] ?? 0
            let rank = ranks[id] ?? 0
            if let (message, rule) = violation(talent, tree: tree, inTree: inTree, rank: rank, ranks: ranks, position: position, index: index) {
                errors.append(Violation(field: field, message: message, rule: rule))
            }
            pointsInTree[treeIndex] = inTree + 1
            ranks[id] = rank + 1
        }
        return errors
    }

    /// The first rule one point breaks: tier, prerequisite, max rank, then the cap.
    private static func violation(
        _ talent: Talent, tree: TalentTree, inTree: Int, rank: Int, ranks: [Int: Int], position: Int, index: PlannerIndex
    ) -> (message: String, rule: String)? {
        let need = TalentRules.pointsPerTier * talent.tier
        if inTree < need {
            return (Copy.planRefuseTierLocked(n: need, tree: tree.name, m: inTree), "tier \(tree.id) \(need)")
        }
        let message: String
        if let prereqID = talent.prereqTalentId, (ranks[prereqID] ?? 0) < (talent.prereqRank ?? 0) {
            message = prereqMissing(talent.name, rank: talent.prereqRank ?? 0, prereq: index.byID[prereqID]?.name ?? String(prereqID))
        } else if rank >= talent.maxRank {
            message = Copy.planRefuseMaxRank(talent: talent.name, max: talent.maxRank)
        } else if position >= TalentRules.maxPoints {
            message = Copy.planRefuseCap(n: TalentRules.maxPoints)
        } else {
            return nil
        }
        return (message, message)
    }

    /// `messages.prereqMissing`: `point` for one, `points` otherwise.
    static func prereqMissing(_ talent: String, rank: Int, prereq: String) -> String {
        rank == 1 ? Copy.planRefusePrereqOne(talent: talent, prereq: prereq) : Copy.planRefusePrereq(talent: talent, n: rank, prereq: prereq)
    }

    // MARK: Adding and removing

    /// `rules.ts` `withoutLastPoint`.
    static func withoutLastPoint(_ order: [Int], _ id: Int) -> [Int] {
        guard let last = order.lastIndex(of: id) else { return order }
        var next = order
        next.remove(at: last)
        return next
    }

    /// `rules.ts` `canAddPoint`: the candidate order's last index must validate.
    static func canAddPoint(_ index: PlannerIndex, order: [Int], talent id: Int) -> PlannerDecision {
        let candidate = order + [id]
        let field = "point_order[\(candidate.count - 1)]"
        guard let failure = validate(index, order: candidate).first(where: { $0.field == field }) else { return .allowed }
        return .refused(failure.message)
    }

    /// The app's add (spec F13.2): the level's budget first, then the site's own rules.
    static func add(_ index: PlannerIndex, order: [Int], talent id: Int, level: Int) -> PlannerDecision {
        if order.count >= pointsAvailable(atLevel: level) { return .refused(Copy.planRefuseNoneLeft) }
        return canAddPoint(index, order: order, talent: id)
    }

    /// `rules.ts` `canRemovePoint`: refused only for a violation the removal itself causes.
    static func canRemovePoint(_ index: PlannerIndex, order: [Int], talent id: Int) -> PlannerDecision {
        let talent = index.byID[id]
        guard order.contains(id) else {
            return .refused(talent.map { Copy.planRefuseNoPoints(talent: $0.name) } ?? Copy.planRefuseUnknownTalent(id: String(id)))
        }
        let newRank = order.count(where: { $0 == id }) - 1
        if let talent, let dependant = dependantNeedingMoreRank(index, order: order, talent: id, newRank: newRank) {
            return .refused(prereqMissing(dependant.name, rank: dependant.prereqRank ?? 0, prereq: talent.name))
        }
        let removedAt = order.lastIndex(of: id) ?? 0
        let before = Dictionary(
            violations(index, order: order).map { (fieldIndex($0.field), $0.rule) }, uniquingKeysWith: { first, _ in first }
        )
        let after = violations(index, order: withoutLastPoint(order, id))
        let new = after.first { error in
            let position = fieldIndex(error.field)
            return before[position < removedAt ? position : position + 1] != error.rule
        }
        return new.map { .refused($0.message) } ?? .allowed
    }

    /// `rules.ts` `dependantNeedingMoreRank`: a talent of the same tree holding points that needs
    /// more rank in `id` than the removal would leave.
    private static func dependantNeedingMoreRank(_ index: PlannerIndex, order: [Int], talent id: Int, newRank: Int) -> Talent? {
        guard let tree = index.treeOf[id] else { return nil }
        return index.trees[tree].talents.first { talent in
            talent.prereqTalentId == id && order.contains(talent.id) && (talent.prereqRank ?? 0) > newRank
        }
    }

    private static func fieldIndex(_ field: String) -> Int {
        Int(field.dropFirst("point_order[".count).dropLast()) ?? -1
    }

    // MARK: Cells

    /// The planner's `cellState`: gold `available` where a point can go now. With no points left
    /// a reachable empty cell reads open (F6's plain state) rather than locked.
    static func cellState(_ talent: Talent, index: PlannerIndex, order: [Int], ranks: [Int: Int], left: Int) -> TalentCellState {
        let rank = ranks[talent.id] ?? 0
        if rank >= talent.maxRank { return .maxed }
        if rank > 0 { return .filled }
        guard canAddPoint(index, order: order, talent: talent.id) == .allowed else { return .locked }
        return left > 0 ? .available : .open
    }

    // MARK: Building an order from ranks

    /// `fs1.ts` `orderFromRanks`: one legal order reaching the ranks, lowest tier first, left to
    /// right, in repeated passes. `dropped` names (sorted) every talent it could not fill.
    static func orderFromRanks(_ index: PlannerIndex, treeRanks: [[Int]]) -> (order: [Int], dropped: [Int]) {
        var wanted: [Int: Int] = [:]
        for (treeIndex, tree) in index.trees.enumerated() {
            let ranks = treeIndex < treeRanks.count ? treeRanks[treeIndex] : []
            for (position, talent) in tree.talents.enumerated() {
                let rank = min(position < ranks.count ? ranks[position] : 0, talent.maxRank)
                if rank > 0 { wanted[talent.id] = rank }
            }
        }
        var spent: [Int: Int] = [:]
        var order: [Int] = []
        var progress = true
        while progress {
            progress = false
            for tree in index.trees {
                let byTier = tree.talents.sorted { ($0.tier, $0.column) < ($1.tier, $1.column) }
                for talent in byTier {
                    let target = wanted[talent.id] ?? 0
                    while (spent[talent.id] ?? 0) < target, canAddPoint(index, order: order, talent: talent.id) == .allowed {
                        order.append(talent.id)
                        spent[talent.id, default: 0] += 1
                        progress = true
                    }
                }
            }
        }
        let dropped = wanted.filter { (spent[$0.key] ?? 0) < $0.value }.map(\.key).sorted()
        return (order, dropped)
    }

    /// Ranks by talent id as tree rank lists in the index's layout (for `orderFromRanks`).
    static func treeRanks(_ index: PlannerIndex, ranks: [Int: Int]) -> [[Int]] {
        index.trees.map { tree in tree.talents.map { ranks[$0.id] ?? 0 } }
    }
}

/// The talent half of an FS1 code (`fs1.ts` `decodeFS1`): class, race and the three trees' ranks.
/// Gear and the version 2 sections are the site's; the planner only reads the talents.
struct FS1Talents: Equatable, Sendable {
    static let prefix = "FS1"
    private static let trees = 3
    private static let fieldCount = 6
    private static let maxLength = 16384
    private static let radix = 36

    let dataBuild: String
    let classSlug: String
    let raceSlug: String
    let treeRanks: [[Int]]

    /// Nil for anything `decodeFS1` refuses: too long, not FS1, too few fields, not three trees,
    /// an unreadable digit.
    init?(code: String) {
        guard code.count <= Self.maxLength else { return nil }
        let head = code.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "|", omittingEmptySubsequences: false).first ?? ""
        let parts = head.split(separator: ":", omittingEmptySubsequences: false).map(String.init)
        guard parts.first == Self.prefix, parts.count >= Self.fieldCount else { return nil }
        let trees = parts[4].split(separator: "/", omittingEmptySubsequences: false)
        guard trees.count == Self.trees else { return nil }
        var ranks: [[Int]] = []
        for tree in trees {
            let digits = tree.map { Int(String($0), radix: Self.radix) }
            guard !digits.contains(nil) else { return nil }
            ranks.append(digits.compactMap(\.self))
        }
        dataBuild = parts[1]
        classSlug = parts[2]
        raceSlug = parts[3]
        treeRanks = ranks
    }
}
