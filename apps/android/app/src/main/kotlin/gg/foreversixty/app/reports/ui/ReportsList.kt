package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import java.util.Locale
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.pillSurface
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.model.ReportRow
import gg.foreversixty.app.reports.logic.ReportState
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.logic.ReportsContent
import gg.foreversixty.app.reports.logic.ReportsCopy
import gg.foreversixty.app.reports.logic.ReportsPhase
import gg.foreversixty.app.reports.logic.ReportsViewer
import gg.foreversixty.app.reports.logic.Visibility
import java.time.Instant
import java.time.ZoneId

/** What the list hands back to its screen. */
data class ReportsListActions(
    val back: () -> Unit = {},
    val refresh: () -> Unit = {},
    val open: (ReportRow) -> Unit = {},
    val loadMore: () -> Unit = {},
    val getCompanion: () -> Unit = {},
    val signIn: () -> Unit = {},
)

object ReportsLayout {
    val rowHeight = 68.dp
    val crest = 36.dp
    const val SKELETON_ROWS = 4

    /** From this font scale the tags and glyph drop under the row's words. */
    const val STACK_SCALE = 1.3f
}

/** F10.1 as drawn from a phase: stateless, so screenshots drive every state. */
@Composable
fun ReportsList(
    phase: ReportsPhase,
    pageFailed: Boolean,
    viewer: ReportsViewer,
    now: Instant,
    actions: ReportsListActions = ReportsListActions(),
    zone: ZoneId = ZoneId.systemDefault(),
) {
    val copy = ReportsCopy(LocalResources.current, zone)
    Column(Modifier.fillMaxWidth().padding(bottom = 24.dp)) {
        ReportsBar(stringResource(R.string.reports_title), stringResource(R.string.tab_home), actions.back, BarAction.refresh(actions.refresh))
        when (phase) {
            ReportsPhase.Loading -> Skeleton()
            ReportsPhase.Failed -> ErrorRegion(stringResource(R.string.reports_failed), actions.refresh)
            ReportsPhase.SignedOut -> Column(Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(14.dp)) {
                Text(stringResource(R.string.reports_signedOut), style = FsType.lead, color = FsColor.strong)
                BrandButton(stringResource(R.string.signin_button), BrandButtonKind.Primary, actions.signIn)
            }
            is ReportsPhase.Loaded -> {
                phase.content.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now, zone))) }
                Loaded(phase.content, pageFailed, viewer, now, copy, actions)
            }
        }
    }
}

@Composable
private fun Loaded(content: ReportsContent, pageFailed: Boolean, viewer: ReportsViewer, now: Instant, copy: ReportsCopy, actions: ReportsListActions) {
    if (content.rows.isEmpty()) {
        Column(Modifier.padding(16.dp).fillMaxWidth().quietCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(14.dp)) {
            Text(stringResource(R.string.reports_empty), style = FsType.body, color = FsColor.text)
            BrandButton(stringResource(R.string.reports_getCompanion), BrandButtonKind.Secondary, actions.getCompanion, height = 48.dp, fill = false)
        }
        return
    }
    val sections = Reports.sections(content.rows, now)
    Column(Modifier.padding(horizontal = 16.dp).padding(top = 8.dp), verticalArrangement = Arrangement.spacedBy(22.dp)) {
        Section(stringResource(R.string.reports_thisWeek), sections.thisWeek, viewer, now, copy, actions)
        Section(stringResource(R.string.reports_earlier), sections.earlier, viewer, now, copy, actions)
        Footer(content, pageFailed, actions)
        Text(stringResource(R.string.reports_footer), style = FsType.meta, color = FsColor.muted)
    }
}

@Composable
private fun Section(title: String, rows: List<ReportRow>, viewer: ReportsViewer, now: Instant, copy: ReportsCopy, actions: ReportsListActions) {
    if (rows.isEmpty()) return
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(title)
        Column(Modifier.fillMaxWidth().quietCard()) {
            rows.forEachIndexed { index, row ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                ReportRowView(row, viewer, now, copy, actions.open)
            }
        }
    }
}

@Composable
private fun Footer(content: ReportsContent, pageFailed: Boolean, actions: ReportsListActions) {
    when {
        pageFailed -> Row(Modifier.fillMaxWidth(), verticalAlignment = Alignment.CenterVertically) {
            Text(stringResource(R.string.reports_failed), style = FsType.meta, color = FsColor.wipe, modifier = Modifier.weight(1f))
            BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, actions.loadMore, height = 48.dp, fill = false)
        }
        content.hasMorePages && content.offlineSince == null -> {
            // The 56 dp spinner row loads the next page as it scrolls into view.
            LaunchedEffect(content.page) { actions.loadMore() }
            Box(Modifier.fillMaxWidth().heightIn(min = 56.dp), contentAlignment = Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold, modifier = Modifier.size(24.dp), strokeWidth = 2.dp)
            }
        }
    }
}

/** One report (68 dp min): crest, name, meta, the state tag, the visibility glyph and a chevron. */
@Composable
fun ReportRowView(row: ReportRow, viewer: ReportsViewer, now: Instant, copy: ReportsCopy, open: (ReportRow) -> Unit) {
    val state = ReportState.of(row.status)
    val label = copy.rowLabel(row, now)
    val stacked = LocalDensity.current.fontScale >= ReportsLayout.STACK_SCALE
    Row(
        Modifier.fillMaxWidth().clickable { open(row) }.heightIn(min = ReportsLayout.rowHeight)
            .padding(horizontal = 16.dp, vertical = 12.dp).clearAndSetSemantics { contentDescription = label },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ClassCrest(Reports.crestClass(row, viewer), ReportsLayout.crest)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(Reports.name(row), style = FsType.bodyStrong, color = FsColor.strong)
            Text(copy.rowMeta(row, now), style = FsType.meta, color = FsColor.muted)
            if (stacked) Row(Modifier.padding(top = 4.dp), horizontalArrangement = Arrangement.spacedBy(10.dp), verticalAlignment = Alignment.CenterVertically) {
                Marks(state, row, copy)
            }
        }
        if (!stacked) Marks(state, row, copy)
        Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(24.dp))
    }
}

@Composable
private fun Marks(state: ReportState, row: ReportRow, copy: ReportsCopy) {
    copy.stateTag(state)?.let { StateTag(state, it) }
    VisibilityGlyph(Visibility.of(row.visibility))
}

@Composable
private fun StateTag(state: ReportState, text: String) {
    when (state) {
        ReportState.Live -> LiveTag(text)
        ReportState.Processing -> Tag(text, FsColor.gold)
        ReportState.Failed -> Tag(text, FsColor.wipe)
        ReportState.Done -> Unit
    }
}

/** The live tag: the kill-green pill led by its 7 dp dot (spec F10.1). */
@Composable
private fun LiveTag(text: String) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Row(
        Modifier.heightIn(min = 22.dp).pillSurface(FsColor.kill.copy(alpha = TAG_FILL), FsColor.kill.copy(alpha = TAG_LINE), shape)
            .padding(horizontal = 8.dp, vertical = 4.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(5.dp),
    ) {
        Box(Modifier.size(7.dp).background(FsColor.kill, CircleShape))
        Text(text.uppercase(Locale.ROOT), style = FsType.micro.copy(letterSpacing = 0.08.em), color = FsColor.kill)
    }
}

/** The shared Tag's ground and border alphas. */
private const val TAG_FILL = 0.12f
private const val TAG_LINE = 0.35f

/** F10.7: four rows at their real height. */
@Composable
private fun Skeleton() {
    Column(Modifier.padding(horizontal = 16.dp).padding(top = 8.dp).clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(stringResource(R.string.reports_thisWeek))
        Column(Modifier.fillMaxWidth().quietCard()) {
            repeat(ReportsLayout.SKELETON_ROWS) { index ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                Row(
                    Modifier.fillMaxWidth().height(ReportsLayout.rowHeight).padding(horizontal = 16.dp),
                    verticalAlignment = Alignment.CenterVertically,
                    horizontalArrangement = Arrangement.spacedBy(12.dp),
                ) {
                    SkeletonDot(ReportsLayout.crest)
                    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) { SkeletonBar(180.dp, 14.dp); SkeletonBar(120.dp, 11.dp) }
                }
            }
        }
    }
}
