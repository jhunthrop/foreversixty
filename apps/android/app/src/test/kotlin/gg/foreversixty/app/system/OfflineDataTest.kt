package gg.foreversixty.app.system

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.network.InMemoryStaticFileCache
import gg.foreversixty.app.network.SavedStaticFile
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.system.offline.IconCache
import gg.foreversixty.app.system.offline.IconCacheSize
import gg.foreversixty.app.system.offline.MemoryOfflineSettings
import gg.foreversixty.app.system.offline.OfflineCategory
import gg.foreversixty.app.system.offline.OfflineData
import gg.foreversixty.app.system.offline.OfflineDataCopy
import gg.foreversixty.app.system.offline.OfflineDataModel
import gg.foreversixty.app.system.offline.UpdateNote
import gg.foreversixty.app.system.offline.UpdateOutcome
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.FixtureEngine
import java.io.File
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** You › Offline data (spec F15.4): rows, sizes, times and the model's actions; tenet 7, no ids or counts. */
@RunWith(AndroidJUnit4::class)
class OfflineDataTest {
    private val zone: ZoneId = ZoneId.of("America/Chicago")
    private val now: Instant = Instant.parse("2026-11-06T03:41:00Z")
    private val build = "1.60.1.70658"
    private val earlier = now.minusSeconds(3_600).toEpochMilli()

    private fun file(path: String, bytes: Long, at: Long = earlier) = SavedStaticFile(path, bytes, at)

    private val files = listOf(
        file("/data/$build/guides/index.json", 40_000),
        file("/data/$build/guides/shaman-elemental.json", 1_200_000),
        file("/data/$build/guides/warrior-fury.json", 1_160_000),
        file("/data/$build/bis/shaman-elemental.json", 2_100_000),
        file("/data/$build/items/shaman.json", 1_000_000),
        file("/data/$build/talents/shaman.json", 600_000),
        file("/data/$build/search.json", 400_000),
        file("/data/active-build.json", 100),
    )

    @Test fun `each file lands in its row and the pointer in none`() {
        assertEquals(OfflineCategory.Guides, OfflineData.category("/data/$build/guides/index.json"))
        assertEquals(OfflineCategory.Guides, OfflineData.category("/data/$build/tiers.json"))
        assertEquals(OfflineCategory.Guides, OfflineData.category("/data/changelog.json"))
        assertEquals(OfflineCategory.Bis, OfflineData.category("/data/$build/item-uses/3.json"))
        assertEquals(OfflineCategory.Talents, OfflineData.category("/data/$build/talents/mage.json"))
        assertEquals(OfflineCategory.Search, OfflineData.category("/data/$build/search.json"))
        assertNull(OfflineData.category("/data/active-build.json"))
    }

    @Test fun `the classes come from the spec and class files`() {
        assertEquals("shaman", OfflineData.classOf("/data/$build/guides/shaman-elemental.json"))
        assertEquals("warrior", OfflineData.classOf("/data/$build/bis/warrior-fury.json"))
        assertEquals("mage", OfflineData.classOf("/data/$build/talents/mage.json"))
        assertNull(OfflineData.classOf("/data/$build/guides/index.json"))
        assertNull(OfflineData.classOf("/data/$build/items/mage.json"))
    }

    @Test fun `sizes read in whole KB, one decimal MB under ten, whole MB above`() {
        assertEquals("0 KB", OfflineData.size(0))
        assertEquals("1 KB", OfflineData.size(120))
        assertEquals("410 KB", OfflineData.size(410_000))
        assertEquals("2.4 MB", OfflineData.size(2_400_000))
        assertEquals("9.9 MB", OfflineData.size(9_940_000))
        assertEquals("41 MB", OfflineData.size(41_000_000))
    }

    @Test fun `the report rows match the board's words`() {
        val report = OfflineData.report(files, IconCacheSize(41_000_000, Instant.ofEpochMilli(earlier)))
        assertEquals(
            listOf(OfflineCategory.Guides, OfflineCategory.Bis, OfflineCategory.Talents, OfflineCategory.Search, OfflineCategory.Icons),
            report.rows.map { it.category },
        )
        val copy = OfflineDataCopy(resources(), zone)
        val metas = report.rows.map { copy.meta(it, now) }
        assertEquals("Shaman, Warrior · 2.4 MB · Updated today", metas[0])
        assertEquals("Shaman · 3.1 MB · Updated today", metas[1])
        assertEquals("Shaman · 600 KB · Updated today", metas[2])
        assertEquals("400 KB · Updated today", metas[3])
        assertEquals("41 MB", metas[4])
        assertEquals("On this phone, 48 MB", copy.section(report))
    }

    @Test fun `updated reads today, yesterday, a weekday, then a date`() {
        val day = 86_400L
        fun updated(daysAgo: Long) = OfflineData.updated(now.minusSeconds(daysAgo * day), now, zone, "today", "yesterday")
        assertEquals("today", updated(0))
        assertEquals("yesterday", updated(1))
        assertEquals("Mon", updated(3))
        assertEquals("21 Oct", updated(15))
    }

    @Test fun `the footer says when it was checked`() {
        val copy = OfflineDataCopy(resources(), zone)
        assertEquals(
            "Checked today at 9:41 pm. Your characters, sign-ups and sims are not stored here; they load when you are online.",
            copy.footer(now, now),
        )
        assertTrue(copy.footer(null, now).startsWith("Not checked yet."))
    }

    @Test fun `nothing saved is the empty state`() {
        assertTrue(OfflineData.report(emptyList(), IconCacheSize(0, null)).isEmpty)
    }

    // The model

    private class FakeIcons(var bytes: Long) : IconCache {
        var cleared = false
        override fun size() = IconCacheSize(bytes, null)
        override fun clear() {
            cleared = true
            bytes = 0
        }
    }

    private fun model(outcome: UpdateOutcome = UpdateOutcome.UpToDate, online: Boolean = true, icons: FakeIcons = FakeIcons(0)): Pair<OfflineDataModel, InMemoryStaticFileCache> {
        val cache = InMemoryStaticFileCache().apply { write("{}", "/data/$build/guides/shaman-elemental.json", earlier) }
        val store = StaticDataStore(StaticDataClient(FixtureEngine.create(ClasspathFixtureSource(requireNotNull(javaClass.classLoader)))), cache)
        val settings = MemoryOfflineSettings()
        return OfflineDataModel(store, icons, settings, Connectivity.fixed(online), { outcome }, { now }) to cache
    }

    @Test fun `remove deletes the store and the icons and shows the empty state`() = runTest {
        val icons = FakeIcons(5_000)
        val (model, cache) = model(icons = icons)
        model.load()
        assertFalse(model.report?.isEmpty ?: true)
        model.remove()
        assertTrue(cache.files().isEmpty())
        assertTrue(icons.cleared)
        assertTrue(model.report?.isEmpty == true)
    }

    @Test fun `update now says up to date and records the check`() = runTest {
        val (model, _) = model(UpdateOutcome.UpToDate)
        model.updateNow()
        assertEquals(UpdateNote.UpToDate, model.note)
        assertEquals(now, model.checkedAt)
        assertFalse(model.updating)
    }

    @Test fun `a failed update says so and keeps the last check`() = runTest {
        val (model, _) = model(UpdateOutcome.Failed)
        model.updateNow()
        assertEquals(UpdateNote.Failed, model.note)
        assertNull(model.checkedAt)
    }

    @Test fun `the switches persist and default on`() {
        val (model, _) = model()
        assertTrue(model.keepClassesReady && model.wifiOnly)
        model.wifiOnly(false)
        assertFalse(model.settings.wifiOnly)
        assertFalse(model.wifiOnly)
    }

    @Test fun `offline, update now is not available`() {
        assertFalse(model(online = false).first.online)
    }

    @Test fun `the icon folder reports its bytes`() {
        val folder = File.createTempFile("icons", "").apply { delete(); mkdirs() }
        File(folder, "a").writeBytes(ByteArray(2_000))
        val size = gg.foreversixty.app.system.offline.FolderIconCache(folder).size()
        assertEquals(2_000L, size.bytes)
        folder.deleteRecursively()
    }
}
