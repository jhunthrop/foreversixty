package gg.foreversixty.app.testing

import gg.foreversixty.app.designsystem.screenGround
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.test.junit4.ComposeContentTestRule
import androidx.compose.ui.test.onRoot
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.LocalParityRecorder
import gg.foreversixty.app.designsystem.ParityRecorder
import org.junit.Assert.assertEquals
import org.robolectric.RuntimeEnvironment

/** The device every parity state renders on (design/parity.md "Native"): Pixel 9 at 412 x 915 dp. */
const val PARITY_PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"

/** Pixel 9's width at a taller height, for a screen the mock draws longer than the window. */
const val PARITY_PIXEL_9_TALL = "w412dp-h1700dp-normal-long-notround-any-420dpi-keyshidden-nonav"

/**
 * One mocked state held to its mock (design/parity.md "Native"): the regions and inventory against
 * design/mocks/app/regions/android-<screen>.json, then the image golden and the text golden. A known
 * drift is named in [expectedDrift] (each line exactly as the comparison prints it) with a date and a
 * lane, and the check fails once that drift is gone. [fontScale] above 1 records the golden only.
 */
object ParityCheck {
    fun state(
        rule: ComposeContentTestRule,
        screen: String,
        fontScale: Float = 1f,
        expectedDrift: Set<String> = emptySet(),
        content: @Composable () -> Unit,
    ) {
        if (fontScale != 1f) RuntimeEnvironment.setFontScale(fontScale)
        val recorder = ParityRecorder()
        var density = 1f
        rule.setContent {
            density = LocalDensity.current.density
            ForeverTheme {
                CompositionLocalProvider(LocalParityRecorder provides recorder) {
                    Box(Modifier.fillMaxSize().screenGround()) { content() }
                }
            }
        }
        rule.waitForIdle()
        val suffix = if (fontScale == 1f) "" else "-x${fontScale.toString().replace(".", "_")}"
        val path = "parity/android-$screen$suffix"
        // The image first, so a failing comparison still leaves the capture to read beside the mock.
        rule.onRoot().captureRoboImage("src/test/screenshots/$path.png", roborazziOptions = STRICT_SNAPSHOT)
        if (fontScale == 1f) {
            val diffs = ParityRegions.compare(ParityRegions.capture(recorder, density), ParityRegions.reference(screen))
            assertEquals("android-$screen against its mock", expectedDrift.sorted(), diffs.sorted())
            ParityText.assertText(rule, path)
        }
    }
}
