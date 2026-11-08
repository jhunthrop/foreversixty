package gg.foreversixty.app.reports

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.model.ReportList
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.reports.logic.CoachPhase
import gg.foreversixty.app.reports.logic.ParsesPhase
import gg.foreversixty.app.reports.logic.ReportPhase
import gg.foreversixty.app.reports.logic.ReportsContent
import gg.foreversixty.app.reports.logic.ReportsPhase
import gg.foreversixty.app.reports.ui.ReportContent
import gg.foreversixty.app.reports.ui.ReportScreenState
import gg.foreversixty.app.reports.ui.ReportsList
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** Your reports and a report's stand-in states (design/mocks/ws_app.py `reports*`, `report-error`, `report-processing`), pushed under Home. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class ReportsParityTest {
    @get:Rule val compose = createComposeRule()

    private val now = ReportsTestData.now
    private val zone = ReportsTestData.zone

    @Test fun reports() = pushed("reports", "Home") { YourReports(loaded(ReportsParityFixtures.oneReport())) }

    @Test fun reportsAt2() = pushed("reports", "Home", fontScale = 2f) { YourReports(loaded(ReportsParityFixtures.oneReport())) }

    @Test fun reportsEmpty() = pushed("reports-empty", "Home") { YourReports(loaded(ReportsTestData.emptyList())) }

    @Test fun reportsFailed() = pushed("reports-failed", "Home") { YourReports(ReportsPhase.Failed) }

    @Test fun reportError() = pushed("report-error", "Logs") { OneReport(ReportPhase.Failed) }

    @Test fun reportProcessing() = pushed("report-processing", "Logs") {
        OneReport(ReportPhase.Loaded(ReportsParityFixtures.processing(now), null))
    }

    @Composable
    private fun YourReports(phase: ReportsPhase) = ReportsList(phase, pageFailed = false, viewer = ReportsTestData.viewer(), now = now, zone = zone)

    @Composable
    private fun OneReport(phase: ReportPhase) =
        ReportContent(ReportScreenState(phase, ParsesPhase.Hidden, CoachPhase.Hidden), now, zone = zone)

    private fun loaded(list: ReportList) = ReportsPhase.Loaded(ReportsContent(list.rows, list.total, list.page, null))

    private fun pushed(screen: String, backTo: String, fontScale: Float = 1f, content: @Composable () -> Unit) =
        ParityCheck.state(compose, screen, fontScale) {
            ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}, badges = TabBadges(raid = 1, activity = 3)) {
                Column {
                    FlowBackBar(backTo, {})
                    content()
                }
            }
        }
}

/** The mocks' data: one finished report, and an upload still being read two minutes after it arrived. */
object ReportsParityFixtures {
    fun oneReport(): ReportList {
        val list = ReportsTestData.list()
        val row = list.rows.single { it.id == ReportsTestData.ID }
        return list.copy(rows = listOf(row.copy(title = "Report, 17 Sept 2026")), total = 1)
    }

    fun processing(now: java.time.Instant) = ReportsTestData.report().copy(
        status = "processing",
        zone = null,
        title = "",
        createdAt = now.minusSeconds(TWO_MINUTES).toString(),
    )

    private const val TWO_MINUTES = 120L
}
