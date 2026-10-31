package gg.foreversixty.app.sims.logic

import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.SimResult
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidCache
import java.time.Instant

sealed interface SimResultPhase {
    data object Loading : SimResultPhase
    data class Loaded(val model: SimResultModel, val offlineSince: Instant?) : SimResultPhase
    data object Failed : SimResultPhase
}

/** Loads one result (F5.3), finds its list row for the headline, and marks it seen once per open. */
class SimResultStore(
    val id: String,
    private val row: SimRow?,
    private val service: SimsService,
    private val cache: RaidCache,
    private val specs: SpecCatalog,
    private val res: Resources,
    private val now: () -> Instant,
    private val markSeen: suspend (String) -> Unit,
) {
    var phase: SimResultPhase by mutableStateOf(SimResultPhase.Loading)
        private set

    private var seenPosted = false
    private val cacheKey get() = "sim-$id"

    suspend fun load() {
        when (val answer = service.result(id)) {
            is ApiResult.Success -> {
                cache.write(cacheKey, SimResult.serializer(), answer.data, now())
                phase = SimResultPhase.Loaded(model(answer.data), null)
                if (!seenPosted) {
                    seenPosted = true
                    markSeen(id)
                }
            }
            is ApiResult.Failure -> {
                val cached = cache.read(cacheKey, SimResult.serializer())
                phase = if (answer.error is ApiError.Network && cached != null) {
                    SimResultPhase.Loaded(model(cached.value), cached.savedAt)
                } else {
                    SimResultPhase.Failed
                }
            }
        }
    }

    /** The list row: the one passed in, else the first page's (a deep link or a push). */
    private suspend fun model(result: SimResult): SimResultModel {
        val found = row ?: (service.list(1, null) as? ApiResult.Success)?.data?.rows?.firstOrNull { it.simId == id }
        return SimResultModel.of(id, result, found, specs, res)
    }
}
