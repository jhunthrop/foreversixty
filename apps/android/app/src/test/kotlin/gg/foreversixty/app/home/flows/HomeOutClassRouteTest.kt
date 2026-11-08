package gg.foreversixty.app.home.flows

import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.bis.SpecInfo
import org.junit.Assert.assertEquals
import org.junit.Test

/** Signed-out Home's class tap (spec R3): the class's BiS at 60 raid, the spec page pushed under Gear. */
class HomeOutClassRouteTest {
    private val specs = SpecCatalog(
        listOf(
            SpecInfo("warrior-arms", "warrior", "Arms", "dps", "ability_warrior_savageblow"),
            SpecInfo("warrior-fury", "warrior", "Fury", "dps", "ability_warrior_innerrage"),
            SpecInfo("warrior-protection", "warrior", "Protection", "tank", "ability_warrior_defensivestance"),
        ),
    )

    private val events = mutableListOf<String>()
    private val requested = mutableListOf<BisSelection>()
    private val route = HomeOutClassRoute(
        specs,
        faction = { "horde" },
        request = { requested += it },
        pushSpecPage = { events += "push" },
        openGear = { events += "gear" },
    )

    private val raid60 = BisRange(BisRange.LAST_BAND, BisRange.RAID)

    @Test
    fun `a class opens its first spec at 60 raid on the pushed spec page under Gear`() {
        route.open("warrior")
        assertEquals(listOf(BisSelection("warrior", "warrior-arms", raid60, "horde")), requested)
        assertEquals(listOf("push", "gear"), events)
    }

    @Test
    fun `the spec the row names wins`() {
        route.open("warrior", "Fury")
        assertEquals("warrior-fury", requested.single().spec)
    }

    @Test
    fun `an unknown spec name falls back to the first spec`() {
        route.open("warrior", "Gladiator")
        assertEquals("warrior-arms", requested.single().spec)
    }

    @Test
    fun `a class the catalogue does not know lands on Gear without a push`() {
        route.open("deathknight")
        assertEquals(emptyList<BisSelection>(), requested)
        assertEquals(listOf("gear"), events)
    }
}
