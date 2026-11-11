package gg.foreversixty.app.push

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.raids.RaidFixture
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith

/** Round 4 a2: the debug push draws any FCM push fixture by name, with the actions its kind carries. */
@RunWith(AndroidJUnit4::class)
class DebugPushTest {
    private fun parsed(name: String) = requireNotNull(DebugPush.parse(RaidFixture.push(name)))

    @Test fun `raid short carries its Mention action`() {
        val push = parsed("push.raid-short.fcm")
        assertEquals("raid.short", push.kind)
        assertEquals(true, push.data.getValue("actions").contains("\"raid_mention_short\""))
    }

    @Test fun `guild applicant carries Accept then Decline`() {
        val push = parsed("push.guild-applicant.fcm")
        assertEquals("guild.applicant", push.kind)
        val actions = push.data.getValue("actions")
        assertEquals(true, actions.indexOf("application_accept") in 0 until actions.indexOf("application_decline"))
    }

    @Test fun `raid posted still draws as before`() {
        assertEquals("raid.posted", parsed("push.raid-posted.fcm").kind)
    }
}
