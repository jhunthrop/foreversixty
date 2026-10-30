package gg.foreversixty.app.bis

import android.content.Context
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.items.ItemTooltipBody
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.ZoneOffset

/**
 * The BiS tab's states (spec F3) at Pixel 9 size, at the default font scale and at 2.0. Icons do
 * not load here (no icon source), so each tile shows its quality ring.
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class BisScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val res = ApplicationProvider.getApplicationContext<Context>().resources
    private val list: BisList = run {
        val file = BisTestData.bis()
        val band = requireNotNull(file.band(BisTestData.raid.range, "horde"))
        BisListBuilder(res, BisTestData.items().items, setOf(12640), ZoneOffset.UTC)
            .list(file, band, "Fury Warrior", res.getString(R.string.bis_range_p1))
    }
    private val context = BisContext(BisTestData.raid, BisTestData.raid, "Obnoxious Yell", BisTestData.classes())
    private val catalog = BisTestData.catalog()

    private fun content(phase: BisPhase) = @Composable {
        BisContent(BisUiState(phase, context), catalog, query = "", nowMillis = BisTestData.NOW)
    }

    @Test fun `list`() = both("bis_list", content(BisPhase.Loaded(list, null)))

    @Test fun `offline`() = both("bis_offline", content(BisPhase.Loaded(list, BisTestData.NOW - DAY)))

    @Test fun `loading`() = both("bis_loading", content(BisPhase.Loading))

    @Test fun `error`() = both("bis_error", content(BisPhase.Failed))

    @Test fun `not cached`() = both("bis_not_cached", content(BisPhase.NotCached))

    @Test fun `picker`() = both("bis_picker") {
        BisPickerContent(context.copy(selection = BisTestData.raid.copy(faction = "alliance")), catalog) {}
    }

    @Test fun `tooltip`() = both("bis_tooltip") {
        val row = list.rows.first()
        ItemTooltipBody(tooltipContent(row, requireNotNull(row.pick), list, res))
    }

    private fun both(name: String, screen: @Composable () -> Unit) {
        capture("${name}_default", 1f, screen)
        capture("${name}_large", LARGE, screen)
    }

    private fun capture(name: String, fontScale: Float, screen: @Composable () -> Unit) {
        RuntimeEnvironment.setFontScale(fontScale)
        compose.setContent {
            ForeverTheme {
                ForeverTabScaffold(selected = ForeverTab.Bis, onSelect = {}) { screen() }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/$name.png", roborazziOptions = TOLERANT)
    }

    private companion object {
        const val DAY = 86_400_000L
        const val LARGE = 2f
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
