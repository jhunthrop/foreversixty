package gg.foreversixty.app.raids.logic

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.flows.RaidRoleLines
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith

/** Round 3 row 1: one count model, `Damage N` is melee plus ranged, and Home reads what the Raid tab reads. */
@RunWith(AndroidJUnit4::class)
class ComingCountsTest {
    private val raid = RaidFixture.raid()

    @Test fun `damage equals melee plus ranged, counting In and Late`() {
        val counts = ComingCounts.of(raid)
        assertEquals(counts.melee + counts.ranged, counts.damage)
        assertEquals(ComingCounts(tanks = 3, healers = 6, melee = 7, ranged = 8), counts)
    }

    @Test fun `the Raid tab's Damage total equals its own melee and ranged split`() {
        val damage = RaidRoleLines(resources()).of(raid).last()
        val match = requireNotNull(Regex("""(\d+) melee, (\d+) ranged""").find(requireNotNull(damage.split)))
        val (melee, ranged) = match.destructured
        assertEquals(melee.toInt() + ranged.toInt(), damage.have)
    }

    @Test fun `Home's tally equals the Raid tab's for every role`() {
        val home = RaidCardWords(resources(), clock(), RaidTime.now).roleTally(raid).map { it.label to it.have }
        val tab = RaidRoleLines(resources()).of(raid).map { it.tally.label to it.tally.have }
        assertEquals(tab, home)
    }

    @Test fun `bench and out never count, and a damage sign-up without a melee group reads ranged`() {
        val odd = raid.copy(signups = raid.signups.map { if (it.state == "in" && it.group == "melee") it.copy(group = "dps") else it })
        val counts = ComingCounts.of(odd)
        assertEquals(1, counts.melee)
        assertEquals(14, counts.ranged)
    }
}
