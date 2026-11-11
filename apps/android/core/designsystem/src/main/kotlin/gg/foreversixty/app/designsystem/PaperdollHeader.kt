package gg.foreversixty.app.designsystem

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.draw.drawWithCache
import androidx.compose.ui.geometry.Rect
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Paint
import androidx.compose.ui.graphics.drawscope.DrawScope
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.graphics.drawscope.drawIntoCanvas
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.ColorFilter
import androidx.compose.ui.graphics.ColorMatrix
import androidx.compose.ui.graphics.Shadow
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp

/** One slot of the phone paperdoll's list: the slot's word, and the item in it or nothing. */
data class PaperdollListSlot(
    /** The slot key (`head`, `finger1`, `off_hand`). */
    val id: String,
    /** The slot's word (`Head`, `Off Hand`). */
    val label: String,
    val item: Item?,
    /** The off hand under a two-hander: the main hand's icon greyed, with the caller's word for it. */
    val coveredBy: Item? = null,
    val coveredNote: String? = null,
) {
    /** What a filled slot shows. */
    data class Item(
        val id: Int,
        val name: String,
        /** The client's quality colour for the name (the items' `ItemQuality`). */
        val quality: Color,
        /** The second line after the slot's word: a gain (`+5.9 DPS`) or `Worn`. */
        val detail: String? = null,
        /** A gain draws its detail in the gain green. */
        val isGain: Boolean = false,
    )
}

/**
 * The paperdoll's phone form (spec round 5, Paperdoll "Phone"; web ui/PaperdollFrame.svelte's list
 * form): under 36 dp slots the frame gives way to its one-line header strip (the medallion with the
 * character's [portrait], or the LX seal; the [title]; the set's [figure]) and the slots in two
 * columns on the panel. A tap on a filled slot calls [onSelect]. [icon] draws a slot's item icon
 * (the app's item icon), which sits in the 44 dp socket. [lead] is a line under the strip (Gear's
 * band, faction and Change). A name wraps and never truncates (spec 3.2), a third line only when its
 * words need it; a gain never shares its line with a cut slot word, it drops under it whole. iOS:
 * `PaperdollHeader`.
 */
@Composable
fun PaperdollHeader(
    title: String,
    figure: String?,
    portrait: Painter?,
    slots: List<PaperdollListSlot>,
    modifier: Modifier = Modifier,
    unit: String = "DPS",
    selected: String? = null,
    onSelect: ((String) -> Unit)? = null,
    lead: (@Composable () -> Unit)? = null,
    icon: @Composable (PaperdollListSlot.Item) -> Unit,
) {
    Column(modifier.fillMaxWidth().panel()) {
        Strip(title, figure, unit, portrait)
        // The line that names the list (the band, the faction, how much of it is worn), under the strip.
        if (lead != null) Box(Modifier.fillMaxWidth().padding(start = FsSpace.space3, end = FsSpace.space3, top = FsSpace.space3)) { lead() }
        Column(
            Modifier.padding(FsSpace.space3).parityRegion(ParityRegion.SLOTS),
            verticalArrangement = Arrangement.spacedBy(ROW_GAP),
        ) {
            // Large text gets one slot per row, so a name never breaks inside a word.
            val columns = if (LocalDensity.current.fontScale >= ONE_COLUMN_SCALE) 1 else 2
            slots.chunked(columns).forEach { pair ->
                Row(horizontalArrangement = Arrangement.spacedBy(COLUMN_GAP)) {
                    pair.forEach { slot ->
                        Box(Modifier.weight(1f)) { Cell(slot, slot.id == selected, onSelect, icon) }
                    }
                    if (pair.size < columns) Spacer(Modifier.weight(1f))
                }
            }
        }
    }
}

@Composable
private fun Strip(title: String, figure: String?, unit: String, portrait: Painter?) {
    Row(
        Modifier
            .fillMaxWidth()
            .background(FsColor.navBar)
            .drawBehind { drawLine(Color.Black, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx()) }
            .padding(horizontal = FsSpace.space3, vertical = FsSpace.space2),
        horizontalArrangement = Arrangement.spacedBy(STRIP_GAP),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        PaperdollMedallion(portrait, MEDALLION_SIZE, ring = true)
        Text(
            title,
            modifier = Modifier.weight(1f).semantics { heading() },
            style = FsType.sectionReading,
            color = FsColor.cream,
        )
        if (figure != null) {
            Text(
                buildAnnotatedString {
                    append(figure)
                    withStyle(FsType.figureUnit.toSpanStyle().copy(color = FsColor.text)) { append(" $unit") }
                },
                style = FsType.figureSm.copy(shadow = Shadow(Color.Black, Offset(0f, 1f), 2f)),
                color = FsColor.strong,
                maxLines = 1,
            )
        }
    }
}

@Composable
private fun Cell(
    slot: PaperdollListSlot,
    selected: Boolean,
    onSelect: ((String) -> Unit)?,
    icon: @Composable (PaperdollListSlot.Item) -> Unit,
) {
    val item = slot.item
    val tap = if (item != null && onSelect != null) {
        Modifier
            .clickable(role = Role.Button) { onSelect(slot.id) }
            .clearAndSetSemantics {
                contentDescription = listOfNotNull(item.name, slot.label, item.detail).joinToString(", ")
                this.selected = selected
            }
    } else {
        Modifier
    }
    Row(
        Modifier.fillMaxWidth().heightIn(min = FsTarget.min).then(tap).parityItem(ParityItemKind.Row),
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space2),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Socket(item ?: slot.coveredBy, covered = item == null, selected = selected, icon = icon)
        Column(Modifier.weight(1f)) {
            if (item != null) {
                Text(
                    item.name,
                    style = FsType.paperdollName,
                    color = item.quality,
                )
                Sub(slot.label, item.detail, item.isGain)
            } else {
                Sub(slot.label, null, false)
                slot.coveredNote?.let { Sub(it, null, false) }
            }
        }
    }
}

/**
 * The slot's word and its detail (`Main Hand · +5.9 DPS`). Each part keeps its words together, so
 * when both do not fit the detail wraps under the word whole and never runs into the next column.
 */
@Composable
private fun Sub(label: String, detail: String?, gain: Boolean) {
    val style = FsType.paperdollDetail
    if (detail == null) {
        Text(label, style = style, color = FsColor.muted, maxLines = 1, overflow = TextOverflow.Ellipsis)
        return
    }
    // The separator joins its word and the detail in one node, so TalkBack never reads a lone `·` (N6).
    FlowRow(Modifier.fillMaxWidth().clearAndSetSemantics { contentDescription = "$label, $detail" }) {
        Text("${label.replace(' ', NBSP)}$NBSP·$NBSP", style = style, color = FsColor.muted, maxLines = 1, softWrap = false)
        Text(
            detail.replace(' ', NBSP),
            style = style.copy(fontWeight = if (gain) FontWeight.Bold else null),
            color = if (gain) FsColor.gain else FsColor.muted,
            maxLines = 1,
            softWrap = false,
        )
    }
}

@Composable
private fun Socket(item: PaperdollListSlot.Item?, covered: Boolean, selected: Boolean, icon: @Composable (PaperdollListSlot.Item) -> Unit) {
    val shape = RoundedCornerShape(FsRadius.pill)
    // The selected ring sits outside the socket (`outline-offset: 1px`), so selecting moves nothing.
    val ring = if (selected) Modifier.drawBehind { drawSelectedRing() } else Modifier
    Box(
        ring
            .size(CELL_SIZE)
            .clip(shape)
            .background(FsColor.bg)
            .border(1.dp, FsColor.line, shape)
            .padding(ICON_INSET),
    ) {
        if (item != null) {
            Box(if (covered) Modifier.alpha(COVERED_ALPHA).greyscale() else Modifier) { icon(item) }
        }
    }
}

/** The cream ring around a selected socket, [SELECTED_GAP] outside its edge. */
private fun DrawScope.drawSelectedRing() {
    val stroke = SELECTED_RING.toPx()
    val outset = SELECTED_GAP.toPx() + stroke / 2
    drawRoundRect(
        FsColor.cream,
        topLeft = Offset(-outset, -outset),
        size = Size(size.width + outset * 2, size.height + outset * 2),
        cornerRadius = CornerRadius(FsRadius.pill.toPx() + outset),
        style = Stroke(stroke),
    )
}

/** The covered off hand draws its icon without colour, as the client greys it. */
private fun Modifier.greyscale(): Modifier = drawWithCache {
    val paint = Paint().apply { colorFilter = ColorFilter.colorMatrix(ColorMatrix().apply { setToSaturation(0f) }) }
    onDrawWithContent {
        drawIntoCanvas { canvas ->
            canvas.saveLayer(Rect(Offset.Zero, size), paint)
            drawContent()
            canvas.restore()
        }
    }
}

/**
 * The paperdoll's medallion: the character's Battle.net [portrait], else the LX seal (owner,
 * 2026-11-08), never the class crest. A circle at [size].
 */
@Composable
fun PaperdollMedallion(portrait: Painter?, size: Dp, ring: Boolean = false) {
    // The strip's medallion wears the frame's thin gold ring (ux-designer, lane M0).
    val edge = if (ring) Modifier.border(MEDALLION_RING, FsColor.goldDeep, CircleShape).padding(MEDALLION_RING) else Modifier
    Image(
        portrait ?: painterResource(BrandArt.Seal.res),
        contentDescription = null,
        contentScale = ContentScale.Crop,
        modifier = Modifier.size(size).clip(CircleShape).then(edge).clip(CircleShape),
    )
}

private val MEDALLION_SIZE = 40.dp
private val MEDALLION_RING = 1.5.dp
private const val ONE_COLUMN_SCALE = 1.5f
private const val NBSP = '\u00a0'
private val CELL_SIZE = 44.dp
private val ICON_INSET = 3.dp
private val ROW_GAP = 6.dp
private val COLUMN_GAP = 10.dp
private val STRIP_GAP = 10.dp
private val SELECTED_RING = 2.dp
private val SELECTED_GAP = 3.dp
private const val COVERED_ALPHA = 0.55f
