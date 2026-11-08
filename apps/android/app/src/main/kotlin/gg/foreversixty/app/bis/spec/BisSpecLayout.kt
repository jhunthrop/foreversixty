package gg.foreversixty.app.bis.spec

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.Layout
import androidx.compose.ui.layout.layout
import androidx.compose.ui.unit.Constraints
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import kotlin.math.roundToInt

/**
 * Rounds the element's height to whole dp, as the page's CSS boxes are: at 2.625 px per dp every line
 * and padding rounds up a little on its own.
 */
fun Modifier.wholeDp(): Modifier = layout { measurable, constraints ->
    val placeable = measurable.measure(constraints)
    val height = (placeable.height / density).roundToInt().dp.roundToPx()
    layout(placeable.width, height) { placeable.place(0, 0) }
}

/**
 * A column that stacks its children on a whole-dp grid: each child takes its height rounded to whole dp
 * and the offsets turn into pixels once, so a long page (seventeen cells, a weights grid) does not drift
 * down by a fraction per child the way pixel-rounded children do.
 */
@Composable
fun DpColumn(modifier: Modifier = Modifier, spacing: Dp = 0.dp, content: @Composable () -> Unit) {
    Layout(content, modifier) { measurables, constraints ->
        val placeables = measurables.map { it.measure(constraints.copy(minHeight = 0, maxHeight = Constraints.Infinity)) }
        val gap = spacing.value.roundToInt()
        val tops = placeables.runningFold(0) { top, placeable -> top + (placeable.height / density).roundToInt() + gap }
        val total = if (placeables.isEmpty()) 0 else tops.last() - gap
        val width = if (constraints.hasBoundedWidth) constraints.maxWidth else placeables.maxOfOrNull { it.width } ?: 0
        val height = (total * density).roundToInt().coerceIn(constraints.minHeight, constraints.maxHeight)
        layout(width.coerceAtLeast(constraints.minWidth), height) {
            placeables.forEachIndexed { i, placeable -> placeable.place(0, (tops[i] * density).roundToInt()) }
        }
    }
}
