package gg.foreversixty.app.bis.spec

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
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

/** A spec's list pushed under Gear for a player not signed in (ws_app.py `bis-spec-*`), inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class BisSpecParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun bisSpec60() = ParityCheck.state(compose, "bis-spec-60") { Pushed { Page(BisSpecFixtures.sixty, BisSpecView.BySlot) } }

    @Test fun bisSpec60At2() = ParityCheck.state(compose, "bis-spec-60", fontScale = 2f) { Pushed { Page(BisSpecFixtures.sixty, BisSpecView.BySlot) } }

    @Test fun bisSpecHealer60() = ParityCheck.state(compose, "bis-spec-healer-60") { Pushed { Page(BisSpecFixtures.healer, BisSpecView.BySlot) } }

    @Test fun bisSpec20() = ParityCheck.state(compose, "bis-spec-20") { Pushed { Page(BisSpecFixtures.twenty, BisSpecView.WhereToGo) } }

    @Composable
    private fun Page(content: BisSpecContent, view: BisSpecView) {
        BisSpecPage(content, view, BisSpecFixtures.BUILD, panelOpen = false, onPanel = {}, actions = BisSpecActions())
    }

    @Composable
    private fun Pushed(content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Gear, onSelect = {}, badges = TabBadges()) {
            Column {
                FlowBackBar("Gear", {}, {})
                content()
            }
        }
    }
}
