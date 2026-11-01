package gg.foreversixty.app.common

import android.os.Build
import android.view.HapticFeedbackConstants
import android.view.View
import androidx.compose.runtime.Composable
import androidx.compose.runtime.compositionLocalOf
import androidx.compose.runtime.remember
import androidx.compose.ui.platform.LocalView

/**
 * The app's one haptics facade (v1 §7, W3 §5.7 Haptics). Every screen calls these moments and
 * nothing else; tests inject a recording fake. The system decides whether a moment is felt (the
 * user's touch-feedback setting), never the app.
 */
interface Haptics {
    /** A segment, chip, picker, tab section or character pick. */
    fun selection()

    /** A long press that removes (the planner's minus). */
    fun impactMedium()

    /** Saved: a sign-up, a share link ready, a watch. */
    fun success()

    /** Refused: a sign-up the server refused, a write that failed. */
    fun error()
}

/**
 * The system haptics on [view]: SEGMENT_TICK on API 34+ else CLOCK_TICK; CONFIRM and REJECT on API
 * 30+ (no fallback below: a wrong buzz is worse than none); LONG_PRESS everywhere.
 */
class ViewHaptics(private val view: View) : Haptics {
    override fun selection() {
        val tick = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.UPSIDE_DOWN_CAKE) {
            HapticFeedbackConstants.SEGMENT_TICK
        } else {
            HapticFeedbackConstants.CLOCK_TICK
        }
        view.performHapticFeedback(tick)
    }

    override fun impactMedium() {
        view.performHapticFeedback(HapticFeedbackConstants.LONG_PRESS)
    }

    override fun success() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) view.performHapticFeedback(HapticFeedbackConstants.CONFIRM)
    }

    override fun error() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) view.performHapticFeedback(HapticFeedbackConstants.REJECT)
    }
}

/** No haptics (previews and logic tests that do not look at them). */
object NoHaptics : Haptics {
    override fun selection() = Unit
    override fun impactMedium() = Unit
    override fun success() = Unit
    override fun error() = Unit
}

/** A test or preview may provide its own facade; null means the system's on the current view. */
val LocalHaptics = compositionLocalOf<Haptics?> { null }

/** The facade a composable uses: the provided one, else [ViewHaptics] on the current view. */
@Composable
fun rememberHaptics(): Haptics {
    val provided = LocalHaptics.current
    val view = LocalView.current
    return provided ?: remember(view) { ViewHaptics(view) }
}
