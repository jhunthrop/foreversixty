package gg.foreversixty.app.widgets

import androidx.datastore.preferences.core.PreferenceDataStoreFactory
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.LiveActivityRegistered
import gg.foreversixty.app.model.LiveActivityRequest
import gg.foreversixty.app.model.Widget
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.FakeRaidService
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.widgets.data.LiveKind
import gg.foreversixty.app.widgets.data.WidgetRefresher
import gg.foreversixty.app.widgets.data.WidgetService
import gg.foreversixty.app.widgets.data.WidgetSignup
import gg.foreversixty.app.widgets.data.WidgetStore
import gg.foreversixty.app.widgets.live.LiveRegistrations
import gg.foreversixty.app.widgets.logic.LivePush
import gg.foreversixty.app.widgets.logic.RaidLiveContent
import gg.foreversixty.app.widgets.logic.SimLive
import gg.foreversixty.app.widgets.logic.SimLiveContent
import gg.foreversixty.app.widgets.logic.SimThrottle
import gg.foreversixty.app.widgets.logic.WidgetFigures
import gg.foreversixty.app.widgets.tile.TileFigures
import java.io.File
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.test.TestScope
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** Lane F14's rules: the store, the refresh, the widget press, the pushes, the live surfaces, the tile. */
@RunWith(AndroidJUnit4::class)
class WidgetLogicTest {
    private val now: Instant = WidgetTestData.boardNow

    private fun TestScope.store(): WidgetStore {
        val file = File.createTempFile("widgets", ".preferences_pb").also { it.delete() }
        return WidgetStore(PreferenceDataStoreFactory.create(scope = backgroundScope) { file })
    }

    private class FakeWidgetService : WidgetService {
        var answer: () -> ApiResult<Widget> = { ok(WidgetTestData.widget("widget.get")) }
        val asked = mutableListOf<String?>()
        val registered = mutableListOf<LiveActivityRequest>()
        val unregistered = mutableListOf<Long>()

        override suspend fun widget(character: String?): ApiResult<Widget> = answer().also { asked += character }
        override suspend fun register(request: LiveActivityRequest): ApiResult<LiveActivityRegistered> {
            registered += request
            return ok(LiveActivityRegistered(41L + registered.size))
        }
        override suspend fun unregister(id: Long): ApiResult<Unit> = ok(Unit).also { unregistered += id }
    }

    private fun fcm(name: String): Map<String, String> =
        RaidFixture.push(name).getValue("message").jsonObject.getValue("data").jsonObject.mapValues { it.value.jsonPrimitive.content }

    // The refresh

    @Test
    fun `a refresh writes one snapshot per character a widget shows`() = runTest {
        val store = store()
        val service = FakeWidgetService()
        store.setCharacter(5, "us/pvp/yellfury")
        var redraws = 0
        WidgetRefresher(service, store, { true }, { "us/pvp/lightningbolltz" }, { now }) { redraws++ }.refresh()
        assertEquals(setOf("us/pvp/lightningbolltz", "us/pvp/yellfury"), service.asked.toSet())
        assertNotNull(store.snapshot())
        assertEquals(now.toString(), store.snapshot("us/pvp/yellfury")?.fetchedAt)
        assertEquals(1, redraws)
    }

    @Test
    fun `offline keeps the old snapshot and a 401 or sign-out clears every surface`() = runTest {
        val store = store()
        val service = FakeWidgetService()
        val refresher = WidgetRefresher(service, store, { true }, { null }, { now })
        refresher.refresh()
        service.answer = { ApiResult.Failure(ApiError.Network(java.io.IOException("offline"))) }
        refresher.refresh()
        assertNotNull("offline keeps what is drawn", store.snapshot())
        service.answer = { ApiResult.Failure(RaidFixture.error("error-unauthorized", 401)) }
        refresher.refresh()
        assertNull(store.snapshot())
        refresher.refresh()
        WidgetRefresher(service, store, { false }, { null }, { now }).refresh()
        assertNull(store.snapshot())
    }

    // The widget press

    @Test
    fun `a press answers in place keeping the note, then the server's snapshot`() = runTest {
        val store = store()
        val raids = FakeRaidService()
        val refresher = WidgetRefresher(FakeWidgetService(), store, { true }, { null }, { now })
        val saved = WidgetSignup(raids, store, refresher, { now }, wait = {}).answer(2, 7, "us/pvp/lightningbolltz", SignupState.Late)
        assertTrue(saved)
        assertEquals("late", raids.signups.single().second.state)
        assertNull(store.overlay())
        assertNotNull(store.snapshot())
    }

    @Test
    fun `a refused press says Not saved for 10 seconds, then the old answer stands`() = runTest {
        val store = store()
        val raids = FakeRaidService().apply { signupAnswer = { ApiResult.Failure(RaidFixture.error("guild.signup.error-raid-locked", 409)) } }
        val refresher = WidgetRefresher(FakeWidgetService(), store, { true }, { null }, { now })
        val waits = mutableListOf<Duration>()
        var during: SignupState? = null
        val signup = WidgetSignup(raids, store, refresher, { now }, wait = {
            waits += it
            during = WidgetFigures(resources()).raid(WidgetTestData.snapshot(), now, store.overlay()).notSaved
        })
        assertFalse(signup.answer(2, 7, "us/pvp/lightningbolltz", SignupState.Out))
        assertEquals(listOf(Duration.ofSeconds(10)), waits)
        assertEquals(SignupState.Out, during)
        assertNull(store.overlay())
    }

    // Pushes

    @Test
    fun `the live-surface FCM fixtures parse`() {
        assertEquals(LivePush.WidgetRefresh, LivePush.parse(fcm("surface.widget-refresh.fcm")))
        assertEquals(LivePush.Raid(LivePush.Event.Start, 2, 7), LivePush.parse(fcm("surface.live-raid-start.fcm")))
        val sim = LivePush.parse(fcm("surface.live-sim-update.fcm")) as LivePush.Sim
        assertEquals(SimLive("k3v9q2m8x1ab", "Elemental · Top Gear", SimLive.State.Running, 40, 64, 120, null), sim.sim)
        assertNull(LivePush.parse(mapOf("kind" to "live.party.start")))
        assertNull(LivePush.parse(mapOf("kind" to "live.raid.start", "raid_id" to "x")))
    }

    // The live surfaces

    @Test
    fun `the raid countdown reads the snapshot's figures`() {
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago"))
        val snapshot = WidgetTestData.snapshot()
        assertNull(RaidLiveContent.of(figures.raid(snapshot, now)))
        val inWindow = figures.raid(snapshot, Instant.parse("2026-12-11T01:21:48Z"))
        val content = requireNotNull(RaidLiveContent.of(inWindow))
        assertEquals("Invites for Onyxia's Lair in 39 minutes", content.title)
        assertEquals("You are In · DPS · Lightningbolltz", content.text)
        assertEquals("until the pull at 8:00 pm", content.subText)
        assertEquals(inWindow.signedUp, content.bigText.lines().last())
        assertEquals(Instant.parse("2026-12-11T02:00:00Z"), content.countdownTo)
    }

    @Test
    fun `the sim notification says progress, the end and a failure`() {
        val res = resources()
        val running = SimLiveContent.of(SimLive("s", "Elemental · Top Gear", SimLive.State.Running, 40, 64, 120, null), res)
        assertEquals("40 of 64 combinations · about 2 min left", running.text)
        assertEquals(40 to 64, running.progress)
        assertTrue(running.ongoing)
        val plain = SimLiveContent.of(SimLive("s", "Elemental · Sim", SimLive.State.Running, 0, null, null, null), res)
        assertEquals("Running", plain.text)
        assertTrue(plain.indeterminate)
        val done = SimLiveContent.of(SimLive("s", "Elemental · Top Gear", SimLive.State.Finished, 64, 64, null, 412.4), res)
        assertEquals("412 DPS · tap for the result", done.text)
        assertEquals("Finished", done.subText)
        assertFalse(done.ongoing)
        assertEquals("Failed", SimLiveContent.of(SimLive("s", "t", SimLive.State.Failed, 0, null, null, null), res).subText)
    }

    @Test
    fun `a running sim updates at most every 15 seconds and its end always shows`() {
        val throttle = SimThrottle()
        val sim = SimLive("s", "t", SimLive.State.Running, 1, 64, null, null)
        assertTrue(throttle.allows(sim, now))
        assertFalse(throttle.allows(sim.copy(done = 2), now.plusSeconds(10)))
        assertTrue(throttle.allows(sim.copy(done = 3), now.plusSeconds(15)))
        assertTrue(throttle.allows(sim.copy(state = SimLive.State.Finished), now.plusSeconds(16)))
    }

    @Test
    fun `registrations need a push token, replace the last of a kind and end with it`() = runTest {
        val store = store()
        val service = FakeWidgetService()
        LiveRegistrations(service, store).start(LiveKind.Raid, "7")
        assertTrue("no FCM token yet: nothing is registered", service.registered.isEmpty())
        val registrations = LiveRegistrations(service, store) { "fcm-token-1234567890" }
        registrations.start(LiveKind.Raid, "7")
        registrations.start(LiveKind.Raid, "7")
        assertEquals(1, service.registered.size)
        assertEquals(LiveActivityRequest("raid", "android", "fcm-token-1234567890", null, "7"), service.registered.single())
        registrations.start(LiveKind.Raid, "8")
        assertEquals(listOf(42L), service.unregistered)
        registrations.end(LiveKind.Raid)
        assertEquals(listOf(42L, 43L), service.unregistered)
        assertTrue(store.registrations().isEmpty())
    }

    @Test
    fun `the tile says the raid and the answer, else No raid planned`() {
        val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago"))
        val tile = TileFigures.of(figures.raid(WidgetTestData.snapshot(), now), resources())
        assertEquals("Next raid", tile.label)
        assertEquals("Thu 8 pm · In", tile.subtitle)
        assertTrue(tile.active)
        val base = WidgetTestData.snapshot()
        val none = TileFigures.of(figures.raid(base.copy(widget = base.widget.copy(nextRaid = null)), now), resources())
        assertEquals("No raid planned", none.subtitle)
        assertFalse(none.active)
    }

    @Test
    fun `the live switches are on by default`() = runTest {
        val store = store()
        assertTrue(store.liveSetting(LiveKind.Raid))
        store.setLiveSetting(LiveKind.Sim, false)
        assertFalse(store.liveSetting(LiveKind.Sim))
        assertTrue(store.liveSetting(LiveKind.Raid))
    }
}
