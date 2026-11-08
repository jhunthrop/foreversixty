package gg.foreversixty.app.gear.flows

import android.content.res.Resources
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisListBuilder
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.flows.FlowItemRow
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.gear.ItemSource
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.model.ItemNeeds

/** One line of the sheet's lists: a lead, the title in its colour, its line. */
data class GearSheetLine(val title: String, val line: String, val lead: FlowLead, val titleColor: Color? = null)

/** The item sheet (spec F5, mock `item-sheet`): the item, the client's tooltip lines, where, the runner-up, the guild. */
data class GearSheetContent(
    val look: ItemLook,
    val slot: String,
    val gain: String?,
    val unitWord: String,
    /** The tooltip's lines in the client's order; green ones are the item's effects. */
    val tooltip: List<Pair<String, Boolean>>,
    val where: List<GearSheetLine>,
    val runnerUp: GearSheetLine?,
    /** Null while the guild's answer loads, or for a player with no guild (the section is left out). */
    val guild: List<GearSheetLine>?,
    val guildNote: String?,
)

/** Words the sheet from what Gear already holds and the guild's needs answer. */
class GearSheetComposer(private val res: Resources, private val build: String) {
    fun content(
        row: GearSlotRow,
        unitWord: String,
        tooltip: ItemTooltipModel?,
        sources: List<ItemSource>,
        needs: ItemNeeds?,
    ): GearSheetContent {
        val pick = requireNotNull(row.row.pick) { "a sheet opens on a pick" }
        val lead = FlowLead.Item(row.look, build)
        return GearSheetContent(
            look = row.look,
            slot = row.row.slotLabel,
            gain = row.gain,
            unitWord = unitWord,
            tooltip = tooltip?.let(::lines).orEmpty(),
            where = sources.map { GearSheetLine(sourceTitle(it), sourceLine(it), lead) },
            runnerUp = pick.runnersUp.firstOrNull()?.let { alt ->
                val delta = BisListBuilder.signedDelta(alt.dpsDelta)?.let { "$it $unitWord" }
                val look = ItemLook(alt.item.id, alt.item.name, alt.item.quality, alt.item.icon)
                GearSheetLine(
                    alt.item.name,
                    listOfNotNull(alt.source.takeIf { it.isNotBlank() }, delta).joinToString(SEPARATOR),
                    FlowLead.Item(look, build),
                    ItemQuality(alt.item.quality).text,
                )
            },
            guild = needs?.let(::guild),
            guildNote = needs?.let { guildNote(it) },
        )
    }

    private fun lines(model: ItemTooltipModel): List<Pair<String, Boolean>> = listOfNotNull(
        model.binds,
        listOfNotNull(model.slotName, model.typeName).joinToString(SEPARATOR).ifBlank { null },
        model.armor,
        listOfNotNull(model.weaponDamage, model.weaponSpeed).joinToString(SEPARATOR).ifBlank { null },
        model.weaponDps,
    ).map { it to false } + model.stats.map { it to false } + model.equipLines.map { it to true } +
        listOfNotNull(model.effect).map { it to true } + listOfNotNull(model.requiresLevel).map { it to false }

    private fun sourceTitle(source: ItemSource): String =
        source.boss?.trim()?.takeIf { it.isNotEmpty() }?.let { "$it, ${source.name}" } ?: source.name

    private fun sourceLine(source: ItemSource): String = listOfNotNull(
        source.chance?.let { res.getString(R.string.gear_sheet_dropChance, GearComposer.chanceWord(it)) },
        source.kind.replace('_', ' '),
        source.standing?.lowercase(),
    ).joinToString(SEPARATOR)

    private fun guild(needs: ItemNeeds): List<GearSheetLine> =
        needs.needs.map { need ->
            val spec = "${need.spec} ${need.className}"
            val words = if (need.pick) R.string.gear_sheet_needsPick else R.string.gear_sheet_needsRunnerUp
            GearSheetLine(need.name, res.getString(words, spec), FlowLead.Crest(need.className), classColor(need.className))
        } + needs.has.map { holder ->
            GearSheetLine(holder.name, res.getString(R.string.gear_sheet_has, "${holder.spec} ${holder.className}"), FlowLead.Crest(holder.className), classColor(holder.className))
        }

    private fun guildNote(needs: ItemNeeds): String? = when {
        needs.needs.isEmpty() && needs.has.isEmpty() -> res.getString(R.string.gear_sheet_nobody)
        needs.unknownGear > 0 -> res.getQuantityString(R.plurals.gear_sheet_unknown, needs.unknownGear, needs.unknownGear)
        else -> null
    }

    private companion object {
        const val SEPARATOR = " · "
    }
}

/** The sheet over Gear: partly open, dragging to full. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun GearItemSheet(content: GearSheetContent, build: String, onDismiss: () -> Unit) {
    ModalBottomSheet(
        onDismissRequest = onDismiss,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = false),
        containerColor = FsColor.raised,
    ) {
        Column(Modifier.verticalScroll(rememberScrollState())) { GearSheetBody(content, build) }
    }
}

/** The sheet's body, pure so the parity test draws it under its grab bar. */
@Composable
fun GearSheetBody(content: GearSheetContent, build: String) {
    Column(Modifier.fillMaxWidth().background(FsColor.raised).padding(start = FsSpace.space4, end = FsSpace.space4, bottom = FsSpace.space6)) {
        val overWorn = stringResource(R.string.gear_sheet_overWorn)
        FlowItemRow(content.look, build, content.slot, content.gain, content.gain?.let { content.unitWord }, content.gain?.let { overWorn }) {}
        if (content.tooltip.isNotEmpty()) {
            Column(Modifier.padding(top = FsSpace.space2)) {
                content.tooltip.forEach { (line, effect) ->
                    Text(line, style = FsType.small, color = if (effect) FsColor.kill else FsColor.text)
                }
            }
        }
        Spacer(Modifier.height(FsSpace.space2))
        if (content.where.isNotEmpty()) Lines(stringResource(R.string.gear_sheet_where), content.where)
        content.runnerUp?.let { Lines(stringResource(R.string.gear_sheet_runnerUp), listOf(it)) }
        content.guild?.let { guild ->
            FlowSection(stringResource(R.string.gear_sheet_needs), top = 0.dp) {
                guild.forEachIndexed { i, line -> SheetRow(line, last = i == guild.lastIndex && content.guildNote == null) }
                content.guildNote?.let { FlowNote(it, Modifier.padding(vertical = FsSpace.space2)) }
            }
        }
    }
}

@Composable
private fun Lines(title: String, lines: List<GearSheetLine>) {
    FlowSection(title, top = 0.dp) {
        lines.forEachIndexed { i, line -> SheetRow(line, last = i == lines.lastIndex) }
    }
}

@Composable
private fun SheetRow(line: GearSheetLine, last: Boolean) {
    FlowRow(line.title, line.line, line.lead, titleColor = line.titleColor ?: FsColor.strong, last = last)
}
