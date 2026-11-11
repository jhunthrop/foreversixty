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
        // The buff's short label from composition (`WF`), else the cover row's name.
        assertEquals("WF 1 of 2", facts.covered(raid.copy(cover = listOf(windfury))))
        val all = raid.copy(cover = emptyList(), composition = raid.composition?.copy(buffs = emptyList()))
        assertNull(facts.covered(all))
        val ok = raid.copy(cover = raid.cover?.map { it.copy(state = "ok") })
        assertNull(facts.covered(ok))
    }

    @Test fun `Covered is one line - missing first, three names, one Missing, then and N more`() {
        val rows = requireNotNull(raid.cover)
        val fortitude = rows.first { it.key == "fortitude" }.copy(state = "missing")
        val windfury = rows[0].copy(state = "short", providersIn = 1, providersLate = 0, need = 2)
        assertEquals("Missing Fort · WF 1 of 2", facts.covered(raid.copy(cover = listOf(windfury, fortitude))))
        val missing = (1..12).map { rows[0].copy(key = "gap$it", name = "Gap $it", state = "missing") }
        assertEquals("Missing Gap 1 · Gap 2 · Gap 3 and 10 more", facts.covered(raid.copy(cover = listOf(windfury) + missing)))
        val fourMissing = missing.take(4)
        assertEquals("Missing Gap 1 · Gap 2 · Gap 3 · Gap 4", facts.covered(raid.copy(cover = fourMissing)))
        val twoMissing = missing.take(2)
        assertEquals("Missing Gap 1 · Gap 2 · WF 1 of 2", facts.covered(raid.copy(cover = listOf(windfury) + twoMissing)))
        assertEquals("Missing Kings · Might · Wisdom and 2 more", facts.covered(raid))
    }

    @Test fun `the Covered row carries each name's spell icon and is not a red wall`() {
        val line = requireNotNull(facts.coverLine(raid))
        assertEquals(listOf("spell_magic_magearmor", "spell_holy_fistofjustice", "spell_holy_sealofwisdom"), line.missing.map { it.icon })
        assertEquals(2, line.more)
        val foot = facts.foot(raid, officer = false, roster = emptyList(), now = RaidTime.now).first { it.cover != null }
        assertEquals(facts.covered(raid), foot.text)
    }

    @Test fun `a Covered name without a spell icon draws none and the short ones carry their count`() {
        val rows = requireNotNull(raid.cover)
        val blank = rows[0].copy(key = "blank", name = "Blank", state = "missing", icon = "")
        val windfury = rows[0].copy(state = "short", providersIn = 1, providersLate = 0, need = 2)
        val names = facts.names(requireNotNull(facts.coverLine(raid.copy(cover = listOf(windfury, blank)))))
        assertEquals(listOf("Blank", "WF 1 of 2"), names.map { it.text })
        assertEquals(listOf(false, true), names.map { it.hasIcon })
        assertEquals(false, CoverName("Kings", null).hasIcon)
    }

    @Test fun `without cover rows Covered falls back to the buffs nobody brings`() {
        val buffs = requireNotNull(raid.composition).buffs.map { it.copy(present = it.key != "fortitude") }
        val line = facts.covered(raid.copy(cover = emptyList(), composition = raid.composition?.copy(buffs = buffs)))
        assertEquals("Missing Fort", line)
    }

    @Test fun `an officer sees who signed In without syncing this week, a member does not`() {
        val home = RaidFixture.load(GuildHome.serializer(), "guild.home")
        val coming = raid.signups.first { it.state == "in" }
        val roster = home.roster.take(1).map { it.copy(characterKey = coming.characterKey, loggedAt = "2026-12-09T22:00:00Z") }
        assertEquals(listOf(coming.name), facts.notSynced(raid, roster, RaidTime.now).map { it.name })
        val foot = facts.foot(raid, officer = true, roster, RaidTime.now).first { it.onlyKeys.isNotEmpty() }
        assertEquals("Not synced this week: ${coming.name}", foot.text)
        assertEquals(setOf(coming.characterKey), foot.onlyKeys)
        assertTrue(facts.foot(raid, officer = false, roster, RaidTime.now).none { it.text.contains("synced") })
    }

    @Test fun `not synced reads the addon's logged_at against seven days, not logged_recently`() {
        val home = RaidFixture.load(GuildHome.serializer(), "guild.home")
        val coming = raid.signups.first { it.state == "in" }
        val row = home.roster.first().copy(characterKey = coming.characterKey)
        // A sync three days ago is this week, though logged_recently (24 h) is false.
        val recent = row.copy(loggedAt = "2026-12-14T22:30:00Z", loggedRecently = false)
        assertTrue(facts.notSynced(raid, listOf(recent), RaidTime.now).isEmpty())
        // No sync time at all reads as not synced.
        val never = row.copy(loggedAt = "")
        assertEquals(listOf(coming.name), facts.notSynced(raid, listOf(never), RaidTime.now).map { it.name })
    }

    @Test fun `the not-synced line names four in full, five as three and 2 more, never and 1 more`() {
        val coming = raid.signups.filter { it.state == "in" }.take(5)
        val four = coming.take(4)
        assertEquals("Not synced this week: ${four.joinToString(", ") { it.name }}", facts.notSyncedLine(four))
        assertEquals("Not synced this week: ${coming[0].name}, ${coming[1].name}, ${coming[2].name} and 2 more", facts.notSyncedLine(coming))
    }

    @Test fun `the head reads the invite time and the size`() {
        val words = RaidCardWords(resources(), clock(), RaidTime.now)
        val card = words.card(raid, "Olympus XXVII", "us/pvp/gorvash", null, null)
        assertEquals("Invites 6:45 pm CST · 40-man", card.facts)
        assertNull(words.card(raid.copy(inviteAt = null), "Olympus XXVII", "us/pvp/gorvash", null, null).facts)
    }
}
