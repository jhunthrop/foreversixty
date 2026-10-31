package gg.foreversixty.app.sims.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateMapOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.SimList
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimResult
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.model.SimStarted
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidCache
import java.time.Instant
import kotlinx.serialization.Serializable

/** The sim calls F5 makes, all on today's API. Live goes through [ApiClient]; tests pass a fake. */
interface SimsService {
    suspend fun list(page: Int, lane: String?): ApiResult<SimList>
    suspend fun result(id: String): ApiResult<SimResult>
    suspend fun progress(id: String): ApiResult<SimProgress>
    suspend fun rerun(id: String): ApiResult<SimStarted>
    suspend fun markSeen(id: String): ApiResult<Unit>
}

class LiveSimsService(private val client: ApiClient) : SimsService {
    override suspend fun list(page: Int, lane: String?) = client.send(ApiEndpoints.listMySims(page = page, lane = lane))
    override suspend fun result(id: String) = client.send(ApiEndpoints.getSim(id))
    override suspend fun progress(id: String) = client.send(ApiEndpoints.getSimProgress(id))
    override suspend fun rerun(id: String) = client.send(ApiEndpoints.rerunSim(id))
    override suspend fun markSeen(id: String) = client.send(ApiEndpoints.markSimSeen(id))
}

/** What Sims draws (spec F5.1): the server rows, your sims, and the offline stamp. */
data class SimsContent(
    val serverRows: List<SimRow>,
    val mine: List<SimRow>,
    /** `GET /v1/sims`'s total, so the list knows whether another page exists. */
    val total: Int,
    val page: Int,
    /** Set while the cached answer shows because a refresh could not connect. */
    val offlineSince: Instant?,
    /** A section that did not load says so on its own; the other keeps its rows (spec F5 States). */
    val serverFailed: Boolean = false,
    val mineFailed: Boolean = false,
) {
    val hasMorePages: Boolean get() = mine.size < total
}

sealed interface SimsPhase {
    data object Loading : SimsPhase
    data class Loaded(val content: SimsContent) : SimsPhase
    data object Failed : SimsPhase
}

/** The two first pages as cached together. */
@Serializable
data class CachedSims(val server: List<SimRow>, val mine: SimList) {
    fun content(offlineSince: Instant?) = SimsContent(server, mine.rows, mine.total, mine.page, offlineSince)
}

/** Loads both lists (cache first), polls in-flight runs, and runs a failed run again. */
class SimsStore(
    private val service: SimsService,
    private val cache: RaidCache,
    private val now: () -> Instant,
    initial: SimsPhase = SimsPhase.Loading,
) {
    var phase: SimsPhase by mutableStateOf(initial)
        private set

    /** The newest poll answer per in-flight run. */
    val progress = mutableStateMapOf<String, SimProgress>()

    /** What each failed row's `Run again` shows. */
    val rerun = mutableStateMapOf<String, RerunStatus>()

    /** The last next-page read failed: the list ends with the error and Try again. */
    var pageFailed: Boolean by mutableStateOf(false)
        private set

    private val content: SimsContent? get() = (phase as? SimsPhase.Loaded)?.content

    suspend fun load() {
        if (phase == SimsPhase.Loading) {
            cache.read(CACHE_KEY, CachedSims.serializer())?.let { phase = SimsPhase.Loaded(it.value.content(null)) }
        }
        val server = service.list(1, ServerRuns.SERVER_LANE)
        val mine = service.list(1, null)
        if (server is ApiResult.Success && mine is ApiResult.Success) {
            val answer = CachedSims(server.data.rows, mine.data)
            cache.write(CACHE_KEY, CachedSims.serializer(), answer, now())
            phase = SimsPhase.Loaded(answer.content(null))
            return
        }
        val error = (server as? ApiResult.Failure)?.error ?: (mine as ApiResult.Failure).error
        fail(error, (server as? ApiResult.Success)?.data, (mine as? ApiResult.Success)?.data)
    }

    private fun fail(error: ApiError, server: SimList?, mine: SimList?) {
        val cached = cache.read(CACHE_KEY, CachedSims.serializer())
        if (error is ApiError.Network && cached != null) {
            phase = SimsPhase.Loaded(cached.value.content(cached.savedAt))
        } else if (server != null || mine != null) {
            phase = SimsPhase.Loaded(
                SimsContent(
                    server?.rows.orEmpty(), mine?.rows.orEmpty(), mine?.total ?: 0, mine?.page ?: 1, null,
                    serverFailed = server == null, mineFailed = mine == null,
                ),
            )
        } else if (content == null) {
            // A refresh that fails keeps what is drawn; only a first load shows the error.
            phase = SimsPhase.Failed
        }
    }

    /** The next page of Your sims, when the list has one. */
    suspend fun loadMore() {
        val current = content ?: return
        if (current.offlineSince != null || !current.hasMorePages) return
        val next = (service.list(current.page + 1, null) as? ApiResult.Success)?.data
        pageFailed = next == null
        if (next == null) return
        phase = SimsPhase.Loaded(current.copy(mine = current.mine + next.rows, total = next.total, page = next.page))
    }

    /** One poll round: each in-flight run's progress; a run that ended reloads the lists. */
    suspend fun pollOnce() {
        val current = content ?: return
        var ended = false
        for (id in ServerRuns.inFlightIds(current.serverRows, now())) {
            val answer = (service.progress(id) as? ApiResult.Success)?.data ?: continue
            progress[id] = answer
            ended = ended || ServerRuns.hasEnded(answer)
        }
        if (ended) load()
    }

    /** Whether anything is still running, so the screen keeps polling. */
    val isPolling: Boolean get() = content?.let { ServerRuns.inFlightIds(it.serverRows, now()).isNotEmpty() } ?: false

    /** `Run again` on a failed row: Starting…, then Started and the new run in Server runs. */
    suspend fun runAgain(id: String) {
        if (rerun[id]?.isLocked == true) return
        rerun[id] = RerunStatus.Pending
        when (val answer = service.rerun(id)) {
            is ApiResult.Success -> {
                rerun[id] = RerunStatus.Started
                load()
            }
            is ApiResult.Failure -> rerun[id] = RerunStatus.refused(answer.error)
        }
    }

    /** Opening a finished result marks it seen, here and on the server. */
    suspend fun markSeen(id: String) {
        service.markSeen(id)
        val current = content ?: return
        fun List<SimRow>.seen() = map { if (it.simId == id) it.copy(seen = true) else it }
        phase = SimsPhase.Loaded(current.copy(serverRows = current.serverRows.seen(), mine = current.mine.seen()))
    }

    companion object {
        const val CACHE_KEY = "sims"
    }
}
