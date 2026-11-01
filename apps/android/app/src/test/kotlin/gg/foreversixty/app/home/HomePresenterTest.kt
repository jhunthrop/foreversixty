package gg.foreversixty.app.home

import android.content.res.Resources
import gg.foreversixty.app.bis.BisRange
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemNumbers
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.FinishedSim
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.FixtureEngine
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.TokenProvider
import gg.foreversixty.app.network.ApiResult
import java.io.IOException
import java.time.ZoneOffset
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/**
 * F2's logic, one test list shared with iOS (`HomePresenterTests.swift`): Home renders the
 * server's answer in the server's order and formats it; it never ranks or computes a gain.
 */
@RunWith(AndroidJUnit4::class)
class HomePresenterTest {
    private val res: Resources = ApplicationProvider.getApplicationContext<android.content.Context>().resources
    private val presenter = HomeFixtures.presenter(res)

    private fun upgrades(model: HomeAnswerModel) = model.body as HomeAnswerBody.Upgrades

    @Test
    fun `rows read measured gains largest first, then not modelled, then empty, and the card shows five`() {
        val home = HomeFixtures.home()
        val ordered = HomePresenter.presentationOrder(home.upgrades).map { it.slot }
        assertEquals(
            listOf("main_hand", "ranged", "legs", "feet", "chest", "head", "waist", "hands", "back", "finger2", "wrist", "trinket2"),
            ordered,
        )
        val body = upgrades(presenter.model(home))
        assertEquals(ordered.take(5), body.rows.map { it.slot })
        assertEquals("7 more upgrades · all 17 slots", body.footer)
    }

    @Test
    fun `the hero shows two figures, the note with the unmodelled count, the sync age`() {
        val model = presenter.model(HomeFixtures.home())
        assertEquals("Orc Beast Mastery Hunter · Level 60 · Olympus XXVII", model.heroMeta)
        assertEquals("12 upgrades", model.count)
        assertEquals("+102.5", model.gainFigure)
        assertEquals("DPS", model.gainUnit)
        assertEquals("Best in slot for 60, together · 2 gains not modelled yet", model.note)
        assertEquals("Synced 2 h ago from the addon", model.synced)
        assertEquals("60 · raid", model.bandHead)
    }

    @Test
    fun `Home's band is the site rule over the character's level, the list BiS opens on`() {
        val home = HomeFixtures.home()
        val cases = listOf(Triple(49, 40, "40 to 49"), Triple(50, 50, "50 to 59"), Triple(59, 50, "50 to 59"), Triple(60, 60, "60 · raid"))
        for ((level, band, head) in cases) {
            val character = HomeFixtures.velaren.copy(level = level)
            val range = HomePresenter.range(home, character)
            assertEquals("level $level", BisRange.forLevel(level), range)
            assertEquals("level $level", band, range?.band)
            assertEquals("level $level", head, HomeFixtures.presenter(res, character = character).model(home).bandHead)
            val tooltip = HomePresenter.tooltipRequest(ItemLook(1), home, character, res)
            assertEquals("level $level", Triple(range?.band, range?.preset, head), Triple(tooltip.band, tooltip.preset, tooltip.bandLabel))
        }
    }

    @Test
    fun `with no level known, the server's band goes through the same rule`() {
        assertEquals(BisRange.forLevel(60), HomePresenter.range(HomeFixtures.home(), HomeFixtures.velaren.copy(level = null)))
    }

    @Test
    fun `a measured gain is the server's figure, the source the server's formatted line`() {
        val home = HomeFixtures.home()
        val rows = upgrades(presenter.model(home)).rows
        assertEquals(HomeUpgradeRow.Gain.Measured("+21.2", "DPS"), rows[0].gain)
        assertEquals("Main hand · Crafted: Blacksmithing", rows[0].meta)
        assertEquals("Gain not modelled yet", HomePresenter.presentationOrder(home.upgrades).last().gainWords)
    }

    @Test
    fun `a negative gain keeps its true minus`() {
        assertEquals("−0.8", ItemNumbers.signedTenths(-0.8))
    }

    @Test
    fun `what a row replaces - the worn item, an empty slot, or an item our data does not know`() {
        val upgrades = HomeFixtures.home().upgrades
        assertEquals("Replaces Dark Iron Rifle, ilvl 55", HomePresenter.replaces(upgrades.first { it.slot == "ranged" }.worn, res))
        assertEquals("Empty slot", HomePresenter.replaces(null, res))
        assertEquals("Not in our item data yet", HomePresenter.replaces(upgrades.first { it.slot == "wrist" }.worn, res))
    }

    @Test
    fun `every slot matching says so in place of the figures, with the Already line`() {
        val model = presenter.model(HomeFixtures.home("character.home.empty"))
        assertEquals(HomeAnswerBody.AllMatch("Every slot matches the 60 list."), model.body)
        assertNull(model.count)
        assertTrue(model.already.orEmpty().startsWith("Already best in slot: Head, Neck, Shoulder"))
    }

    @Test
    fun `no gear export shows the sentence, Get the addon, the list button and the note`() {
        val model = presenter.model(HomeFixtures.home("character.home.no-gear"))
        val body = model.body as HomeAnswerBody.Message
        assertEquals("Add the addon to see which slots you can upgrade.", body.text)
        assertEquals(listOf(HomeAction.Kind.GetAddon, HomeAction.Kind.OpenList), body.buttons.map { it.kind })
        assertEquals(listOf("Get the addon", "Beast Mastery 60 list"), body.buttons.map { it.title })
        assertTrue(body.note.orEmpty().startsWith("The addon sends your gear"))
        assertEquals(
            "Not available yet: no talent export for this character. Open the addon once to send it.",
            model.next.first().line,
        )
        assertEquals("No sim yet.", model.next.last().line)
    }

    @Test
    fun `under 20, no spec and no list each have their sentence`() {
        mapOf(
            "under_20" to "Best in slot starts at 20.",
            "no_spec" to "Set a spec in the planner to see this.",
            "no_list" to "No best in slot list published yet for this spec and level range.",
        ).forEach { (state, sentence) ->
            assertEquals(sentence, (presenter.model(HomeFixtures.home(state = state)).body as HomeAnswerBody.Message).text)
        }
    }

    @Test
    fun `talents show the differ count and its tag`() {
        val row = presenter.model(HomeFixtures.home()).next.first()
        assertEquals("1 of your 51 points differ from the 60 guide build", row.line)
        assertEquals(HomeNextRow.Tag.Differs, row.tag)
    }

    @Test
    fun `sims show the newest finished run for this spec`() {
        val sim = FinishedSim("s1", "hunter-beast-mastery", "t", 812.4, "2026-10-29T23:15:00Z")
        val row = HomeFixtures.presenter(res, listOf(sim)).model(HomeFixtures.home()).next.last()
        assertEquals("Your Beast Mastery sim finished: 812 DPS", row.line)
    }

    @Test
    fun `the offline stamp reads today, this week, older`() {
        val time = HomeTime(HomeFixtures.now, ZoneOffset.UTC)
        val hour = 3600L
        assertEquals("9:15 pm", time.stamp(HomeFixtures.now.minusSeconds(2 * hour)))
        assertEquals("Tue 11:15 pm", time.stamp(HomeFixtures.now.minusSeconds(48 * hour)))
        assertEquals("29 Sep, 11:15 pm", time.stamp(HomeFixtures.now.minusSeconds(30 * 24 * hour)))
    }
}

@RunWith(AndroidJUnit4::class)
class HomeStoreTest {
    private fun store(result: ApiResult<CharacterHome>, cache: HomeCache) =
        HomeStore(fetch = { result }, cache = cache, now = { HomeFixtures.now })

    @Test
    fun `a first load with nothing cached shows the answer and caches it`() = runTest {
        val home = HomeFixtures.home()
        val cache = MemoryHomeCache()
        val store = store(ApiResult.Success(home, "r"), cache)
        store.load(HomeFixtures.velaren)
        assertEquals(HomePhase.Loaded(home), store.phase.value)
        assertEquals(home, cache.load(HomeFixtures.velaren.key)?.home)
    }

    @Test
    fun `no connection with a cached answer keeps it under the offline strip`() = runTest {
        val home = HomeFixtures.home()
        val at = HomeFixtures.now.minusSeconds(600)
        val cache = MemoryHomeCache(mapOf(HomeFixtures.velaren.key to HomeCacheEntry(home, at.toEpochMilli())))
        val store = store(ApiResult.Failure(ApiError.Network(IOException("offline"))), cache)
        store.load(HomeFixtures.velaren)
        assertEquals(HomePhase.Loaded(home, offlineSince = at), store.phase.value)
    }

    @Test
    fun `no connection and nothing cached is the error state`() = runTest {
        val store = store(ApiResult.Failure(ApiError.Network(IOException("offline"))), MemoryHomeCache())
        store.load(HomeFixtures.velaren)
        assertEquals(HomePhase.Failed(null), store.phase.value)
    }

    @Test
    fun `a server error is the error state, never the offline strip`() = runTest {
        val home = HomeFixtures.home()
        val cache = MemoryHomeCache(mapOf(HomeFixtures.velaren.key to HomeCacheEntry(home, HomeFixtures.now.toEpochMilli())))
        val store = store(ApiResult.Failure(ApiError.Server(500, "internal", "x", emptyMap(), "r")), cache)
        store.load(HomeFixtures.velaren)
        assertEquals(HomePhase.Failed(home), store.phase.value)
    }

    @Test
    fun `fixture mode answers the home route and the tooltip's static files`() = runTest {
        val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))
        val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
        val client = ApiClient(ForeverHttpClient.create(config, NoToken, FixtureEngine.create(source)))
        val store = HomeStore.live(client, MemoryHomeCache())
        store.load(HomeFixtures.velaren)
        val home = (store.phase.value as HomePhase.Loaded).home
        assertEquals(12, home.upgradeCount)
        val catalog = ItemCatalog(StaticDataStore(StaticDataClient(FixtureEngine.create(source))))
        val sources = catalog.sources(HomePresenter.tooltipRequest(ItemLook(252605), home, HomeFixtures.velaren, ApplicationProvider.getApplicationContext<android.content.Context>().resources))
        assertTrue(sources?.items?.containsKey(252605) == true)
        assertTrue(sources?.bis != null && sources.sets.isNotEmpty())
    }

    private object NoToken : TokenProvider {
        override suspend fun bearerToken(): String? = null
        override suspend fun discardBearerToken() = Unit
    }

    @Test
    fun `the cache entry survives a round trip through its file`() {
        val entry = HomeCacheEntry(HomeFixtures.home(), HomeFixtures.now.toEpochMilli())
        assertEquals(entry, HomeCacheEntry.decode(entry.encode()))
    }
}

@RunWith(AndroidJUnit4::class)
class ItemTooltipTest {
    private val res: Resources = ApplicationProvider.getApplicationContext<android.content.Context>().resources
    private val sources = HomeFixtures.tooltipSources()

    private fun request(id: Int, wornId: Int? = null, sourceLine: String? = null) = HomePresenter.tooltipRequest(
        ItemLook(id, "Fallback", 4, null, 60), HomeFixtures.home(), HomeFixtures.velaren, res, sourceLine, wornId,
    )

    @Test
    fun `a pick shows binds, armor type, stats, levels, its source, the next best and the footer`() {
        val model = ItemTooltipModel.build(request(252605), sources, res)
        assertEquals("Epic · Head", model.headerMeta)
        assertEquals("Binds when equipped", model.binds)
        assertEquals("Head", model.slotName)
        assertEquals("Mail", model.typeName)
        assertTrue(model.armor.orEmpty().endsWith(" Armor"))
        assertTrue(model.stats.all { it.startsWith("+") })
        assertEquals("Item Level 61", model.itemLevel)
        assertEquals("Crafted: Leatherworking", model.source)
        assertEquals(2, model.alternatives.size)
        assertEquals("Best in slot for Head · Beast Mastery 60 · raid", model.pickFooter)
    }

    @Test
    fun `a vendor item adds the standing it requires to the site's line`() {
        assertEquals("Vendor: Friz Frazzlespark · requires Friendly", ItemTooltipModel.build(request(276202), sources, res).source)
    }

    @Test
    fun `the worn runner-up leads with You wear this`() {
        val model = ItemTooltipModel.build(request(9533, wornId = 11862), sources, res)
        assertTrue(model.alternatives.map { it.meta }.contains("You wear this · −2.3 DPS vs the pick"))
        assertEquals("2% from Ancient Treasure", model.alternatives.first().source)
    }

    @Test
    fun `a weapon reads Two-Hand and its type, as the client prints it`() {
        val model = ItemTooltipModel.build(request(250619), sources, res)
        assertEquals("Two-Hand", model.slotName)
        assertEquals("Polearm", model.typeName)
        assertTrue(model.weaponDamage.orEmpty().endsWith(" Damage"))
        assertTrue(model.weaponDps.orEmpty().endsWith("damage per second)"))
    }

    @Test
    fun `an item the build does not list keeps the caller's look and says so`() {
        val model = ItemTooltipModel.build(request(999999, sourceLine = "Somewhere"), sources, res)
        assertFalse(model.known)
        assertEquals("Fallback", model.look.name)
        assertEquals("Somewhere", model.source)
    }

    @Test
    fun `Your gear lists 17 rows in frame order, and the off hand names the two-hander`() {
        val home = HomeFixtures.home()
        val blocker = GearModel.offHandBlocker(home, emptyMap(), sources.items, ItemTooltipModel.band(request(0), sources.bis))
        assertEquals("Legionite Glaive", blocker)
        val rows = GearModel.rows(home, emptyMap(), sources.items, res, blocker)
        assertEquals(ItemVocabulary.frameSlots, rows.map { it.slot })
        assertEquals(17, rows.size)
        assertEquals("none: Legionite Glaive is two-handed.", rows.first { it.slot == "off_hand" }.wornName)
        val trinket = rows.first { it.slot == "trinket2" }
        assertNull(trinket.worn)
        assertEquals("Empty", trinket.wornName)
        assertEquals(249469, trinket.pick?.id)
        assertTrue(rows.first { it.slot == "shoulder" }.isBestInSlot)
        assertEquals("Wrist · Item 999999", rows.first { it.slot == "wrist" }.wornName)
        assertEquals("Mourning Handguards", rows.first { it.slot == "hands" }.wornName)
    }
}
