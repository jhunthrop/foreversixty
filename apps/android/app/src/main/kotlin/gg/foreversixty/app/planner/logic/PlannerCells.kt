package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.model.Talent
import gg.foreversixty.app.model.TalentTree
import gg.foreversixty.app.talents.TalentCellModel
import gg.foreversixty.app.talents.TalentCellState
import gg.foreversixty.app.talents.TalentConnector
import gg.foreversixty.app.talents.TalentRules
import gg.foreversixty.app.talents.TalentTabModel

/** The stepper strip's talent (spec F13.1): icon, name, `{rank}/{max}`, the client's rank text. */
data class StripModel(
    val talent: Talent,
    val rank: Int,
    val state: TalentCellState,
    /** The rank held's text, or the first rank's before any point. */
    val text: String?,
    val refusal: Refusal?,
)

/** F6's drawing fed with the planner's ranks and its gold `available` state. iOS: PlannerCells.swift. */
object PlannerCells {
    fun cells(tree: TalentTree, ranks: Map<Int, Int>): List<TalentCellModel> = tree.talents.map { talent ->
        TalentCellModel(talent, ranks[talent.id] ?: 0, state(talent, tree, ranks), delta = null, guideRank = null)
    }

    /** The planner's `cellState`: F6's, with an open cell drawn as available (gold). */
    fun state(talent: Talent, tree: TalentTree, ranks: Map<Int, Int>): TalentCellState =
        TalentRules.cellState(talent, tree, ranks).let { if (it == TalentCellState.Open) TalentCellState.Available else it }

    fun connectors(tree: TalentTree, ranks: Map<Int, Int>): List<TalentConnector> = TalentRules.connectors(tree, ranks)

    /** `{Tree} {points}` for the three tabs, the totals summed from the drawn build. */
    fun tabs(index: TalentIndex, order: List<Int>): List<TalentTabModel> {
        val points = PlannerRules.pointsPerTree(index, order)
        return index.trees.mapIndexed { i, tree -> TalentTabModel(i, tree.name, points[i], differing = 0) }
    }

    /** The talents in reading order: lowest tier first, left to right. */
    fun readingOrder(tree: TalentTree): List<Talent> = tree.talents.sortedWith(compareBy({ it.tier }, { it.column }))

    /** The strip: the touched talent, else the shown tree's first. */
    fun strip(index: TalentIndex, state: EditorState, shownTree: Int): StripModel? {
        val talent = state.touched?.let(index.byId::get)
            ?: index.trees.getOrNull(shownTree)?.let { readingOrder(it).firstOrNull() }
            ?: return null
        val tree = index.treeOf[talent.id] ?: return null
        val ranks = state.ranks
        val rank = ranks[talent.id] ?: 0
        val text = talent.ranks.getOrNull((rank - 1).coerceAtLeast(0))?.description
        return StripModel(talent, rank, state(talent, tree, ranks), text, state.refusal)
    }

    /** The tree a talent sits in, left to right, or null. */
    fun treeOf(index: TalentIndex, talentId: Int): Int? = index.trees.indexOfFirst { tree -> tree.talents.any { it.id == talentId } }.takeIf { it >= 0 }
}

/** A hardware arrow key over the tree (spec F13.1). */
enum class TreeDirection { Left, Right, Up, Down }

/** Moving the strip's talent with the arrows: left and right in reading order, up and down to the nearest talent a tier away. */
object PlannerKeys {
    fun move(tree: TalentTree, current: Int?, direction: TreeDirection): Int? {
        val ordered = PlannerCells.readingOrder(tree)
        val here = ordered.firstOrNull { it.id == current } ?: return ordered.firstOrNull()?.id
        return when (direction) {
            TreeDirection.Left -> ordered.getOrNull(ordered.indexOf(here) - 1) ?: here
            TreeDirection.Right -> ordered.getOrNull(ordered.indexOf(here) + 1) ?: here
            TreeDirection.Up -> nearest(ordered.filter { it.tier < here.tier }.maxOfOrNull { it.tier }, ordered, here) ?: here
            TreeDirection.Down -> nearest(ordered.filter { it.tier > here.tier }.minOfOrNull { it.tier }, ordered, here) ?: here
        }.id
    }

    private fun nearest(tier: Int?, ordered: List<Talent>, here: Talent): Talent? =
        tier?.let { t -> ordered.filter { it.tier == t }.minWithOrNull(compareBy({ kotlin.math.abs(it.column - here.column) }, { it.column })) }
}
