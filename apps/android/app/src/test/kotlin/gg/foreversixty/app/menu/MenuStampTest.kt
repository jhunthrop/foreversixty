package gg.foreversixty.app.menu

import androidx.compose.foundation.layout.width
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.test.assertIsNotEnabled
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.assertHeightIsAtLeast
import androidx.compose.ui.unit.Density
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.isStale
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

/** The drawer's Battle.net stamp (row 12, a4): the in-flight state, the stale tone, Refresh on its own row at large text. */
@RunWith(AndroidJUnit4::class)
class MenuStampTest {
    @get:Rule val compose = createComposeRule()

    private val model = MenuFixtures.model(listOf(MenuFixtures.thandrel))

    private fun menu(refresh: MenuRefresh = MenuRefresh.Idle, scale: Float = 1f) {
        compose.setContent {
            val density = LocalDensity.current
            CompositionLocalProvider(LocalDensity provides Density(density.density, scale)) {
                ProfileMenu(model, MenuFixtures.lines, MenuActions(), Modifier.width(DRAWER), refresh = refresh)
            }
        }
    }

    @Test fun `a refresh under way reads Refreshing and takes no second press`() {
        menu(MenuRefresh.Refreshing)
        compose.onNodeWithText("Refreshing…", substring = true).assertIsNotEnabled()
    }

    @Test fun `from 1_5x Refresh takes its own 48 dp row under the stamp`() {
        menu(scale = LARGE)
        val stamp = compose.onNodeWithText("Updated", substring = true).fetchSemanticsNode().boundsInRoot
        val refresh = compose.onNodeWithText("Refresh").assertHeightIsAtLeast(FsTarget.min).fetchSemanticsNode()
        assertTrue("refresh ${refresh.boundsInRoot} under stamp $stamp", refresh.boundsInRoot.top >= stamp.bottom)
    }

    @Test fun `a stamp past a week reads stale`() {
        assertEquals(false, RelativeAge.Days(7).isStale())
        assertEquals(true, RelativeAge.Days(8).isStale())
        assertEquals(false, RelativeAge.Hours(20).isStale())
    }

    private companion object {
        const val LARGE = 1.5f
        val DRAWER = 350.dp
    }
}
