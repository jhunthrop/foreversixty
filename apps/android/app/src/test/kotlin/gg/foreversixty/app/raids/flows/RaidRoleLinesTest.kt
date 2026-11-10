package gg.foreversixty.app.raids.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.RaidCompTarget
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith

/** Who is coming (spec 4.2): each role with the names of those In or Late, yours first, the short role's gap. */
@RunWith(AndroidJUnit4::class)
class RaidRoleLinesTest {
    private val lines = RaidRoleLines(resources())
    private val raid = RaidFixture.raid()

    @Test fun `each role counts those In or Late, never Bench or Out`() {
        val roles = lines.of(raid)
        assertEquals(listOf("Tanks" to 3, "Healers" to 6, "Damage" to 15), roles.map { it.role to it.have })
    }

    @Test fun `five names show in full, more read three and the rest`() {
        val roles = lines.of(raid)
        assertEquals("Kraggor, Ironhide, Bramblehoof", roles[0].names)
        assertEquals("Mendara, Gravemend, Rainspeaker and 3 more", roles[1].names)
    }

    @Test fun `your own character leads its role`() {
        assertEquals(true, lines.of(raid)[2].names.startsWith("Gorvash, "))
    }

    @Test fun `a Late answer carries its minutes`() {
        val late = raid.copy(signups = raid.signups.map { if (it.name == "Kraggor") it.copy(state = "late", lateMinutes = 15) else it })
        assertEquals("Kraggor (Late\u00a0·\u00a015\u00a0min), Ironhide, Bramblehoof", lines.of(late)[0].names)
    }

    @Test fun `the comp target sets each count and names the short role's gap`() {
        val target = raid.copy(compTarget = RaidCompTarget(tanks = 2, healers = 8, melee = 6, ranged = 8, source = "guild"))
        val roles = lines.of(target)
        assertEquals(listOf(2, 8, 14), roles.map { it.want })
        assertEquals("Mendara, Gravemend, Rainspeaker and 3 more · 2 short", roles[1].names)
        assertEquals(false, roles[0].names.contains("short"))
    }
}
