import Foundation
import Models

/// The planner's read-only rules (web/src/lib/planner), applied to a talent file and a set of
/// ranks. Nothing here is new: each function cites its web twin, so the app and the site draw the
/// same cells, lines and counts from the same strings.
enum TalentRules {
    /// A build spends at most this many points (`types.ts` `MAX_POINTS`).
    static let maxPoints = 51
    /// Tier `t` opens at `5 * t` points in its tree (`POINTS_PER_TIER`).
    static let pointsPerTier = 5
    /// The level of a character with no points spent (`BASE_LEVEL`).
    static let baseLevel = 9
    /// The separator between the three trees of a band's `talents` string.
    private static let treeSeparator: Character = "-"
    private static let radix = 36

    /// One base-36 digit per talent, in the file's list order (`band-compare.ts`
    /// `bandTreeRanksFrom`, `account-character.ts` `ranksOfTree`); a bad digit reads as 0.
    static func ranks(ofTree digits: some StringProtocol) -> [Int] {
        digits.map { Int(String($0), radix: radix) ?? 0 }
    }

    /// A band's `talents` (`"342…-0515…-000…"`) as one rank list per tree.
    static func ranks(ofBand talents: String) -> [[Int]] {
        talents.split(separator: treeSeparator, omittingEmptySubsequences: false).map { ranks(ofTree: $0) }
    }

    /// The trees left to right (`rules.ts` `indexTalents`).
    static func trees(_ file: TalentsFile) -> [TalentTree] {
        file.trees.sorted { $0.position < $1.position }
    }

    /// Talent id to rank, each clamped to the talent's maximum; a missing digit is 0.
    static func rankByID(_ file: TalentsFile, treeRanks: [[Int]]) -> [Int: Int] {
        var result: [Int: Int] = [:]
        for (treeIndex, tree) in trees(file).enumerated() {
            let digits = treeIndex < treeRanks.count ? treeRanks[treeIndex] : []
            for (index, talent) in tree.talents.enumerated() {
                let rank = index < digits.count ? digits[index] : 0
                result[talent.id] = min(max(rank, 0), talent.maxRank)
            }
        }
        return result
    }

    /// Points in each tree, left to right (`derive.ts` `pointsPerTree`).
    static func pointsPerTree(_ file: TalentsFile, ranks: [Int: Int]) -> [Int] {
        trees(file).map { tree in tree.talents.reduce(0) { $0 + (ranks[$1.id] ?? 0) } }
    }

    /// The points a character of `level` has earned (`derive.ts` `pointsAvailableAt`).
    static func pointsAvailable(atLevel level: Int) -> Int {
        min(maxPoints, max(0, level - baseLevel))
    }

    /// The planner's `cellState` without its gold `available` (spec F6 §4.5): an empty cell is
    /// open when its tier is reached and its prerequisite held, else locked.
    static func cellState(_ talent: Talent, in tree: TalentTree, ranks: [Int: Int]) -> TalentCellState {
        let rank = ranks[talent.id] ?? 0
        if rank >= talent.maxRank { return .maxed }
        if rank > 0 { return .filled }
        return canTakePoint(talent, in: tree, ranks: ranks) ? .open : .locked
    }

    /// Whether an empty talent's tier is open and its prerequisite met (`rules.ts` rules 3 and 4).
    static func canTakePoint(_ talent: Talent, in tree: TalentTree, ranks: [Int: Int]) -> Bool {
        let inTree = tree.talents.reduce(0) { $0 + (ranks[$1.id] ?? 0) }
        return inTree >= pointsPerTier * talent.tier && prerequisiteMet(talent, ranks: ranks)
    }

    /// True once the prerequisite holds the rank the talent needs (`connectors.ts` `met`).
    static func prerequisiteMet(_ talent: Talent, ranks: [Int: Int]) -> Bool {
        guard let prereq = talent.prereqTalentId else { return true }
        return (ranks[prereq] ?? 0) >= (talent.prereqRank ?? 0)
    }

    /// The tree's prerequisite lines (`connectors.ts` `connectorsFor`), sorted the same way.
    static func connectors(_ tree: TalentTree, ranks: [Int: Int]) -> [TalentConnector] {
        let byID = Dictionary(tree.talents.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
        return tree.talents.compactMap { talent -> TalentConnector? in
            guard let prereqID = talent.prereqTalentId, let prereq = byID[prereqID] else { return nil }
            return TalentConnector(
                from: TalentCell(tier: prereq.tier, column: prereq.column),
                to: TalentCell(tier: talent.tier, column: talent.column),
                fromID: prereq.id,
                toID: talent.id,
                met: prerequisiteMet(talent, ranks: ranks)
            )
        }
        .sorted { ($0.fromID, $0.toID) < ($1.fromID, $1.toID) }
    }

    /// The tree's grid size (`grid.ts` `gridSize`).
    static func gridSize(_ tree: TalentTree) -> (tiers: Int, columns: Int) {
        (tree.talents.map { $0.tier + 1 }.max() ?? 0, tree.talents.map { $0.column + 1 }.max() ?? 0)
    }

    /// The planner's `diffAgainstBand`. `stray = Σ max(yours - guide, 0)` is your points spent where
    /// the guide spends fewer; `missingCount = Σ max(guide - yours, 0)` the guide's points you have not
    /// taken; `extraCount = max(0, yourTotal - guideTotal)` your points past the guide's own total.
    /// Points past the guide's total are not differences (owner ruling 2026-10-31), so
    /// `diffCount = max(0, stray - extraCount)`: stray points differ only up to the guide's total.
    /// `differing` is every talent whose two ranks differ.
    static func compare(_ file: TalentsFile, yours: [Int: Int], guide: [Int: Int]) -> TalentDiff {
        var stray = 0
        var missingCount = 0
        var differing: Set<Int> = []
        for talent in trees(file).flatMap(\.talents) {
            let mine = yours[talent.id] ?? 0
            let theirs = guide[talent.id] ?? 0
            stray += max(mine - theirs, 0)
            missingCount += max(theirs - mine, 0)
            if mine != theirs { differing.insert(talent.id) }
        }
        // stray - missing = yourTotal - guideTotal, so the extra points are that difference floored at 0.
        let extraCount = max(0, stray - missingCount)
        return TalentDiff(
            diffCount: max(0, stray - extraCount), missingCount: missingCount, extraCount: extraCount, differing: differing
        )
    }

    /// The tab that opens: the tree with the most points in the build first shown, ties to the
    /// first (spec F6 §5.3).
    static func openingTree(_ points: [Int]) -> Int {
        points.indices.max { points[$0] < points[$1] || (points[$0] == points[$1] && $0 > $1) } ?? 0
    }
}

/// A cell's state: the planner's `CellState`. F6's read-only tree never draws `available` (it reads
/// as open there); the planner (F13) draws it gold where a point can go now.
enum TalentCellState: Equatable, Sendable {
    case locked, open, available, filled, maxed
}

/// A talent's place on the grid, 0-based.
struct TalentCell: Equatable, Sendable {
    let tier: Int
    let column: Int
}

/// The three line shapes the client draws (`connectors.ts` `connectorPath`).
enum TalentConnectorShape: Equatable, Sendable {
    case down, across, elbow
}

/// One prerequisite line: from the prerequisite to the talent that needs it.
struct TalentConnector: Equatable, Sendable {
    let from: TalentCell
    let to: TalentCell
    let fromID: Int
    let toID: Int
    let met: Bool

    var shape: TalentConnectorShape {
        if from.column == to.column { return .down }
        return from.tier == to.tier ? .across : .elbow
    }
}

/// Your build against the guide's: the counts and the talents to mark.
struct TalentDiff: Equatable, Sendable {
    let diffCount: Int
    let missingCount: Int
    /// Your points past the guide's own total: the "+n beyond the guide" note.
    let extraCount: Int
    let differing: Set<Int>

    /// The status line's state (spec F6 §4.4): `Matches` once you hold every guide point, however
    /// many points you have beyond its total.
    var status: TalentStatus {
        if missingCount == 0 { return .matches }
        return diffCount > 0 ? .differs(diffCount) : .toTake(missingCount)
    }
}

/// What the status line says about your build against the guide's.
enum TalentStatus: Equatable, Sendable {
    case matches
    /// `{n} of your {m} points differ`: n is the diffCount.
    case differs(Int)
    /// `{n} of the {band} build's points still to take`: your build is inside the guide's.
    case toTake(Int)
}
