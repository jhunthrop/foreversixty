package gg.foreversixty.app.reports

import gg.foreversixty.app.model.CoachNight
import gg.foreversixty.app.model.CoachReview
import gg.foreversixty.app.model.FightMeters
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.Report
import gg.foreversixty.app.model.ReportList
import gg.foreversixty.app.model.ReportMe
import gg.foreversixty.app.model.TrashGroup
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.reports.logic.Metric
import gg.foreversixty.app.reports.logic.ReportsService
import gg.foreversixty.app.reports.logic.ReportsViewer
import java.time.Instant
import java.time.ZoneId

/**
 * F10's test data, the same story as iOS's ReportsTestData: the contract fixtures (the guild Onyxia
 * night ewlrx3enxbeb), plus variants derived from them with the arithmetic kept. No figure is typed here.
 */
object ReportsTestData {
    const val ID = "ewlrx3enxbeb"
    const val ME_KEY = "us/pvp/obnoxious-yell"

    /** Twenty seconds after the live report's heartbeat (2026-12-16 00:31:40Z). */
    val now: Instant = Instant.parse("2026-12-16T00:32:00Z")
    val zone: ZoneId = ZoneId.of("America/Chicago")

    fun list(): ReportList = RaidFixture.load(ReportList.serializer(), "reports.list")
    fun emptyList(): ReportList = RaidFixture.load(ReportList.serializer(), "reports.list.empty")
    fun report(): Report = RaidFixture.load(Report.serializer(), "reports.get")
    fun me(): ReportMe = RaidFixture.load(ReportMe.serializer(), "reports.me")
    fun meEmpty(): ReportMe = RaidFixture.load(ReportMe.serializer(), "reports.me.empty")
    fun meters(): FightMeters = RaidFixture.load(FightMeters.serializer(), "reports.meters")
    fun night(): CoachNight = RaidFixture.load(CoachNight.serializer(), "coach.night")
    fun coach(): CoachReview = RaidFixture.load(CoachReview.serializer(), "coach.fight")
    fun account(premium: Boolean): Me = RaidFixture.load(Me.serializer(), if (premium) "me.premium" else "me")

    fun viewer(premium: Boolean = true): ReportsViewer = ReportsViewer.of(account(premium), ME_KEY)

    /** The night without its kill: fight 9 and its /me row dropped. */
    fun onlyWipes(): Report = report().let { r -> r.copy(fights = r.fights.filter { it.index != KILL_FIGHT }) }

    /**
     * The night with a second zone after it: the Onyxia fights, then Molten Core trash as the
     * server folds it, its pulls copied from the Onyxia trash pulls with their own lengths.
     */
    fun twoZones(): Report {
        val base = report()
        val trash = base.fights.filter { it.kind == "trash" }.take(MC_PULLS)
        val offset = base.fights.maxOf { it.index }
        val mc = trash.mapIndexed { i, f -> f.copy(index = offset + i + 1, zone = MOLTEN_CORE) }
        val group = TrashGroup(mc.first().index, MOLTEN_CORE, mc.map { it.index }, mc.sumOf { it.durationMs })
        return base.copy(fights = base.fights + mc, trashGroups = base.trashGroups.orEmpty() + group)
    }

    /** Healing: the fixture's players sorted by hps, those with none left off (as the server does). */
    fun healing(): FightMeters = meters().let { m -> m.copy(metric = "healing", players = m.players.filter { it.hps > 0 }.sortedByDescending { it.hps }) }

    fun taken(): FightMeters = meters().let { m -> m.copy(metric = "taken", players = m.players.sortedByDescending { it.dtps }) }

    /** The trash view: the six Onyxia trash pulls folded (`reports.meters.trash`). */
    fun trashMeters(): FightMeters = RaidFixture.load(FightMeters.serializer(), "reports.meters.trash")

    const val KILL_FIGHT = 9
    private const val MC_PULLS = 3
    private const val MOLTEN_CORE = "Molten Core"
}

/** A scripted reports service: each call answers from its lambda; every call is recorded. */
class FakeReportsService : ReportsService {
    var listAnswer: (Int) -> ApiResult<ReportList> = { ok(ReportsTestData.list()) }
    var reportAnswer: () -> ApiResult<Report> = { ok(ReportsTestData.report()) }
    var meAnswer: () -> ApiResult<ReportMe> = { ok(ReportsTestData.me()) }
    var metersAnswer: (Metric) -> ApiResult<FightMeters> = { ok(ReportsTestData.meters()) }
    var nightAnswer: () -> ApiResult<CoachNight?> = { ok(ReportsTestData.night()) }
    var coachAnswer: () -> ApiResult<CoachReview> = { ok(ReportsTestData.coach()) }
    val calls = mutableListOf<String>()

    override suspend fun list(page: Int): ApiResult<ReportList> {
        calls += "list $page"
        return listAnswer(page)
    }
    override suspend fun report(id: String): ApiResult<Report> {
        calls += "report"
        return reportAnswer()
    }
    override suspend fun me(id: String): ApiResult<ReportMe> {
        calls += "me"
        return meAnswer()
    }
    override suspend fun meters(id: String, n: Int, metric: Metric, scope: String?): ApiResult<FightMeters> {
        calls += "meters $n ${metric.query} ${scope.orEmpty()}".trim()
        return metersAnswer(metric)
    }
    override suspend fun coachNight(id: String, player: String): ApiResult<CoachNight?> {
        calls += "night $player"
        return nightAnswer()
    }
    override suspend fun coachFight(id: String, n: Int, player: String): ApiResult<CoachReview> {
        calls += "coach $n $player"
        return coachAnswer()
    }
}
