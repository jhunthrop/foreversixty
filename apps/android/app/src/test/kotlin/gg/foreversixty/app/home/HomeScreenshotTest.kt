package gg.foreversixty.app.home

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipContent
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipPhase
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import java.time.Instant
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * F2's states inside the tab shell at Pixel 9 size, at the default text size and font scale 2.0,
 * from the contract fixtures with the repo's own icons. The same list is drawn on iOS
 * (`HomeSnapshotTests.swift`); the PNGs are also the review captures.
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class HomeScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    @Composable
    private fun Answer(phase: HomePhase, offline: String? = null) {
        val presenter = HomeFixtures.presenter(LocalResources.current)
        Box(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) {
            HomeAnswer(phase, presenter, HomeFixtures.BUILD, offline)
        }
    }

    @Test
    fun `F2_1 the answer`() = capture("home-default") { Answer(HomePhase.Loaded(HomeFixtures.home())) }

    @Test
    fun `F2_1 the answer at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        capture("home-default-xl") { Answer(HomePhase.Loaded(HomeFixtures.home())) }
    }

    @Test
    fun `F2_2 loading`() = capture("home-loading") { Answer(HomePhase.Loading) }

    @Test
    fun `F2_3 offline with the stamp`() = capture("home-offline") {
        Answer(HomePhase.Loaded(HomeFixtures.home(), Instant.parse("2026-10-29T21:41:00Z")), "Thu 9:41 pm")
    }

    @Test
    fun `F2_4 no gear export`() = capture("home-nogear") { Answer(HomePhase.Loaded(HomeFixtures.home("character.home.no-gear"))) }

    @Test
    fun `under 20`() = capture("home-under20") { Answer(HomePhase.Loaded(HomeFixtures.home(state = "under_20"))) }

    @Test
    fun `no spec`() = capture("home-nospec") { Answer(HomePhase.Loaded(HomeFixtures.home(state = "no_spec"))) }

    @Test
    fun `no list`() = capture("home-nolist") { Answer(HomePhase.Loaded(HomeFixtures.home(state = "no_list"))) }

    @Test
    fun `every slot matches`() = capture("home-allmatch") { Answer(HomePhase.Loaded(HomeFixtures.home("character.home.empty"))) }

    @Test
    fun `F2_5 error with the cached Next rows`() = capture("home-error") { Answer(HomePhase.Failed(HomeFixtures.home())) }

    @Test
    fun `F2_6 the tooltip partly expanded`() = capture("tooltip-medium") { TooltipOver(252605, 470) }

    @Test
    fun `F2_6 the tooltip fully expanded`() = capture("tooltip-large") { TooltipOver(252605, 760) }

    @Test
    fun `F2_6 the tooltip at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        capture("tooltip-large-xl") { TooltipOver(252605, 760) }
    }

    @Test
    fun `F2_6 a vendor item with a standing`() = capture("tooltip-vendor") { TooltipOver(276202, 760) }

    @Test
    fun `F2_6 a weapon`() = capture("tooltip-weapon") { TooltipOver(250619, 760) }

    @Test
    fun `F2_7 Your gear`() = capture("gear-default") { Gear() }

    @Test
    @Config(qualifiers = PIXEL_9_TALL)
    fun `F2_7 Your gear, all 17 slots`() = capture("gear-full") { Gear() }

    @Test
    @Config(qualifiers = PIXEL_9_TALL)
    fun `F2_1 the answer, full length`() = capture("home-full") { Answer(HomePhase.Loaded(HomeFixtures.home())) }

    @Test
    fun `F2_7 Your gear at font scale 2`() {
        RuntimeEnvironment.setFontScale(2f)
        capture("gear-default-xl") { Gear() }
    }

    @Composable
    private fun Gear() {
        val res = LocalResources.current
        val home = HomeFixtures.home()
        val sources = HomeFixtures.tooltipSources()
        val probe = HomePresenter.tooltipRequest(ItemLook(0), home, HomeFixtures.velaren, res)
        val blocker = GearModel.offHandBlocker(home, emptyMap(), sources.items, ItemTooltipModel.band(probe, sources.bis))
        Box(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) {
            GearList(
                GearModel.rows(home, emptyMap(), sources.items, res, blocker),
                "What you wear, and the pick for 60 · raid. Tap an item for its tooltip.",
                home.dataBuild,
            )
        }
    }

    /** A sheet at a height over the dimmed Home, for the tooltip captures. */
    @Composable
    private fun TooltipOver(id: Int, height: Int) {
        val res = LocalResources.current
        val home = HomeFixtures.home()
        val sources = HomeFixtures.tooltipSources()
        val item = sources.items[id]
        val request = HomePresenter.tooltipRequest(
            ItemLook(id, item?.name, item?.quality, item?.icon, item?.itemLevel), home, HomeFixtures.velaren, res,
            wornId = 279253,
        )
        val model = ItemTooltipModel.build(request, sources, res)
        Box(Modifier.fillMaxSize(), contentAlignment = Alignment.BottomCenter) {
            Answer(HomePhase.Loaded(home))
            Box(Modifier.fillMaxSize().background(Color.Black.copy(alpha = 0.45f)))
            Column(
                Modifier.fillMaxWidth().height(height.dp).clip(RoundedCornerShape(topStart = 28.dp, topEnd = 28.dp))
                    .background(FsColor.raised),
            ) {
                Box(Modifier.fillMaxWidth().height(22.dp), contentAlignment = Alignment.Center) {
                    Box(Modifier.fillMaxWidth(0.1f).height(4.dp).clip(RoundedCornerShape(2.dp)).background(FsColor.muted))
                }
                ItemTooltipContent(request, ItemTooltipPhase.Loaded(model))
            }
        }
    }

    private fun capture(name: String, screen: @Composable () -> Unit) {
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) { screen() }
                }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/$name.png", roborazziOptions = TOLERANT)
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(
    compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F),
)

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"

/** Pixel 9's width at a tall height, for the full-length captures. */
private const val PIXEL_9_TALL = "w412dp-h2000dp-normal-long-notround-any-420dpi-keyshidden-nonav"
