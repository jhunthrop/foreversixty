package gg.foreversixty.app.rankings

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.hasContentDescription
import androidx.compose.ui.test.longClick
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.test.performTouchInput
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.ExperimentalRoborazziApi
import com.github.takahirom.roborazzi.captureRoboImage
import com.github.takahirom.roborazzi.captureScreenRoboImage
import gg.foreversixty.app.account.CharacterSelector
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.home.GearList
import gg.foreversixty.app.home.GearModel
import gg.foreversixty.app.home.HomeAnswer
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.home.HomeLogs
import gg.foreversixty.app.home.HomePhase
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.home.HomeRankings
import gg.foreversixty.app.home.HomeTime
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.RankingsPage
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.rankings.logic.Board
import gg.foreversixty.app.rankings.logic.BoardPhase
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.Rankings
import gg.foreversixty.app.rankings.logic.RankingsCopy
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.logic.WatchState
import gg.foreversixty.app.rankings.ui.CharacterContent
import gg.foreversixty.app.rankings.ui.CharacterScreenState
import gg.foreversixty.app.rankings.ui.FiltersActions
import gg.foreversixty.app.rankings.ui.FiltersSheetContent
import gg.foreversixty.app.rankings.ui.FiltersSheetState
import gg.foreversixty.app.rankings.ui.RankingsContent
import gg.foreversixty.app.rankings.ui.RankingsScreenState
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.ReportsTestData
import gg.foreversixty.app.reports.ui.MetersDragHandle
import java.time.Instant
import java.time.ZoneOffset
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Every F11 screen and state on a Pixel 9 (412 x 915 dp), at the default size and at font scale 2.0
 * (spec W3 §9). Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 * The PNGs double as the review captures in design/reviews/2026-10-31-mobile-wave3-rankings.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class RankingsScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val now = RankingsTestData.now
    private val zone = RankingsTestData.zone
    private val offlineAt: Instant = Instant.parse("2026-12-17T03:10:00Z")
    private val specs = RankingsTestData.specs()
    private val phases = RankingsTestData.phases().phases
    private val names = RankingsTestData.encounters().encounters.associate { it.slug to it.name }

    private fun capture(name: String, large: Boolean = false, scrolls: Boolean = true, tabs: Boolean = true, screen: @Composable () -> Unit) {
        setContent(large, scrolls, tabs, screen)
        compose.onRoot().captureRoboImage("src/test/screenshots/rankings/$name.png", roborazziOptions = TOLERANT)
    }

    private fun setContent(large: Boolean, scrolls: Boolean, tabs: Boolean, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    // The screens are given the scroll their route gives them; a sheet scrolls itself.
                    val ground = Modifier.fillMaxSize().background(FsColor.bg)
                    val body = @Composable { Box(if (scrolls) ground.verticalScroll(rememberScrollState()) else ground) { screen() } }
                    // A modal sheet covers the tab bar, as ModalBottomSheet does in the app.
                    if (tabs) ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) { body() } else body()
                }
            }
        }
    }

    // F11.1, F11.3 and the board's states

    private fun board(name: String, phase: BoardPhase, viewer: RankingsViewer = RankingsTestData.viewer(), large: Boolean = false) =
        capture(name, large) {
            val filters = Rankings.defaults(viewer, specs, phases, now, "onyxia")
            RankingsContent(RankingsScreenState(phase, filters, names, viewer, watched = setOf("us/normal/voltaris"), reportsOn = true), now, zone = zone)
        }

    private fun loaded(page: RankingsPage, offline: Instant? = null) = BoardPhase.Loaded(Board.of(page, offline))

    @Test fun onPage() = board("rankings-onpage-and", loaded(RankingsTestData.page()))

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun onPageLarge() = board("rankings-onpage-large-and", loaded(RankingsTestData.page()), large = true)

    @Test fun offPage() = board("rankings-offpage-and", loaded(RankingsTestData.offPage()), RankingsTestData.viewer(RankingsTestData.sparkwick))

    /** The long press on Stormwrath (watched): Stop watching, Open the report, Share. */
    @OptIn(ExperimentalRoborazziApi::class)
    @Test fun rowMenu() {
        setContent(large = false, scrolls = true, tabs = true) {
            val viewer = RankingsTestData.viewer()
            val filters = Rankings.defaults(viewer, specs, phases, now, "onyxia")
            val state = RankingsScreenState(loaded(RankingsTestData.page()), filters, names, viewer, watched = setOf("us/normal/stormwrath"), reportsOn = true)
            RankingsContent(state, now, zone = zone)
        }
        compose.onNode(hasContentDescription("Rank 2 of", substring = true)).performTouchInput { longClick() }
        compose.waitForIdle()
        captureScreenRoboImage("src/test/screenshots/rankings/rankings-menu-and.png", roborazziOptions = TOLERANT)
    }

    @Test fun unranked() = board("rankings-empty-and", loaded(RankingsTestData.empty()))

    @Test fun signedOut() = board("rankings-signedout-and", loaded(mixed()), RankingsViewer.SignedOut)

    @Test fun offline() = board("rankings-offline-and", loaded(RankingsTestData.page(), offlineAt))

    @Test fun error() = board("rankings-error-and", BoardPhase.Failed)

    @Test fun loading() = board("rankings-loading-and", BoardPhase.Loading)

    /** Signed out the board is every class: the fixture's rows with every other one a mage, so the crests show. */
    private fun mixed(): RankingsPage = RankingsTestData.page().let { page ->
        page.copy(
            characterRank = null,
            rows = page.rows.mapIndexed { i, r -> if (i % 2 == 1) r.copy(player = r.player.copy(className = "Mage", spec = "Fire")) else r },
        )
    }

    // F11.2 Filters

    /** The sheet at its large detent over the whole window (it covers the tab bar); [atEnd] scrolls to `Reset filters`. */
    private fun sheet(name: String, choosingEncounter: Boolean, today: Boolean = false, atEnd: Boolean = false, large: Boolean = false) =
        capture(name, large, scrolls = false, tabs = false) {
            val filters = Rankings.defaults(RankingsTestData.viewer(), specs, phases, now, "onyxia").copy(todayOnly = today)
            val scroll = rememberScrollState()
            if (atEnd) LaunchedEffect(scroll.maxValue) { scroll.scrollTo(scroll.maxValue) }
            Column(Modifier.padding(top = 24.dp).fillMaxSize().background(GameColors.sheet, RoundedCornerShape(topStart = 28.dp, topEnd = 28.dp))) {
                Box(Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) { MetersDragHandle() }
                FiltersSheetContent(FiltersSheetState(filters, RankingsTestData.encounters().encounters, Rankings.phaseChoices(phases, now), choosingEncounter), FiltersActions(), scroll)
            }
        }

    @Test fun filters() = sheet("filters-and", choosingEncounter = false)

    @Test fun filtersTodayOff() = sheet("filters-today-off-and", choosingEncounter = false, atEnd = true)

    @Test fun filtersTodayOn() = sheet("filters-today-on-and", choosingEncounter = false, today = true, atEnd = true)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun filtersLarge() = sheet("filters-large-and", choosingEncounter = false, large = true)

    @Test fun encounters() = sheet("filters-encounters-and", choosingEncounter = true)

    // F11.4 to F11.7 A character

    private fun character(
        name: String,
        phase: CharacterPhase,
        watch: WatchState? = WatchState(watching = false),
        route: CharacterRoute = CharacterRoute("us", "normal", "Voltaris"),
        large: Boolean = false,
    ) = capture(name, large) {
        val state = CharacterScreenState(
            route, phase, watch, RankingsTestData.viewer(), RankingsTestData.items().items.associateBy { it.id }, RankingsTestData.BUILD, reportsOn = true,
        )
        CharacterContent(state, specs, now, zone = zone)
    }

    private fun profile(p: CharacterProfile, offline: Instant? = null) = CharacterPhase.Loaded(p, offline)

    @Test fun other() = character("character-other-and", profile(RankingsTestData.voltaris()))

    @Test fun watchOn() = character("character-watch-on-and", profile(RankingsTestData.voltaris()), WatchState(watching = true))

    @Test fun watchFailed() = character("character-watch-failed-and", profile(RankingsTestData.voltaris()), WatchState(watching = false, failed = true))

    @Test @Config(qualifiers = PIXEL_9_XTALL)
    fun yours() = character(
        "character-yours-and", profile(RankingsTestData.yours()), route = CharacterRoute("us", "normal", "Lightningbolltz"),
    )

    @Test @Config(qualifiers = PIXEL_9_XTALL)
    fun gear() = character("character-gear-and", profile(RankingsTestData.voltaris()))

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun noGear() = character("character-nogear-and", profile(RankingsTestData.voltaris().copy(gear = emptyList(), gearGapTotal = null)))

    @Test fun noLogs() = character("character-nologs-and", profile(RankingsTestData.noLogs()), route = CharacterRoute("us", "pvp", "Freshling"))

    @Test fun notFound() = character("character-notfound-and", CharacterPhase.NotFound, watch = null, route = CharacterRoute("us", "normal", "Voltxris"))

    @Test fun characterOffline() = character("character-offline-and", profile(RankingsTestData.voltaris(), offlineAt))

    @Test fun characterLoading() = character("character-loading-and", CharacterPhase.Loading, watch = null)

    @Test fun characterError() = character("character-error-and", CharacterPhase.Failed, watch = null)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun otherLarge() = character("character-other-large-and", profile(RankingsTestData.voltaris()), large = true)

    @Test @Config(qualifiers = PIXEL_9_XXL)
    fun gearLarge() = character("character-gear-large-and", profile(RankingsTestData.voltaris()), large = true)

    // Home: the Rankings row under Logs, and Your gear as the Paperdoll

    @Test @Config(qualifiers = PIXEL_9_HOME)
    fun homeRankings() = capture("home-rankings-and") {
        val res = LocalResources.current
        val viewer = RankingsTestData.viewer()
        val line = RankingsCopy(res, zone).homeLine(RankingsTestData.page(), Rankings.defaults(viewer, specs, phases, now, "onyxia"), names)
        val presenter = HomePresenter(
            LIGHTNINGBOLLTZ, emptyList(), HomeTime(HomeFixtures.now, ZoneOffset.UTC), res,
            logs = HomeLogs(Reports.lastRaid(ReportsTestData.list().rows)), rankings = HomeRankings(line),
        )
        // Lightningbolltz's own Home answer (character.home.yours): the six upgrades behind his six gold dots.
        val home = HomeFixtures.home("character.home.yours")
        Column {
            // The selector row as CharacterSelectorBar draws it for the session's current character.
            Row(Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 8.dp).heightIn(min = 48.dp), verticalAlignment = Alignment.CenterVertically) {
                CharacterSelector(LIGHTNINGBOLLTZ_ACCOUNT, open = false, onClick = {})
            }
            HomeAnswer(HomePhase.Loaded(home), presenter, home.dataBuild, null)
        }
    }

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun homeGear() = capture("home-gear-and") { HomeGear() }

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun homeGearLarge() = capture("home-gear-large-and", large = true) { HomeGear() }

    @Composable
    private fun HomeGear() {
        val res = LocalResources.current
        val home = HomeFixtures.home()
        val sources = HomeFixtures.tooltipSources()
        val probe = HomePresenter.tooltipRequest(ItemLook(0), home, HomeFixtures.velaren, res)
        val blocker = GearModel.offHandBlocker(home, emptyMap(), sources.items, ItemTooltipModel.band(probe, sources.bis))
        GearList(
            GearModel.rows(home, emptyMap(), sources.items, res, blocker),
            "What you wear, and the pick for 60 · raid. Tap an item for its tooltip.",
            home.dataBuild,
        )
    }
}

/** Home's character for the Rankings row: the board's own Elemental Shaman. */
private val LIGHTNINGBOLLTZ = gg.foreversixty.app.home.HomeCharacter(
    key = "us/normal/lightningbolltz", region = "us", ruleset = "normal", name = "Lightningbolltz", race = "troll",
    className = "shaman", specName = "Elemental", level = 60, guild = "Hordebreakers", faction = "horde",
)

/** The same character as the account's, for the selector. */
private val LIGHTNINGBOLLTZ_ACCOUNT = gg.foreversixty.app.model.AccountCharacter(
    key = LIGHTNINGBOLLTZ.key, region = LIGHTNINGBOLLTZ.region, ruleset = LIGHTNINGBOLLTZ.ruleset, name = LIGHTNINGBOLLTZ.name,
    source = "addon", className = LIGHTNINGBOLLTZ.className, level = LIGHTNINGBOLLTZ.level, faction = LIGHTNINGBOLLTZ.faction,
    race = LIGHTNINGBOLLTZ.race, spec = LIGHTNINGBOLLTZ.specName,
)

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_TALL = "w412dp-h1300dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XTALL = "w412dp-h2000dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_HOME = "w412dp-h2400dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XL = "w412dp-h2600dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XXL = "w412dp-h4400dp-normal-long-notround-any-420dpi-keyshidden-nonav"
