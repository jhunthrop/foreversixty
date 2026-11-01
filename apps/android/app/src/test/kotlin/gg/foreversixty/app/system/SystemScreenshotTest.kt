package gg.foreversixty.app.system

import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.talents.TalentBuildKind
import gg.foreversixty.app.talents.TalentsPhase
import gg.foreversixty.app.talents.TalentsPage
import gg.foreversixty.app.talents.TalentTestData
import gg.foreversixty.app.planner.ui.SharedPageState
import gg.foreversixty.app.planner.ui.SharedBuildPage
import gg.foreversixty.app.planner.ui.SharedActions
import gg.foreversixty.app.planner.logic.SharedCompare
import gg.foreversixty.app.planner.logic.SharedBuildPhase
import gg.foreversixty.app.planner.logic.SharedBuildView
import gg.foreversixty.app.rankings.ui.CharacterScreenState
import gg.foreversixty.app.rankings.ui.CharacterContent
import gg.foreversixty.app.rankings.ui.CharacterActions
import gg.foreversixty.app.rankings.logic.WatchState
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.CharacterPhase as CharacterPagePhase
import gg.foreversixty.app.rankings.RankingsTestData
import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.model.WatchingList
import gg.foreversixty.app.raids.RaidFixture
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.account.CharacterSelector
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.NewDataStrip
import gg.foreversixty.app.guides.GuidesRootLayout
import gg.foreversixty.app.guides.GuidesTestData
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesSegment
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.guides.ui.GuidesHome
import gg.foreversixty.app.guides.ui.GuidesHomeActions
import gg.foreversixty.app.guides.ui.GuidesHomeState
import gg.foreversixty.app.home.HomeAnswer
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.home.HomePhase
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.search.SearchTestData
import gg.foreversixty.app.search.logic.CharacterPhase
import gg.foreversixty.app.search.logic.IndexPhase
import gg.foreversixty.app.search.logic.SearchScope
import gg.foreversixty.app.search.ui.ForCharacter
import gg.foreversixty.app.search.ui.RootSearchBar
import gg.foreversixty.app.search.ui.SearchActions
import gg.foreversixty.app.search.ui.SearchScreen
import gg.foreversixty.app.search.ui.SearchScreenState
import gg.foreversixty.app.system.offline.IconCacheSize
import gg.foreversixty.app.system.offline.OfflineData
import gg.foreversixty.app.system.offline.OfflineDataActions
import gg.foreversixty.app.system.offline.OfflineDataScreen
import gg.foreversixty.app.system.offline.OfflineDataState
import gg.foreversixty.app.system.offline.OfflineReport
import gg.foreversixty.app.network.SavedStaticFile
import java.io.File
import java.time.Instant
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Lane F15's in-app surfaces on a Pixel 9 (412 x 915 dp) at the default size and font scale 2.0:
 * Offline data (F15.4), the three strip states (F15.5), the Guides root bar's search button and
 * Search with Watching and Rankings. Record: ./gradlew :app:recordRoborazziDebug. Each PNG is also
 * copied to design/reviews/2026-11-01-mobile-integration/build/and for the design review.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class SystemScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    /** Thursday 5 Nov 2026, 9:41 pm in Chicago (the spec's one story). */
    private val now: Instant = Instant.parse("2026-11-06T03:41:00Z")
    private val saved: Long = now.minusSeconds(SAVED_AGO).toEpochMilli()
    private val build = "1.60.1.70658"

    private fun capture(name: String, large: Boolean, tab: ForeverTab? = ForeverTab.You, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(LARGE_SCALE)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    if (tab == null) {
                        Box(Modifier.fillMaxSize().background(FsColor.bg)) { screen() }
                    } else {
                        ForeverTabScaffold(selected = tab, onSelect = {}, tabs = ForeverTab.visible(guides = true)) {
                            Box(Modifier.fillMaxSize().background(FsColor.bg)) { screen() }
                        }
                    }
                }
            }
        }
        val file = "$name${if (large) "-large" else ""}-and"
        val path = "src/test/screenshots/system/$file.png"
        compose.onRoot().captureRoboImage(path, roborazziOptions = TOLERANT)
        File(path).takeIf { it.exists() }?.copyTo(File(File(REVIEW_DIR).apply { mkdirs() }, "$file.png"), overwrite = true)
    }

    /** The guides Offline data counts: Elemental Shaman and Fury Warrior, 2.4 MB together. */
    private val savedGuides = mapOf(GuidesTestData.ELEMENTAL to ELEMENTAL_BYTES, "warrior-fury" to FURY_BYTES)

    // F15.4 Offline data

    private val report: OfflineReport = OfflineData.report(
        listOf(
            SavedStaticFile("/data/$build/guides/shaman-elemental.json", ELEMENTAL_BYTES, saved),
            SavedStaticFile("/data/$build/guides/warrior-fury.json", FURY_BYTES, saved),
            SavedStaticFile("/data/$build/bis/shaman-elemental.json", 2_100_000, saved),
            SavedStaticFile("/data/$build/items/shaman.json", 1_000_000, saved),
            SavedStaticFile("/data/$build/talents/shaman.json", 600_000, saved),
            SavedStaticFile("/data/$build/search.json", 400_000, saved),
        ),
        IconCacheSize(41_000_000, Instant.ofEpochMilli(saved)),
    )

    private fun offlineData(name: String, large: Boolean, report: OfflineReport, confirming: Boolean = false, online: Boolean = true) =
        capture(name, large) {
            OfflineDataScreen(
                OfflineDataState(report, keepClassesReady = true, wifiOnly = true, online = online, updating = false, note = null, checkedAt = now, confirmingRemove = confirming),
                now,
                OfflineDataActions(),
            )
        }

    @Test fun offlineData() = offlineData("f15-offline-data", false, report)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun offlineDataLarge() = offlineData("f15-offline-data", true, report)

    @Test fun offlineDataEmpty() = offlineData("f15-offline-data-empty", false, OfflineReport(emptyList()))

    @Test fun offlineDataEmptyLarge() = offlineData("f15-offline-data-empty", true, OfflineReport(emptyList()))

    @Test fun offlineDataRemove() = offlineData("f15-offline-data-remove", false, report, confirming = true)

    @Test fun offlineDataRemoveLarge() = offlineData("f15-offline-data-remove", true, report, confirming = true)

    @Test fun offlineDataNoConnection() = offlineData("f15-offline-data-offline", false, report, online = false)

    // F15.5 the strip's three states, on Home and a guide

    private fun home(name: String, large: Boolean, stale: Boolean) = capture(name, large, ForeverTab.Home) {
        val presenter = HomeFixtures.presenter(LocalResources.current)
        val at = Instant.parse("2026-11-06T03:41:00Z")
        Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            HomeAnswer(HomePhase.Loaded(HomeFixtures.home(), at, stale = stale), presenter, HomeFixtures.BUILD, "9:41 pm")
        }
    }

    @Test fun stripOffline() = home("f15-strip-offline", false, stale = false)

    @Test fun stripOfflineLarge() = home("f15-strip-offline", true, stale = false)

    @Test fun stripStale() = home("f15-strip-stale", false, stale = true)

    @Test fun stripStaleLarge() = home("f15-strip-stale", true, stale = true)

    private fun newData(large: Boolean) = capture("f15-strip-newdata", large, ForeverTab.Guides) {
        val index = GuidesTestData.index()
        Column {
            NewDataStrip("New game data is ready.", "Update", {})
            GuidesRootLayout(GuidesSegment.visible(tiersOn = true), GuidesSegment.Guides, {}) {
                GuidesHome(GuidesHomeState(Loadable.Loaded(index, null), savedGuides, Guides.yourGuide(index, GuidesTestData.lightningbolltz)), GuidesHomeActions())
            }
        }
    }

    @Test fun stripNewData() = newData(false)

    @Test fun stripNewDataLarge() = newData(true)

    // §3.2 the Guides root bar's search button

    private fun guidesBar(large: Boolean) = capture("f15-guides-search", large, ForeverTab.Guides) {
        val index = GuidesTestData.index()
        GuidesRootLayout(GuidesSegment.visible(tiersOn = true), GuidesSegment.Guides, {}, selector = {
            RootSearchBar({}) { Row(Modifier.padding(horizontal = 16.dp, vertical = 8.dp)) { CharacterSelector(GuidesTestData.lightningbolltz, open = false) {} } }
        }) {
            GuidesHome(GuidesHomeState(Loadable.Loaded(index, null), savedGuides, Guides.yourGuide(index, GuidesTestData.lightningbolltz)), GuidesHomeActions())
        }
    }

    @Test fun guidesSearch() = guidesBar(false)

    @Test fun guidesSearchLarge() = guidesBar(true)

    // Search: Watching and the Rankings link (follow-up of F11 and F12)

    private fun search(large: Boolean) = capture("f15-search-watching", large, tab = null) {
        val data = SearchTestData
        SearchScreen(
            SearchScreenState(
                query = "",
                scope = SearchScope.All,
                indexPhase = IndexPhase.Ready(data.realPrepared),
                sections = emptyList(),
                characterPhase = CharacterPhase.Idle,
                recents = emptyList(),
                forCharacter = ForCharacter("Lightningbolltz", "Elemental Shaman guide"),
                build = SearchTestData.BUILD,
                watching = RaidFixture.load(WatchingList.serializer(), "watching.list").rows.map {
                    CharacterSearchRow(it.key, it.region, it.ruleset, it.name, it.className, it.level)
                },
                guidesInApp = true,
                rankingsInApp = true,
            ),
            SearchActions(),
            focusField = false,
        )
    }

    @Test fun searchWatching() = search(false)

    @Test fun searchWatchingLarge() = search(true)

    // Round 2: F11.4 Builds seen's `Open in planner`

    private fun openPlanner(large: Boolean) = capture("f15-char-open-planner", large, ForeverTab.Home) {
        val state = CharacterScreenState(
            CharacterRoute("us", "normal", "Lightningbolltz"),
            CharacterPagePhase.Loaded(RankingsTestData.yours()),
            WatchState(watching = true),
            RankingsTestData.viewer(),
            RankingsTestData.items().items.associateBy { it.id },
            RankingsTestData.BUILD,
            reportsOn = true,
        )
        val list = rememberScrollState(Int.MAX_VALUE)
        Box(Modifier.fillMaxSize().verticalScroll(list)) {
            CharacterContent(state, RankingsTestData.specs(), now, CharacterActions(openPlanner = {}))
        }
    }

    @Test fun charOpenPlanner() = openPlanner(false)

    @Test fun charOpenPlannerLarge() = openPlanner(true)

    // Round 2: F13.5 `Compare with mine` and the compare it opens

    private val sharedView = SharedBuildView(RaidFixture.load(Build.serializer(), "builds.get.compare"), TalentTestData.warrior(), "warrior", "Warrior", "Orc")
    private val mine = RaidFixture.load(Me.serializer(), "me").characters.first { it.name == "Obnoxious Yell" }

    private fun compareButton(name: String, large: Boolean, canCompare: Boolean) = capture(name, large, ForeverTab.Home) {
        SharedBuildPage(SharedPageState(SharedBuildPhase.Loaded(sharedView), 0, canCompare = canCompare), SharedActions())
    }

    @Test fun compareButton() = compareButton("f15-shared-build-compare-button", false, canCompare = true)

    @Test fun compareButtonLarge() = compareButton("f15-shared-build-compare-button", true, canCompare = true)

    @Test fun compareButtonNoExport() = compareButton("f15-shared-build-compare-button-noexport", false, canCompare = false)

    @Test fun compareButtonNoExportLarge() = compareButton("f15-shared-build-compare-button-noexport", true, canCompare = false)

    private fun compare(large: Boolean) = capture("f15-shared-build-compare", large, ForeverTab.Home) {
        val presenter = SharedCompare.presenter(sharedView, mine)
        TalentsPage(
            phase = TalentsPhase.Loaded(presenter, null),
            build = BisTestData.BUILD,
            offlineStamp = null,
            shown = TalentBuildKind.Yours,
            tree = presenter.openingTree,
            title = "Compare with mine",
        )
    }

    @Test fun sharedCompare() = compare(false)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun sharedCompareLarge() = compare(true)
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

private const val LARGE_SCALE = 2f
private const val SAVED_AGO = 3_600L
private const val ELEMENTAL_BYTES = 1_240_000L
private const val FURY_BYTES = 1_160_000L

/** From apps/android/app (Gradle's test working directory) to the review folder. */
private const val REVIEW_DIR = "../../../design/reviews/2026-11-01-mobile-integration/build/and"

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XL = "w412dp-h2000dp-normal-long-notround-any-420dpi-keyshidden-nonav"
