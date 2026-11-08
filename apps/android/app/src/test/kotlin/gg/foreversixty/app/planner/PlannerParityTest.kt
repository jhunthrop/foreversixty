package gg.foreversixty.app.planner

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.planner.ui.EditorActions
import gg.foreversixty.app.planner.ui.EditorPageState
import gg.foreversixty.app.planner.ui.PlannerEditorPage
import gg.foreversixty.app.planner.ui.SharedActions
import gg.foreversixty.app.planner.ui.SharedBuildPage
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** The planner and a shared build pushed from the menu under Home (ws_app.py `planner-*`, `build-shared`). */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class PlannerParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun plannerYou() = ParityCheck.state(compose, "planner-you") { Editor(PlannerPageFixtures.you) }

    @Test fun plannerYouAt2() = ParityCheck.state(compose, "planner-you", fontScale = 2f) { Editor(PlannerPageFixtures.you) }

    @Test fun plannerGuide() = ParityCheck.state(compose, "planner-guide") { Editor(PlannerPageFixtures.guide) }

    @Test fun plannerLoading() = ParityCheck.state(compose, "planner-loading") { Editor(PlannerPageFixtures.loading) }

    @Test fun plannerError() = ParityCheck.state(compose, "planner-error") { Editor(PlannerPageFixtures.failed) }

    @Test fun buildShared() = ParityCheck.state(compose, "build-shared") {
        Pushed { SharedBuildPage(PlannerPageFixtures.shared, SharedActions()) }
    }

    @Composable
    private fun Editor(state: EditorPageState) = Pushed { PlannerEditorPage(state, EditorActions()) }

    @Composable
    private fun Pushed(content: @Composable () -> Unit) {
        CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
            ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}, badges = TabBadges()) { content() }
        }
    }
}
