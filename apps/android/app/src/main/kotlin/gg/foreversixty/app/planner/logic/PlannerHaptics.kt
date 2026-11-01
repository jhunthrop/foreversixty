package gg.foreversixty.app.planner.logic

import android.os.Build
import android.view.HapticFeedbackConstants
import android.view.View

/** The planner's four haptics (spec F13.1, F13.3). iOS: the matching UIFeedbackGenerator calls. */
interface PlannerHaptics {
    /** A point added: `SEGMENT_TICK`. */
    fun added()

    /** A point removed: `LONG_PRESS`. */
    fun removed()

    /** A refusal: `REJECT`. */
    fun refused()

    /** A build shared: `CONFIRM`. */
    fun shared()

    /** The haptic one edit plays; an undo, a clear or a no-op plays none. */
    fun play(outcome: EditOutcome) = when (outcome) {
        EditOutcome.Added -> added()
        EditOutcome.Removed -> removed()
        EditOutcome.Refused -> refused()
        EditOutcome.Undone, EditOutcome.Cleared, EditOutcome.Unchanged -> Unit
    }
}

/** The system haptics on the editor's view; below API 30 confirm and reject play nothing. */
class ViewPlannerHaptics(private val view: View) : PlannerHaptics {
    override fun added() {
        val tick = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.UPSIDE_DOWN_CAKE) {
            HapticFeedbackConstants.SEGMENT_TICK
        } else {
            HapticFeedbackConstants.CLOCK_TICK
        }
        view.performHapticFeedback(tick)
    }

    override fun removed() {
        view.performHapticFeedback(HapticFeedbackConstants.LONG_PRESS)
    }

    override fun refused() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) view.performHapticFeedback(HapticFeedbackConstants.REJECT)
    }

    override fun shared() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) view.performHapticFeedback(HapticFeedbackConstants.CONFIRM)
    }
}

object NoPlannerHaptics : PlannerHaptics {
    override fun added() = Unit
    override fun removed() = Unit
    override fun refused() = Unit
    override fun shared() = Unit
}
