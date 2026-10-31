package gg.foreversixty.app.search.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisLabels
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.ItemUseSource
import gg.foreversixty.app.model.ItemUseSpec
import gg.foreversixty.app.model.SearchItem
import kotlin.math.roundToInt

/** `Your Back pick` or `Runner-up 2 for Back`, with its meta (spec F12.4). */
data class ItemPickCard(val title: String, val meta: String, val className: String?)

/** One `Where it comes from` row: `Blackrock Depths · Dungeon · Drop chance not known`. */
data class ItemSourceLine(val kind: String, val title: String, val meta: String)

/** One `Best in slot for` row; [band], [preset] and [faction] open that list in the BiS tab. */
data class ItemSpecLine(
    val spec: String,
    val classSlug: String,
    val icon: String?,
    val title: String,
    val meta: String,
    val tag: String,
    val isPick: Boolean,
    val band: Int,
    val preset: String,
    val faction: String?,
)

/** The item page (spec F12.4) once its files are in: the head, your pick, the F2.6 tooltip body. */
data class ItemPageModel(
    val look: ItemLook,
    val headMeta: String,
    val pick: ItemPickCard?,
    val tooltip: ItemTooltipModel,
)

/** The page's words and rules, pure over the files so the tests read them. iOS: ItemPageModel.swift. */
object ItemPageCopy {
    /** Six spec rows show, then `{n} more specs`. */
    const val VISIBLE_SPECS = 6

    private val bossKinds = setOf("raid", "dungeon")
    private val worldKinds = setOf("world", "world_drop")

    fun look(item: SearchItem): ItemLook = ItemLook(item.id, item.name, item.quality, item.icon, item.itemLevel)

    /** `Rare · Back · item level 52`. */
    fun headMeta(item: SearchItem, res: Resources): String = res.getString(
        R.string.itemPage_meta,
        ItemQuality(item.quality).name?.let(res::getString).orEmpty(),
        ItemVocabulary.slotLabel(res, item.slot),
        item.itemLevel,
    )

    /**
     * The card only when the item is on the character's own list for [band] (the tooltip's band
     * selection): `Your {Slot} pick`, or `Runner-up {n} for {Slot}`. Nothing otherwise.
     */
    fun pickCard(itemId: Int, band: BisBand?, own: BisSelection?, catalog: SpecCatalog, res: Resources): ItemPickCard? {
        if (band == null || own == null) return null
        val placement = band.slots.firstNotNullOfOrNull { slot ->
            when {
                slot.itemId == itemId -> slot.slot to 0
                else -> slot.alternatives.orEmpty().indexOfFirst { it.itemId == itemId }.takeIf { it >= 0 }?.let { slot.slot to it + 1 }
            }
        } ?: return null
        val (slotKey, rank) = placement
        val slot = ItemVocabulary.slotLabel(res, slotKey)
        val title = if (rank == 0) res.getString(R.string.itemPage_pick, slot) else res.getString(R.string.itemPage_runnerUp, rank, slot)
        val spec = catalog.spec(own.spec)?.name ?: own.spec
        val bandWords = BisLabels.band(res, BisRange(band.band, band.preset))
        val faction = BisLabels.faction(res, band.faction)
        val meta = if (band.preset == BisRange.RAID) {
            res.getString(R.string.itemPage_pickMeta, spec, bandWords, res.getString(R.string.itemPage_preset_raid), faction)
        } else {
            res.getString(R.string.itemPage_pickMetaBare, spec, bandWords, faction)
        }
        return ItemPickCard(title, meta, own.classSlug)
    }

    /** `20%`, `<1%`: the site's one way to print a chance (a 0.3% drop never reads `0%`). */
    fun percent(chance: Double): String = if (chance < 1) "<1%" else "${chance.roundToInt()}%"

    /**
     * One row per source, in the file's order: the place and the kind word, then for a boss drop
     * `{n}% from {boss}` or `Drop chance not known`, and a quest's `Requires level {n}`.
     */
    fun sources(sources: List<ItemUseSource>, res: Resources): List<ItemSourceLine> = sources.map { source ->
        val kind = kindWord(source.kind, res)?.takeIf { word -> !source.title.startsWith(word) }
        val known = source.chance
        val chance = when {
            known != null -> res.getString(R.string.itemPage_dropChance, percent(known), source.title)
            source.kind in bossKinds -> res.getString(R.string.itemPage_chanceUnknown)
            else -> null
        }
        val quest = source.minLevel?.takeIf { it > 0 }?.let { res.getString(R.string.itemPage_questLevel, it) }
        ItemSourceLine(source.kind, source.title, listOfNotNull(source.place, kind, chance, quest).joinToString(SEPARATOR))
    }

    /** The kind word; null for a kind the app does not know yet. */
    fun kindWord(kind: String, res: Resources): String? = when (kind) {
        "raid" -> R.string.itemPage_kind_raid
        "dungeon" -> R.string.itemPage_kind_dungeon
        in worldKinds -> R.string.itemPage_kind_world
        "zone" -> R.string.itemPage_kind_zone
        "vendor" -> R.string.itemPage_kind_vendor
        "crafted" -> R.string.itemPage_kind_crafted
        "rep" -> R.string.itemPage_kind_rep
        "pvp" -> R.string.itemPage_kind_pvp
        "quest" -> R.string.itemPage_kind_quest
        else -> null
    }?.let(res::getString)

    /** The specs that list it, in the file's order (picks first): `Balance Druid`, `60 · Phase 1 raid`, `Pick`. */
    fun specs(specs: List<ItemUseSpec>, catalog: SpecCatalog, copy: SearchCopy, res: Resources): List<ItemSpecLine> = specs.map { row ->
        val isPick = row.rank == 0
        ItemSpecLine(
            spec = row.spec,
            classSlug = row.className,
            icon = catalog.spec(row.spec)?.icon,
            title = res.getString(R.string.bis_specChip, row.specName, copy.className(row.className)),
            meta = res.getString(if (row.preset == BisRange.RAID) R.string.itemPage_specMeta_raid else R.string.itemPage_specMeta_bare, row.band),
            tag = if (isPick) res.getString(R.string.itemPage_tag_pick) else res.getString(R.string.itemPage_tag_runnerUp, row.rank),
            isPick = isPick,
            band = row.band,
            preset = row.preset,
            faction = row.faction,
        )
    }

    /** Six rows, or all of them once `{n} more specs` is open; the count left behind it. */
    fun visibleSpecs(all: List<ItemSpecLine>, expanded: Boolean): Pair<List<ItemSpecLine>, Int> =
        if (expanded || all.size <= VISIBLE_SPECS) all to 0 else all.take(VISIBLE_SPECS) to all.size - VISIBLE_SPECS

    private const val SEPARATOR = " · "
}
