package gg.foreversixty.app.bis

import gg.foreversixty.app.designsystem.LoadError
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.surfaceBand
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.screenGround
import android.content.res.Resources
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ArrowDropDown
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.PathEffect
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.onClick
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.pillSurface
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.FsFont
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.SpecIcon
import java.time.Instant

/** Heights the spec fixes (§4.1, F3.1). */
private val ROW_MIN = 72.dp
private val ICON = 44.dp
private const val SKELETON_ROWS = 17

/** The BiS list (spec F3.1), drawn from plain state so screenshot tests render each state. */
@Composable
fun BisContent(
    state: BisUiState,
    catalog: SpecCatalog,
    query: String,
    onOpenPicker: () -> Unit = {},
    onTryAgain: () -> Unit = {},
    onOpenItem: (ItemTooltipRequest) -> Unit = {},
    nowMillis: Long = System.currentTimeMillis(),
) {
    val phase = state.phase
    val offlineSince = when (phase) {
        is BisPhase.Loaded -> phase.offlineSinceMillis
        is BisPhase.NoList -> phase.offlineSinceMillis
        else -> null
    }
    Column(Modifier.fillMaxSize().screenGround()) {
        BandHead(state.context, catalog, state.build.orEmpty(), onOpenPicker)
        offlineSince?.let { saved ->
            OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(Instant.ofEpochMilli(saved), Instant.ofEpochMilli(nowMillis))))
        }
        LazyColumn(
            Modifier.fillMaxSize(),
            contentPadding = PaddingValues(horizontal = 16.dp, vertical = 16.dp),
            verticalArrangement = Arrangement.spacedBy(0.dp),
        ) {
            when (phase) {
                BisPhase.Loading -> item { Skeleton() }
                is BisPhase.Loaded -> {
                    phase.list.setDpsLine?.let { line -> item { SetDpsLine(line) } }
                    val rows = phase.list.filtered(query)
                    if (rows.isEmpty()) {
                        item { Notice(stringResource(R.string.bis_searchEmpty)) }
                    } else {
                        itemsIndexed(rows, key = { _, row -> row.slot }) { index, row ->
                            RowCell(index, rows.size) {
                                val pick = row.pick
                                if (pick != null) {
                                    PickRow(row, pick, state.build.orEmpty()) {
                                        state.build?.let { build -> onOpenItem(tooltipRequest(row, pick, phase.list, state.context, build)) }
                                    }
                                } else {
                                    EmptyRow(row.emptyReason.orEmpty())
                                }
                            }
                        }
                    }
                }
                is BisPhase.NoList -> item { StateNotice(stringResource(R.string.home_answer_noList), FsColor.mutedGround, onTryAgain = null) }
                BisPhase.NotCached -> item { LoadError(stringResource(R.string.bis_notCached), stringResource(R.string.common_tryAgain), onTryAgain, Modifier.padding(top = 10.dp)) }
                BisPhase.Failed -> item { LoadError(stringResource(R.string.bis_failed), stringResource(R.string.common_tryAgain), onTryAgain, Modifier.padding(top = 10.dp)) }
            }
        }
    }
}

/**
 * The list's head on the painted band (the web's /bis page head): the spec's icon, its name in the
 * class colour in the display face and `Best in slot`; then the control bar, one 48 dp row with the
 * selection in words and `Change` (spec 3.4, ControlBar, phone form), which opens the picker.
 */
@Composable
private fun BandHead(context: BisContext, catalog: SpecCatalog, build: String, onOpenPicker: () -> Unit) {
    val res = LocalResources.current
    val selection = context.selection
    val spec = catalog.spec(selection.spec)
    val specLabel = stringResource(R.string.bis_specChip, spec?.name ?: selection.spec, context.className(selection.classSlug))
    // One line in the control bar: the short range (`60 · Phase 1 raid`); the opening date stays in the picker.
    val range = BisLabels.range(res, selection.range)
    val faction = BisLabels.faction(res, selection.faction)
    Column(Modifier.fillMaxWidth().surfaceBand()) {
        Row(
            Modifier.fillMaxWidth().padding(start = 16.dp, end = 16.dp, top = 8.dp, bottom = 12.dp),
            horizontalArrangement = Arrangement.spacedBy(12.dp),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            spec?.let { SpecIcon(it.icon, build, 44.dp) }
            Column(Modifier.weight(1f).semantics(mergeDescendants = true) { heading() }) {
                Text(specLabel, style = ScaledType.barTitle, color = classColor(selection.classSlug))
                Text(stringResource(R.string.bis_title), style = ScaledType.meta, color = FsColor.mutedGround)
            }
        }
        val description = stringResource(R.string.bis_a11y_picker, specLabel, range, faction)
        Row(
            Modifier
                .fillMaxWidth()
                .heightIn(min = 48.dp)
                .clearAndSetSemantics {
                    contentDescription = description
                    role = Role.Button
                    onClick { onOpenPicker(); true }
                }
                .clickable(onClick = onOpenPicker)
                .padding(horizontal = 16.dp),
            horizontalArrangement = Arrangement.spacedBy(8.dp),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            FactionMark(selection.faction, size = 20.dp)
            Text(listOf(range, faction).joinToString(SUMMARY_SEPARATOR), style = ScaledType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
            Text(stringResource(R.string.gear_change), style = ScaledType.bodyStrong, color = FsColor.gold)
            Icon(Icons.Filled.ArrowDropDown, contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(18.dp))
        }
    }
}

private const val SUMMARY_SEPARATOR = " · "

/** A BiS chip (spec F3.1): 40 tall, pill, the 4.5% fill, the card border and a chevron. */
@Composable
fun Chip(label: String, onClick: () -> Unit, leading: @Composable () -> Unit) {
    Row(
        Modifier
            .heightIn(min = 40.dp)
            .pillSurface(GameColors.controlFill, FsColor.lineCard)
            .clickable(onClick = onClick)
            .padding(start = 8.dp, end = 10.dp, top = 4.dp, bottom = 4.dp),
        horizontalArrangement = Arrangement.spacedBy(8.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        leading()
        Text(label, style = MaterialTheme.typography.bodyLarge.copy(fontWeight = FontWeight.SemiBold), color = FsColor.strong)
        Icon(Icons.Filled.ArrowDropDown, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(18.dp))
    }
}

/** `894.1 DPS at best in slot for your level range · Updated 30 Oct 2026`, the figure in mono. */
@Composable
private fun SetDpsLine(line: String) {
    val figureEnd = line.indexOf(' ').coerceAtLeast(0)
    val text = buildAnnotatedString {
        withStyle(ScaledType.tabular.toSpanStyle().copy(color = FsColor.strong)) {
            append(line.substring(0, figureEnd))
        }
        append(line.substring(figureEnd))
    }
    Text(text, style = ScaledType.meta, color = FsColor.mutedGround, modifier = Modifier.padding(bottom = 14.dp))
}

/** One cell of the quiet card: the card's ground, rounded at its ends, divided between rows. */
@Composable
private fun RowCell(index: Int, count: Int, content: @Composable () -> Unit) {
    val radius = FsRadius.legacyPanel
    val shape = RoundedCornerShape(
        topStart = if (index == 0) radius else 0.dp,
        topEnd = if (index == 0) radius else 0.dp,
        bottomStart = if (index == count - 1) radius else 0.dp,
        bottomEnd = if (index == count - 1) radius else 0.dp,
    )
    Column(Modifier.fillMaxWidth().background(FsColor.raised, shape)) {
        if (index > 0) HorizontalDivider(color = FsColor.lineSoft, thickness = 1.dp)
        content()
    }
}

@Composable
private fun PickRow(row: BisRow, pick: BisPick, build: String, onOpen: (Resources) -> Unit) {
    val res = LocalResources.current
    val large = LocalDensity.current.fontScale >= LARGE_FONT_SCALE
    val description = rowDescription(res, row, pick)
    val hint = stringResource(R.string.bis_a11y_showsItem)
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = ROW_MIN)
            .clickable(onClickLabel = hint) { onOpen(res) }
            .clearAndSetSemantics {
                contentDescription = description
                role = Role.Button
                onClick(label = hint) { onOpen(res); true }
            }
            .padding(horizontal = 16.dp, vertical = 12.dp),
        horizontalArrangement = Arrangement.spacedBy(12.dp),
        verticalAlignment = Alignment.Top,
    ) {
        ItemIcon(pick.item.icon, pick.item.quality, build, ICON)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalAlignment = Alignment.Top) {
                Text(pick.item.name, style = MaterialTheme.typography.bodyLarge, color = ItemQuality(pick.item.quality).text, modifier = Modifier.weight(1f))
                if (pick.isWorn && !large) WornTag()
            }
            Meta(pick.slotSource)
            Meta(pick.next)
            if (pick.isUnmodelled) Meta(stringResource(R.string.home_upgrades_notModelled))
            if (pick.isWorn && large) {
                Spacer(Modifier.height(4.dp))
                WornTag()
            }
        }
    }
}

/** From font scale 1.5 the worn tag leaves the name's line, so the name keeps the width. */
private const val LARGE_FONT_SCALE = 1.5f

@Composable
private fun WornTag() = Tag(stringResource(R.string.bis_youWear), FsColor.kill, checked = true)

@Composable
private fun Meta(text: String) {
    Text(text, style = MaterialTheme.typography.labelMedium.copy(fontWeight = FontWeight.Normal), color = FsColor.muted)
}

/** `Head. Blue Dragonscale Helm, Epic, item level 61. Crafted: Leatherworking. Next: …` */
private fun rowDescription(res: Resources, row: BisRow, pick: BisPick): String {
    val parts = mutableListOf("${row.slotLabel}.")
    val item = pick.item.item
    parts += if (item != null) {
        res.getString(R.string.bis_a11y_item, pick.item.name, ItemQuality(item.quality).name?.let(res::getString).orEmpty(), item.itemLevel)
    } else {
        "${pick.item.name}."
    }
    if (pick.source.isNotEmpty()) parts += "${pick.source}."
    parts += "${pick.next}."
    if (pick.isUnmodelled) parts += "${res.getString(R.string.home_upgrades_notModelled)}."
    if (pick.isWorn) parts += "${res.getString(R.string.bis_youWear)}."
    return parts.joinToString(" ")
}

@Composable
private fun EmptyRow(reason: String) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = ROW_MIN).padding(horizontal = 16.dp, vertical = 12.dp),
        horizontalArrangement = Arrangement.spacedBy(12.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        DashedBox(ICON)
        Text(reason, style = MaterialTheme.typography.labelMedium.copy(fontWeight = FontWeight.Normal), color = FsColor.muted, modifier = Modifier.weight(1f))
    }
}

@Composable
private fun DashedBox(size: Dp) {
    val line = FsColor.line
    Box(
        Modifier.size(size).drawBehind {
            drawRoundRect(
                color = line,
                cornerRadius = androidx.compose.ui.geometry.CornerRadius(FsRadius.legacyPill.toPx()),
                style = Stroke(width = 1.dp.toPx(), pathEffect = PathEffect.dashPathEffect(floatArrayOf(6f, 6f))),
            )
        },
    )
}

/** 17 rows at their real height while nothing is saved (spec §5.0 Loading). */
@Composable
private fun Skeleton() {
    val loading = stringResource(R.string.bis_loading)
    Column(
        Modifier
            .fillMaxWidth()
            .background(FsColor.raised, RoundedCornerShape(FsRadius.legacyPanel))
            .semantics(mergeDescendants = true) {
                contentDescription = loading
                liveRegion = LiveRegionMode.Polite
            },
    ) {
        repeat(SKELETON_ROWS) { index ->
            if (index > 0) HorizontalDivider(color = FsColor.lineSoft, thickness = 1.dp)
            Row(
                Modifier.fillMaxWidth().heightIn(min = ROW_MIN).padding(horizontal = 16.dp, vertical = 12.dp),
                horizontalArrangement = Arrangement.spacedBy(12.dp),
            ) {
                Box(Modifier.size(ICON).background(FsColor.cardTop, RoundedCornerShape(FsRadius.legacyPill)))
                Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
                    Bar(170.dp, 12.dp)
                    Bar(220.dp, 10.dp)
                    Bar(190.dp, 10.dp)
                }
            }
        }
    }
}

@Composable
private fun Bar(width: Dp, height: Dp) =
    Box(Modifier.width(width).height(height).background(FsColor.cardTop, RoundedCornerShape(percent = 50)))

@Composable
private fun Notice(text: String) {
    Text(text, style = MaterialTheme.typography.bodyLarge, color = FsColor.text, modifier = Modifier.padding(top = 8.dp))
}

/**
 * A state of the list region (not saved, no list, error): its sentence, 24 below the chips (the
 * chips' 14 plus 10), and `Try again` where a retry can help (spec §5.0).
 */
@Composable
private fun StateNotice(text: String, color: Color, onTryAgain: (() -> Unit)?) {
    Column(Modifier.padding(top = 10.dp), verticalArrangement = Arrangement.spacedBy(16.dp)) {
        Text(
            text,
            style = MaterialTheme.typography.bodyLarge,
            color = color,
            modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite },
        )
        onTryAgain?.let { BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, it, height = 48.dp) }
    }
}

/** The tooltip request of a list row, in F2's entry point's terms (`ItemTooltipSheet(request, catalog, onDismiss)`). */
fun tooltipRequest(row: BisRow, pick: BisPick, list: BisList, context: BisContext, build: String): ItemTooltipRequest {
    val selection = context.selection
    val item = pick.item.item
    return ItemTooltipRequest(
        itemId = pick.item.id,
        fallback = ItemLook(pick.item.id, pick.item.name, item?.quality, item?.icon, item?.itemLevel),
        build = build,
        classSlug = selection.classSlug,
        spec = selection.spec,
        specName = list.specLabel,
        band = selection.range.band,
        bandLabel = list.rangeLabel,
        preset = selection.range.preset,
        faction = selection.faction,
        sourceLine = pick.source,
        wornId = pick.runnersUp.firstOrNull { it.isWorn }?.item?.id ?: pick.item.id.takeIf { pick.isWorn },
    )
}
