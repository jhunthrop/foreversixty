package gg.foreversixty.app.planner

import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.semantics.SemanticsActions
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.longClick
import androidx.compose.ui.test.onNodeWithContentDescription
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTouchInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorState
import gg.foreversixty.app.planner.logic.EditorView
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.planner.ui.EditorActions
import gg.foreversixty.app.planner.ui.EditorPageState
import gg.foreversixty.app.planner.ui.PlannerEditorPage
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** A cell adds on tap, removes on long press, and TalkBack gets both as named actions (spec F13.1). */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav")
class PlannerAccessibilityTest {
    @get:Rule
    val compose = createComposeRule()

    private val shaman = PlannerTestData.shaman()
    private val index = TalentIndex(shaman)
    private val convection = index.trees[0].talents.first { it.name == "Convection" }
    private val added = mutableListOf<Int>()
    private val removed = mutableListOf<Int>()
    private val opened = mutableListOf<Int>()

    private fun show() {
        val state = EditorState("shaman", "orc", 60, emptyList())
        val phase = EditorPhase.Ready(EditorView(shaman, index, "Shaman", emptyList(), state))
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    PlannerEditorPage(EditorPageState(phase, BisTestData.BUILD, 0), EditorActions(add = { added += it }, remove = { removed += it }, openTalent = { opened += it }))
                }
            }
        }
    }

    private val cell get() = compose.onNodeWithContentDescription("Convection, rank 0 of 5")

    @Test fun `tap adds and long press removes`() {
        show()
        cell.performClick()
        cell.performTouchInput { longClick() }
        assertEquals(listOf(convection.id), added)
        assertEquals(listOf(convection.id), removed)
    }

    @Test fun `TalkBack offers Add a point and Remove a point`() {
        show()
        val actions = cell.fetchSemanticsNode().config[SemanticsActions.CustomActions]
        assertEquals(listOf("Add a point", "Remove a point"), actions.map { it.label })
        compose.runOnUiThread { actions.forEach { it.action() } }
        assertEquals(listOf(convection.id), added)
        assertEquals(listOf(convection.id), removed)
    }

    @Test fun `the strip's words open the rank sheet`() {
        show()
        compose.onNodeWithText(convection.name, substring = true).performClick()
        assertEquals(listOf(convection.id), opened)
    }

    @Test fun `the class crest names the class`() {
        show()
        compose.onNodeWithContentDescription("Shaman").assertExists()
    }
}
