package gg.foreversixty.app.bis

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.items.ItemWords
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.BisSlot
import gg.foreversixty.app.model.Item
import java.math.BigDecimal
import java.math.RoundingMode
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.format.DateTimeParseException
import java.util.Locale
import kotlin.math.abs

/**
 * The measure a band's figures are in: a healer band's are healing per second, a tank band's its
 * tank score (the site's `scoreUnitWord`: DPS, HPS, score).
 */
enum class BisUnit(val word: Int) {
    Dps(R.string.bis_unit_dps),
    Hps(R.string.bis_unit_hps),
    Score(R.string.bis_unit_score);

    companion object {
        fun forRole(role: String?): BisUnit = when (role) {
            "healer" -> Hps
            "tank" -> Score
            else -> Dps
        }
    }
}

/** An item as a row draws it: the file's id and name, and the build's item facts when known. */
data class BisItemRef(val id: Int, val name: String, val item: Item?) {
    val quality: Int get() = item?.quality ?: 1
    val icon: String? get() = item?.icon
}

/** A runner-up and its gap to the pick, as published. */
data class BisRunnerUp(val item: BisItemRef, val source: String, val dpsDelta: Double, val isWorn: Boolean)

/** A slot's pick, ready to draw. */
data class BisPick(
    val item: BisItemRef,
    /** `Head · Crafted: Leatherworking`. */
    val slotSource: String,
    /** `Next: Warlord's Mail Helm, −1.3 DPS`, or `Nothing else ranked`. */
    val next: String,
    val isUnmodelled: Boolean,
    val isWorn: Boolean,
    val runnersUp: List<BisRunnerUp>,
    val source: String,
)

/** One of the 17 rows, in the file's (the character frame's) order. */
data class BisRow(val slot: String, val slotLabel: String, val pick: BisPick?, val emptyReason: String?)

/** The list for one band, ready to draw. Every figure and order is the published file's. */
data class BisList(
    val setDpsLine: String?,
    val rows: List<BisRow>,
    val unit: BisUnit,
    val specLabel: String,
    val rangeLabel: String,
) {
    /** The rows whose pick or runner-up names contain [query] (spec F3.1 search). */
    fun filtered(query: String): List<BisRow> {
        val needle = query.trim()
        if (needle.isEmpty()) return rows
        return rows.filter { row ->
            val pick = row.pick ?: return@filter false
            (listOf(pick.item.name) + pick.runnersUp.map { it.item.name }).any { it.contains(needle, ignoreCase = true) }
        }
    }
}

/**
 * Turns a published band into rows. Nothing is re-ranked or recomputed: the order, the picks, the
 * runners-up and every delta are the file's; this only words them (spec §5.3).
 */
class BisListBuilder(
    private val res: Resources,
    items: List<Item>,
    /** Item ids the current character wears, any slot (a ring may sit in either finger). */
    private val worn: Set<Int>,
    private val zone: ZoneId = ZoneId.systemDefault(),
) {
    private val items: Map<Int, Item> = items.associateBy { it.id }

    fun list(file: BisFile, band: BisBand, specLabel: String, rangeLabel: String): BisList {
        val unit = BisUnit.forRole(band.role)
        val mainHand = band.slots.firstOrNull { it.slot == MAIN_HAND }?.itemName
        return BisList(
            setDpsLine = headline(file, band, unit),
            rows = band.slots.map { row(it, unit, mainHand) },
            unit = unit,
            specLabel = specLabel,
            rangeLabel = rangeLabel,
        )
    }

    /**
     * The line above the rows: a DPS or HPS band's `set_dps` with its Updated date; a tank band's
     * `set_dps` is its own damage, not what a tank is judged on (the site's `headlineDpsOf`), so it
     * says what a score point is instead.
     */
    fun headline(file: BisFile, band: BisBand, unit: BisUnit): String? {
        val date = updated(file.generatedAt)
        if (unit == BisUnit.Score) return res.getString(R.string.bis_tankLine, date)
        return band.setDps?.let { res.getString(R.string.bis_setDps, figure(it), res.getString(unit.word), date) }
    }

    fun row(slot: BisSlot, unit: BisUnit, mainHand: String?): BisRow {
        val label = ItemWords.slot(res, slot.slot)
        val id = slot.itemId
        val name = slot.itemName
        if (id == null || id <= 0 || name == null) {
            val reason = if (slot.emptyReason == TWO_HAND_EQUIPPED && mainHand != null) {
                res.getString(R.string.bis_offHandTwoHand, mainHand)
            } else {
                res.getString(R.string.bis_noPick, label)
            }
            return BisRow(slot.slot, label, pick = null, emptyReason = reason)
        }
        val source = sourceLine(slot.source.orEmpty(), slot.sourceKind)
        val runnersUp = slot.alternatives.orEmpty().map {
            BisRunnerUp(ref(it.itemId, it.itemName), sourceLine(it.source, it.sourceKind), it.dpsDelta, it.itemId in worn)
        }
        val pick = BisPick(
            item = ref(id, name),
            slotSource = if (source.isEmpty()) label else res.getString(R.string.bis_slotSource, label, source),
            next = nextLine(runnersUp.firstOrNull(), unit),
            isUnmodelled = slot.effectUnmodelled == true,
            isWorn = id in worn,
            runnersUp = runnersUp,
            source = source,
        )
        return BisRow(slot.slot, label, pick, emptyReason = null)
    }

    private fun ref(id: Int, name: String) = BisItemRef(id, name, items[id])

    /**
     * The source in the site's words: a quest, a recipe and a vendor say what they are; every
     * other kind already reads in the game's words (`Stratholme: Balzaphon`, a PvP rank).
     */
    fun sourceLine(source: String, kind: String?): String = when {
        source.isEmpty() -> source
        kind == "quest" -> res.getString(R.string.bis_source_quest, source)
        kind == "crafted" -> res.getString(R.string.bis_source_crafted, source)
        kind == "vendor" && !source.startsWith(PVP_RANK_PREFIX) -> res.getString(R.string.bis_source_vendor, source)
        else -> source
    }

    fun nextLine(runnerUp: BisRunnerUp?, unit: BisUnit): String {
        runnerUp ?: return res.getString(R.string.tooltip_nothingElse)
        val unitWord = res.getString(unit.word)
        val delta = signedDelta(runnerUp.dpsDelta) ?: return res.getString(R.string.bis_nextSame, runnerUp.item.name, unitWord)
        return res.getString(R.string.bis_next, runnerUp.item.name, delta, unitWord)
    }

    /** `29 Oct 2026`, the site's Updated date. */
    fun updated(generatedAt: String): String = try {
        UPDATED.withZone(zone).format(Instant.parse(generatedAt))
    } catch (e: DateTimeParseException) {
        generatedAt
    }

    companion object {
        const val TWO_HAND_EQUIPPED = "two_hand_equipped"
        const val MAIN_HAND = "main_hand"
        private const val PVP_RANK_PREFIX = "PvP rank"
        const val SAME_THRESHOLD = 0.05
        private const val MINUS = "−"
        private val UPDATED: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM yyyy", Locale.US)

        /**
         * Below 0.05 reads `same` (the site's `alternativeGapLabel`); otherwise a signed figure to
         * one decimal with a true minus sign.
         */
        fun signedDelta(delta: Double): String? {
            if (abs(delta) < SAME_THRESHOLD) return null
            return (if (delta < 0) MINUS else "+") + figure(abs(delta))
        }

        /** One decimal, halves away from zero, as the site's `toFixed(1)` reads. */
        fun figure(value: Double): String = BigDecimal.valueOf(value).setScale(1, RoundingMode.HALF_UP).toPlainString()
    }
}
