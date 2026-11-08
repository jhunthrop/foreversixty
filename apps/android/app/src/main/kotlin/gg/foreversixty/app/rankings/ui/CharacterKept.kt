package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Path
import androidx.compose.ui.graphics.StrokeCap
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPlate
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.rankings.logic.CharacterPageView
import gg.foreversixty.app.rankings.logic.RecentRow
import gg.foreversixty.app.rankings.logic.ScoreBar
import gg.foreversixty.app.rankings.logic.ScoreView

// What the character page kept from before the restyle (lane 5 ruling, 2026-11-08): the score's trend
// and its five bars, the recent fights with their wipes, and the gear gap.

/** `As of 17 Dec` with the score's sparkline in the tier's colour, then the five parts as bars over their caps. */
@Composable
fun CharacterScoreTrend(score: ScoreView) {
    val tierColor = ItemQuality(score.tierQuality).text
    if (score.asOf != null || score.sparkline.size >= 2) {
        Row(Modifier.fillMaxWidth().padding(top = FsSpace.space3), verticalAlignment = Alignment.CenterVertically) {
            Box(Modifier.weight(1f)) { score.asOf?.let { FlowNote(it) } }
            Sparkline(score.sparkline, tierColor)
        }
    }
    if (score.bars.isNotEmpty()) {
        Column(Modifier.padding(top = FsSpace.space3).clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
            score.bars.forEach { ComponentBar(it) }
        }
    }
}

@Composable
private fun ComponentBar(bar: ScoreBar) {
    Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
        Text(stringResource(bar.component.label), style = FsType.small, color = FsColor.muted, modifier = Modifier.width(BAR_LABEL))
        Box(Modifier.weight(1f).height(BAR).background(FsColor.lineSoft, CircleShape)) {
            val fraction = if (bar.cap > 0) (bar.value.toFloat() / bar.cap).coerceIn(0f, 1f) else 0f
            Box(Modifier.fillMaxWidth(fraction).height(BAR).background(FsColor.goldHover, CircleShape))
        }
        Text(bar.value.toString(), style = FsType.small.copy(fontWeight = FontWeight.Bold), color = FsColor.strong, modifier = Modifier.width(BAR_VALUE), textAlign = TextAlign.End)
    }
}

/** `score_history[]`, oldest first. */
@Composable
private fun Sparkline(points: List<Float>, color: Color) {
    if (points.size < 2) return
    Canvas(Modifier.width(SPARK_WIDTH).height(SPARK_HEIGHT)) {
        val step = size.width / (points.size - 1)
        val path = Path()
        points.forEachIndexed { i, p ->
            val at = Offset(i * step, size.height * (1f - p))
            if (i == 0) path.moveTo(at.x, at.y) else path.lineTo(at.x, at.y)
        }
        drawPath(path, color, style = Stroke(width = 2.dp.toPx(), cap = StrokeCap.Round))
    }
}

/** Recent fights, kills and wipes, each opening its report when reports are on. */
@Composable
fun CharacterRecentFights(view: CharacterPageView, reportsOn: Boolean, openFight: (RecentRow) -> Unit) {
    if (view.recent.isEmpty()) return
    FlowSection(stringResource(R.string.char_recent), top = SECTION_GAP) {
        FlowGap(FsSpace.space3)
        view.recent.forEachIndexed { i, row ->
            FlowRow(
                row.boss,
                row.sub,
                lead = FlowLead.Check(ok = row.kill, miss = !row.kill),
                last = i == view.recent.lastIndex,
                trail = { Figure(row.value, row.unit) },
                onClick = if (reportsOn) ({ openFight(row) }) else null,
            )
        }
    }
}

@Composable
private fun Figure(value: String, unit: String) {
    Text(
        buildAnnotatedString {
            append(value)
            withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" $unit") }
        },
        style = FsType.figureSm,
        color = FsColor.strong,
    )
}

/** `Gear gap  36.0 points` under the gear: what the band's picks would add to the Gear part. */
@Composable
fun CharacterGearGap(gap: String) {
    FlowGap(FsSpace.space3)
    FlowPlate {
        Row(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4, vertical = GAP_PAD), verticalAlignment = Alignment.CenterVertically) {
            Text(stringResource(R.string.char_gap), style = FsType.item, color = FsColor.strong, modifier = Modifier.weight(1f))
            Text(gap, style = FsType.figureSm, color = FsColor.strong)
        }
    }
}

private val BAR = 8.dp
private val BAR_LABEL = 96.dp
private val BAR_VALUE = 36.dp
private val SPARK_WIDTH = 120.dp
private val SPARK_HEIGHT = 32.dp
private val GAP_PAD = 14.dp
