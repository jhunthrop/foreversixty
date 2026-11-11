package gg.foreversixty.app.gear.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisList
import gg.foreversixty.app.bis.BisPhase
import gg.foreversixty.app.bis.BisRow
import gg.foreversixty.app.bis.BisUnit
import gg.foreversixty.app.common.HouseFigure
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.model.HomeAlreadyBest
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.isStale
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
    /** `+14.5`, `+0.0` for a measured slot that gains nothing, or null when equipped or not modelled. */
    val gain: String?,
    val worn: Boolean,
    val row: BisRow,
    /** What the character wears in the slot today, which the pick replaces. */
    val equipped: ItemLook? = null,
    /** Round 4 row 3: the pick's gain is not modelled yet, so the cell reads `Not modelled`. */
    val notModelled: Boolean = false,
    /** The gain is above zero: drawn in the gain green; a measured `+0.0` stays plain. */
    val gains: Boolean = gain != null,
)

/**
 * Gear's header line, `+157.6 DPS from 13 upgrades · synced 2 h ago · 2 gains not modelled yet`: the
 * [synced] part takes the warning tone when [syncedStale] (round 4 row 2, the drawer's rule).
 */
data class GearHeader(val gain: String?, val synced: String?, val syncedStale: Boolean, val notModelled: String?) {
    /** The line as one string: what TalkBack reads and the tests compare. */
    val text: String get() = parts().joinToString("") { it.first }

    /** The line's parts in order, each with whether it is the stale stamp. */
    fun parts(): List<Pair<String, Boolean>> {
        val head = listOfNotNull(gain?.let { it to false }, synced?.let { it to syncedStale })
        val joined = head.flatMapIndexed { i, part -> if (i == 0) listOf(part) else listOf(SEPARATOR to false, part) }
        val tail = notModelled?.let { (if (joined.isEmpty()) it.removePrefix(SEPARATOR) else it) to false }
        return joined + listOfNotNull(tail)
    }

    private companion object {
        const val SEPARATOR = " · "
    }
}

/**
 * Gear's answer (spec 2026-11-08 4.3, round 5 mock `m-android-gear-frame`): the paperdoll strip (the
 * spec's list and its set figure), the line naming the list, every slot, then where to go tonight.
 */
data class GearContent(
    /** `Fury Warrior`. */
    val title: String,
    /** The set's figure in the house format, `1,004`, in [unitWord]; null for a tank list. */
    val figure: String?,
    /** `60 raid · Alliance · 4 of 17 slots on the list`. */
    val facts: String,
    val where: List<GearWhereRow>,
    val slots: List<GearSlotRow>,
    val unitWord: String,
    /** `+157.6 DPS from 13 upgrades · synced 2 h ago · 2 gains not modelled yet`, under the band line. */
    val header: GearHeader? = null,
)

/**
 * Gear with no list to draw (round 2): the spec as the title, the line naming the list with Change (so
 * the player can pick another), and one row: no published list, or `We could not load this list` with
 * Try again ([retry]).
 */
data class GearNotice(
    /** `Beast Mastery Hunter`. */
    val title: String,
    /** `30 to 39 · Horde`. */
    val facts: String,
    val row: String,
    val sub: String? = null,
    val retry: Boolean = false,
)

/**
 * Builds Gear from the BiS list on screen, the character's measured upgrades (Home's answer, for the
 * same list) and the picks [gg.foreversixty.app.gear.WhereToGo] placed at their sources. [setDps] is the
 * band's set figure, which Gear heads with in the house format (round 7 row 1; the BiS screen keeps its own).
 */
class GearComposer(private val res: Resources, private val now: Instant = Instant.now()) {
    fun content(
        list: BisList,
        faction: String,
        upgrades: List<HomeUpgrade>,
        picks: List<RunPick>,
        home: CharacterHome? = null,
        setDps: Double? = null,
    ): GearContent {
        val bySlot = upgrades.associateBy { it.slot }
        // Row 10: with the home answer, `already_best` decides Equipped, so the counter, the Equipped
        // cells, the header's upgrade count and the gain cells read the same figures.
        val best = home?.alreadyBest?.associateBy { it.slot }
        val slots = list.rows.map { row -> slotRow(row, bySlot[row.slot], best) }
        val onList = res.getString(R.string.gear_onList, slots.count { it.worn }, list.rows.size)
        return GearContent(
            title = list.specLabel,
            // A tank band's set_dps is not its score: no figure, as the BiS list has none.
            figure = setDps?.takeIf { list.unit != BisUnit.Score }?.let(HouseFigure::of),
            facts = facts(list.rangeLabel, faction, onList),
            where = where(picks, list),
            slots = slots,
            unitWord = res.getString(list.unit.word),
            header = home?.let { header(it, res.getString(list.unit.word)) },
        )
    }

    /** Gear's [GearNotice] for a [phase] with no list to draw; null while it loads or once it has one. */
    fun notice(phase: BisPhase, title: String, rangeLabel: String, faction: String): GearNotice? {
        val facts = facts(rangeLabel, faction)
        return when (phase) {
            is BisPhase.NoList -> GearNotice(title, facts, res.getString(R.string.home_answer_noList))
            BisPhase.Failed -> GearNotice(title, facts, res.getString(R.string.gear_error_title), res.getString(R.string.home_error_sub), retry = true)
            BisPhase.NotCached -> GearNotice(title, facts, res.getString(R.string.gear_error_title), res.getString(R.string.bis_notCached), retry = true)
            BisPhase.Loading, is BisPhase.Loaded -> null
        }
    }

    /** `60 raid · Alliance · 4 of 17 slots match the list`: the band, the faction, then what follows. */
    private fun facts(rangeLabel: String, faction: String, vararg more: String): String =
        (listOf(rangeLabel, faction.replaceFirstChar { it.titlecase(Locale.US) }) + more)
            .filter { it.isNotBlank() }
            .joinToString(SEPARATOR)

    /**
     * The header line (round 2): the measured total and count, when the gear synced (stale past a week,
     * round 4 row 2), the gains not modelled.
     */
    fun header(home: CharacterHome, unit: String): GearHeader? {
        val upgrades = res.getQuantityString(R.plurals.shared_home_upgradesCount, home.upgradeCount, home.upgradeCount)
        val gain = home.totalGain?.takeIf { it > 0 }?.let { res.getString(R.string.gear_header_gain, "+" + HouseFigure.of(it), unit, upgrades) }
        val age = home.syncedAt?.let { runCatching { Instant.parse(it) }.getOrNull() }?.let { RelativeAge.of(it, now) }
        val synced = age?.let { res.getString(R.string.gear_header_synced, it.phrase(res)) }
        val notModelled = home.notModelledCount.takeIf { it > 0 }?.let { res.getQuantityString(R.plurals.home_answer_notModelled, it, it) }
        return GearHeader(gain, synced, age?.isStale() == true, notModelled).takeIf { it.text.isNotBlank() }
    }

    /**
     * A slot's cell: the list's pick, its measured gain, or Equipped. [best] is the home answer's
     * `already_best` by slot (null without a home answer, when the list's own worn mark decides);
     * the label is the API's `slot_label` (`Finger 1`, `Trinket 2`) when the answer names the slot.
     */
    private fun slotRow(row: BisRow, upgrade: HomeUpgrade?, best: Map<String, HomeAlreadyBest>?): GearSlotRow {
        val pick = row.pick
        val item = pick?.item
        val worn = best?.containsKey(row.slot) ?: (pick?.isWorn == true)
        val measured = upgrade?.takeIf { it.pick.id == item?.id }
        val notModelled = measured?.modelled == false
        // Round 4 row 3 (as iOS): with the home answer a pick that is not Equipped says what it gives,
        // its gain, `Not modelled`, else `+0.0`.
        val gain = measured?.takeIf { it.modelled }?.gain?.takeIf { it > 0 } ?: ZERO.takeIf { best != null && !notModelled }
        val equipped = upgrade?.worn?.takeIf { it.known }?.let { ItemLook(it.id, it.name, it.quality, it.icon, it.itemLevel) }
        val label = upgrade?.slotLabel ?: best?.get(row.slot)?.slotLabel ?: row.slotLabel
        return GearSlotRow(
            slot = row.slot,
            label = label,
            look = item?.let { ItemLook(it.id, it.name, it.quality, it.icon, it.item?.itemLevel) },
            source = pick?.slotSource ?: label,
            gain = if (worn || item == null) null else gain?.let { "+" + HouseFigure.of(it) },
            worn = worn,
            row = row,
            equipped = equipped.takeIf { !worn },
            notModelled = !worn && notModelled,
            gains = !worn && gain != null && gain > ZERO,
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
        return res.getString(R.string.gear_where_figure, count, "+" + HouseFigure.of(group.sumOf { it.gain }), unit)
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
        private const val ZERO = 0.0
        private const val DUNGEON = "dungeon"
        private const val RAID = "raid"
        private const val CRAFTED = "crafted"
        private const val REP = "rep"

        /** The kinds Where to go tonight lists (spec 4.3: instance, profession, reputation), in its order. */
        private val PLACE_KINDS = listOf(DUNGEON, CRAFTED, REP)

        private fun kindOrder(kind: String) = if (kind == RAID) DUNGEON else kind

        /**
         * A drop chance as the site prints it (web/src/lib/items/drop-chance.ts `dropPercentLabel`, the
         * web Gear's and tooltip's one way): a whole percent, half up (`8%`, `1%` for 1.3), and `<1%`
         * under one, so a rare drop never reads `0%`.
         */
        fun chanceWord(chance: Double): String = if (chance < 1) UNDER_ONE else "${chance.roundToInt()}%"

        private const val UNDER_ONE = "<1%"
    }
}
