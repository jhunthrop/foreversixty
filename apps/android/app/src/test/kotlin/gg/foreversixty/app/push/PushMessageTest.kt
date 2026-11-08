package gg.foreversixty.app.push

import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.push.RaidPushCategory
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test

class PushMessageTest {
    private fun message(data: Map<String, String>, title: String? = null, body: String? = null) =
        PushMessage.parse(data, title, body, collapse = null)

    @Test
    fun `data title and body win over the notification block`() {
        val parsed = message(mapOf("kind" to "sim.finished", "title" to "Sim done", "body" to "948.6 DPS"), "old", "old body")
        assertEquals("Sim done", parsed?.title)
        assertEquals("948.6 DPS", parsed?.body)
    }

    @Test
    fun `an older send reads the notification block`() {
        val parsed = message(mapOf("kind" to "raid.posted"), "Onyxia's Lair", "Are you in?")
        assertEquals("Onyxia's Lair", parsed?.title)
        assertEquals("Are you in?", parsed?.body)
    }

    @Test
    fun `a message with no title is not drawn`() {
        assertNull(message(mapOf("kind" to "raid.posted")))
    }

    @Test
    fun `the group field names the channel, else the kind does`() {
        assertEquals(PushGroup.Ladder, PushGroup.of("ladder", "raid.posted"))
        assertEquals(PushGroup.Raids, PushGroup.of(null, "raid.posted"))
        assertEquals(PushGroup.Officer, PushGroup.of(null, "raid.short"))
        assertEquals(PushGroup.Officer, PushGroup.of(null, "roster.pending"))
        assertEquals(PushGroup.Gear, PushGroup.of(null, "loot.awarded"))
        assertEquals(PushGroup.Gear, PushGroup.of(null, "sim.finished"))
        assertEquals(PushGroup.Logs, PushGroup.of(null, "report.ready"))
        assertEquals(PushGroup.Logs, PushGroup.of(null, "raid.summary"))
        assertEquals(PushGroup.Account, PushGroup.of(null, "billing.failed"))
        assertEquals(PushGroup.Account, PushGroup.of("nonsense", null))
    }

    @Test
    fun `a raid event lands on its raid page`() {
        val parsed = message(
            mapOf("kind" to "raid.posted", "title" to "Onyxia's Lair", "guild_id" to "2", "raid_id" to "31", "link" to "https://foreversixty.gg/guild/eu/olympus/raids/31"),
        )
        assertEquals(DeepLink.Raid(2, 31).url(), parsed?.landing())
        assertEquals("raid.posted-31", parsed?.collapse)
    }

    @Test
    fun `the server's collapse id wins`() {
        val parsed = message(mapOf("kind" to "raid.changed", "title" to "Moved", "raid_id" to "7", "collapse" to "raid.posted-7"))
        assertEquals("raid.posted-7", parsed?.collapse)
    }

    @Test
    fun `a sim and a report land on their result`() {
        val sim = message(mapOf("kind" to "sim.finished", "title" to "Sim done", "link" to "https://foreversixty.gg/sim/abc123"))
        assertEquals(DeepLink.Sim("abc123").url(), sim?.landing())
        val report = message(mapOf("kind" to "report.ready", "title" to "Report ready", "link" to "https://foreversixty.gg/reports/r9"))
        assertEquals(DeepLink.Report("r9").url(), report?.landing())
    }

    @Test
    fun `an app link the server sent opens as itself`() {
        val link = DeepLink.Notifications.url()
        assertEquals(link, message(mapOf("kind" to "billing.failed", "title" to "Payment failed", "link" to link))?.landing())
    }

    @Test
    fun `a site page with no app screen yet opens Home`() {
        assertNull(message(mapOf("kind" to "bis.changed", "title" to "BiS changed", "link" to "https://foreversixty.gg/bis/warrior-fury/60"))?.landing())
    }

    @Test
    fun `the four answers and Late's minutes`() {
        assertEquals(SignupState.Bench, RaidPushCategory.stateFor(RaidPushCategory.ACTION_BENCH))
        assertEquals(SignupState.Late, RaidPushCategory.stateFor("signup_late_30"))
        assertEquals(30, RaidPushCategory.lateMinutes("signup_late_30"))
        assertNull(RaidPushCategory.lateMinutes("signup_late"))
        assertNull(RaidPushCategory.lateMinutes("signup_late_45"))
        assertNull(RaidPushCategory.stateFor("signup_late_45"))
    }
}
