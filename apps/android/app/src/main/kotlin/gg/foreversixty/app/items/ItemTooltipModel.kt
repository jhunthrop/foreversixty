package gg.foreversixty.app.items

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisListBuilder
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisUnit
import gg.foreversixty.app.bis.band
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.BisSlot
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.model.ItemSet

// MERGE NOTE (lanes F2 and F3): F2's ItemTooltipModel.kt replaces this file at integration. The
// request, the look and the sources carry F2's names and fields exactly, so BiS's call site
// (`ItemTooltipSheet(request, catalog, onDismiss)`) compiles against either. F2's model must word
// the next best in the band's unit (`tooltip_delta` and `items_sameDps` take the unit) and drop
// the preset from the footer (`tooltip_pickOf` has no preset).

/** An item's name, quality, icon and level; every field but id may be null. */
data class ItemLook(
    val id: Int,
    val name: String? = null,
    val quality: Int? = null,
    val icon: String? = null,
    val itemLevel: Int? = null,
)

/**
 * Who asked for a tooltip and what they already know: the item, the build its static data lives
 * in, and the band the pick belongs to. Built by Home, the gear list and BiS.
 */
data class ItemTooltipRequest(
    val itemId: Int,
    /** The name and look to draw while the static file loads, or when it lacks the item. */
    val fallback: ItemLook,
    val build: String,
    val classSlug: String,
    val spec: String?,
    val specName: String?,
    val band: Int?,
    val bandLabel: String?,
    val preset: String?,
    val faction: String?,
    /** The caller's source words for this item, when it has them. */
    val sourceLine: String? = null,
    /** The item worn in the same slot, so the next best can say "You wear this". */
    val wornId: Int? = null,
)

/** The static files a tooltip reads, already loaded. */
data class ItemTooltipSources(
    val items: Map<Int, Item>,
    val sets: List<ItemSet>,
    val bis: BisFile?,
)

/** One runner-up under "The next best". */
data class ItemTooltipAlternative(val look: ItemLook, val meta: String)

/** What F3's minimal sheet draws: header, level lines, source, the next best in the band's unit, the pick footer. */
data class ItemTooltipModel(
    val look: ItemLook,
    val item: Item?,
    val slotKey: String?,
    val source: String?,
    val alternatives: List<ItemTooltipAlternative>,
    val pickFooter: String?,
) {
    companion object {
        /** Spec F2.6: up to two alternatives. */
        private const val MAX_ALTERNATIVES = 2

        fun build(request: ItemTooltipRequest, sources: ItemTooltipSources, res: Resources): ItemTooltipModel {
            val item = sources.items[request.itemId]
            val band = band(request, sources.bis)
            val slot = band?.slots?.firstOrNull { it.itemId == request.itemId }
            val fallback = request.fallback
            return ItemTooltipModel(
                look = ItemLook(
                    id = request.itemId,
                    name = item?.name ?: fallback.name,
                    quality = item?.quality ?: fallback.quality,
                    icon = item?.icon ?: fallback.icon,
                    itemLevel = item?.itemLevel ?: fallback.itemLevel,
                ),
                item = item,
                slotKey = slot?.slot ?: item?.slot,
                source = request.sourceLine,
                alternatives = alternatives(slot, request, sources.items, BisUnit.forRole(band?.role), res),
                pickFooter = footer(slot, request, res),
            )
        }

        /** The band the request names, by the site's selection rule. */
        fun band(request: ItemTooltipRequest, bis: BisFile?): BisBand? {
            val file = bis ?: return null
            val band = request.band ?: return null
            val preset = request.preset ?: return null
            val faction = request.faction ?: return null
            return file.band(BisRange(band, preset), faction)
        }

        private fun alternatives(
            slot: BisSlot?,
            request: ItemTooltipRequest,
            items: Map<Int, Item>,
            unit: BisUnit,
            res: Resources,
        ): List<ItemTooltipAlternative> {
            val unitWord = res.getString(unit.word)
            return slot?.alternatives.orEmpty().take(MAX_ALTERNATIVES).map { alt ->
                val row = items[alt.itemId]
                val look = ItemLook(alt.itemId, row?.name ?: alt.itemName, row?.quality, row?.icon, row?.itemLevel)
                val delta = BisListBuilder.signedDelta(alt.dpsDelta)?.let { res.getString(R.string.tooltip_delta, it, unitWord) }
                    ?: res.getString(R.string.items_sameDps, unitWord)
                val worn = if (alt.itemId == request.wornId) res.getString(R.string.tooltip_youWear) else ""
                ItemTooltipAlternative(look, worn + delta)
            }
        }

        /** `Best in slot for Head · Fury Warrior 60 · Phase 1 raid`, only for the pick; never a preset id. */
        private fun footer(slot: BisSlot?, request: ItemTooltipRequest, res: Resources): String? {
            slot ?: return null
            val specName = request.specName ?: return null
            val bandLabel = request.bandLabel ?: return null
            return res.getString(R.string.tooltip_pickOf, ItemWords.slot(res, slot.slot), specName, bandLabel)
        }
    }
}
