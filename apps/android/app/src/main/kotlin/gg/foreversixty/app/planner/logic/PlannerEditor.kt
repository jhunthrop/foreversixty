package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.talents.TalentRules

/**
 * The editor's state (spec F13.1): the class, race and level, the point order, the undo stack of
 * earlier orders (newest last, at most [PlannerEditor.UNDO_LIMIT]), the strip's talent and its line.
 */
data class EditorState(
    val classSlug: String,
    val raceSlug: String?,
    val level: Int,
    val order: List<Int>,
    val undo: List<List<Int>> = emptyList(),
    /** The strip's talent: the last touched, or null before any touch (the strip shows the first). */
    val touched: Int? = null,
    /** Why the last touch changed nothing; null once a touch succeeds. */
    val refusal: Refusal? = null,
) {
    val spent: Int get() = order.size
    val left: Int get() = (PlannerRules.pointsAvailable(level) - spent).coerceAtLeast(0)
    val ranks: Map<Int, Int> get() = PlannerRules.ranks(order)
    val canUndo: Boolean get() = undo.isNotEmpty()

    /** What the phone keeps per class: never the undo stack or the strip. */
    val draft: PlannerDraft get() = PlannerDraft(classSlug, raceSlug, level, order)
}

/** What one edit did, so the screen plays the right haptic (spec F13.1 Interactions). */
enum class EditOutcome { Added, Removed, Refused, Undone, Cleared, Unchanged }

/** An edit's new state and what happened. */
data class Edit(val state: EditorState, val outcome: EditOutcome)

/** The editor's pure transitions: each returns a new state and never mutates. iOS: PlannerEditor.swift. */
object PlannerEditor {
    const val UNDO_LIMIT = 100
    const val MIN_LEVEL = TalentRules.BASE_LEVEL + 1
    const val MAX_LEVEL = TalentRules.BASE_LEVEL + TalentRules.MAX_POINTS

    /** Tap or `+`: one point, or the refusal and nothing else. */
    fun add(index: TalentIndex, state: EditorState, talentId: Int): Edit {
        val refusal = PlannerRules.refuseAdd(index, state.order, talentId, state.level)
        return if (refusal != null) refused(state, talentId, refusal) else changed(state, talentId, state.order + talentId, EditOutcome.Added)
    }

    /** Long press or `−`: the talent's last point, or the refusal and nothing else. */
    fun remove(index: TalentIndex, state: EditorState, talentId: Int): Edit {
        val refusal = PlannerRules.refuseRemove(index, state.order, talentId)
        return if (refusal != null) {
            refused(state, talentId, refusal)
        } else {
            changed(state, talentId, PlannerRules.withoutLastPoint(state.order, talentId), EditOutcome.Removed)
        }
    }

    /** One step back; each point is one step. */
    fun undo(state: EditorState): Edit {
        val previous = state.undo.lastOrNull() ?: return Edit(state, EditOutcome.Unchanged)
        return Edit(state.copy(order = previous, undo = state.undo.dropLast(1), refusal = null), EditOutcome.Undone)
    }

    /** Reset, after its alert: every point off, as one undoable step. */
    fun reset(state: EditorState): Edit =
        if (state.order.isEmpty()) Edit(state, EditOutcome.Unchanged)
        else Edit(state.copy(order = emptyList(), undo = pushed(state), refusal = null), EditOutcome.Cleared)

    /** The level sheet's stepper: 10 to 60, never below the spent points' own level (spent + 9). */
    fun setLevel(state: EditorState, level: Int): EditorState = state.copy(level = level.coerceIn(minLevel(state), MAX_LEVEL))

    fun minLevel(state: EditorState): Int = maxOf(MIN_LEVEL, state.spent + TalentRules.BASE_LEVEL)

    fun setRace(state: EditorState, raceSlug: String): EditorState = state.copy(raceSlug = raceSlug)

    /** A loaded order (a guide's code, a shared build's copy) replaces the points; the old ones stay one undo away. */
    fun load(state: EditorState, order: List<Int>, raceSlug: String?): EditorState = state.copy(
        order = order,
        raceSlug = raceSlug ?: state.raceSlug,
        level = maxOf(state.level, order.size + TalentRules.BASE_LEVEL).coerceAtMost(MAX_LEVEL),
        undo = if (state.order.isEmpty()) state.undo else pushed(state),
        touched = null,
        refusal = null,
    )

    private fun changed(state: EditorState, talentId: Int, order: List<Int>, outcome: EditOutcome) =
        Edit(state.copy(order = order, undo = pushed(state), touched = talentId, refusal = null), outcome)

    private fun refused(state: EditorState, talentId: Int, refusal: Refusal) =
        Edit(state.copy(touched = talentId, refusal = refusal), EditOutcome.Refused)

    private fun pushed(state: EditorState): List<List<Int>> = (state.undo + listOf(state.order)).takeLast(UNDO_LIMIT)
}
