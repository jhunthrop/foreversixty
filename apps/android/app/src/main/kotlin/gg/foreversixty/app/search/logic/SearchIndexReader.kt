package gg.foreversixty.app.search.logic

import gg.foreversixty.app.model.SearchGuide
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.network.dataOrNull
import kotlin.coroutines.CoroutineContext
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import kotlinx.coroutines.withContext

/**
 * The search index reader (Wave 3 spec §14, lane F12): `search.json` for the active build through
 * the [StaticDataStore] (saved on the phone, so it searches offline once fetched), prepared once
 * per build off the main thread and kept for the session. A newer build swaps the store, so the
 * old build's files leave the phone.
 */
class SearchIndexReader(
    private val store: StaticDataStore,
    private val activeBuild: suspend () -> String,
    private val guideTitle: (SearchGuide) -> String,
    private val compute: CoroutineContext = Dispatchers.Default,
) : SearchIndexSource {
    private val lock = Mutex()
    private var prepared: PreparedIndex? = null

    override suspend fun load(): PreparedIndex? = lock.withLock {
        val build = activeBuild()
        prepared?.takeIf { it.build == build }?.let { return@withLock it }
        val file = store.get(StaticEndpoints.getSearchIndex(build)).dataOrNull() ?: return@withLock null
        store.swap(build)
        withContext(compute) {
            PreparedIndex(build, file.itemUsesShards, file.items, file.talents, file.guides, guideTitle)
        }.also { prepared = it }
    }
}
