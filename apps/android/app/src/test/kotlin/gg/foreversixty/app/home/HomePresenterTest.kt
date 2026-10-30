package gg.foreversixty.app.home

import android.content.res.Resources
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
    fun `the rows keep the server's order and the card shows the first five`() {
        val home = HomeFixtures.home()
        val body = upgrades(presenter.model(home))
        assertEquals(home.upgrades.take(5).map { it.slot }, body.rows.map { it.slot })
        assertEquals("6 more upgrades · all 17 slots", body.footer)
    }

    @Test
    fun `the hero shows two figures, the note with the unmodelled count, the sync age`() {
        val model = presenter.model(HomeFixtures.home())
        assertEquals("Orc Beast Mastery Hunter · Level 60 · Olympus XXVII", model.heroMeta)
        assertEquals("11 upgrades", model.count)
        assertEquals("+76.2", model.gainFigure)
        assertEquals("DPS", model.gainUnit)
        assertEquals("Best in slot for 60, together · 6 gains not modelled yet", model.note)
        assertEquals("Synced 2 h ago from the addon", model.synced)
        assertEquals("60 · raid", model.bandHead)
    }

    @Test
    fun `a measured gain is the server's figure, an unmodelled one the server's words`() {
        val rows = upgrades(presenter.model(HomeFixtures.home())).rows
        assertEquals(HomeUpgradeRow.Gain.Measured("+11.1", "DPS"), rows[0].gain)
        assertEquals(HomeUpgradeRow.Gain.Words("Gain not modelled yet"), rows[1].gain)
        assertEquals("Hands · Creeg Bothunk", rows[0].meta)
    }

    @Test
    fun `a negative gain keeps its true minus`() {
        assertEquals("−0.8", ItemNumbers.signedTenths(-0.8))
        assertEquals("−5.3", ItemNumbers.delta(-5.31))
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
        assertEquals("1 of your 51 points differ from the 60 build", row.line)
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
        assertEquals(11, home.upgradeCount)
        val catalog = ItemCatalog(StaticDataClient(FixtureEngine.create(source)))
        val sources = catalog.sources(HomePresenter.tooltipRequest(ItemLook(226887), home, HomeFixtures.velaren))
        assertTrue(sources?.items?.containsKey(226887) == true)
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
        ItemLook(id, "Fallback", 4, null, 60), HomeFixtures.home(), HomeFixtures.velaren, sourceLine, wornId,
    )

    @Test
    fun `a pick shows binds, armor type, stats, levels, its source, the next best and the footer`() {
        val model = ItemTooltipModel.build(request(226887), sources, res)
        assertEquals("Epic · Head", model.headerMeta)
        assertEquals("Binds when picked up", model.binds)
        assertEquals("Head", model.slotName)
        assertEquals("Mail", model.typeName)
        assertTrue(model.armor.orEmpty().endsWith(" Armor"))
        assertTrue(model.stats.all { it.startsWith("+") })
        assertEquals("Item Level 60", model.itemLevel)
        assertEquals("Saving the Best for Last", model.source)
        assertEquals(2, model.alternatives.size)
        assertEquals("Best in slot for Head · Beast Mastery 60 · Raid-ready, Phase 1", model.pickFooter)
    }

    @Test
    fun `a rep reward names its standing, a vendor adds what it requires`() {
        assertEquals("The Defilers (Exalted)", ItemTooltipModel.build(request(20068), sources, res).source)
        assertEquals("Creeg Bothunk · requires Exalted", ItemTooltipModel.build(request(272099), sources, res).source)
    }

    @Test
    fun `the worn runner-up leads with You wear this`() {
        val model = ItemTooltipModel.build(request(234202, wornId = 9533), sources, res)
        assertTrue(model.alternatives.map { it.meta }.contains("You wear this · −8.2 DPS vs the pick"))
    }

    @Test
    fun `a two-handed weapon names its handedness and its damage`() {
        val model = ItemTooltipModel.build(request(250619), sources, res)
        assertEquals("Two-Handed Weapon", model.typeName)
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
    fun `Your gear lists 17 rows in character-frame order, worn then the pick`() {
        val rows = GearModel.rows(HomeFixtures.home(), mapOf("off_hand" to 16004), sources.items, res)
        assertEquals(ItemVocabulary.frameSlots, rows.map { it.slot })
        assertEquals(17, rows.size)
        val hands = rows.first { it.slot == "hands" }
        assertNull(hands.worn)
        assertEquals("Empty", hands.wornName)
        assertEquals(272099, hands.pick?.id)
        assertTrue(rows.first { it.slot == "shoulder" }.isBestInSlot)
        assertEquals("Not in our item data yet", rows.first { it.slot == "wrist" }.wornName)
        assertEquals("Dark Iron Rifle", rows.first { it.slot == "off_hand" }.wornName)
    }
}
