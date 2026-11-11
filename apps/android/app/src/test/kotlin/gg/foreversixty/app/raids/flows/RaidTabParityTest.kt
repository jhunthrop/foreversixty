package gg.foreversixty.app.raids.flows

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.activity.ActivityFixtures
import gg.foreversixty.app.home.flows.AnswerControl
import gg.foreversixty.app.home.flows.HomeFixtures
import gg.foreversixty.app.raids.logic.StatusLine
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.home.flows.OfficerItems
import gg.foreversixty.app.home.flows.RaidLiveModel
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.CharacterBar
import gg.foreversixty.app.menu.LocalPortraitPainter
import gg.foreversixty.app.menu.MenuFixtures
import gg.foreversixty.app.model.AccountCharacter
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
import gg.foreversixty.app.flows.TallyItem

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

    @Test fun noRaid() = state("raid-no-raid", RaidTabFixtures.guild().copy(next = null, roles = emptyList(), foot = emptyList(), notes = null, ready = emptyList(), readyStamp = null))

    @Test fun live() = state("raid-live", RaidTabFixtures.live())

    @Test fun noGuild() = state("raid-none", RaidTabFixtures.noGuild, activity = 2)

    @Test fun leveler() = state("raid-leveler", RaidTabView.Leveler(26), activity = 1)

    @Test fun signedOut() = state("raid-out", RaidTabView.SignedOut, activity = 0)

    private fun state(screen: String, view: RaidTabView, raid: Int = 0, activity: Int = 3, fontScale: Float = 1f) =
        ParityCheck.state(compose, screen, fontScale = fontScale) {
            Tab(TabBadges(raid = raid, activity = activity)) {
                RaidTabScreen(view, RaidTabActions(), topBar = { CharacterBar(barFor(view), "Raid") })
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
                LocalPortraitPainter provides MenuFixtures.repoPortraits,
            ) { content() }
        }
    }

    /** The bar's character: none signed out, Felsnap while leveling, else Thandrel (the mocks' portraits). */
    private fun barFor(view: RaidTabView): AccountCharacter? = when (view) {
        RaidTabView.SignedOut -> null
        is RaidTabView.Leveler -> MenuFixtures.all.first { it.name == "Felsnap" }
        else -> MenuFixtures.thandrel
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

    /** A Who is coming row as the presenter builds it: the tally, the names in one class, the rest. */
    private fun role(label: String, have: Int, want: Int, names: String, more: String? = null, split: String? = null) =
        RoleLine(TallyItem(label, have, want), names.split(", ").map { RoleName(it, CLASS) }, more, split)

    private val coming = listOf(
        role("Tanks", 2, 2, "Grunvald, Brakka"),
        role("Healers", 5, 8, "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen (Late · 15 min)"),
        role("Damage", 11, 30, "Thandrel, Kaelra, Thalrin", " and 8 more", "5 melee, 6 ranged"),
    )

    /** The lines under Who is coming: the other answers and what the comp misses; officers also see readiness. */
    // Row 6: the foot's not answered is the same population as Needs you's `5 have not answered`.
    private fun foot(officer: Boolean) = listOfNotNull(
        RoleFoot("Bench 1 · Out 3 · 5 not answered"),
        RoleFoot("Missing Fort · WF 1 of 2", cover = COVER),
        RoleFoot(
            "Not synced this week: Kaelra, Thalrin, Brakka, Mirethil",
            onlyKeys = setOf("us/pvp/kaelra", "us/pvp/thalrin", "us/pvp/brakka", "us/pvp/mirethil"),
        ).takeIf { officer },
    )

    /** Row 4: `Missing` once, then each gap in wipe red beside its spell icon. */
    private val COVER = CoverLine(
        missing = listOf(CoverItem("Fort", "spell_holy_wordfortitude")),
        short = listOf(CoverItem("WF", "spell_nature_windfury", have = 1, need = 2)),
        more = 0,
    )

    private const val NOTES = "Bring fire resistance gear and flasks. Whelps go to the left wall, the tanks call the deep " +
        "breath in voice; consumables check at 7:50 pm."

    private val ready = listOf(
        ReadyItem(ReadyTone.Ok, "Drakefire Amulet", "Attuned", null),
        ReadyItem(ReadyTone.Ok, "Elixir of the Mongoose", "3 in your bags", null),
        ReadyItem(ReadyTone.Miss, "Flask of the Titans", "None in your bags", "Auction house"),
        ReadyItem(ReadyTone.Warn, "Boots not enchanted", "", null),
    )

    private fun head(answer: SignupState, lateMinutes: Int? = null, control: AnswerControl = AnswerControl()) =
        HomeFixtures.raid(answer, lateOpen = answer == SignupState.Late).copy(
            headLine = "8:00 pm CST · pull in 2 h 48 min",
            facts = "Invites 7:45 pm · 40-man",
            who = "As Thandrel, Damage",
            lateMinutes = lateMinutes,
            control = control,
        )

    fun guild(officer: Boolean = false) = RaidTabView.Guild(
        guild = "Olympus XXVII",
        next = head(SignupState.In),
        needs = if (officer) OfficerItems("5 have not answered, 2 of them healers", "Healers short by 3", "1 applicant", "Mirelle, Holy Priest 60, 2 min ago") else null,
        needsCount = 6,
        live = null,
        roles = coming,
        foot = foot(officer),
        notes = NOTES,
        needsFirst = officer,
        ready = ready,
        readyStamp = "From your last log-out, 9 h ago",
        pages = pages,
    )

    /** F1 change the answer (mock m-android-signup): Late, 30 min chosen, the minutes open under the answers. */
    fun signup() = guild().copy(next = head(SignupState.Late, lateMinutes = 30, control = AnswerControl(status = StatusLine.Saved("Saved."))))

    fun live() = guild().copy(
        foot = emptyList(),
        notes = null,
        next = null,
        live = RaidLiveModel(41, "Olympus XXVII · Onyxia's Lair in progress", "Pulled 1 h 12 min ago · 18 in, 3 late · you are In as Thandrel, Damage"),
        roles = listOf(
            role("Tanks", 2, 2, "Grunvald, Brakka"),
            role("Healers", 5, 8, "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen (Late\u00a0·\u00a015\u00a0min)"),
            role("Damage", 11, 30, "Thandrel, Kaelra, Thalrin", " and 8 more", "5 melee, 6 ranged"),
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

/** The mock's Who is coming draws every name in one class colour. */
private const val CLASS = "warrior"
