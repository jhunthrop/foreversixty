package gg.foreversixty.app.flows

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType

/** One figure (the mock's Figure): its label over the value, the unit small beside it. */
data class FlowFigure(val label: String, val value: String, val unit: String? = null)

/** Figures in two columns at phone width (`.figs`: 20 between rows, 24 between columns). */
@Composable
fun FlowFigures(figures: List<FlowFigure>, modifier: Modifier = Modifier) {
    Column(modifier.fillMaxWidth(), verticalArrangement = Arrangement.spacedBy(FsSpace.space5)) {
        figures.chunked(COLUMNS).forEach { pair ->
            Row(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space6)) {
                pair.forEach { Figure(it, Modifier.weight(1f)) }
                repeat(COLUMNS - pair.size) { Spacer(Modifier.weight(1f)) }
            }
        }
    }
}

@Composable
private fun Figure(figure: FlowFigure, modifier: Modifier) {
    Column(modifier, verticalArrangement = Arrangement.spacedBy(LABEL_GAP)) {
        Text(figure.label, style = LABEL, color = FsColor.mutedGround)
        Text(
            buildAnnotatedString {
                append(figure.value)
                figure.unit?.let { withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.mutedGround)) { append(" $it") } }
            },
            style = FsType.figure,
            color = FsColor.strong,
        )
    }
}

private const val COLUMNS = 2
private val LABEL_GAP = 2.dp
private val LABEL = FsType.label
