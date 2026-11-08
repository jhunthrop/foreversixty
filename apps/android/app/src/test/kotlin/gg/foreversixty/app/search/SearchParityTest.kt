package gg.foreversixty.app.search

import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.model.SearchItem
import gg.foreversixty.app.search.logic.CharacterPhase
import gg.foreversixty.app.search.logic.IndexPhase
import gg.foreversixty.app.search.logic.SearchGroup
import gg.foreversixty.app.search.logic.SearchScope
import gg.foreversixty.app.search.logic.SearchSection
import gg.foreversixty.app.search.logic.SearchTarget
import gg.foreversixty.app.search.ui.SearchActions
import gg.foreversixty.app.search.ui.SearchScreen
import gg.foreversixty.app.search.ui.SearchScreenState
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** Search against its mocks (design/mocks/ws_app.py `search`): the empty field, and `Thal`. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class SearchParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun search() = ParityCheck.state(compose, "search") { SearchScreen(SearchFixtures.empty, SearchActions(), focusField = false) }

    @Test fun searchResults() = ParityCheck.state(compose, "search-results") { SearchScreen(SearchFixtures.thal, SearchActions(), focusField = false) }

    @Test fun searchResultsAt2() =
        ParityCheck.state(compose, "search-results", fontScale = 2f) { SearchScreen(SearchFixtures.thal, SearchActions(), focusField = false) }
}

/** The mock's searches: nothing typed, then `Thal` finding two warriors and a pair of epic shoulders. */
object SearchFixtures {
    private val index = IndexPhase.Ready(SearchTestData.prepared())

    private fun warrior(name: String) = CharacterSearchRow("us-normal-${name.lowercase()}", "us", "normal", name, "warrior", 60)

    private val characters = listOf(warrior("Thandrel"), warrior("Thalrin"))
    private val pauldrons = SearchItem(19_375, "Thalassian Pauldrons", "inv_shoulder_15", 4, "shoulder", 71, 60)

    val empty = SearchScreenState(
        query = "",
        scope = SearchScope.All,
        indexPhase = index,
        sections = emptyList(),
        characterPhase = CharacterPhase.Idle,
        recents = emptyList(),
        forCharacter = null,
        build = SearchTestData.BUILD,
        rankingsInApp = true,
        guidesInApp = true,
    )

    val thal = empty.copy(
        query = "Thal",
        characterPhase = CharacterPhase.Loaded(characters),
        sections = listOf(
            SearchSection(SearchGroup.Characters, characters.map { SearchTarget.Character(it) }, characters.size),
            SearchSection(SearchGroup.Items, listOf(SearchTarget.Item(pauldrons)), 1),
        ),
    )
}
