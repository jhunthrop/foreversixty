package gg.foreversixty.app.gear.flows

import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.BisItemRef
import gg.foreversixty.app.bis.BisList
import gg.foreversixty.app.bis.BisPick
import gg.foreversixty.app.bis.BisRow
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.raids.RaidFixture
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith

/** Round 3 row 10: Gear's counter, Equipped cells, header count and gain cells read one set of figures. */
@RunWith(AndroidJUnit4::class)
class GearCountsTest {
    private val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
    private val home = RaidFixture.load(CharacterHome.serializer(), "character.home.gorvash")

    /** A list whose picks are exactly the home answer's: its upgrades' picks, then what is already best. */
    private val list: BisList = GearFixtures.list.copy(
        rows = home.upgrades.map { BisRow(it.slot, it.slot, pick(it.pick.id, it.pick.name.orEmpty()), null) } +
            home.alreadyBest.map { BisRow(it.slot, it.slot, pick(it.item.id, it.item.name.orEmpty()), null) },
    )

    private fun pick(id: Int, name: String) =
        BisPick(BisItemRef(id, name, null), "", "", isUnmodelled = false, isWorn = false, runnersUp = emptyList(), source = "")

    @Test fun `the counter equals the Equipped cells, and the gain cells equal the header's upgrade count`() {
        val content = GearComposer(res).content(list, "horde", home.upgrades, emptyList(), home)
        val worn = content.slots.count { it.worn }
        assertEquals(home.alreadyBest.size, worn)
        assertEquals(true, content.facts.endsWith("$worn of ${list.rows.size} slots match the list"))
        assertEquals(home.upgradeCount, content.slots.count { it.gains })
        assertEquals(true, content.header?.text?.contains("from ${home.upgradeCount} upgrades"))
    }

    @Test fun `with the home answer no cell is silent - a pick reads its gain, Not modelled or +0,0`() {
        val extra = BisRow("shoulder", "shoulder", pick(999_001, "Officer's Pauldrons"), null)
        val unmodelled = home.upgrades.first().copy(modelled = false)
        val upgrades = listOf(unmodelled) + home.upgrades.drop(1)
        val content = GearComposer(res).content(list.copy(rows = list.rows + extra), "horde", upgrades, emptyList(), home.copy(upgrades = upgrades))
        val shoulder = content.slots.first { it.slot == "shoulder" }
        assertEquals("+0.0", shoulder.gain)
        // The cell draws it as `No gain` (round 8 row 3): measured, nothing gained.
        assertEquals(false, shoulder.gains)
        val first = content.slots.first { it.slot == unmodelled.slot }
        assertEquals(true, first.notModelled)
        assertEquals(null, first.gain)
        assertEquals(true, content.slots.all { it.worn || it.notModelled || it.gain != null })
    }

    @Test fun `cells take the API's slot label, so the two rings and trinkets read apart`() {
        val content = GearComposer(res).content(list, "horde", home.upgrades, emptyList(), home)
        val labels = content.slots.map { it.label }
        assertEquals(true, labels.containsAll(listOf("Finger 1", "Finger 2", "Trinket 1", "Trinket 2")))
    }
}
