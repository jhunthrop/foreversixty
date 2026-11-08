package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.layout
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.LineBreak
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.TextUnit
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.items.GameIconImage

/**
 * The flows screens' layout (design/mocks/ws_app_flows.py `.ap`, spec 2026-11-08): one scrolling column
 * under the profile bar, 8 above, 16 at the sides, 24 below. Home, Raid and Gear are built from these.
 * [top] is 0 for a page the web draws (signed-out Home, whose head starts at the bar).
 */
@Composable
fun FlowPage(modifier: Modifier = Modifier, top: Dp = FsSpace.space2, content: @Composable ColumnScope.() -> Unit) {
    Column(
        modifier
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
            .parityContent()
            .padding(start = FsSpace.space4, end = FsSpace.space4, top = top, bottom = FsSpace.space6),
        content = content,
    )
}

/**
 * The page head (the mock's PageHeader): an optional 56 crest, the title in its colour and the fact line;
 * 16 above the title block, 14 below it. [actions] sit under the title block inside the head (`.ph-actions`,
 * 16 above), as signed-out Home's Sign in. [balanced] breaks a long title over balanced lines (a guild's name).
 */
@Composable
fun FlowHead(
    title: String,
    facts: List<String>,
    titleColor: Color = FsColor.strong,
    crest: String? = null,
    lead: (@Composable () -> Unit)? = null,
    /** The head's answer under the title block, 16 below it (the mock's `.ans`): a secondary page's figure and its line. */
    answer: (@Composable ColumnScope.() -> Unit)? = null,
    actions: (@Composable () -> Unit)? = null,
    balanced: Boolean = false,
) {
    Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD).padding(top = FsSpace.space4, bottom = HEAD_BOTTOM)) {
        Row(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE), verticalAlignment = Alignment.CenterVertically) {
            val image: (@Composable () -> Unit)? = lead ?: crest?.let { { ClassCrest(it, CREST, Modifier.parityItem(ParityItemKind.Image)) } }
            if (image != null) {
                image()
                Spacer(Modifier.width(FsSpace.space4))
            }
            // The title block is as wide as its widest line (the mock's shrink-wrapped head), the facts across it.
            val bounded = if (image != null) Modifier.weight(1f, fill = false) else Modifier
            Column(bounded.width(IntrinsicSize.Max)) {
                // A title that wraps keeps two lines of 38 (Marcellus overshoots its box on the outer lines); one line is unchanged.
                var lines by remember { mutableIntStateOf(1) }
                Text(
                    title, style = if (balanced) BALANCED_TITLE else FsType.title, color = titleColor, onTextLayout = { lines = it.lineCount },
                    modifier = Modifier.semantics { heading() }.then(if (lines > 1) Modifier.linesBox(FsType.title.lineHeight, lines) else Modifier),
                )
                if (facts.isNotEmpty()) {
                    // The facts wrap (`.ph-meta`, gap 6 by 16) when they outgrow one line.
                    FlowRow(
                        Modifier.padding(top = META_GAP).fillMaxWidth().parityRegion(ParityRegion.HEAD_META),
                        horizontalArrangement = Arrangement.spacedBy(FsSpace.space4),
                        verticalArrangement = Arrangement.spacedBy(META_GAP),
                    ) {
                        facts.forEach { Text(it, style = FsType.small, color = FsColor.muted) }
                    }
                }
            }
        }
        if (answer != null) {
            Column(Modifier.padding(top = FsSpace.space4).fillMaxWidth().parityRegion(ParityRegion.HEAD_ANSWER), content = answer)
        }
        if (actions != null) Box(Modifier.padding(top = FsSpace.space4)) { actions() }
    }
}

/** The head's game icon (a raid's zone), 56 in the rounded frame of a row's lead. */
@Composable
fun FlowHeadIcon(icon: String?, build: String?) {
    Box(
        Modifier
            .size(CREST)
            .parityItem(ParityItemKind.Image)
            .clip(RoundedCornerShape(FsRadius.control))
            .background(FsColor.raised)
            .border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.control)),
    ) { GameIconImage(icon, build, CREST) }
}

/**
 * A titled section (the mock's Section): 28 above, the title (18/24) with its action or meta on the
 * right, 10 under it and a rule, then [content]. Its region is `section-<slug of the title>`.
 */
@Composable
fun FlowSection(
    title: String,
    action: String? = null,
    onAction: () -> Unit = {},
    meta: String? = null,
    top: Dp = FsSpace.space7,
    content: @Composable ColumnScope.() -> Unit,
) {
    Column(Modifier.padding(top = top).fillMaxWidth().parityRegion(ParityRegion.section(title))) {
        Row(Modifier.fillMaxWidth().padding(bottom = SECTION_TITLE_GAP), verticalAlignment = Alignment.Top) {
            Text(title, style = FsType.section, color = FsColor.strong, modifier = Modifier.semantics { heading() })
            // The meta sits beside the title (`Needs you  6 open`); the action at the far end.
            if (meta != null) Text(meta, style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(start = FsSpace.space3, top = ACTION_DROP))
            Spacer(Modifier.weight(1f))
            if (action != null) {
                Text(
                    action,
                    style = ACTION,
                    color = FsColor.strong,
                    modifier = Modifier.padding(top = ACTION_DROP).clickable(role = Role.Button, onClick = onAction),
                )
            }
        }
        HorizontalDivider(color = FsColor.line)
        content()
    }
}

/** A block without a title (the mock's `.blk`): 24 above, its own region (`countdowns`, `section-error`). */
@Composable
fun FlowBlock(region: String, top: Dp = FsSpace.space6, content: @Composable ColumnScope.() -> Unit) {
    Column(Modifier.padding(top = top).fillMaxWidth().parityRegion(region), content = content)
}

/** A small line of text (`.t-small`, 13/20, muted). */
@Composable
fun FlowNote(text: String, modifier: Modifier = Modifier, color: Color = FsColor.muted) {
    Text(text, style = FsType.small, color = color, modifier = modifier)
}

/** A loading placeholder (`.skel`). */
@Composable
fun FlowSkeleton(width: Dp, height: Dp, modifier: Modifier = Modifier) {
    Box(modifier.width(width).height(height).background(FsColor.raised, RoundedCornerShape(SKELETON_RADIUS)))
}

/** A thin rule between rows. */
@Composable
fun FlowRule() {
    HorizontalDivider(color = FsColor.line)
}

/** A full-width spacer of [height]. */
@Composable
fun FlowGap(height: Dp) {
    Spacer(Modifier.size(width = 0.dp, height = height))
}

/** Holds wrapped text to [lines] lines of [lineHeight], centred, as a CSS line box does. */
fun Modifier.linesBox(lineHeight: TextUnit, lines: Int): Modifier = layout { measurable, constraints ->
    val placeable = measurable.measure(constraints)
    val box = minOf(placeable.height, (lineHeight * lines).roundToPx())
    layout(placeable.width, box) { placeable.place(0, (box - placeable.height) / 2) }
}

private val CREST = 56.dp
private val BALANCED_TITLE = FsType.title.copy(lineBreak = LineBreak.Heading)
private val HEAD_BOTTOM = 14.dp
private val META_GAP = 6.dp
private val SECTION_TITLE_GAP = 10.dp
private val ACTION_DROP = 2.dp
private val SKELETON_RADIUS = 6.dp
private val ACTION = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 24.sp)
