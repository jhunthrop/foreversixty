package gg.foreversixty.app.gear.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowItemRow
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import androidx.compose.ui.unit.dp

data class GearActions(
    val onChange: () -> Unit = {},
    val onWhere: (GearWhereRow) -> Unit = {},
    val onSlot: (GearSlotRow) -> Unit = {},
)

/**
 * Gear (spec 2026-11-08 4.3, mock `gear`): the list it compares against with Change, Where to go
 * tonight (left out when nothing is worth a trip), then every slot, Worn or the gain.
 */
@Composable
fun GearFlowScreen(content: GearContent, build: String, actions: GearActions) {
    FlowPage {
        Column(Modifier.fillMaxWidth().parityRegion(CONTROLS)) {
            FlowRow(content.listTitle, content.onList, last = true, trail = {
                FlowLink(stringResource(R.string.gear_change), onClick = actions.onChange)
            })
        }
        if (content.where.isNotEmpty()) {
            FlowSection(stringResource(R.string.gear_whereTonight)) {
                content.where.forEachIndexed { i, row ->
                    FlowRow(
                        row.title, row.line, FlowLead.Item(row.lead, build),
                        last = i == content.where.lastIndex,
                        onClick = { actions.onWhere(row) },
                    )
                }
            }
        }
        FlowSection(stringResource(R.string.gear_everySlot)) {
            Column(Modifier.fillMaxWidth().parityRegion(SLOTS)) {
                content.slots.forEach { slot ->
                    FlowItemRow(
                        look = slot.look,
                        build = build,
                        source = slot.source,
                        gain = if (slot.worn) stringResource(R.string.gear_worn) else slot.gain,
                        unit = slot.gain?.takeIf { !slot.worn }?.let { content.unitWord },
                    ) { actions.onSlot(slot) }
                }
            }
        }
    }
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

private const val CONTROLS = "controls"
private const val SLOTS = "slots"
