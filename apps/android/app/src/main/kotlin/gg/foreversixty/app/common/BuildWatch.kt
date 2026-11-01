package gg.foreversixty.app.common

import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.compose.LifecycleResumeEffect
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.NewDataStrip
import kotlinx.coroutines.launch

/**
 * The new-data rule for a screen showing build data (W3 §5.7, extended by lane F15 from Guides to
 * BiS, Talents and the item page): remember the build the screen opened on; a newer active build
 * found later waits behind the strip, and nothing on screen moves until `Update` ([accept]).
 */
class BuildWatch(private val latestBuild: suspend () -> String?) {
    var shown: String? by mutableStateOf(null)
        private set
    var newer: String? by mutableStateOf(null)
        private set

    val hasNewData: Boolean get() = newer != null

    /** On open and on every resume: the first answer is what is shown; a different one later is news. */
    suspend fun check() {
        val latest = latestBuild() ?: return
        val current = shown
        when {
            current == null -> shown = latest
            latest != current -> newer = latest
            else -> newer = null
        }
    }

    /** `Update`: the screen now shows the newer build. */
    fun accept() {
        newer?.let { shown = it }
        newer = null
    }
}

/**
 * The fresh active build (a network read, never the saved copy), for [BuildWatch]; provided by the
 * app shell. Null (tests, previews) means no screen ever shows the strip.
 */
val LocalLatestBuild = staticCompositionLocalOf<(suspend () -> String?)?> { null }

/**
 * The strip for a screen that shows build data: it checks on open and resume, and [onUpdate] runs
 * when the player taps `Update` (the caller reloads on the new build).
 */
@Composable
fun NewBuildStrip(key: Any?, onUpdate: () -> Unit) {
    val source = LocalLatestBuild.current ?: return
    val watch = remember(key, source) { BuildWatch(source) }
    val scope = rememberCoroutineScope()
    LifecycleResumeEffect(watch) {
        val job = scope.launch { watch.check() }
        onPauseOrDispose { job.cancel() }
    }
    if (watch.hasNewData) {
        NewDataStrip(stringResource(R.string.offline_newData), stringResource(R.string.offline_update), {
            watch.accept()
            onUpdate()
        })
    }
}
