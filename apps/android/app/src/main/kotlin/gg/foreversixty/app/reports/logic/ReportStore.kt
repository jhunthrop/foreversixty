package gg.foreversixty.app.reports.logic

import gg.foreversixty.app.common.StripDecision
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.common.CacheLimits
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.CoachReview
import gg.foreversixty.app.model.FightMeters
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.model.ReportMe
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidCache
import java.time.Instant

sealed interface ReportPhase {
    data object Loading : ReportPhase
    /** [offlineSince] is set while the saved report shows under a strip: offline, or [stale] (W3 §5.7). */
    data class Loaded(val report: Report, val offlineSince: Instant?, val stale: Boolean = false) : ReportPhase
    data object Failed : ReportPhase

    /** A 404: private, removed, or not the viewer's to see (spec F10.5). */
    data object NotYours : ReportPhase
}

/** The `Your parses` card: hidden signed out, else loading, loaded or failed in the card only. */
sealed interface ParsesPhase {
    data object Hidden : ParsesPhase
    data object Loading : ParsesPhase
    data class Loaded(val me: ReportMe) : ParsesPhase
    data object Failed : ParsesPhase
}

/** The coach card (spec F10.4): hidden unless Premium with a night line. */
sealed interface CoachPhase {
    data object Hidden : CoachPhase
    data object Loading : CoachPhase
    data class Ready(val review: CoachReview) : CoachPhase
    data object Writing : CoachPhase
    data object Preparing : CoachPhase
    data object Failed : CoachPhase

    companion object {
        fun of(review: CoachReview): CoachPhase = when (review.status) {
            "ready" -> Ready(review)
            "writing" -> Writing
            "preparing" -> Preparing
            else -> Failed
        }
    }
}

/** One coach section, only those the review holds, in the card's order. */
sealed interface CoachSection {
    data class FixFirst(val title: String, val why: String) : CoachSection
    data class Practice(val text: String) : CoachSection
    data class Then(val items: List<Pair<String, String>>) : CoachSection
    data class DoingWell(val lines: List<String>) : CoachSection
    data class CouldntCheck(val lines: List<String>) : CoachSection

    companion object {
        fun of(review: CoachReview): List<CoachSection> {
            val doc = review.review
            val fix = doc.fixFirst
            return listOfNotNull(
                fix?.let { FixFirst(it.title, it.why) },
                fix?.drillText?.takeIf { it.isNotEmpty() }?.let(::Practice),
                doc.thenFix.takeIf { it.isNotEmpty() }?.let { list -> Then(list.map { it.title to it.why }) },
                doc.doingWell.takeIf { it.isNotEmpty() }?.let { list -> DoingWell(list.map { it.text }) },
                doc.couldntCheck.takeIf { it.isNotEmpty() }?.let(::CouldntCheck),
            )
        }
    }
}

/**
 * One report (F10.2): the report itself (cache first, offline from cache), the viewer's parses,
 * and the coach card. A free or signed-out viewer makes no coach call at all.
 */
class ReportStore(
    val id: String,
    private val service: ReportsService,
    private val cache: RaidCache,
    private val viewer: ReportsViewer,
    private val connectivity: Connectivity = Connectivity.Unknown,
    private val now: () -> Instant,
) {
    var phase: ReportPhase by mutableStateOf(ReportPhase.Loading)
        private set
    var parses: ParsesPhase by mutableStateOf(if (viewer.signedIn) ParsesPhase.Loading else ParsesPhase.Hidden)
        private set
    var coach: CoachPhase by mutableStateOf(CoachPhase.Hidden)
        private set

    private val cacheKey get() = "report-$id"
    private val loaded: Report? get() = (phase as? ReportPhase.Loaded)?.report

    /** The whole screen: the report, then the parses and, with them, the coach. */
    suspend fun load() {
        loadReport()
        if (phase is ReportPhase.Loaded) loadParses()
    }

    /** Whether the open screen keeps polling (a live report). */
    val isLive: Boolean get() = loaded?.let(Reports::polls) ?: false

    suspend fun loadReport() {
        if (phase == ReportPhase.Loading) {
            cache.read(cacheKey, Report.serializer())?.let { phase = ReportPhase.Loaded(it.value, null) }
        }
        when (val answer = service.report(id)) {
            is ApiResult.Success -> {
                cache.write(cacheKey, Report.serializer(), answer.data, now())
                phase = ReportPhase.Loaded(answer.data, null)
            }
            is ApiResult.Failure -> {
                val error = answer.error
                val cached = cache.read(cacheKey, Report.serializer())
                val saved = StripDecision.savedView(error, cached?.value, cached?.savedAt, now(), CacheLimits.REPORTS, connectivity.isConnected())
                phase = when {
                    error is ApiError.Server && error.status == NOT_FOUND -> ReportPhase.NotYours
                    saved != null -> ReportPhase.Loaded(saved.value, saved.since, saved.stale)
                    loaded != null -> phase
                    else -> ReportPhase.Failed
                }
            }
        }
    }

    suspend fun loadParses() {
        if (!viewer.signedIn) return
        if (parses !is ParsesPhase.Loaded) parses = ParsesPhase.Loading
        when (val answer = service.me(id)) {
            is ApiResult.Success -> {
                parses = ParsesPhase.Loaded(answer.data)
                loadCoach(answer.data)
            }
            is ApiResult.Failure -> if (parses !is ParsesPhase.Loaded) parses = ParsesPhase.Failed
        }
    }

    /** Premium only: the first character with fights, their night line, then that pull's review. */
    suspend fun loadCoach(me: ReportMe) {
        if (!viewer.logCoach) return
        val player = Parses.coachPlayer(me) ?: return
        val night = when (val answer = service.coachNight(id, player)) {
            is ApiResult.Success -> answer.data ?: run { coach = CoachPhase.Hidden; return }
            is ApiResult.Failure -> { coach = CoachPhase.Failed; return }
        }
        if (coach !is CoachPhase.Ready) coach = CoachPhase.Loading
        coach = when (val review = service.coachFight(id, night.fightIndex, player)) {
            is ApiResult.Success -> CoachPhase.of(review.data)
            is ApiResult.Failure -> CoachPhase.Failed
        }
    }

    companion object {
        private const val NOT_FOUND = 404
    }
}

sealed interface MetersPhase {
    data object Loading : MetersPhase
    data class Loaded(val meters: FightMeters) : MetersPhase
    data object Failed : MetersPhase

    /** 404 `flagged`: `This report cannot be shown.` */
    data object Flagged : MetersPhase
}

/** The fight sheet (F10.3): one segment's meters at a time; Show all lasts for the sheet. */
class MetersStore(
    private val reportId: String,
    val target: MetersTarget,
    private val service: ReportsService,
) {
    var metric: Metric by mutableStateOf(Metric.Damage)
        private set
    var phase: MetersPhase by mutableStateOf(MetersPhase.Loading)
        private set
    var showAll: Boolean by mutableStateOf(false)

    suspend fun select(next: Metric) {
        metric = next
        load()
    }

    suspend fun load() {
        phase = MetersPhase.Loading
        phase = when (val answer = service.meters(reportId, target.index, metric, target.scope)) {
            is ApiResult.Success -> MetersPhase.Loaded(answer.data)
            is ApiResult.Failure -> {
                val error = answer.error
                if (error is ApiError.Server && error.code == FLAGGED) MetersPhase.Flagged else MetersPhase.Failed
            }
        }
    }

    private companion object {
        const val FLAGGED = "flagged"
    }
}
