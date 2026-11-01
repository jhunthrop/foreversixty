package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.common.Haptics

/**
 * The planner's haptics through the app's one facade (W3 §5.7 Haptics, planner rows): a point added
 * is a selection tick, a point removed (long press) impact medium / LONG_PRESS, a refusal error /
 * REJECT. An undo, a clear or a no-op plays none. A shared build plays success ([Haptics.success]).
 */
fun Haptics.play(outcome: EditOutcome) = when (outcome) {
    EditOutcome.Added -> selection()
    EditOutcome.Removed -> impactMedium()
    EditOutcome.Refused -> error()
    EditOutcome.Undone, EditOutcome.Cleared, EditOutcome.Unchanged -> Unit
}
