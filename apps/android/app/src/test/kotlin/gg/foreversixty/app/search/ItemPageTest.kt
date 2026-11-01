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

    private val sash = SearchTestData.SATYRMANE_SASH
    private val pauldrons = SearchTestData.DEFILERS_MAIL_PAULDRONS
    private val raid get() = owner(BisRange.RAID)

    private fun ItemPageStore.elementalRow() = (uses as ItemUsesPhase.Loaded).specs.firstOrNull { it.spec == "shaman-elemental" }

    @Test fun `the raid list's pick reads Your Back pick with the band, preset and faction`() = runTest {
        val store = load(cape, raid)
        assertEquals("Your Back pick", store.loaded().pick?.title)
        assertEquals("Elemental · 60 · Phase 1 raid · Horde", store.loaded().pick?.meta)
        assertEquals("Rare · Back · item level 52", store.loaded().headMeta)
        assertEquals("Pick", store.elementalRow()?.tag)
    }

    @Test fun `a runner-up reads its place for the slot, and Elemental's row agrees`() = runTest {
        val store = load(sash, raid)
        assertEquals("Runner-up 2 for Waist", store.loaded().pick?.title)
        assertEquals("Elemental · 60 · Phase 1 raid · Horde", store.loaded().pick?.meta)
        assertEquals("Runner-up 2", store.elementalRow()?.tag)
        assertEquals("60 · Phase 1 raid", store.elementalRow()?.meta)
    }

    @Test fun `an item no list places has no card and no row, and signed out has no card`() = runTest {
        val store = load(pauldrons, raid)
        assertNull(store.loaded().pick)
        assertNull(store.elementalRow())
        assertNull(load(cape, ItemPageOwner()).loaded().pick)
    }

    @Test fun `the tooltip is the F2 6 body with the client's caster line`() = runTest {
        val tooltip = load(cape, raid).loaded().tooltip
        assertEquals("Spritecaster Cape", tooltip.look.name)
        assertEquals("Binds when picked up", tooltip.binds)
        assertEquals("Back", tooltip.slotName)
        assertEquals("Cloth", tooltip.typeName)
        assertEquals(listOf("+5 Stamina", "+6 Intellect", "+6 Spirit"), tooltip.stats)
        assertEquals(listOf("Equip: Increases damage and healing done by magical spells and effects by up to 14."), tooltip.equipLines)
        assertEquals("Requires Level 47", tooltip.requiresLevel)
    }

    @Test fun `a set piece names its set, and mp5 reads as the client's Equip line`() = runTest {
        val tooltip = load(pauldrons, raid).loaded().tooltip
        assertEquals("The Defiler's Fortitude", tooltip.setName)
        assertTrue(tooltip.setBonuses.isNotEmpty())
        assertEquals(listOf("Equip: Restores 4 mana per 5 sec."), tooltip.equipLines)
    }

    @Test fun `an id the index does not carry is unknown`() = runTest {
        assertEquals(ItemPagePhase.Unknown, load(999_999, raid).phase)
    }

    @Test fun `uses that fail show the error in their region only`() = runTest {
        val store = load(cape, raid, SearchTestSource(uses = ApiResult.Failure(ApiError.Status(500))))
        assertTrue(store.phase is ItemPagePhase.Loaded)
        assertEquals(ItemUsesPhase.Failed, store.uses)
    }

    @Test fun `the cape's sources read place and kind, with the unknown chance said once`() = runTest {
        val uses = load(cape, raid).uses as ItemUsesPhase.Loaded
        assertEquals("High Interrogator Gerstahn", uses.sources.first().title)
        assertEquals("Blackrock Depths · Dungeon", uses.sources.first().meta)
        assertTrue(uses.chanceUnknown)
        assertEquals("Balance Druid", uses.specs.first().title)
        assertEquals("60 · Phase 1 raid", uses.specs.first().meta)
        assertEquals("Pick", uses.specs.first().tag)
        assertEquals("Runner-up 3", uses.specs.last().tag)
        assertEquals("50", uses.specs.last().meta)
    }

    @Test fun `a vendor and a reputation source read their kind`() = runTest {
        val uses = load(pauldrons, raid).uses as ItemUsesPhase.Loaded
        assertEquals(listOf("vendor", "rep"), uses.sources.map { it.kind })
        assertEquals("Vendor", uses.sources[0].meta)
        assertEquals("The Defilers — Revered", uses.sources[1].title)
        assertEquals("Reputation", uses.sources[1].meta)
        assertTrue(!uses.chanceUnknown)
        assertTrue(uses.specs.isEmpty())
    }

    @Test fun `a known chance prints its percent and an unknown one is not on the row`() {
        val sources = listOf(
            ItemUseSource("raid", "Onyxia", "Onyxia's Lair", chance = 19.6),
            ItemUseSource("dungeon", "Baron Rivendare", "Stratholme", chance = 0.3),
            ItemUseSource("dungeon", "Houndmaster Grebmar", "Blackrock Depths"),
            ItemUseSource("quest", "The Great Fras Siabi", "Horde", minLevel = 58),
            ItemUseSource("world_drop", "World drop"),
        )
        val rows = ItemPageCopy.sources(sources, res)
        assertEquals("Onyxia's Lair · Raid · 20% from Onyxia", rows[0].meta)
        assertEquals("Stratholme · Dungeon · <1% from Baron Rivendare", rows[1].meta)
        assertEquals("Blackrock Depths · Dungeon", rows[2].meta)
        assertEquals("Horde · Quest · Requires level 58", rows[3].meta)
        assertEquals("", rows[4].meta)
        assertTrue(ItemPageCopy.chanceUnknown(sources))
        assertTrue(!ItemPageCopy.chanceUnknown(sources.take(2)))
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
