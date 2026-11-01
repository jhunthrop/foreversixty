package gg.foreversixty.app.paperdoll

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.PathEffect
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalWindowInfo
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.role
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.reports.ui.isLargeText

/**
 * The Paperdoll (spec W3 §4, design/specs/2026-10-31-paperdoll.md phone frame): two mirrored
 * columns with the sockets on the outer edges and the names toward the centre in quality colour,
 * the weapons in a row under them. Largest text: one column of 17 rows. [open] shows a socket's item.
 */
@Composable
fun PaperdollFrame(slots: List<PaperdollSlot>, build: String?, open: (PaperdollSlot) -> Unit, modifier: Modifier = Modifier) {
    val width = with(LocalDensity.current) { LocalWindowInfo.current.containerSize.width.toDp() }
    val socket = if (width < NARROW_WIDTH) PaperdollLayout.socketNarrow else PaperdollLayout.socket
    if (isLargeText()) {
        Column(modifier.fillMaxWidth(), verticalArrangement = Arrangement.spacedBy(PaperdollLayout.rowGap)) {
            val bySlot = slots.associateBy { it.slot }
            (Paperdoll.LEFT + Paperdoll.RIGHT + Paperdoll.WEAPONS).mapNotNull(bySlot::get).forEach { Cell(it, build, socket, mirrored = false, open) }
        }
        return
    }
    val columns = Paperdoll.columns(slots)
    Column(modifier.fillMaxWidth(), verticalArrangement = Arrangement.spacedBy(PaperdollLayout.rowGap)) {
        Row(horizontalArrangement = Arrangement.spacedBy(PaperdollLayout.columnGap)) {
            Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(PaperdollLayout.rowGap)) {
                columns.left.forEach { Cell(it, build, socket, mirrored = false, open) }
            }
            Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(PaperdollLayout.rowGap)) {
                columns.right.forEach { Cell(it, build, socket, mirrored = true, open) }
            }
        }
        HorizontalDivider(Modifier.padding(vertical = 4.dp), color = FsColor.lineSoft)
        Row(horizontalArrangement = Arrangement.spacedBy(PaperdollLayout.columnGap)) {
            columns.weapons.forEach { WeaponCell(it, build, socket, open, Modifier.weight(1f)) }
        }
    }
}

/** A weapon under the columns: the socket over its words, centred, so a name keeps the cell's width. */
@Composable
private fun WeaponCell(slot: PaperdollSlot, build: String?, socket: Dp, open: (PaperdollSlot) -> Unit, modifier: Modifier) {
    val label = Paperdoll.label(slot, LocalResources.current)
    val tap = if (slot.item != null) Modifier.clickable(role = Role.Button) { open(slot) } else Modifier
    Column(
        modifier.heightIn(min = PaperdollLayout.cellHeight).clip(RoundedCornerShape(8.dp)).then(tap)
            .padding(4.dp).clearAndSetSemantics { contentDescription = label; if (slot.item != null) role = Role.Button },
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.spacedBy(6.dp),
    ) {
        Socket(slot, build, socket)
        Words(slot, Alignment.CenterHorizontally, TextAlign.Center, Modifier)
    }
}

/** One socket and its words; the right column mirrors so the socket stays on the outer edge. */
@Composable
private fun Cell(slot: PaperdollSlot, build: String?, socket: Dp, mirrored: Boolean, open: (PaperdollSlot) -> Unit, modifier: Modifier = Modifier) {
    val label = Paperdoll.label(slot, LocalResources.current)
    val tap = if (slot.item != null) Modifier.clickable(role = Role.Button) { open(slot) } else Modifier
    Row(
        modifier.fillMaxWidth().heightIn(min = PaperdollLayout.cellHeight).clip(RoundedCornerShape(8.dp)).then(tap)
            .padding(4.dp).clearAndSetSemantics { contentDescription = label; if (slot.item != null) role = Role.Button },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(8.dp, if (mirrored) Alignment.End else Alignment.Start),
    ) {
        if (!mirrored) Socket(slot, build, socket)
        Words(slot, if (mirrored) Alignment.End else Alignment.Start, if (mirrored) TextAlign.End else TextAlign.Start, Modifier.weight(1f, fill = false))
        if (mirrored) Socket(slot, build, socket)
    }
}

@Composable
private fun Words(slot: PaperdollSlot, horizontal: Alignment.Horizontal, align: TextAlign, modifier: Modifier) {
    val nameColor = if (slot.isEmpty) FsColor.muted else ItemQuality(slot.item?.quality).text
    Column(modifier, horizontalAlignment = horizontal) {
        Text(
            slot.name, style = PaperdollLayout.name, color = nameColor, textAlign = align,
            maxLines = PaperdollLayout.NAME_LINES, overflow = TextOverflow.Ellipsis,
        )
        Text(slot.detail, style = PaperdollLayout.detail, color = FsColor.muted, textAlign = align)
    }
}

/** A 40 dp socket in its 2 dp quality border on `#14171f`; an empty slot is a dashed outline. The gold dot sits on its corner. */
@Composable
private fun Socket(slot: PaperdollSlot, build: String?, size: Dp) {
    val shape = RoundedCornerShape(PaperdollLayout.socketRadius)
    val item = slot.item
    Box {
        val frame = if (item == null) {
            Modifier.drawBehind {
                drawRoundRect(
                    color = FsColor.line,
                    cornerRadius = CornerRadius(PaperdollLayout.socketRadius.toPx()),
                    style = Stroke(width = 1.5.dp.toPx(), pathEffect = PathEffect.dashPathEffect(floatArrayOf(DASH, GAP))),
                )
            }
        } else {
            Modifier.background(PaperdollLayout.ground, shape).border(2.dp, ItemQuality(item.quality).border, shape)
        }
        Box(Modifier.size(size).then(frame).clip(shape), contentAlignment = Alignment.Center) {
            if (item != null) {
                GameIconImage(item.icon, build, size - 4.dp)
                Box(Modifier.size(size).padding(2.dp).border(1.dp, Color.Black.copy(alpha = INNER_LINE), shape))
            }
        }
        if (slot.better) {
            Box(
                Modifier.align(Alignment.TopEnd).offset(x = DOT_OFFSET, y = -DOT_OFFSET).size(PaperdollLayout.dot)
                    .background(FsColor.gold, CircleShape).border(1.5.dp, FsColor.raised, CircleShape),
            )
        }
    }
}

/** The Paperdoll's sizes (spec W3 §4 and §9). */
object PaperdollLayout {
    val socket = 40.dp
    val socketNarrow = 36.dp
    val socketRadius = 6.dp
    val dot = 10.dp
    val columnGap = 8.dp
    val rowGap = 4.dp
    val cellHeight = 48.dp
    val ground = Color(0xFF14171F)
    const val NAME_LINES = 3
    val name = FsType.meta.copy(fontWeight = FsType.bodyStrong.fontWeight, fontSize = 13.sp, lineHeight = 16.sp)
    val detail = FsType.meta.copy(fontSize = 12.sp, lineHeight = 15.sp)
}

/** Below this screen width the sockets drop to 36 dp (spec W3 §9). */
private val NARROW_WIDTH = 360.dp
private const val DASH = 8f
private const val GAP = 6f
private const val INNER_LINE = 0.65f
private val DOT_OFFSET = 3.dp
