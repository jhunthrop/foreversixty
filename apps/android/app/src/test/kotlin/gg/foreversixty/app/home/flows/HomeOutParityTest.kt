package gg.foreversixty.app.home.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.activity.ActivityFixtures
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.model.FeedSegment
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

/** Signed-out Home (design/mocks/ws_app.py `home-out`, the web home R3) inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class HomeOutParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun homeOut() = state()

    @Test fun homeOutAt2() = state(fontScale = 2f)

    private fun state(fontScale: Float = 1f) = ParityCheck.state(compose, "home-out", fontScale = fontScale) {
        ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}, badges = TabBadges()) {
            Column {
                CompositionLocalProvider(LocalOpenMenu provides {}) { ProfileBar(null) }
                HomeOutScreen(HomeOutFixtures.view, signingIn = false, actions = HomeOutActions())
            }
        }
    }
}

/** The mock's scenario: 32 days to the first raids, the site feed's seven lines. */
object HomeOutFixtures {
    private const val BUILD = ActivityFixtures.BUILD

    private val classes = listOf(
        ClassPick("warrior", "Warrior", "Arms, Fury, Protection"),
        ClassPick("paladin", "Paladin", "Holy, Protection, Retribution"),
        ClassPick("hunter", "Hunter", "Beast Mastery, Marksmanship, Survival"),
        ClassPick("rogue", "Rogue", "Assassination, Combat, Subtlety"),
        ClassPick("priest", "Priest", "Discipline, Holy, Shadow"),
        ClassPick("shaman", "Shaman", "Elemental, Enhancement, Restoration"),
        ClassPick("mage", "Mage", "Arcane, Fire, Frost"),
        ClassPick("warlock", "Warlock", "Affliction, Demonology, Destruction"),
        ClassPick("druid", "Druid", "Balance, Feral, Feral Bear, Restoration"),
    )

    private val raids = FirstRaidsCard(
        "First raids open · Wed 9 Dec",
        listOf(ClockPart("32", "days"), ClockPart("06", "hours"), ClockPart("00", "min")),
        "32 days 6 hours 0 minutes",
        HomeOutComposer.FIRST_RAIDS,
    )

    private fun name(text: String, cls: String) = FeedSegment(text, style = "name", className = cls)
    private fun strong(text: String) = FeedSegment(text, style = "strong")
    private fun item(text: String, quality: Int) = FeedSegment(text, style = "item", quality = quality)
    private fun plain(text: String) = FeedSegment(text)

    private fun line(id: String, lead: FlowLead, segments: List<FeedSegment>, time: String, kind: String) =
        FeedLine(id, lead, HomeOutComposer.sentence(segments), time, kind, "/")

    private val feed = listOf(
        line("1", FlowLead.Crest("warlock"), listOf(name("Nightfall", "warlock"), plain(" uploaded a full Shadowfang Keep, 8 bosses in 41 min")), "4 min ago", "Logs"),
        line("2", FlowLead.Crest("warrior"), listOf(strong("Fury Warrior, 20 to 29:"), plain(" "), name("Grukk", "warrior"), plain(" climbed 6 places to #3")), "17 min ago", "Rankings"),
        line(
            "3",
            FlowLead.Item(ItemLook(0, null, 3, "inv_chest_leather_07"), BUILD),
            listOf(strong("Best in slot changed"), plain(" for Fury Warrior 20 to 29: "), item("Blackened Defias Armor", 3), plain(" is the new chest")),
            "1 h ago",
            "Best in slot",
        ),
        line("4", FlowLead.Crest("hunter"), listOf(name("Ellowen", "hunter"), plain(" uploaded Razorfen Kraul, 5 bosses in 52 min")), "1 h ago", "Logs"),
        line("5", FlowLead.Crest("warrior"), listOf(name("Brakka", "warrior"), plain(" took #1 for Arms Warrior, 30 to 39")), "2 h ago", "Rankings"),
        line("6", FlowLead.Crest("priest"), listOf(name("Saelis", "priest"), plain(" uploaded The Deadmines, 6 bosses in 38 min")), "3 h ago", "Logs"),
        line(
            "7",
            FlowLead.Item(ItemLook(0, null, 1, "spell_fire_flamebolt"), BUILD),
            listOf(strong("Best in slot refreshed"), plain(" for Fire Mage 40 to 49: 3 new picks")),
            "5 h ago",
            "Best in slot",
        ),
    )

    val view = HomeOutView(listOf("Pick your class.", "First raids 9 Dec"), classes, raids, HomeFeed.Rows(feed), BUILD)
}
