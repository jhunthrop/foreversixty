package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsRadius
import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.common.SavedAnswerStrip
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowPlate
import gg.foreversixty.app.flows.FlowPlateMessage
import gg.foreversixty.app.flows.FlowPlateRetry
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.phrase
import gg.foreversixty.app.reports.logic.CoachPhase
import gg.foreversixty.app.reports.logic.ParsesPhase
import gg.foreversixty.app.reports.logic.ReportPhase
import gg.foreversixty.app.reports.logic.ReportState
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.logic.ReportsCopy
import gg.foreversixty.app.reports.logic.MetersTarget
import java.time.Instant
import java.time.ZoneId

/** What the report screen hands back. */
data class ReportActions(
    val openWeb: () -> Unit = {},
    val retry: () -> Unit = {},
    val retryParses: () -> Unit = {},
    val retryCoach: () -> Unit = {},
    val openMeters: (MetersTarget) -> Unit = {},
)

/** Everything F10.2 draws, apart from the store so the tests drive every state. */
data class ReportScreenState(
    val phase: ReportPhase,
    val parses: ParsesPhase,
    val coach: CoachPhase,
    /** The head's title before the report answers: the list row's zone, when pushed from it. */
    val title: String = "",
)

/**
 * A report (design/mocks `report-*`): the head (its name, date and visibility, the Summary tab and More,
 * which opens the full report on the site), then your parses, the fights and the coach; or the state
 * that stands in for it (`report-error`, `report-processing`).
 */
@Composable
fun ReportContent(state: ReportScreenState, now: Instant, actions: ReportActions = ReportActions(), zone: ZoneId = ZoneId.systemDefault()) {
    val copy = ReportsCopy(LocalResources.current, zone)
    FlowPage(top = 0.dp) {
        when (val phase = state.phase) {
            ReportPhase.Loading -> Loading(state.title)
            ReportPhase.Failed -> {
                FlowHead(stringResource(R.string.report_title), emptyList())
                FlowGap(FsSpace.space5)
                FlowPlate { FlowPlateRetry(stringResource(R.string.report_failed), stringResource(R.string.common_tryAgain), actions.retry) }
            }
            ReportPhase.NotYours -> {
                FlowHead(stringResource(R.string.report_title), emptyList())
                FlowGap(FsSpace.space5)
                FlowPlate { FlowPlateMessage(listOf(stringResource(R.string.report_notYours_title), stringResource(R.string.report_notYours_body))) }
            }
            is ReportPhase.Loaded -> {
                Head(phase.report, now, copy, actions.openWeb)
                phase.offlineSince?.let { SavedAnswerStrip(OfflineStamp.format(it, now, zone), phase.stale, actions.retry) }
                Body(phase.report, state, now, copy, actions)
            }
        }
    }
}

@Composable
private fun Head(report: Report, now: Instant, copy: ReportsCopy, openWeb: () -> Unit) {
    val res = LocalResources.current
    val processing = ReportState.of(report.status) == ReportState.Processing
    val title = Reports.headTitle(report) ?: stringResource(R.string.report_upload)
    // A report still being read has no date or logger yet: its upload's age stands in.
    val facts = if (processing) {
        listOfNotNull(Reports.instantOrNull(report.createdAt)?.let { stringResource(R.string.report_uploaded, RelativeAge.of(it, now).phrase(res)) })
    } else {
        listOf(copy.reportMeta(report))
    }
    FlowHead(
        title,
        facts,
        meta = if (processing) null else ({ Tag(copy.visibility(report.visibility), FsColor.muted) }),
        tabs = { ReportTabs(openWeb) },
    )
}

@Composable
private fun Body(report: Report, state: ReportScreenState, now: Instant, copy: ReportsCopy, actions: ReportActions) {
    when (ReportState.of(report.status)) {
        ReportState.Processing -> Message(listOf(stringResource(R.string.report_processing_title), stringResource(R.string.report_processing_body)))
        ReportState.Failed -> Message(listOf(stringResource(R.string.report_unreadable_title), stringResource(R.string.report_unreadable_body)))
        ReportState.Live, ReportState.Done -> {
            if (ReportState.of(report.status) == ReportState.Live) {
                val seconds = Reports.secondsSinceUpdate(report.heartbeatAt, report.lastUploadAt, now) ?: 0
                Message(listOf(stringResource(R.string.report_live_title, seconds.toInt()), stringResource(R.string.report_live_body)))
            }
            ParsesSection(state.parses, copy, actions.retryParses)
            FightsSection(report, copy, actions.openMeters)
            CoachBlock(state.coach, copy, actions.retryCoach)
            FlowNote(stringResource(R.string.report_footerMore), Modifier.padding(top = FsSpace.space6))
        }
    }
}

@Composable
private fun Message(lines: List<String>) {
    FlowGap(FsSpace.space5)
    FlowPlate { FlowPlateMessage(lines) }
}

/** Loading: the head's title as the list row named it, then the plates' placeholders. */
@Composable
private fun Loading(title: String) {
    FlowHead(title.ifEmpty { stringResource(R.string.report_title) }, emptyList())
    FlowGap(FsSpace.space5)
    Column(Modifier.clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(FsSpace.space5)) {
        Placeholder(SKELETON_PARSES)
        Placeholder(SKELETON_FIGHTS)
    }
}

@Composable
private fun Placeholder(height: Dp) {
    Box(Modifier.fillMaxWidth().height(height).background(FsColor.raised, RoundedCornerShape(FsRadius.plate)))
}

private val SKELETON_PARSES = 96.dp
private val SKELETON_FIGHTS = 220.dp
