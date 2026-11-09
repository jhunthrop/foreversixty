package gg.foreversixty.app.widgets

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.widgets.logic.RaidCard
import gg.foreversixty.app.widgets.logic.SignupOverlay
import gg.foreversixty.app.widgets.logic.WidgetFigures
import gg.foreversixty.app.widgets.logic.WidgetSnapshot
import java.time.Instant
import java.time.ZoneId
import kotlinx.serialization.json.JsonNull
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The restyled widget's card and upgrade lines (mocks `widgets-guild`, `-noguild`, `-open`). */
@RunWith(AndroidJUnit4::class)
class RaidCardsTest {
    private val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago"))
    private val now = WidgetTestData.boardNow
    private val afterOpening: Instant = Instant.parse("2026-12-10T18:00:00Z")
    private val zones = "Barrow Deeps, Hyjal Summit, Onyxia's Lair"

    private fun card(snapshot: WidgetSnapshot?, at: Instant = now, overlay: SignupOverlay? = null): RaidCard? =
        figures.raid(snapshot, at, overlay).card

    private fun noRaid() = WidgetTestData.snapshot().let { it.copy(widget = it.widget.copy(nextRaid = null)) }

    private fun my(state: String?, spec: String? = "Elemental") = JsonObject(
        mapOf("my" to JsonObject(mapOf("state" to (state?.let(::JsonPrimitive) ?: JsonNull), "spec" to (spec?.let(::JsonPrimitive) ?: JsonNull), "role" to JsonPrimitive("dps")))),
    )

    @Test
    fun `a planned raid shows the guild, the countdown, the zone, the time and the answer`() {
        assertEquals(
            RaidCard(
                key = "OLYMPUS XXVII", big = "2d 2h", name = "Onyxia's Lair",
                line = "Thu 10 Dec, 8:00 pm CST · You are In", lock = "2d 2h · Onyxia's Lair", counts = "24 in, 0 late",
            ),
            card(WidgetTestData.snapshot()),
        )
    }

    @Test
    fun `the line says the answer in words, and a pending press stands until the server answers`() {
        assertEquals("Thu 10 Dec, 8:00 pm CST · Not signed up", card(WidgetTestData.snapshot(patch = my(null)))?.line)
        assertEquals("Thu 10 Dec, 8:00 pm CST · Choose a spec", card(WidgetTestData.snapshot(patch = my(null, spec = null)))?.line)
        assertEquals("Thu 10 Dec, 8:00 pm CST · You are Late", card(WidgetTestData.snapshot(), overlay = SignupOverlay(7, "late"))?.line)
    }

    @Test
    fun `a cancelled raid says Cancelled in the big figure and only the time below`() {
        val cancelled = card(WidgetTestData.snapshot(patch = WidgetTestData.cancelled))
        assertEquals("Cancelled", cancelled?.big)
        assertEquals("Thu 10 Dec, 8:00 pm CST", cancelled?.line)
        assertTrue(cancelled?.cancelled == true)
    }

    @Test
    fun `under an hour the countdown is minutes, and Now once the pull has passed`() {
        assertEquals("21m", card(WidgetTestData.snapshot(), at = Instant.parse("2026-12-11T01:39:00Z"))?.big)
        assertEquals("Now", card(WidgetTestData.snapshot(), at = Instant.parse("2026-12-11T02:10:00Z"))?.big)
    }

    @Test
    fun `a guildless player counts down to the first raids, then sees them open`() {
        val before = card(WidgetTestData.snapshot("widget.get.no-guild"))
        assertEquals(RaidCard("First raids", "6h 21m", "Wed 9 Dec", zones, "6h 21m · Wed 9 Dec"), before)
        val after = card(WidgetTestData.snapshot("widget.get.no-guild"), at = afterOpening)
        assertEquals(RaidCard("Raids open now", "Open", zones, "No guild yet", "Open · $zones"), after)
    }

    @Test
    fun `a guild with nothing planned counts down to the first raids, then says raids are open`() {
        assertEquals(
            RaidCard("OLYMPUS XXVII", "6h 21m", "First raids Wed 9 Dec", "No raid planned yet", "6h 21m · First raids Wed 9 Dec"),
            card(noRaid()),
        )
        assertEquals(RaidCard("Raids open now", "Open", zones, "No raid planned yet", "Open · $zones"), card(noRaid(), at = afterOpening))
    }

    @Test
    fun `signed out, and an old snapshot with no first raids fields, have no card`() {
        assertNull(card(null))
        val old = noRaid().let { it.copy(widget = it.widget.copy(firstRaidsAt = null, firstRaidsZones = null)) }
        assertNull(card(old, at = afterOpening))
    }

    @Test
    fun `the upgrade card shows the next upgrade's gain and where its item comes from`() {
        val upgrades = figures.upgrades(WidgetTestData.snapshot())
        assertEquals("+5.3 DPS", upgrades.gain)
        assertEquals("Knowledge of the Timbermaw, Blacksmithing", upgrades.item)
        assertNull(figures.upgrades(WidgetTestData.snapshot("widget.get.no-guild")).gain)
    }
}
