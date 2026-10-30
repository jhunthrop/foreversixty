package gg.foreversixty.app.items

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.MicroStyle
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.GamePalette
import gg.foreversixty.app.model.Item

// MERGE NOTE (lanes F2 and F3): this is F3's minimal item tooltip sheet (spec F2.6). Lane F2
// builds the full one in parallel; the two merge into this one file, F2's body winning. F3 reads
// only `ItemTooltipContent` and `ItemTooltipBody(content)`, shown inside a ModalBottomSheet.

/** A runner-up as the sheet's `The next best` lists it. */
data class ItemTooltipAlternative(
    val id: Int,
    val name: String,
    val icon: String?,
    val quality: Int,
    /** `−1.3 DPS vs the pick`, led by `You wear this · ` for the worn one. */
    val meta: String,
)

/** Everything the sheet draws, already worded by the screen that opens it. */
data class ItemTooltipContent(
    val id: Int,
    val name: String,
    /** The build's item facts; null when the class's items file is not on the phone. */
    val item: Item?,
    val slotKey: String,
    val source: String?,
    val nextBest: List<ItemTooltipAlternative>,
    /** `Best in slot for Head · Fury Warrior 60 · Raid-ready, Phase 1`, when a pick. */
    val pickOf: String?,
) {
    val quality: Int get() = item?.quality ?: 1
}

/** The item tooltip (spec F2.6): header, the game-ordered panel, where it comes from, the next best, the pick footer. */
@Composable
fun ItemTooltipBody(content: ItemTooltipContent, modifier: Modifier = Modifier) {
    val res = LocalResources.current
    Column(
        modifier.background(GamePalette.sheet).verticalScroll(rememberScrollState()).padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(16.dp),
    ) {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
            ItemIcon(content.item?.icon, content.quality, size = 56.dp)
            Column {
                Text(content.name, style = MaterialTheme.typography.titleMedium, color = GamePalette.qualityText(content.quality))
                Text(
                    stringResource(R.string.bis_qualitySlot, ItemWords.quality(res, content.quality), ItemWords.slot(res, content.slotKey)),
                    style = MaterialTheme.typography.labelMedium,
                    color = FsColor.muted,
                )
            }
        }
        content.item?.let { Panel(it) }
        content.source?.takeIf { it.isNotEmpty() }?.let { source ->
            Section(stringResource(R.string.tooltip_source)) {
                Text(source, style = MaterialTheme.typography.bodyMedium, color = FsColor.strong)
            }
        }
        Section(stringResource(R.string.tooltip_nextBest)) {
            if (content.nextBest.isEmpty()) {
                Text(stringResource(R.string.tooltip_nothingElse), style = MaterialTheme.typography.labelMedium, color = FsColor.muted)
            }
            content.nextBest.forEach { alternative ->
                Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
                    ItemIcon(alternative.icon, alternative.quality, size = 36.dp)
                    Column {
                        Text(alternative.name, style = MaterialTheme.typography.bodyMedium, color = GamePalette.qualityText(alternative.quality))
                        Text(alternative.meta, style = MaterialTheme.typography.labelMedium, color = FsColor.muted)
                    }
                }
            }
        }
        content.pickOf?.let { Text(it, style = MaterialTheme.typography.labelMedium, color = FsColor.muted) }
    }
}

@Composable
private fun Panel(item: Item) {
    val res = LocalResources.current
    val shape = RoundedCornerShape(FsRadius.control)
    val style = MaterialTheme.typography.bodyMedium
    Column(
        Modifier.fillMaxWidth().background(GamePalette.tooltipPanel, shape).border(1.dp, GamePalette.iconLine, shape)
            .padding(horizontal = 14.dp, vertical = 12.dp),
    ) {
        Text(item.name, style = style, color = GamePalette.qualityText(item.quality))
        Text(ItemWords.slot(res, item.slot), style = style, color = FsColor.strong)
        if (item.armor > 0) Text(stringResource(R.string.tooltip_armor, item.armor), style = style, color = FsColor.strong)
        if (item.effectText.isNotEmpty()) Text(item.effectText, style = style, color = GamePalette.gameGreen)
        if (item.requiredLevel > 0) {
            Text(stringResource(R.string.tooltip_requires, item.requiredLevel), style = style, color = FsColor.strong)
        }
        Text(stringResource(R.string.tooltip_itemLevel, item.itemLevel), style = style, color = GamePalette.itemLevel)
    }
}

@Composable
private fun Section(title: String, body: @Composable () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        Text(title.uppercase(), style = MicroStyle, color = FsColor.muted, modifier = Modifier.semantics { heading() })
        body()
    }
}
