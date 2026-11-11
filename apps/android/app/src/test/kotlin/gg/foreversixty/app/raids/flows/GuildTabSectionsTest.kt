package gg.foreversixty.app.raids.flows

import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.guild.GuildFixture
import gg.foreversixty.app.model.LootNights
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.logic.RaidClock
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith
import java.time.Instant
import java.time.ZoneOffset

/** The Guild tab's sections (spec 2026-11-10-guild-tab.md section 2) from guild.home.roster: 38 members, 30 at 60, 5 at 55 to 59. */
@RunWith(AndroidJUnit4::class)
class GuildTabSectionsTest {
    private val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
    private val now = Instant.parse("2026-12-17T12:00:00Z")
    private val presenter = GuildTabPresenter(res, RaidClock(res, ZoneOffset.UTC), now, ZoneOffset.UTC)
    private val viewer = "us/pvp/obnoxious-yell"

    private fun sections(level: Int, loot: LootNights? = null) =
        requireNotNull(presenter.sections(GuildTabInputs(GuildFixture.home(), level = level, characterKey = viewer, lootNights = loot), "OLYMPUS XXVII"))

    @Test fun `under 60 the guild leads, the spread names your level, Close to 60 never repeats Recently on`() {
        val s = sections(36)
        assertTrue(s.leveler)
        assertEquals(GuildSpread(30, 5, 36, "your level", own = true), s.spread)
        assertTrue(s.close.all { (it.trail.toIntOrNull() ?: 0) in 55..59 })
        assertTrue(s.close.map { it.name }.intersect(s.recent.map { it.name }.toSet()).isEmpty())
        assertNull(s.lootNight)
    }

    @Test fun `at 60 the spread names the raiders and the newest loot night lists its drops`() {
        val s = sections(60, RaidFixture.load(LootNights.serializer(), "guild.loot.nights"))
        assertFalse(s.leveler)
        assertEquals("raiders", s.spread?.thirdLabel)
        val drops = requireNotNull(s.lootNight).drops
        assertEquals(listOf("Kraggor", null), drops.map { it.winner })
        assertTrue(s.close.isEmpty())
    }

    @Test fun `the head names the realm in the site's words`() {
        assertEquals("US PvP", GuildTabPresenter.realm("us", "pvp"))
        assertEquals("EU Hardcore", GuildTabPresenter.realm("eu", "hardcore"))
        assertTrue(sections(60).head.line.endsWith("US PvP"))
    }

    @Test fun `seen ages read without a Seen word`() {
        assertTrue(sections(60).recent.none { it.trail.startsWith("Seen") })
    }

    @Test fun `a report's length in hours and minutes`() {
        assertEquals("1 h 32 min", presenter.duration(5_520_000))
        assertEquals("41 min", presenter.duration(2_460_000))
    }
}
