package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clipToBounds
import androidx.compose.ui.layout.layout
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion

/** One choice of a control bar group: `DPS` with its count `20`, or `Alliance` led by its [mark]. */
class FlowChoice(val label: String, val count: String? = null, val mark: (@Composable () -> Unit)? = null)

/** A labelled group of the bar's panel (`Role`, `Faction`): equal segments, [selected] the one on. */
class FlowChoiceGroup(val label: String, val choices: List<FlowChoice>, val selected: Int, val onSelect: (Int) -> Unit)

/**
 * The phone control bar (the web's ControlBar below 768, `details.bar-m`): one full-bleed line, the
 * choice in words and `Change ▾`, ruled above and below; the line opens the panel of [groups] in place.
 * Closed, the panel is still laid out at no height (the web's closed `details` keeps its content), so
 * its segments count in the `controls` region. [inset] is the line's padding above and below; [end]
 * sits after `Change ▾` (a share button).
 */
@Composable
fun FlowControlBar(summary: String, groups: List<FlowChoiceGroup>, inset: Dp = 0.dp, end: (@Composable () -> Unit)? = null) {
    var open by rememberSaveable { mutableStateOf(false) }
    Column(Modifier.fullBleed(FsSpace.space4).parityRegion(ParityRegion.CONTROLS)) {
        HorizontalDivider(color = FsColor.lineSoft)
        Column(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4, vertical = inset)) {
            Row(
                Modifier
                    .fillMaxWidth()
                    .heightIn(min = LINE)
                    .clickable(role = Role.Button) { open = !open }
                    .semantics { selected = open },
                verticalAlignment = Alignment.CenterVertically,
            ) {
                Text(summary, style = SUMMARY, color = FsColor.strong, modifier = Modifier.weight(1f))
                Text(stringResource(R.string.flows_change), style = CHANGE, color = FsColor.gold)
                end?.invoke()
            }
            Box(Modifier.collapsed(open)) {
                ParityScrollRegion(ParityRegion.CONTROLS) {
                    Column(Modifier.padding(bottom = FsSpace.space3), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
                        groups.forEach { ChoiceGroup(it) }
                    }
                }
            }
        }
        HorizontalDivider(color = FsColor.lineSoft)
    }
}

@Composable
private fun ChoiceGroup(group: FlowChoiceGroup) {
    Column {
        FlowLabel(group.label, Modifier.padding(bottom = FsSpace.space1))
        Row(Modifier.fillMaxWidth().border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.control))) {
            group.choices.forEachIndexed { index, choice ->
                val on = index == group.selected
                Row(
                    Modifier
                        .weight(1f)
                        .heightIn(min = SEGMENT)
                        .parityItem(ParityItemKind.Segment)
                        .background(if (on) FsColor.selected else FsColor.raised)
                        .clickable(role = Role.RadioButton) { if (!on) group.onSelect(index) }
                        .semantics { selected = on },
                    horizontalArrangement = Arrangement.Center,
                    verticalAlignment = Alignment.CenterVertically,
                ) {
                    choice.mark?.let { mark ->
                        Box(Modifier.parityItem(ParityItemKind.Image)) { mark() }
                        Spacer(Modifier.width(FsSpace.space2))
                    }
                    Text(choice.label, style = SEGMENT_TEXT, color = if (on) FsColor.gold else FsColor.strong)
                    choice.count?.let { Text(it, style = COUNT, color = FsColor.muted, modifier = Modifier.padding(start = 6.dp)) }
                }
            }
        }
    }
}

/** A small bold label (`.t-label`, 12/16 700, muted): `Role`, `Your spec · DPS`, `Sources`. */
@Composable
fun FlowLabel(text: String, modifier: Modifier = Modifier) {
    Text(text, style = LABEL, color = FsColor.muted, modifier = modifier)
}

/** A tag (`.pill.tag`, 20 tall, 12/16 600): a spec's role in a list, `Your spec` on a tier row. */
@Composable
fun FlowTag(text: String, color: androidx.compose.ui.graphics.Color = FsColor.muted) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(Modifier.heightIn(min = TAG).border(1.dp, color.copy(alpha = TAG_LINE), shape).padding(horizontal = 6.dp), contentAlignment = Alignment.Center) {
        Text(text, style = TAG_TEXT, color = color)
    }
}

/** Widens the element by [gutter] on each side, past the page's side padding (`.full-bleed`). */
fun Modifier.fullBleed(gutter: Dp): Modifier = layout { measurable, constraints ->
    val extra = gutter.roundToPx() * 2
    val width = constraints.maxWidth + extra
    val placeable = measurable.measure(constraints.copy(minWidth = width, maxWidth = width))
    layout(constraints.maxWidth, placeable.height) { placeable.place(-extra / 2, 0) }
}

/** Closed: laid out and placed at no height, clipped and silent; open: as tall as its content. */
private fun Modifier.collapsed(open: Boolean): Modifier =
    clipToBounds()
        .then(if (open) Modifier else Modifier.clearAndSetSemantics {})
        .layout { measurable, constraints ->
            val placeable = measurable.measure(constraints.copy(minHeight = 0))
            layout(placeable.width, if (open) placeable.height else 0) { placeable.place(0, 0) }
        }

private val LINE = 48.dp
private val SEGMENT = 44.dp
private val TAG = 20.dp
private const val TAG_LINE = 0.4f
private val SUMMARY = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val CHANGE = FsType.bodyStrong.copy(fontSize = 13.sp, lineHeight = 20.sp)
private val SEGMENT_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 16.sp)
private val COUNT = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
private val LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val TAG_TEXT = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
