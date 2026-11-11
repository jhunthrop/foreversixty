package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.ui.test.getBoundsInRoot
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithContentDescription
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.compose.ui.unit.height
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.designsystem.FsTarget
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

/** Round 5 rows 1 and 4: the not-synced row and the Covered `and N more` are 48 dp taps. */
@RunWith(AndroidJUnit4::class)
class RaidTabFootTest {
    @get:Rule val compose = createComposeRule()

    private val keys = setOf("us/pvp/kaelra", "us/pvp/thalrin")
    private val cover = CoverLine(listOf(CoverItem("Kings", null), CoverItem("Might", null), CoverItem("Wisdom", null)), emptyList(), more = 2)
    private val foot = listOf(
        RoleFoot("Missing Kings · Might · Wisdom and 2 more", cover = cover),
        RoleFoot("Not synced this week: Kaelra, Thalrin", onlyKeys = keys),
    )

    private fun draw(rosterOf: (Set<String>) -> Unit = {}, onCover: () -> Unit = {}) =
        compose.setContent { Column { RoleRows(emptyList(), foot, rosterOf, onCover) } }

    @Test fun `the not-synced row opens the roster on the raiders it names, 48 dp tall`() {
        val opened = mutableListOf<Set<String>>()
        draw(rosterOf = { opened += it })
        val row = compose.onNodeWithText("Not synced this week: Kaelra, Thalrin", useUnmergedTree = true)
        row.performClick()
        assertEquals(listOf(keys), opened)
        compose.onNodeWithText("Not synced this week: Kaelra, Thalrin").let { node ->
            assertTrue(node.getBoundsInRoot().height >= FsTarget.min)
        }
    }

    @Test fun `the Covered line with names left out opens the cover list`() {
        var opened = 0
        draw(onCover = { opened++ })
        compose.onNodeWithContentDescription("Missing Kings · Might · Wisdom and 2 more").performClick()
        assertEquals(1, opened)
        assertTrue(compose.onNodeWithContentDescription("Missing Kings · Might · Wisdom and 2 more").getBoundsInRoot().height >= FsTarget.min)
    }
}
