package gg.foreversixty.app.guides.ui

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
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.HeadingFamily
import gg.foreversixty.app.designsystem.MicroLabel
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
                Loadable.NotCached -> Text(stringResource(R.string.guides_listNotCached), style = FsType.body, color = FsColor.muted, modifier = Modifier.padding(top = GuideSpace.rowGap))
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
        QuietRows(rows) { row -> TierRowView(row, rows.size, file.build) { actions.open(row.row.spec) } }
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
                MicroLabel(stringResource(R.string.tiers_yours), color = FsColor.gold)
                Text(callout.row.title, style = FsType.bodyStrong, color = FsColor.strong)
                Text(callout.row.detail, style = FsType.meta, color = FsColor.muted)
            }
        }
        callout.note?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
    }
}

/** One row (68 min): rank, icon, name, race and bar; the tier chip and gap (or the metric) at the right. */
@Composable
private fun TierRowView(model: TierRowModel, count: Int, build: String, open: () -> Unit) {
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
        Text(row.rank.toString(), style = FsType.mono, color = FsColor.muted, modifier = Modifier.widthIn(min = RANK_WIDTH))
        SpecTile(row.icon, row.classSlug, build, ROW_ICON)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(4.dp)) {
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalAlignment = Alignment.CenterVertically) {
                Text(row.name, style = FsType.bodyStrong, color = classColor(row.classSlug), modifier = Modifier.weight(1f, fill = false))
                if (model.isYours) Tag(stringResource(R.string.tiers_yoursTag), FsColor.gold)
            }
            Text(row.race, style = FsType.meta, color = FsColor.muted)
            Bar(row.fraction.toFloat(), toneColor(model.tone) ?: classColor(row.classSlug))
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
            TierChip(row.tier.orEmpty(), tone)
            Text(model.gap, style = FsType.mono, color = FsColor.text, textAlign = TextAlign.End)
        } else {
            Text(row.metricText, style = FsType.mono, color = FsColor.text)
            Text(row.unit, style = FsType.meta, color = FsColor.muted)
        }
    }
    if (inline) {
        Row(horizontalArrangement = Arrangement.spacedBy(GuideSpace.labelGap), verticalAlignment = Alignment.CenterVertically) { parts() }
    } else {
        // A fixed minimum width keeps every row's bar on the same scale.
        Column(Modifier.widthIn(min = FIGURE_WIDTH), horizontalAlignment = Alignment.End, verticalArrangement = Arrangement.spacedBy(6.dp)) { parts() }
    }
}

/** The tier chip (spec §4): 28 dp, pill radius, a 1.5 dp border and the letter in the tier colour. */
@Composable
private fun TierChip(letter: String, tone: TierTone) {
    val color = toneColor(tone) ?: FsColor.muted
    val shape = RoundedCornerShape(percent = PILL_PERCENT)
    // 28 dp, growing with the letter at large text so it is never cut.
    Box(Modifier.defaultMinSize(CHIP, CHIP).border(1.5.dp, color, shape).padding(horizontal = 4.dp), contentAlignment = Alignment.Center) {
        Text(letter, style = FsType.bodyStrong.copy(fontFamily = HeadingFamily), color = color)
    }
}

/** The 8 dp bar: the fraction of the top spec, in the tier colour. */
@Composable
private fun Bar(fraction: Float, color: Color) {
    Box(Modifier.fillMaxWidth(BAR_WIDTH).height(8.dp).clip(CircleShape).background(FsColor.lineSoft)) {
        Box(Modifier.fillMaxHeight().fillMaxWidth(fraction.coerceIn(0f, 1f)).clip(CircleShape).background(color))
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
        Text(stringResource(R.string.tiers_loading), style = FsType.meta, color = FsColor.muted, modifier = Modifier.cardInside())
    }
}

private val ROW_MIN = 68.dp
private val ROW_ICON = 36.dp
private val CALLOUT_ICON = 40.dp
private val FACTION_MARK = 20.dp
private val CHIP = 28.dp
private val RANK_WIDTH = 20.dp
private val FIGURE_WIDTH = 72.dp
private val GOLD_EDGE = 3.dp
private const val BAR_WIDTH = 0.85f
private const val PILL_PERCENT = 50
private const val SKELETON_ROWS = 6
