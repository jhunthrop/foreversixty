package gg.foreversixty.app.rankings

import androidx.compose.runtime.CompositionLocalProvider
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.rankings.logic.CharacterExtras
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.ui.CharacterContent
import gg.foreversixty.app.rankings.ui.CharacterScreenState
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** The character page (design/mocks/ws_app.py `character-*`, the web character mocks) under its back bar in the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class CharacterParityTest {
    @get:Rule val compose = createComposeRule()

    private val elyra = CharacterRoute("us", "hardcore", "Elyra Duskvale")
    private val felsnap = CharacterRoute("us", "normal", "Felsnap")

    @Test fun characterOwner60() = pushed("character-owner-60", "Home", share = true) { Owner60() }

    @Test fun characterOwner60At2() = pushed("character-owner-60", "Home", share = true, fontScale = 2f) { Owner60() }

    @Test fun characterOwner34() = pushed("character-owner-34", "Home", share = true) {
        Page(felsnap, CharacterPhase.Loaded(CharacterFixtures.felsnap), CharacterFixtures.owner(CharacterFixtures.FELSNAP_KEY), CharacterFixtures.felsnapOwnerExtras)
    }

    @Test fun characterVisitor60() = pushed("character-visitor-60", "Search", share = true) {
        Page(elyra, CharacterPhase.Loaded(CharacterFixtures.elyra), RankingsViewer.SignedOut, CharacterFixtures.elyraVisitorExtras)
    }

    @Test fun characterNoScore() = pushed("character-no-score", "Home") {
        Page(felsnap, CharacterPhase.Loaded(CharacterFixtures.felsnap), RankingsViewer.SignedOut, CharacterFixtures.felsnapVisitorExtras)
    }

    @Test fun characterNotFound() = pushed("character-not-found", "Search") {
        Page(CharacterRoute("us", "hardcore", "Thalrin"), CharacterPhase.NotFound, RankingsViewer.SignedOut)
    }

    @Test fun characterError() = pushed("character-error", "Home") { Page(elyra, CharacterPhase.Failed, RankingsViewer.SignedOut) }

    @Test fun characterLoading() = pushed("character-loading", "Home") { Page(elyra, CharacterPhase.Loading, RankingsViewer.SignedOut) }

    @Composable
    private fun Owner60() =
        Page(elyra, CharacterPhase.Loaded(CharacterFixtures.elyra), CharacterFixtures.owner(CharacterFixtures.ELYRA_KEY), CharacterFixtures.elyraOwnerExtras)

    @Composable
    private fun Page(route: CharacterRoute, phase: CharacterPhase, viewer: RankingsViewer, extras: CharacterExtras = CharacterExtras()) {
        val state = CharacterScreenState(route, phase, null, viewer, CharacterFixtures.items, CharacterFixtures.BUILD, reportsOn = true, extras = extras)
        CharacterContent(state, RankingsTestData.specs(), CharacterFixtures.now, zone = CharacterFixtures.zone)
    }

    @Composable
    private fun Shell(backTo: String, share: Boolean, content: @Composable () -> Unit) {
        // The repo's icons, so every gear slot draws its item.
        CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
            ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}, badges = TabBadges(raid = 1, activity = 3)) {
                Column {
                    FlowBackBar(backTo, {}, if (share) ({}) else null)
                    content()
                }
            }
        }
    }

    private fun pushed(screen: String, backTo: String, share: Boolean = false, fontScale: Float = 1f, content: @Composable () -> Unit) =
        ParityCheck.state(compose, screen, fontScale) { Shell(backTo, share, content) }
}
