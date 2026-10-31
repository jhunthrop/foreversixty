package gg.foreversixty.app.sims

import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.sims.logic.SimGearCell
import gg.foreversixty.app.sims.logic.SimGearLook
import gg.foreversixty.app.sims.logic.SimResultModel

/** The gear grid's names and icons, from the build's items file through the shared catalog. */
object SimGearLooks {
    suspend fun load(model: SimResultModel, build: String, catalog: ItemCatalog): Map<Int, SimGearLook> {
        val probe = model.gear.firstOrNull() ?: return emptyMap()
        val items = catalog.sources(request(probe, null, model, build))?.items ?: return emptyMap()
        return model.gear.mapNotNull { cell -> items[cell.itemId]?.let { cell.itemId to SimGearLook(it.name, it.icon, it.quality) } }.toMap()
    }

    fun request(cell: SimGearCell, look: SimGearLook?, model: SimResultModel, build: String) = ItemTooltipRequest(
        itemId = cell.itemId,
        fallback = ItemLook(cell.itemId, look?.name, look?.quality, look?.icon),
        build = build,
        classSlug = model.classSlug,
        spec = null,
        specName = model.specName,
        band = null,
        bandLabel = null,
        preset = null,
        faction = null,
    )
}
