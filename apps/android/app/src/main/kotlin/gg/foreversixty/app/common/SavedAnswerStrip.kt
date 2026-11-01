package gg.foreversixty.app.common

import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.Shape
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.StaleStrip

/**
 * The strip over a saved answer (W3 §5.7): offline (v1 words, unchanged) or, when [stale], `Forever
 * Sixty did not answer. Showing what you had at {time}.` with `Retry`, which re-runs the load.
 * [stamp] is the saved answer's time in the strip's words ([OfflineStamp]).
 */
@Composable
fun SavedAnswerStrip(stamp: String, stale: Boolean, onRetry: () -> Unit, shape: Shape? = null) {
    if (stale) {
        StaleStrip(stringResource(R.string.offline_stale, stamp), stringResource(R.string.offline_retry), onRetry, shape)
    } else {
        OfflineStrip(stringResource(R.string.offline_strip, stamp), shape)
    }
}
