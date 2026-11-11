package gg.foreversixty.app.push

import android.app.NotificationManager
import android.content.Context
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.Shadows.shadowOf

/** Round 5 a1: the drawer's Notifications row reads the groups' real switches. */
@RunWith(AndroidJUnit4::class)
class NotificationsLineTest {
    private val context: Context = ApplicationProvider.getApplicationContext()
    private val manager = context.getSystemService(NotificationManager::class.java)

    @Test fun `the groups that are on read in order, sentence case`() {
        assertEquals("Raids, gear and sims, logs on", NotificationsLine.of(resources(), listOf(PushGroup.Raids, PushGroup.Gear, PushGroup.Logs)))
        assertEquals("All off", NotificationsLine.of(resources(), emptyList()))
    }

    @Test fun `Account is never named and Officer only for an officer`() {
        assertEquals(false, PushGroup.Account in NotificationsLine.shown(officer = true))
        assertEquals(false, PushGroup.Officer in NotificationsLine.shown(officer = false))
        assertEquals(true, PushGroup.Officer in NotificationsLine.shown(officer = true))
    }

    @Test fun `a channel turned off drops out, and app notifications off reads all off`() {
        PushGroup.createChannels(context)
        val logs = manager.getNotificationChannel(PushGroup.Logs.channelId).apply { importance = NotificationManager.IMPORTANCE_NONE }
        manager.createNotificationChannel(logs)
        shadowOf(manager).setNotificationsEnabled(true)
        assertEquals(false, PushGroup.Logs in NotificationsLine.on(context, officer = false))
        assertEquals(true, PushGroup.Raids in NotificationsLine.on(context, officer = false))
        shadowOf(manager).setNotificationsEnabled(false)
        assertEquals(emptyList<PushGroup>(), NotificationsLine.on(context, officer = false))
    }
}
