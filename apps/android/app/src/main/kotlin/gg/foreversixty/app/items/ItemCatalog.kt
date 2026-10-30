package gg.foreversixty.app.items

import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedStaticDataClient
import gg.foreversixty.app.network.FileStaticFileCache
import gg.foreversixty.app.network.InMemoryStaticFileCache
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticEndpoints
import java.io.File

// MERGE NOTE (lanes F2 and F3): F2's ItemCatalog.kt replaces this file at integration. Same name,
// same constructor and `sources(request)`, so BiS's call site compiles against either.

/**
 * The build's static files a tooltip reads (`items/{class}.json`, `bis/{spec}.json`) through the
 * offline cache, so a tooltip opens offline from what the list already saved. [directory] null
 * keeps them in memory (tests and previews).
 */
class ItemCatalog(client: StaticDataClient, directory: File? = null) {
    private val files = CachedStaticDataClient(
        client,
        if (directory == null) InMemoryStaticFileCache() else FileStaticFileCache(directory),
    )

    /** Items are required (null when they cannot load); the BiS file only adds lines. */
    suspend fun sources(request: ItemTooltipRequest): ItemTooltipSources? {
        val items = (files.load(StaticEndpoints.getItems(request.build, request.classSlug)) as? ApiResult.Success)
            ?.data?.value ?: return null
        val bis: BisFile? = request.spec?.let { spec ->
            (files.load(StaticEndpoints.getBis(request.build, spec)) as? ApiResult.Success)?.data?.value
        }
        return ItemTooltipSources(items.items.associateBy { it.id }, emptyList(), bis)
    }
}
