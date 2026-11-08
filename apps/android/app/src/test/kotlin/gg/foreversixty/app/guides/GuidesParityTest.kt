package gg.foreversixty.app.guides

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.guides.logic.TierChoice
import gg.foreversixty.app.guides.ui.GuidesHome
import gg.foreversixty.app.guides.ui.GuidesHomeActions
import gg.foreversixty.app.guides.ui.GuidesHomeState
import gg.foreversixty.app.guides.ui.TierListActions
import gg.foreversixty.app.guides.ui.TierListContent
import gg.foreversixty.app.guides.ui.TierListState
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.TierRow
import gg.foreversixty.app.model.TiersFile
import gg.foreversixty.app.guides.ui.specTitle
import java.util.Locale
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * The Guides and Tiers tools (design/mocks/ws_app.py `guides`, `guides-you`, `tiers`, `tiers-you`: the
 * site's phone layout pushed under Home from the menu) against their regions, from the W3-C1 fixtures.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class GuidesParityTest {
    @get:Rule val compose = createComposeRule()

    private val index = GuidesTestData.index()

    @Test fun guides() = state("guides") { Guides(null) }

    @Test fun guidesAt2() = state("guides", fontScale = 2f) { Guides(null) }

    @Test fun guidesYou() = state("guides-you") { Guides(GuidesParityFixtures.felsnap) }

    @Test fun tiers() = state("tiers") { Tiers(null) }

    @Test fun tiersAt2() = state("tiers", fontScale = 2f) { Tiers(null) }

    @Test fun tiersYou() = state("tiers-you") { Tiers(GuidesParityFixtures.thandrel) }

    @Composable
    private fun Guides(character: AccountCharacter?) {
        GuidesHome(
            GuidesHomeState(Loadable.Loaded(index, null), emptyMap(), character, character?.let { Guides.yourGuide(index, it) }),
            GuidesHomeActions(),
        )
    }

    @Composable
    private fun Tiers(character: AccountCharacter?) {
        TierListContent(
            TierListState(
                Loadable.Loaded(GuidesParityFixtures.tiers(), null),
                TierChoice(TierChoice.DPS, TierChoice.ALLIANCE),
                character?.let { Guides.yourGuide(index, it)?.spec },
                index,
                signedIn = character != null,
            ),
            GuidesTestData.now,
            TierListActions(openBis = { _, _ -> }, openPlanner = {}, openSim = { _, _ -> }, openChecks = {}, signIn = {}),
        )
    }

    /** Pushed under Home from the menu: the Back bar over the page, inside the tab shell. */
    private fun state(screen: String, fontScale: Float = 1f, content: @Composable () -> Unit) =
        ParityCheck.state(compose, screen, fontScale) {
            CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) {
                    Column {
                        FlowBackBar(BACK, {})
                        content()
                    }
                }
            }
        }

    private companion object {
        const val BACK = "Home"
    }
}

/**
 * The mocks' data: Felsnap (Beast Mastery Hunter 34, Dwarf) for Guides; Thandrel (Fury Warrior 60) for
 * Tiers, over the mock's Alliance DPS list (ws_tiers.py at build 1.60.1.70790) in the contract fixture.
 */
object GuidesParityFixtures {
    /** The mock's list, row by row: spec id, race, DPS, gap, letter, tie and the bar's share. */
    private val MOCK_DPS = listOf(
        MockRow("hunter-marksmanship", "Night Elf", 1083.6, 0.0, "S", null, 1.0),
        MockRow("warrior-fury", "Night Elf", 1004.2, 7.3, "A", null, 0.926),
        MockRow("hunter-survival", "Night Elf", 968.5, 10.6, "B", "below", 0.893),
        MockRow("hunter-beast-mastery", "Night Elf", 958.8, 11.5, "B", "below", 0.886),
        MockRow("warrior-arms", "Night Elf", 954.2, 11.9, "B", "below", 0.882),
        MockRow("druid-feral", "Night Elf", 947.6, 12.6, "B", "above", 0.875),
        MockRow("shaman-enhancement", "Dwarf", 922.3, 14.9, "B", null, 0.849),
        MockRow("rogue-assassination", "Night Elf", 911.5, 15.9, "B", null, 0.842),
        MockRow("rogue-combat", "Gnome", 889.5, 17.9, "B", null, 0.82),
        MockRow("rogue-subtlety", "Night Elf", 878.6, 18.9, "B", null, 0.812),
        MockRow("warlock-destruction", "Human", 868.9, 19.8, "B", null, 0.801),
        MockRow("paladin-retribution", "Human", 853.7, 21.2, "C", null, 0.787),
        MockRow("warlock-affliction", "Human", 840.4, 22.4, "C", null, 0.776),
        MockRow("mage-fire", "Human", 810.9, 25.2, "C", "below", 0.746),
        MockRow("mage-arcane", "Gnome", 803.9, 25.8, "C", "above", 0.743),
        MockRow("warlock-demonology", "Human", 795.2, 26.6, "C", null, 0.735),
        MockRow("priest-shadow", "Human", 783.4, 27.7, "C", null, 0.724),
        MockRow("shaman-elemental", "Dwarf", 731.4, 32.5, "D", "below", 0.676),
        MockRow("druid-balance", "Night Elf", 725.5, 33.0, "D", "above", 0.669),
        MockRow("mage-frost", "Gnome", 679.1, 37.3, "D", null, 0.629),
    )

    private class MockRow(val spec: String, val race: String, val dps: Double, val gap: Double, val tier: String, val tie: String?, val fraction: Double)

    /** The contract fixture with the mock's Alliance DPS list (names and icons from the guides index). */
    fun tiers(): TiersFile {
        val file = GuidesTestData.tiers()
        val index = GuidesTestData.index()
        val rows = MOCK_DPS.mapIndexed { i, mock ->
            val spec = requireNotNull(Guides.spec(index, mock.spec)) { "no ${mock.spec} in the index" }
            val cls = requireNotNull(Guides.classOf(index, mock.spec))
            TierRow(
                rank = i + 1,
                spec = mock.spec,
                classSlug = cls.classSlug,
                specSlug = spec.specSlug,
                name = specTitle(spec.specName, cls.className),
                race = mock.race,
                icon = spec.icon,
                metric = mock.dps,
                metricText = String.format(Locale.US, "%,.1f", mock.dps),
                unit = "DPS",
                gapPercent = mock.gap,
                gapText = if (i == 0) "Top" else String.format(Locale.US, "−%.1f%%", mock.gap),
                fraction = mock.fraction,
                tie = mock.tie,
                tier = mock.tier,
                lowConfidence = false,
                title = "",
                detail = "",
            )
        }
        val alliance = file.factions.alliance
        return file.copy(factions = file.factions.copy(alliance = alliance.copy(dps = alliance.dps.copy(rows = rows))))
    }

    val felsnap = AccountCharacter(
        key = "us/normal/felsnap",
        region = "us",
        ruleset = "normal",
        name = "Felsnap",
        source = "bnet",
        className = "hunter",
        level = 34,
        faction = "alliance",
        race = "Dwarf",
        spec = "Beast Mastery",
    )

    /** Fury Warrior, 2nd of 20 on the mock's list. */
    val thandrel = AccountCharacter(
        key = "us/normal/thandrel",
        region = "us",
        ruleset = "normal",
        name = "Thandrel",
        source = "bnet",
        className = "warrior",
        level = 60,
        faction = "alliance",
        race = "Night Elf",
        spec = "Fury",
    )
}
