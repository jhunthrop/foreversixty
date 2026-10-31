package gg.foreversixty.app.reports

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.CoachNight
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.reports.logic.CoachPhase
import gg.foreversixty.app.reports.logic.CoachSection
import gg.foreversixty.app.reports.logic.FightRow
import gg.foreversixty.app.reports.logic.FightRows
import gg.foreversixty.app.reports.logic.Meters
import gg.foreversixty.app.reports.logic.MetersPhase
import gg.foreversixty.app.reports.logic.MetersStore
import gg.foreversixty.app.reports.logic.MetersTarget
import gg.foreversixty.app.reports.logic.Metric
import gg.foreversixty.app.reports.logic.ParseTier
import gg.foreversixty.app.reports.logic.Parses
import gg.foreversixty.app.reports.logic.ParsesPhase
import gg.foreversixty.app.reports.logic.ReportPhase
import gg.foreversixty.app.reports.logic.ReportState
import gg.foreversixty.app.reports.logic.ReportStore
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.logic.ReportsCopy
import gg.foreversixty.app.reports.logic.ReportsPhase
import gg.foreversixty.app.reports.logic.ReportsStore
import gg.foreversixty.app.reports.logic.ReportsViewer
import java.io.IOException
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The one F10 test list (spec §5.2), mirrored in iOS's ReportsLogicTests and ReportStoreTests. */
@RunWith(AndroidJUnit4::class)
class ReportsLogicTest {
    private val now = ReportsTestData.now
    private fun copy() = ReportsCopy(resources(), ReportsTestData.zone)
    private val offline = ApiError.Network(IOException("offline"))
    private fun store(service: FakeReportsService, viewer: ReportsViewer = ReportsTestData.viewer(), cache: MemoryRaidCache = MemoryRaidCache()) =
        ReportStore(ReportsTestData.ID, service, cache, viewer) { now }

    // F10.1 list

    @Test fun `sections split at seven days`() {
        val sections = Reports.sections(ReportsTestData.list().rows, now)
        assertEquals(listOf("m4xk9d2pqz7a", "ewlrx3enxbeb", "q7c2v9tk4mzh"), sections.thisWeek.map { it.id })
        assertEquals(listOf("t3n8b5xw2kqe"), sections.earlier.map { it.id })
    }

    @Test fun `row words follow the state`() {
        val rows = ReportsTestData.list().rows.associateBy { it.id }
        assertEquals("Updated 20 s ago", copy().rowMeta(rows.getValue("m4xk9d2pqz7a"), now))
        assertEquals("14 Dec · 1 boss · 1 kill", copy().rowMeta(rows.getValue("ewlrx3enxbeb"), now))
        assertEquals("29 Nov · could not be read", copy().rowMeta(rows.getValue("t3n8b5xw2kqe"), now))
        assertEquals(ReportState.Processing, ReportState.of(rows.getValue("q7c2v9tk4mzh").status))
        assertEquals("Onyxia, Monday", Reports.name(rows.getValue("ewlrx3enxbeb")))
        assertEquals("Blackrock Spire", Reports.name(rows.getValue("m4xk9d2pqz7a")))
    }

    @Test fun `row label reads the date, counts, visibility and live age`() {
        val live = ReportsTestData.list().rows.first()
        assertEquals("Blackrock Spire, Tuesday 15 December, 0 bosses, 0 kills, Private, live, updated 20 seconds ago", copy().rowLabel(live, now))
    }

    @Test fun `crest is the logging character's class, else the current character's`() {
        val row = ReportsTestData.list().rows.first()
        assertEquals("warrior", Reports.crestClass(row, ReportsTestData.viewer()))
        val stranger = row.copy(loggingCharacter = row.loggingCharacter?.copy(key = "us/pvp/someone"))
        assertEquals("warrior", Reports.crestClass(stranger, ReportsTestData.viewer()))
        assertNull(Reports.crestClass(stranger, ReportsViewer.SignedOut))
    }

    @Test fun `list keeps the cache offline and says signed out on 401`() = runTest {
        val cache = MemoryRaidCache()
        val service = FakeReportsService()
        ReportsStore(service, cache, { now }).load()
        service.listAnswer = { fail(offline) }
        val offlineStore = ReportsStore(service, cache, { now })
        offlineStore.load()
        assertEquals(now, (offlineStore.phase as ReportsPhase.Loaded).content.offlineSince)
        service.listAnswer = { fail(ApiError.Server(401, "unauthorized", "", emptyMap(), "t")) }
        val signedOut = ReportsStore(service, MemoryRaidCache(), { now })
        signedOut.load()
        assertEquals(ReportsPhase.SignedOut, signedOut.phase)
        service.listAnswer = { fail(offline) }
        val failed = ReportsStore(service, MemoryRaidCache(), { now })
        failed.load()
        assertEquals(ReportsPhase.Failed, failed.phase)
    }

    @Test fun `next page appends and a failed page says so`() = runTest {
        val service = FakeReportsService()
        val first = ReportsTestData.list()
        service.listAnswer = { page -> ok(first.copy(total = first.rows.size * 2, page = page)) }
        val store = ReportsStore(service, MemoryRaidCache(), { now })
        store.load()
        store.loadMore()
        assertEquals(first.rows.size * 2, (store.phase as ReportsPhase.Loaded).content.rows.size)
        assertEquals(listOf("list 1", "list 2"), service.calls)
    }

    // F10.2 fights

    @Test fun `fights are boss pulls in order then the zone's one trash row`() {
        val rows = FightRows.of(ReportsTestData.report())
        assertEquals(listOf("boss 4", "boss 6", "boss 9", "trash 1"), rows.map(::describe))
        val trash = rows.last() as FightRow.Trash
        assertEquals("Trash · Onyxia's Lair", copy().trashTitle(trash.group.zone))
        assertEquals("6 pulls · 7:40", copy().trashMeta(trash))
        assertEquals(3, FightRows.bossPulls(ReportsTestData.report()))
        assertEquals("3 boss pulls", copy().bossPulls(3))
        assertEquals("1 boss pull", copy().bossPulls(1))
        assertEquals("No boss pulls", copy().bossPulls(0))
    }

    @Test fun `two zones each keep their trash after their own boss pulls`() {
        val rows = FightRows.of(ReportsTestData.twoZones())
        assertEquals(listOf("boss 4", "boss 6", "boss 9", "trash 1", "trash 10"), rows.map(::describe))
    }

    @Test fun `trash with no group keeps its own row`() {
        val report = ReportsTestData.report().let { r -> r.copy(trashGroups = emptyList()) }
        assertEquals(listOf("pull 1", "pull 2", "pull 3", "boss 4", "pull 5", "boss 6", "pull 7", "pull 8", "boss 9"), FightRows.of(report).map(::describe))
    }

    @Test fun `kill and wipe meta`() {
        val bosses = FightRows.of(ReportsTestData.report()).filterIsInstance<FightRow.Boss>()
        assertEquals("Wipe · 1:12 · boss at 61%", copy().bossMeta(bosses[0]))
        assertEquals("Wipe · 4:52 · boss at 22%", copy().bossMeta(bosses[1]))
        assertEquals("Kill · 4:28 · 40 players · 3 deaths", copy().bossMeta(bosses[2]))
        val unseen = FightRow.Boss(bosses[0].fight.copy(bossHealthPct = -1.0))
        assertEquals("Wipe · 1:12", copy().bossMeta(unseen))
        assertEquals("Wipe · 1:12", copy().bossMeta(FightRow.Boss(bosses[0].fight.copy(bossHealthPct = null))))
        assertEquals("Onyxia, wipe, 1 minute, 12 seconds", copy().bossLabel(FightRow.Boss(unseen.fight)))
        assertEquals("Onyxia · pull 1 of 3, wipe, 1 minute, 12 seconds, boss at 61 percent", copy().bossLabel(bosses[0]))
    }

    @Test fun `a boss pulled more than once is numbered in fight order`() {
        val bosses = FightRows.of(ReportsTestData.report()).filterIsInstance<FightRow.Boss>()
        assertEquals(listOf("Onyxia · pull 1 of 3", "Onyxia · pull 2 of 3", "Onyxia · pull 3 of 3"), bosses.map { copy().bossTitle(it.fight, it.pull) })
        val single = FightRows.of(ReportsTestData.onlyWipes().let { r -> r.copy(fights = r.fights.filter { it.index != 6 }) })
        assertEquals(listOf("Onyxia"), single.filterIsInstance<FightRow.Boss>().map { copy().bossTitle(it.fight, it.pull) })
        val kill = bosses.last()
        assertEquals("Onyxia · pull 3 of 3", copy().headName(MetersTarget.Pull(kill.fight, kill.pull)))
    }

    @Test fun `home's last raid is the newest finished report with a boss pull`() {
        val rows = ReportsTestData.list().rows
        assertEquals("ewlrx3enxbeb", Reports.lastRaid(rows)?.id)
        assertNull(Reports.lastRaid(rows.filter { it.id != "ewlrx3enxbeb" }))
    }

    @Test fun `home's logs line names the last raid, never the live run`() = runTest {
        val store = ReportsStore(FakeReportsService(), MemoryRaidCache(), { now })
        assertEquals("Onyxia's Lair", store.lastRaid()?.zone)
        val none = FakeReportsService().apply { listAnswer = { ok(ReportsTestData.emptyList()) } }
        assertNull(ReportsStore(none, MemoryRaidCache(), { now }).lastRaid())
    }

    // Parses

    @Test fun `parse ladder floors`() {
        assertEquals(ParseTier.Grey, ParseTier.of(24))
        assertEquals(ParseTier.Green, ParseTier.of(25))
        assertEquals(ParseTier.Blue, ParseTier.of(50))
        assertEquals(ParseTier.Purple, ParseTier.of(75))
        assertEquals(ParseTier.Purple, ParseTier.of(94))
        assertEquals(ParseTier.Orange, ParseTier.of(95))
        assertEquals(ParseTier.Pink, ParseTier.of(99))
        assertEquals(ParseTier.Pink, ParseTier.of(100))
    }

    @Test fun `parse row words`() {
        val row = Parses.rows(ReportsTestData.me()).single()
        assertEquals("Kill 4:28 · Obnoxious Yell, Fury", copy().parseMeta(row))
        assertEquals("712", Parses.figure(row.fight.value))
        assertEquals("DPS", copy().unit(row.fight.unit))
        assertEquals("Onyxia, kill, 4 minutes, 28 seconds, 712 damage per second, 91st percentile", copy().parseLabel(row))
        assertTrue(Parses.rows(ReportsTestData.meEmpty()).isEmpty())
    }

    // Meters

    @Test fun `meter bars are each value over the top`() {
        val rows = Meters.rows(ReportsTestData.meters().players, Metric.Damage)
        assertEquals(1f, rows.first().fraction)
        val mine = rows.single { it.player.mine }
        assertEquals(6, mine.rank)
        assertEquals((mine.player.dps / rows.first().player.dps).toFloat(), mine.fraction)
        assertEquals("Rank 6, Obnoxious Yell, Fury Warrior, 712, your character", copy().meterLabel(mine))
    }

    @Test fun `eleven rows then show all, four at the largest text with your row after them`() {
        val rows = Meters.rows(ReportsTestData.meters().players, Metric.Damage)
        assertEquals(11, Meters.visible(rows, Meters.cut(1f), showAll = false).size)
        assertTrue(Meters.hasMore(rows, Meters.cut(1f), showAll = false))
        val large = Meters.visible(rows, Meters.cut(2f), showAll = false)
        assertEquals(listOf(1, 2, 3, 4, 6), large.map { it.rank })
        assertTrue(large.last().player.mine)
        assertTrue(Meters.hasMore(rows, Meters.cut(2f), showAll = false))
        assertEquals(rows.size, Meters.visible(rows, Meters.cut(1f), showAll = true).size)
        assertFalse(Meters.hasMore(rows, Meters.cut(1f), showAll = true))
    }

    @Test fun `meters ask by segment and trash scope, flagged and failed apart`() = runTest {
        val service = FakeReportsService()
        val group = requireNotNull(ReportsTestData.report().trashGroups).first()
        val trash = MetersStore(ReportsTestData.ID, MetersTarget.Trash(group), service)
        trash.load()
        trash.select(Metric.Healing)
        assertEquals(listOf("meters 1 damage trash", "meters 1 healing trash"), service.calls)
        service.metersAnswer = { fail(ApiError.Server(404, "flagged", "", emptyMap(), "t")) }
        trash.load()
        assertEquals(MetersPhase.Flagged, trash.phase)
        service.metersAnswer = { fail(offline) }
        trash.load()
        assertEquals(MetersPhase.Failed, trash.phase)
        assertEquals("6 pulls · 7:40", copy().metersHead(trash.target))
        val kill = ReportsTestData.report().fights.single { it.index == ReportsTestData.KILL_FIGHT }
        assertEquals("Kill · 4:28 · 40 players", copy().metersHead(MetersTarget.Pull(kill)))
        // Healing leaves off players with none of it; the head still counts the pull's roster.
        assertTrue(ReportsTestData.healing().players.size < kill.players.size)
        assertEquals("Kill · 4:28 · 40 players", copy().metersHead(MetersTarget.Pull(kill)))
    }

    // Report, parses and coach

    @Test fun `premium viewer gets the coach of the first character with fights`() = runTest {
        val service = FakeReportsService()
        val store = store(service)
        store.load()
        assertTrue(store.parses is ParsesPhase.Loaded)
        assertTrue(store.coach is CoachPhase.Ready)
        assertEquals(listOf("report", "me", "night ${ReportsTestData.ME_KEY}", "coach 9 ${ReportsTestData.ME_KEY}"), service.calls)
    }

    @Test fun `free viewer makes no coach call`() = runTest {
        val service = FakeReportsService()
        val store = store(service, ReportsTestData.viewer(premium = false))
        store.load()
        assertEquals(CoachPhase.Hidden, store.coach)
        assertEquals(listOf("report", "me"), service.calls)
    }

    @Test fun `signed out shows no parses and calls nothing but the report`() = runTest {
        val service = FakeReportsService()
        val store = store(service, ReportsViewer.SignedOut)
        store.load()
        assertEquals(ParsesPhase.Hidden, store.parses)
        assertEquals(listOf("report"), service.calls)
    }

    @Test fun `no night line means no card, coach states and errors stay in the card`() = runTest {
        val service = FakeReportsService()
        service.nightAnswer = { ok<CoachNight?>(null) }
        val none = store(service)
        none.load()
        assertEquals(CoachPhase.Hidden, none.coach)
        service.nightAnswer = { ok(ReportsTestData.night()) }
        service.coachAnswer = { ok(ReportsTestData.coach().copy(status = "writing")) }
        val writing = store(service)
        writing.load()
        assertEquals(CoachPhase.Writing, writing.coach)
        service.coachAnswer = { ok(ReportsTestData.coach().copy(status = "preparing")) }
        val preparing = store(service)
        preparing.load()
        assertEquals(CoachPhase.Preparing, preparing.coach)
        service.coachAnswer = { fail(offline) }
        val failed = store(service)
        failed.load()
        assertEquals(CoachPhase.Failed, failed.coach)
        assertTrue(failed.parses is ParsesPhase.Loaded)
    }

    @Test fun `coach sections only those present`() {
        val review = ReportsTestData.coach()
        val all = CoachSection.of(review)
        assertEquals(5, all.size)
        val bare = review.copy(review = review.review.copy(fixFirst = null, thenFix = emptyList(), couldntCheck = emptyList()))
        assertEquals(listOf(CoachSection.DoingWell::class), CoachSection.of(bare).map { it::class })
        val labels = review.basisLabels.map { it.trimEnd('.') }.distinct()
        assertEquals("Reviewed 14 Dec · Based on: ${labels.joinToString("; ")}.", copy().coachFooter(review.reviewedAt, review.basisLabels))
        assertEquals("Reviewed 14 Dec · Based on: a; b.", copy().coachFooter(review.reviewedAt, listOf("a.", "b.", "a.")))
    }

    @Test fun `parses failing stay in their card`() = runTest {
        val service = FakeReportsService()
        service.meAnswer = { fail(offline) }
        val store = store(service)
        store.load()
        assertTrue(store.phase is ReportPhase.Loaded)
        assertEquals(ParsesPhase.Failed, store.parses)
    }

    @Test fun `404 is not yours, offline shows the cache`() = runTest {
        val service = FakeReportsService()
        service.reportAnswer = { fail(RaidFixture.error("reports.get.error-not-found", 404)) }
        val missing = store(service)
        missing.load()
        assertEquals(ReportPhase.NotYours, missing.phase)
        val cache = MemoryRaidCache()
        store(FakeReportsService(), cache = cache).load()
        service.reportAnswer = { fail(offline) }
        val cached = store(service, cache = cache)
        cached.load()
        assertEquals(now, (cached.phase as ReportPhase.Loaded).offlineSince)
    }

    @Test fun `only a live report polls`() = runTest {
        val service = FakeReportsService()
        val done = store(service)
        done.loadReport()
        assertFalse(done.isLive)
        service.reportAnswer = { ok(ReportsTestData.report().copy(status = "live")) }
        val live = store(service)
        live.loadReport()
        assertTrue(live.isLive)
        assertEquals(20L, Reports.LIVE_POLL.inWholeSeconds)
    }

    private fun describe(row: FightRow): String = when (row) {
        is FightRow.Boss -> "boss ${row.fight.index}"
        is FightRow.Trash -> "trash ${row.group.key}"
        is FightRow.LoosePull -> "pull ${row.fight.index}"
    }
}
