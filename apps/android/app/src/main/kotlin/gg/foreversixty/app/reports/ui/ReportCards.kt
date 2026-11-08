package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.onClick
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowPlate
import gg.foreversixty.app.flows.FlowPlateMessage
import gg.foreversixty.app.flows.FlowPlateRetry
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowRule
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.SUB
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.reports.logic.FightRow
import gg.foreversixty.app.reports.logic.FightRows
import gg.foreversixty.app.reports.logic.MetersTarget
import gg.foreversixty.app.reports.logic.ParseRow
import gg.foreversixty.app.reports.logic.ParseTier
import gg.foreversixty.app.reports.logic.Parses
import gg.foreversixty.app.reports.logic.ParsesPhase
import gg.foreversixty.app.reports.logic.ReportsCopy

/**
 * The report's tab strip (`.strip-tabs`): Summary, this screen, and More, which opens the whole report
 * (meters, casts, deaths, the timeline) on foreversixty.gg.
 */
@Composable
fun ReportTabs(openWeb: () -> Unit) {
    Column(Modifier.fillMaxWidth()) {
        Row(Modifier.horizontalScroll(rememberScrollState())) {
            Tab(stringResource(R.string.report_summary), current = true, onClick = null)
            Tab(stringResource(R.string.report_more), current = false, onClick = openWeb)
        }
        HorizontalDivider(color = FsColor.line)
    }
}

@Composable
private fun Tab(label: String, current: Boolean, onClick: (() -> Unit)?) {
    val tap = onClick?.let { Modifier.parityItem(ParityItemKind.Button).clickable(role = Role.Button, onClick = it) } ?: Modifier.semantics { selected = current }
    // The current tab's gold underline sits on the strip's rule.
    val underline = if (current) {
        Modifier.drawBehind {
            val stroke = UNDERLINE.toPx()
            drawRect(FsColor.gold, topLeft = Offset(0f, size.height - stroke), size = Size(size.width, stroke))
        }
    } else {
        Modifier
    }
    Box(Modifier.height(TAB).then(underline).then(tap).padding(horizontal = FsSpace.space3), contentAlignment = Alignment.Center) {
        Text(if (onClick != null) "$label$MORE_MARK" else label, style = TAB_TEXT, color = if (current) FsColor.strong else FsColor.text)
    }
}

/** `Your parses` (spec F10.2): each kill of yours with its figure and parse; hidden signed out. */
@Composable
fun ParsesSection(phase: ParsesPhase, copy: ReportsCopy, retry: () -> Unit) {
    if (phase == ParsesPhase.Hidden) return
    FlowSection(stringResource(R.string.report_yourParses)) {
        FlowGap(FsSpace.space3)
        FlowPlate {
            when (phase) {
                ParsesPhase.Hidden -> Unit
                ParsesPhase.Loading -> Column(Modifier.padding(FsSpace.space4).clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
                    FlowSkeleton(SKELETON_TITLE, SKELETON_LINE)
                    FlowSkeleton(SKELETON_SUB, SKELETON_SMALL)
                }
                ParsesPhase.Failed -> FlowPlateRetry(stringResource(R.string.report_parsesFailed), stringResource(R.string.common_tryAgain), retry)
                is ParsesPhase.Loaded -> {
                    val rows = Parses.rows(phase.me)
                    if (rows.isEmpty()) FlowPlateMessage(listOf(stringResource(R.string.report_noKills)))
                    rows.forEachIndexed { index, row -> ParseRowView(row, copy, last = index == rows.lastIndex) }
                }
            }
        }
    }
}

@Composable
private fun ParseRowView(row: ParseRow, copy: ReportsCopy, last: Boolean) {
    val label = copy.parseLabel(row)
    Box(Modifier.clearAndSetSemantics { contentDescription = label }) { ParseRowBody(row, copy, last) }
}

@Composable
private fun ParseRowBody(row: ParseRow, copy: ReportsCopy, last: Boolean) {
    FlowRow(
        row.fight.encounter,
        copy.parseMeta(row),
        lead = FlowLead.Crest(row.className),
        last = last,
        inset = FsSpace.space4,
        trail = { ParseFigure(row, copy) },
    )
}

/** `712 DPS · 91 percentile`, the parse in its ladder colour, and `New best` when it is one. */
@Composable
private fun ParseFigure(row: ParseRow, copy: ReportsCopy) {
    val fight = row.fight
    val percentileWord = stringResource(R.string.report_percentile)
    Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space3), verticalAlignment = Alignment.CenterVertically) {
        Text(
            buildAnnotatedString {
                append(Parses.figure(fight.value))
                withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" " + copy.unit(fight.unit)) }
                fight.percentile?.let { percentile ->
                    withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.Normal, color = FsColor.muted)) { append(" · ") }
                    withStyle(SpanStyle(color = parseColor(ParseTier.of(percentile)))) { append(percentile.toString()) }
                    withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.Normal, color = FsColor.muted)) { append(" $percentileWord") }
                }
            },
            style = FsType.figureSm,
            color = FsColor.strong,
        )
        if (fight.newBest) Tag(stringResource(R.string.report_newBest), FsColor.gold)
    }
}

/** `Fights`: every boss pull, each zone's trash as one row, and the boss-pull count; a row opens its meters. */
@Composable
fun FightsSection(report: Report, copy: ReportsCopy, open: (MetersTarget) -> Unit) {
    val rows = FightRows.of(report)
    FlowSection(stringResource(R.string.report_fights), meta = copy.bossPulls(FightRows.bossPulls(report))) {
        if (rows.isEmpty()) return@FlowSection
        FlowGap(FsSpace.space3)
        FlowPlate {
            rows.forEachIndexed { index, row ->
                val last = index == rows.lastIndex
                when (row) {
                    is FightRow.Boss -> FightRowView(
                        copy.bossTitle(row.fight, row.pull), copy.bossMeta(row), if (row.fight.kill) FsColor.kill else FsColor.wipe, last, copy.bossLabel(row),
                        if (row.inProgress) null else { { open(MetersTarget.Pull(row.fight, row.pull)) } },
                    )
                    is FightRow.Trash -> FightRowView(copy.trashTitle(row.group.zone), copy.trashMeta(row), FsColor.muted, last, copy.trashLabel(row)) {
                        open(MetersTarget.Trash(row.group))
                    }
                    is FightRow.LoosePull -> FightRowView(copy.looseTitle(), copy.looseMeta(row), FsColor.muted, last, null) { open(MetersTarget.Pull(row.fight)) }
                }
            }
        }
    }
}

/** A fight: its name and outcome line, the line in the kill or wipe colour's dot ahead of it. */
@Composable
private fun FightRowView(title: String, meta: String, dot: Color, last: Boolean, label: String?, open: (() -> Unit)?) {
    val hint = stringResource(R.string.report_a11y_showsMeters)
    val words = label ?: "$title, $meta"
    Column(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .then(open?.let { Modifier.clickable(role = Role.Button, onClickLabel = hint, onClick = it) } ?: Modifier)
            .clearAndSetSemantics {
                contentDescription = words
                if (open != null) onClick(hint) { open(); true }
            }
            .padding(horizontal = FsSpace.space4, vertical = ROW_PAD),
    ) {
        Text(title, style = FsType.item, color = FsColor.strong)
        Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
            OutcomeDot(dot)
            Text(meta, style = SUB, color = FsColor.muted)
        }
    }
    if (!last) FlowRule()
}

private val TAB = 46.dp
private val UNDERLINE = 2.dp
private val ROW_PAD = 10.dp
private val SKELETON_TITLE = 160.dp
private val SKELETON_SUB = 220.dp
private val SKELETON_LINE = 16.dp
private val SKELETON_SMALL = 12.dp
private const val MORE_MARK = " ▾"
private val TAB_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 16.sp)
