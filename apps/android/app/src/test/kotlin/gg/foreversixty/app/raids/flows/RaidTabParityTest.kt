package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.home.flows.HomeFixtures
import gg.foreversixty.app.home.flows.OfficerItems
import gg.foreversixty.app.home.flows.RaidLiveModel
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** The Raid tab's mocked states (design/mocks/ws_app_flows.py `raid`) inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class RaidTabParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun member() = state("raid", RaidTabFixtures.guild(), raid = 1)

    @Test fun memberAt2() = state("raid", RaidTabFixtures.guild(), raid = 1, fontScale = 2f)

    @Test fun officer() = state("raid-officer", RaidTabFixtures.guild(officer = true), raid = 6)

    @Test fun noRaid() = state("raid-no-raid", RaidTabFixtures.guild().copy(next = null))

    @Test fun live() = state("raid-live", RaidTabFixtures.live())

    @Test fun noGuild() = state("raid-none", RaidTabFixtures.noGuild, activity = 2)

    @Test fun leveler() = state("raid-leveler", RaidTabView.Leveler(26), activity = 1)

    @Test fun signedOut() = state("raid-out", RaidTabView.SignedOut, activity = 0)

    private fun state(screen: String, view: RaidTabView, raid: Int = 0, activity: Int = 3, fontScale: Float = 1f) =
        ParityCheck.state(compose, screen, fontScale = fontScale) {
            Tab(TabBadges(raid = raid, activity = activity)) { RaidTabScreen(view, RaidTabActions()) }
        }

    @Composable
    private fun Tab(badges: TabBadges, content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Raid, onSelect = {}, badges = badges) {
            Column {
                CompositionLocalProvider(LocalOpenMenu provides {}) { ProfileBar(null) }
                content()
            }
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

    fun guild(officer: Boolean = false) = RaidTabView.Guild(
        guild = "Olympus XXVII",
        next = HomeFixtures.raid(),
        needs = if (officer) OfficerItems("5 have not answered, 2 of them healers", "Healers short by 3", "1 applicant", "Mirelle, Holy Priest 60, 2 min ago") else null,
        needsCount = 6,
        live = null,
        roles = emptyList(),
        ready = emptyList(),
        readyStamp = null,
        pages = pages,
    )

    fun live() = guild().copy(
        next = null,
        live = RaidLiveModel(41, "Olympus XXVII · Onyxia's Lair in progress", "Pulled 1 h 12 min ago · 18 in, 3 late · you are In as Thandrel, Damage"),
        roles = listOf(
            "Tanks 2 of 2" to "Grunvald, Brakka",
            "Healers 5 of 8" to "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen (Late, 15 min)",
        ),
        ready = listOf(
            ReadyItem(ReadyTone.Ok, "Drakefire Amulet", "Attuned", null),
            ReadyItem(ReadyTone.Miss, "Flask of the Titans", "None in your bags", "Auction house"),
            ReadyItem(ReadyTone.Warn, "Boots not enchanted", "Enchant Boots: Greater Agility", "Where to get it"),
        ),
        readyStamp = "From your last log-out, 9 h ago",
    )

    val noGuild = RaidTabView.NoGuild(
        "Fury Warrior",
        listOf(
            RecruitingGuild(2, "Olympus XXVII", "Alliance · Wed and Sat, 8:00 pm CST · needs damage"),
            RecruitingGuild(3, "Dawnbreak", "Alliance · Tue and Thu, 9:00 pm EST · needs damage"),
        ),
    )
}
