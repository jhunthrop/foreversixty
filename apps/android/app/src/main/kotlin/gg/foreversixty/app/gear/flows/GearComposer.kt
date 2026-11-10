package gg.foreversixty.app.gear.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisList
import gg.foreversixty.app.bis.BisListBuilder
import gg.foreversixty.app.bis.BisRow
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.phrase
import java.time.Instant
import java.util.Locale
import kotlin.math.roundToInt

/** One Where to go tonight row: the place, what it gives, the best item's look as its lead, `2 upgrades · +20.4 DPS`. */
data class GearWhereRow(val title: String, val line: String, val lead: ItemLook, val figure: String? = null)

/** One slot of the paperdoll's list: the list's pick (null when the list has none), its gain or Equipped. */
data class GearSlotRow(
    val slot: String,
    /** The slot's word, `Head`, `Main hand`. */
    val label: String,
    val look: ItemLook?,
    /** `Head · Blacksmithing`. */
    val source: String,
    /** `+14.5`, or null when equipped or not measured. */
    val gain: String?,
    val worn: Boolean,
    val row: BisRow,
    /** What the character wears in the slot today, which the pick replaces. */
    val equipped: ItemLook? = null,
)

/**
 * Gear's answer (spec 2026-11-08 4.3, round 5 mock `m-android-gear-frame`): the paperdoll strip (the
 * spec's list and its set figure), the line naming the list, every slot, then where to go tonight.
 */
data class GearContent(
    /** `Fury Warrior`. */
    val title: String,
    /** The set's figure, `1004.2`, in [unitWord]; null for a tank list. */
    val figure: String?,
    /** `60 raid · Alliance · 4 of 17 slots on the list`. */
    val facts: String,
    val where: List<GearWhereRow>,
    val slots: List<GearSlotRow>,
    val unitWord: String,
    /** `+157.6 DPS from 13 upgrades · synced 2 h ago · 2 gains not modelled yet`, under the band line. */
    val header: String? = null,
)

/**
 * Builds Gear from the BiS list on screen, the character's measured upgrades (Home's answer, for the
 * same list) and the picks [gg.foreversixty.app.gear.WhereToGo] placed at their sources.
 */
class GearComposer(private val res: Resources, private val now: Instant = Instant.now()) {
    fun content(list: BisList, faction: String, upgrades: List<HomeUpgrade>, picks: List<RunPick>, home: CharacterHome? = null): GearContent {
        val bySlot = upgrades.associateBy { it.slot }
        val slots = list.rows.map { row -> slotRow(row, bySlot[row.slot]) }
        val onList = res.getString(R.string.gear_onList, slots.count { it.worn }, list.rows.size)
        return GearContent(
            title = list.specLabel,
            figure = list.setFigure,
            facts = listOf(list.rangeLabel, faction.replaceFirstChar { it.titlecase(Locale.US) }, onList)
                .filter { it.isNotBlank() }
                .joinToString(SEPARATOR),
            where = where(picks, list),
            slots = slots,
            unitWord = res.getString(list.unit.word),
            header = home?.let { header(it, res.getString(list.unit.word)) },
        )
    }

    /** The header line (round 2): the measured total and count, when the gear synced, the gains not modelled. */
    fun header(home: CharacterHome, unit: String): String? {
        val upgrades = res.getQuantityString(R.plurals.shared_home_upgradesCount, home.upgradeCount, home.upgradeCount)
        val gain = home.totalGain?.takeIf { it > 0 }?.let { res.getString(R.string.gear_header_gain, "+" + BisListBuilder.figure(it), unit, upgrades) }
        val synced = home.syncedAt?.let { runCatching { Instant.parse(it) }.getOrNull() }
            ?.let { res.getString(R.string.gear_header_synced, RelativeAge.of(it, now).phrase(res)) }
        val line = listOfNotNull(gain, synced).joinToString(SEPARATOR)
        val notModelled = home.notModelledCount.takeIf { it > 0 }?.let { res.getQuantityString(R.plurals.home_answer_notModelled, it, it) }.orEmpty()
        return (line + notModelled).removePrefix(SEPARATOR).ifBlank { null }
    }

    private fun slotRow(row: BisRow, upgrade: HomeUpgrade?): GearSlotRow {
        val pick = row.pick
        val item = pick?.item
        val gain = upgrade?.takeIf { it.modelled && it.pick.id == item?.id }?.gain?.takeIf { it > 0 }
        val equipped = upgrade?.worn?.takeIf { it.known }?.let { ItemLook(it.id, it.name, it.quality, it.icon, it.itemLevel) }
        return GearSlotRow(
            slot = row.slot,
            label = row.slotLabel,
            look = item?.let { ItemLook(it.id, it.name, it.quality, it.icon, it.item?.itemLevel) },
            source = pick?.slotSource ?: row.slotLabel,
            gain = if (pick?.isWorn != false) null else gain?.let { "+" + BisListBuilder.figure(it) },
            worn = pick?.isWorn == true,
            row = row,
            equipped = equipped.takeIf { pick?.isWorn != true },
        )
    }

    /**
     * The places worth a trip tonight: instances (most pieces first), then professions, then reputations;
     * each slot counted once per place, its best candidate there.
     */
    fun where(picks: List<RunPick>, list: BisList): List<GearWhereRow> {
        val looks = list.rows.flatMap { row ->
            val pick = row.pick ?: return@flatMap emptyList()
            listOf(pick.item) + pick.runnersUp.map { it.item }
        }.associateBy { it.id }
        return picks
            .filter { kindOrder(it.source.kind) in PLACE_KINDS }
            .groupBy { it.source.kind to it.source.name }
            .map { (_, group) -> group.groupBy { it.slot }.map { (_, slot) -> slot.maxBy { it.gain } }.sortedByDescending { it.gain } }
            .sortedWith(compareBy<List<RunPick>> { PLACE_KINDS.indexOf(kindOrder(it.first().source.kind)) }.thenByDescending { it.size })
            .map { group ->
                val best = group.first()
                val ref = looks[best.itemId]
                GearWhereRow(
                    title = best.source.name,
                    line = whereLine(group),
                    lead = ItemLook(best.itemId, best.itemName, ref?.quality, ref?.icon),
                    figure = whereFigure(group, res.getString(list.unit.word)),
                )
            }
    }

    /** `2 upgrades · +20.4 DPS`: the slots a trip fills and their summed gain. */
    private fun whereFigure(group: List<RunPick>, unit: String): String {
        val count = res.getQuantityString(R.plurals.shared_home_upgradesCount, group.size, group.size)
        return res.getString(R.string.gear_where_figure, count, "+" + BisListBuilder.figure(group.sumOf { it.gain }), unit)
    }

    private fun whereLine(group: List<RunPick>): String {
        val items = group.joinToString(", ") { it.itemName }
        val source = group.first().source
        return when (source.kind) {
            CRAFTED -> crafter(source.profession ?: source.name)?.let { res.getString(R.string.gear_where_ask, items, it) } ?: items
            REP -> source.standing?.let { res.getString(R.string.gear_where_rep, items, it.lowercase(Locale.US)) } ?: items
            else -> {
                val boss = source.boss?.trim()?.takeIf { group.size == 1 } ?: return items
                val chance = source.chance?.let { chanceWord(it) }
                if (chance == null) res.getString(R.string.gear_where_boss, items, boss) else res.getString(R.string.gear_where_drop, items, boss, chance)
            }
        }
    }

    private fun crafter(profession: String): String? = when (profession.lowercase(Locale.US)) {
        "blacksmithing" -> R.string.gear_crafter_blacksmithing
        "leatherworking" -> R.string.gear_crafter_leatherworking
        "tailoring" -> R.string.gear_crafter_tailoring
        "engineering" -> R.string.gear_crafter_engineering
        "alchemy" -> R.string.gear_crafter_alchemy
        "enchanting" -> R.string.gear_crafter_enchanting
        else -> null
    }?.let(res::getString)

    companion object {
        private const val SEPARATOR = " · "
        private const val DUNGEON = "dungeon"
        private const val RAID = "raid"
        private const val CRAFTED = "crafted"
        private const val REP = "rep"

        /** The kinds Where to go tonight lists (spec 4.3: instance, profession, reputation), in its order. */
        private val PLACE_KINDS = listOf(DUNGEON, CRAFTED, REP)

        private fun kindOrder(kind: String) = if (kind == RAID) DUNGEON else kind

        /** `8%`, or `0.5%` under one. */
        fun chanceWord(chance: Double): String =
            if (chance >= 1) "${chance.roundToInt()}%" else "${BisListBuilder.figure(chance)}%"
    }
}
