package gg.foreversixty.app.gear.flows

import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomeTalents
import java.time.Instant
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test
import org.junit.runner.RunWith

/** Gear's header line (round 2, info-density Gear): the measured total, the sync, the gains not modelled. */
@RunWith(AndroidJUnit4::class)
class GearHeaderTest {
    private val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
    private val now = Instant.parse("2026-11-06T18:00:00Z")
    private val composer = GearComposer(res, now)
    private val home = CharacterHome(
        character = "us/normal/thandrel", state = "ok", dataBuild = "1.60.1.70790", upgradeCount = 13,
        notSimCheckedCount = 0, notModelledCount = 0, talents = HomeTalents("optimized"),
    )

    @Test fun `the total, the count and when the gear synced`() {
        val synced = home.copy(totalGain = 157.6, syncedAt = "2026-11-06T16:00:00Z")
        assertEquals("+158 DPS from 13 upgrades · synced 2 h ago", composer.header(synced, "DPS")?.text)
        assertEquals(false, composer.header(synced, "DPS")?.syncedStale)
    }

    @Test fun `the gains not modelled close the line`() {
        assertEquals("+158 DPS from 13 upgrades · 2 gains not modelled yet", composer.header(home.copy(totalGain = 157.6, notModelledCount = 2), "DPS")?.text)
    }

    @Test fun `past a week the synced part alone takes the warning tone`() {
        val header = requireNotNull(composer.header(home.copy(totalGain = 18.0, syncedAt = "2026-10-29T18:00:00Z", notModelledCount = 1), "DPS"))
        assertEquals("+18.0 DPS from 13 upgrades · synced 8 d ago · 1 gain not modelled yet", header.text)
        assertEquals(listOf("synced 8 d ago"), header.parts().filter { it.second }.map { it.first })
        val week = requireNotNull(composer.header(home.copy(syncedAt = "2026-10-30T18:00:00Z"), "DPS"))
        assertEquals(false, week.syncedStale)
    }

    @Test fun `nothing measured and never synced leaves no line`() {
        assertNull(composer.header(home, "DPS"))
    }
}
