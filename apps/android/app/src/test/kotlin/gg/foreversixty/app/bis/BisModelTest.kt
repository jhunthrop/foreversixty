package gg.foreversixty.app.bis

import android.content.Context
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.ItemsFile
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith
import java.io.IOException
import java.time.ZoneOffset

/** A data source answering from the contract fixtures, with switches for offline and errors. */
private class FakeBisSource(
    var bis: ApiResult<CachedFile<BisFile>>,
    var character: AccountCharacter?,
) : BisDataSource {
    private val items = BisTestData.items()
    private val classes = BisTestData.classes()

    override suspend fun activeBuild() = BisTestData.BUILD
    override suspend fun bis(build: String, spec: String) = bis
    override suspend fun items(build: String, classSlug: String): ApiResult<CachedFile<ItemsFile>> =
        ApiResult.Success(CachedFile(items, BisTestData.NOW, isFromCache = false), requestId = "")
    override suspend fun classes(build: String): ApiResult<CachedFile<List<ClassInfo>>> =
        ApiResult.Success(CachedFile(classes, BisTestData.NOW, isFromCache = false), requestId = "")
    override fun cachedBis(build: String, spec: String): CachedFile<BisFile>? = null
    override fun cachedItems(build: String, classSlug: String): CachedFile<ItemsFile>? = null
    override fun cachedClasses(build: String): CachedFile<List<ClassInfo>>? = null
    override suspend fun currentCharacter() = character
}

private class MemoryPreferences : BisPreferences {
    override var lastSelection: BisSelection? = null
    override var hasOpenedPicker: Boolean = false
}

/** The same list as iOS's BisModelTests. */
@RunWith(AndroidJUnit4::class)
class BisModelTest {
    private val res = ApplicationProvider.getApplicationContext<Context>().resources
    private val file = BisTestData.bis()
    private val main = requireNotNull(BisTestData.me().mainCharacter)
    private val catalog = BisTestData.catalog()

    private fun fresh(): ApiResult<CachedFile<BisFile>> = ApiResult.Success(CachedFile(file, BisTestData.NOW, isFromCache = false), "")

    private fun model(source: BisDataSource, preferences: BisPreferences = MemoryPreferences()) =
        BisModel(source, catalog, preferences, res)

    @Test
    fun `signed in shows the character's list, marked Yours`() = runTest {
        val model = model(FakeBisSource(fresh(), main))
        model.start()
        val state = model.state.value
        assertTrue(state.context.isYours)
        assertEquals(BisRange(50, "bare"), state.context.selection.range)
        assertFalse(state.isPickerOpen)
        val phase = state.phase as BisPhase.Loaded
        assertEquals(null, phase.offlineSinceMillis)
        assertEquals(17, phase.list.rows.size)
    }

    @Test
    fun `signed out on the first visit preselects Warrior and opens the picker once`() = runTest {
        val preferences = MemoryPreferences()
        val source = FakeBisSource(fresh(), null)
        val model = model(source, preferences)
        model.start()
        assertTrue(model.state.value.isPickerOpen)
        assertEquals("warrior", model.state.value.context.selection.classSlug)
        assertFalse(model.state.value.context.isYours)
        val again = model(source, preferences)
        again.start()
        assertFalse(again.state.value.isPickerOpen)
    }

    @Test
    fun `a choice applies at once and Back returns to the character's list`() = runTest {
        val model = model(FakeBisSource(fresh(), main))
        model.start()
        model.select(model.state.value.context.selection.copy(range = BisRange(60, "raid")))
        assertFalse(model.state.value.context.isYours)
        val back = requireNotNull(model.state.value.context.backTarget)
        model.select(back.second)
        assertTrue(model.state.value.context.isYours)
    }

    @Test
    fun `offline with a saved copy shows the list stamped with the save time`() = runTest {
        val saved = 1_793_300_000_000L
        val model = model(FakeBisSource(ApiResult.Success(CachedFile(file, saved, isFromCache = true), ""), main))
        model.start()
        assertEquals(saved, (model.state.value.phase as BisPhase.Loaded).offlineSinceMillis)
    }

    @Test
    fun `offline and never fetched says the list is not on the phone`() = runTest {
        val model = model(FakeBisSource(ApiResult.Failure(ApiError.Network(IOException("offline"))), main))
        model.start()
        assertEquals(BisPhase.NotCached, model.state.value.phase)
    }

    @Test
    fun `a CDN error is the error state, and Try again recovers`() = runTest {
        val source = FakeBisSource(ApiResult.Failure(ApiError.Status(500)), main)
        val model = model(source)
        model.start()
        assertEquals(BisPhase.Failed, model.state.value.phase)
        source.bis = fresh()
        model.refresh()
        assertTrue(model.state.value.phase is BisPhase.Loaded)
    }

    @Test
    fun `a range the file does not publish says so`() = runTest {
        val model = model(FakeBisSource(fresh(), main))
        model.start()
        model.select(model.state.value.context.selection.copy(faction = "alliance"))
        assertEquals(BisPhase.NoList(null), model.state.value.phase)
    }

    @Test
    fun `the offline stamp reads today, this week, else the date`() {
        val day = 86_400_000L
        val stamp = { ago: Long -> OfflineStamp.format(BisTestData.NOW - ago, BisTestData.NOW, ZoneOffset.UTC) }
        assertEquals("9:41 pm", stamp(0))
        assertEquals("Thu 9:41 pm", stamp(day))
        assertEquals("Sat 9:41 pm", stamp(6 * day))
        assertEquals("22 Jul, 9:41 pm", stamp(100 * day))
    }
}
