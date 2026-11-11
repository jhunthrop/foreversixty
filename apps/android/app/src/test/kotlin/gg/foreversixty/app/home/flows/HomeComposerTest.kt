package gg.foreversixty.app.home.flows

import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.gear.ItemSource
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.gear.WhereToGo
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.model.BisAlternative
import gg.foreversixty.app.model.BisSlot
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomePick
import gg.foreversixty.app.model.HomeSection
import gg.foreversixty.app.model.HomeTalents
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.LootBoss
import gg.foreversixty.app.model.LootFile
import gg.foreversixty.app.model.LootSource
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith
import gg.foreversixty.app.activity.ActivityTime
import java.time.Instant
import java.time.ZoneId

/** Home's composition from `/v1/home` sections, and Where to go's Best run pick. */
@RunWith(AndroidJUnit4::class)
class HomeComposerTest {
    private val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
    private val now = Instant.parse("2026-11-06T18:00:00Z")
    private val composer = HomeComposer(res, now, ZoneId.of("America/Chicago"))

    private fun upgrade(slot: String, id: Int, name: String, gain: Double?) =
        HomeUpgrade(slot, slot.replaceFirstChar(Char::uppercaseChar), HomePick(id, name, 4, "icon", 60, "Crafted: Blacksmithing"), null, gain, "+$gain DPS", gain != null)

    private val home = CharacterHome(
        character = "us/normal/thandrel", state = "ok", spec = "warrior-fury", band = 60, preset = "raid", unitWord = "DPS",
        dataBuild = "1.60.1.70790", upgradeCount = 4, notSimCheckedCount = 0, notModelledCount = 1, talents = HomeTalents("optimized"),
        upgrades = listOf(upgrade("finger1", 3, "Band", 6.8), upgrade("head", 1, "Lionheart Helm", 14.5), upgrade("trinket1", 9, "Odd", null), upgrade("off_hand", 2, "Greenhammer", 9.7)),
    )

    private fun row(id: Long, kind: String, at: String, group: String = "gear") =
        ActivityRow(id = id, kind = kind, group = group, title = "t$id", body = "b$id", link = "", occurredAt = at, read = id % 2 == 0L, icon = "i")

    private fun inputs(sections: List<HomeSection>, card: RaidCardModel? = null, rows: List<ActivityRow> = emptyList(), home: CharacterHome = this.home) =
        HomeInputs(sections, home, RunPick("back", 7, "Blackcrow", 1.94, ItemSource(WhereToGo.DUNGEON, "Blackrock Spire", 8.0)), card, null, rows, 3, null)

    @Test fun `the head names level, spec and class, then the guild or the levels to the next band`() {
        val warrior = HomeCharacter("k", "us", "normal", "Thandrel", null, "warrior", "Fury", 60, "Olympus XXVII", "alliance")
        assertEquals(listOf("Level 60 Fury Warrior", "Olympus XXVII"), composer.head(warrior).facts)
        val hunter = warrior.copy(name = "Felsnap", className = "hunter", specName = "Beast Mastery", level = 34)
        assertEquals(listOf("Level 34 Beast Mastery Hunter", "6 levels to the 40 list"), composer.head(hunter).facts)
        assertEquals("1 level to the 40 list", composer.head(hunter.copy(level = 39)).facts[1])
    }

    @Test fun `blocks follow the sections, upgrades largest first with the Best run line`() {
        val blocks = composer.blocks(inputs(listOf(HomeSection("head"), HomeSection("upgrades", limit = 3), HomeSection("mystery"), HomeSection("guild_cta"))))
        val upgrades = blocks[0] as HomeBlock.Upgrades
        assertEquals(listOf("Lionheart Helm", "Greenhammer", "Band"), upgrades.rows.map { it.look.name })
        assertEquals("+14.5", upgrades.rows[0].gain)
        // Round 2: Best run is a row (the item, the dungeon and its chance, the gain).
        val run = requireNotNull(upgrades.bestRun)
        assertEquals("Blackcrow", run.look.name)
        assertEquals("Best run · Blackrock Spire · 8%", run.source)
        assertEquals("+1.9", run.gain)
        assertTrue(blocks[1] is HomeBlock.FindGuild)
        assertEquals(2, blocks.size)
    }

    @Test fun `the section meta counts the upgrades and their total, the stamp says when gear synced`() {
        val full = home.copy(totalGain = 157.6, syncedAt = "2026-11-06T16:00:00Z", syncedFrom = "battlenet")
        assertEquals("4 upgrades · +158 DPS", composer.upgradesMeta(full, "DPS"))
        assertEquals("4 upgrades", composer.upgradesMeta(home, "DPS"))
        assertEquals(true, composer.syncedStamp(full)?.startsWith("Synced 2"))
        assertEquals(true, composer.syncedStamp(full.copy(syncedFrom = "addon"))?.endsWith("from the addon"))
        assertNull(composer.syncedStamp(home))
    }

    @Test fun `an upgrade's line ends with its drop chance and the worn item it replaces`() {
        val worn = gg.foreversixty.app.model.HomeWorn(id = 5, name = "Crown of Destruction", known = true, label = "Crown of Destruction")
        val pick = HomePick(1, "Lionheart Helm", 4, "icon", 60, "Onyxia's Lair", dropChance = 4.2)
        val upgrades = listOf(HomeUpgrade("head", "Head", pick, worn, 14.5, "+14.5 DPS", true))
        val one = composer.blocks(
            HomeInputs(listOf(HomeSection("upgrades")), home.copy(upgrades = upgrades), null, null, null, emptyList(), 0, null),
        ).first() as HomeBlock.Upgrades
        assertEquals("Head · Onyxia's Lair · 4% · replaces Crown of Destruction", one.rows.single().source)
    }

    @Test fun `the raid day card leaves the Best run line out`() {
        val card = HomeFixtures.raid()
        val blocks = composer.blocks(inputs(listOf(HomeSection("next_raid", style = "card"), HomeSection("upgrades", limit = 3)), card))
        assertTrue(blocks[0] is HomeBlock.NextRaidCard)
        assertNull((blocks[1] as HomeBlock.Upgrades).bestRun)
    }

    @Test fun `the day after, Last night takes the excluded kinds and New for you reads past them`() {
        val rows = listOf(
            row(5, "loot.awarded", "2026-11-06T03:00:00Z"),
            row(4, "guildsim.passed", "2026-11-06T02:00:00Z", "ladder"),
            row(3, "report.ready", "2026-11-06T03:10:00Z", "logs"),
            row(2, "sim.finished", "2026-11-05T15:00:00Z"),
        )
        val sections = listOf(HomeSection("raid_summary"), HomeSection("activity", limit = 3, excludeKinds = listOf("loot.awarded", "report.ready")))
        val blocks = composer.blocks(inputs(sections, rows = rows))
        val last = blocks[0] as HomeBlock.LastNight
        assertEquals(listOf("t5", "t3"), last.rows.map { it.row.title })
        assertTrue(last.simAgain)
        val preview = blocks[1] as HomeBlock.Activity
        assertEquals(listOf("t4", "t2"), preview.rows.map { it.row.title })
        // Row 3: the body alone, the when on its own meta line in Activity's buckets.
        assertEquals("b2", preview.rows[1].line)
        assertEquals("9:00 am", preview.rows[1].age)
        assertEquals(ActivityTime.label(res, Instant.parse("2026-11-05T15:00:00Z"), now, ZoneId.of("America/Chicago")), preview.rows[1].age)
    }

    @Test fun `a sync over a week old takes the warning tone (row 12)`() {
        val fresh = composer.blocks(inputs(listOf(HomeSection("upgrades")), home = home.copy(syncedAt = "2026-11-04T18:00:00Z")))
        val stale = composer.blocks(inputs(listOf(HomeSection("upgrades")), home = home.copy(syncedAt = "2026-10-27T18:00:00Z")))
        assertEquals(false, (fresh.first() as HomeBlock.Upgrades).syncedStale)
        assertEquals(true, (stale.first() as HomeBlock.Upgrades).syncedStale)
    }

    @Test fun `Best run is the dungeon candidate with the largest gain over the worn item`() {
        val slots = listOf(
            BisSlot("head", true, itemId = 1, itemName = "Lionheart Helm"),
            BisSlot("back", true, itemId = 20, itemName = "Cape", alternatives = listOf(BisAlternative(7, "Blackcrow", "Ubrs", "drop", -0.5))),
        )
        val loot = LootFile(
            sources = listOf(
                LootSource("dungeon:ubrs", "dungeon", "Blackrock Spire", bosses = listOf(LootBoss("b", "Drakkisath", items = listOf(7), itemChances = mapOf("7" to 8.0)))),
                LootSource("crafted:bs", "crafted", "Blacksmithing", items = listOf(1), requiredSkill = mapOf("1" to 245)),
            ),
        )
        // Round 8 row 1: a crafted source carries the profession's skill when the loot file has it.
        assertEquals(245, WhereToGo.sources(loot)[1]?.single()?.skill)
        val upgrades = listOf(upgrade("head", 1, "Lionheart Helm", 14.5), upgrade("back", 20, "Cape", 2.4))
        val picks = WhereToGo.picks(upgrades, slots, WhereToGo.sources(loot))
        assertEquals(listOf("Lionheart Helm", "Blackcrow"), picks.map { it.itemName })
        val best = WhereToGo.bestRun(picks)
        assertEquals("Blackcrow", best?.itemName)
        assertEquals(1.9, best?.gain ?: 0.0, 1e-9)
        assertEquals(8.0, best?.source?.chance)
    }
}
