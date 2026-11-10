package gg.foreversixty.app.activity

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.ActivityPage
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.ReadyCheck
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.flows.RaidTabPresenter
import gg.foreversixty.app.raids.flows.ReadyRows
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/**
 * Round 2 (iOS ActivityRowTests): one Activity row on Activity and Home's New for you (the feed's
 * styled body, the other player's crest, the trailing figure), and the Raid tab's ready icons and
 * Loot history line.
 */
@RunWith(AndroidJUnit4::class)
class ActivityRowTest {
    private val rows = RaidFixture.load(ActivityPage.serializer(), "activity.list").rows

    private fun row(kind: String): ActivityRow = requireNotNull(rows.firstOrNull { it.kind == kind }) { "no $kind row in activity.list" }

    @Test fun `loot names the winner in class colour and the item in its quality, and its text is the body`() {
        val loot = row("loot.awarded")
        val parts = requireNotNull(ActivityBody.parts(loot))
        assertTrue(ActivityBody.Part("Gorvash", ActivityBody.Tone.ClassName("warrior")) in parts)
        assertTrue(ActivityBody.Part("Deathbringer", ActivityBody.Tone.Quality(4)) in parts)
        assertEquals(loot.body, parts.joinToString("") { it.text })
        assertNull(ActivityBody.figure(loot))
    }

    @Test fun `the guild ladder has strong figures, the passer's crest and the figure trailing`() {
        val passed = row("guildsim.passed")
        assertEquals(ActivityBody.Part("1196", ActivityBody.Tone.Strong), requireNotNull(ActivityBody.parts(passed)).first())
        assertEquals("warrior", Activity.crestClass(passed, SpecCatalog(emptyList())))
        assertEquals("1196.2", ActivityBody.figure(passed))
    }

    @Test fun `a row without segments keeps its plain body and a sim keeps its figure`() {
        val sim = row("sim.finished")
        assertNull(ActivityBody.parts(sim))
        assertNull(ActivityBody.annotated(sim))
        assertEquals("1184.3", ActivityBody.figure(sim))
    }

    @Test fun `Home's age follows the styled body`() {
        val loot = row("loot.awarded")
        assertEquals("${loot.body} · 3 h ago", ActivityBody.annotated(loot, "3 h ago")?.text)
    }

    @Test fun `Before you pull carries the item's icon where the check names one`() {
        val check = RaidFixture.load(ReadyCheck.serializer(), "character.ready")
        val (items, _) = ReadyRows.of(resources(), check, RaidTime.now)
        val icons = items.map { it.title to it.icon }.toMap()
        assertEquals("inv_jewelry_talisman_11", icons["Drakefire Amulet"])
        assertEquals("inv_boots_07", icons["Boots not enchanted"])
        assertNull(icons["Grilled Squid"])
        assertNull(icons["Flask or elixir"])
    }

    @Test fun `Loot history reads the last night's zone, kills and wipes`() {
        val list = RaidFixture.load(GuildRaids.serializer(), "guild.raids")
        val night = list.rows.first { it.raidId != null }
        val presenter = RaidTabPresenter(resources(), clock(), RaidTime.now)
        assertEquals("Last: Onyxia's Lair, 1 kill, 2 wipes", presenter.lootLine(night, null))
        assertNull(presenter.lootLine(null, null))
    }
}
