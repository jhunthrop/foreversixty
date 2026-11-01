package gg.foreversixty.app.reports.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.CoachNight
import gg.foreversixty.app.model.CoachReview
import gg.foreversixty.app.model.FightMeters
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.model.ReportList
import gg.foreversixty.app.model.ReportMe
import gg.foreversixty.app.model.ReportRow
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.Endpoint
import gg.foreversixty.app.raids.logic.RaidCache
import java.time.Instant
import kotlinx.serialization.builtins.nullable

/** The report calls F10 makes. Live goes through [ApiClient]; tests pass a fake. */
interface ReportsService {
    suspend fun list(page: Int): ApiResult<ReportList>
    suspend fun report(id: String): ApiResult<Report>
    suspend fun me(id: String): ApiResult<ReportMe>
    suspend fun meters(id: String, n: Int, metric: Metric, scope: String?): ApiResult<FightMeters>

    /** The night's coach line; a success with null data is "no line" (no card). */
    suspend fun coachNight(id: String, player: String): ApiResult<CoachNight?>
    suspend fun coachFight(id: String, n: Int, player: String): ApiResult<CoachReview>
}

class LiveReportsService(private val client: ApiClient) : ReportsService {
    override suspend fun list(page: Int) = client.send(ApiEndpoints.listMyReports(page = page))
    override suspend fun report(id: String) = client.send(ApiEndpoints.getReport(id, fold = TRASH_FOLD))
    override suspend fun me(id: String) = client.send(ApiEndpoints.getReportMe(id))
    override suspend fun meters(id: String, n: Int, metric: Metric, scope: String?) =
        client.send(ApiEndpoints.getFightMeters(id, n, metric.query, scope))
    override suspend fun coachNight(id: String, player: String): ApiResult<CoachNight?> =
        client.send(ApiEndpoints.getCoachNight(id, player).nullableData())
    override suspend fun coachFight(id: String, n: Int, player: String) = client.send(ApiEndpoints.getFightCoach(id, n, player))

    private companion object {
        const val TRASH_FOLD = "trash"

        /** The same call, decoding `data: null` as null instead of failing. */
        fun <T : Any> Endpoint<T>.nullableData(): Endpoint<T?> = Endpoint(method, path, response.nullable, query, body)
    }
}

/** What Your reports draws (spec F10.1). */
data class ReportsContent(
    val rows: List<ReportRow>,
    val total: Int,
    val page: Int,
    /** Set while the cached answer shows because a refresh could not connect. */
    val offlineSince: Instant?,
) {
    val hasMorePages: Boolean get() = rows.size < total
}

sealed interface ReportsPhase {
    data object Loading : ReportsPhase
    data class Loaded(val content: ReportsContent) : ReportsPhase
    data object Failed : ReportsPhase
    data object SignedOut : ReportsPhase
}

/** Loads Your reports (cache first), pages on, and keeps what is drawn when a refresh fails. */
class ReportsStore(
    private val service: ReportsService,
    private val cache: RaidCache,
    private val now: () -> Instant,
    initial: ReportsPhase = ReportsPhase.Loading,
) {
    var phase: ReportsPhase by mutableStateOf(initial)
        private set

    /** The last next-page read failed: the list ends with the error and Try again. */
    var pageFailed: Boolean by mutableStateOf(false)
        private set

    private val content: ReportsContent? get() = (phase as? ReportsPhase.Loaded)?.content

    suspend fun load() {
        if (phase == ReportsPhase.Loading) {
            cache.read(CACHE_KEY, ReportList.serializer())?.let { phase = ReportsPhase.Loaded(it.value.content(null)) }
        }
        when (val answer = service.list(1)) {
            is ApiResult.Success -> {
                cache.write(CACHE_KEY, ReportList.serializer(), answer.data, now())
                pageFailed = false
                phase = ReportsPhase.Loaded(answer.data.content(null))
            }
            is ApiResult.Failure -> fail(answer.error)
        }
    }

    private fun fail(error: ApiError) {
        val cached = cache.read(CACHE_KEY, ReportList.serializer())
        phase = when {
            error is ApiError.Server && error.isUnauthorized -> ReportsPhase.SignedOut
            error is ApiError.Network && cached != null -> ReportsPhase.Loaded(cached.value.content(cached.savedAt))
            // A refresh that fails keeps what is drawn; only a first load shows the error.
            content != null -> phase
            else -> ReportsPhase.Failed
        }
    }

    /** The next page of 100, when the list has one. */
    suspend fun loadMore() {
        val current = content ?: return
        if (current.offlineSince != null || !current.hasMorePages) return
        val next = (service.list(current.page + 1) as? ApiResult.Success)?.data
        pageFailed = next == null
        if (next == null) return
        phase = ReportsPhase.Loaded(current.copy(rows = current.rows + next.rows, total = next.total, page = next.page))
    }

    /** Home's Logs row: the last raid with a boss pull, from the network or the cache; null when there is none. */
    suspend fun lastRaid(): ReportRow? {
        val fresh = (service.list(1) as? ApiResult.Success)?.data
        if (fresh != null) cache.write(CACHE_KEY, ReportList.serializer(), fresh, now())
        return (fresh ?: cache.read(CACHE_KEY, ReportList.serializer())?.value)?.rows?.let(Reports::lastRaid)
    }

    companion object {
        const val CACHE_KEY = "reports"

        private fun ReportList.content(offlineSince: Instant?) = ReportsContent(rows, total, page, offlineSince)
    }
}
