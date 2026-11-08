package gg.foreversixty.app.raids.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.RaidCompTarget
import gg.foreversixty.app.model.RaidNightSummary
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.resources
import java.time.Instant
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The raid page on raid night from me.raid-night.before: Late 30 min, 36 in and 2 late, the Ready check exported. */
@RunWith(AndroidJUnit4::class)
class RaidNightPageTest {
    private val night = RaidFixture.load(RaidNightSummary.serializer(), "me.raid-night.before")
    private val now = Instant.parse("2026-12-16T01:18:00Z")
    private val target = RaidCompTarget(tanks = 2, healers = 8, melee = 15, ranged = 15, source = "guild")

    @Test
    fun `the countdown, your answer and who is coming against the target`() {
        val c = RaidNightPagePresenter(resources(), clock(), now).content(night, target)
        assertEquals(listOf("00", "42"), c.clock.map { it.value })
        assertTrue(c.live)
        assertEquals(SignupState.Late, c.selected)
        assertEquals("You are Late", c.tag)
        assertEquals("Tanks 1 of 2" to "Grunvald · 1 short", c.roles[0])
        assertEquals("Healers 1 of 8" to "Brightwen · 7 short", c.roles[1])
        assertEquals("Damage 1 of 30", c.roles[2].first)
        assertTrue(c.aside.endsWith("36 in, 2 late"))
    }

    @Test
    fun `before you pull lists the export's checks with where to get each miss`() {
        val c = RaidNightPagePresenter(resources(), clock(), now).content(night, null, pressed = SignupState.In)
        assertEquals(SignupState.In, c.selected)
        assertEquals("Drakefire Amulet", c.ready.first().title)
        assertEquals(ReadyTone.Miss, c.ready.first { it.title == "Flask or elixir" }.tone)
        assertEquals("Auction house", c.ready.first { it.title == "Flask or elixir" }.where)
        assertEquals("Tanks", c.roles[0].first)
    }
}
