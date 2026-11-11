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
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
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
import gg.foreversixty.app.home.LARGE_TEXT_SCALE
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
 * the nav bar. [page] carries the scroll a route keeps across a push and its pull to refresh (R12).
 */
@Composable
fun FlowBandPage(
    topBar: @Composable () -> Unit,
    band: BandKind = BandKind.Band,
    head: @Composable () -> Unit = {},
    notice: @Composable () -> Unit = {},
    page: FlowPageState = FlowPageState(rememberScrollState()),
    content: @Composable ColumnScope.() -> Unit,
) {
    val scroll = page.scroll
    val scrolled by remember(scroll) { derivedStateOf { scroll.value > 0 } }
    val statusBar = WindowInsets.statusBars.asPaddingValues().calculateTopPadding()
    Box(Modifier.reachUp(statusBar).fillMaxSize()) {
        FlowRefreshBox(page.refresh, Modifier.fillMaxSize().screenGround()) {
            BandPageColumn(scroll, statusBar, band, topBar, head, notice, content)
        }
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
    Column(Modifier.fillMaxSize().verticalScroll(scroll)) {
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
 * A [stamp] (`Synced 2 h ago`, round 3 row 8) sits under the meta line, outside the panel, in the
 * warning tone when [stampWarn].
 */
@Composable
fun FlowSection(
    title: String,
    action: String? = null,
    onAction: () -> Unit = {},
    meta: String? = null,
    top: Dp = FsSpace.space7,
    panel: PanelFrame? = null,
    stamp: String? = null,
    stampWarn: Boolean = false,
    content: @Composable ColumnScope.() -> Unit,
) {
    Column(Modifier.padding(top = top).fillMaxWidth().parityRegion(ParityRegion.section(title))) {
        val round5 = panel != null
        val style = SectionHeadStyle(round5)
        val actionText: (@Composable () -> Unit)? = action?.let { { SectionAction(it, style, onAction) } }
        Column(Modifier.fillMaxWidth().padding(bottom = style.bottom)) {
            if (meta != null && LocalDensity.current.fontScale >= LARGE_TEXT_SCALE) {
                StackedSectionHead(title, meta, style, actionText)
            } else {
                InlineSectionHead(title, meta, style, actionText)
            }
            stamp?.let { Text(it, style = FsType.small, color = if (stampWarn) FsColor.warn else FsColor.mutedGround) }
        }
        if (panel == null) {
            HorizontalDivider(color = FsColor.line)
            CompositionLocalProvider(LocalSubInk provides FsColor.mutedGround) { content() }
        } else {
            FlowPanel(frame = panel) {
                // Under a stamp the panel's first row sits 8 lower (the mock's `.stmp + *`).
                if (stamp != null) Spacer(Modifier.height(FsSpace.space2))
                content()
            }
        }
    }
}

/** A section head's look: the round-5 cream display title with a gold action, or the plain one. */
private class SectionHeadStyle(val round5: Boolean) {
    val title get() = if (round5) PANEL_TITLE else FsType.section
    val titleColor get() = if (round5) FsColor.cream else FsColor.strong
    val actionColor get() = if (round5) FsColor.gold else FsColor.strong
    val actionDrop get() = if (round5) 0.dp else ACTION_DROP
    val bottom get() = if (round5) FsSpace.space2 else SECTION_TITLE_GAP
}

@Composable
private fun SectionTitle(title: String, style: SectionHeadStyle, modifier: Modifier = Modifier) {
    Text(title, style = style.title, color = style.titleColor, modifier = modifier.semantics { heading() })
}

@Composable
private fun SectionAction(action: String, style: SectionHeadStyle, onAction: () -> Unit) {
    Text(
        action,
        style = ACTION,
        color = style.actionColor,
        modifier = Modifier.padding(top = style.actionDrop).clickable(role = Role.Button, onClick = onAction),
    )
}

/** The title, the meta beside it (`Needs you  6 open`) and the action at the far end, on one line. */
@Composable
private fun InlineSectionHead(title: String, meta: String?, style: SectionHeadStyle, action: (@Composable () -> Unit)?) {
    Row(
        Modifier.fillMaxWidth(),
        verticalAlignment = if (style.round5) Alignment.CenterVertically else Alignment.Top,
    ) {
        SectionTitle(title, style)
        if (meta != null) Text(meta, style = FsType.small, color = FsColor.mutedGround, modifier = Modifier.padding(start = FsSpace.space3, top = ACTION_DROP))
        Spacer(Modifier.weight(1f))
        action?.invoke()
    }
}

/**
 * From font scale 1.5 (spec 3.2) the meta moves under the title, both wrapping in the space the
 * action leaves; the action keeps its natural width at the far end of the title's line.
 */
@Composable
private fun StackedSectionHead(title: String, meta: String, style: SectionHeadStyle, action: (@Composable () -> Unit)?) {
    Row(Modifier.fillMaxWidth(), verticalAlignment = Alignment.Top) {
        Column(Modifier.weight(1f)) {
            SectionTitle(title, style)
            Text(meta, style = FsType.small, color = FsColor.mutedGround, modifier = Modifier.padding(top = ACTION_DROP))
        }
        if (action != null) {
            Spacer(Modifier.width(FsSpace.space3))
            action()
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
private val ACTION = FsType.chip

/** A round-5 section title over the ground: the display face at 21 (the mock's `.sh h2`). */
private val PANEL_TITLE = ScaledType.sectionTitle
