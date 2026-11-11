package gg.foreversixty.app.push

import android.app.NotificationManager
import android.content.Context
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.NotificationSettings
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.Shadows.shadowOf

/** The drawer's Notifications row: the account's per-group choices and the channels' switches (round 7 row 4). */
@RunWith(AndroidJUnit4::class)
class NotificationsLineTest {
    private val context: Context = ApplicationProvider.getApplicationContext()
    private val manager = context.getSystemService(NotificationManager::class.java)
    private val settings = RaidFixture.load(NotificationSettings.serializer(), "app.notifications")

    @Test fun `the row says what is missed - all on, all off, or the groups that are off`() {
        val shown = NotificationsLine.shown(officer = false)
        assertEquals("All on", NotificationsLine.of(resources(), shown, shown))
        assertEquals("All off", NotificationsLine.of(resources(), shown, emptyList()))
        assertEquals("Logs, tiers off", NotificationsLine.of(resources(), shown, shown - setOf(PushGroup.Logs, PushGroup.Tiers)))
        assertEquals("Gear and sims off", NotificationsLine.of(resources(), shown, shown - PushGroup.Gear))
    }

    @Test fun `the account's groups are the ones it gets events of, on when every event is on for push`() {
        val account = NotificationsLine.account(settings)
        assertEquals(listOf(PushGroup.Raids, PushGroup.Gear, PushGroup.Ladder, PushGroup.Tiers, PushGroup.Officer), account.keys.toList())
        assertEquals(listOf(PushGroup.Raids, PushGroup.Officer), account.filterValues { it }.keys.toList())
    }

    @Test fun `a group is on only when the account and its channel are`() {
        val account = NotificationsLine.account(settings)
        val shown = account.keys.toList()
        // The fixture with every channel on: the account's off groups.
        assertEquals("Gear and sims, guild ladder, tiers off", NotificationsLine.line(resources(), shown, account, shown))
        // The account's on group whose channel is off is off too.
        assertEquals("Raids, gear and sims, guild ladder, tiers off", NotificationsLine.line(resources(), shown, account, shown - PushGroup.Raids))
        // Without the account's answer the channels decide.
        assertEquals("All on", NotificationsLine.line(resources(), shown, null, shown))
    }

    @Test fun `Account is never named and Officer only for an officer`() {
        assertEquals(false, PushGroup.Account in NotificationsLine.shown(officer = true))
        assertEquals(false, PushGroup.Officer in NotificationsLine.shown(officer = false))
        assertEquals(true, PushGroup.Officer in NotificationsLine.shown(officer = true))
    }

    @Test fun `a channel turned off drops out, and app notifications off reads all off`() {
        PushGroup.createChannels(context)
        val shown = NotificationsLine.shown(officer = false)
        val logs = manager.getNotificationChannel(PushGroup.Logs.channelId).apply { importance = NotificationManager.IMPORTANCE_NONE }
        manager.createNotificationChannel(logs)
        shadowOf(manager).setNotificationsEnabled(true)
        assertEquals(false, PushGroup.Logs in NotificationsLine.channelsOn(context, shown))
        assertEquals(true, PushGroup.Raids in NotificationsLine.channelsOn(context, shown))
        shadowOf(manager).setNotificationsEnabled(false)
        assertEquals(emptyList<PushGroup>(), NotificationsLine.channelsOn(context, shown))
    }
}
