package gg.foreversixty.app.bis.spec

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemQuality

/** The plate (`.plate`): every slot's cell, or the cells grouped by where to go. */
@Composable
fun BisSpecPlate(content: BisSpecContent, view: BisSpecView, build: String, actions: BisSpecActions) {
    val shape = RoundedCornerShape(FsRadius.plate)
    Column(
        Modifier.fillMaxWidth().clip(shape).background(FsColor.raised).border(1.dp, FsColor.line, shape).parityRegion(ParityRegion.PLATE),
    ) {
        DpColumn(Modifier.padding(1.dp).fillMaxWidth().parityRegion(ParityRegion.SLOTS)) {
            val shown = if (view == BisSpecView.BySlot) content.cells.size else content.groups.size
            if (shown == 0) {
                Text(stringResource(R.string.bis_searchEmpty), style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(FsSpace.space4))
            }
            when (view) {
                BisSpecView.BySlot -> content.cells.forEach { Cell(it, build, actions.onItem) }
                BisSpecView.WhereToGo -> content.groups.forEach { group ->
                    Row(Modifier.fillMaxWidth().padding(start = FsSpace.space4, end = FsSpace.space4, top = GROUP_TOP, bottom = GROUP_BOTTOM)) {
                        Text(group.place, style = FsType.section, color = FsColor.strong, modifier = Modifier.weight(1f))
                        Text(group.pieces, style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(start = FsSpace.space3, top = PIECES_DROP))
                    }
                    group.cells.forEach { Cell(it, build, actions.onItem) }
                }
            }
        }
    }
}

/** One cell (`.cell`): the icon, the slot and level, the name with its stats, the source, the gain. */
@Composable
private fun Cell(cell: BisSpecCell, build: String, onItem: (gg.foreversixty.app.items.ItemLook) -> Unit) {
    val look = cell.look
    val tap = look?.let { Modifier.clickable(role = Role.Button) { onItem(it) } } ?: Modifier
    Row(
        Modifier
            .fillMaxWidth()
            .wholeDp()
            .parityItem(ParityItemKind.Row)
            .drawBehind { drawRect(FsColor.lineSoft, topLeft = Offset(0f, size.height - 1.dp.toPx()), size = Size(size.width, 1.dp.toPx())) }
            .then(tap)
            .padding(start = FsSpace.space4, end = FsSpace.space4, top = CELL_PAD, bottom = CELL_PAD + 1.dp),
        verticalAlignment = Alignment.Top,
    ) {
        Box(Modifier.padding(top = ICON_DROP)) {
            if (look != null) {
                ItemIcon(look.icon, look.quality, build, ICON, Modifier.parityItem(ParityItemKind.Image))
            } else {
                Box(Modifier.size(ICON).border(1.dp, FsColor.lineSoft, RoundedCornerShape(ICON_RADIUS)))
            }
        }
        Spacer(Modifier.width(FsSpace.space3))
        Column(Modifier.weight(1f)) {
            Row(Modifier.padding(top = 1.dp), horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
                Text(cell.slot, style = SLOT, color = FsColor.muted)
                cell.level?.let { Text(it, style = SLOT.copy(fontWeight = FontWeight.Bold), color = FsColor.muted) }
            }
            Row(Modifier.heightIn(min = NAME_LINE)) {
                if (look != null) {
                    Text(look.name.orEmpty(), style = FsType.item, color = ItemQuality(look.quality).text)
                    if (cell.stats.isNotEmpty()) {
                        Text(
                            cell.stats, style = STATS, color = FsColor.muted, maxLines = 1, overflow = TextOverflow.Ellipsis,
                            modifier = Modifier.weight(1f).padding(start = STATS_GAP, top = 1.dp),
                        )
                    }
                } else {
                    Text(stringResource(R.string.bisSpec_nothing), style = FsType.item.copy(fontWeight = FontWeight.Normal), color = FsColor.muted)
                }
            }
            Text(cell.source, style = SOURCE, color = FsColor.muted)
        }
        Spacer(Modifier.width(FsSpace.space3))
        cell.gain?.let { gain -> Gain(gain, cell.unit, cell.versus, cell.versusItem) }
    }
}

@Composable
private fun Gain(gain: String, unit: String?, versus: String?, versusItem: gg.foreversixty.app.items.ItemLook?) {
    Column(Modifier.widthIn(max = GAIN_WIDTH), horizontalAlignment = Alignment.End) {
        Text(
            buildAnnotatedString {
                append(gain)
                unit?.let { withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" $it") } }
            },
            style = FsType.figureSm, color = FsColor.strong,
        )
        if (versus != null) {
            Text(
                buildAnnotatedString {
                    append(versus)
                    versusItem?.let { withStyle(SpanStyle(color = ItemQuality(it.quality).text)) { append(" " + it.name.orEmpty()) } }
                },
                style = VERSUS, color = FsColor.muted, maxLines = 1, overflow = TextOverflow.Ellipsis,
            )
        }
    }
}

/** What to stack: the sentence, the legend, each stat's scale factor in two columns, the addon line. */
@Composable
fun BisSpecStack(weights: BisSpecWeights, onAddon: () -> Unit) {
    FlowSection(stringResource(R.string.bisSpec_stack), top = FsSpace.space8) {
        DpColumn(Modifier.padding(top = FsSpace.space3)) {
            Text(weights.sentence, style = FsType.body.copy(lineHeight = 24.sp), color = FsColor.text, modifier = Modifier.wholeDp())
            Text(weights.legend, style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(top = FsSpace.space2).wholeDp())
            DpColumn(Modifier.padding(top = FsSpace.space3), spacing = FsSpace.space4) {
                weights.stats.chunked(2).forEach { pair ->
                    Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space4)) {
                        pair.forEach { Weight(it, Modifier.weight(1f)) }
                        if (pair.size == 1) Spacer(Modifier.weight(1f))
                    }
                }
            }
            Text(
                linked(stringResource(R.string.bisSpec_stack_addon) + " ", stringResource(R.string.bisSpec_stack_addonLink), onAddon),
                style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(top = FsSpace.space2).wholeDp(),
            )
        }
    }
}

@Composable
private fun Weight(weight: BisSpecWeight, modifier: Modifier) {
    DpColumn(modifier, spacing = WEIGHT_GAP) {
        Text(weight.name, style = WEIGHT_NAME, color = FsColor.text)
        if (weight.factor != null) {
            Text(weight.factor, style = FsType.figureSm, color = FsColor.strong)
            Box(Modifier.fillMaxWidth().height(METER).background(FsColor.lineSoft)) {
                Box(Modifier.fillMaxWidth(weight.fraction.coerceIn(0f, 1f)).fillMaxHeight().background(FsColor.gold))
            }
            weight.perPoint?.let { Text(it, style = PER_POINT, color = FsColor.muted) }
        } else {
            Text(stringResource(R.string.bisSpec_stack_insignificant), style = PER_POINT, color = FsColor.muted)
        }
    }
}

private val CELL_PAD = 9.dp
private val ICON = 40.dp
private val ICON_DROP = 2.dp
private val ICON_RADIUS = 5.dp
private val NAME_LINE = 21.dp
private val STATS_GAP = 10.dp
private val GAIN_WIDTH = 104.dp
private val GROUP_TOP = 14.dp
private val GROUP_BOTTOM = 8.dp
private val PIECES_DROP = 2.dp
private val WEIGHT_GAP = 6.dp
private val METER = 8.dp
private val SLOT = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
private val STATS = FsType.small
private val SOURCE = FsType.small.copy(lineHeight = 18.sp)
private val VERSUS = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp)
private val WEIGHT_NAME = FsType.bodyStrong.copy(lineHeight = 23.sp)
private val PER_POINT = FsType.body.copy(fontSize = 14.sp, lineHeight = 21.sp)
