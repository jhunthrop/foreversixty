package gg.foreversixty.app.guild

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.hasScrollToIndexAction
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.performScrollToIndex
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.ExperimentalRoborazziApi
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureScreenRoboImage
import gg.foreversixty.app.account.CharacterLine
import gg.foreversixty.app.account.AccountCopy
import gg.foreversixty.app.account.YouActions
import gg.foreversixty.app.account.YouScreen
import gg.foreversixty.app.account.YouSummary
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.guild.logic.CalendarAlert
import gg.foreversixty.app.guild.logic.CalendarLinkNotice
import gg.foreversixty.app.guild.logic.CalendarLinkState
import gg.foreversixty.app.guild.logic.CalendarSettings
import gg.foreversixty.app.guild.logic.CalendarStatus
import gg.foreversixty.app.guild.logic.CalendarWhich
import gg.foreversixty.app.guild.logic.HistoryPager
import gg.foreversixty.app.guild.logic.Loadable
import gg.foreversixty.app.guild.logic.MoreState
import gg.foreversixty.app.guild.logic.RosterQuery
import gg.foreversixty.app.guild.logic.RosterRole
import gg.foreversixty.app.guild.logic.RosterSort
import gg.foreversixty.app.guild.ui.CalendarActions
import gg.foreversixty.app.guild.ui.CalendarContent
import gg.foreversixty.app.guild.ui.CalendarScreenState
import gg.foreversixty.app.guild.ui.HistoryContent
import gg.foreversixty.app.guild.ui.RosterContent
import gg.foreversixty.app.guild.ui.RosterControls
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidNightRow
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.premium.PlanSummary
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.fail
import gg.foreversixty.app.raids.makeController
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.raids.logic.GuildHead
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.ui.RaidsTabContentView
import java.io.IOException
import java.time.Instant
import kotlinx.coroutines.test.runTest
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Every F16 screen and state on a Pixel 9 (412 x 915 dp) inside the tab bar, at the default size
 * and at font scale 2.0 (`-xl`). Record: ./gradlew :app:recordRoborazziDebug
 * Verify: ./gradlew :app:verifyRoborazziDebug. The PNGs named by board cell double as the review
 * captures in design/reviews/2026-10-31-mobile-wave3-guild/build/android.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class GuildScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    /** The whole screen, so menus and dialogs (their own windows) are in the picture. */
    @OptIn(ExperimentalRoborazziApi::class)
    private fun capture(name: String, large: Boolean = false, tab: ForeverTab = ForeverTab.Raids, scrollTo: Int? = null, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(2f)
        compose.setContent {
            ForeverTheme {
                // The spec icons from the repo's own build folder, so a capture shows the real art.
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons, LocalRaidIconBuild provides HomeFixtures.BUILD) {
                    ForeverTabScaffold(selected = tab, onSelect = {}) { screen() }
                }
            }
        }
        scrollTo?.let { compose.onNode(hasScrollToIndexAction()).performScrollToIndex(it) }
        compose.waitForIdle()
        captureScreenRoboImage("src/test/screenshots/guild/$name${if (large) "-xl" else ""}.png", roborazziOptions = TOLERANT)
    }

    private val stripTime = { at: Instant -> clock().stripTime(at, RaidTime.now) }

    // F16.1 and F16.2 Roster

    @Composable
    private fun Roster(state: Loadable<GuildHome> = Loadable.Loaded(GuildFixture.home()), controls: RosterControls = RosterControls()) =
        RosterContent(state, controls, stripTime, open = {}, retry = {}, back = {})

    @Test fun roster_all() = capture("f16.1-roster-all") { Roster() }

    @Test fun roster_all_xl() = capture("f16.1-roster-all", large = true) { Roster() }

    @Test fun roster_tanks() = capture("f16.1-roster-tanks") { Roster(controls = RosterControls(RosterQuery(RosterRole.Tanks))) }

    @Test fun roster_tanks_xl() = capture("f16.1-roster-tanks", large = true) { Roster(controls = RosterControls(RosterQuery(RosterRole.Tanks))) }

    @Test fun roster_healers() = capture("f16.1-roster-healers") { Roster(controls = RosterControls(RosterQuery(RosterRole.Healers))) }

    @Test fun roster_dps() = capture("f16.1-roster-dps") { Roster(controls = RosterControls(RosterQuery(RosterRole.Dps))) }

    @Test fun roster_sortAttendance() =
        capture("f16.1-roster-sort-attendance") { Roster(controls = RosterControls(RosterQuery(sort = RosterSort.Attendance))) }

    @Test fun roster_sortAttendance_xl() =
        capture("f16.1-roster-sort-attendance", large = true) { Roster(controls = RosterControls(RosterQuery(sort = RosterSort.Attendance))) }

    @Test fun roster_sortClass() = capture("f16.1-roster-sort-class") { Roster(controls = RosterControls(RosterQuery(sort = RosterSort.ClassName))) }

    @Test fun roster_sortName() = capture("f16.1-roster-sort-name") { Roster(controls = RosterControls(RosterQuery(sort = RosterSort.Name))) }

    @Test fun roster_sortLevel() = capture("f16.1-roster-sort-level") { Roster(controls = RosterControls(RosterQuery(sort = RosterSort.Level))) }

    @Test fun roster_sortMenu() = capture("f16.1-roster-sort-menu") { Roster(controls = RosterControls(sortMenuOpen = true)) }

    /** The foot of the 38-row list: the waiting and no-account rows, and the footer. */
    @Test fun roster_foot() = capture("f16.1-roster-foot", scrollTo = 40) { Roster() }

    @Test fun roster_noMatch() =
        capture("f16.2-no-match") { Roster(controls = RosterControls(RosterQuery(RosterRole.Tanks, search = "Mendara"), searching = true)) }

    @Test fun roster_noMatch_xl() =
        capture("f16.2-no-match", large = true) { Roster(controls = RosterControls(RosterQuery(RosterRole.Tanks, search = "Mendara"), searching = true)) }

    @Test fun roster_loading() = capture("f16.1-roster-loading") { Roster(Loadable.Loading) }

    @Test fun roster_error() = capture("f16.1-roster-error") { Roster(Loadable.Failed) }

    @Test fun roster_offline() =
        capture("f16.1-roster-offline") { Roster(Loadable.Loaded(GuildFixture.home(), RaidTime.now.minusSeconds(900))) }

    // F16.3 Raid history

    @Composable
    private fun History(state: Loadable<List<RaidNightRow>>, more: MoreState = MoreState.Idle) =
        HistoryContent(state, more, GuildFixture.mine, clock(), stripTime, open = {}, loadMore = {}, retry = {}, back = {})

    private fun page(n: Int) = GuildFixture.history(n).rows

    @Test fun history_page1() = capture("f16.3-history-page1") { History(Loadable.Loaded(page(1))) }

    @Test fun history_page1_xl() = capture("f16.3-history-page1", large = true) { History(Loadable.Loaded(page(1))) }

    /** Load more answered: 23 nights and no more row, scrolled to the foot. */
    @Test fun history_page2() = runTest {
        val pager = HistoryPager(2, FakeGuildService(), MemoryRaidCache()) { RaidTime.now }
        pager.load()
        pager.loadMore()
        capture("f16.3-history-page2", scrollTo = 25) { History(pager.state, pager.more) }
    }

    @Test fun history_page2_xl() = capture("f16.3-history-page2", large = true, scrollTo = 25) {
        History(Loadable.Loaded(page(1) + page(2)), MoreState.End)
    }

    @Test fun history_moreLoading() = capture("f16.3-history-more-loading", scrollTo = 22) { History(Loadable.Loaded(page(1)), MoreState.Loading) }

    @Test fun history_moreFailed() = capture("f16.3-history-more-failed", scrollTo = 22) { History(Loadable.Loaded(page(1)), MoreState.Failed) }

    @Test fun history_loading() = capture("f16.3-history-loading") { History(Loadable.Loading) }

    @Test fun history_error() = capture("f16.3-history-error") { History(Loadable.Failed) }

    @Test fun history_empty() = capture("f16.3-history-empty") { History(Loadable.Loaded(emptyList()), MoreState.End) }

    @Test fun history_offline() = capture("f16.3-history-offline") { History(Loadable.Loaded(page(1), RaidTime.now.minusSeconds(900))) }

    @Test fun history_failsWithNothingCached() = runTest {
        val service = FakeGuildService().apply { historyAnswer = { fail(ApiError.Network(IOException("offline"))) } }
        val pager = HistoryPager(2, service, MemoryRaidCache()) { RaidTime.now }
        pager.load()
        capture("f16.3-history-error-loaded") { History(pager.state, pager.more) }
    }

    // F16.4 to F16.5 Calendar

    private val url = GuildFixture.feed().url.orEmpty()
    private val on = CalendarScreenState(CalendarStatus.On, CalendarSettings(enabled = true), CalendarLinkState.Shown(url))
    private val off = CalendarScreenState(CalendarStatus.Off, CalendarSettings(), CalendarLinkState.Shown(url))
    private val mine = on.copy(settings = on.settings.copy(which = CalendarWhich.Mine))

    @Composable
    private fun Calendar(state: CalendarScreenState, confirming: Boolean = false) =
        CalendarContent(state, CalendarActions({}, {}, {}, {}, {}, {}), back = {}, confirmingReset = confirming)

    @Test fun calendar_on() = capture("f16.4-calendar-on", tab = ForeverTab.You) { Calendar(on) }

    @Test fun calendar_on_xl() = capture("f16.4-calendar-on", large = true, tab = ForeverTab.You) { Calendar(on) }

    @Test fun calendar_off() = capture("f16.4-calendar-off", tab = ForeverTab.You) { Calendar(off) }

    @Test fun calendar_off_xl() = capture("f16.4-calendar-off", large = true, tab = ForeverTab.You) { Calendar(off) }

    @Test fun calendar_whichMine() = capture("f16.4-which-mine", tab = ForeverTab.You) { Calendar(mine) }

    @Test fun calendar_whichMine_xl() = capture("f16.4-which-mine", large = true, tab = ForeverTab.You) { Calendar(mine) }

    @Test fun calendar_alertDay() =
        capture("f16.4-alert-day", tab = ForeverTab.You) { Calendar(on.copy(settings = on.settings.copy(alert = CalendarAlert.Day))) }

    @Test fun calendar_resetAlert() = capture("f16.4-reset-alert", tab = ForeverTab.You) { Calendar(on, confirming = true) }

    @Test fun calendar_resetAlert_xl() = capture("f16.4-reset-alert", large = true, tab = ForeverTab.You) { Calendar(on, confirming = true) }

    @Test fun calendar_copied() = capture("f16.4-link-copied", tab = ForeverTab.You) { Calendar(on.copy(linkNotice = CalendarLinkNotice.Copied)) }

    @Test fun calendar_elsewhere() = capture("f16.4-link-elsewhere", tab = ForeverTab.You) { Calendar(on.copy(link = CalendarLinkState.Elsewhere)) }

    @Test fun calendar_offline() = capture("f16.4-link-offline", tab = ForeverTab.You) { Calendar(on.copy(offline = true)) }

    @Test fun calendar_linkFailed() =
        capture("f16.4-link-failed", tab = ForeverTab.You) { Calendar(on.copy(link = CalendarLinkState.Failed, linkNotice = CalendarLinkNotice.LoadFailed)) }

    @Test fun calendar_resetFailed() =
        capture("f16.4-reset-failed", tab = ForeverTab.You) { Calendar(on.copy(linkNotice = CalendarLinkNotice.ResetFailed)) }

    @Test fun calendar_removed() =
        capture("f16.4-calendar-removed", tab = ForeverTab.You) { Calendar(off.copy(status = CalendarStatus.Removed)) }

    @Test fun calendar_syncFailed() = capture("f16.4-sync-failed", tab = ForeverTab.You) { Calendar(on.copy(status = CalendarStatus.Failed)) }

    @Test fun calendar_accessOff() = capture("f16.5-access-off", tab = ForeverTab.You) { Calendar(off.copy(status = CalendarStatus.AccessOff)) }

    @Test fun calendar_accessOff_xl() =
        capture("f16.5-access-off", large = true, tab = ForeverTab.You) { Calendar(off.copy(status = CalendarStatus.AccessOff)) }

    // F16.7 Raids tab rows, and You › Calendar

    private fun raidsContent(): RaidsTabContent {
        val me = RaidFixture.load(Me.serializer(), "me")
        val raids = RaidFixture.load(GuildRaids.serializer(), "guild.raids")
        val character = requireNotNull(me.mainCharacter)
        return RaidsTabContent(
            head = GuildHead.of(requireNotNull(character.guild), me.guilds.firstOrNull(), resources()),
            character = character,
            planned = raids.planned.orEmpty(),
            nights = emptyList(),
            waiting = null,
            offlineSince = null,
        )
    }

    @Composable
    private fun RaidsTab() {
        val content = raidsContent()
        RaidsTabContentView(
            state = RaidsTabState.Loaded(content),
            clock = clock(),
            now = { RaidTime.now },
            nextController = makeController(content.planned.first()),
            openRaid = { _, _ -> },
            openNight = {},
            openWeb = {},
            retry = {},
            requestPermission = {},
            guildRows = { GuildExtrasRows(onRoster = GuildFixture.home().counts.onRoster, openRoster = {}, openHistory = {}) },
        )
    }

    @Test fun raidsTab() = capture("f16.7-raids-tab") { RaidsTab() }

    @Test fun raidsTab_xl() = capture("f16.7-raids-tab", large = true) { RaidsTab() }

    @Composable
    private fun You() {
        val me = RaidFixture.load(Me.serializer(), "me")
        val current = me.characters.first()
        Column(Modifier.fillMaxSize()) {
            YouScreen(
                summary = YouSummary(
                    battletag = me.user.battletag.orEmpty(),
                    plan = PlanSummary.of(me, Instant.parse("2025-11-01T00:00:00Z")),
                    character = CharacterLine(current.name, current.className, AccountCopy.place(current)),
                    devices = "This phone and 2 more",
                    version = "Version 0.1.0 (1) · data 1.60.1.70658",
                ),
                actions = YouActions({}, {}, {}, {}, {}, {}, {}, {}, onCalendar = {}),
            )
        }
    }

    @Test fun youCalendarRow() = capture("you-calendar-row", tab = ForeverTab.You) { You() }

    @Test fun youCalendarRow_xl() = capture("you-calendar-row", large = true, tab = ForeverTab.You) { You() }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
