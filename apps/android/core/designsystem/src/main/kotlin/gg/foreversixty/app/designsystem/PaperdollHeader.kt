package gg.foreversixty.app.designsystem

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
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
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

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
 * (the app's item icon), which sits in the 44 dp socket. iOS: `PaperdollHeader`.
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
    icon: @Composable (PaperdollListSlot.Item) -> Unit,
) {
    Column(modifier.fillMaxWidth().panel()) {
        Strip(title, figure, unit, portrait)
        Column(
            Modifier.padding(FsSpace.space3).parityRegion(ParityRegion.SLOTS),
            verticalArrangement = Arrangement.spacedBy(ROW_GAP),
        ) {
            slots.chunked(2).forEach { pair ->
                Row(horizontalArrangement = Arrangement.spacedBy(COLUMN_GAP)) {
                    pair.forEach { slot ->
                        Box(Modifier.weight(1f)) { Cell(slot, slot.id == selected, onSelect, icon) }
                    }
                    if (pair.size == 1) Spacer(Modifier.weight(1f))
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
        PaperdollMedallion(portrait, MEDALLION_SIZE)
        Text(
            title,
            modifier = Modifier.weight(1f).semantics { heading() },
            style = TextStyle(fontFamily = FsFont.title, fontSize = TITLE_SIZE, lineHeight = TITLE_LINE),
            color = FsColor.cream,
        )
        if (figure != null) {
            Text(
                buildAnnotatedString {
                    append(figure)
                    withStyle(SpanStyle(fontSize = UNIT_SIZE, color = FsColor.text)) { append(" $unit") }
                },
                style = TextStyle(
                    fontFamily = FsFont.body,
                    fontWeight = FontWeight.SemiBold,
                    fontSize = FIGURE_SIZE,
                    fontFeatureSettings = "tnum",
                    shadow = Shadow(Color.Black, Offset(0f, 1f), 2f),
                ),
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
                contentDescription = "${item.name}, ${slot.label}"
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
                    style = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.SemiBold, fontSize = 13.sp, lineHeight = 16.sp),
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

@Composable
private fun Sub(label: String, detail: String?, gain: Boolean) {
    Text(
        buildAnnotatedString {
            append(label)
            if (detail != null) {
                append(" · ")
                withStyle(SpanStyle(color = if (gain) FsColor.gain else FsColor.muted, fontWeight = if (gain) FontWeight.Bold else null)) {
                    append(detail)
                }
            }
        },
        style = TextStyle(fontFamily = FsFont.body, fontSize = 12.sp, lineHeight = 16.sp),
        color = FsColor.muted,
        maxLines = 1,
        overflow = TextOverflow.Ellipsis,
    )
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
fun PaperdollMedallion(portrait: Painter?, size: Dp) {
    Image(
        portrait ?: painterResource(BrandArt.Seal.res),
        contentDescription = null,
        contentScale = ContentScale.Crop,
        modifier = Modifier.size(size).clip(CircleShape),
    )
}

private val MEDALLION_SIZE = 40.dp
private val CELL_SIZE = 44.dp
private val ICON_INSET = 3.dp
private val ROW_GAP = 6.dp
private val COLUMN_GAP = 10.dp
private val STRIP_GAP = 10.dp
private val SELECTED_RING = 2.dp
private val SELECTED_GAP = 3.dp
private val TITLE_SIZE = 19.sp
private val TITLE_LINE = 23.sp
private val FIGURE_SIZE = 17.sp
private val UNIT_SIZE = 11.sp
private const val COVERED_ALPHA = 0.55f
