package gg.foreversixty.app.items

import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.network.dataOrNull

/**
 * The build's static files the tooltip reads (`items/{class}.json`, `sets.json`,
 * `bis/{spec}.json`), through the per-build [StaticDataStore]: fetched once, kept in memory for
 * the session and on disk until the store swaps to a newer build (spec §5.0 Caching).
 */
class ItemCatalog(private val store: StaticDataStore) {
    /**
     * The files a tooltip needs. Items are required; sets and the BiS file only add lines, so a
     * missing one leaves its lines out rather than failing the sheet. Null when the items file
     * did not load.
     */
    suspend fun sources(request: ItemTooltipRequest): ItemTooltipSources? {
        val items = store.get(StaticEndpoints.getItems(request.build, request.classSlug)).dataOrNull() ?: return null
        val sets = store.get(StaticEndpoints.getSets(request.build)).dataOrNull().orEmpty()
        val bis = request.spec?.let { store.get(StaticEndpoints.getBis(request.build, it)).dataOrNull() }
        return ItemTooltipSources(items.items.associateBy { it.id }, sets, bis)
    }
}
