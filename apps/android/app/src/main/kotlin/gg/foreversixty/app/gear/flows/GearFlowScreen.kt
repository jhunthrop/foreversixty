package gg.foreversixty.app.gear.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.designsystem.FsSpace
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowPanel
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.PaperdollHeader
import gg.foreversixty.app.designsystem.PaperdollListSlot
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowBandPage
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.FlowSubLink
import gg.foreversixty.app.flows.FlowLinkLine
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemQuality

data class GearActions(
    val onChange: () -> Unit = {},
    val onRetry: () -> Unit = {},
    val onWhere: (GearWhereRow) -> Unit = {},
    val onSlot: (GearSlotRow) -> Unit = {},
)

/**
 * Gear in the round-5 look (spec 2026-11-08 4.3; spec 2026-11-09 Paperdoll "Mobile Gear tab", mock
 * m-android-gear-frame): the band behind [topBar], then the paperdoll strip (the character's
 * [portrait] or the seal, the spec, the set's figure), the line naming the list with Change, every
 * slot in two columns, and Where to go tonight (left out when nothing is worth a trip).
 */
@Composable
fun GearFlowScreen(content: GearContent, build: String, actions: GearActions, portrait: Painter? = null, topBar: @Composable () -> Unit = {}) {
    val equipped = stringResource(R.string.gear_worn)
    val looks = content.slots.mapNotNull { it.look }.associateBy { it.id }
    FlowBandPage(topBar) {
        PaperdollHeader(
            title = content.title,
            figure = content.figure,
            portrait = portrait,
            slots = content.slots.map { it.listSlot(content.unitWord, equipped) },
            unit = content.unitWord,
            onSelect = { id -> content.slots.firstOrNull { it.slot == id }?.let(actions.onSlot) },
            lead = { Lead(content.facts, content.header, actions.onChange) },
        ) { item ->
            val look = looks[item.id]
            ItemIcon(look?.icon, look?.quality, build, SOCKET_ICON, Modifier.parityItem(ParityItemKind.Image))
        }
        if (content.where.isNotEmpty()) {
            FlowSection(stringResource(R.string.gear_whereTonight), top = 0.dp, panel = PanelFrame.None) {
                content.where.forEachIndexed { i, row ->
                    FlowRow(
                        row.title, row.line, FlowLead.Item(row.lead, build),
                        last = i == content.where.lastIndex,
                        onClick = { actions.onWhere(row) },
                        meta = row.figure,
                    )
                }
            }
        }
    }
}

/**
 * Gear with no list for the spec and band (round 2): the spec, the line naming the list with Change, and
 * one row saying why, with Try again when the list did not load.
 */
@Composable
fun GearNoticeScreen(notice: GearNotice, actions: GearActions, topBar: @Composable () -> Unit = {}) {
    FlowBandPage(topBar) {
        Column(Modifier.fillMaxWidth()) {
            FlowHead(notice.title, emptyList())
            Lead(notice.facts, null, actions.onChange)
        }
        FlowBlock(ParityRegion.section(if (notice.retry) "Error" else "Empty"), top = 0.dp) {
            FlowPanel {
                FlowRow(notice.row, notice.sub, last = true, trail = if (notice.retry) {
                    { FlowButton(stringResource(R.string.common_tryAgain), actions.onRetry, Modifier.heightIn(min = FsTarget.min)) }
                } else {
                    null
                })
            }
        }
    }
}

/**
 * The line that names the list, `60 raid · Alliance · 4 of 17 slots match the list · Change`, and under
 * it the measured [header] (`+157.6 DPS from 13 upgrades · synced 2 h ago`).
 */
@Composable
private fun Lead(facts: String, header: String?, onChange: () -> Unit) {
    Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.CONTROLS)) {
        FlowLinkLine(facts, FlowSubLink(stringResource(R.string.gear_change), onChange), LEAD, FsColor.text, Modifier.fillMaxWidth())
        header?.let { FlowNote(it, Modifier.padding(top = FsSpace.space1)) }
    }
}

/** A slot as the paperdoll's list draws it: the pick in its quality colour, its gain or Equipped. */
private fun GearSlotRow.listSlot(unit: String, equipped: String): PaperdollListSlot {
    val pick = look ?: return PaperdollListSlot(slot, label, item = null)
    val detail = when {
        worn -> equipped
        gain != null -> "$gain $unit"
        else -> null
    }
    val item = PaperdollListSlot.Item(pick.id, pick.name.orEmpty(), ItemQuality(pick.quality).text, detail, isGain = gain != null && !worn)
    return PaperdollListSlot(slot, label, item)
}

/** Gear while its list loads. */
@Composable
fun GearLoading() {
    FlowPage {
        FlowSkeleton(240.dp, 20.dp)
        FlowSkeleton(160.dp, 18.dp, Modifier.fillMaxWidth())
    }
}

/** Signed out (mock `gear-out`): the list for a class; a tap opens the picker on that class. */
@Composable
fun GearSignedOut(classes: List<Pair<String, String>>, onClass: (String) -> Unit) {
    FlowPage {
        FlowHead(stringResource(R.string.gear_out_title), listOf(stringResource(R.string.gear_out_meta)))
        FlowSection(stringResource(R.string.gear_out_classes)) {
            classes.forEachIndexed { i, (slug, name) ->
                FlowRow(name, lead = FlowLead.Crest(slug), titleColor = classColor(slug), last = i == classes.lastIndex, onClick = { onClass(slug) })
            }
        }
    }
}

/** The item icon inside the 44 dp socket, as the design-system sheet draws it. */
private val SOCKET_ICON = 38.dp

/** The list's lead line, 14 on 20 (the mock's `.pdl .lead`). */
private val LEAD = FsType.factLine
