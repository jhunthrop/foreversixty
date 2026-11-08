package gg.foreversixty.app.raids.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.GuildOfficerApplications
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.resources
import java.time.Instant
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test
import org.junit.runner.RunWith

/** The applicant's words from the officers' list fixture (application 41, Test Applicant, Holy Paladin 60). */
@RunWith(AndroidJUnit4::class)
class ApplicantTest {
    private val page = RaidFixture.load(GuildOfficerApplications.serializer(), "guild.applications")
    private val app = page.rows.first { it.id == 41L }
    private val now = Instant.parse("2026-12-14T19:04:11Z")

    @Test
    fun `the head, the numbers, raid ready and the history`() {
        val c = ApplicantPresenter(resources(), clock(), now).content(app, "Olympus XXVII", accepted = false)
        assertEquals(listOf("Holy Paladin · 60", "Alliance Dwarf", "applied 2 min ago"), c.facts)
        assertEquals(listOf("74", "612", "63"), c.figures.map { it.value })
        assertEquals("HPS", c.figures[1].unit)
        assertEquals("Attuned to Onyxia's Lair", c.ready[0].title)
        assertEquals("Drakefire Amulet", c.ready[0].sub)
        assertEquals("2 of 4 slots · cloak and boots missing", c.ready[1].sub)
        assertEquals("9 pieces from the Holy Paladin, 60 raid list", c.ready[2].sub)
        assertEquals(listOf("Dawnbreak", "Ashen Vow"), c.history.map { it.title })
        assertNull(c.accepted)
    }

    @Test
    fun `after Accept the page hands over the ginvite line, and the officer's row names the applicant`() {
        val presenter = ApplicantPresenter(resources(), clock(), now)
        val c = presenter.content(app, "Olympus XXVII", accepted = true)
        assertEquals("/ginvite Test Applicant", c.accepted?.command)
        assertEquals("Accepted into Olympus XXVII. Invite Test Applicant in game:", c.accepted?.line)
        assertEquals("Test Applicant, Holy Paladin 60, 2 min ago", presenter.line(app))
    }
}
