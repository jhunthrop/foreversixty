package gg.foreversixty.app.raids.flows

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.activity.ActivityFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.menu.CharacterBar
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.LocalPortraitPainter
import gg.foreversixty.app.menu.MenuFixtures
import gg.foreversixty.app.model.RaidAssignedName
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.raids.logic.AssignmentEncounterLine
import gg.foreversixty.app.raids.logic.RaidJobLine
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * The Guild tab's four states (spec 2026-11-10-guild-tab.md section 2) against their mocks
 * (design/mocks/ws_app_guild.py) inside the tab shell: at 60, officer, below 60, no guild (at 60 and below).
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class GuildTabParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun atCap() = state("guild-tab", GuildTabFixtures.atCap(), guild = 0, activity = 3)

    // 2026-11-11, lane guild-tab: the officer's roster is 11 rows under four group labels; their 18.29 line boxes run 2 to 3
    // taller than the mock's over the roster, so the sections under it sit 12 to 13 lower; each section's own height holds.
    @Test fun officer() = state(
        "guild-tab-officer", GuildTabFixtures.atCap(officer = true), guild = 6, activity = 3,
        drift = setOf(
            "section-loot: y 3139, the mock has 3127 (within 8)",
            "section-pages: y 3769, the mock has 3756 (within 8)",
            "section-progression: y 3640, the mock has 3627 (within 8)",
            "section-recent-logs: y 3391, the mock has 3379 (within 8)",
        ),
    )

    @Test fun leveler() = state("guild-tab-leveler", GuildTabFixtures.leveler(), activity = 1, felsnap = true)

    // 2026-11-11, lane guild-tab: the recruiting rows' trail (Apply 44 dp, Applied's 20 line) is 10 shorter than the
    // mock's `.tr` with its 8 margin and 48 px button, so Find your guild sits 10 to 11 higher.
    @Test fun noGuild() = state(
        "guild-tab-none", GuildTabFixtures.noGuild(leveler = false), activity = 2,
        drift = setOf("section-find-your-guild: y 289, the mock has 300 (within 8)"),
    )

    @Test fun noGuildLeveler() = state(
        "guild-tab-none-leveler", GuildTabFixtures.noGuild(leveler = true), activity = 1, felsnap = true,
        drift = setOf("section-find-your-guild: y 356, the mock has 366 (within 8)"),
    )

    private fun state(
        screen: String,
        view: RaidTabView,
        guild: Int = 0,
        activity: Int = 0,
        felsnap: Boolean = false,
        drift: Set<String> = emptySet(),
    ) =
        ParityCheck.state(compose, screen, expectedDrift = drift) {
            Tab(TabBadges(guild = guild, activity = activity)) {
                val bar = if (felsnap) MenuFixtures.all.first { it.name == "Felsnap" } else MenuFixtures.thandrel
                RaidTabScreen(view, RaidTabActions(), topBar = { CharacterBar(bar, "Guild") })
            }
        }

    @Composable
    private fun Tab(badges: TabBadges, content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Guild, onSelect = {}, badges = badges) {
            CompositionLocalProvider(
                LocalOpenMenu provides {},
                LocalRaidIconBuild provides ActivityFixtures.BUILD,
                LocalItemIconPainter provides gg.foreversixty.app.home.HomeFixtures.repoIcons,
                LocalPortraitPainter provides MenuFixtures.repoPortraits,
            ) { content() }
        }
    }
}

/** The mock's guild for the Guild tab: Olympus XXVII, 38 members, 31 at 60; Thandrel at 60, Felsnap at 36. */
object GuildTabFixtures {
    private val head = GuildTabHead("Olympus XXVII", "alliance", null, "38 members · Alliance · US PvP")

    private val recent = listOf(
        GuildMemberLine("Grunvald", "warrior", null, "60 · Protection Warrior · item level 66", "just now"),
        GuildMemberLine("Mirethil", "priest", null, "60 · Holy Priest · item level 63", "12 min ago"),
        GuildMemberLine("Kaelra", "mage", null, "60 · Frost Mage · item level 61", "40 min ago"),
        GuildMemberLine("Thornwick", "druid", null, "57 · Feral Druid", "1 h ago"),
        GuildMemberLine("Selka", "druid", null, "60 · Restoration Druid · item level 62", "2 h ago"),
    )

    private val missing = listOf(
        GuildMemberLine("Brakka", "druid", null, "Feral Druid · item level 60", "2 of 8 nights"),
        GuildMemberLine("Thalrin", "hunter", null, "Marksmanship Hunter · item level 58", "3 of 7 nights"),
    )

    private val lootNight = GuildLootNight(
        "Onyxia's Lair", "Wed 30 Dec",
        listOf(
            GuildLootNight.Drop(18205, "Eskhandar's Collar", "inv_jewelry_necklace_07", 4, "Kraggor", "warrior"),
            GuildLootNight.Drop(17067, "Ancient Cornerstone Grimoire", "inv_misc_book_07", 4, "Thandrel", "warrior", yours = true),
            GuildLootNight.Drop(18422, "Head of Onyxia", "inv_misc_head_dragon_01", 4, null, null),
        ),
    )

    private val close = listOf(
        GuildMemberLine("Vexa", "warlock", null, "Affliction Warlock", "56"),
        GuildMemberLine("Pell", "mage", null, "Frost Mage", "55"),
        GuildMemberLine("Orrin", "paladin", null, "Holy Paladin", "55"),
    )

    private val gearGaps = listOf(
        GuildMemberLine("Selka", "druid", null, "Restoration Druid · item level 62", "41 HPS"),
        GuildMemberLine("Thalrin", "hunter", null, "Marksmanship Hunter · item level 58", "37 DPS"),
        GuildMemberLine("Kaelra", "mage", null, "Frost Mage · item level 61", "22 DPS"),
    )

    private val yourAttendance = GuildMemberLine("Thandrel", "warrior", null, "Fury Warrior", "7 of 8 nights")

    private val logs = listOf(
        GuildLogLine("r3", "UBRS, Sunday", "Blackrock Spire · 7/8 kills · 1 h 32 min", "Sun 3 Jan", clean = false),
        GuildLogLine("r2", "Scholomance, Saturday", "Scholomance · 10/10 kills · 1 h 08 min", "Sat 2 Jan", clean = true),
        GuildLogLine("r1", "Onyxia, Wednesday", "Onyxia's Lair · 1/1 kills · 41 min", "Wed 30 Dec", clean = true),
    )

    private val officerRows = listOf(
        OfficerShortcut(OfficerKind.Plan, "Plan a raid", "Onyxia Wed, Barrow Deeps Sat", null),
        OfficerShortcut(OfficerKind.Applications, "Applications", "Mirelle, Holy Priest 60", "1 waiting"),
        OfficerShortcut(OfficerKind.Setup, "Setup", "Discord bot, raid schedule and addon sync to go", "4 of 8 done"),
    )

    private fun sections(leveler: Boolean, officer: Boolean) = GuildTabSections(
        head = head,
        spread = if (leveler) GuildSpread(31, 4, 36, "your level", own = true) else GuildSpread(31, 4, 24, "raiders", own = false),
        leveler = leveler,
        rosterMeta = "31 of 38 at 60",
        recent = recent,
        close = if (leveler) close else emptyList(),
        logs = logs,
        progression = GuildProgressLine("Onyxia's Lair", "1 of 1 down · 6 pulls"),
        lootNight = if (leveler) null else lootNight,
        officer = if (officer) officerRows else emptyList(),
        missing = if (officer) missing else emptyList(),
        gearGaps = if (officer) gearGaps else emptyList(),
        yourAttendance = if (leveler) null else yourAttendance,
    )

    private val job = RaidJobLine(
        "Onyxia", "whelps_west", "Whelps, west cave with Kaelra and Thalrin",
        "Onyxia · Pick up the west cave's whelps in phases 2 and 3 and kill them away from the raid.",
    )

    private fun picked(name: String, className: String, own: Boolean = false) =
        RaidAssignedName("us/pvp/${name.lowercase()}", name, className, "", "in", own, false)

    private fun slot(key: String, label: String, count: Int, vararg people: RaidAssignedName) =
        AssignmentEncounterLine.Slot(key, label, people.toList(), "${people.size} of $count")

    private val assignments = listOf(
        AssignmentEncounterLine(
            1084, "Onyxia",
            listOf(
                slot("main_tank", "Main tank", 1, picked("Grunvald", "warrior")),
                slot("off_tank", "Off tank (phase 3, fear)", 2, picked("Brakka", "druid")),
                slot(
                    "main_tank_heals", "Main tank healers", 4,
                    picked("Mirethil", "priest"), picked("Selka", "druid"), picked("Hollowmend", "priest"), picked("Aenwyn", "paladin"),
                ),
                slot("off_tank_heals", "Off tank healers", 2),
                slot("fear_cover", "Fear cover", 2, picked("Mirethil", "priest")),
                slot("whelps_west", "Whelps, west cave", 3, picked("Thandrel", "warrior", own = true), picked("Kaelra", "mage"), picked("Thalrin", "hunter")),
                slot("whelps_east", "Whelps, east cave", 3),
            ),
        ),
    )

    /** At 60 (spec 2.2), and for an officer (2.3, 2.3a). */
    fun atCap(officer: Boolean = false) = RaidTabFixtures.guild(officer).copy(
        tab = sections(leveler = false, officer = officer),
        jobs = listOf(job),
        assignments = if (officer) assignments else null,
        pages = GuildTabPresenter.pagesLeft(RaidTabFixtures.pages, leveler = false),
        pagesTitle = "Pages",
    )

    /** Below 60 (spec 2.1): the guild leads, the raid head answers as at 60, one quiet line. */
    fun leveler() = RaidTabFixtures.guild().copy(
        next = RaidTabFixtures.head(SignupState.In).copy(who = "As Felsnap, Damage"),
        ready = emptyList(),
        readyStamp = null,
        foot = emptyList(),
        roles = RaidTabFixtures.guild().roles.map { it.countOnly() },
        tab = sections(leveler = true, officer = false),
        raidsOpen = "Raids open at 60 · you're 36",
        pages = GuildTabPresenter.pagesLeft(RaidTabFixtures.pages, leveler = true),
        pagesTitle = "Pages",
    )

    /** No guild (spec 2.4): Olympus already applied to, Dawnbreak with Apply; under 60 the quiet line. */
    fun noGuild(leveler: Boolean) = RaidTabView.NoGuild(
        if (leveler) "Beast Mastery Hunter" else "Fury Warrior",
        listOf(
            RecruitingGuild(2, "Olympus XXVII", "Alliance · Wed and Sat, 8:00 pm CST · needs damage", applied = "Applied Mon"),
            RecruitingGuild(3, "Dawnbreak", "Alliance · Tue and Thu, 9:00 pm EST · needs damage"),
        ),
        raidsOpen = if (leveler) "Raids open at 60 · you're 36" else null,
    )
}
