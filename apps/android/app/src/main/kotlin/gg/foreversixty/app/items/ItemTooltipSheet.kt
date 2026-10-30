package gg.foreversixty.app.items

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.navigationBars
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.windowInsetsPadding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.MicroStyle
import gg.foreversixty.app.common.SecondaryButton
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.GamePalette
import gg.foreversixty.app.model.Item

// MERGE NOTE (lanes F2 and F3): F2's full sheet replaces this file at integration (coordinator
// ruling, 2026-10-31). The entry point is F2's: `ItemTooltipSheet(request, catalog, onDismiss)`,
// with the pure body `ItemTooltipContent(request, phase, retry)`. F3's body here is minimal on
// purpose: no stats, weapon line, Binds, Unique, armor type or set lines.

/** The sheet's phases: the fallback header while the static files load, then the model. */
sealed interface ItemTooltipPhase {
    data object Loading : ItemTooltipPhase
    data class Loaded(val model: ItemTooltipModel) : ItemTooltipPhase
    data object Failed : ItemTooltipPhase
}

/** The item tooltip as a bottom sheet (spec F2.6). It reads only the build's static files. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ItemTooltipSheet(request: ItemTooltipRequest, catalog: ItemCatalog, onDismiss: () -> Unit) {
    val res = LocalResources.current
    var phase by remember(request) { mutableStateOf<ItemTooltipPhase>(ItemTooltipPhase.Loading) }
    var attempt by remember(request) { mutableIntStateOf(0) }
    LaunchedEffect(request, attempt) {
        phase = ItemTooltipPhase.Loading
        val sources = catalog.sources(request)
        phase = if (sources == null) ItemTooltipPhase.Failed else ItemTooltipPhase.Loaded(ItemTooltipModel.build(request, sources, res))
    }
    ModalBottomSheet(
        onDismissRequest = onDismiss,
        sheetState = rememberModalBottomSheetState(),
        containerColor = GamePalette.sheet,
    ) {
        ItemTooltipContent(request, phase, retry = { attempt++ })
    }
}

/** The sheet's body, pure over its phase so the screenshot tests draw it directly. Clears the gesture bar. */
@Composable
fun ItemTooltipContent(request: ItemTooltipRequest, phase: ItemTooltipPhase, retry: () -> Unit = {}) {
    Column(
        Modifier.fillMaxWidth().background(GamePalette.sheet).verticalScroll(rememberScrollState())
            .windowInsetsPadding(WindowInsets.navigationBars)
            .padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(16.dp),
    ) {
        when (phase) {
            ItemTooltipPhase.Loading -> {
                Header(request.fallback, slotKey = null)
                Text(stringResource(R.string.items_loading), style = MaterialTheme.typography.labelMedium, color = FsColor.muted)
            }
            ItemTooltipPhase.Failed -> {
                Header(request.fallback, slotKey = null)
                Text(stringResource(R.string.items_failed), style = MaterialTheme.typography.bodyMedium, color = FsColor.wipe)
                SecondaryButton(stringResource(R.string.common_tryAgain), retry)
            }
            is ItemTooltipPhase.Loaded -> Loaded(phase.model)
        }
    }
}

@Composable
private fun Loaded(model: ItemTooltipModel) {
    Header(model.look, model.slotKey)
    model.item?.let { Panel(it) }
    model.source?.takeIf { it.isNotEmpty() }?.let { source ->
        Section(stringResource(R.string.tooltip_source)) {
            Text(source, style = MaterialTheme.typography.bodyMedium, color = FsColor.strong)
        }
    }
    Section(stringResource(R.string.tooltip_nextBest)) {
        if (model.alternatives.isEmpty()) {
            Text(stringResource(R.string.tooltip_nothingElse), style = MaterialTheme.typography.labelMedium, color = FsColor.muted)
        }
        model.alternatives.forEach { alternative ->
            val quality = alternative.look.quality ?: 1
            Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
                ItemIcon(alternative.look.icon, quality, size = 36.dp)
                Column {
                    Text(alternative.look.name.orEmpty(), style = MaterialTheme.typography.bodyMedium, color = GamePalette.qualityText(quality))
                    Text(alternative.meta, style = MaterialTheme.typography.labelMedium, color = FsColor.muted)
                }
            }
        }
    }
    model.pickFooter?.let { Text(it, style = MaterialTheme.typography.labelMedium, color = FsColor.muted) }
}

@Composable
private fun Header(look: ItemLook, slotKey: String?) {
    val res = LocalResources.current
    val quality = look.quality ?: 1
    val qualityName = ItemWords.quality(res, quality)
    Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
        ItemIcon(look.icon, quality, size = 56.dp)
        Column {
            Text(look.name.orEmpty(), style = MaterialTheme.typography.titleMedium, color = GamePalette.qualityText(quality))
            Text(
                slotKey?.let { stringResource(R.string.bis_qualitySlot, qualityName, ItemWords.slot(res, it)) } ?: qualityName,
                style = MaterialTheme.typography.labelMedium,
                color = FsColor.muted,
            )
        }
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
