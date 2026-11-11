package gg.foreversixty.app.flows

import androidx.compose.foundation.ScrollState
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.compositionLocalOf
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
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.asPaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.statusBars
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.derivedStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.layout
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.BandKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.fadingBand
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.designsystem.panel
import gg.foreversixty.app.designsystem.screenGround
import gg.foreversixty.app.designsystem.surfaceBand
import gg.foreversixty.app.items.GameIconImage

/**
 * The secondary ink of the surface a flow row sits on: the panel's `muted` by default; a section
 * without a panel (rows on the painted ground) provides `mutedGround`, AA there.
 */
val LocalSubInk = compositionLocalOf { FsColor.muted }

/**
 * The flows screens' layout (design/mocks/ws_app_flows.py `.ap`, spec 2026-11-08): one scrolling column
 * under the profile bar, 8 above, 16 at the sides, 24 below. Home, Raid and Gear are built from these.
 */
@Composable
fun FlowPage(modifier: Modifier = Modifier, content: @Composable ColumnScope.() -> Unit) {
    Column(
        modifier
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
            .parityContent()
            .padding(start = FsSpace.space4, end = FsSpace.space4, top = FsSpace.space2, bottom = FsSpace.space6),
        content = content,
    )
}

/**
 * A tab root in the round-5 look (spec Mobile): the painted [band] behind the status bar, [topBar]
 * (and [head] on Home's hero), fading into the ground below its foot as the site's head does
 * (`fadingBand`), then [notice] and the blocks over the fade, 16 under the band, 14 at the sides and
 * 20 apart. Once the page scrolls, a strip of the band holds the status bar so the clock never sits
 * on the content. Parity regions are measured from the bottom of [topBar], as the mocks' are from
 * the nav bar.
 */
@Composable
fun FlowBandPage(
    topBar: @Composable () -> Unit,
    band: BandKind = BandKind.Band,
    head: @Composable () -> Unit = {},
    notice: @Composable () -> Unit = {},
    content: @Composable ColumnScope.() -> Unit,
) {
    val scroll = rememberScrollState()
    val scrolled by remember { derivedStateOf { scroll.value > 0 } }
    val statusBar = WindowInsets.statusBars.asPaddingValues().calculateTopPadding()
    Box(Modifier.reachUp(statusBar).fillMaxSize()) {
        BandPageColumn(scroll, statusBar, band, topBar, head, notice, content)
        if (scrolled) Box(Modifier.fillMaxWidth().height(statusBar).surfaceBand(bottomLine = false, kind = band))
    }
}

@Composable
private fun BandPageColumn(
    scroll: ScrollState,
    statusBar: Dp,
    band: BandKind,
    topBar: @Composable () -> Unit,
    head: @Composable () -> Unit,
    notice: @Composable () -> Unit,
    content: @Composable ColumnScope.() -> Unit,
) {
    Column(Modifier.fillMaxSize().screenGround().verticalScroll(scroll)) {
        Column(Modifier.fillMaxWidth().fadingBand(band)) {
            Spacer(Modifier.height(statusBar))
            topBar()
            Spacer(Modifier.fillMaxWidth().height(0.dp).parityContent())
            head()
        }
        notice()
        Column(
            Modifier.fillMaxWidth().padding(start = BAND_PAGE_SIDE, end = BAND_PAGE_SIDE, top = FsSpace.space4, bottom = BAND_PAGE_BOTTOM),
            verticalArrangement = Arrangement.spacedBy(FsSpace.space5),
            content = content,
        )
    }
}

/**
 * Grows this element up by [reach] past its top: the tab scaffold (ForeverTabScaffold) pads its
 * content below the status bar, and a band page draws its band back up under it, edge to edge.
 * With no status bar inset (the screenshot tests) it changes nothing.
 */
private fun Modifier.reachUp(reach: Dp): Modifier = layout { measurable, constraints ->
    val px = reach.roundToPx()
    val grown = constraints.copy(
        minHeight = constraints.minHeight + px,
        maxHeight = if (constraints.hasBoundedHeight) constraints.maxHeight + px else constraints.maxHeight,
    )
    val placeable = measurable.measure(grown)
    layout(placeable.width, placeable.height - px) { placeable.place(0, -px) }
}

/**
 * The page head (the mock's PageHeader): an optional 56 crest, the title in its colour and the fact line;
 * 16 above the title block, 14 below it.
 */
@Composable
fun FlowHead(
    title: String,
    facts: List<String>,
    titleColor: Color = FsColor.strong,
    crest: String? = null,
    lead: (@Composable () -> Unit)? = null,
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
                Text(title, style = FsType.title, color = titleColor, modifier = Modifier.semantics { heading() })
                if (facts.isNotEmpty()) {
                    // The facts wrap (`.ph-meta`, gap 6 by 16) when they outgrow one line.
                    FlowRow(
                        Modifier.padding(top = META_GAP).fillMaxWidth().parityRegion(ParityRegion.HEAD_META),
                        horizontalArrangement = Arrangement.spacedBy(FsSpace.space4),
                        verticalArrangement = Arrangement.spacedBy(META_GAP),
                    ) {
                        facts.forEach { Text(it, style = FsType.small, color = FsColor.mutedGround) }
                    }
                }
            }
        }
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
 * With a [panel] (round 5, spec Mobile) the title is the cream display face over the ground with a
 * gold action, 8 under it, and [content] sits in the round-5 panel with that frame, no rule.
 */
@Composable
fun FlowSection(
    title: String,
    action: String? = null,
    onAction: () -> Unit = {},
    meta: String? = null,
    top: Dp = FsSpace.space7,
    panel: PanelFrame? = null,
    content: @Composable ColumnScope.() -> Unit,
) {
    Column(Modifier.padding(top = top).fillMaxWidth().parityRegion(ParityRegion.section(title))) {
        val round5 = panel != null
        Row(
            Modifier.fillMaxWidth().padding(bottom = if (round5) FsSpace.space2 else SECTION_TITLE_GAP),
            verticalAlignment = if (round5) Alignment.CenterVertically else Alignment.Top,
        ) {
            Text(
                title,
                style = if (round5) ScaledType.sectionTitle else FsType.section,
                color = if (round5) FsColor.cream else FsColor.strong,
                modifier = Modifier.semantics { heading() },
            )
            // The meta sits beside the title (`Needs you  6 open`); the action at the far end.
            if (meta != null) Text(meta, style = FsType.small, color = FsColor.mutedGround, modifier = Modifier.padding(start = FsSpace.space3, top = ACTION_DROP))
            Spacer(Modifier.weight(1f))
            if (action != null) {
                Text(
                    action,
                    style = ACTION,
                    color = if (round5) FsColor.gold else FsColor.strong,
                    modifier = Modifier.padding(top = if (round5) 0.dp else ACTION_DROP).clickable(role = Role.Button, onClick = onAction),
                )
            }
        }
        if (panel == null) {
            HorizontalDivider(color = FsColor.line)
            CompositionLocalProvider(LocalSubInk provides FsColor.mutedGround) { content() }
        } else {
            FlowPanel(frame = panel, content = content)
        }
    }
}

/**
 * The round-5 panel a section's rows sit in (spec Mobile): the brown-black panel with [frame], 16 at
 * the sides and 8 above and below, so a row's own padding finishes the inset.
 */
@Composable
fun FlowPanel(modifier: Modifier = Modifier, frame: PanelFrame = PanelFrame.None, content: @Composable ColumnScope.() -> Unit) {
    Box(modifier.fillMaxWidth().panel(frame)) {
        Column(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4, vertical = FsSpace.space2), content = content)
    }
}

/** A block without a title (the mock's `.blk`): 24 above, its own region (`countdowns`, `section-error`). */
@Composable
fun FlowBlock(region: String, top: Dp = FsSpace.space6, content: @Composable ColumnScope.() -> Unit) {
    Column(Modifier.padding(top = top).fillMaxWidth().parityRegion(region), content = content)
}

/** A small line of text on the ground (`.t-small`, 13/20, mutedGround: AA on the painted art). */
@Composable
fun FlowNote(text: String, modifier: Modifier = Modifier, color: Color = FsColor.mutedGround) {
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

private val CREST = 56.dp
private val BAND_PAGE_SIDE = 14.dp
private val BAND_PAGE_BOTTOM = 22.dp
private val HEAD_BOTTOM = 14.dp
private val META_GAP = 6.dp
private val SECTION_TITLE_GAP = 10.dp
private val ACTION_DROP = 2.dp
private val SKELETON_RADIUS = 6.dp
private val ACTION = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 24.sp)
