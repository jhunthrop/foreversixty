package gg.foreversixty.app.talents

import gg.foreversixty.app.model.Talent
import gg.foreversixty.app.model.TalentTree
import gg.foreversixty.app.model.TalentsFile

/**
 * The planner's read-only rules (web/src/lib/planner), applied to a talent file and a set of
 * ranks. Nothing here is new: each function names its web twin, so the app and the site draw the
 * same cells, lines and counts from the same strings. iOS: TalentRules.swift.
 */
object TalentRules {
    /** A build spends at most this many points (`types.ts` `MAX_POINTS`). */
    const val MAX_POINTS = 51
    /** Tier `t` opens at `5 * t` points in its tree (`POINTS_PER_TIER`). */
    const val POINTS_PER_TIER = 5
    /** The level of a character with no points spent (`BASE_LEVEL`). */
    const val BASE_LEVEL = 9
    private const val TREE_SEPARATOR = '-'
    private const val RADIX = 36

    /** One base-36 digit per talent, in the file's list order; a bad digit reads as 0. */
    fun ranksOfTree(digits: CharSequence): List<Int> = digits.map { it.digitToIntOrNull(RADIX) ?: 0 }

    /** A band's `talents` (`"342…-0515…-000…"`) as one rank list per tree. */
    fun ranksOfBand(talents: String): List<List<Int>> = talents.split(TREE_SEPARATOR).map(::ranksOfTree)

    /** The trees left to right (`rules.ts` `indexTalents`). */
    fun trees(file: TalentsFile): List<TalentTree> = file.trees.sortedBy { it.position }

    /** Talent id to rank, each clamped to the talent's maximum; a missing digit is 0. */
    fun rankById(file: TalentsFile, treeRanks: List<List<Int>>): Map<Int, Int> = buildMap {
        trees(file).forEachIndexed { treeIndex, tree ->
            val digits = treeRanks.getOrElse(treeIndex) { emptyList() }
            tree.talents.forEachIndexed { index, talent ->
                put(talent.id, digits.getOrElse(index) { 0 }.coerceIn(0, talent.maxRank))
            }
        }
    }

    /** Points in each tree, left to right (`derive.ts` `pointsPerTree`). */
    fun pointsPerTree(file: TalentsFile, ranks: Map<Int, Int>): List<Int> =
        trees(file).map { tree -> tree.talents.sumOf { ranks[it.id] ?: 0 } }

    /** The points a character of [level] has earned (`derive.ts` `pointsAvailableAt`). */
    fun pointsAvailable(level: Int): Int = (level - BASE_LEVEL).coerceIn(0, MAX_POINTS)

    /**
     * The planner's `cellState` without its gold `available` (spec F6 §4.5): an empty cell is
     * open when its tier is reached and its prerequisite held, else locked.
     */
    fun cellState(talent: Talent, tree: TalentTree, ranks: Map<Int, Int>): TalentCellState {
        val rank = ranks[talent.id] ?: 0
        return when {
            rank >= talent.maxRank -> TalentCellState.Maxed
            rank > 0 -> TalentCellState.Filled
            canTakePoint(talent, tree, ranks) -> TalentCellState.Open
            else -> TalentCellState.Locked
        }
    }

    /** Whether an empty talent's tier is open and its prerequisite met (`rules.ts` rules 3 and 4). */
    fun canTakePoint(talent: Talent, tree: TalentTree, ranks: Map<Int, Int>): Boolean =
        tree.talents.sumOf { ranks[it.id] ?: 0 } >= POINTS_PER_TIER * talent.tier && prerequisiteMet(talent, ranks)

    /** True once the prerequisite holds the rank the talent needs (`connectors.ts` `met`). */
    fun prerequisiteMet(talent: Talent, ranks: Map<Int, Int>): Boolean {
        val prereq = talent.prereqTalentId ?: return true
        return (ranks[prereq] ?: 0) >= (talent.prereqRank ?: 0)
    }

    /** The tree's prerequisite lines (`connectors.ts` `connectorsFor`), sorted the same way. */
    fun connectors(tree: TalentTree, ranks: Map<Int, Int>): List<TalentConnector> {
        val byId = tree.talents.associateBy { it.id }
        return tree.talents.mapNotNull { talent ->
            val prereq = talent.prereqTalentId?.let(byId::get) ?: return@mapNotNull null
            TalentConnector(
                from = TalentCell(prereq.tier, prereq.column),
                to = TalentCell(talent.tier, talent.column),
                fromId = prereq.id,
                toId = talent.id,
                met = prerequisiteMet(talent, ranks),
            )
        }.sortedWith(compareBy({ it.fromId }, { it.toId }))
    }

    /** The tree's grid size (`grid.ts` `gridSize`). */
    fun gridSize(tree: TalentTree): TalentCell =
        TalentCell(tree.talents.maxOfOrNull { it.tier + 1 } ?: 0, tree.talents.maxOfOrNull { it.column + 1 } ?: 0)

    /**
     * The planner's `diffAgainstBand`: `diffCount` is your points spent where the guide spends
     * fewer (`Σ max(yours - guide, 0)`), `missingCount` the guide's points you have not taken
     * (`Σ max(guide - yours, 0)`), and `differing` every talent whose two ranks differ.
     */
    fun compare(file: TalentsFile, yours: Map<Int, Int>, guide: Map<Int, Int>): TalentDiff {
        var diffCount = 0
        var missingCount = 0
        val differing = mutableSetOf<Int>()
        trees(file).flatMap { it.talents }.forEach { talent ->
            val mine = yours[talent.id] ?: 0
            val theirs = guide[talent.id] ?: 0
            diffCount += (mine - theirs).coerceAtLeast(0)
            missingCount += (theirs - mine).coerceAtLeast(0)
            if (mine != theirs) differing += talent.id
        }
        return TalentDiff(diffCount, missingCount, differing)
    }

    /** The tab that opens: the tree with the most points in the build first shown, ties to the first. */
    fun openingTree(points: List<Int>): Int = points.indices.maxWithOrNull(compareBy({ points[it] }, { -it })) ?: 0
}

/** A cell's state: the planner's `CellState` with `available` read as open (spec F6 §5.2). */
enum class TalentCellState { Locked, Open, Filled, Maxed }

/** A talent's place on the grid, 0-based; also a grid's size in tiers and columns. */
data class TalentCell(val tier: Int, val column: Int)

/** The three line shapes the client draws (`connectors.ts` `connectorPath`). */
enum class TalentConnectorShape { Down, Across, Elbow }

/** One prerequisite line: from the prerequisite to the talent that needs it. */
data class TalentConnector(
    val from: TalentCell,
    val to: TalentCell,
    val fromId: Int,
    val toId: Int,
    val met: Boolean,
) {
    val shape: TalentConnectorShape
        get() = when {
            from.column == to.column -> TalentConnectorShape.Down
            from.tier == to.tier -> TalentConnectorShape.Across
            else -> TalentConnectorShape.Elbow
        }
}

/** Your build against the guide's: the counts and the talents to mark. */
data class TalentDiff(val diffCount: Int, val missingCount: Int, val differing: Set<Int>) {
    /** The status line's state (spec F6 §4.4): `Matches` exactly when no cell differs. */
    val status: TalentStatus
        get() = when {
            differing.isEmpty() -> TalentStatus.Matches
            diffCount > 0 -> TalentStatus.Differs(diffCount)
            else -> TalentStatus.ToTake(missingCount)
        }
}

/** What the status line says about your build against the guide's. */
sealed interface TalentStatus {
    data object Matches : TalentStatus
    /** `{n} of your {m} points differ`: n is the diffCount. */
    data class Differs(val count: Int) : TalentStatus
    /** `{n} of the {band} build's points still to take`: your build is inside the guide's. */
    data class ToTake(val count: Int) : TalentStatus
}
