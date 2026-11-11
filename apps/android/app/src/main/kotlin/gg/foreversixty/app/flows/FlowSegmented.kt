package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import gg.foreversixty.app.home.LARGE_TEXT_SCALE
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.parityItem

/**
 * Material 3's segmented buttons in the round-5 look (spec Mobile, Android): one 48 dp pill with a
 * rule between the segments, the selected one filled teal inside the teal edge, as iOS and the web
 * primary (a1). The raid's four answers. Each segment is a button in the parity inventory, as the
 * mock's answers are. From font scale 1.5 the four stack two by two, as iOS does at accessibility
 * sizes (a2). Not [enabled] (offline, locked) the pill dims to the quiet tokens, keeps the answer
 * marked and reads as disabled to TalkBack (R16).
 */
@Composable
fun FlowSegmented(labels: List<String>, selected: String?, modifier: Modifier = Modifier, enabled: Boolean = true, onPress: (Int) -> Unit) {
    val colors = ChoiceColors.of(enabled)
    val stacked = LocalDensity.current.fontScale >= LARGE_TEXT_SCALE && labels.size > STACK_COLUMNS
    // Round 4 row 4: the 2 x 2 stack at radius 4, as the web draws a grid; the one row stays the pill.
    val shape = if (stacked) RoundedCornerShape(FsRadius.control) else RoundedCornerShape(percent = 50)
    val rows = if (stacked) labels.withIndex().chunked(STACK_COLUMNS) else listOf(labels.withIndex().toList())
    Column(modifier.fillMaxWidth().clip(shape).border(1.dp, colors.edge, shape)) {
        rows.forEachIndexed { rowIndex, row ->
            if (rowIndex > 0) Box(Modifier.fillMaxWidth().height(1.dp).background(colors.edge))
            Row(Modifier.fillMaxWidth().height(IntrinsicSize.Min)) {
                row.forEachIndexed { at, (index, label) ->
                    if (at > 0) Box(Modifier.width(1.dp).fillMaxHeight().background(colors.edge))
                    Segment(label, label == selected, enabled, colors, Modifier.weight(1f)) { onPress(index) }
                }
            }
        }
    }
}

@Composable
private fun Segment(label: String, on: Boolean, enabled: Boolean, colors: ChoiceColors, modifier: Modifier, onPress: () -> Unit) {
    Box(
        modifier
            .height(FsTarget.min)
            .parityItem(ParityItemKind.Button)
            .background(if (on) colors.selected else Color.Transparent)
            .then(if (on) Modifier.border(1.dp, colors.selectedEdge) else Modifier)
            .clickable(enabled = enabled, role = Role.Button, onClick = onPress)
            .semantics { this.selected = on },
        contentAlignment = Alignment.Center,
    ) {
        Text(label, style = ScaledType.segment, color = colors.text, maxLines = 1)
    }
}

/** Two answers to a row when the four stack. */
private const val STACK_COLUMNS = 2

/**
 * A choice's outline, selected fill, the selected one's edge and label, from the tokens (a1: the
 * fill is teal inside the teal edge); the disabled set is the quiet one.
 */
internal data class ChoiceColors(val edge: Color, val selected: Color, val selectedEdge: Color, val text: Color) {
    companion object {
        private val ENABLED = ChoiceColors(FsColor.rule, FsColor.teal, FsColor.tealEdge, FsColor.strong)
        private val DISABLED = ChoiceColors(FsColor.line, FsColor.teal, FsColor.line, FsColor.muted)

        fun of(enabled: Boolean): ChoiceColors = if (enabled) ENABLED else DISABLED
    }
}
