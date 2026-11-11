package gg.foreversixty.app.flows

import androidx.compose.ui.test.SemanticsNodeInteraction
import androidx.compose.ui.test.assertIsNotSelected
import androidx.compose.ui.test.assertIsNotEnabled
import androidx.compose.ui.test.assertIsSelected
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.test.performScrollTo
import androidx.compose.ui.test.performTouchInput
import androidx.compose.ui.test.swipeDown
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithContentDescription
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.RaidsDestination
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.StatusLine
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import kotlinx.coroutines.CompletableDeferred
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Changing the answer on the Raid tab root (spec 2026-11-08 F1) through the real presenter, screen and
 * SignupController over the fixture-mode service. Each step is a capture under screenshots/flows.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class RaidSignupFlowTest {
    @get:Rule val compose = createComposeRule()

    private val flow = RaidFlow()
    private val shots = FlowCapture(compose, FLOW)

    private fun show() {
        compose.setContent { FlowFrame(ForeverTab.Raid) { RaidFlowScreen(flow) } }
        compose.waitUntil(SETTLE_MILLIS) { flow.next != null }
        compose.waitForIdle()
    }

    private fun answer(state: SignupState): SemanticsNodeInteraction =
        compose.onNode(hasText(resources().getString(state.label)) and SELECTABLE)

    private fun minutes(value: Int): SemanticsNodeInteraction =
        compose.onNode(hasText(resources().getString(R.string.push_late_minutes, value)) and SELECTABLE)

    private fun lateChip(value: Int) = resources().getString(R.string.raid_answer_lateMinutes, value)

    private fun assertOnly(state: SignupState) = SignupState.entries.forEach {
        if (it == state) answer(it).assertIsSelected() else answer(it).assertIsNotSelected()
    }

    private val sent get() = flow.service.requests.map { it.second }
    private val status get() = flow.next?.block(FlowPlayer.KEY)?.status

    @Test
    fun `In, Late, 30 min, Bench and Out from the tab`() {
        show()
        assertOnly(SignupState.In)
        compose.onNodeWithContentDescription(resources().getString(R.string.raid_tab_howLate)).assertDoesNotExist()
        shots.step(1, "in")

        shots.pressed(answer(SignupState.Late), 2, "late-pressed")
        compose.settle(flow.service)
        assertOnly(SignupState.Late)
        assertEquals(listOf("late"), sent.map { it.state })
        assertNull(sent.last().lateMinutes)
        assertEquals(listOf("selection", "success"), flow.haptics.calls)
        assertTrue(status is StatusLine.Saved)
        // Item 2: the head says so under the answers.
        compose.onNodeWithText(text(R.string.common_saved)).assertExists()
        compose.onNodeWithContentDescription(resources().getString(R.string.raid_tab_howLate)).assertExists()
        listOf(15, 30, 60).forEach { minutes(it).assertIsNotSelected() }
        shots.step(3, "late-chips")

        shots.pressed(minutes(30), 4, "late-30-pressed")
        compose.settle(flow.service)
        assertEquals("late", sent.last().state)
        assertEquals(30, sent.last().lateMinutes)
        minutes(30).assertIsSelected()
        minutes(15).assertIsNotSelected()
        // AS5: the head carries no answer chip; the segmented control and the minutes are the state.
        compose.onNodeWithText(lateChip(30)).assertDoesNotExist()
        assertEquals(listOf("selection", "success", "selection", "success"), flow.haptics.calls)
        shots.step(5, "late-30")

        answer(SignupState.Bench).performClick()
        compose.settle(flow.service)
        assertOnly(SignupState.Bench)
        assertEquals("bench", sent.last().state)
        assertNull(sent.last().lateMinutes)
        compose.onNodeWithContentDescription(resources().getString(R.string.raid_tab_howLate)).assertDoesNotExist()
        shots.step(6, "bench")

        answer(SignupState.Out).performClick()
        compose.settle(flow.service)
        assertOnly(SignupState.Out)
        assertEquals(listOf("late", "late", "bench", "out"), sent.map { it.state })
        assertTrue(flow.service.requests.all { it.first == FlowPlayer.KEY })
        assertEquals(List(4) { listOf("selection", "success") }.flatten(), flow.haptics.calls)
        shots.step(7, "out")
    }

    @Test
    fun `a failed write returns to the answer before it, then offline the answers stop`() {
        show()
        val gate = CompletableDeferred<Unit>()
        flow.service.gate = gate
        flow.service.failWith = ApiError.Server(500, "internal", "boom", emptyMap(), "flow")
        answer(SignupState.Late).performClick()
        compose.waitForIdle()
        // In flight: Late drawn at once, the write still out.
        assertOnly(SignupState.Late)
        assertEquals(1, flow.service.inFlight)
        assertTrue(flow.next?.block(FlowPlayer.KEY)?.busy == true)
        compose.onNodeWithText(text(R.string.raids_note_saving)).assertExists()
        shots.step(8, "late-saving")

        gate.complete(Unit)
        compose.settle(flow.service)
        assertOnly(SignupState.In)
        assertEquals(listOf("selection", "error"), flow.haptics.calls)
        assertEquals(resources().getString(R.string.common_failed), status?.text)
        assertTrue(status is StatusLine.Error)
        compose.onNodeWithText(text(R.string.common_failed)).assertExists()
        compose.onNodeWithText(text(R.string.common_tryAgain)).assertExists()
        shots.step(9, "late-failed")

        // Offline: the tab's refresh finds no connection and draws the saved answer under the strip.
        flow.service.gate = null
        flow.service.failWith = null
        flow.service.offline = true
        compose.runOnIdle { flow.reloadNow() }
        compose.waitUntil(SETTLE_MILLIS) { flow.next?.offline == true }
        compose.waitForIdle()
        compose.onNodeWithText(resources().getString(R.string.offline_strip, "").substringBefore(" Showing"), substring = true).assertExists()
        assertTrue(flow.next?.segmentsEnabled(FlowPlayer.KEY) == false)
        assertOnly(SignupState.In)
        // Items 3 and 4: the answers read as disabled, and the failed write's line gave way to the reason.
        SignupState.entries.forEach { answer(it).assertIsNotEnabled() }
        compose.onNodeWithText(text(R.string.raid_offline)).assertExists()
        compose.onNodeWithText(text(R.string.common_failed)).assertDoesNotExist()
        shots.step(10, "offline")

        val before = flow.service.requests.size
        answer(SignupState.Bench).performClick()
        compose.waitForIdle()
        assertOnly(SignupState.In)
        assertEquals(before, flow.service.requests.size)
        assertEquals(listOf("selection", "error"), flow.haptics.calls)
        compose.onNodeWithText(text(R.string.raid_offline)).assertExists()
        shots.step(11, "offline-tap")
    }

    @Test
    fun `Try again repeats the failed write and says Saved`() {
        show()
        flow.service.failWith = ApiError.Server(500, "internal", "boom", emptyMap(), "flow")
        failLate()
        flow.service.failWith = null
        compose.onNodeWithText(text(R.string.common_tryAgain)).performClick()
        compose.settle(flow.service)
        assertEquals(listOf("late", "late"), sent.map { it.state })
        assertOnly(SignupState.Late)
        compose.onNodeWithText(text(R.string.common_saved)).assertExists()
        compose.onNodeWithText(text(R.string.common_tryAgain)).assertDoesNotExist()
    }

    @Test
    @Config(qualifiers = SHORT_PHONE)
    fun `an answer on the raid page shows on the tab after back, the root where it was`() {
        show()
        val scroll = checkNotNull(flow.rootScroll)
        compose.runOnIdle { runBlocking { scroll.scrollTo(ROOT_OFFSET) } }
        compose.waitForIdle()
        val before = scroll.value
        assertTrue(before > 0)

        // The head scrolled off, so the page is pushed as the head's tap pushes it.
        compose.runOnIdle { flow.navigator.push(RaidsDestination.Raid(FlowPlayer.GUILD, FlowPlayer.RAID)) }
        compose.waitUntil(SETTLE_MILLIS) { compose.onAllNodes(hasText(text(SignupState.Bench.label)) and SELECTABLE).fetchSemanticsNodes().isNotEmpty() }
        answer(SignupState.Bench).performScrollTo().performClick()
        compose.settle(flow.service)
        compose.runOnIdle { flow.navigator.pop() }
        compose.waitForIdle()

        // Item 1: the tab's controller took the page's answer; item 12: the root kept its place.
        assertEquals(SignupState.Bench, flow.next?.block(FlowPlayer.KEY)?.shown)
        assertOnly(SignupState.Bench)
        assertEquals(before, checkNotNull(flow.rootScroll).value)
    }

    @Test
    fun `a pull on the root reloads the raids with the light tick`() {
        show()
        val reads = flow.service.raidsReads
        compose.onRoot().performTouchInput { swipeDown(startY = top + PULL_START, endY = bottom) }
        compose.waitUntil(SETTLE_MILLIS) { flow.service.raidsReads > reads }
        compose.waitForIdle()
        assertEquals(listOf("selection"), flow.haptics.calls)
    }

    /** A failed Late from the tab: its error line and Try again under the answers. */
    private fun failLate() {
        answer(SignupState.Late).performClick()
        compose.settle(flow.service)
        assertOnly(SignupState.In)
    }

    private fun text(id: Int): String = resources().getString(id)

    private companion object {
        const val FLOW = "raid-signup"

        /** Far enough down the root that a lost scroll shows. */
        const val ROOT_OFFSET = 200

        /** A phone short enough that the root scrolls. */
        const val SHORT_PHONE = "w412dp-h480dp-normal-long-notround-any-420dpi-keyshidden-nonav"

        /** Under the tab bar's top, inside the root. */
        const val PULL_START = 400f
    }
}
