package gg.foreversixty.app.bis.spec

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clipToBounds
import androidx.compose.ui.geometry.Rect
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.TextLayoutResult
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.IntOffset
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.items.GameIconImage
import kotlin.math.roundToInt

/**
 * The Change panel (`details.bar-m .stack`): spec, level, faction, profile, the race the list was simmed
 * as, and the sources. Closed it keeps its place in the bar's group at no height, as the page's closed
 * `details` does, and is hidden from accessibility.
 */
@Composable
fun BisSpecPanel(controls: BisSpecControls, open: Boolean, actions: BisSpecActions) {
    val closed = if (open) Modifier else Modifier.height(0.dp).clipToBounds().clearAndSetSemantics { }
    Box(Modifier.fillMaxWidth().then(closed)) {
        Column(
            Modifier.fillMaxWidth().wrapContentHeight(Alignment.Top, unbounded = true).padding(start = FsSpace.space4, end = FsSpace.space4, bottom = FsSpace.space3),
            verticalArrangement = Arrangement.spacedBy(FsSpace.space3),
        ) {
            Choices(stringResource(R.string.bis_picker_spec), controls.specs, PER_ROW_SPEC, actions)
            Choices(stringResource(R.string.bisSpec_level), controls.levels, controls.levels.size, actions)
            Choices(stringResource(R.string.bis_picker_faction), controls.factions, controls.factions.size, actions)
            if (controls.profiles.isNotEmpty()) Choices(stringResource(R.string.bisSpec_profile), controls.profiles, controls.profiles.size, actions)
            Labelled(stringResource(R.string.bisSpec_race)) {
                Field(Modifier.parityItem(ParityItemKind.Field)) { Text(controls.race, style = FIELD_TEXT, color = FsColor.strong) }
            }
            Labelled(stringResource(R.string.bisSpec_sources)) {
                Field(
                    Modifier
                        .parityItem(ParityItemKind.Field)
                        .parityItem(ParityItemKind.Button)
                        .clickable(role = Role.Button, onClick = actions.onSources),
                ) { Text(controls.sources, style = FIELD_TEXT.copy(fontSize = 13.sp), color = FsColor.strong) }
            }
        }
    }
}

@Composable
private fun Choices(label: String, options: List<BisSpecOption>, perRow: Int, actions: BisSpecActions) {
    Labelled(label) {
        BisSpecSegment(options, height = CHOICE_HEIGHT, fill = true, perRow = perRow) { index -> options[index].pick?.let(actions.onPick) }
    }
}

@Composable
private fun Labelled(label: String, body: @Composable () -> Unit) {
    Column {
        Text(label, style = LABEL, color = FsColor.muted, modifier = Modifier.padding(bottom = FsSpace.space1))
        body()
    }
}

@Composable
private fun Field(modifier: Modifier, content: @Composable () -> Unit) {
    val shape = RoundedCornerShape(FsRadius.control)
    Box(
        modifier.fillMaxWidth().heightIn(min = FIELD_HEIGHT).border(1.dp, FsColor.line, shape).padding(horizontal = FsSpace.space3),
        contentAlignment = Alignment.CenterStart,
    ) { content() }
}

/**
 * A segmented control (`.seg`): options [height] tall in a 1 px frame, the selected one gold on its wash.
 * [fill] shares each row among [perRow] options; otherwise the options keep their own widths.
 */
@OptIn(ExperimentalLayoutApi::class)
@Composable
fun BisSpecSegment(options: List<BisSpecOption>, height: Dp, fill: Boolean, perRow: Int = options.size, onPick: (Int) -> Unit) {
    val shape = RoundedCornerShape(FsRadius.control)
    val frame = Modifier.border(1.dp, FsColor.line, shape).padding(1.dp)
    if (fill) {
        FlowRow(frame.fillMaxWidth(), maxItemsInEachRow = perRow) {
            options.forEachIndexed { i, option -> Option(option, height, Modifier.weight(1f)) { onPick(i) } }
        }
    } else {
        Row(frame.width(IntrinsicSize.Max)) {
            options.forEachIndexed { i, option -> Option(option, height, Modifier) { onPick(i) } }
        }
    }
}

@Composable
private fun Option(option: BisSpecOption, height: Dp, modifier: Modifier, onClick: () -> Unit) {
    Row(
        modifier
            .height(height)
            .parityItem(ParityItemKind.Segment)
            .background(if (option.on) FsColor.selected else FsColor.bg.copy(alpha = 0f))
            .clickable(role = Role.RadioButton, onClick = onClick)
            .semantics { selected = option.on }
            .padding(horizontal = OPTION_PAD),
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space2, Alignment.CenterHorizontally),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        option.icon?.let { icon ->
            Box(Modifier.size(SPEC_ICON).border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.control))) { GameIconImage(icon, null, SPEC_ICON) }
        }
        Text(option.label, style = OPTION_TEXT, color = if (option.on) FsColor.gold else FsColor.text, maxLines = 1)
    }
}

/**
 * Text that ends with an inline button (`Changes from 20 to 29: see 30 to 39 ›`): the button's words in
 * gold after [text], with a tappable button laid over exactly where they fall.
 */
@Composable
fun InlineButtonText(text: AnnotatedString, button: String?, style: TextStyle, onClick: () -> Unit) {
    var layout by remember { mutableStateOf<TextLayoutResult?>(null) }
    val full = if (button == null) {
        text
    } else {
        buildAnnotatedString {
            append(text)
            withStyle(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold)) { append(button) }
        }
    }
    Box {
        Text(full, style = style, color = FsColor.text, onTextLayout = { layout = it })
        val bounds = button?.let { layout?.let { l -> rangeBounds(l, text.length, full.length) } }
        if (bounds != null) {
            val density = LocalDensity.current
            Box(
                Modifier
                    .offset { IntOffset(bounds.left.roundToInt(), bounds.top.roundToInt()) }
                    .size(with(density) { bounds.width.toDp() }, with(density) { bounds.height.toDp() })
                    .parityItem(ParityItemKind.Button)
                    .clickable(role = Role.Button, onClick = onClick),
            )
        }
    }
}

private fun rangeBounds(layout: TextLayoutResult, start: Int, end: Int): Rect? {
    if (end <= start || end > layout.layoutInput.text.length) return null
    return layout.getPathForRange(start, end).getBounds()
}

private const val PER_ROW_SPEC = 2
private val CHOICE_HEIGHT = 40.dp
private val FIELD_HEIGHT = 44.dp
private val OPTION_PAD = 14.dp
private val SPEC_ICON = 22.dp
private val LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val OPTION_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 16.sp)
private val FIELD_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 16.sp)
