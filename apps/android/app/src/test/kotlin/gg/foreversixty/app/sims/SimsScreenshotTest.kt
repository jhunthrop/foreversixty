package gg.foreversixty.app.sims

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.sims.logic.RerunStatus
import gg.foreversixty.app.sims.logic.SimResultModel
import gg.foreversixty.app.sims.logic.SimResultPhase
import gg.foreversixty.app.sims.logic.SimsContent
import gg.foreversixty.app.sims.logic.SimsPhase
import gg.foreversixty.app.sims.ui.SimResultContent
import gg.foreversixty.app.sims.ui.SimsList
import gg.foreversixty.app.sims.ui.SimsListState
import java.time.Instant
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Every F5 screen and state on a Pixel 9 (412 x 915 dp), at the default size and at font scale 2.0
 * (spec 15). Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 * The PNGs double as the review captures in design/reviews/2026-10-31-mobile-sims.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class SimsScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val specs = SimsTestData.specs()
    private val now = SimsTestData.now
    private val offlineAt: Instant = Instant.parse("2026-12-15T04:40:00Z")
    private val bulk = mapOf("q8wd2c9vh4tz" to SimProgress("running", 0, null, 1, 40, 64))

    private fun capture(name: String, large: Boolean = false, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) {
                        Box(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) { screen() }
                    }
                }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/sims/$name.png", roborazziOptions = TOLERANT)
    }

    private fun list(name: String, phase: SimsPhase, state: SimsListState = SimsListState(entitled = true), large: Boolean = false) =
        capture(name, large) { SimsList(phase, state, specs, SimsTestData.BUILD, now) }

    private fun result(name: String, phase: SimResultPhase, looks: Boolean = true, large: Boolean = false) =
        capture(name, large) { SimResultContent(phase, if (looks) SimsTestData.looks() else emptyMap(), SimsTestData.BUILD, now) }

    private fun runModel() = SimResultModel.of("km3xibdrsiyr", SimsTestData.result(), SimsTestData.row("km3xibdrsiyr"), specs, resources())

    // F5.1 Sims

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun premium() = list("sims-premium-and", SimsPhase.Loaded(SimsTestData.content()), SimsListState(true, progress = bulk))

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun premiumLarge() = list("sims-premium-large-and", SimsPhase.Loaded(SimsTestData.content()), SimsListState(true, progress = bulk), large = true)

    @Test
    fun free() = list("sims-free-and", SimsPhase.Loaded(SimsTestData.content(server = emptyList(), mine = listOf(SimsTestData.browserRow))), SimsListState(false))

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun lapsed() = list(
        "sims-lapsed-and", SimsPhase.Loaded(SimsTestData.content()),
        SimsListState(false, rerun = mapOf("b7nf5x1pa3le" to RerunStatus.Lapsed), mayLinkToPlans = true),
    )

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun runAgainStates() {
        val failed = SimsTestData.row("b7nf5x1pa3le")
        val rows = (0 until 4).map { failed.copy(simId = "f$it") }
        val rerun = mapOf("f0" to RerunStatus.Pending, "f1" to RerunStatus.Started, "f2" to RerunStatus.Failed, "f3" to RerunStatus.Lapsed)
        val content = SimsContent(rows, listOf(SimsTestData.browserRow), 1, 1, null)
        list("sims-rerun-states-and", SimsPhase.Loaded(content), SimsListState(true, rerun = rerun))
    }

    @Test fun loading() = list("sims-loading-and", SimsPhase.Loading)

    @Test fun empty() = list("sims-empty-and", SimsPhase.Loaded(SimsTestData.content(server = emptyList(), mine = emptyList())))

    @Test fun error() = list("sims-error-and", SimsPhase.Failed)

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun offline() = list("sims-offline-and", SimsPhase.Loaded(SimsTestData.content(offline = offlineAt)))

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun pageFailed() = list(
        "sims-page-failed-and", SimsPhase.Loaded(SimsTestData.content().copy(total = 200)), SimsListState(true, pageFailed = true),
    )

    // F5.3 A result

    @Test fun resultRun() = result("result-run-and", SimResultPhase.Loaded(runModel(), null))

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun resultRunLarge() = result("result-run-large-and", SimResultPhase.Loaded(runModel(), null), large = true)

    @Test fun resultGear() {
        val model = SimResultModel.of("g5tq8n2rk1wa", SimsTestData.result(topGear = true, untitled = true), SimsTestData.topGearRow, specs, resources())
        result("result-gear-and", SimResultPhase.Loaded(model, null))
    }

    @Test fun resultOffline() = result("result-offline-and", SimResultPhase.Loaded(runModel(), offlineAt))

    @Test fun resultLoading() = result("result-loading-and", SimResultPhase.Loading)

    @Test fun resultFailed() = result("result-failed-and", SimResultPhase.Failed)
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_TALL = "w412dp-h1300dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XL = "w412dp-h2600dp-normal-long-notround-any-420dpi-keyshidden-nonav"
