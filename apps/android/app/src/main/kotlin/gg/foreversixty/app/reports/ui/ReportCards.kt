package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.onClick
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.reports.logic.FightRow
import gg.foreversixty.app.reports.logic.FightRows
import gg.foreversixty.app.reports.logic.MetersTarget
import gg.foreversixty.app.reports.logic.ParseRow
import gg.foreversixty.app.reports.logic.ParseTier
import gg.foreversixty.app.reports.logic.Parses
import gg.foreversixty.app.reports.logic.ParsesPhase
import gg.foreversixty.app.reports.logic.ReportsCopy

object ReportCardLayout {
    val parseRowHeight = 72.dp
    val fightRowHeight = 64.dp
    val crest = 36.dp
    val figureSize = 20.sp

    /** From this font scale the parse figure drops under the boss name. */
    const val STACK_SCALE = 1.3f
}

/** `Your parses`, the report's one primary card (spec F10.2); hidden when nobody is signed in. */
@Composable
fun ParsesCard(phase: ParsesPhase, copy: ReportsCopy, retry: () -> Unit) {
    if (phase == ParsesPhase.Hidden) return
    Column(Modifier.fillMaxWidth().primaryCard().padding(vertical = 14.dp)) {
        MicroLabel(stringResource(R.string.report_yourParses), Modifier.padding(horizontal = 16.dp), color = FsColor.gold)
        when (phase) {
            ParsesPhase.Hidden -> Unit
            ParsesPhase.Loading -> Row(Modifier.padding(16.dp), horizontalArrangement = Arrangement.spacedBy(12.dp)) {
                SkeletonDot(ReportCardLayout.crest)
                Column(verticalArrangement = Arrangement.spacedBy(8.dp)) { SkeletonBar(140.dp, 15.dp); SkeletonBar(200.dp, 12.dp) }
            }
            ParsesPhase.Failed -> ErrorRegion(stringResource(R.string.report_parsesFailed), retry)
            is ParsesPhase.Loaded -> {
                val rows = Parses.rows(phase.me)
                if (rows.isEmpty()) {
                    Text(stringResource(R.string.report_noKills), style = FsType.body, color = FsColor.text, modifier = Modifier.padding(horizontal = 16.dp, vertical = 10.dp))
                }
                rows.forEachIndexed { index, row ->
                    if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                    ParseRowView(row, copy)
                }
            }
        }
    }
}

@Composable
private fun ParseRowView(row: ParseRow, copy: ReportsCopy) {
    val label = copy.parseLabel(row)
    val stacked = LocalDensity.current.fontScale >= ReportCardLayout.STACK_SCALE
    Row(
        Modifier.fillMaxWidth().heightIn(min = ReportCardLayout.parseRowHeight).padding(horizontal = 16.dp, vertical = 10.dp)
            .clearAndSetSemantics { contentDescription = label },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ClassCrest(row.className, ReportCardLayout.crest)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(row.fight.encounter, style = FsType.bodyStrong.copy(fontWeight = FontWeight.Bold), color = FsColor.strong)
            Text(copy.parseMeta(row), style = FsType.meta, color = FsColor.muted)
            if (row.fight.newBest) Row(Modifier.padding(top = 4.dp)) { Tag(stringResource(R.string.report_newBest), FsColor.gold) }
            if (stacked) Figure(row, copy, Alignment.Start)
        }
        if (!stacked) Figure(row, copy, Alignment.End)
    }
}

/** `712 DPS` in mono 20, the parse number in its ladder colour and `percentile` under it. */
@Composable
private fun Figure(row: ParseRow, copy: ReportsCopy, alignment: Alignment.Horizontal) {
    val fight = row.fight
    Column(horizontalAlignment = alignment) {
        Text(
            buildAnnotatedString {
                append(Parses.figure(fight.value))
                withStyle(SpanStyle(fontSize = FsType.meta.fontSize, color = FsColor.muted)) { append(" " + copy.unit(fight.unit)) }
            },
            style = FsType.mono.copy(fontSize = ReportCardLayout.figureSize), color = FsColor.strong,
        )
        fight.percentile?.let { percentile ->
            Text(
                buildAnnotatedString {
                    withStyle(SpanStyle(color = parseColor(ParseTier.of(percentile)), fontFamily = FsType.mono.fontFamily, fontWeight = FontWeight.Bold)) {
                        append(percentile.toString())
                    }
                    append(" " + stringResource(R.string.report_percentile))
                },
                style = FsType.meta, color = FsColor.muted,
            )
        }
    }
}

/** `Fights`: every boss pull, each zone's trash as one row, and the boss-pull count (spec F10.2). */
@Composable
fun FightsCard(report: Report, copy: ReportsCopy, open: (MetersTarget) -> Unit) {
    val rows = FightRows.of(report)
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        Row(verticalAlignment = Alignment.CenterVertically) {
            MicroLabel(stringResource(R.string.report_fights), Modifier.weight(1f))
            MicroLabel(copy.bossPulls(FightRows.bossPulls(report)))
        }
        if (rows.isEmpty()) return@Column
        Column(Modifier.fillMaxWidth().quietCard()) {
            rows.forEachIndexed { index, row ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                when (row) {
                    is FightRow.Boss -> FightRowView(
                        row.fight.name, copy.bossMeta(row), if (row.fight.kill) FsColor.kill else FsColor.wipe, copy.bossLabel(row),
                        if (row.inProgress) null else { { open(MetersTarget.Pull(row.fight)) } },
                    )
                    is FightRow.Trash -> FightRowView(
                        copy.trashTitle(row.group.zone), copy.trashMeta(row), FsColor.muted, copy.trashLabel(row),
                    ) { open(MetersTarget.Trash(row.group)) }
                    is FightRow.LoosePull -> FightRowView(copy.looseTitle(), copy.looseMeta(row), FsColor.muted, null) { open(MetersTarget.Pull(row.fight)) }
                }
            }
        }
    }
}

@Composable
private fun FightRowView(title: String, meta: String, dot: Color, label: String?, open: (() -> Unit)?) {
    val hint = stringResource(R.string.report_a11y_showsMeters)
    val words = label ?: "$title, $meta"
    val tap = if (open != null) Modifier.clickable(onClickLabel = hint, onClick = open) else Modifier
    Row(
        Modifier.fillMaxWidth().then(tap).heightIn(min = ReportCardLayout.fightRowHeight).padding(horizontal = 16.dp, vertical = 10.dp)
            .clearAndSetSemantics {
                contentDescription = words
                if (open != null) onClick(hint) { open(); true }
            },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(14.dp),
    ) {
        OutcomeDot(dot)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(title, style = FsType.bodyStrong, color = FsColor.strong)
            Text(meta, style = FsType.meta, color = FsColor.muted)
        }
        if (open != null) Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(24.dp))
    }
}
