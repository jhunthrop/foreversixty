package gg.foreversixty.app.gear.flows

import android.content.res.Resources
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
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
import androidx.compose.ui.text.font.FontWeight
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisListBuilder
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.gear.ItemSource
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipPanel
import gg.foreversixty.app.model.ItemNeeds

/** One line of the sheet's list: a lead (none for a source row), the title in its colour, its line. */
data class GearSheetLine(val title: String, val line: String?, val lead: FlowLead?, val titleColor: Color? = null)

/**
 * The item sheet (spec F5, round 5 mock `m-android-item`): the client's tooltip with the gain closing
 * it, then where it comes from, what it replaces, the runner-up and who in the guild needs or has it.
 */
data class GearSheetContent(
    /** The tooltip; while the static file loads, the name alone. */
    val tooltip: ItemTooltipModel,
    /** `+12.6 DPS` over what is equipped, or null when it is equipped or not measured. */
    val gain: String?,
    /** `Skum, Wailing Caverns (15-24), 50%`, one line per source. */
    val sources: List<GearSheetLine>,
    /** `Replaces Ironwood Blade`, Equipped: what the character wears in the slot today. */
    val replaces: GearSheetLine?,
    val runnerUp: GearSheetLine?,
    /** Null while the guild's answer loads, or for a player with no guild (left out). */
    val guild: List<GearSheetLine>?,
    val guildNote: String?,
) {
    val lines: List<GearSheetLine> get() = sources + listOfNotNull(replaces, runnerUp) + guild.orEmpty()
}

/** Words the sheet from what Gear already holds and the guild's needs answer. */
class GearSheetComposer(private val res: Resources, private val build: String) {
    fun content(
        row: GearSlotRow,
        unitWord: String,
        tooltip: ItemTooltipModel?,
        sources: List<ItemSource>,
        needs: ItemNeeds?,
        /** The runner-up's first place in the loot file, when it has one. */
        runnerUpSource: ItemSource? = null,
    ): GearSheetContent {
        val pick = requireNotNull(row.row.pick) { "a sheet opens on a pick" }
        val look = requireNotNull(row.look) { "a sheet opens on a pick" }
        return GearSheetContent(
            // The client's tooltip has no item level (spec 2026-11-08, round 4 cut), so the sheet drops it.
            tooltip = (tooltip ?: ItemTooltipModel(look = look, headerMeta = "", known = true)).copy(itemLevel = null),
            gain = row.gain?.let { if (row.gains) "$it $unitWord" else res.getString(R.string.gear_cell_noGain) },
            // Round 8 row 1: a source row is the source alone, self-describing, no caption and no icon
            // (it must not read as an item row); the loot file's places, else the list's own line.
            sources = sourceRows(sources, pick.source),
            replaces = row.equipped?.let { worn ->
                GearSheetLine(res.getString(R.string.gear_sheet_replaces, worn.name.orEmpty()), res.getString(R.string.gear_worn), FlowLead.Item(worn, build))
            },
            runnerUp = pick.runnersUp.firstOrNull()?.let { alt ->
                // Round 8 row 4 (as iOS): the runner-up always says what it gives against the pick.
                val delta = BisListBuilder.signedDelta(alt.dpsDelta)?.let { "$it $unitWord" } ?: res.getString(R.string.items_sameDps, unitWord)
                val altLook = ItemLook(alt.item.id, alt.item.name, alt.item.quality, alt.item.icon)
                val place = runnerUpPlace(runnerUpSource, alt.source).takeIf { !samePlace(it, runnerUpSource, sources, pick.source) }
                GearSheetLine(
                    alt.item.name,
                    listOfNotNull(res.getString(R.string.gear_sheet_runnerUp), place, delta).joinToString(SEPARATOR),
                    FlowLead.Item(altLook, build),
                    ItemQuality(alt.item.quality).text,
                )
            },
            guild = needs?.let(::guild),
            guildNote = needs?.let { guildNote(it) },
        )
    }

    /** The pick's source rows: the loot file's places, else the list's own line (round 8 row 1). */
    private fun sourceRows(sources: List<ItemSource>, listSource: String): List<GearSheetLine> =
        sources.map { GearSheetLine(sourceLine(it), null, null) }
            .ifEmpty { listOfNotNull(listWords(listSource)?.let { GearSheetLine(it, null, null) }) }

    /**
     * Where the runner-up comes from in the source row's shared words (round 9 row 2): a boss drop as
     * `Ancient Treasure · 1%`, any other place as its source row reads, else the list's own words.
     */
    private fun runnerUpPlace(source: ItemSource?, listSource: String): String? = when {
        source == null -> listWords(listSource)
        source.boss?.isNotBlank() == true -> listOfNotNull(source.boss.trim(), source.chance?.let(GearComposer::chanceWord)).joinToString(SEPARATOR)
        else -> sourceLine(source)
    }

    /** The runner-up's place is the pick's own: its source row already says it, so the row leaves it out. */
    private fun samePlace(place: String?, source: ItemSource?, sources: List<ItemSource>, listSource: String): Boolean = when {
        place == null -> false
        source != null -> sources.any { it.kind == source.kind && it.name == source.name && it.boss == source.boss }
        else -> place == listWords(listSource)
    }

    /**
     * The list's own source words on the sheet (round 9 row 3): a PvP source drops its faction word, the
     * rank title already belongs to one side (`PvP rank 11 · Blood Guard`).
     */
    private fun listWords(source: String): String? {
        if (source.isBlank()) return null
        val parts = source.split(SEPARATOR)
        return if (source.startsWith(PVP_PREFIX) && parts.last() in FACTIONS) parts.dropLast(1).joinToString(SEPARATOR) else source
    }

    /**
     * The source in the shared source words with its figure when the data has one (round 8 row 1):
     * `Skum, Wailing Caverns (15-24), 50%`, `Crafted: Blacksmithing (245)`, `Vendor: Pix Xizzix`,
     * `Quest: A Collection of Heads`, `Stormpike Guard, revered`, `World drop`.
     */
    fun sourceLine(source: ItemSource): String = when (source.kind) {
        VENDOR -> res.getString(R.string.bis_source_vendor, source.name)
        QUEST -> res.getString(R.string.bis_source_quest, source.name)
        CRAFTED -> source.skill?.let { res.getString(R.string.bis_source_craftedSkill, source.name, it) }
            ?: res.getString(R.string.bis_source_crafted, source.name)
        else -> listOfNotNull(
            source.boss?.trim()?.takeIf { it.isNotEmpty() },
            source.name + source.levels.orEmpty(),
            source.chance?.let { GearComposer.chanceWord(it) },
            source.standing?.lowercase(),
        ).joinToString(", ")
    }

    private fun IntRange?.orEmpty(): String = when {
        this == null -> ""
        first == last -> ""
        else -> " ($first-$last)"
    }

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
        const val VENDOR = "vendor"
        const val QUEST = "quest"
        const val CRAFTED = "crafted"
        const val PVP_PREFIX = "PvP rank"

        /** The faction words a PvP source ends with in the list (the server's). */
        val FACTIONS = setOf("Horde", "Alliance")
    }
}

/** The sheet over Gear: partly open, dragging to full. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun GearItemSheet(content: GearSheetContent, onDismiss: () -> Unit) {
    ModalBottomSheet(
        onDismissRequest = onDismiss,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = false),
        containerColor = FsColor.raised,
    ) {
        Column(Modifier.verticalScroll(rememberScrollState())) { GearSheetBody(content) }
    }
}

/** The sheet's body, pure so the parity test draws it under its grab bar. */
@Composable
fun GearSheetBody(content: GearSheetContent) {
    Column(Modifier.fillMaxWidth().background(FsColor.raised).padding(start = FsSpace.space4, end = FsSpace.space4, bottom = FsSpace.space6)) {
        ItemTooltipPanel(content.tooltip, Modifier.parityRegion(ParityRegion.INSPECTOR)) {
            content.gain?.let { Text(it, style = ScaledType.tooltip.copy(fontWeight = FontWeight.SemiBold), color = FsColor.gold) }
        }
        val lines = content.lines
        if (lines.isNotEmpty() || content.guildNote != null) {
            FlowBlock(SHEET_LINES, top = FsSpace.space3) {
                lines.forEachIndexed { i, line ->
                    FlowRow(
                        line.title, line.line, line.lead,
                        titleColor = line.titleColor ?: FsColor.strong,
                        last = i == lines.lastIndex && content.guildNote == null,
                    )
                }
                content.guildNote?.let { FlowNote(it, Modifier.padding(vertical = FsSpace.space2)) }
            }
        }
    }
}

/** The region of the sheet's list under the tooltip. */
private const val SHEET_LINES = "section-where-to-get-it"
