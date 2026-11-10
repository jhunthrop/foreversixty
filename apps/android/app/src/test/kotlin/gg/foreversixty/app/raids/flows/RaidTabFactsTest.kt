package gg.foreversixty.app.raids.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** Round 2 on the Raid tab: the answers' counts, the Covered line (G4), readiness (G1), the head facts. */
@RunWith(AndroidJUnit4::class)
class RaidTabFactsTest {
    private val facts = RaidTabFacts(resources())
    private val raid = RaidFixture.raid()

    @Test fun `the counts line names bench, out and not answered, and drops the empty ones`() {
        val counts = raid.counts.copy(bench = 1, signedOut = 4)
        val line = facts.counts(raid.copy(counts = counts))
        val parts = listOfNotNull("Bench 1", "Out 4", raid.notSigned.size.takeIf { it > 0 }?.let { "$it not answered" })
        assertEquals(parts.joinToString(" · "), line)
        assertNull(facts.counts(raid.copy(counts = raid.counts.copy(bench = 0, signedOut = 0), notSigned = emptyList())))
    }

    @Test fun `Covered names what is missing or short and is empty when all is covered`() {
        val windfury = requireNotNull(raid.cover).first().copy(state = "short", providersIn = 1, providersLate = 0, need = 2)
        val short = raid.copy(cover = listOf(windfury))
        assertTrue(requireNotNull(facts.covered(short)).startsWith("${windfury.name} 1 of 2"))
        assertTrue(requireNotNull(facts.covered(raid)).contains("Missing Blessings"))
        val all = raid.copy(cover = emptyList(), composition = raid.composition?.copy(buffs = emptyList()))
        assertNull(facts.covered(all))
    }

    @Test fun `an officer sees who signed In without syncing this week, a member does not`() {
        val home = RaidFixture.load(GuildHome.serializer(), "guild.home")
        val coming = raid.signups.first { it.state == "in" }
        val roster = home.roster.take(1).map { it.copy(characterKey = coming.characterKey, loggedRecently = false) }
        assertEquals(1, facts.notSynced(raid, roster))
        assertTrue(facts.foot(raid, officer = true, roster).any { it.warn && it.text == "1 signed In has not synced this week" })
        assertTrue(facts.foot(raid, officer = false, roster).none { it.text.contains("synced") })
    }

    @Test fun `the head reads the invite time and the size`() {
        val words = RaidCardWords(resources(), clock(), RaidTime.now)
        val card = words.card(raid, "Olympus XXVII", "us/pvp/gorvash", null, null)
        assertEquals("Invites 6:45 pm CST · 40", card.facts)
        assertNull(words.card(raid.copy(inviteAt = null), "Olympus XXVII", "us/pvp/gorvash", null, null).facts)
    }
}
