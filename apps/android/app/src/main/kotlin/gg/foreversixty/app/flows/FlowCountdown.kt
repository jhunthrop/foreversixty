package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.layout.layout
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.unit.TextUnit
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType

/** One part of a countdown clock: `02` over `hours`. */
data class ClockPart(val value: String, val unit: String)

/**
 * The raid's countdown panel (the mock's Countdown, tone guild): the eyebrow, the clock read aloud as
 * one timer ([clockLabel]), and the aside under it. [live] puts the green dot before the eyebrow.
 */
@Composable
fun FlowCountdown(eyebrow: String, clock: List<ClockPart>, clockLabel: String, live: Boolean = false, aside: @Composable () -> Unit) {
    val shape = RoundedCornerShape(PANEL_RADIUS)
    Column(
        Modifier
            .fillMaxWidth()
            .clip(shape)
            .background(FsColor.raised)
            .drawBehind {
                drawRect(Brush.radialGradient(listOf(FsColor.gold.copy(alpha = GLOW), Color.Transparent), center = Offset(size.width, 0f), radius = GLOW_RADIUS.toPx()))
            }
            .border(1.dp, FsColor.line, shape)
            .padding(horizontal = FsSpace.space6, vertical = FsSpace.space5),
    ) {
        Row(verticalAlignment = Alignment.CenterVertically) {
            if (live) {
                Box(Modifier.size(LIVE_DOT).background(FsColor.kill, CircleShape))
                Spacer(Modifier.width(FsSpace.space2))
            }
            Text(eyebrow, style = EYEBROW, color = FsColor.muted)
        }
        Column(Modifier.padding(top = EYEBROW_GAP)) {
            if (clock.isNotEmpty()) {
                Row(Modifier.semantics(mergeDescendants = true) {}.clearAndSetSemantics { contentDescription = clockLabel }, horizontalArrangement = Arrangement.spacedBy(CLOCK_GAP)) {
                    clock.forEach { part ->
                        Column(horizontalAlignment = Alignment.CenterHorizontally) {
                            Text(part.value, style = CLOCK, color = FsColor.strong, modifier = Modifier.lineBox(CLOCK.lineHeight))
                            Text(part.unit, style = UNIT, color = FsColor.muted)
                        }
                    }
                }
                Spacer(Modifier.height(FsSpace.space4))
            }
            aside()
        }
    }
}

/**
 * A one-line text held to its line height: a figure whose glyphs are taller than the role's line keeps
 * the line's box (the mock's 40/42 clock), centred, so the panel does not grow.
 */
fun Modifier.lineBox(lineHeight: TextUnit): Modifier = layout { measurable, constraints ->
    val placeable = measurable.measure(constraints)
    val box = minOf(placeable.height, lineHeight.roundToPx())
    layout(placeable.width, box) { placeable.place(0, (box - placeable.height) / 2) }
}

/**
 * Who is coming (the mock's Tally): three roles, 16 apart, each its count coloured by how close it is.
 * At large text the roles go one per line, full width, so no role name is clipped.
 */
@Composable
fun FlowTally(items: List<TallyItem>, modifier: Modifier = Modifier) {
    if (LocalDensity.current.fontScale >= STACK_SCALE) {
        Column(modifier.fillMaxWidth(), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
            items.forEach { TallyRole(it, Modifier.fillMaxWidth()) }
        }
    } else {
        Row(modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space4)) {
            items.forEach { TallyRole(it, Modifier.weight(1f)) }
        }
    }
}

@Composable
private fun TallyRole(item: TallyItem, modifier: Modifier) {
    val tone = item.tone.color
    val want = item.want ?: item.have
    Column(modifier.semantics(mergeDescendants = true) {}) {
        Row(verticalAlignment = Alignment.CenterVertically) {
            Text(item.label, style = FsType.item, color = FsColor.strong, modifier = Modifier.weight(1f), maxLines = 1)
            Text("${item.have}/$want", style = FsType.item, color = tone)
        }
        val share = if (want > 0) (item.have.toFloat() / want).coerceIn(0f, 1f) else 1f
        Box(Modifier.padding(top = METER_GAP).fillMaxWidth().height(METER).clip(RoundedCornerShape(METER_RADIUS)).background(FsColor.lineSoft)) {
            Box(Modifier.fillMaxWidth(share).height(METER).clip(RoundedCornerShape(METER_RADIUS)).background(if (item.tone == TallyTone.Met) FsColor.kill else FsColor.gold))
        }
    }
}

private val PANEL_RADIUS = 8.dp
private val LIVE_DOT = 8.dp
private val EYEBROW_GAP = 10.dp
private val CLOCK_GAP = 14.dp
private val METER = 8.dp
private val METER_GAP = 6.dp
private val METER_RADIUS = 3.dp
private const val GLOW = 0.1f
private val GLOW_RADIUS = 300.dp
private const val STACK_SCALE = 1.3f
private val EYEBROW = FsType.tag
private val CLOCK = FsType.countdown
private val UNIT = FsType.figureUnit
