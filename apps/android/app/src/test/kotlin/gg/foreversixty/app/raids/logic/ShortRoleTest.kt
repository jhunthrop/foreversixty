package gg.foreversixty.app.raids.logic

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.RaidCompTarget
import gg.foreversixty.app.raids.RaidFixture
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test
import org.junit.runner.RunWith

/** Round 4 s1: the short role is the first one short in tank, healer, DPS order, as the server's raid.short push. */
@RunWith(AndroidJUnit4::class)
class ShortRoleTest {
    private val raid = RaidFixture.raid("guild.raid.officer")
    private val coming = ComingCounts.of(raid)

    private fun target(tanks: Int = coming.tanks, healers: Int = coming.healers, dps: Int = coming.melee + coming.ranged) =
        RaidCompTarget(tanks = tanks, healers = healers, melee = dps, ranged = 0, source = "guild")

    @Test fun `healers lead a larger damage gap, the role an officer can fill`() {
        val short = ShortRole.of(raid.copy(compTarget = target(healers = coming.healers + 1, dps = coming.melee + coming.ranged + 19)))
        assertEquals(ShortRole.HEALER, short?.wire)
        assertEquals(1, short?.missing)
    }

    @Test fun `a tank short comes before healers and damage`() {
        val short = ShortRole.of(raid.copy(compTarget = target(tanks = coming.tanks + 1, healers = coming.healers + 3)))
        assertEquals(ShortRole.TANK, short?.wire)
    }

    @Test fun `damage leads only when tanks and healers are met`() {
        assertEquals(ShortRole.DPS, ShortRole.of(raid.copy(compTarget = target(dps = coming.melee + coming.ranged + 2)))?.wire)
        assertNull(ShortRole.of(raid.copy(compTarget = target())))
    }
}
