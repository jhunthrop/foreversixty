package gg.foreversixty.app.widgets

import android.app.Notification
import android.content.Context
import android.widget.RemoteViews
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.wrapContentSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.unit.dp
import androidx.compose.ui.viewinterop.AndroidView
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.widgets.live.LiveNotifier
import gg.foreversixty.app.widgets.logic.RaidLiveContent
import gg.foreversixty.app.widgets.logic.SimLive
import gg.foreversixty.app.widgets.logic.SimLiveContent
import gg.foreversixty.app.widgets.logic.WidgetFigures
import gg.foreversixty.app.widgets.tile.TileFigures
import java.time.Instant
import java.time.ZoneId
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * The ongoing notifications as the system's own expanded template draws them (spec §5.6, board
 * F14.3), and You › Live surfaces. Record with recordRoborazziDebug; the PNGs are review captures.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class LiveScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val context: Context get() = ApplicationProvider.getApplicationContext()
    private var notification by mutableStateOf<RemoteViews?>(null)

    // The expanded template is the system's own picture of the notification; nothing else draws it.
    @Suppress("DEPRECATION")
    private fun notification(name: String, built: Notification) {
        val views = requireNotNull(Notification.Builder.recoverBuilder(context, built).createBigContentView())
        if (notification == null) {
            compose.setContent {
                notification?.let { remote ->
                    Box(Modifier.wrapContentSize(Alignment.TopStart).testTag(TAG).background(SHADE).padding(12.dp)) {
                        androidx.compose.runtime.key(remote) {
                            AndroidView(factory = { remote.apply(it, null) }, modifier = Modifier.width(388.dp).background(CARD))
                        }
                    }
                }
            }
        }
        notification = views
        compose.waitForIdle()
        compose.onNodeWithTag(TAG).captureRoboImage("src/test/screenshots/widgets/$name.png", roborazziOptions = TOLERANT)
    }

    /** The raid card at [at]; the chronometer counts against the device clock, so it is re-based on the capture's own now. */
    private fun raid(name: String, at: Instant, patch: kotlinx.serialization.json.JsonObject = WidgetTestData.signedUp(38)) {
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago")).raid(WidgetTestData.snapshot(patch = patch), at)
        val content = requireNotNull(RaidLiveContent.of(figures, resources()))
        val left = content.countdownTo?.let { java.time.Duration.between(at, it).seconds }
        val rebased = left?.let { content.copy(countdownTo = Instant.ofEpochMilli(org.robolectric.shadows.ShadowSystem.currentTimeMillis()).plusSeconds(it)) } ?: content
        notification(name, LiveNotifier.raidNotification(context, rebased))
    }

    @Test
    fun raidCountdown() = raid("and-live-raid", Instant.parse("2026-12-11T01:21:48Z"))

    @Test
    fun raidCountdownFiveMinutes() = raid("and-live-raid-5-min", Instant.parse("2026-12-11T01:55:00Z"))

    @Test
    fun raidCancelled() = raid("and-live-raid-cancelled", Instant.parse("2026-12-11T01:21:48Z"), WidgetTestData.cancelled)

    @Test
    fun simFailed() {
        val failed = SimLive("k3v9q2m8x1ab", "Elemental · Top Gear", SimLive.State.Failed, 12, 64, null, null)
        notification("and-live-sim-failed", LiveNotifier.simNotification(context, failed, SimLiveContent.of(failed, resources())))
    }

    @Test
    fun tileActive() = settings("and-tile-active", large = false) {
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago")).raid(WidgetTestData.snapshot(), WidgetTestData.boardNow)
        TileRender(TileFigures.of(figures, resources()))
    }

    @Test
    fun tileNoRaid() = settings("and-tile-no-raid", large = false) {
        val base = WidgetTestData.snapshot()
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago")).raid(base.copy(widget = base.widget.copy(nextRaid = null)), WidgetTestData.boardNow)
        TileRender(TileFigures.of(figures, resources()))
    }

    @Test
    fun simProgress() {
        val running = SimLive("k3v9q2m8x1ab", "Elemental · Top Gear", SimLive.State.Running, 40, 64, 120, null)
        notification("and-live-sim-running", LiveNotifier.simNotification(context, running, SimLiveContent.of(running, resources())))
    }

    @Test
    fun simFinished() {
        val done = SimLive("k3v9q2m8x1ab", "Elemental · Top Gear", SimLive.State.Finished, 64, 64, null, 412.0)
        notification("and-live-sim-finished", LiveNotifier.simNotification(context, done, SimLiveContent.of(done, resources())))
    }

    private fun settings(name: String, large: Boolean, content: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        compose.setContent {
            ForeverTheme {
                Box(Modifier.testTag(TAG).fillMaxWidth().background(FsColor.bg).padding(16.dp)) { content() }
            }
        }
        compose.onNodeWithTag(TAG).captureRoboImage("src/test/screenshots/widgets/$name.png", roborazziOptions = TOLERANT)
    }

    @Test
    fun liveSurfacesSettings() = settings("and-live-settings", large = false) {
        LiveSurfacesSection(raidOn = true, simOn = true, showsSim = true, onRaid = {}, onSim = {})
    }

    @Test
    fun liveSurfacesSettingsLarge() = settings("and-live-settings-fs2", large = true) {
        LiveSurfacesSection(raidOn = true, simOn = true, showsSim = true, onRaid = {}, onSim = {})
    }
}

/**
 * The quick-settings tile's state as Android 13's shade draws a tile (the board's F14.3 pill): the
 * system UI cannot be rendered on the JVM, so this pictures what [TileFigures] hands the tile, the
 * glyph, the label and the subtitle, active in the accent, inactive in the tile grey.
 */
@Composable
private fun TileRender(tile: TileFigures) {
    androidx.compose.foundation.layout.Row(
        Modifier.width(196.dp).tileGround(tile.active).padding(horizontal = 18.dp, vertical = 16.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        androidx.compose.material3.Icon(
            androidx.compose.ui.res.painterResource(gg.foreversixty.app.R.drawable.ic_tile_raid), contentDescription = null,
            tint = if (tile.active) FsColor.bg else FsColor.text,
        )
        androidx.compose.foundation.layout.Spacer(Modifier.width(12.dp))
        androidx.compose.foundation.layout.Column {
            androidx.compose.material3.Text(tile.label, color = if (tile.active) FsColor.bg else FsColor.strong)
            androidx.compose.material3.Text(tile.subtitle, color = if (tile.active) FsColor.bg else FsColor.muted)
        }
    }
}

private fun Modifier.tileGround(active: Boolean) =
    background(if (active) FsColor.gold else Color(0xFF283041), androidx.compose.foundation.shape.RoundedCornerShape(40.dp))

private const val TAG = "live"
private val SHADE = Color(0xFFDADDE3)
/** The system template draws dark text: a light shade, as on a light-theme phone. */
private val CARD = Color(0xFFF3F4F7)
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))
