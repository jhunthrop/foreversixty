package gg.foreversixty.app.guides.ui

import androidx.compose.ui.unit.sp
import androidx.compose.foundation.layout.width
import gg.foreversixty.app.designsystem.Label12
import androidx.compose.material.icons.filled.Share
import androidx.compose.material3.IconButton
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.defaultMinSize
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ArrowDropDown
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.onClick
import androidx.compose.ui.semantics.role
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.FsFont
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.guides.logic.TierCallout
import gg.foreversixty.app.guides.logic.TierChoice
import gg.foreversixty.app.guides.logic.TierRowModel
import gg.foreversixty.app.guides.logic.TierTone
import gg.foreversixty.app.guides.logic.Tiers
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.model.TiersFile
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.SkeletonBar
import gg.foreversixty.app.reports.ui.isLargeText
import java.time.Instant

data class TierListState(
    val tiers: Loadable<TiersFile>,
    val choice: TierChoice,
    /** The current character's spec id; null with no character (no callout, no `Yours`). */
    val yourSpec: String?,
)

class TierListActions(
    val choose: (TierChoice) -> Unit = {},
    val open: (String) -> Unit = {},
    val retry: () -> Unit = {},
    /** Lane F15: shares the role's tier list page; null draws no share button. */
    val share: ((TierChoice) -> Unit)? = null,
)

/** F9.4: role and faction chips, the callout for your spec, then the file's rows as they come. */
@Composable
fun TierListContent(state: TierListState, now: Instant, actions: TierListActions) {
    Column(Modifier.fillMaxWidth()) {
        (state.tiers as? Loadable.Loaded)?.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now))) }
        Column(Modifier.screenMargins().padding(top = GuideSpace.labelGap)) {
            Chips(state.choice, actions.choose, actions.share)
            when (val tiers = state.tiers) {
                Loadable.Loading -> TierSkeleton()
                Loadable.NotCached -> Text(stringResource(R.string.guides_listNotCached), style = ScaledType.body, color = FsColor.mutedGround, modifier = Modifier.padding(top = GuideSpace.rowGap))
                Loadable.Failed -> ErrorRegion(stringResource(R.string.tiers_error), actions.retry, inset = false)
                is Loadable.Loaded -> Loaded(tiers.value, state, actions)
            }
        }
    }
}

@Composable
private fun Loaded(file: TiersFile, state: TierListState, actions: TierListActions) {
    val list = Tiers.list(file, state.choice)
    val rows = Tiers.rows(list, state.yourSpec)
    Column(Modifier.padding(top = GuideSpace.rowGap, bottom = GuideSpace.section), verticalArrangement = Arrangement.spacedBy(GuideSpace.rowGap)) {
        Tiers.callout(list, state.yourSpec)?.let { Callout(it, file.build) }
        // The web's anatomy: the tier letter once at the head of its group, then rank, crest, name and the figures.
        rows.firstOrNull()?.row?.unit?.takeIf { it.isNotEmpty() }?.let { unit ->
            Box(Modifier.fillMaxWidth().padding(end = 16.dp), contentAlignment = Alignment.CenterEnd) { Label12(unit, color = FsColor.mutedGround) }
        }
        val heads = rows.indices.filter { i -> i == 0 || rows[i - 1].row.tier != rows[i].row.tier }.toSet()
        QuietRows(rows.withIndex().toList()) { (i, row) ->
            TierRowView(row, rows.size, file.build, letter = row.row.tier.takeIf { i in heads }) { actions.open(row.row.spec) }
        }
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
private fun Chips(choice: TierChoice, choose: (TierChoice) -> Unit, share: ((TierChoice) -> Unit)?) {
    val copy = GuidesCopy(LocalResources.current)
    var factionMenu by remember { mutableStateOf(false) }
    FlowRow(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
        TierChoice.roles.forEach { role ->
            PillChip(copy.role(role), selected = role == choice.role, onClick = { choose(choice.copy(role = role)) })
        }
        Box {
            PillChip(
                copy.faction(choice.faction), selected = false, onClick = { factionMenu = true },
                leading = { FactionMark(choice.faction, FACTION_MARK) },
                trailing = { Icon(Icons.Filled.ArrowDropDown, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(18.dp)) },
            )
            DropdownMenu(expanded = factionMenu, onDismissRequest = { factionMenu = false }) {
                TierChoice.factions.forEach { faction ->
                    DropdownMenuItem(text = { Text(copy.faction(faction)) }, leadingIcon = { FactionMark(faction, FACTION_MARK) }, onClick = {
                        factionMenu = false
                        choose(choice.copy(faction = faction))
                    })
                }
            }
        }
        share?.let {
            IconButton(onClick = { it(choice) }) {
                Icon(Icons.Filled.Share, contentDescription = stringResource(R.string.share_action), tint = FsColor.strong)
            }
        }
    }
}

/** Your spec, first (the one warm card): `Fire Mage is 6th of 20 DPS specs.`, its gap, and the race line. */
@Composable
private fun Callout(callout: TierCallout, build: String) {
    Column(verticalArrangement = Arrangement.spacedBy(GuideSpace.labelGap)) {
        Row(
            Modifier.fillMaxWidth().primaryCard().cardInside(),
            horizontalArrangement = Arrangement.spacedBy(GuideSpace.inCard),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            SpecTile(callout.row.icon, callout.row.classSlug, build, CALLOUT_ICON)
            Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
                Label12(stringResource(R.string.tiers_yours), color = FsColor.gold)
                Text(callout.row.title, style = ScaledType.bodyStrong, color = FsColor.strong)
                Text(callout.row.detail, style = ScaledType.meta, color = FsColor.muted)
            }
        }
        callout.note?.let { Text(it, style = ScaledType.meta, color = FsColor.mutedGround) }
    }
}

/** One row (68 min): the tier [letter] at its group's head, rank, icon, name and race; the metric over the gap at the right. */
@Composable
private fun TierRowView(model: TierRowModel, count: Int, build: String, letter: String?, open: () -> Unit) {
    val copy = GuidesCopy(LocalResources.current)
    val row = model.row
    val label = copy.tierLabel(model, count)
    val opens = stringResource(R.string.tiers_a11y_opens)
    val stacked = isLargeText()
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = ROW_MIN)
            .then(if (model.isYours) Modifier.yoursGround() else Modifier)
            .clickable(onClickLabel = opens, role = Role.Button, onClick = open)
            .clearAndSetSemantics {
                contentDescription = label
                role = Role.Button
                onClick(label = opens) { open(); true }
            }
            .cardInside(),
        horizontalArrangement = Arrangement.spacedBy(GuideSpace.rowGap),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        if (model.tone != null) {
            Box(Modifier.width(LETTER_WIDTH)) {
                letter?.let { Text(it, style = LETTER, color = toneColor(model.tone) ?: FsColor.muted) }
            }
        }
        Text(row.rank.toString(), style = ScaledType.tabular, color = FsColor.muted, modifier = Modifier.widthIn(min = RANK_WIDTH))
        ClassCrest(row.classSlug, ROW_ICON)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(4.dp)) {
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalAlignment = Alignment.CenterVertically) {
                Text(row.name, style = ScaledType.bodyStrong, color = classColor(row.classSlug), modifier = Modifier.weight(1f, fill = false))
                if (model.isYours) Tag(stringResource(R.string.tiers_yoursTag), FsColor.gold)
            }
            Text(row.race, style = ScaledType.meta, color = FsColor.muted)
            if (stacked) Figure(model, inline = true)
        }
        if (!stacked) Figure(model, inline = false)
    }
}

/**
 * The tier chip and the gap, or for tank and healer lists (no letters) the metric and its unit: a
 * right column, or at the largest text one line under the name ([inline], spec W3 §9).
 */
@Composable
private fun Figure(model: TierRowModel, inline: Boolean) {
    val parts = @Composable {
        val row = model.row
        val tone = model.tone
        if (tone != null && row.tier != null) {
            Text(row.metricText, style = ScaledType.tabular, color = FsColor.strong, textAlign = TextAlign.End)
            Text(model.gap, style = ScaledType.tabularMeta, color = FsColor.muted, textAlign = TextAlign.End)
        } else {
            Text(row.metricText, style = ScaledType.tabular, color = FsColor.text)
            Text(row.unit, style = ScaledType.meta, color = FsColor.muted)
        }
    }
    if (inline) {
        Row(horizontalArrangement = Arrangement.spacedBy(GuideSpace.labelGap), verticalAlignment = Alignment.CenterVertically) { parts() }
    } else {
        // A fixed minimum width keeps every row's bar on the same scale.
        Column(Modifier.widthIn(min = FIGURE_WIDTH), horizontalAlignment = Alignment.End, verticalArrangement = Arrangement.spacedBy(6.dp)) { parts() }
    }
}

/** A spec icon, or its class crest when the spec has none (never a letter tile). */
@Composable
private fun SpecTile(icon: String?, classSlug: String, build: String, size: androidx.compose.ui.unit.Dp) {
    if (icon != null) SpecIcon(icon, build, size) else ClassCrest(classSlug, size)
}

/** Your row: the row ground and a 3 dp gold left edge. */
private fun Modifier.yoursGround(): Modifier = background(GameColors.rowWash)
    .drawBehind { drawRect(FsColor.gold, size = Size(GOLD_EDGE.toPx(), size.height)) }

/** The app's rarity tokens, as the web's tier page maps them. */
private fun toneColor(tone: TierTone?): Color? = when (tone) {
    TierTone.Legendary -> FsColor.rarityLegendary
    TierTone.Epic -> FsColor.rarityEpicText
    TierTone.Rare -> FsColor.rarityRareText
    TierTone.Uncommon -> FsColor.rarityUncommon
    TierTone.Muted -> FsColor.muted
    null -> null
}

@Composable
private fun TierSkeleton() {
    Column(Modifier.padding(top = GuideSpace.rowGap).fillMaxWidth().quietCard()) {
        repeat(SKELETON_ROWS) {
            Row(Modifier.fillMaxWidth().heightIn(min = ROW_MIN).cardInside(), horizontalArrangement = Arrangement.spacedBy(GuideSpace.rowGap), verticalAlignment = Alignment.CenterVertically) {
                SkeletonBar(ROW_ICON, ROW_ICON)
                Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
                    SkeletonBar(160.dp, 14.dp)
                    SkeletonBar(200.dp, 8.dp)
                }
            }
        }
        Text(stringResource(R.string.tiers_loading), style = ScaledType.meta, color = FsColor.muted, modifier = Modifier.cardInside())
    }
}

private val ROW_MIN = 68.dp
private val ROW_ICON = 36.dp
private val CALLOUT_ICON = 40.dp
private val FACTION_MARK = 20.dp
private val RANK_WIDTH = 20.dp
private val FIGURE_WIDTH = 72.dp
private val LETTER_WIDTH = 24.dp

/** The tier letter: the display face, as the web's tier column. */
private val LETTER = ScaledType.sectionTitle.copy(fontSize = 24.sp, lineHeight = 28.sp)
private val GOLD_EDGE = 3.dp
private const val SKELETON_ROWS = 6
