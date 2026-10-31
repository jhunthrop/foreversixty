package gg.foreversixty.app.reports

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.home.HomeAnswer
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.home.HomeLogs
import gg.foreversixty.app.home.HomePhase
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.home.HomeTime
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.model.ReportList
import gg.foreversixty.app.model.ReportMe
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.reports.logic.CoachPhase
import gg.foreversixty.app.reports.logic.FightRow
import gg.foreversixty.app.reports.logic.FightRows
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.logic.MetersPhase
import gg.foreversixty.app.reports.logic.MetersTarget
import gg.foreversixty.app.reports.logic.Metric
import gg.foreversixty.app.reports.logic.ParsesPhase
import gg.foreversixty.app.reports.logic.ReportPhase
import gg.foreversixty.app.reports.logic.ReportsContent
import gg.foreversixty.app.reports.logic.ReportsPhase
import gg.foreversixty.app.reports.ui.MetersActions
import gg.foreversixty.app.reports.ui.MetersDragHandle
import gg.foreversixty.app.reports.ui.MetersSheetContent
import gg.foreversixty.app.reports.ui.MetersSheetState
import gg.foreversixty.app.reports.ui.ReportContent
import gg.foreversixty.app.reports.ui.ReportScreenState
import gg.foreversixty.app.reports.ui.ReportsList
import gg.foreversixty.app.sims.SimsTestData
import java.time.Instant
import java.time.ZoneOffset
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Every F10 screen and state on a Pixel 9 (412 x 915 dp), at the default size and at font scale 2.0
 * (spec W3 §9). Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 * The PNGs double as the review captures in design/reviews/2026-10-31-mobile-wave3-reports.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class ReportsScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val now = ReportsTestData.now
    private val zone = ReportsTestData.zone
    private val offlineAt: Instant = Instant.parse("2026-12-15T22:10:00Z")
    private val specs = SimsTestData.specs()

    private fun capture(name: String, large: Boolean = false, scrolls: Boolean = true, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) {
                        // A route scrolls itself; the stateless screens are given the scroll the route would.
                        val ground = Modifier.fillMaxSize().background(FsColor.bg)
                        Box(if (scrolls) ground.verticalScroll(rememberScrollState()) else ground) { screen() }
                    }
                }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/reports/$name.png", roborazziOptions = TOLERANT)
    }

    private fun list(name: String, phase: ReportsPhase, large: Boolean = false, pageFailed: Boolean = false) =
        capture(name, large) { ReportsList(phase, pageFailed, ReportsTestData.viewer(), now, zone = zone) }

    private fun loaded(list: ReportList, offline: Instant? = null) =
        ReportsPhase.Loaded(ReportsContent(list.rows, list.total, list.page, offline))

    private fun report(
        name: String,
        report: Report = ReportsTestData.report(),
        me: ReportMe? = ReportsTestData.me(),
        coach: CoachPhase = CoachPhase.Hidden,
        offline: Instant? = null,
        large: Boolean = false,
    ) = screen(name, ReportScreenState(ReportPhase.Loaded(report, offline), me?.let { ParsesPhase.Loaded(it) } ?: ParsesPhase.Hidden, coach), large)

    private fun screen(name: String, state: ReportScreenState, large: Boolean = false) =
        capture(name, large) { ReportContent(state, now, zone = zone) }

    private fun meters(name: String, target: MetersTarget, metric: Metric, phase: MetersPhase, large: Boolean = false) =
        capture(name, large) {
            // The sheet as Material draws it at its large detent: the rounded top and the drag handle.
            Column(Modifier.padding(top = 24.dp).fillMaxWidth().background(GameColors.sheet, RoundedCornerShape(topStart = 28.dp, topEnd = 28.dp))) {
                Box(Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) { MetersDragHandle() }
                MetersSheetContent(MetersSheetState(target, metric, phase, showAll = false), specs, SimsTestData.BUILD, MetersActions(), zone)
            }
        }

    private val kill get() = FightRows.of(ReportsTestData.report()).filterIsInstance<FightRow.Boss>()
        .single { it.fight.index == ReportsTestData.KILL_FIGHT }.let { MetersTarget.Pull(it.fight, it.pull) }
    private val trash get() = MetersTarget.Trash(requireNotNull(ReportsTestData.report().trashGroups).first())

    // F10.1 Your reports, F10.6, F10.7

    @Test fun list() = list("reports-list-and", loaded(ReportsTestData.list()))

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun listLarge() = list("reports-list-large-and", loaded(ReportsTestData.list()), large = true)

    @Test fun empty() = list("reports-empty-and", loaded(ReportsTestData.emptyList()))

    @Test fun loading() = list("reports-loading-and", ReportsPhase.Loading)

    @Test fun error() = list("reports-error-and", ReportsPhase.Failed)

    @Test fun signedOut() = list("reports-signedout-and", ReportsPhase.SignedOut)

    @Test fun offline() = list("reports-offline-and", loaded(ReportsTestData.list(), offlineAt))

    // F10.2 A report

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun reportKill() = report("report-kill-and")

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun reportWipes() = report("report-wipes-and", ReportsTestData.onlyWipes(), ReportsTestData.meEmpty())

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun reportTwoZones() = report("report-twozones-and", ReportsTestData.twoZones())

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun reportOffline() = report("report-offline-and", offline = offlineAt)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun reportLarge() = report("report-kill-large-and", large = true)

    @Test fun reportLoading() = screen("report-loading-and", ReportScreenState(ReportPhase.Loading, ParsesPhase.Loading, CoachPhase.Hidden, "Onyxia's Lair"))

    @Test fun parsesFailed() = screen("report-parsesfailed-and", ReportScreenState(ReportPhase.Loaded(ReportsTestData.report(), null), ParsesPhase.Failed, CoachPhase.Hidden))

    // F10.4 Coach

    @Test @Config(qualifiers = PIXEL_9_XTALL)
    fun reportPremium() = report("report-premium-and", coach = CoachPhase.Ready(ReportsTestData.coach()))

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun reportFree() = report("report-free-and")

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun coachWriting() = report("report-coachwriting-and", coach = CoachPhase.Writing)

    // F10.5 Report states

    /** The live Blackrock Spire run, never the finished Onyxia night. */
    @Test fun live() = report("report-live-and", ReportsTestData.live(), ReportsTestData.meEmpty())

    @Test fun processing() = report("report-processing-and", ReportsTestData.report().copy(status = "processing", fights = emptyList(), trashGroups = emptyList()), me = null)

    @Test fun unreadable() = report("report-failed-and", ReportsTestData.report().copy(status = "failed", fights = emptyList(), trashGroups = emptyList()), me = null)

    @Test fun notYours() = screen("report-notyours-and", ReportScreenState(ReportPhase.NotYours, ParsesPhase.Hidden, CoachPhase.Hidden))

    @Test fun reportFailed() = screen("report-error-and", ReportScreenState(ReportPhase.Failed, ParsesPhase.Hidden, CoachPhase.Hidden))

    // F10.3 A fight

    @Test fun metersDamage() = meters("meters-damage-and", kill, Metric.Damage, MetersPhase.Loaded(ReportsTestData.meters()))

    @Test fun metersHealing() = meters("meters-healing-and", kill, Metric.Healing, MetersPhase.Loaded(ReportsTestData.healing()))

    @Test fun metersTaken() = meters("meters-taken-and", kill, Metric.Taken, MetersPhase.Loaded(ReportsTestData.taken()))

    @Test fun metersTrash() = meters("meters-trash-and", trash, Metric.Damage, MetersPhase.Loaded(ReportsTestData.trashMeters()))

    @Test fun metersLoading() = meters("meters-loading-and", kill, Metric.Damage, MetersPhase.Loading)

    @Test fun metersFailed() = meters("meters-failed-and", kill, Metric.Damage, MetersPhase.Failed)

    @Test fun metersFlagged() = meters("meters-flagged-and", kill, Metric.Damage, MetersPhase.Flagged)

    @Test
    fun metersLarge() = meters("meters-damage-large-and", kill, Metric.Damage, MetersPhase.Loaded(ReportsTestData.meters()), large = true)

    // The /app/report/{id} link: Home > Your reports > the report, through the navigator the link uses.

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun deepLink() {
        val feature = ReportsFeature(
            FakeReportsService(), MemoryRaidCache(), specs, { SimsTestData.BUILD }, { true }, { now }, zone,
        )
        assertTrue(feature.navigator.open(requireNotNull(DeepLink.parse(DeepLink.Report(ReportsTestData.ID).url()))))
        assertTrue(feature.navigator.showsList && feature.navigator.wantsHome)
        assertEquals(ReportsTestData.ID, feature.navigator.report?.id)
        capture("report-deeplink-and", scrolls = false) { ReportsRoute(feature, ReportsTestData.viewer(premium = false), signIn = {}, onBack = {}) }
    }

    // F10.8 Home with the Logs row

    @Test @Config(qualifiers = PIXEL_9_HOME)
    fun homeLogs() = capture("home-logs-and") {
        val res = LocalResources.current
        val presenter = HomePresenter(
            HomeFixtures.velaren, emptyList(), HomeTime(HomeFixtures.now, ZoneOffset.UTC), res,
            logs = HomeLogs(Reports.lastRaid(ReportsTestData.list().rows)),
        )
        HomeAnswer(HomePhase.Loaded(HomeFixtures.home()), presenter, HomeFixtures.BUILD, null)
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_TALL = "w412dp-h1300dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XTALL = "w412dp-h2000dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_HOME = "w412dp-h2400dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XL = "w412dp-h2600dp-normal-long-notround-any-420dpi-keyshidden-nonav"
