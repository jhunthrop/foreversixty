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
import gg.foreversixty.app.items.ItemTooltipContent
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipPhase
import gg.foreversixty.app.items.ItemTooltipSources
import gg.foreversixty.app.model.KeyDate
import kotlinx.serialization.builtins.ListSerializer
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performScrollTo
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
    private val context = BisContext(
        BisTestData.raid,
        BisTestData.raid,
        "Obnoxious Yell",
        BisTestData.classes(),
        RaidsOpening.from(BisTestData.fixture("cdn.dates", ListSerializer(KeyDate.serializer()))),
    )
    private val catalog = BisTestData.catalog()

    private fun content(phase: BisPhase) = @Composable {
        BisContent(BisUiState(phase, context, build = BisTestData.BUILD), catalog, query = "", nowMillis = BisTestData.NOW)
    }

    @Test fun `list`() = both("bis_list", content(BisPhase.Loaded(list, null)))

    @Test fun `offline`() = both("bis_offline", content(BisPhase.Loaded(list, BisTestData.NOW - DAY)))

    @Test fun `loading`() = both("bis_loading", content(BisPhase.Loading))

    @Test fun `error`() = both("bis_error", content(BisPhase.Failed))

    @Test fun `not cached`() = both("bis_not_cached", content(BisPhase.NotCached))

    /** The sheets draw over the tab bar on a device, so they are captured alone, as the sheet's content. */
    @Test fun `picker`() = both("bis_picker", screen = { PickerSheet() }, inShell = false)

    @Test fun `picker scrolled to Back at the largest size`() {
        capture("bis_picker_bottom_large", LARGE, inShell = false, scrollTo = "Back to Obnoxious Yell's list") { PickerSheet() }
    }

    @Test fun `tooltip`() = both("bis_tooltip", inShell = false, screen = {
        val row = list.rows.first()
        val request = tooltipRequest(row, requireNotNull(row.pick), list, context, BisTestData.BUILD)
        val sources = ItemTooltipSources(BisTestData.items().items.associateBy { it.id }, emptyList(), BisTestData.bis())
        ItemTooltipContent(request, ItemTooltipPhase.Loaded(ItemTooltipModel.build(request, sources, res)))
    })

    @Composable
    private fun PickerSheet() = BisPickerContent(context.copy(selection = BisTestData.raid.copy(faction = "alliance")), catalog) {}

    private fun both(name: String, screen: @Composable () -> Unit, inShell: Boolean = true) {
        capture("${name}_default", 1f, inShell, screen = screen)
        capture("${name}_large", LARGE, inShell, screen = screen)
    }

    private fun capture(
        name: String,
        fontScale: Float,
        inShell: Boolean,
        scrollTo: String? = null,
        screen: @Composable () -> Unit,
    ) {
        RuntimeEnvironment.setFontScale(fontScale)
        compose.setContent {
            ForeverTheme {
                if (inShell) ForeverTabScaffold(selected = ForeverTab.Bis, onSelect = {}) { screen() } else screen()
            }
        }
        scrollTo?.let { compose.onNodeWithText(it).performScrollTo() }
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
