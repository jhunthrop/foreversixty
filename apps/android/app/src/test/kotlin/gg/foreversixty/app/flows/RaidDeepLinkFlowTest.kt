package gg.foreversixty.app.flows

import android.Manifest
import android.app.Application
import android.app.Notification
import android.app.NotificationManager
import android.app.PendingIntent
import android.content.Intent
import androidx.activity.ComponentActivity
import androidx.compose.ui.test.assertIsSelected
import androidx.compose.ui.semantics.SemanticsProperties
import androidx.compose.ui.test.SemanticsMatcher
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.v2.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.R
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.push.PushMessage
import gg.foreversixty.app.push.PushNotifier
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidsDestination
import gg.foreversixty.app.raids.logic.RaidOpenAction
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.push.RaidActionReceiver
import gg.foreversixty.app.raids.push.RaidNotificationActions
import gg.foreversixty.app.raids.push.RaidPushCategory
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.widgets.live.LiveNotifier
import gg.foreversixty.app.widgets.logic.RaidLiveContent
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.Shadows.shadowOf
import kotlinx.coroutines.CompletableDeferred
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * A raid push or app link opening the Raids stack (spec 2026-11-08 F1, F4): the link through
 * RaidsNavigator.open, the raid page with the press applied (or the spec sheet), back to the tab
 * root; and the notifications' content intents and buttons resolving to those links.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class RaidDeepLinkFlowTest {
    @get:Rule val compose = createAndroidComposeRule<ComponentActivity>()

    private val flow = RaidFlow()
    private val shots = FlowCapture(compose, FLOW)
    private val app: Application = ApplicationProvider.getApplicationContext()

    private fun show() {
        compose.setContent { FlowFrame(ForeverTab.Raid) { RaidFlowScreen(flow) } }
        compose.waitUntil(SETTLE_MILLIS) { flow.next != null }
        compose.waitForIdle()
    }

    /** The link as MainActivity.handleAppLink reads it: the URL parsed, then the Raids navigator. */
    private fun openLink(link: DeepLink) {
        val parsed = requireNotNull(DeepLink.parse(link.url())) { "not an app link: ${link.url()}" }
        assertEquals(link, parsed)
        compose.runOnIdle { assertTrue(flow.navigator.open(parsed)) }
        compose.waitForIdle()
    }

    private fun back() {
        compose.runOnIdle { compose.activity.onBackPressedDispatcher.onBackPressed() }
        compose.waitForIdle()
    }

    /**
     * Waits for the page on top to finish loading: no progress indicator left in the tree. A page pushed
     * or popped back to builds a new model that loads again, so a capture straight after it could catch
     * the spinner; this waits on the state, never on a sleep.
     */
    private fun awaitLoaded() {
        compose.waitUntil(SETTLE_MILLIS) { compose.onAllNodes(LOADING, useUnmergedTree = true).fetchSemanticsNodes().isEmpty() }
        compose.waitForIdle()
    }

    private fun answer(state: SignupState) = compose.onNode(hasText(resources().getString(state.label)) and SELECTABLE)

    @Test
    fun `a Late link opens the raid page with Late saved, back returns to the tab`() {
        show()
        shots.step(1, "tab")

        openLink(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID, state = "late"))
        assertEquals(listOf(RaidsDestination.Raid(FlowPlayer.GUILD, FlowPlayer.RAID, RaidOpenAction(SignupState.Late))), flow.navigator.stack.toList())
        assertTrue(flow.navigator.wantsTab)
        compose.waitUntil(SETTLE_MILLIS) { flow.service.requests.isNotEmpty() }
        compose.settle(flow.service)
        assertEquals("late", flow.service.requests.single().second.state)
        assertEquals(FlowPlayer.KEY, flow.service.requests.single().first)
        assertEquals(listOf("selection", "success"), flow.haptics.calls)
        answer(SignupState.Late).assertIsSelected()
        compose.onNodeWithText(resources().getString(R.string.common_saved)).assertExists()
        shots.step(2, "raid-late-saved")

        back()
        assertTrue(flow.navigator.stack.isEmpty())
        shots.step(3, "back-to-tab")
    }

    @Test
    fun `a spec=ask link opens the spec sheet, the pick saves In with the spec`() {
        flow.service.raidOverride = RaidFixture.raid("guild.raid.ask-spec")
        show()
        openLink(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID, spec = "ask"))
        assertEquals(RaidOpenAction(SignupState.In), (flow.navigator.stack.single() as RaidsDestination.Raid).open)
        compose.onNodeWithText(resources().getString(R.string.spec_title)).assertExists()
        assertTrue(flow.service.requests.isEmpty())
        assertEquals(listOf("selection"), flow.haptics.calls)
        shots.screen(4, "spec-sheet")

        // The capture waits for the saved line itself (it shows the frame the write returns), then
        // stops the clock, so the frame always lands inside the line's 4 s hold instead of racing it.
        val saved = resources().getString(R.string.spec_kept, "Restoration", "Lightningbolltz")
        compose.onNodeWithText("Restoration").performClick()
        compose.waitUntil(SETTLE_MILLIS) { compose.onAllNodes(hasText(saved, substring = true)).fetchSemanticsNodes().isNotEmpty() }
        compose.mainClock.autoAdvance = false
        val (key, request) = flow.service.requests.single()
        assertEquals("us/pvp/lightningbolltz", key)
        assertEquals("in", request.state)
        assertEquals("shaman-restoration", request.spec)
        assertEquals(0, flow.service.inFlight)
        answer(SignupState.In).assertIsSelected()
        shots.screen(5, "spec-picked")
        compose.mainClock.autoAdvance = true
    }

    /**
     * The flake's cause, held open: the pick closes the sheet at once, and the write's answer comes back
     * after that recomposition. It used to run in the sheet's own scope and was cancelled with it.
     */
    @Test
    fun `a spec pick whose write answers after the sheet closes still shows Saved`() {
        flow.service.raidOverride = RaidFixture.raid("guild.raid.ask-spec")
        show()
        openLink(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID, spec = "ask"))
        val gate = CompletableDeferred<Unit>()
        flow.service.gate = gate
        compose.onNodeWithText("Restoration").performClick()
        compose.waitUntil(SETTLE_MILLIS) { flow.service.inFlight == 1 }
        compose.waitForIdle()
        compose.onNodeWithText(resources().getString(R.string.spec_title)).assertDoesNotExist()
        gate.complete(Unit)
        val saved = resources().getString(R.string.spec_kept, "Restoration", "Lightningbolltz")
        compose.waitUntil(SETTLE_MILLIS) { compose.onAllNodes(hasText(saved, substring = true)).fetchSemanticsNodes().isNotEmpty() }
    }

    @Test
    fun `a spec=ask link from a Late press saves Late with the spec`() {
        flow.service.raidOverride = RaidFixture.raid("guild.raid.ask-spec")
        show()
        openLink(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID, spec = "ask", state = "late"))
        assertEquals(RaidOpenAction(SignupState.Late), (flow.navigator.stack.single() as RaidsDestination.Raid).open)
        compose.onNodeWithText(resources().getString(R.string.spec_title)).assertExists()
        assertTrue(flow.service.requests.isEmpty())

        compose.onNodeWithText("Restoration").performClick()
        compose.waitUntil(SETTLE_MILLIS) { flow.service.requests.isNotEmpty() }
        compose.settle(flow.service)
        val request = flow.service.requests.single().second
        assertEquals("late", request.state)
        assertEquals("shaman-restoration", request.spec)
        answer(SignupState.Late).assertIsSelected()
    }

    @Test
    fun `a link answers for the session's character when it is on the raid`() {
        flow.service.raidOverride = RaidFixture.raid("guild.raid.two-characters")
        flow.currentKey = YELLFURY
        show()
        openLink(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID, state = "out"))
        compose.waitUntil(SETTLE_MILLIS) { flow.service.requests.isNotEmpty() }
        compose.settle(flow.service)
        assertEquals(YELLFURY, flow.service.requests.single().first)
        assertEquals("out", flow.service.requests.single().second.state)
    }

    @Test
    fun `a raid night link pushes the raid then the night, back pops each`() {
        show()
        openLink(DeepLink.RaidNight(FlowPlayer.GUILD, FlowPlayer.RAID))
        assertEquals(
            listOf(RaidsDestination.Raid(FlowPlayer.GUILD, FlowPlayer.RAID), RaidsDestination.Night(FlowPlayer.RAID)),
            flow.navigator.stack.toList(),
        )
        awaitLoaded()
        shots.step(6, "night")

        back()
        assertTrue(flow.navigator.stack.single() is RaidsDestination.Raid)
        assertTrue(flow.service.requests.isEmpty())
        awaitLoaded()
        shots.step(7, "night-back-raid")

        back()
        assertTrue(flow.navigator.stack.isEmpty())
    }

    @Test
    fun `the raid push opens its raid page and carries the four answers`() {
        shadowOf(app).grantPermissions(Manifest.permission.POST_NOTIFICATIONS)
        val message = requireNotNull(PushMessage.parse(raidPush(RaidPushCategory.SIGNUP), null, null, collapse = null))
        PushNotifier.show(app, message)

        val posted = postedNotification()
        assertEquals(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID).url(), opened(posted.contentIntent).dataString)
        val actions = posted.actions.map { savedIntent(it.actionIntent).getStringExtra(RaidActionReceiver.EXTRA_ACTION) }
        assertEquals(listOf("signup_in", "signup_late", "signup_bench", "signup_out"), actions)
        assertEquals(listOf("In", "Late", "Bench", "Out"), posted.actions.map { it.title.toString() })
    }

    @Test
    fun `the ask-spec button opens the raid page asking for the spec, with no receiver trampoline`() {
        shadowOf(app).grantPermissions(Manifest.permission.POST_NOTIFICATIONS)
        val message = requireNotNull(PushMessage.parse(raidPush(RaidPushCategory.ASK_SPEC), null, null, collapse = null))
        PushNotifier.show(app, message)

        val button = postedNotification().actions.first()
        // Android 12+: an action that opens the app is an activity PendingIntent itself.
        assertEquals(MainActivity::class.java.name, savedIntent(button.actionIntent).component?.className)
        val link = DeepLink.parse(requireNotNull(opened(button.actionIntent).dataString))
        assertEquals(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID, spec = "ask"), link)
        assertTrue(flow.navigator.open(requireNotNull(link)))
        assertEquals(RaidOpenAction(SignupState.In), (flow.navigator.stack.single() as RaidsDestination.Raid).open)
    }

    @Test
    fun `the ongoing raid notification opens the raid and answers through the receiver`() {
        val content = RaidLiveContent(
            title = "Onyxia's Lair", text = "OLYMPUS XXVII · You are In", line = null, zoneIcon = null,
            countdownTo = FlowTime.now.plusSeconds(LIVE_LEAD_SECONDS), guildId = FlowPlayer.GUILD, raidId = FlowPlayer.RAID,
            description = "Onyxia's Lair", answer = SignupState.In,
        )
        val live = LiveNotifier.raidNotification(app, content, now = FlowTime.now)
        assertEquals(DeepLink.Raid(FlowPlayer.GUILD, FlowPlayer.RAID).url(), opened(live.contentIntent).dataString)

        val data = mapOf("guild_id" to FlowPlayer.GUILD.toString(), "raid_id" to FlowPlayer.RAID.toString())
        val pills = RaidNotificationActions.actions(app, RaidPushCategory.LIVE, data, LiveNotifier.RAID_TAG, 1)
        val intents = pills.map { savedIntent(it.actionIntent) }
        assertEquals(listOf("signup_in", "signup_late", "signup_bench", "signup_out"), intents.map { it.getStringExtra(RaidActionReceiver.EXTRA_ACTION) })
        intents.forEach {
            val bundle = requireNotNull(it.getBundleExtra(RaidActionReceiver.EXTRA_DATA))
            assertEquals(FlowPlayer.RAID.toString(), bundle.getString("raid_id"))
            assertNull(bundle.getString("character_key"))
        }
    }

    private fun raidPush(category: String) = mapOf(
        "kind" to "raid.posted",
        "title" to "Onyxia's Lair",
        "body" to "Saturday 7:00 pm CST. Are you in?",
        "guild_id" to FlowPlayer.GUILD.toString(),
        "raid_id" to FlowPlayer.RAID.toString(),
        "category" to category,
    )

    private fun postedNotification(): Notification {
        val manager = app.getSystemService(NotificationManager::class.java)
        return requireNotNull(shadowOf(manager).allNotifications.singleOrNull()) { "expected one posted notification" }
    }

    private fun savedIntent(pending: PendingIntent?): Intent = shadowOf(requireNotNull(pending) { "no pending intent" }).savedIntent

    private fun opened(pending: PendingIntent?): Intent = savedIntent(pending).also { assertEquals(Intent.ACTION_VIEW, it.action) }

    private companion object {
        const val FLOW = "raid-link"

        /** The fixture account's second verified character on the raid, listed after the main. */
        const val YELLFURY = "us/pvp/yellfury"

        /** The ongoing notification's countdown: 15 min to the invites. */
        const val LIVE_LEAD_SECONDS = 15L * 60

        /** A progress indicator: a page still loading. */
        val LOADING = SemanticsMatcher.keyIsDefined(SemanticsProperties.ProgressBarRangeInfo)
    }
}
