package gg.foreversixty.app.flows

import androidx.compose.ui.test.SemanticsNodeInteraction
import androidx.compose.ui.test.assertIsNotSelected
import androidx.compose.ui.test.assertIsSelected
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithContentDescription
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.StatusLine
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import kotlinx.coroutines.CompletableDeferred
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** Changing the answer from Home's raid-day card (spec 2026-11-08 F1) through the Raid tab's controller. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class HomeSignupFlowTest {
    @get:Rule val compose = createComposeRule()

    private val flow = HomeFlow()
    private val shots = FlowCapture(compose, FLOW)

    private fun show() {
        compose.setContent { FlowFrame(ForeverTab.Home) { HomeFlowRoot(flow) } }
        compose.waitUntil(SETTLE_MILLIS) { flow.controller != null }
        compose.waitForIdle()
    }

    private fun answer(state: SignupState): SemanticsNodeInteraction =
        compose.onNode(hasText(resources().getString(state.label)) and SELECTABLE)

    private fun minutes(value: Int): SemanticsNodeInteraction =
        compose.onNode(hasText(resources().getString(R.string.push_late_minutes, value)) and SELECTABLE)

    private fun assertOnly(state: SignupState) = SignupState.entries.forEach {
        if (it == state) answer(it).assertIsSelected() else answer(it).assertIsNotSelected()
    }

    private val sent get() = flow.service.requests.map { it.second }
    private val status get() = flow.controller?.block(flow.character.key)?.status
    private val howLate get() = resources().getString(R.string.raid_tab_howLate)
    private fun text(id: Int): String = resources().getString(id)

    @Test
    fun `In, Late and 30 min from the raid card`() {
        show()
        assertOnly(SignupState.In)
        compose.onNodeWithContentDescription(howLate).assertDoesNotExist()
        shots.step(1, "in")

        shots.pressed(answer(SignupState.Late), 2, "late-pressed")
        compose.settle(flow.service)
        assertOnly(SignupState.Late)
        assertEquals("late", sent.single().state)
        assertNull(sent.single().lateMinutes)
        assertEquals(listOf("selection", "success"), flow.haptics.calls)
        assertTrue(status is StatusLine.Saved)
        compose.onNodeWithText(text(R.string.common_saved)).assertExists()
        compose.onNodeWithContentDescription(howLate).assertExists()
        listOf(15, 30, 60).forEach { minutes(it).assertIsNotSelected() }
        shots.step(3, "late-chips")

        shots.pressed(minutes(30), 4, "late-30-pressed")
        compose.settle(flow.service)
        assertEquals(listOf("late", "late"), sent.map { it.state })
        assertEquals(30, sent.last().lateMinutes)
        assertTrue(flow.service.requests.all { it.first == flow.character.key })
        assertEquals(listOf("selection", "success", "selection", "success"), flow.haptics.calls)
        assertOnly(SignupState.Late)
        // Item 5: the chosen minutes selected, the Raid tab's list told; AS5: no line repeats the answer.
        minutes(30).assertIsSelected()
        listOf(15, 60).forEach { minutes(it).assertIsNotSelected() }
        compose.onNodeWithText(resources().getString(R.string.raid_answer_lateMinutes, 30), substring = true).assertDoesNotExist()
        val own = flow.state.nextRaid?.signups?.first { it.own && it.characterKey == flow.character.key }
        assertEquals(30, own?.lateMinutes)
        shots.step(5, "late-30")
    }

    @Test
    fun `a failed write from the card returns to In`() {
        show()
        val gate = CompletableDeferred<Unit>()
        flow.service.gate = gate
        flow.service.failWith = ApiError.Server(500, "internal", "boom", emptyMap(), "flow")
        answer(SignupState.Late).performClick()
        compose.waitForIdle()
        assertOnly(SignupState.Late)
        assertEquals(1, flow.service.inFlight)
        compose.onNodeWithText(text(R.string.raids_note_saving)).assertExists()
        shots.step(6, "late-saving")

        gate.complete(Unit)
        compose.settle(flow.service)
        assertOnly(SignupState.In)
        assertEquals(listOf("selection", "error"), flow.haptics.calls)
        assertEquals(resources().getString(R.string.common_failed), status?.text)
        compose.onNodeWithText(text(R.string.common_failed)).assertExists()
        compose.onNodeWithText(text(R.string.common_tryAgain)).assertExists()
        shots.step(7, "late-failed")
    }

    private companion object {
        const val FLOW = "home-signup"
    }
}
