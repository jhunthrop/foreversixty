package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.shape.RoundedCornerShape
import gg.foreversixty.app.designsystem.FsRadius
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.reports.logic.CoachPhase
import gg.foreversixty.app.reports.logic.MetersTarget
import gg.foreversixty.app.reports.logic.ParsesPhase
import gg.foreversixty.app.reports.logic.ReportPhase
import gg.foreversixty.app.reports.logic.ReportState
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.logic.ReportsCopy
import java.time.Instant
import java.time.ZoneId

/** What the report screen hands back. */
data class ReportActions(
    val back: () -> Unit = {},
    val openWeb: () -> Unit = {},
    val retry: () -> Unit = {},
    val retryParses: () -> Unit = {},
    val retryCoach: () -> Unit = {},
    val openMeters: (MetersTarget) -> Unit = {},
)

/** Everything F10.2 draws, apart from the store so screenshots drive every state. */
data class ReportScreenState(
    val phase: ReportPhase,
    val parses: ParsesPhase,
    val coach: CoachPhase,
    /** The bar's title before the report answers: the list row's zone, when pushed from it. */
    val title: String = "",
)

/** F10.2 and F10.5: a report, or the state that stands in for it. */
@Composable
fun ReportContent(state: ReportScreenState, now: Instant, actions: ReportActions = ReportActions(), zone: ZoneId = ZoneId.systemDefault()) {
    val copy = ReportsCopy(LocalResources.current, zone)
    val report = (state.phase as? ReportPhase.Loaded)?.report
    Column(Modifier.fillMaxWidth().padding(bottom = 24.dp)) {
        ReportsBar(report?.zone ?: state.title, stringResource(R.string.reports_back), actions.back, BarAction.openWeb(actions.openWeb))
        when (val phase = state.phase) {
            ReportPhase.Loading -> ReportSkeleton()
            ReportPhase.Failed -> ErrorRegion(stringResource(R.string.report_failed), actions.retry)
            ReportPhase.NotYours -> Notice(stringResource(R.string.report_notYours_title), stringResource(R.string.report_notYours_body), FsColor.strong)
            is ReportPhase.Loaded -> {
                phase.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now, zone))) }
                Loaded(phase.report, state, now, copy, actions)
            }
        }
    }
}

@Composable
private fun Loaded(report: Report, state: ReportScreenState, now: Instant, copy: ReportsCopy, actions: ReportActions) {
    Column(Modifier.padding(horizontal = 16.dp), verticalArrangement = Arrangement.spacedBy(22.dp)) {
        Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(10.dp)) {
                MicroLabel(stringResource(R.string.report_tag), color = FsColor.gold)
                Tag(copy.visibility(report.visibility), FsColor.nav)
            }
            Text(copy.reportMeta(report), style = FsType.meta, color = FsColor.muted)
        }
        when (ReportState.of(report.status)) {
            ReportState.Processing -> {
                Notice(stringResource(R.string.report_processing_title), stringResource(R.string.report_processing_body), FsColor.strong, inset = false)
                return@Column
            }
            ReportState.Failed -> {
                Notice(stringResource(R.string.report_unreadable_title), stringResource(R.string.report_unreadable_body), FsColor.wipe, inset = false)
                return@Column
            }
            ReportState.Live -> {
                val seconds = Reports.secondsSinceUpdate(report.heartbeatAt, report.lastUploadAt, now) ?: 0
                Notice(stringResource(R.string.report_live_title, seconds.toInt()), stringResource(R.string.report_live_body), FsColor.kill, inset = false)
            }
            ReportState.Done -> Unit
        }
        ParsesCard(state.parses, copy, actions.retryParses)
        FightsCard(report, copy, actions.openMeters)
        CoachCard(state.coach, copy, actions.retryCoach)
        Text(stringResource(R.string.report_footer), style = FsType.meta, color = FsColor.muted)
    }
}

/** A report state's two lines in a quiet card (spec F10.5). */
@Composable
fun Notice(title: String, body: String, color: Color, inset: Boolean = true) {
    val outer = if (inset) Modifier.padding(16.dp) else Modifier
    Column(outer.fillMaxWidth().quietCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(6.dp)) {
        Text(title, style = FsType.lead, color = color)
        Text(body, style = FsType.meta, color = FsColor.muted)
    }
}

@Composable
private fun ReportSkeleton() {
    Column(Modifier.padding(16.dp).clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(14.dp)) {
        SkeletonBar(80.dp, 18.dp)
        SkeletonBar(260.dp, 12.dp)
        Box(Modifier.fillMaxWidth().height(96.dp).background(FsColor.cardTop, RoundedCornerShape(FsRadius.panel)))
        Box(Modifier.fillMaxWidth().height(220.dp).background(FsColor.cardTop, RoundedCornerShape(FsRadius.panel)))
    }
}
