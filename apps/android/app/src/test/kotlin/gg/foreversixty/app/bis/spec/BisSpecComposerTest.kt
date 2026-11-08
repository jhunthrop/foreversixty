package gg.foreversixty.app.bis.spec

import android.content.Context
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.bis.band
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotEquals
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith
import java.time.ZoneOffset

@RunWith(AndroidJUnit4::class)
class BisSpecComposerTest {
    private val res = ApplicationProvider.getApplicationContext<Context>().resources
    private val composer = BisSpecComposer(res, BisTestData.catalog(), ZoneOffset.UTC)
    private val file = BisTestData.bis()
    private val items = BisTestData.items().items

    private fun input(selection: gg.foreversixty.app.bis.BisSelection = BisTestData.raid, off: Set<BisSource> = emptySet()) =
        BisSpecInput(file, requireNotNull(file.band(selection.range, selection.faction)), items, selection, "Warrior", off)

    @Test fun `the raid list words every slot and its head`() {
        val content = composer.content(input())
        val band = requireNotNull(file.band(BisTestData.raid.range, "horde"))
        assertEquals("Fury Warrior", content.title)
        assertEquals(band.slots.size, content.cells.size)
        val picked = band.slots.count { (it.itemId ?: 0) > 0 }
        assertEquals("$picked of ${band.slots.size} slots have a pick at this level.", content.summary.count)
        assertNull("the top band has no next band", content.summary.nextBand)
        assertEquals(listOf("Raid-ready", "Bare"), content.controls.profiles.map { it.label })
        assertTrue(content.controls.levels.single { it.on }.label == "60")
        assertEquals("All sources", content.controls.sources)
        assertEquals("Fury Warrior guide ›", content.guide)
        assertNotNull(content.answer.value)
    }

    @Test fun `where to go holds every picked cell once`() {
        val content = composer.content(input())
        assertEquals(content.cells.count { it.look != null }, content.groups.sumOf { it.cells.size })
    }

    @Test fun `a lower band links the next one`() {
        val selection = BisTestData.raid.copy(range = BisRange(40, BisRange.BARE))
        val band = file.band(selection.range, selection.faction) ?: return
        val content = composer.content(BisSpecInput(file, band, items, selection, "Warrior"))
        assertEquals("Changes from 40 to 49: see 50 to 59 ›", content.summary.nextBand)
        assertEquals(BisRange(50, BisRange.BARE), content.summary.nextPick?.range)
        assertTrue(content.controls.profiles.isEmpty())
    }

    @Test fun `a source turned off falls to the next runner-up`() {
        val all = composer.content(input())
        val kinds = composer.sources(input().band)
        val off = BisSource.of(input().band.slots.first { (it.itemId ?: 0) > 0 && !it.alternatives.isNullOrEmpty() }.sourceKind)
        val filtered = composer.content(input(off = setOf(off)))
        assertNotEquals(all.cells.map { it.look?.id }, filtered.cells.map { it.look?.id })
        assertEquals("Sources (${kinds.size - 1})", filtered.controls.sources)
    }
}
