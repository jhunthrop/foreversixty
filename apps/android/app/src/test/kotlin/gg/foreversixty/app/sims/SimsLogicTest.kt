package gg.foreversixty.app.sims

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.sims.logic.RerunStatus
import gg.foreversixty.app.sims.logic.RunEstimate
import gg.foreversixty.app.sims.logic.ServerRunState
import gg.foreversixty.app.sims.logic.ServerRuns
import gg.foreversixty.app.sims.logic.SimResultModel
import gg.foreversixty.app.sims.logic.SimResultPhase
import gg.foreversixty.app.sims.logic.SimResultStore
import gg.foreversixty.app.sims.logic.SimsCopy
import gg.foreversixty.app.sims.logic.SimsPhase
import gg.foreversixty.app.sims.logic.SimsStore
import java.io.IOException
import java.time.Instant
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The one F5 test list (spec §5.5), mirrored in iOS's SimsLogicTests, SimsStoreTests and SimResultTests. */
@RunWith(AndroidJUnit4::class)
class SimsLogicTest {
    private val now = SimsTestData.now
    private fun copy() = SimsCopy(resources())
    private val offline = ApiError.Network(IOException("offline"))

    @Test fun `states follow the site`() {
        assertEquals(ServerRunState.DoneUnseen, ServerRuns.state(SimsTestData.row("km3xibdrsiyr"), now))
        assertEquals(ServerRunState.Running, ServerRuns.state(SimsTestData.row("q8wd2c9vh4tz"), now))
        assertEquals(ServerRunState.Failed, ServerRuns.state(SimsTestData.row("b7nf5x1pa3le"), now))
    }

    @Test fun `running turns overdue past one and a half estimates`() {
        val running = SimsTestData.row("q8wd2c9vh4tz")
        assertEquals(ServerRunState.Running, ServerRuns.state(running, Instant.parse("2026-12-15T05:22:55Z")))
        assertEquals(ServerRunState.RunningOverdue, ServerRuns.state(running, Instant.parse("2026-12-15T05:22:56Z")))
    }

    @Test fun `running without estimate has no suffix`() {
        val row = SimRow(simId = "x", state = "running", createdAt = now.toString())
        assertEquals(ServerRunState.RunningNoEstimate, ServerRuns.state(row, now))
        assertEquals("", copy().stateSuffix(row, ServerRunState.RunningNoEstimate))
    }

    @Test fun `estimate rounds to one unit`() {
        assertEquals(RunEstimate.AboutAMinute, RunEstimate.of(50))
        assertEquals(RunEstimate.Minutes(2), RunEstimate.of(90))
        assertEquals(RunEstimate.Minutes(4), RunEstimate.of(240))
        assertEquals(RunEstimate.Minutes(89), RunEstimate.of(89 * 60))
        assertEquals(RunEstimate.Hours(2), RunEstimate.of(90 * 60))
        assertEquals(" · takes about 2 minutes", copy().stateSuffix(SimsTestData.row("q8wd2c9vh4tz"), ServerRunState.Running))
    }

    @Test fun `visible is the newest five`() {
        val many = (0 until 7).map { SimRow(simId = "r$it", createdAt = now.plusSeconds(it.toLong()).toString()) }
        val (rows, hasMore) = ServerRuns.visible(many)
        assertEquals(listOf("r6", "r5", "r4", "r3", "r2"), rows.map { it.simId })
        assertTrue(hasMore)
        assertFalse(ServerRuns.visible(many.take(5)).second)
    }

    @Test fun `section shows for premium or old server rows`() {
        assertTrue(ServerRuns.showsSection(true, emptyList()))
        assertTrue(ServerRuns.showsSection(false, SimsTestData.list().rows))
        assertFalse(ServerRuns.showsSection(false, emptyList()))
    }

    @Test fun `poll asks only about runs in flight`() {
        assertEquals(listOf("q8wd2c9vh4tz"), ServerRuns.inFlightIds(SimsTestData.list().rows, now))
    }

    @Test fun `bar only for combinations`() {
        val plain = RaidFixture.load(SimProgress.serializer(), "sims.progress")
        assertNull(ServerRuns.fraction(plain))
        assertEquals(0.625f, ServerRuns.fraction(SimProgress("running", 0, null, 1, 40, 64)))
        assertTrue(ServerRuns.hasEnded(SimProgress("done", 1, null, 0, 0, 0)))
        assertFalse(ServerRuns.hasEnded(plain))
    }

    @Test fun `rerun 402 is a lapsed plan`() {
        assertEquals(RerunStatus.Lapsed, RerunStatus.refused(RaidFixture.error("sims.rerun.error-premium-required", 402)))
        assertEquals(RerunStatus.Failed, RerunStatus.refused(RaidFixture.error("sims.rerun.error-not-found", 404)))
        assertEquals(RerunStatus.Failed, RerunStatus.refused(offline))
        assertTrue(RerunStatus.Pending.isLocked && RerunStatus.Started.isLocked && !RerunStatus.Failed.isLocked)
    }

    @Test fun `failed line names the reason when there is one`() {
        assertEquals("Failed: the engine stopped", copy().failedLine(SimsTestData.row("b7nf5x1pa3le")))
        assertEquals("Failed", copy().failedLine(SimRow(simId = "x", state = "error")))
    }

    // Store

    private fun store(service: FakeSimsService, cache: MemoryRaidCache = MemoryRaidCache()) = SimsStore(service, cache, { now })

    private fun SimsStore.content() = (phase as SimsPhase.Loaded).content

    @Test fun `load reads server runs and your sims`() = runTest {
        val service = FakeSimsService()
        val store = store(service)
        store.load()
        assertEquals(3, store.content().serverRows.size)
        assertEquals(3, store.content().mine.size)
        assertFalse(store.content().hasMorePages)
        assertEquals(setOf("server", "all"), service.lists.map { it.second ?: "all" }.toSet())
    }

    @Test fun `offline shows the cache under the strip`() = runTest {
        val cache = MemoryRaidCache()
        store(FakeSimsService(), cache).load()
        val offlineService = FakeSimsService().apply { listAnswer = { _, _ -> fail(offline) } }
        val store = store(offlineService, cache)
        store.load()
        assertEquals(now, store.content().offlineSince)
    }

    @Test fun `server error with nothing cached fails`() = runTest {
        val store = store(FakeSimsService().apply { listAnswer = { _, _ -> fail(ApiError.Status(500)) } })
        store.load()
        assertEquals(SimsPhase.Failed, store.phase)
    }

    @Test fun `empty list loads empty`() = runTest {
        val service = FakeSimsService().apply { listAnswer = { _, _ -> ok(RaidFixture.load(gg.foreversixty.app.model.SimList.serializer(), "sims.list.empty")) } }
        val store = store(service)
        store.load()
        assertTrue(store.content().mine.isEmpty())
    }

    @Test fun `poll asks the running row and reloads when it ends`() = runTest {
        val service = FakeSimsService()
        val store = store(service)
        store.load()
        assertTrue(store.isPolling)
        store.pollOnce()
        assertEquals(listOf("q8wd2c9vh4tz"), service.polled)
        assertEquals(1200, store.progress["q8wd2c9vh4tz"]?.iterationsDone)
        assertEquals(2, service.lists.size)
        service.progressAnswer = { ok(SimProgress("done", 3000, 990.0, 0, 0, 0)) }
        store.pollOnce()
        assertEquals(4, service.lists.size)
    }

    @Test fun `run again starts then reloads`() = runTest {
        val service = FakeSimsService()
        val store = store(service)
        store.load()
        store.runAgain("b7nf5x1pa3le")
        assertEquals(listOf("b7nf5x1pa3le"), service.reruns)
        assertEquals(RerunStatus.Started, store.rerun["b7nf5x1pa3le"])
        store.runAgain("b7nf5x1pa3le")
        assertEquals(1, service.reruns.size)
    }

    @Test fun `run again on a lapsed plan says premium`() = runTest {
        val service = FakeSimsService().apply { rerunAnswer = { fail(RaidFixture.error("sims.rerun.error-premium-required", 402)) } }
        val store = store(service)
        store.load()
        store.runAgain("b7nf5x1pa3le")
        assertEquals(RerunStatus.Lapsed, store.rerun["b7nf5x1pa3le"])
    }

    @Test fun `mark seen clears the dot`() = runTest {
        val service = FakeSimsService()
        val store = store(service)
        store.load()
        store.markSeen("km3xibdrsiyr")
        assertEquals(listOf("km3xibdrsiyr"), service.seen)
        assertEquals(true, store.content().serverRows.first { it.simId == "km3xibdrsiyr" }.seen)
    }

    @Test fun `load more appends the next page`() = runTest {
        val service = FakeSimsService().apply {
            listAnswer = { page, _ ->
                val list = SimsTestData.list()
                ok(list.copy(rows = if (page == 1) list.rows else list.rows.take(1), total = 4, page = page))
            }
        }
        val store = store(service)
        store.load()
        assertTrue(store.content().hasMorePages)
        store.loadMore()
        assertEquals(4, store.content().mine.size)
        assertFalse(store.content().hasMorePages)
    }

    // Result

    private val specs = SimsTestData.specs()

    @Test fun `plain run draws the mean and the error line`() {
        val model = SimResultModel.of("km3xibdrsiyr", RaidFixture.load(gg.foreversixty.app.model.SimResult.serializer(), "sims.result"), SimsTestData.row("km3xibdrsiyr"), specs, resources())
        assertEquals("Tuesday", model.title)
        assertEquals("Sim", model.kind)
        assertEquals("Fury Warrior · level 60", model.meta(resources()))
        assertEquals(SimResultModel.Answer.Figure("101.4"), model.answer)
        assertEquals("± 0.1 · 3,000 iterations · engine 1f542def3", model.errorLine)
        assertEquals("https://foreversixty.gg/sim/km3xibdrsiyr", model.breakdownUrl)
    }

    @Test fun `untitled is named by spec`() {
        assertEquals("Fury", SimResultModel.of("x", SimsTestData.result(untitled = true), null, specs, resources()).title)
    }

    @Test fun `other kinds show the row headline and mark the changed slot`() {
        val model = SimResultModel.of("g", SimsTestData.result(topGear = true, untitled = true), SimsTestData.topGearRow, specs, resources())
        assertEquals("Top Gear", model.kind)
        assertEquals(SimResultModel.Answer.Sentence("+4 DPS from Howler's Furs"), model.answer)
        assertEquals(listOf("back"), model.gear.filter { it.changed }.map { it.slot })
        assertTrue(model.showsChangedLegend)
    }

    @Test fun `gear is in frame order`() {
        val model = SimResultModel.of("x", RaidFixture.load(gg.foreversixty.app.model.SimResult.serializer(), "sims.result"), null, specs, resources())
        assertEquals(listOf("head", "main_hand"), model.gear.map { it.slot })
        assertFalse(model.showsChangedLegend)
    }

    @Test fun `result load marks seen once`() = runTest {
        val seen = mutableListOf<String>()
        val store = SimResultStore("km3xibdrsiyr", null, FakeSimsService(), MemoryRaidCache(), specs, resources(), { now }) { seen += it }
        store.load()
        store.load()
        assertEquals(listOf("km3xibdrsiyr"), seen)
        assertEquals("Tuesday", (store.phase as SimResultPhase.Loaded).model.title)
    }

    @Test fun `result offline draws the cache else fails`() = runTest {
        val cache = MemoryRaidCache()
        SimResultStore("a", null, FakeSimsService(), cache, specs, resources(), { now }) {}.load()
        val offlineService = FakeSimsService().apply {
            resultAnswer = { fail(offline) }
            listAnswer = { _, _ -> fail(offline) }
        }
        val cached = SimResultStore("a", null, offlineService, cache, specs, resources(), { now }) {}
        cached.load()
        assertEquals(now, (cached.phase as SimResultPhase.Loaded).offlineSince)
        val fresh = SimResultStore("b", null, offlineService, cache, specs, resources(), { now }) {}
        fresh.load()
        assertEquals(SimResultPhase.Failed, fresh.phase)
    }
}
