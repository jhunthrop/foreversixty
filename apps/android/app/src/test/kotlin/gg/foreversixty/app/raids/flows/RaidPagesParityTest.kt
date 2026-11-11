package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.ui.Modifier
import gg.foreversixty.app.designsystem.screenGround
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import androidx.compose.runtime.CompositionLocalProvider
import gg.foreversixty.app.items.LocalItemIconPainter
import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.flows.FlowFigure
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

/** The screens pushed under the Raid tab (design/mocks/ws_app_flows.py `raid_unanswered`, `applicant`), inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class RaidPagesParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun unanswered() = pushed("raid-unanswered", "Raid", share = true) {
        NotAnsweredScreen(RaidPagesFixtures.unanswered, BUILD, NotAnsweredActions())
    }

    @Test fun applicant() = pushed("raid-applicant", "Raid") {
        ApplicantScreen(RaidPagesFixtures.applicant(), ApplicantActions())
    }

    @Test fun applicantAccepted() = pushed("raid-applicant-accepted", "Raid") {
        ApplicantScreen(RaidPagesFixtures.applicant(accepted = true), ApplicantActions())
    }

    @Test fun nightLive() = pushed("raid-night-live", "Raid", share = true) {
        RaidNightPageScreen(RaidPagesFixtures.night(answered = true), BUILD, RaidNightPageActions())
    }

    @Test fun nightSignup() = pushed("raid-night-signup", "Raid", share = true) {
        RaidNightPageScreen(RaidPagesFixtures.night(answered = false), BUILD, RaidNightPageActions())
    }

    private fun pushed(
        screen: String,
        backTo: String,
        share: Boolean = false,
        drift: Set<String> = emptySet(),
        content: @Composable () -> Unit,
    ) =
        ParityCheck.state(compose, screen, expectedDrift = drift) {
            // The repo's game icons, as the tab roots' parity tests paint them (never an empty square).
            CompositionLocalProvider(LocalItemIconPainter provides gg.foreversixty.app.home.HomeFixtures.repoIcons, LocalRaidIconBuild provides gg.foreversixty.app.home.HomeFixtures.BUILD) {
                ForeverTabScaffold(selected = ForeverTab.Raid, onSelect = {}, badges = TabBadges(raid = 6, activity = 3)) {
                    Column(Modifier.fillMaxSize().screenGround()) {
                        FlowBackBar(backTo, {}, if (share) ({}) else null)
                        content()
                    }
                }
            }
        }

    private companion object {
        const val BUILD = "1.60.1.70790"
    }
}

/** The mock's officer page: healers short by 3, two healers unanswered, Grunvald marked In, Aurelith on the bench. */
object RaidPagesFixtures {
    /** Onyxia tonight in 42 minutes (ws_app.py `raid_night`): you are In, or not signed up yet. */
    fun night(answered: Boolean) = RaidNightPageContent(
        title = "Onyxia's Lair",
        icon = "inv_misc_head_dragon_black",
        facts = listOf("Olympus XXVII", "Wed 6 Jan, 8:00 pm CST"),
        eyebrow = "Olympus XXVII · Onyxia's Lair tonight",
        clock = listOf(ClockPart("00", "hours"), ClockPart("42", "min")),
        clockLabel = "0 hours 42 minutes",
        live = true,
        aside = "Wed 6 Jan, 8:00 pm CST · 18 in, 3 late",
        tag = if (answered) "You are In" else "Not signed up",
        answered = answered,
        selected = if (answered) SignupState.In else null,
        lateOpen = false,
        roles = listOf(
            "Tanks 2 of 2" to "Grunvald, Brakka",
            "Healers 5 of 8" to "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen · 3 short",
            "Damage 11 of 30" to "Thandrel, Kaelra, Thalrin and 8 more",
        ),
        ready = listOf(
            ReadyItem(ReadyTone.Ok, "Drakefire Amulet", "Attuned to Onyxia's Lair", null),
            ReadyItem(ReadyTone.Ok, "Elixir of the Mongoose", "3 in your bags, from the addon's last export", null),
            ReadyItem(ReadyTone.Miss, "Flask of the Titans", "None in your bags", "Auction house"),
            ReadyItem(ReadyTone.Warn, "Boots not enchanted", "", null),
        ),
        readyStamp = "From your last log-out, 9 h ago",
    )

    /** Mirelle, Holy Priest 60, applied 2 min ago (mock `applicant`; the API has no guild rank, so the history shows raids and the last one). */
    fun applicant(accepted: Boolean = false) = ApplicantContent(
        name = "Mirelle",
        className = "priest",
        facts = listOf("Holy Priest · 60", "Alliance Dwarf", "applied 2 min ago"),
        figures = listOf(FlowFigure("Forever Score", "74"), FlowFigure("Sim", "612", "HPS"), FlowFigure("Item level", "63")),
        ready = listOf(
            ApplicantLine("Attuned to Onyxia's Lair", "Drakefire Amulet"),
            ApplicantLine("Enchants", "6 of 8 slots · cloak and boots missing"),
            ApplicantLine("Best in slot gap", "9 pieces from the Holy Priest 60 raid list"),
        ),
        history = listOf(
            ApplicantLine("Dawnbreak", "14 raids · last raid Mon 2 Nov"),
            ApplicantLine("Ashen Vow", "6 raids · last raid Sun 18 Oct"),
        ),
        note = "Holy main since launch. Wednesday and Sunday nights work for me.",
        accepted = if (accepted) {
            AcceptedContent(
                "Accepted into Olympus XXVII. Invite Mirelle in game:",
                "/ginvite Mirelle",
                "Mirelle shows on the roster once the addon or Battle.net sees the character in the guild.",
            )
        } else {
            null
        },
    )

    private fun row(name: String, cls: String, sub: String, role: String, marked: SignupState? = null) =
        NotAnsweredRow("us/pvp/${name.lowercase()}", name, cls, sub, role, marked, "https://discord.com/users/${name.length}")

    val unanswered = NotAnsweredContent(
        title = "Onyxia's Lair",
        icon = "inv_misc_head_dragon_black",
        facts = listOf("Olympus XXVII", "Wed 6 Jan, 8:00 pm CST", "healers short by 3"),
        chips = listOf(
            NotAnsweredChip(NotAnsweredFilter.NotAnswered, "Not answered 4"),
            NotAnsweredChip(NotAnsweredFilter.In, "In 19"),
            NotAnsweredChip(NotAnsweredFilter.Late, "Late 3"),
            NotAnsweredChip(NotAnsweredFilter.Bench, "Bench 2"),
            NotAnsweredChip(NotAnsweredFilter.Out, "Out 3"),
            NotAnsweredChip(NotAnsweredFilter.All, "All 31"),
        ),
        filter = NotAnsweredFilter.NotAnswered,
        mention = MentionContent(
            role = "Mention the 2 healers in Discord",
            all = "Mention all 4",
            preview = "@Mirethil @Selka: Onyxia's Lair is Wed 6 Jan, 8:00 pm CST and we need healers. Answer In, Late, Bench or Out in the app or on the raid page.",
        ),
        listTitle = "Not answered",
        listMeta = "4 of 31, healers first",
        rows = listOf(
            row("Mirethil", "priest", "Holy Priest · was In last week", "healer"),
            row("Selka", "druid", "Restoration Druid · was Late last week", "healer"),
            row("Grunvald", "warrior", "Protection Warrior · was In last week", "tank", SignupState.In),
            row("Varro", "rogue", "Combat Rogue · missed last week", "dps"),
            row("Ysolde", "mage", "Fire Mage · new this week", "dps"),
        ),
        actionable = true,
        benchTitle = "Healers on the bench",
        bench = listOf(row("Aurelith", "paladin", "Holy Paladin · on the bench", "healer")),
        benchNote = "If Mirethil, Selka and Aurelith are In, healers are full.",
    )
}
