package gg.foreversixty.app.home.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.activity.ActivityFixtures
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.TallyItem
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** Home's mocked moments (design/mocks/ws_app_flows.py `HOME`) inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class HomeParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun raidDay() = state("home-raid-day", HomeFixtures.raidDay(), raid = 1)

    @Test fun raidDayAt2() = state("home-raid-day", HomeFixtures.raidDay(), raid = 1, fontScale = 2f)

    @Test fun raidDayLate() = state("home-raid-day-late", HomeFixtures.raidDay(answer = SignupState.Late, lateOpen = true))

    @Test fun raidDayOfficer() = state("home-raid-day-officer", HomeFixtures.raidDay(officer = true), raid = 6)

    @Test fun raidLive() = state("home-raid-live", HomeFixtures.raidLive())

    @Test fun afterRaid() = state("home-after-raid", HomeFixtures.afterRaid())

    @Test fun normalDay() = state("home-60", HomeFixtures.normalDay(), raid = 1)

    @Test fun normalDayAt1_5() = state("home-60", HomeFixtures.normalDay(), raid = 1, fontScale = 1.5f)

    @Test fun normalDayAnswered() = state("home-60-answered", HomeFixtures.normalDay(SignupState.In))

    @Test fun noGuild() = state("home-noguild", HomeFixtures.noGuild(), activity = 2)

    @Test fun leveling() = state("home-34", HomeFixtures.leveling(), activity = 1)

    @Test fun loading() = state("home-loading", HomeView.Loading(HomeFixtures.head), activity = 0)

    @Test fun error() = state("home-error", HomeView.Failed(HomeFixtures.head), activity = 0)

    @Test fun noCharacter() = state("home-new", HomeView.NoCharacter, activity = 0)

    private fun state(screen: String, view: HomeView, raid: Int = 0, activity: Int = 3, fontScale: Float = 1f) =
        ParityCheck.state(compose, screen, fontScale = fontScale) {
            Tab(TabBadges(raid = raid, activity = activity)) { HomeFlowScreen(view, HomeActions()) }
        }

    @Composable
    private fun Tab(badges: TabBadges, content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}, badges = badges) {
            Column {
                CompositionLocalProvider(LocalOpenMenu provides {}) { ProfileBar(null) }
                content()
            }
        }
    }
}

/** The mock's scenario: Thandrel, Fury Warrior 60 of Olympus XXVII; Onyxia's Lair, Wed 6 Jan, 8:00 pm CST. */
object HomeFixtures {
    private const val BUILD = ActivityFixtures.BUILD
    private const val ONY = "inv_misc_head_dragon_black"

    val head = HomeHead("Thandrel", "warrior", listOf("Level 60 Fury Warrior", "Olympus XXVII"))
    private val felsnap = HomeHead("Felsnap", "hunter", listOf("Level 34 Beast Mastery Hunter", "6 levels to band 40"))

    private fun upgrade(name: String, quality: Int, icon: String, source: String, gain: String) =
        UpgradeLine(ItemLook(name.hashCode(), name, quality, icon), source, gain, "DPS", "head", null, source)

    private val raidUpgrades = listOf(
        upgrade("Lionheart Helm", 4, "inv_helmet_36", "Head · Blacksmithing", "+14.5"),
        upgrade("Greenhammer", 4, "inv_mace_1h_raidhyjalc60_d_01", "Off hand · Blacksmithing", "+9.7"),
        upgrade("Don Julio's Band", 4, "inv_jewelry_ring_35", "Finger · Stormpike Guard, needs revered", "+6.8"),
    )

    private fun upgrades(bestRun: Boolean) =
        HomeBlock.Upgrades(raidUpgrades, if (bestRun) "Best run: Blackrock Spire, Blackcrow +1.9 DPS, 8%" else null, BUILD)

    private fun act(id: Long, icon: String, title: String, line: String, unread: Boolean) = HomeActivityRow(
        ActivityRow(id = id, kind = "x", group = "gear", title = title, body = "", link = "", occurredAt = "2026-11-06T10:00:00Z", read = !unread, icon = icon),
        line,
    )

    private val dayRows = listOf(
        act(8, "ability_warrior_innerrage", "Kaelra passed you in the guild sims", "1,012.3 DPS, you are #7 of 14 · 1 h ago", true),
        act(4, "inv_misc_pocketwatch_01", "Sim done: 948.6 DPS ±6", "+21 vs your last · 3 h ago", true),
        act(1, "inv_misc_book_09", "Fury Warrior BiS changed", "A new pick for Hands · yesterday", false),
    )

    private fun newForYou(rows: List<HomeActivityRow> = dayRows, unread: Int = 3) = HomeBlock.Activity(rows, unread)

    private val tally = listOf(TallyItem("Tanks", 2, 2), TallyItem("Healers", 5, 8), TallyItem("Damage", 11, 30))

    fun raid(answer: SignupState? = null, lateOpen: Boolean = false, officer: Boolean = false) = RaidCardModel(
        raidId = 41,
        icon = ONY,
        eyebrow = "Olympus XXVII · Onyxia's Lair tonight",
        clock = listOf(ClockPart("02", "hours"), ClockPart("48", "min")),
        clockLabel = "2 hours 48 minutes",
        aside = "Wed 6 Jan, 8:00 pm CST · answer by 6:00 pm CST",
        title = "Onyxia's Lair · Wed 6 Jan, 8:00 pm CST",
        sub = if (answer == null) "18 in, 3 late · you have not answered, as Thandrel, Damage" else "You are In, as Thandrel, Damage",
        me = if (answer == null) "You have not answered, as Thandrel, Damage" else "You are ${answer.name}, as Thandrel, Damage",
        answer = answer,
        lateOpen = lateOpen,
        tally = tally,
        ready = ReadyLine("2 things before you pull", "No flask in your bags; boots not enchanted · from your last log-out, 9\u00a0h ago"),
        officer = if (officer) OfficerItems("5 have not answered, 2 of them healers", "Healers short by 3", "1 applicant", "Mirelle, Holy Priest 60, 2 min ago") else null,
    )

    fun raidDay(answer: SignupState? = null, lateOpen: Boolean = false, officer: Boolean = false) =
        HomeView.Blocks(head, listOf(HomeBlock.NextRaidCard(raid(answer, lateOpen, officer)), upgrades(bestRun = false), newForYou()), BUILD)

    fun raidLive() = HomeView.Blocks(
        head,
        listOf(
            HomeBlock.RaidLive(RaidLiveModel(41, "Olympus XXVII · Onyxia's Lair in progress", "Pulled 1 h 12 min ago · 18 in, 3 late · you are In as Thandrel, Damage")),
            newForYou(
                listOf(
                    act(9, ONY, "Brightwen is Late, 15 min", "Onyxia's Lair · 4 min ago", true),
                    act(10, ONY, "Selka is In", "Restoration Druid · 20 min ago", true),
                    act(3, ONY, "Healers short for Onyxia tonight", "3 of 8 missing · 3 h ago", false),
                ),
            ),
        ),
        BUILD,
    )

    fun afterRaid() = HomeView.Blocks(
        head,
        listOf(
            HomeBlock.LastNight(
                listOf(
                    act(7, "inv_axe_09", "You got Deathbringer", "You got it from Onyxia · main hand, +12.6 DPS over your worn Ravencrest's Legacy", true),
                    act(6, ONY, "Onyxia's Lair report ready", "87th percentile, your best on Onyxia so far · the coach has 2 notes", true),
                ),
                simAgain = true,
            ),
            upgrades(bestRun = true),
            newForYou(
                listOf(
                    act(8, "ability_warrior_innerrage", "Kaelra passed you in the guild sims", "1,012.3 DPS, you are #7 of 14 · 8 h ago", true),
                    act(4, "inv_misc_pocketwatch_01", "Sim done: 948.6 DPS ±6", "+21 vs your last · yesterday", false),
                    act(1, "inv_misc_book_09", "Fury Warrior BiS changed", "A new pick for Hands · Tue", false),
                ),
            ),
        ),
        BUILD,
    )

    fun normalDay(answer: SignupState? = null) =
        HomeView.Blocks(head, listOf(upgrades(bestRun = true), HomeBlock.NextRaidLine(raid(answer)), newForYou()), BUILD)

    fun noGuild() = HomeView.Blocks(
        head,
        listOf(
            upgrades(bestRun = true),
            newForYou(unread = 2),
            HomeBlock.FindGuild("Raid with a guild that plans its nights here, and answer its raids from the app."),
        ),
        BUILD,
    )

    fun leveling() = HomeView.Blocks(
        felsnap,
        listOf(
            HomeBlock.Upgrades(
                listOf(
                    upgrade("Felbough Arm", 3, "inv_spear_08", "Main hand · Fel Ancient, Ashenvale", "+1.3"),
                    upgrade("Cartilage Shapers", 3, "inv_gauntlets_18", "Hands · quest, Lyn the Ignored", "+0.8"),
                    upgrade("Ring of Precision", 3, "inv_jewelry_ring_14", "Finger · Blackfathom Deeps", "+0.6"),
                ),
                "Best run: Blackfathom Deeps, Ring of Precision +0.6 DPS",
                BUILD,
            ),
            newForYou(
                listOf(
                    act(5, "inv_misc_book_09", "Level 34: your 30 to 39 list", "Beast Mastery Hunter · 2 h ago", true),
                    act(4, "inv_misc_pocketwatch_01", "Sim done: 58.4 DPS ±1", "+3.1 vs your last · yesterday", false),
                    act(2, "inv_misc_book_09", "Beast Mastery Hunter moved to A", "Leveling tiers, 30 to 39 · 2 d ago", false),
                ),
                unread = 1,
            ),
        ),
        BUILD,
    )
}
