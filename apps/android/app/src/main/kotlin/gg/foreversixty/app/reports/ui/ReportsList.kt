package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.common.SavedAnswerStrip
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowPlate
import gg.foreversixty.app.flows.FlowPlateMessage
import gg.foreversixty.app.flows.FlowPlateRetry
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowRule
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.model.ReportRow
import gg.foreversixty.app.reports.logic.ReportState
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.logic.ReportsContent
import gg.foreversixty.app.reports.logic.ReportsCopy
import gg.foreversixty.app.reports.logic.ReportsPhase
import java.time.Instant
import java.time.ZoneId

/** What the list hands back to its screen. */
data class ReportsListActions(
    val refresh: () -> Unit = {},
    val open: (ReportRow) -> Unit = {},
    val loadMore: () -> Unit = {},
    /** Upload a log: the site's upload, the companion's other half. */
    val upload: () -> Unit = {},
    /** The empty list's `Open the sample report`. */
    val openSample: () -> Unit = {},
    val signIn: () -> Unit = {},
)

/**
 * Your reports (design/mocks `reports`, `reports-empty`, `reports-failed`): the head with Upload a log,
 * then every report in one plate, newest first, each row its name, counts and visibility with its date
 * (and a state tag) under them. Stateless, so the parity tests drive every state.
 */
@Composable
fun ReportsList(
    phase: ReportsPhase,
    pageFailed: Boolean,
    now: Instant,
    actions: ReportsListActions = ReportsListActions(),
    zone: ZoneId = ZoneId.systemDefault(),
) {
    val copy = ReportsCopy(LocalResources.current, zone)
    val loaded = (phase as? ReportsPhase.Loaded)?.content
    FlowPage(top = 0.dp) {
        val count = loaded?.total?.takeIf { it > 0 }?.let { pluralStringResource(R.plurals.reports_count, it, it) }
        FlowHead(
            stringResource(R.string.reports_title),
            listOfNotNull(stringResource(R.string.reports_lead), count),
            actions = { FlowButton(stringResource(R.string.reports_upload), actions.upload, small = false) },
        )
        loaded?.offlineSince?.let { SavedAnswerStrip(OfflineStamp.format(it, now, zone), loaded.stale, actions.refresh) }
        FlowGap(FsSpace.space5)
        when (phase) {
            ReportsPhase.Loading -> Skeleton()
            ReportsPhase.Failed -> FlowPlate { FlowPlateMessage(listOf(stringResource(R.string.reports_failed))) { TryAgain(actions.refresh) } }
            ReportsPhase.SignedOut -> FlowPlate {
                FlowPlateMessage(listOf(stringResource(R.string.reports_signedOut))) {
                    FlowButton(stringResource(R.string.signin_button), actions.signIn, primary = true, small = false)
                }
            }
            is ReportsPhase.Loaded -> Loaded(phase.content, pageFailed, now, copy, actions)
        }
    }
}

@Composable
private fun TryAgain(retry: () -> Unit) {
    FlowButton(stringResource(R.string.common_tryAgain), retry)
}

@Composable
private fun Loaded(content: ReportsContent, pageFailed: Boolean, now: Instant, copy: ReportsCopy, actions: ReportsListActions) {
    if (content.rows.isEmpty()) {
        FlowPlate {
            FlowPlateMessage(listOf(stringResource(R.string.reports_empty))) {
                FlowButton(stringResource(R.string.reports_upload), actions.upload, primary = true, small = false)
                FlowLink(stringResource(R.string.reports_sample), onClick = actions.openSample)
            }
        }
        return
    }
    FlowPlate {
        content.rows.forEachIndexed { index, row -> ReportRowView(row, now, copy, last = index == content.rows.lastIndex, open = actions.open) }
    }
    Footer(content, pageFailed, actions)
}

@Composable
private fun Footer(content: ReportsContent, pageFailed: Boolean, actions: ReportsListActions) {
    when {
        pageFailed -> Box(Modifier.padding(top = FsSpace.space3)) {
            FlowPlate { FlowPlateRetry(stringResource(R.string.reports_failed), stringResource(R.string.common_tryAgain), actions.loadMore) }
        }
        content.hasMorePages && content.offlineSince == null -> {
            // The 56 dp spinner row loads the next page as it scrolls into view.
            LaunchedEffect(content.page) { actions.loadMore() }
            Box(Modifier.fillMaxWidth().heightIn(min = SPINNER_ROW), contentAlignment = Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold, modifier = Modifier.size(SPINNER), strokeWidth = 2.dp)
            }
        }
    }
}

/** One report (`.plate .row`, no lead): its name, counts and visibility, then its date and any state tag. */
@Composable
fun ReportRowView(row: ReportRow, now: Instant, copy: ReportsCopy, last: Boolean, open: (ReportRow) -> Unit) {
    val state = ReportState.of(row.status)
    FlowRow(
        Reports.name(row),
        copy.rowSub(row, now),
        last = last,
        inset = FsSpace.space4,
        trail = {
            Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space3), verticalAlignment = Alignment.CenterVertically) {
                FlowNote(copy.rowDate(row))
                copy.stateTag(state)?.let { StateTag(state, it) }
            }
        },
        onClick = { open(row) },
    )
}

@Composable
private fun StateTag(state: ReportState, text: String) {
    when (state) {
        ReportState.Live -> Tag(text, FsColor.kill)
        ReportState.Processing -> Tag(text, FsColor.gold)
        ReportState.Failed -> Tag(text, FsColor.wipe)
        ReportState.Done -> Unit
    }
}

/** Loading: three rows' worth of placeholders in the plate. */
@Composable
private fun Skeleton() {
    FlowPlate(Modifier.clearAndSetSemantics {}) {
        repeat(SKELETON_ROWS) { index ->
            Column(Modifier.padding(horizontal = FsSpace.space4, vertical = ROW_PAD), verticalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
                FlowSkeleton(SKELETON_TITLE, SKELETON_LINE)
                FlowSkeleton(SKELETON_SUB, SKELETON_SMALL)
                FlowSkeleton(SKELETON_DATE, SKELETON_SMALL)
            }
            if (index < SKELETON_ROWS - 1) FlowRule()
        }
    }
}

private const val SKELETON_ROWS = 3
private val ROW_PAD = 10.dp
private val SKELETON_TITLE = 180.dp
private val SKELETON_SUB = 220.dp
private val SKELETON_DATE = 48.dp
private val SKELETON_LINE = 16.dp
private val SKELETON_SMALL = 12.dp
private val SPINNER_ROW = 56.dp
private val SPINNER = 24.dp
