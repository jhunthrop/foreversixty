package gg.foreversixty.app.rankings

import androidx.compose.runtime.CompositionLocalProvider
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.onRoot
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.testing.PARITY_PIXEL_9_TALL
import gg.foreversixty.app.testing.STRICT_SNAPSHOT
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

    /**
     * Owner and visitor at 60 are held to their own goldens, not the mock: the mock's Forever Score
     * section carries the Rating figure from GET .../rating, which api/openapi.yaml (and so the app's
     * client) does not have yet.
     */
    @Test @Config(qualifiers = PARITY_PIXEL_9_TALL)
    fun characterOwner60() = golden("character-owner-60") { Owner60() }

    @Test fun characterOwner60At2() = pushed("character-owner-60", "Home", share = true, fontScale = 2f) { Owner60() }

    @Test fun characterOwner34() = pushed("character-owner-34", "Home", share = true) {
        Page(felsnap, CharacterPhase.Loaded(CharacterFixtures.felsnap), CharacterFixtures.owner(CharacterFixtures.FELSNAP_KEY), CharacterFixtures.felsnapOwnerExtras)
    }

    @Test @Config(qualifiers = PARITY_PIXEL_9_TALL)
    fun characterVisitor60() = golden("character-visitor-60", "Search") {
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

    private fun golden(screen: String, backTo: String = "Home", content: @Composable () -> Unit) {
        compose.setContent {
            ForeverTheme {
                Box(Modifier.fillMaxSize().background(FsColor.bg)) { Shell(backTo, share = true, content) }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/rankings/$screen-and.png", roborazziOptions = STRICT_SNAPSHOT)
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
