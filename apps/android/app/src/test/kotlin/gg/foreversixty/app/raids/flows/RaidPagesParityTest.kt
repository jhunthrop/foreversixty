package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
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

    private fun pushed(screen: String, backTo: String, share: Boolean = false, content: @Composable () -> Unit) =
        ParityCheck.state(compose, screen) {
            ForeverTabScaffold(selected = ForeverTab.Raid, onSelect = {}, badges = TabBadges(raid = 6, activity = 3)) {
                Column {
                    FlowBackBar(backTo, {}, if (share) ({}) else null)
                    content()
                }
            }
        }

    private companion object {
        const val BUILD = "1.60.1.70790"
    }
}

/** The mock's officer page: healers short by 3, two healers unanswered, Grunvald marked In, Aurelith on the bench. */
object RaidPagesFixtures {
    private fun row(name: String, cls: String, sub: String, role: String, marked: SignupState? = null) =
        NotAnsweredRow("us/pvp/${name.lowercase()}", name, cls, sub, role, marked)

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
