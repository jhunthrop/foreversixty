package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.model.Talent
import gg.foreversixty.app.model.TalentTree
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.talents.TalentRules

/**
 * Why the planner refuses a point (`rules.ts` `messages`, plus the app's level budget). Pure data:
 * [PlannerCopy] turns each into the planner's own sentence, so the rules stay JVM-testable and the
 * words come from the shared catalog.
 */
sealed interface Refusal {
    /** `Tier {tier} of {tree} needs {n} points in {tree} first`. */
    data class TierLocked(val tier: Int, val treeName: String) : Refusal {
        val points: Int get() = TalentRules.POINTS_PER_TIER * tier
    }

    /** `{talent} needs {n} points in {prereq} first`. */
    data class PrereqMissing(val talentName: String, val rank: Int, val prereqName: String) : Refusal

    /** `{talent} is already at {max} of {max} points`. */
    data class MaxRank(val talentName: String, val maxRank: Int) : Refusal

    /** `A build spends at most {n} points`. */
    data object Cap : Refusal

    /** `Talent {id} is not in this class`. */
    data class UnknownTalent(val id: Int) : Refusal

    /** `{talent} has no points to remove`. */
    data class NoPoints(val talentName: String) : Refusal

    /** The app's budget (spec F13.2): the level's points are all spent. */
    data object NoneLeft : Refusal
}

/** One refused index of a point order: the API's `point_order[n]` field. */
data class OrderError(val index: Int, val refusal: Refusal) {
    val field: String get() = "point_order[$index]"
}

/** The talent file indexed the way `rules.ts` `indexTalents` does: trees left to right. */
class TalentIndex(val file: TalentsFile) {
    val trees: List<TalentTree> = TalentRules.trees(file)
    val byId: Map<Int, Talent> = trees.flatMap { it.talents }.associateBy { it.id }
    val treeOf: Map<Int, TalentTree> = trees.flatMap { tree -> tree.talents.map { it.id to tree } }.toMap()
}

/** A legal order rebuilt from final ranks, and the talents whose ranks could not be placed. */
data class OrderFromRanks(val order: List<Int>, val dropped: List<Int>)

/**
 * The web planner's point rules (`web/src/lib/planner/rules.ts`, `derive.ts`, `fs1.ts`), ported
 * function for function so the app refuses exactly what the site and the API refuse. Every
 * decision is replayed against `apps/shared/fixtures/planner-rules.json` in the unit tests.
 * Nothing here mutates: each change returns a new order. iOS: PlannerRules.swift.
 */
object PlannerRules {
    /** The points a character of [level] has earned (`derive.ts` `pointsAvailableAt`). */
    fun pointsAvailable(level: Int): Int = TalentRules.pointsAvailable(level)

    /** Rules 2 to 5, point by point (`validateOrder`); the counters advance even on a refused point. */
    fun validate(index: TalentIndex, order: List<Int>): List<OrderError> {
        val errors = mutableListOf<OrderError>()
        val pointsInTree = mutableMapOf<Int, Int>()
        val ranks = mutableMapOf<Int, Int>()
        order.forEachIndexed { i, id ->
            val talent = index.byId[id]
            val tree = index.treeOf[id]
            if (talent == null || tree == null) {
                errors += OrderError(i, Refusal.UnknownTalent(id))
                return@forEachIndexed
            }
            val inTree = pointsInTree[tree.id] ?: 0
            val rank = ranks[id] ?: 0
            refusalAt(index, talent, tree, inTree, rank, ranks, i)?.let { errors += OrderError(i, it) }
            pointsInTree[tree.id] = inTree + 1
            ranks[id] = rank + 1
        }
        return errors
    }

    private fun refusalAt(
        index: TalentIndex,
        talent: Talent,
        tree: TalentTree,
        inTree: Int,
        rank: Int,
        ranks: Map<Int, Int>,
        position: Int,
    ): Refusal? {
        val prereqId = talent.prereqTalentId
        return when {
            inTree < TalentRules.POINTS_PER_TIER * talent.tier -> Refusal.TierLocked(talent.tier, tree.name)
            prereqId != null && (ranks[prereqId] ?: 0) < (talent.prereqRank ?: 0) ->
                Refusal.PrereqMissing(talent.name, talent.prereqRank ?: 0, index.byId[prereqId]?.name ?: prereqId.toString())
            rank >= talent.maxRank -> Refusal.MaxRank(talent.name, talent.maxRank)
            position >= TalentRules.MAX_POINTS -> Refusal.Cap
            else -> null
        }
    }

    /** The app's add (spec F13.2): the level budget first, then `canAddPoint`. Null means allowed. */
    fun refuseAdd(index: TalentIndex, order: List<Int>, talentId: Int, level: Int): Refusal? {
        if (order.size >= pointsAvailable(level)) return Refusal.NoneLeft
        return refuseAddIgnoringLevel(index, order, talentId)
    }

    /** `canAddPoint`: the candidate order's last point, validated. */
    fun refuseAddIgnoringLevel(index: TalentIndex, order: List<Int>, talentId: Int): Refusal? {
        val candidate = order + talentId
        return validate(index, candidate).firstOrNull { it.index == candidate.lastIndex }?.refusal
    }

    /**
     * `canRemovePoint`: a dependant that needs more rank than the removal leaves, else a violation the
     * removal itself causes (an older one carried over is not this removal's fault). Null means allowed.
     */
    fun refuseRemove(index: TalentIndex, order: List<Int>, talentId: Int): Refusal? {
        val talent = index.byId[talentId]
        if (talentId !in order) return talent?.let { Refusal.NoPoints(it.name) } ?: Refusal.UnknownTalent(talentId)
        if (talent != null) {
            val newRank = order.count { it == talentId } - 1
            val dependant = index.treeOf[talentId]?.talents?.firstOrNull {
                it.prereqTalentId == talentId && it.id in order && (it.prereqRank ?: 0) > newRank
            }
            if (dependant != null) return Refusal.PrereqMissing(dependant.name, dependant.prereqRank ?: 0, talent.name)
        }
        val removedAt = order.lastIndexOf(talentId)
        val before = validate(index, order).associate { it.index to it.refusal }
        val after = validate(index, withoutLastPoint(order, talentId))
        return after.firstOrNull { error ->
            val beforeIndex = if (error.index < removedAt) error.index else error.index + 1
            before[beforeIndex] != error.refusal
        }?.refusal
    }

    /** A new order without the last point in [talentId] (`withoutLastPoint`). */
    fun withoutLastPoint(order: List<Int>, talentId: Int): List<Int> {
        val last = order.lastIndexOf(talentId)
        return if (last == -1) order.toList() else order.filterIndexed { i, _ -> i != last }
    }

    /** Points per tree, left to right (`derive.ts` `pointsPerTree`); an unknown id counts nowhere. */
    fun pointsPerTree(index: TalentIndex, order: List<Int>): List<Int> {
        val counts = order.mapNotNull { index.treeOf[it]?.id }.groupingBy { it }.eachCount()
        return index.trees.map { counts[it.id] ?: 0 }
    }

    /** Talent id to rank (`derive.ts` `ranksByTalent`). */
    fun ranks(order: List<Int>): Map<Int, Int> = order.groupingBy { it }.eachCount()

    /**
     * `fs1.ts` `orderFromRanks`: a legal order reaching [treeRanks], spending lowest tier first, left
     * to right, in repeated passes over the trees, each point asked of [refuseAddIgnoringLevel].
     */
    fun orderFromRanks(index: TalentIndex, treeRanks: List<List<Int>>): OrderFromRanks {
        val wanted = wantedRanks(index, treeRanks)
        val spent = mutableMapOf<Int, Int>()
        var order = emptyList<Int>()
        var progress = true
        while (progress) {
            progress = false
            index.trees.forEach { tree ->
                tree.talents.sortedWith(compareBy({ it.tier }, { it.column })).forEach { talent ->
                    val target = wanted[talent.id] ?: 0
                    while ((spent[talent.id] ?: 0) < target && refuseAddIgnoringLevel(index, order, talent.id) == null) {
                        order = order + talent.id
                        spent[talent.id] = (spent[talent.id] ?: 0) + 1
                        progress = true
                    }
                }
            }
        }
        val dropped = wanted.filter { (id, target) -> (spent[id] ?: 0) < target }.keys.sorted()
        return OrderFromRanks(order, dropped)
    }

    private fun wantedRanks(index: TalentIndex, treeRanks: List<List<Int>>): Map<Int, Int> = buildMap {
        index.trees.forEachIndexed { treeIndex, tree ->
            val ranks = treeRanks.getOrElse(treeIndex) { emptyList() }
            tree.talents.forEachIndexed { talentIndex, talent ->
                val rank = minOf(ranks.getOrElse(talentIndex) { 0 }, talent.maxRank)
                if (rank > 0) put(talent.id, rank)
            }
        }
    }
}
