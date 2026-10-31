package gg.foreversixty.app.bis

import android.content.res.Resources
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.BisSlot
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.model.ItemsFile
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.model.KeyDate
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSources
import kotlinx.serialization.builtins.ListSerializer
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith
import java.io.File
import java.time.ZoneOffset

/** The BiS list's rules (spec F3.1); the same list as iOS's BisListTests. */
@RunWith(AndroidJUnit4::class)
class BisListTest {
    private val res: Resources = ApplicationProvider.getApplicationContext<android.content.Context>().resources

    private fun raidList(worn: Set<Int> = emptySet()): BisList {
        val file = BisTestData.bis()
        val band = requireNotNull(file.band(BisTestData.raid.range, "horde"))
        return BisListBuilder(res, BisTestData.items().items, worn, ZoneOffset.UTC)
            .list(file, band, "Fury Warrior", res.getString(gg.foreversixty.app.R.string.bis_range_p1))
    }

    @Test
    fun `the level range follows the site's band rule and 60 reads the raid list`() {
        assertEquals(BisRange(50, "bare"), BisRange.forLevel(58))
        assertEquals(BisRange(60, "raid"), BisRange.forLevel(60))
        assertEquals(BisRange(20, "bare"), BisRange.forLevel(12))
        assertEquals(BisRange(20, "bare"), BisRange.forLevel(29))
        assertEquals(
            listOf("20 to 29", "30 to 39", "40 to 49", "50 to 59", "60", "60 · Phase 1 raid"),
            BisRange.all.map { BisLabels.range(res, it) },
        )
    }

    @Test
    fun `a band is the file's own entry for range, faction and preset, else the bare list`() {
        val file = BisTestData.bis()
        assertEquals("raid", file.band(BisRange(60, "raid"), "horde")?.preset)
        assertEquals("bare", file.band(BisRange(50, "raid"), "horde")?.preset)
        assertNull(file.band(BisRange(60, "raid"), "alliance"))
    }

    @Test
    fun `17 rows in the file's order, never re-ranked`() {
        val list = raidList()
        val band = requireNotNull(BisTestData.bis().band(BisTestData.raid.range, "horde"))
        assertEquals(band.slots.map { it.slot }, list.rows.map { it.slot })
        assertEquals(17, list.rows.size)
        assertEquals(band.slots.mapNotNull { it.itemId }, list.rows.mapNotNull { it.pick?.item?.id })
    }

    @Test
    fun `the set line carries the figure, the unit and the Updated date`() {
        assertEquals("894.1 DPS at best in slot for your level range · Updated 30 Oct 2026", raidList().setDpsLine)
    }

    @Test
    fun `the next line is a signed gap with a true minus, same below 0_05, or nothing else ranked`() {
        val rows = raidList().rows
        assertEquals("Next: Champion's Plate Helm, −15.0 DPS", rows[0].pick?.next)
        assertEquals("Next: Warlord's Plate Shoulders, same DPS", rows[2].pick?.next)
        assertEquals("Nothing else ranked", BisListBuilder(res, emptyList(), emptySet()).nextLine(null, BisUnit.Dps))
        assertNull(BisListBuilder.signedDelta(-0.049))
        assertEquals("+0.1", BisListBuilder.signedDelta(0.05))
    }

    @Test
    fun `healer and tank bands name their own measure`() {
        assertEquals("HPS", res.getString(BisUnit.forRole("healer").word))
        assertEquals("score", res.getString(BisUnit.forRole("tank").word))
        assertEquals("DPS", res.getString(BisUnit.forRole(null).word))
    }

    @Test
    fun `the source line in the site's words`() {
        val rows = raidList().rows
        assertEquals("Head · Crafted: Blacksmithing", rows[0].pick?.slotSource)
        assertEquals("Neck · Quest: Falrin's Vendetta", rows[1].pick?.slotSource)
        assertEquals("Shoulder · Vendor: Creeg Bothunk", rows[2].pick?.slotSource)
        assertEquals("Wrist · Warsong Outriders", rows[5].pick?.slotSource)
        assertEquals("Finger 1", rows[10].slotLabel)
        assertEquals("Finger 2", rows[11].slotLabel)
        assertEquals("Trinket 1", rows[12].slotLabel)
        assertEquals(
            "PvP rank 14 · Champion · Horde",
            BisListBuilder(res, emptyList(), emptySet()).sourceLine("PvP rank 14 · Champion · Horde", "vendor"),
        )
    }

    @Test
    fun `an unmodelled proc is said and the worn pick is tagged`() {
        val rows = raidList(worn = setOf(12640)).rows
        assertEquals(true, rows[10].pick?.isUnmodelled)
        assertEquals(false, rows[0].pick?.isUnmodelled)
        assertEquals(true, rows[0].pick?.isWorn)
        assertEquals(false, rows[1].pick?.isWorn)
    }

    @Test
    fun `an off hand under a two-hander and a slot with no pick`() {
        val builder = BisListBuilder(res, emptyList(), emptySet())
        val offHand = builder.row(BisSlot("off_hand", verified = false, emptyReason = "two_hand_equipped"), BisUnit.Dps, "Argent Crusader")
        assertEquals("Off hand · none: Argent Crusader is two-handed.", offHand.emptyReason)
        val trinket = builder.row(BisSlot("trinket1", verified = false, emptyReason = "no_sourced_item"), BisUnit.Dps, null)
        assertEquals("Trinket 1 · no pick for this range.", trinket.emptyReason)
    }

    @Test
    fun `search keeps rows whose pick or runner-up matches`() {
        val list = raidList()
        assertEquals(listOf("head"), list.filtered("lionheart").map { it.slot })
        assertEquals(listOf("head"), list.filtered("Champion's Plate Helm").map { it.slot })
        assertEquals(17, list.filtered(" ").size)
    }

    @Test
    fun `every published BiS file of the active build decodes`() {
        val files = File(BisTestData.repo, "data/builds/${BisTestData.BUILD}/bis").listFiles { f -> f.extension == "json" }.orEmpty()
        assertTrue(files.size >= 28)
        files.forEach { ApiJson.decodeFromString(BisFile.serializer(), it.readText()) }
    }

    @Test
    fun `the character's own list is its spec, range and faction`() {
        val catalog = BisTestData.catalog()
        val main = requireNotNull(BisTestData.me().mainCharacter)
        assertEquals(BisSelection("warrior", "warrior-fury", BisRange(50, "bare"), "horde"), BisSelection.forCharacter(main, catalog))
        assertEquals("warrior", BisSelection.signedOutDefault(catalog).classSlug)
        assertEquals(listOf("Elemental", "Enhancement", "Restoration"), catalog.specsOf("shaman").map { it.name })
    }

    @Test
    fun `Yours and the way back follow the character's list`() {
        val own = BisTestData.raid
        val context = BisContext(selection = own, characterSelection = own, characterName = "Obnoxious Yell")
        assertTrue(context.isYours)
        assertNull(context.backTarget)
        val away = context.copy(selection = own.copy(faction = "alliance"))
        assertFalse(away.isYours)
        assertEquals("Obnoxious Yell", away.backTarget?.first)
        val signedOut = BisContext(selection = own)
        assertFalse(signedOut.isYours)
        assertNull(signedOut.backTarget)
    }

    private fun real(spec: String): BisFile =
        ApiJson.decodeFromString(BisFile.serializer(), File(BisTestData.repo, "data/builds/${BisTestData.BUILD}/bis/$spec.json").readText())

    private fun request(file: BisFile, band: Int, preset: String, slotIndex: Int, wornId: Int? = null): ItemTooltipRequest {
        val slot = requireNotNull(file.band(BisRange(band, preset), "horde")).slots[slotIndex]
        val id = requireNotNull(slot.itemId)
        return ItemTooltipRequest(
            itemId = id, fallback = ItemLook(id, slot.itemName), build = BisTestData.BUILD, classSlug = "warrior",
            spec = file.spec, specName = "Fury Warrior", band = band, bandLabel = BisLabels.range(res, BisRange(band, preset)),
            preset = preset, faction = "horde", sourceLine = slot.source, wornId = wornId,
        )
    }

    private fun realItems(classSlug: String): List<Item> =
        ApiJson.decodeFromString(ItemsFile.serializer(), File(BisTestData.repo, "data/builds/${BisTestData.BUILD}/items/$classSlug.json").readText()).items

    private fun model(file: BisFile, request: ItemTooltipRequest, items: List<Item> = BisTestData.items().items) =
        ItemTooltipModel.build(request, ItemTooltipSources(items.associateBy { it.id }, emptyList(), file), res)

    @Test
    fun `the tooltip's next best reads in the band's unit`() {
        val fury = BisTestData.bis()
        assertEquals("\u221215.0 DPS vs the pick", model(fury, request(fury, 60, "raid", 0)).alternatives.first().meta)
        listOf("priest-holy" to "HPS", "warrior-protection" to "score").forEach { (spec, word) ->
            val file = real(spec)
            val index = requireNotNull(file.band(BisRange(60, "raid"), "horde")).slots
                .indexOfFirst { s -> s.alternatives.orEmpty().take(2).any { kotlin.math.abs(it.dpsDelta) >= 0.05 } }
            val items = realItems(spec.substringBefore('-'))
            assertTrue(model(file, request(file, 60, "raid", index), items).alternatives.any { it.meta.endsWith(" $word vs the pick") })
        }
    }

    @Test
    fun `the tooltip footer names the slot, spec and band, never a preset id`() {
        val fury = BisTestData.bis()
        assertEquals("Best in slot for Head · Fury Warrior 50 to 59", model(fury, request(fury, 50, "bare", 0)).pickFooter)
        assertEquals("Best in slot for Finger 1 · Fury Warrior 60 · Phase 1 raid", model(fury, request(fury, 60, "raid", 10)).pickFooter)
    }

    @Test
    fun `the worn runner-up leads with You wear this`() {
        val fury = BisTestData.bis()
        val alt = requireNotNull(fury.band(BisRange(60, "raid"), "horde")).slots[0].alternatives.orEmpty().first()
        assertTrue(model(fury, request(fury, 60, "raid", 0, wornId = alt.itemId)).alternatives.first().meta.startsWith("You wear this · "))
    }

    @Test
    fun `a tank list explains its score instead of a damage headline`() {
        val file = real("warrior-protection")
        val band = requireNotNull(file.band(BisRange(60, "raid"), "horde"))
        val list = BisListBuilder(res, emptyList(), emptySet(), ZoneOffset.UTC).list(file, band, "Protection Warrior", "60")
        assertEquals(BisUnit.Score, list.unit)
        assertTrue(list.setDpsLine.orEmpty().startsWith("Tank lists rank by score"))
        assertTrue(list.rows.mapNotNull { it.pick }.any { it.next.endsWith(" score") })
    }

    @Test
    fun `the raid chip says when raids open, from the data, until they do`() {
        val dates = BisTestData.fixture("cdn.dates", ListSerializer(KeyDate.serializer()))
        val opening = requireNotNull(RaidsOpening.from(dates))
        assertEquals("Dec 9", opening.label)
        val raid = BisRange(60, "raid")
        assertEquals("60 · Phase 1 raid · opens Dec 9", BisLabels.range(res, raid, opening, BisTestData.NOW))
        assertEquals("60 · Phase 1 raid", BisLabels.range(res, raid, opening, opening.startsMillis))
        assertEquals("60", BisLabels.range(res, BisRange(60, "bare"), opening, BisTestData.NOW))
        assertNull(RaidsOpening.from(emptyList()))
    }

    @Test
    fun `the app carries the spec catalogue and the fallback build`() {
        val context = ApplicationProvider.getApplicationContext<android.content.Context>()
        val specs = context.assets.open(SpecCatalog.ASSET).bufferedReader().use { it.readText() }
        assertTrue(SpecCatalog.decode(specs).specs.size >= 28)
        assertNotNull(context.assets.open(BUNDLED_BUILD).bufferedReader().use { it.readText() })
    }
}
