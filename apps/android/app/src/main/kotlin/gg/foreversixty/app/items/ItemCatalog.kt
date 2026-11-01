package gg.foreversixty.app.items

import gg.foreversixty.app.model.Item
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
        val items = items(request.build, request.classSlug) ?: return null
        val sets = store.get(StaticEndpoints.getSets(request.build)).dataOrNull().orEmpty()
        val bis = request.spec?.let { store.get(StaticEndpoints.getBis(request.build, it)).dataOrNull() }
        return ItemTooltipSources(items, sets, bis)
    }

    /** A class's items by id (lane F11: a character page's gear), or null when the file did not load. */
    suspend fun items(build: String, classSlug: String): Map<Int, Item>? =
        store.get(StaticEndpoints.getItems(build, classSlug)).dataOrNull()?.items?.associateBy { it.id }
}
