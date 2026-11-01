package gg.foreversixty.app.search

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.model.ItemUseSource
import gg.foreversixty.app.model.ItemUseSpec
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.search.logic.ItemPageCopy
import gg.foreversixty.app.search.logic.ItemPageOwner
import gg.foreversixty.app.search.logic.ItemPagePhase
import gg.foreversixty.app.search.logic.ItemPageStore
import gg.foreversixty.app.search.logic.ItemUsesPhase
import gg.foreversixty.app.search.logic.SearchCopy
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The item page (spec F12.4, F12.5), mirrored in iOS's ItemPageTests. */
@RunWith(AndroidJUnit4::class)
class ItemPageTest {
    private val res get() = resources()
    private val specs = SearchTestData.specs()
    private val cape = SearchTestData.SPRITECASTER_CAPE

    private fun owner(preset: String) =
        ItemPageOwner("shaman", BisSelection("shaman", "shaman-elemental", BisRange(60, preset), "horde"))

    private suspend fun load(id: Int, owner: ItemPageOwner, source: SearchTestSource = SearchTestSource()): ItemPageStore =
        ItemPageStore(id, ItemLook(id), source, owner, specs, res).also { it.load() }

    private fun ItemPageStore.loaded() = (phase as ItemPagePhase.Loaded).model

    @Test fun `the raid list's pick reads Your Back pick with the band, preset and faction`() = runTest {
        val model = load(cape, owner(BisRange.RAID)).loaded()
        assertEquals("Your Back pick", model.pick?.title)
        assertEquals("Elemental · 60 · Phase 1 raid · Horde", model.pick?.meta)
        assertEquals("Rare · Back · item level 52", model.headMeta)
    }

    @Test fun `a runner-up reads its place for the slot`() = runTest {
        val model = load(cape, owner(BisRange.BARE)).loaded()
        assertEquals("Runner-up 1 for Back", model.pick?.title)
        assertEquals("Elemental · 60 · Horde", model.pick?.meta)
    }

    @Test fun `an item not on your list has no card, and signed out has none`() = runTest {
        assertNull(load(16866, owner(BisRange.RAID)).loaded().pick)
        assertNull(load(cape, ItemPageOwner()).loaded().pick)
    }

    @Test fun `the tooltip is the F2 6 body`() = runTest {
        val tooltip = load(cape, owner(BisRange.RAID)).loaded().tooltip
        assertEquals("Spritecaster Cape", tooltip.look.name)
        assertEquals("Back", tooltip.slotName)
        assertTrue(tooltip.stats.contains("+14 Spell power"))
        assertEquals("Requires Level 47", tooltip.requiresLevel)
    }

    @Test fun `an id the index does not carry is unknown`() = runTest {
        assertEquals(ItemPagePhase.Unknown, load(999_999, owner(BisRange.RAID)).phase)
    }

    @Test fun `uses that fail show the error in their region only`() = runTest {
        val store = load(cape, owner(BisRange.RAID), SearchTestSource(uses = ApiResult.Failure(ApiError.Status(500))))
        assertTrue(store.phase is ItemPagePhase.Loaded)
        assertEquals(ItemUsesPhase.Failed, store.uses)
    }

    @Test fun `the cape's sources and specs come from its shard`() = runTest {
        val uses = load(cape, owner(BisRange.RAID)).uses as ItemUsesPhase.Loaded
        assertEquals("High Interrogator Gerstahn", uses.sources.first().title)
        assertEquals("Blackrock Depths · Dungeon · Drop chance not known", uses.sources.first().meta)
        assertEquals("Balance Druid", uses.specs.first().title)
        assertEquals("60 · Phase 1 raid", uses.specs.first().meta)
        assertEquals("Pick", uses.specs.first().tag)
        assertEquals("Runner-up 3", uses.specs.last().tag)
        assertEquals("50", uses.specs.last().meta)
    }

    @Test fun `a known chance prints its percent and never 0 percent`() {
        val rows = ItemPageCopy.sources(
            listOf(
                ItemUseSource("raid", "Onyxia", "Onyxia's Lair", chance = 19.6),
                ItemUseSource("dungeon", "Baron Rivendare", "Stratholme", chance = 0.3),
                ItemUseSource("vendor", "Meilosh"),
                ItemUseSource("quest", "The Great Fras Siabi", "Horde", minLevel = 58),
                ItemUseSource("world_drop", "World drop"),
            ),
            res,
        )
        assertEquals("Onyxia's Lair · Raid · 20% from Onyxia", rows[0].meta)
        assertEquals("Stratholme · Dungeon · <1% from Baron Rivendare", rows[1].meta)
        assertEquals("Vendor", rows[2].meta)
        assertEquals("Horde · Quest · Requires level 58", rows[3].meta)
        assertEquals("", rows[4].meta)
    }

    @Test fun `six spec rows show, then the count of the rest`() {
        val all = ItemPageCopy.specs(
            (1..11).map { ItemUseSpec("mage-fire", "mage", "Fire", 60, "raid", null, 0) }, specs, SearchCopy(res), res,
        )
        val (shown, more) = ItemPageCopy.visibleSpecs(all, expanded = false)
        assertEquals(ItemPageCopy.VISIBLE_SPECS, shown.size)
        assertEquals(5, more)
        assertEquals(11 to 0, ItemPageCopy.visibleSpecs(all, expanded = true).let { it.first.size to it.second })
        assertEquals("5 more specs", res.getQuantityString(gg.foreversixty.app.R.plurals.itemPage_moreSpecs, more, more))
    }
}
