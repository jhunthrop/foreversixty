package gg.foreversixty.app.raids.flows

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.activity.ActivityFixtures
import gg.foreversixty.app.home.flows.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.home.flows.OfficerItems
import gg.foreversixty.app.home.flows.RaidLiveModel
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** The Raid tab's mocked states (design/mocks/ws_app_flows.py `raid`, round 5 on Android) inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class RaidTabParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun member() = state("raid", RaidTabFixtures.guild(), raid = 0)

    @Test fun memberAt2() = state("raid", RaidTabFixtures.guild(), raid = 0, fontScale = 2f)

    @Test fun signup() = state("raid-signup", RaidTabFixtures.signup(), raid = 0)

    @Test fun signupAt2() = state("raid-signup", RaidTabFixtures.signup(), raid = 0, fontScale = 2f)

    @Test fun officer() = state("raid-officer", RaidTabFixtures.guild(officer = true), raid = 6)

    @Test fun noRaid() = state("raid-no-raid", RaidTabFixtures.guild().copy(next = null, roles = emptyList(), ready = emptyList(), readyStamp = null))

    @Test fun live() = state("raid-live", RaidTabFixtures.live())

    @Test fun noGuild() = state("raid-none", RaidTabFixtures.noGuild, activity = 2)

    @Test fun leveler() = state("raid-leveler", RaidTabView.Leveler(26), activity = 1)

    @Test fun signedOut() = state("raid-out", RaidTabView.SignedOut, activity = 0)

    private fun state(screen: String, view: RaidTabView, raid: Int = 0, activity: Int = 3, fontScale: Float = 1f) =
        ParityCheck.state(compose, screen, fontScale = fontScale) {
            Tab(TabBadges(raid = raid, activity = activity)) {
                RaidTabScreen(view, RaidTabActions(), topBar = { ProfileBar(null, "Raid") })
            }
        }

    @Composable
    private fun Tab(badges: TabBadges, content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Raid, onSelect = {}, badges = badges) {
            // The game icons from the repo's build data, as Home's parity test paints them.
            CompositionLocalProvider(
                LocalOpenMenu provides {},
                LocalRaidIconBuild provides ActivityFixtures.BUILD,
                LocalItemIconPainter provides gg.foreversixty.app.home.HomeFixtures.repoIcons,
            ) { content() }
        }
    }
}

/** The mock's guild, Olympus XXVII, and its raid. */
object RaidTabFixtures {
    private val pages = listOf(
        GuildPageRow(GuildPage.Roster, "Roster", "38 members, 31 at 60"),
        GuildPageRow(GuildPage.Calendar, "Calendar", "Onyxia Wed, Barrow Deeps Sat"),
        GuildPageRow(GuildPage.Loot, "Loot history", "Last raid: 4 pieces"),
        GuildPageRow(GuildPage.Sims, "Guild sims", "You are #7 of 14 DPS"),
    )

    private val coming = listOf(
        RoleLine("Tanks", 2, 2, "Grunvald, Brakka"),
        RoleLine("Healers", 5, 8, "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen · 3 short"),
        RoleLine("Damage", 11, 30, "Thandrel, Kaelra, Thalrin and 8 more"),
    )

    private val ready = listOf(
        ReadyItem(ReadyTone.Ok, "Drakefire Amulet", "Attuned", null),
        ReadyItem(ReadyTone.Ok, "Elixir of the Mongoose", "3 in your bags", null),
        ReadyItem(ReadyTone.Miss, "Flask of the Titans", "None in your bags", "Auction house"),
        ReadyItem(ReadyTone.Warn, "Boots not enchanted", "Enchant Boots: Greater Agility", "Where to get it"),
    )

    private fun head(answer: SignupState, chip: String, lateMinutes: Int? = null) = HomeFixtures.raid(answer, lateOpen = answer == SignupState.Late).copy(
        headLine = "8:00 pm CST · pull in 2 h 48 min",
        who = "As Thandrel, Damage",
        chip = chip,
        lateMinutes = lateMinutes,
    )

    fun guild(officer: Boolean = false) = RaidTabView.Guild(
        guild = "Olympus XXVII",
        next = head(SignupState.In, "In"),
        needs = if (officer) OfficerItems("5 have not answered, 2 of them healers", "Healers short by 3", "1 applicant", "Mirelle, Holy Priest 60, 2 min ago") else null,
        needsCount = 6,
        live = null,
        roles = coming,
        ready = ready,
        readyStamp = "From your last log-out, 9 h ago",
        pages = pages,
    )

    /** F1 change the answer (mock m-android-signup): Late, 30 min chosen, the minutes open under the answers. */
    fun signup() = guild().copy(next = head(SignupState.Late, "Late\u00a0·\u00a030\u00a0min", lateMinutes = 30))

    fun live() = guild().copy(
        next = null,
        live = RaidLiveModel(41, "Olympus XXVII · Onyxia's Lair in progress", "Pulled 1 h 12 min ago · 18 in, 3 late · you are In as Thandrel, Damage"),
        roles = listOf(
            RoleLine("Tanks", 2, 2, "Grunvald, Brakka"),
            RoleLine("Healers", 5, 8, "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen (Late\u00a0·\u00a015\u00a0min) · 3 short"),
            RoleLine("Damage", 11, 30, "Thandrel, Kaelra, Thalrin and 8 more"),
        ),
    )

    val noGuild = RaidTabView.NoGuild(
        "Fury Warrior",
        listOf(
            RecruitingGuild(2, "Olympus XXVII", "Alliance · Wed and Sat, 8:00 pm CST · needs damage"),
            RecruitingGuild(3, "Dawnbreak", "Alliance · Tue and Thu, 9:00 pm EST · needs damage"),
        ),
    )
}
