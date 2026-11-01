package gg.foreversixty.app.guides

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.guides.logic.GuideStore
import gg.foreversixty.app.guides.logic.GuidesNavigator
import gg.foreversixty.app.guides.logic.GuidesScreen
import gg.foreversixty.app.guides.logic.GuidesSegment
import gg.foreversixty.app.guides.logic.GuidesStore
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.auth.session.MemoryValue
import java.time.Instant
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** F9's data and offline rules (spec §5.1 Sync, §5.7): cache first, saved for offline, new data waits. */
@RunWith(AndroidJUnit4::class)
class GuidesStoreTest {
    private val savedAt: Instant = Instant.ofEpochMilli(FakeGuideFiles.SAVED_AT)

    private fun Loadable<GuideIndex>.index(): GuideIndex = (this as Loadable.Loaded).value
    private fun Loadable<GuideFile>.guide(): GuideFile = (this as Loadable.Loaded).value

    @Test
    fun `the tab reads the active build, then the index`() = runTest {
        val files = FakeGuideFiles()
        val store = GuidesStore(files)
        store.start()
        assertEquals("1.60.1.70658", store.build)
        assertEquals(9, store.index.index().classes.size)
        assertNull((store.index as Loadable.Loaded).offlineSince)
        assertEquals(listOf("/data/active-build.json", "/data/1.60.1.70658/guides/index.json"), files.loads)
    }

    @Test
    fun `offline with a saved index shows it under the offline strip`() = runTest {
        val files = FakeGuideFiles()
        GuidesStore(files).start()
        files.offline = true
        val store = GuidesStore(files)
        store.start()
        assertEquals(savedAt, (store.index as Loadable.Loaded).offlineSince)
    }

    @Test
    fun `offline with nothing saved says so, and a server error is the error state`() = runTest {
        val offline = GuidesStore(FakeGuideFiles().apply { offline = true })
        offline.start()
        assertEquals(Loadable.NotCached, offline.index)
        val store = GuidesStore(FakeGuideFiles())
        store.start()
        val guide = GuideStore("priest-shadow", FakeGuideFiles())
        guide.load(requireNotNull(store.build))
        assertEquals(Loadable.Failed, guide.phase)
    }

    @Test
    fun `a guide once opened is saved and opens offline from the cache`() = runTest {
        val files = FakeGuideFiles()
        val store = GuidesStore(files)
        store.start()
        val build = requireNotNull(store.build)
        GuideStore(GuidesTestData.ELEMENTAL, files).load(build)
        files.offline = true
        val reopened = GuideStore(GuidesTestData.ELEMENTAL, files)
        reopened.load(build)
        assertEquals("Elemental", reopened.phase.guide().specName)
        assertEquals(savedAt, (reopened.phase as Loadable.Loaded).offlineSince)
        store.refreshSaved()
        assertEquals(setOf(GuidesTestData.ELEMENTAL), store.saved.keys)
    }

    @Test
    fun `a saved guide draws at once, before the network answers`() = runTest {
        val files = FakeGuideFiles()
        GuideStore(GuidesTestData.ELEMENTAL, files).load(GuidesTestData.BUILD)
        val store = GuideStore(GuidesTestData.ELEMENTAL, files)
        store.drawSaved(GuidesTestData.BUILD)
        assertEquals("Elemental", store.phase.guide().specName)
    }

    @Test
    fun `an offline guide never saved says it is not on the phone`() = runTest {
        val store = GuideStore(GuidesTestData.FIRE, FakeGuideFiles().apply { offline = true })
        store.load(GuidesTestData.BUILD)
        assertEquals(Loadable.NotCached, store.phase)
    }

    @Test
    fun `prefetch saves every guide of the player's classes`() = runTest {
        val files = FakeGuideFiles()
        val store = GuidesStore(files)
        store.start()
        store.prefetch(setOf("shaman", "warrior"))
        // The fixtures carry three guides; the others answer 404 and stay unsaved.
        assertEquals(setOf(GuidesTestData.ELEMENTAL, GuidesTestData.PROTECTION), store.saved.keys)
    }

    @Test
    fun `a new build waits behind the new-data strip until Update`() = runTest {
        val files = FakeGuideFiles()
        val store = GuidesStore(files)
        store.start()
        assertFalse(store.hasNewData)
        files.activeBuild = GuidesTestData.BUILD
        store.checkForNewBuild()
        assertTrue(store.hasNewData)
        assertEquals("1.60.1.70658", store.build)
        assertTrue(store.applyUpdate())
        assertEquals(GuidesTestData.BUILD, store.build)
        assertFalse(store.hasNewData)
        assertFalse(store.updateFailed)
    }

    @Test
    fun `an update that fails keeps the open build and says so`() = runTest {
        val files = FakeGuideFiles()
        val store = GuidesStore(files)
        store.start()
        files.activeBuild = GuidesTestData.BUILD
        store.checkForNewBuild()
        files.offline = true
        assertFalse(store.applyUpdate())
        assertTrue(store.updateFailed)
        assertTrue(store.hasNewData)
        assertEquals("1.60.1.70658", store.build)
    }

    @Test
    fun `tiers and what changed load their own regions`() = runTest {
        val files = FakeGuideFiles()
        val store = GuidesStore(files)
        store.start()
        store.loadTiers()
        store.loadChanges()
        assertTrue(store.tiers is Loadable.Loaded)
        assertTrue(store.changelog is Loadable.Loaded)
        assertTrue(store.dates is Loadable.Loaded)
        val offline = GuidesStore(FakeGuideFiles().apply { offline = true })
        offline.loadChanges()
        assertEquals(Loadable.NotCached, offline.changelog)
    }

    @Test
    fun `guide links open the tab and the guide, and the segment persists`() {
        val navigator = GuidesNavigator()
        assertTrue(navigator.open(requireNotNull(DeepLink.parse(DeepLink.Guide("shaman", "elemental").url()))))
        assertTrue(navigator.wantsTab)
        assertEquals(listOf(GuidesScreen.Guide(GuidesTestData.ELEMENTAL)), navigator.stack)
        assertTrue(navigator.open(DeepLink.Guides))
        assertTrue(navigator.stack.isEmpty())
        assertFalse(navigator.open(DeepLink.Plan))
        navigator.push(GuidesScreen.ClassPage("mage"))
        navigator.push(GuidesScreen.Guide(GuidesTestData.FIRE))
        navigator.pop()
        assertEquals(listOf(GuidesScreen.ClassPage("mage")), navigator.stack)

        val stored = MemoryValue()
        assertEquals(GuidesSegment.Guides, GuidesSegment.load(stored))
        GuidesSegment.save(stored, GuidesSegment.Changes)
        assertEquals(GuidesSegment.Changes, GuidesSegment.load(stored))
        GuidesSegment.save(stored, GuidesSegment.Tiers)
        assertEquals(GuidesSegment.Tiers, GuidesSegment.load(stored))
        // With guide_tiers off a saved Tier lists segment opens Guides.
        assertEquals(GuidesSegment.Guides, GuidesSegment.load(stored, tiersOn = false))
        assertEquals(listOf(GuidesSegment.Guides, GuidesSegment.Changes), GuidesSegment.visible(tiersOn = false))
    }
}
