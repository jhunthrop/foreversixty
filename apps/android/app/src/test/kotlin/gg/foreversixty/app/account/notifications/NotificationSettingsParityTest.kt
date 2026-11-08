package gg.foreversixty.app.account.notifications

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.model.NotificationChannels
import gg.foreversixty.app.model.NotificationEvent
import gg.foreversixty.app.model.NotificationSettings
import gg.foreversixty.app.model.NotificationsPatch
import gg.foreversixty.app.model.QuietHours
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** Notifications (design/mocks/ws_app_flows.py `notification_settings`), pushed under Home. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class NotificationSettingsParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun settings() = ParityCheck.state(compose, "notification-settings") { Pushed() }

    @Test fun settingsAt2() = ParityCheck.state(compose, "notification-settings", fontScale = 2f) { Pushed() }

    @Test fun aMemberSeesFiveGroupsInOrder() {
        val summary = NotificationGroupsSummary.of(NotificationFixtures.member)
        assertEquals(listOf(PushGroup.Raids, PushGroup.Gear, PushGroup.Logs, PushGroup.Ladder, PushGroup.Tiers), summary.rows.map { it.group })
        assertEquals(listOf(true, true, true, false, false), summary.rows.map { it.on })
        assertEquals("3:00 pm", NotificationGroupsSummary.clock(summary.reminder))
        assertFalse(summary.officerDms)
    }

    @Test fun aGroupIsOnOnlyWhenEveryKindPushes() {
        val mixed = NotificationFixtures.member.copy(events = NotificationFixtures.member.events + event("raid.changed", "raids", false))
        assertFalse(NotificationGroupsSummary.of(mixed).rows.first().on)
    }

    @Test fun aGroupSwitchPatchesEveryKindInIt() {
        val patch = NotificationGroupsSummary.patch(PushGroup.Raids, false, NotificationFixtures.member)
        assertEquals(listOf("raid.posted", "raid.starting"), patch.events?.map { it.kind })
        assertTrue(patch.events.orEmpty().all { it.channel == "push" && !it.enabled })
    }

    @Test fun aRefusedSavePutsTheChoiceBack() = runTest {
        val service = FakeService(save = ApiResult.Failure(ApiError.Status(500)))
        val feature = NotificationSettingsFeature(service)
        feature.load()
        feature.setOfficerDms(true)
        assertEquals(NotificationsPatch(officerDms = true), service.saved.single())
        assertFalse(NotificationGroupsSummary.of((feature.state as NotificationSettingsState.Loaded).settings).officerDms)
        assertTrue(feature.saveFailed)
    }

    @Test fun aSavedReminderIsTheServersAnswer() = runTest {
        val answer = NotificationFixtures.member.copy(raidReminder = "18:00")
        val feature = NotificationSettingsFeature(FakeService(save = ApiResult.Success(answer, "r")))
        feature.load()
        feature.setReminder("18:00")
        assertEquals("18:00", (feature.state as NotificationSettingsState.Loaded).settings.raidReminder)
        assertFalse(feature.saveFailed)
    }

    @Composable
    private fun Pushed() {
        ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}, badges = TabBadges()) {
            Column {
                FlowBackBar("Home", {})
                NotificationSettingsScreen(NotificationSettingsState.Loaded(NotificationFixtures.member), false, NotificationSettingsActions())
            }
        }
    }

    private class FakeService(private val save: ApiResult<NotificationSettings>) : NotificationSettingsService {
        val saved = mutableListOf<NotificationsPatch>()
        override suspend fun load(): ApiResult<NotificationSettings> = ApiResult.Success(NotificationFixtures.member, "r")
        override suspend fun save(patch: NotificationsPatch): ApiResult<NotificationSettings> = save.also { saved += patch }
    }
}

private fun event(kind: String, group: String, push: Boolean) =
    NotificationEvent(kind, group, "player", officer = false, quietExempt = false, timeSensitive = false, channels = NotificationChannels(push, push))

/** Thandrel, a guild member (no Officer group): Raids, Gear and sims, Logs on; Guild ladder, Tiers off. */
object NotificationFixtures {
    val member = NotificationSettings(
        timeZone = "America/Chicago",
        quietHours = QuietHours(enabled = false, start = "22:00", end = "08:00"),
        officerDms = false,
        raidReminder = "15:00",
        events = listOf(
            event("raid.posted", "raids", true),
            event("raid.starting", "raids", true),
            event("loot.awarded", "gear", true),
            event("report.ready", "logs", true),
            event("guildsim.passed", "ladder", false),
            event("tier.moved", "tiers", false),
            event("billing.failed", "account", true),
        ),
    )
}
