package gg.foreversixty.app.widgets

import android.content.Context
import android.widget.RemoteViews
import androidx.compose.runtime.getValue
import androidx.compose.runtime.key
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.wrapContentSize
import androidx.compose.ui.Alignment
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.platform.testTag
import android.graphics.BitmapFactory
import gg.foreversixty.app.home.HomeFixtures
import java.io.File
import androidx.compose.ui.unit.DpSize
import androidx.compose.ui.unit.dp
import androidx.compose.ui.viewinterop.AndroidView
import androidx.glance.appwidget.ExperimentalGlanceRemoteViewsApi
import androidx.glance.appwidget.GlanceRemoteViews
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.widgets.glance.RaidWidgetContent
import gg.foreversixty.app.widgets.glance.UpgradeWidgetContent
import gg.foreversixty.app.widgets.glance.WidgetIcons
import gg.foreversixty.app.widgets.logic.SignupOverlay
import gg.foreversixty.app.widgets.logic.WidgetFigures
import gg.foreversixty.app.widgets.logic.WidgetSnapshot
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.json.JsonNull
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * The real Glance widgets, rendered to RemoteViews by Glance and inflated as the launcher would,
 * over a wallpaper ground (spec §5.6, boards F14.1 and F14.2), at the default size and font scale
 * 2.0. Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 * The PNGs double as the review captures in design/reviews/2026-10-31-mobile-wave3-widgets/build.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
@OptIn(ExperimentalGlanceRemoteViewsApi::class)
class WidgetScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val zone = ZoneId.of("America/Chicago")
    private val now = WidgetTestData.boardNow
    private val small = DpSize(170.dp, 170.dp)
    private val medium = DpSize(360.dp, 170.dp)
    private var shown by mutableStateOf<Pair<RemoteViews, DpSize>?>(null)

    private fun capture(name: String, size: DpSize, large: Boolean = false, content: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        val context = ApplicationProvider.getApplicationContext<Context>()
        val views = runBlocking { GlanceRemoteViews().compose(context, size) { content() }.remoteViews }
        // The rule takes one setContent per test; each capture swaps the inflated widget in.
        if (shown == null) {
            compose.setContent {
                shown?.let { (remote, dp) ->
                    Box(Modifier.wrapContentSize(Alignment.TopStart).testTag(CAPTURE).background(WALLPAPER).padding(16.dp)) {
                        key(remote) { AndroidView(factory = { remote.apply(it, null) }, modifier = Modifier.size(dp)) }
                    }
                }
            }
        }
        shown = views to size
        compose.waitForIdle()
        compose.onNodeWithTag(CAPTURE).captureRoboImage("src/test/screenshots/widgets/$name.png", roborazziOptions = TOLERANT)
    }

    private fun figures() = WidgetFigures(resources(), zone)

    private fun raid(name: String, snapshot: WidgetSnapshot?, medium: Boolean, at: Instant = now, overlay: SignupOverlay? = null, large: Boolean = false) {
        val figures = figures().raid(snapshot, at, overlay)
        capture(name, if (medium) this.medium else small, large) { RaidWidgetContent(figures, medium, snapshot?.widget?.character?.key) }
    }

    private fun patch(vararg pairs: Pair<String, Any?>) = JsonObject(
        pairs.associate { (k, v) ->
            k to when (v) {
                null -> JsonNull
                is JsonObject -> v
                is Number -> JsonPrimitive(v)
                is Boolean -> JsonPrimitive(v)
                else -> JsonPrimitive(v.toString())
            }
        },
    )

    private val unsigned = patch("my" to patch("state" to null, "spec" to "Elemental", "role" to "dps"))
    private val noSpec = patch("my" to patch("state" to null, "spec" to null, "role" to "dps"))
    private val locked = patch("status" to "locked", "can_sign_up" to false)
    private val cancelled = patch("status" to "cancelled", "can_sign_up" to false)
    private fun answered(state: String) = patch("my" to patch("state" to state, "spec" to "Elemental", "role" to "dps"))

    private fun noRaid(): WidgetSnapshot = WidgetTestData.snapshot().let { it.copy(widget = it.widget.copy(nextRaid = null)) }

    @Test
    fun nextRaidSmall() {
        raid("and-raid2x2-in", WidgetTestData.snapshot(), medium = false)
        raid("and-raid2x2-late", WidgetTestData.snapshot(patch = answered("late")), medium = false)
        raid("and-raid2x2-out", WidgetTestData.snapshot(patch = answered("out")), medium = false)
        raid("and-raid2x2-not-signed-up", WidgetTestData.snapshot(patch = unsigned), medium = false)
        raid("and-raid2x2-choose-spec", WidgetTestData.snapshot(patch = noSpec), medium = false)
        raid("and-raid2x2-locked", WidgetTestData.snapshot(patch = locked), medium = false)
        raid("and-raid2x2-cancelled", WidgetTestData.snapshot(patch = cancelled), medium = false)
        raid("and-raid2x2-stale", WidgetTestData.snapshot(fetchedAt = "2026-12-08T09:20:00Z"), medium = false)
        raid("and-raid2x2-no-guild", WidgetTestData.snapshot("widget.get.no-guild"), medium = false)
        raid("and-raid2x2-no-raid", noRaid(), medium = false)
        raid("and-raid2x2-signed-out", null, medium = false)
    }

    @Test
    fun nextRaidMedium() {
        raid("and-raid4x2-in", WidgetTestData.snapshot(), medium = true)
        raid("and-raid4x2-late", WidgetTestData.snapshot(patch = answered("late")), medium = true)
        raid("and-raid4x2-out", WidgetTestData.snapshot(patch = answered("out")), medium = true)
        raid("and-raid4x2-not-signed-up", WidgetTestData.snapshot(patch = unsigned), medium = true)
        raid("and-raid4x2-choose-spec", WidgetTestData.snapshot(patch = noSpec), medium = true)
        raid("and-raid4x2-locked", WidgetTestData.snapshot(patch = locked), medium = true)
        raid("and-raid4x2-cancelled", WidgetTestData.snapshot(patch = cancelled), medium = true)
        raid("and-raid4x2-stale", WidgetTestData.snapshot(fetchedAt = "2026-12-08T09:20:00Z"), medium = true)
        raid(
            "and-raid4x2-not-saved", WidgetTestData.snapshot(), medium = true,
            overlay = SignupOverlay(7, "out", failedUntil = now.plusSeconds(10).toString()),
        )
    }

    @Test
    fun nextRaidMediumLargeText() {
        raid("and-raid4x2-in-fs2", WidgetTestData.snapshot(), medium = true, large = true)
    }

    @Test
    fun nextRaidMediumLargeTextNotSaved() {
        raid(
            "and-raid4x2-not-saved-fs2", WidgetTestData.snapshot(), medium = true, large = true,
            overlay = SignupOverlay(7, "out", failedUntil = now.plusSeconds(10).toString()),
        )
    }

    private fun upgrades(name: String, snapshot: WidgetSnapshot, medium: Boolean, large: Boolean = false) {
        val figures = figures().upgrades(snapshot)
        // The repo's own icon files stand in for the cached CDN files.
        val icons = figures.rows.associate { row ->
            val file = File(repoRoot, "data/builds/${HomeFixtures.BUILD}/icons/${row.icon}.webp")
            row.itemId to WidgetIcons.framed(BitmapFactory.decodeFile(file.path), row.quality)
        }
        capture(name, if (medium) this.medium else small, large) { UpgradeWidgetContent(figures, medium, icons) }
    }

    @Test
    fun nextUpgrade() {
        upgrades("and-upgrade2x2", WidgetTestData.snapshot(), medium = false)
        upgrades("and-upgrade4x2", WidgetTestData.snapshot(), medium = true)
        upgrades("and-upgrade2x2-no-gear", WidgetTestData.snapshot("widget.get.no-guild"), medium = false)
        upgrades("and-upgrade2x2-no-home", WidgetTestData.snapshot().let { it.copy(widget = it.widget.copy(upgrades = null)) }, medium = false)
    }

    @Test
    fun nextUpgradeLargeText() {
        upgrades("and-upgrade4x2-fs2", WidgetTestData.snapshot(), medium = true, large = true)
    }
}

/** A dark wallpaper behind the widgets, like the board's home screen. */
private val WALLPAPER = Brush.verticalGradient(listOf(Color(0xFF3A2414), Color(0xFF0B1220)))

private const val CAPTURE = "widget"
private val repoRoot: File = generateSequence(File("").absoluteFile) { it.parentFile }.first { File(it, "data/builds").isDirectory }

private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))
