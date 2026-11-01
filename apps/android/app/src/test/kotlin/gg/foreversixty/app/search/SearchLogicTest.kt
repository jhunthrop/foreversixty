package gg.foreversixty.app.search

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.CharacterSearch
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.search.logic.CharacterPhase
import gg.foreversixty.app.search.logic.MatchTier
import gg.foreversixty.app.search.logic.MemoryRecentStorage
import gg.foreversixty.app.search.logic.RecentSearches
import gg.foreversixty.app.search.logic.SearchContext
import gg.foreversixty.app.search.logic.SearchCopy
import gg.foreversixty.app.search.logic.SearchGroup
import gg.foreversixty.app.search.logic.SearchLogic
import gg.foreversixty.app.search.logic.SearchName
import gg.foreversixty.app.search.logic.SearchScope
import gg.foreversixty.app.search.logic.SearchStore
import gg.foreversixty.app.search.logic.SearchTarget
import java.io.IOException
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.test.StandardTestDispatcher
import kotlinx.coroutines.test.TestScope
import kotlinx.coroutines.test.advanceTimeBy
import kotlinx.coroutines.test.runCurrent
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The F12 search rules (spec F12.1, F12.2), mirrored in iOS's SearchLogicTests. */
@OptIn(ExperimentalCoroutinesApi::class)
@RunWith(AndroidJUnit4::class)
class SearchLogicTest {
    private val data = SearchTestData

    @Test fun `a name matches exact, then prefix, then word prefix, then contains`() {
        val name = SearchName("Elemental Fury")
        assertEquals(MatchTier.Exact, name.tier("elemental fury"))
        assertEquals(MatchTier.Prefix, name.tier("ele"))
        assertEquals(MatchTier.WordPrefix, name.tier("fu"))
        assertEquals(MatchTier.Contains, name.tier("ment"))
        assertNull(name.tier("frost"))
    }

    @Test fun `the query is trimmed, lowercased and its spaces collapsed`() {
        assertEquals("devout robe", SearchLogic.normalize("  Devout   ROBE "))
    }

    @Test fun `groups order by match tier before anything else`() {
        val index = data.prepared(
            talents = listOf(
                data.talent(1, "Improved Elements"), data.talent(2, "Elemental Fury"), data.talent(3, "Telemetry"), data.talent(4, "Ele"),
            ),
        )
        val names = SearchLogic.filter(index, "ele", "shaman").talents.map { it.name }
        assertEquals(listOf("Ele", "Elemental Fury", "Improved Elements", "Telemetry"), names)
    }

    @Test fun `items rank your class's usable items, then quality, then item level`() {
        val index = data.prepared(
            items = listOf(
                data.item(1, "Devout Robe", quality = 4, level = 63, classes = listOf("priest")),
                data.item(2, "Devout Bracers", quality = 3, level = 57, classes = listOf("shaman")),
                data.item(3, "Devout Mantle", quality = 3, level = 60, classes = listOf("shaman")),
                data.item(4, "Devout Crown", quality = 4, level = 62, classes = listOf("shaman", "priest")),
            ),
        )
        val ids = SearchLogic.filter(index, "devout", "shaman").items.map { it.id }
        assertEquals(listOf(4, 3, 2, 1), ids)
    }

    @Test fun `guides match on their spec and class, as the row reads`() {
        val index = data.prepared(guides = listOf(data.guide("shaman", "elemental", "Elemental"), data.guide("mage", "fire", "Fire")))
        assertEquals(listOf("elemental"), SearchLogic.filter(index, "shaman", null).guides.map { it.spec })
        assertEquals(listOf("fire"), SearchLogic.filter(index, "fire mage", null).guides.map { it.spec })
    }

    @Test fun `All caps each group at five in the order Guides, Talents, Items, Characters`() {
        val index = data.prepared(
            items = (1..7).map { data.item(it, "Ele Item $it") },
            talents = (1..2).map { data.talent(it, "Ele Talent $it") },
            guides = listOf(data.guide("shaman", "elemental", "Elemental")),
        )
        val local = SearchLogic.filter(index, "ele", "shaman")
        val characters = (1..3).map { data.character("Ele$it") }
        val sections = SearchLogic.sections(local, characters, SearchScope.All)
        assertEquals(listOf(SearchGroup.Guides, SearchGroup.Talents, SearchGroup.Items, SearchGroup.Characters), sections.map { it.group })
        val items = sections.first { it.group == SearchGroup.Items }
        assertEquals(5, items.rows.size)
        assertEquals(7, items.total)
        assertTrue(items.showsAll)
    }

    @Test fun `a scope shows only its group and lifts the cap`() {
        val index = data.prepared(items = (1..7).map { data.item(it, "Ele Item $it") }, talents = listOf(data.talent(1, "Ele")))
        val sections = SearchLogic.sections(SearchLogic.filter(index, "ele", null), emptyList(), SearchScope.Items)
        assertEquals(listOf(SearchGroup.Items), sections.map { it.group })
        assertEquals(7, sections.single().rows.size)
        assertTrue(!sections.single().showsAll)
    }

    @Test fun `characters are asked from two letters, under All and Characters only`() {
        assertTrue(!SearchLogic.asksCharacters("e", SearchScope.All))
        assertTrue(SearchLogic.asksCharacters("el", SearchScope.All))
        assertTrue(SearchLogic.asksCharacters("el", SearchScope.Characters))
        assertTrue(!SearchLogic.asksCharacters("el", SearchScope.Items))
    }

    @Test fun `characters in your region and ruleset come first, up to eight`() {
        val rows = (1..9).map { data.character("Eu$it", region = "eu") } + data.character("Us", region = "us")
        val ordered = SearchLogic.orderCharacters(rows, "us", "normal")
        assertEquals("Us", ordered.first().name)
        assertEquals(SearchLogic.CHARACTER_LIMIT, ordered.size)
    }

    @Test fun `typing waits 250 ms before characters are asked, and only the last query is asked`() = runTest {
        val asked = mutableListOf<String>()
        val store = store(this) { q ->
            asked += q
            ApiResult.Success(CharacterSearch(listOf(data.character("Elementalx"))), requestId = "")
        }
        store.loadIndex()
        store.type("el")
        advanceTimeBy(100)
        store.type("ele")
        advanceTimeBy(SearchLogic.CHARACTER_DEBOUNCE_MILLIS - 1)
        runCurrent()
        assertEquals(emptyList<String>(), asked)
        advanceTimeBy(2)
        runCurrent()
        assertEquals(listOf("ele"), asked)
        assertTrue(store.characterPhase is CharacterPhase.Loaded)
    }

    @Test fun `items filter at once without waiting`() = runTest {
        val store = store(this) { ApiResult.Success(CharacterSearch(), requestId = "") }
        store.loadIndex()
        store.type("spritec")
        runCurrent()
        assertEquals(listOf(SearchTestData.SPRITECASTER_CAPE), store.local.items.map { it.id })
    }

    @Test fun `no connection marks characters offline, and one letter asks nothing`() = runTest {
        var calls = 0
        val store = store(this) {
            calls++
            ApiResult.Failure(ApiError.Network(IOException("offline")))
        }
        store.loadIndex()
        store.type("x")
        advanceTimeBy(SearchLogic.CHARACTER_DEBOUNCE_MILLIS + 1)
        assertEquals(0, calls)
        store.type("xqz")
        advanceTimeBy(SearchLogic.CHARACTER_DEBOUNCE_MILLIS + 1)
        runCurrent()
        assertEquals(CharacterPhase.Offline, store.characterPhase)
    }

    @Test fun `Recent keeps the last 20, newest first, without repeats, and clears`() {
        val storage = MemoryRecentStorage()
        val recent = RecentSearches(storage)
        (1..25).forEach { recent.add(SearchTarget.Item(data.item(it, "Item $it"))) }
        recent.add(SearchTarget.Item(data.item(10, "Item 10")))
        assertEquals(RecentSearches.CAP, recent.entries.size)
        assertEquals("item:10", recent.entries.first().key)
        assertEquals(1, recent.entries.count { it.key == "item:10" })
        assertEquals(recent.entries, RecentSearches(storage).entries)
        recent.clear()
        assertTrue(RecentSearches(storage).entries.isEmpty())
    }

    @Test fun `row words come from the copy catalog`() {
        val copy = SearchCopy(resources())
        assertEquals("Back · item level 52 · requires level 47", copy.itemMeta(data.item(1, "Cape")))
        assertEquals("Mage · Level 58 · US Normal", copy.characterMeta(data.character("Elementalx")))
        assertEquals("Elemental Shaman · Elemental · 1 rank", copy.talentMeta(data.talent(1, "Elemental Focus", maxRank = 1)))
        assertEquals("Elemental Shaman · Elemental · 3 ranks", copy.talentMeta(data.talent(1, "Elemental Fury", maxRank = 3)))
        assertEquals("Guide · DPS", copy.guideMeta(data.guide("shaman", "elemental", "Elemental")))
    }

    private fun store(test: TestScope, search: suspend (String) -> ApiResult<CharacterSearch>): SearchStore {
        val fixture = data.fixtureIndex()
        val index = data.prepared(fixture.items, fixture.talents, fixture.guides)
        return SearchStore(
            index = { index },
            characters = { search(it) },
            context = SearchContext("shaman", "us", "normal"),
            scope = test.backgroundScope,
            compute = StandardTestDispatcher(test.testScheduler),
        )
    }
}
