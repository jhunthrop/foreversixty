package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.widthIn
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.SegmentedButton
import androidx.compose.material3.SegmentedButtonDefaults
import androidx.compose.material3.SingleChoiceSegmentedButtonRow
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.reports.logic.MeterRow
import gg.foreversixty.app.reports.logic.Meters
import gg.foreversixty.app.reports.logic.MetersPhase
import gg.foreversixty.app.reports.logic.MetersTarget
import gg.foreversixty.app.reports.logic.Metric
import gg.foreversixty.app.reports.logic.Parses
import gg.foreversixty.app.reports.logic.ReportsCopy
import java.time.ZoneId
import java.util.Locale

/** What the sheet draws: its target, segment, answer and whether Show all was pressed. */
data class MetersSheetState(val target: MetersTarget, val metric: Metric, val phase: MetersPhase, val showAll: Boolean)

/** What the sheet hands back. */
data class MetersActions(val select: (Metric) -> Unit = {}, val showAll: () -> Unit = {}, val retry: () -> Unit = {})

object MeterLayout {
    val rowHeight = 44.dp
    val icon = 28.dp
    val goldEdge = 3.dp
    val barGap = 3.dp
    const val BAR_ALPHA = 0.22f
}

/** F10.3: the fight sheet at its large detent. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun MetersSheet(state: MetersSheetState, specs: SpecCatalog, build: String?, actions: MetersActions, onDismiss: () -> Unit) {
    ModalBottomSheet(
        onDismissRequest = onDismiss,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
        containerColor = GameColors.sheet,
    ) {
        Box(Modifier.fillMaxHeight()) { MetersSheetContent(state, specs, build, actions) }
    }
}

/** The sheet's body, apart from the sheet so screenshot tests draw it. */
@Composable
fun MetersSheetContent(state: MetersSheetState, specs: SpecCatalog, build: String?, actions: MetersActions, zone: ZoneId = ZoneId.systemDefault()) {
    val copy = ReportsCopy(LocalResources.current, zone)
    Column(Modifier.fillMaxWidth().padding(bottom = 24.dp)) {
        Column(Modifier.padding(horizontal = 16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
            Text(
                copy.metersTitleFor(state.target, state.metric).uppercase(Locale.ROOT), style = FsType.sec, color = FsColor.strong,
                modifier = Modifier.semantics { heading() },
            )
            Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(10.dp)) {
                val fight = (state.target as? MetersTarget.Pull)?.fight
                OutcomeDot(if (fight == null) FsColor.muted else if (fight.kill) FsColor.kill else FsColor.wipe)
                copy.headName(state.target)?.let { Text(it, style = FsType.bodyStrong, color = FsColor.strong) }
                Text(copy.metersHead(state.target), style = FsType.meta, color = FsColor.muted)
            }
            Segments(state.metric, copy, actions.select)
        }
        Box(Modifier.padding(top = 12.dp)) {
            when (val phase = state.phase) {
                MetersPhase.Loading -> Column(Modifier.clearAndSetSemantics {}) { repeat(Meters.SKELETON_ROWS) { SkeletonMeterRow() } }
                MetersPhase.Failed -> ErrorRegion(stringResource(R.string.meters_failed), actions.retry)
                MetersPhase.Flagged -> Message(stringResource(R.string.meters_flagged))
                is MetersPhase.Loaded -> if (phase.meters.players.isEmpty()) {
                    Message(stringResource(R.string.meters_empty))
                } else {
                    Rows(Meters.rows(phase.meters.players, state.metric), state, specs, build, copy, actions)
                }
            }
        }
    }
}

@Composable
private fun Segments(selected: Metric, copy: ReportsCopy, select: (Metric) -> Unit) {
    SingleChoiceSegmentedButtonRow(Modifier.fillMaxWidth()) {
        Metric.entries.forEachIndexed { index, metric ->
            SegmentedButton(
                selected = metric == selected,
                onClick = { if (metric != selected) select(metric) },
                shape = SegmentedButtonDefaults.itemShape(index, Metric.entries.size),
                modifier = Modifier.heightIn(min = 48.dp),
                icon = {},
                colors = SegmentedButtonDefaults.colors(
                    activeContainerColor = GameColors.pressedWash,
                    activeContentColor = FsColor.strong,
                    activeBorderColor = FsColor.line,
                    inactiveContainerColor = GameColors.controlFill,
                    inactiveContentColor = FsColor.muted,
                    inactiveBorderColor = FsColor.line,
                ),
            ) {
                Text(copy.segment(metric), style = FsType.segment, maxLines = 2, overflow = TextOverflow.Clip)
            }
        }
    }
}

@Composable
private fun Rows(rows: List<MeterRow>, state: MetersSheetState, specs: SpecCatalog, build: String?, copy: ReportsCopy, actions: MetersActions) {
    val scale = LocalDensity.current.fontScale
    val cut = Meters.cut(scale)
    val stacked = scale >= Meters.LARGE_SCALE
    Column {
        Meters.visible(rows, cut, state.showAll).forEach { row -> MeterRowView(row, state.metric, specs, build, copy, stacked) }
        if (Meters.hasMore(rows, cut, state.showAll)) {
            HorizontalDivider(color = FsColor.lineSoft)
            Text(
                stringResource(R.string.meters_showAll, rows.size), style = FsType.bodyStrong, color = FsColor.gold,
                modifier = Modifier.fillMaxWidth().heightIn(min = 48.dp).clickable(onClick = actions.showAll).padding(horizontal = 16.dp, vertical = 14.dp),
            )
        }
    }
}

/** One meter row (spec §4): the bar behind, rank, spec icon, class-coloured name, the figure and its unit. */
@Composable
private fun MeterRowView(row: MeterRow, metric: Metric, specs: SpecCatalog, build: String?, copy: ReportsCopy, stacked: Boolean) {
    val player = row.player
    val color = classColor(player.className)
    val label = copy.meterLabel(row)
    // Your row: the --row ground and a 3 dp gold left edge.
    val ground = if (player.mine) {
        Modifier.background(GameColors.rowWash).drawBehind { drawRect(FsColor.gold, size = Size(MeterLayout.goldEdge.toPx(), size.height)) }
    } else {
        Modifier
    }
    Box(Modifier.fillMaxWidth().heightIn(min = MeterLayout.rowHeight).then(ground).clearAndSetSemantics { contentDescription = label }) {
        Box(Modifier.matchParentSize().padding(vertical = MeterLayout.barGap)) {
            Box(Modifier.fillMaxWidth(row.fraction).fillMaxHeight().background(color.copy(alpha = MeterLayout.BAR_ALPHA)))
        }
        Row(
            Modifier.fillMaxWidth().heightIn(min = MeterLayout.rowHeight).padding(horizontal = 16.dp, vertical = 6.dp),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(10.dp),
        ) {
            Text(row.rank.toString(), style = FsType.monoMeta, color = FsColor.muted, modifier = Modifier.widthIn(min = 20.dp))
            PlayerMark(player.spec, player.className, specs, build)
            if (stacked) {
                Column(Modifier.weight(1f)) {
                    Text(player.name, style = FsType.bodyStrong, color = color)
                    Figure(row, metric, copy)
                }
            } else {
                Text(player.name, style = FsType.bodyStrong, color = color, modifier = Modifier.weight(1f), maxLines = 1, overflow = TextOverflow.Clip)
                Figure(row, metric, copy)
            }
        }
    }
}

@Composable
private fun Figure(row: MeterRow, metric: Metric, copy: ReportsCopy) {
    Row(verticalAlignment = Alignment.Bottom, horizontalArrangement = Arrangement.spacedBy(6.dp)) {
        Text(Parses.figure(row.value), style = FsType.mono, color = FsColor.strong)
        Text(copy.unit(metric.unit), style = FsType.micro, color = FsColor.muted, modifier = Modifier.padding(bottom = 2.dp))
    }
}

/** The spec icon, or the class crest when the log named no spec the data knows. */
@Composable
private fun PlayerMark(spec: String?, className: String, specs: SpecCatalog, build: String?) {
    val icon = spec?.let(specs::spec)?.icon
    if (icon != null && build != null) SpecIcon(icon, build, MeterLayout.icon) else ClassCrest(className, MeterLayout.icon)
}

@Composable
private fun SkeletonMeterRow() {
    Row(
        Modifier.fillMaxWidth().height(MeterLayout.rowHeight).padding(horizontal = 16.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(10.dp),
    ) {
        SkeletonBar(20.dp, 12.dp)
        SkeletonBar(MeterLayout.icon, MeterLayout.icon)
        SkeletonBar(140.dp, 14.dp)
    }
}

@Composable
private fun Message(text: String) {
    Text(text, style = FsType.body, color = FsColor.text, modifier = Modifier.padding(16.dp))
}
