package gg.foreversixty.app.search

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.search.logic.CharacterPhase
import gg.foreversixty.app.search.logic.IndexPhase
import gg.foreversixty.app.search.logic.ItemPageCopy
import gg.foreversixty.app.search.logic.ItemPageOwner
import gg.foreversixty.app.search.logic.ItemPagePhase
import gg.foreversixty.app.search.logic.ItemPageStore
import gg.foreversixty.app.search.logic.ItemUsesPhase
import gg.foreversixty.app.search.logic.LocalResults
import gg.foreversixty.app.search.logic.SearchLogic
import gg.foreversixty.app.search.logic.SearchScope
import gg.foreversixty.app.search.logic.SearchTarget
import gg.foreversixty.app.search.ui.ForCharacter
import gg.foreversixty.app.search.ui.ItemPage
import gg.foreversixty.app.search.ui.ItemPageActions
import gg.foreversixty.app.search.ui.ItemSection
import gg.foreversixty.app.search.ui.SearchActions
import gg.foreversixty.app.search.ui.SearchScreen
import gg.foreversixty.app.search.ui.SearchScreenState
import kotlinx.coroutines.runBlocking
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Every F12 board on a Pixel 9 (412 x 915 dp), at the default size and at font scale 2.0 (spec
 * §12). Search draws full screen over the tab bar, so no navigation bar is drawn. Record:
 * ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 * The PNGs double as the review captures in design/reviews/2026-10-31-mobile-wave3-search.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class SearchScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val data = SearchTestData
    private val index = data.realPrepared
    private val build = SearchTestData.BUILD

    private fun capture(name: String, large: Boolean = false, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    Box(Modifier.fillMaxSize().background(FsColor.bg)) { screen() }
                }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/search/$name.png", roborazziOptions = TOLERANT)
    }

    private fun search(name: String, state: SearchScreenState, large: Boolean = false) =
        capture(name, large) { SearchScreen(state, SearchActions(), focusField = false) }

    private fun state(
        query: String,
        scope: SearchScope = SearchScope.All,
        characters: CharacterPhase = CharacterPhase.Idle,
        recents: List<SearchTarget> = emptyList(),
    ): SearchScreenState {
        val local = if (query.isBlank()) LocalResults() else SearchLogic.filter(index, SearchLogic.normalize(query), "shaman")
        val rows = (characters as? CharacterPhase.Loaded)?.rows.orEmpty()
        return SearchScreenState(
            query = query,
            scope = scope,
            indexPhase = IndexPhase.Ready(index),
            sections = SearchLogic.sections(local, rows, scope),
            characterPhase = characters,
            recents = recents,
            forCharacter = ForCharacter(LIGHTNINGBOLLTZ, "Elemental Shaman guide"),
            build = build,
            ownKeys = setOf(CURRENT.key),
        )
    }

    /** Invented strangers (captions say so) for `ele`: no row is the account's. */
    private val strangers = CharacterPhase.Loaded(
        listOf(data.character("Elementalx", "mage", level = 58), data.character("Elemental", "druid", "eu", level = 44)),
    )

    /** `lightn`: the current character first (the `You` row), then two invented strangers. */
    private val lightn = CharacterPhase.Loaded(
        listOf(CURRENT, data.character("Lightningrod", "mage", level = 60), data.character("Lightnin", "warrior", "eu", level = 52)),
    )

    private fun cape() = requireNotNull(index.item(SearchTestData.SPRITECASTER_CAPE))

    // F12.1 The field and the empty state

    @Test fun empty() = search(
        "f12-1-empty-and",
        state("", recents = listOf(SearchTarget.Item(cape()), SearchTarget.Guide(data.realIndex().guides.first()))),
    )

    // F12.2 All results

    @Test fun all() = search("f12-2-all-and", state("ele", characters = strangers))

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun allFull() = search("f12-2-all-full-and", state("ele", characters = strangers))

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun allLarge() = search("f12-2-all-large-and", state("ele", characters = strangers), large = true)

    @Test fun characters() = search("f12-2-characters-and", state("lightn", SearchScope.Characters, lightn))

    // F12.3 Items

    @Test fun items() = search("f12-3-items-and", state("devout", SearchScope.Items))

    // F12.6 No results

    @Test fun noResults() = search("f12-6-none-and", state("xqz", characters = CharacterPhase.Loaded(emptyList())))

    @Test fun charactersAllScope() = search("f12-2-all-lightn-and", state("lightn", characters = lightn))

    @Test fun noResultsOffline() = search("f12-6-offline-and", state("xqz", characters = CharacterPhase.Offline))

    @Test fun indexFailed() = search("f12-index-failed-and", state("ele").copy(indexPhase = IndexPhase.Failed, sections = emptyList()))

    // F12.4 An item, F12.5 Best in slot for

    private fun owner(preset: String) = ItemPageOwner("shaman", BisSelection("shaman", "shaman-elemental", BisRange(60, preset), "horde"))

    private fun page(id: Int, owner: ItemPageOwner, source: SearchTestSource = SearchTestSource()): ItemPageStore = runBlocking {
        ItemPageStore(id, ItemLook(id), source, owner, data.specs(), resources()).also { it.load() }
    }

    private fun item(name: String, store: ItemPageStore, large: Boolean = false, startAt: ItemSection? = null) = capture(name, large) {
        ItemPage(store.phase, store.uses, build, resources().getString(R.string.itemPage_back), ItemPageActions(), startAt)
    }

    @Test fun pick() = item("f12-4-pick-and", page(SearchTestData.SPRITECASTER_CAPE, owner(BisRange.RAID)))

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun runnerUp() = item("f12-4-runner-up-and", page(SearchTestData.SATYRMANE_SASH, owner(BisRange.RAID)))

    /** A Horde vendor and reputation set piece no list places: no card, the set lines, two source kinds. */
    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun notOnList() = item("f12-4-not-listed-and", page(SearchTestData.DEFILERS_MAIL_PAULDRONS, owner(BisRange.RAID)))

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun usesFailed() = item(
        "f12-4-uses-failed-and",
        page(SearchTestData.SPRITECASTER_CAPE, owner(BisRange.RAID), SearchTestSource(uses = ApiResult.Failure(ApiError.Status(HTTP_UNAVAILABLE)))),
    )

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun pickLarge() = item("f12-4-pick-large-and", page(SearchTestData.SPRITECASTER_CAPE, owner(BisRange.RAID)), large = true)

    @Test fun bisFor() = item("f12-5-bis-for-and", page(SearchTestData.SPRITECASTER_CAPE, owner(BisRange.RAID)), startAt = ItemSection.BisFor)

    @Test fun unknown() = item("f12-unknown-and", page(999_999, owner(BisRange.RAID)))

    @Test fun loading() = capture("f12-4-loading-and") {
        ItemPage(ItemPagePhase.Loading(ItemPageCopy.look(cape())), ItemUsesPhase.Loading, build, "Search", ItemPageActions())
    }

    private companion object {
        const val LIGHTNINGBOLLTZ = "Lightningbolltz"
        const val HTTP_UNAVAILABLE = 503

        /** The signed-in account's current character, an Elemental Shaman: the one `You` row. */
        val CURRENT = SearchTestData.character(LIGHTNINGBOLLTZ, "shaman", level = 60)
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"

/** Pixel 9's width, tall enough for a whole item page. */
private const val PIXEL_9_TALL = "w412dp-h2000dp-normal-long-notround-any-420dpi-keyshidden-nonav"

/** Pixel 9's width, tall enough for every group and the largest text. */
private const val PIXEL_9_XL = "w412dp-h3200dp-normal-long-notround-any-420dpi-keyshidden-nonav"
