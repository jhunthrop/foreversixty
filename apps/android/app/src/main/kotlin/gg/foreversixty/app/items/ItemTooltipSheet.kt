package gg.foreversixty.app.items

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.HorizontalDivider
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
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.home.HomeType
import gg.foreversixty.app.home.MicroLabel
import gg.foreversixty.app.home.SecondaryButton

/** The sheet's phases: the fallback header while the static files load, then the model. */
sealed interface ItemTooltipPhase {
    data object Loading : ItemTooltipPhase
    data class Loaded(val model: ItemTooltipModel) : ItemTooltipPhase
    data object Failed : ItemTooltipPhase
}

/**
 * The item tooltip as a bottom sheet (spec F2.6), opening partly expanded and dragging to full.
 * It reads only the build's static files, cached per build, so it needs no API request.
 */
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
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = false),
        containerColor = FsColor.raised,
    ) {
        ItemTooltipContent(request, phase, retry = { attempt++ })
    }
}

/** The sheet's body, pure over its phase so the screenshot tests draw it directly. */
@Composable
fun ItemTooltipContent(request: ItemTooltipRequest, phase: ItemTooltipPhase, retry: () -> Unit = {}) {
    Column(
        Modifier.fillMaxWidth().background(FsColor.raised).verticalScroll(rememberScrollState())
            .padding(start = 16.dp, end = 16.dp, top = 8.dp, bottom = 32.dp),
        verticalArrangement = Arrangement.spacedBy(16.dp),
    ) {
        val fallbackMeta = ItemQuality(request.fallback.quality).name?.let { stringResource(it) }.orEmpty()
        when (phase) {
            ItemTooltipPhase.Loading -> {
                Header(request.fallback, fallbackMeta, request.build)
                Text(stringResource(R.string.items_loading), style = HomeType.meta, color = FsColor.muted)
            }
            ItemTooltipPhase.Failed -> {
                Header(request.fallback, fallbackMeta, request.build)
                Text(stringResource(R.string.items_failed), style = HomeType.body, color = FsColor.wipe)
                SecondaryButton(stringResource(R.string.common_tryAgain), onClick = retry, withRetryGlyph = true)
            }
            is ItemTooltipPhase.Loaded -> Loaded(phase.model, request.build)
        }
    }
}

@Composable
private fun Loaded(model: ItemTooltipModel, build: String) {
    Header(model.look, model.headerMeta, build)
    ItemTooltipPanel(model)
    model.source?.let { source ->
        Column(Modifier.semantics(mergeDescendants = true) {}, verticalArrangement = Arrangement.spacedBy(6.dp)) {
            MicroLabel(stringResource(R.string.tooltip_source))
            Text(source, style = HomeType.body, color = FsColor.strong)
        }
    }
    if (model.pickFooter != null || model.alternatives.isNotEmpty()) {
        Column {
            MicroLabel(stringResource(R.string.tooltip_nextBest))
            Spacer(Modifier.height(8.dp))
            if (model.alternatives.isEmpty()) {
                Text(stringResource(R.string.tooltip_nothingElse), style = HomeType.meta, color = FsColor.muted)
            }
            model.alternatives.forEach { alt ->
                Row(
                    Modifier.fillMaxWidth().padding(vertical = 10.dp).semantics(mergeDescendants = true) {},
                    verticalAlignment = Alignment.CenterVertically,
                    horizontalArrangement = Arrangement.spacedBy(12.dp),
                ) {
                    ItemIcon(alt.look.icon, alt.look.quality, build, 36.dp)
                    Column {
                        Text(alt.look.name ?: stringResource(R.string.items_notInData), style = HomeType.body, color = ItemQuality(alt.look.quality).text)
                        Text(alt.meta, style = HomeType.meta, color = FsColor.muted)
                    }
                }
                HorizontalDivider(color = FsColor.lineSoft)
            }
        }
    }
    model.pickFooter?.let { Text(it, style = HomeType.meta, color = FsColor.muted) }
}

@Composable
private fun Header(look: ItemLook, meta: String, build: String) {
    Row(
        Modifier.semantics(mergeDescendants = true) { heading() },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ItemIcon(look.icon, look.quality, build, 56.dp)
        Column {
            Text(look.name ?: stringResource(R.string.items_notInData), style = HomeType.sectionTitle, color = ItemQuality(look.quality).text)
            Text(meta, style = HomeType.meta, color = FsColor.muted)
        }
    }
}

/** The game-tooltip panel: near-black, `#3a4256` hairline, the client's line order. */
@Composable
fun ItemTooltipPanel(model: ItemTooltipModel) {
    val shape = RoundedCornerShape(FsRadius.control)
    Column(
        Modifier.fillMaxWidth().background(PANEL, shape).border(1.dp, PANEL_LINE, shape)
            .padding(horizontal = 14.dp, vertical = 12.dp).semantics(mergeDescendants = true) {},
        verticalArrangement = Arrangement.spacedBy(4.dp),
    ) {
        Text(
            model.look.name ?: stringResource(R.string.items_notInData),
            style = HomeType.tooltip.copy(fontWeight = FontWeight.Bold, fontSize = 17.sp),
            color = ItemQuality(model.look.quality).text,
        )
        if (!model.known) Line(stringResource(R.string.items_notInData), FsColor.muted)
        model.binds?.let { Line(it) }
        if (model.unique) Line(stringResource(R.string.tooltip_unique))
        if (model.slotName != null || model.typeName != null) Split(model.slotName.orEmpty(), model.typeName.orEmpty())
        model.armor?.let { Line(it) }
        val damage = model.weaponDamage
        val speed = model.weaponSpeed
        if (damage != null && speed != null) Split(damage, speed)
        model.weaponDps?.let { Line(it) }
        model.stats.forEach { Line(it) }
        model.effect?.let { Line(it, FsColor.parseGreen) }
        model.setName?.let {
            Spacer(Modifier.height(4.dp))
            Line(it, FsColor.parseGold)
            model.setBonuses.forEach { bonus -> Line(bonus, FsColor.muted) }
        }
        model.requiresLevel?.let {
            Spacer(Modifier.height(4.dp))
            Line(it)
        }
        model.itemLevel?.let { Line(it, FsColor.parseGold) }
    }
}

/** A left and right pair on one line, wrapping to two lines when the text is large. */
@Composable
private fun Split(left: String, right: String) {
    FlowRow(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
        Text(left, style = HomeType.tooltip, color = FsColor.strong)
        Text(right, style = HomeType.tooltip, color = FsColor.strong)
    }
}

@Composable
private fun Line(text: String, color: Color = FsColor.strong) {
    Text(text, style = HomeType.tooltip, color = color)
}

private val PANEL = Color(0xFF0A0D14)
private val PANEL_LINE = Color(0xFF3A4256)
