package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.RowScope
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.layout
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion

/**
 * The web-derived screens' plate (the mock's `.plate`): the raised ground in a 1 dp line, radius 6, its
 * rows edge to edge ([FlowRow] with `inset = 16`); [padded] adds the 16 of `.plate.pad`.
 */
@Composable
fun FlowPlate(modifier: Modifier = Modifier, padded: Boolean = false, content: @Composable ColumnScope.() -> Unit) {
    val shape = RoundedCornerShape(FsRadius.plate)
    Column(
        modifier
            .fillMaxWidth()
            .parityRegion(ParityRegion.PLATE)
            .clip(shape)
            .background(FsColor.raised)
            .border(1.dp, FsColor.line, shape)
            .padding(PLATE_LINE)
            .padding(if (padded) FsSpace.space4 else 0.dp),
        content = content,
    )
}

/** A plate's message (`.empty`): its lines (14/20 muted), then its actions 12 under them, 16 apart. */
@Composable
fun FlowPlateMessage(lines: List<String>, actions: (@Composable RowScope.() -> Unit)? = null) {
    Column(Modifier.fillMaxWidth().padding(FsSpace.space4)) {
        lines.forEach { Text(it, style = PLATE_TEXT, color = FsColor.muted) }
        if (actions != null) {
            Row(
                Modifier.padding(top = FsSpace.space3),
                horizontalArrangement = Arrangement.spacedBy(FsSpace.space4),
                verticalAlignment = Alignment.CenterVertically,
                content = actions,
            )
        }
    }
}

/** A load that failed inside a plate (`.empty`, `.err`): the sentence and Try again beside it. */
@Composable
fun FlowPlateRetry(text: String, retryLabel: String, retry: () -> Unit) {
    FlowRow(
        Modifier.fillMaxWidth().padding(FsSpace.space4).heightIn(min = RETRY_ROW),
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space3),
        verticalArrangement = Arrangement.spacedBy(FsSpace.space2),
        itemVerticalAlignment = Alignment.CenterVertically,
    ) {
        Text(text, style = PLATE_TEXT.copy(lineHeight = 24.sp), color = FsColor.muted)
        FlowButton(retryLabel, retry)
    }
}

/** A line with its link (`.notice`): 15/24, the link gold after it; [color] is the sentence's. */
@Composable
fun FlowNotice(text: String?, link: String?, onLink: () -> Unit, modifier: Modifier = Modifier, color: Color = FsColor.text) {
    FlowRow(
        modifier.fillMaxWidth().padding(vertical = NOTICE_PAD),
        horizontalArrangement = Arrangement.spacedBy(NOTICE_GAP),
    ) {
        if (text != null) Text(text, style = NOTICE, color = color)
        if (link != null) FlowLink(link, onClick = onLink)
    }
}

/**
 * The controls bar (`.cbar.full-bleed`): edge to edge across the page, 62 tall between two soft lines;
 * its [content] (fields and segments) scrolls sideways and counts in `controls` wherever it scrolls.
 */
@Composable
fun FlowControls(content: @Composable RowScope.() -> Unit) {
    Column(Modifier.fillMaxWidth().flowBleed().parityRegion(ParityRegion.CONTROLS).background(FsColor.bg.copy(alpha = BAR_ALPHA))) {
        HorizontalDivider(color = FsColor.lineSoft)
        ParityScrollRegion(ParityRegion.CONTROLS) {
            Row(
                Modifier.fillMaxWidth().height(BAR_BODY).horizontalScroll(rememberScrollState()).padding(horizontal = FsSpace.space4),
                horizontalArrangement = Arrangement.spacedBy(FsSpace.space2),
                verticalAlignment = Alignment.CenterVertically,
                content = content,
            )
        }
        HorizontalDivider(color = FsColor.lineSoft)
    }
}

/** A select in the controls bar (`.sel`): its value and the open mark; [button] also counts it as a button. */
@Composable
fun FlowField(value: String, onClick: () -> Unit, button: Boolean = false) {
    val shape = RoundedCornerShape(FsRadius.field)
    Box(
        Modifier
            .height(FIELD)
            .parityItem(ParityItemKind.Field)
            .then(if (button) Modifier.parityItem(ParityItemKind.Button) else Modifier)
            .clip(shape)
            .border(1.dp, FsColor.line, shape)
            .clickable(role = Role.DropdownList, onClick = onClick)
            .padding(horizontal = FsSpace.space3),
        contentAlignment = Alignment.Center,
    ) {
        Text("$value$OPEN_MARK", style = FIELD_TEXT, color = FsColor.strong, maxLines = 1)
    }
}

/** A segmented choice (`.seg`): one bordered group, the selected option gold on its tint. */
@Composable
fun FlowSegments(options: List<String>, selected: Int, onSelect: (Int) -> Unit) {
    val shape = RoundedCornerShape(FsRadius.field)
    Row(Modifier.height(FIELD).clip(shape).border(1.dp, FsColor.line, shape)) {
        options.forEachIndexed { index, label ->
            val on = index == selected
            Box(
                Modifier
                    .height(FIELD)
                    .parityItem(ParityItemKind.Segment)
                    .background(if (on) FsColor.selected else Color.Transparent)
                    .clickable(role = Role.RadioButton) { onSelect(index) }
                    .semantics { this.selected = on }
                    .padding(horizontal = FsSpace.space3),
                contentAlignment = Alignment.Center,
            ) {
                Text(label, style = FIELD_TEXT, color = if (on) FsColor.gold else FsColor.text, maxLines = 1)
            }
        }
    }
}

/** Widens an element past the page's 16 dp gutters to the screen's edges (`.full-bleed`). */
fun Modifier.flowBleed(gutter: Dp = FsSpace.space4): Modifier = layout { measurable, constraints ->
    val extra = gutter.roundToPx() * 2
    val placeable = measurable.measure(constraints.copy(minWidth = constraints.minWidth + extra, maxWidth = constraints.maxWidth + extra))
    layout(placeable.width - extra, placeable.height) { placeable.place(-extra / 2, 0) }
}

private val PLATE_LINE = 1.dp
private val RETRY_ROW = 40.dp
private val NOTICE_PAD = 9.dp
private val NOTICE_GAP = 14.dp
private val BAR_BODY = 60.dp
private val FIELD = 40.dp
private const val BAR_ALPHA = 0.94f
private const val OPEN_MARK = " ▾"
private val PLATE_TEXT = FsType.small.copy(fontSize = 14.sp)
private val NOTICE = FsType.body.copy(lineHeight = 24.sp)
private val FIELD_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 16.sp)
