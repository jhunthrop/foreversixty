package gg.foreversixty.app.raids

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.R
import gg.foreversixty.app.home.flows.AnswerControl
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.menu.MenuRefresh
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.SignupAnswer
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.push.PushMessage
import gg.foreversixty.app.push.PushLanding
import gg.foreversixty.app.raids.flows.NextRaidSignup
import gg.foreversixty.app.raids.logic.NotificationPermission
import gg.foreversixty.app.raids.logic.RaidOpenAction
import gg.foreversixty.app.raids.logic.RaidPageModel
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.StatusLine
import gg.foreversixty.app.raids.push.RaidPushAnswerer
import gg.foreversixty.app.raids.push.RaidPushOutcome
import gg.foreversixty.app.raids.push.RaidPushPayload
import java.io.IOException
import kotlin.time.Duration.Companion.seconds
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/**
 * The raid cards' sign-up after the 2026-11-10 flow walk: the status line and Try again (items 2 and
 * 4), the shared next-raid controller (1, 5), Late's minutes on the card (5, 6) and the links (7 to 9).
 */
@RunWith(AndroidJUnit4::class)
class SignupCardTest {
    private val key = "us/pvp/gorvash"
    private val offline = ApiError.Network(IOException("offline"))
    private val boom = ApiError.Status(500)

    private fun deps(service: FakeRaidService, haptics: RecordingHaptics = RecordingHaptics()) = SignupController.Dependencies(
        service = service,
        res = resources(),
        haptics = haptics,
        permission = NotificationPermission { false },
        askStore = MemoryAskStore(),
        now = { RaidTime.now },
        scope = null,
        hold = 4.seconds,
    )

    /** The fixture's answer to a sign-up, with the request's state and minutes in the player's own row. */
    private fun echo(view: RaidView, request: SignupRequest): RaidView = view.copy(
        viewer = view.viewer.copy(characters = view.viewer.characters.map { if (it.characterKey == key) it.copy(state = request.state) else it }),
        signups = view.signups.map { if (it.characterKey == key) it.copy(state = request.state, lateMinutes = request.lateMinutes) else it },
    )

    private fun echoing(service: FakeRaidService) {
        val canned = RaidFixture.load(SignupAnswer.serializer(), "guild.signup")
        service.signupAnswer = { request -> ok(canned.copy(raid = echo(RaidFixture.raid(), request), moved = null)) }
    }

    // Status line and Try again

    @Test
    fun `a write out reads Saving, a failure offers Try again, and Try again sends the same write`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { fail(boom) }
        val controller = makeController(service = service)
        controller.pressLate(key, 30)
        assertEquals(StatusLine.Error(resources().getString(R.string.common_failed)), controller.answerStatus(key))
        assertTrue(controller.canRetry(key))

        echoing(service)
        controller.retry(key)
        assertEquals(listOf(30, 30), service.signups.map { it.second.lateMinutes })
        assertEquals(listOf("late", "late"), service.signups.map { it.second.state })
        assertTrue(controller.block(key).status is StatusLine.Saved)
        assertFalse(controller.canRetry(key))
    }

    @Test
    fun `while a write is out the cards read Saving`() = runTest {
        val service = FakeRaidService()
        val controller = makeController(service = service)
        val during = mutableListOf<StatusLine?>()
        val answer = service.signupAnswer
        service.signupAnswer = { request -> answer(request).also { during += controller.answerStatus(key) } }
        assertNull(controller.answerStatus(key))
        controller.press(key, SignupState.Bench)
        assertEquals(listOf(StatusLine.Info(resources().getString(R.string.raids_note_saving))), during)
    }

    @Test
    fun `an error line gives way to the offline sentence on the next press`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { fail(offline) }
        val controller = makeController(service = service)
        controller.press(key, SignupState.Out)
        assertTrue(controller.restingStatus(key) is StatusLine.Error)

        controller.press(key, SignupState.Bench)
        assertEquals(StatusLine.Info(resources().getString(R.string.raid_offline)), controller.restingStatus(key))
        assertEquals(1, service.signups.size)
    }

    @Test
    fun `a newer view clears a failed write's line`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { fail(boom) }
        val controller = makeController(service = service)
        controller.press(key, SignupState.Late)
        assertTrue(controller.block(key).status is StatusLine.Error)
        controller.refresh()
        assertNull(controller.block(key).status)
        assertFalse(controller.canRetry(key))
    }

    // The card

    @Test
    fun `the card carries the controller's state, its line, and dims offline`() = runTest {
        val service = FakeRaidService()
        service.signupAnswer = { fail(offline) }
        val controller = makeController(service = service)
        controller.press(key, SignupState.Out)
        val card = RaidCardWords(resources(), clock(), RaidTime.now).card(controller.raid, "Olympus XXVII", key, controller, null)
        assertEquals(SignupState.In, card.answer)
        assertFalse(card.control.enabled)
        assertTrue(card.control.status is StatusLine.Error)
        assertTrue(card.control.canRetry)
    }

    @Test
    fun `Late with its minutes reads Late 30 min on the raid-day card and the line, the chip selected`() = runTest {
        val service = FakeRaidService()
        echoing(service)
        val controller = makeController(service = service)
        controller.pressLate(key, 30)
        val card = RaidCardWords(resources(), clock(), RaidTime.now).card(RaidFixture.raid(), "Olympus XXVII", key, controller, null)
        val late = resources().getString(R.string.raid_answer_lateMinutes, 30)
        assertEquals(SignupState.Late, card.answer)
        assertEquals(30, card.lateMinutes)
        assertTrue(card.lateOpen)
        // AS5: the segmented control and its minutes are the answer; with one character no line repeats it.
        assertEquals(null, card.me)
        assertEquals(false, card.sub.orEmpty().contains(late))
        assertEquals(late, RaidCardWords(resources(), clock(), RaidTime.now).answerWord(SignupState.Late, card.lateMinutes))
    }

    @Test
    fun `without a controller the card is the raid's own sign-up, answers enabled`() {
        val card = RaidCardWords(resources(), clock(), RaidTime.now).card(RaidFixture.raid(), "Olympus XXVII", key, null, null)
        assertEquals(SignupState.In, card.answer)
        assertEquals(AnswerControl(), card.control)
    }

    // The shared next-raid controller

    private suspend fun loadedSignup(service: FakeRaidService): Pair<RaidsTabModel, NextRaidSignup> {
        val tab = RaidsTabModel(service, MemoryRaidCache(), resources(), { RaidTime.now })
        tab.currentKey = key
        tab.load()
        val signup = NextRaidSignup(tab, { RaidTime.now })
        signup.sync(deps(service))
        return tab to signup
    }

    @Test
    fun `the raid page's answer reaches the tab's controller and list`() = runTest {
        val service = FakeRaidService()
        val (tab, signup) = loadedSignup(service)
        val controller = requireNotNull(signup.controller)
        val bench = echo(controller.raid, SignupRequest(state = "bench"))

        signup.replaced(bench)
        assertEquals(SignupState.Bench, controller.block(key).shown)
        val listed = (tab.state as RaidsTabState.Loaded).content.planned.first { it.id == bench.id }
        assertEquals("bench", listed.signups.first { it.characterKey == key }.state)
    }

    @Test
    fun `a write on the card reaches the tab's list`() = runTest {
        val service = FakeRaidService()
        echoing(service)
        val (tab, signup) = loadedSignup(service)
        requireNotNull(signup.controller).pressLate(key, 30)
        val raidId = requireNotNull(signup.controller).raid.id
        val listed = (tab.state as RaidsTabState.Loaded).content.planned.first { it.id == raidId }
        assertEquals(30, listed.signups.first { it.characterKey == key }.lateMinutes)
    }

    // Links and pushes

    @Test
    fun `a Late that needs a spec opens spec=ask with state=late, and the link presses Late`() = runTest {
        val service = FakeRaidService()
        service.raidAnswer = { ok(RaidFixture.raid("guild.raid.ask-spec")) }
        val data = mapOf("guild_id" to "2", "raid_id" to "31", "character_key" to "us/pvp/lightningbolltz")
        val payload = requireNotNull(RaidPushPayload.parse(data))
        val outcome = RaidPushAnswerer(service, resources()).answer(payload, SignupState.Late)
        val link = DeepLink.Raid(2, 31, spec = "ask", state = "late")
        assertEquals(RaidPushOutcome.Open(link), outcome)

        val navigator = RaidsNavigator()
        navigator.open(requireNotNull(DeepLink.parse(link.url())))
        assertEquals(listOf(RaidsDestination.Raid(2, 31, open = RaidOpenAction(SignupState.Late))), navigator.stack.toList())
    }

    @Test
    fun `a link's press answers for the session's character when it is on the raid`() = runTest {
        val service = FakeRaidService()
        service.raidAnswer = { ok(RaidFixture.raid("guild.raid.two-characters")) }
        val yellfury = "us/pvp/yellfury"
        val model = RaidPageModel(2, 31, deps(service), MemoryRaidCache(), RaidOpenAction(SignupState.Bench), currentKey = yellfury)
        model.load()
        val controller = (model.state as RaidPageModel.State.Loaded).controller
        val pressed = controller.specAsk?.characterKey ?: service.signups.single().first
        assertEquals(yellfury, pressed)
    }

    @Test
    fun `a raid push's tap keeps the ids and carries the link's state and spec`() {
        val data = mapOf(
            "kind" to "raid.posted", "title" to "Onyxia's Lair", "guild_id" to "2", "raid_id" to "31",
            "link" to DeepLink.Raid(9, 99, spec = "ask", state = "bench").url(),
        )
        val message = requireNotNull(PushMessage.parse(data, null, null, null))
        assertEquals(DeepLink.Raid(2, 31, spec = "ask", state = "bench"), PushLanding.of(message))
        val plain = requireNotNull(PushMessage.parse(data - "link", null, null, null))
        assertEquals(DeepLink.Raid(2, 31), PushLanding.of(plain))
    }

    @Test
    fun `the menu's Refresh shows Refreshing, then rests or fails`() = runTest {
        val shown = mutableListOf<MenuRefresh>()
        MenuRefresh.run({ shown += it }) { true }
        MenuRefresh.run({ shown += it }) { false }
        assertEquals(listOf(MenuRefresh.Refreshing, MenuRefresh.Idle, MenuRefresh.Refreshing, MenuRefresh.Failed), shown)
    }
}
