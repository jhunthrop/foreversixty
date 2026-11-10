package gg.foreversixty.app.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.height
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.semantics.SemanticsProperties
import androidx.compose.ui.test.SemanticsMatcher
import androidx.compose.ui.test.assert
import androidx.compose.ui.test.assertIsEnabled
import androidx.compose.ui.test.assertIsNotEnabled
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.click
import androidx.compose.ui.test.performTouchInput
import androidx.compose.ui.unit.Density
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.ItemLook
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

/** The shared primitives' fixes: the disabled answers, the padded inline link, the 1.5x item row. */
@RunWith(AndroidJUnit4::class)
class FlowPrimitivesTest {
    @get:Rule val compose = createComposeRule()

    private val labels = listOf("In", "Late", "Bench", "Out")

    @Test
    fun `disabled answers read as disabled and take no press`() {
        val pressed = mutableListOf<Int>()
        compose.setContent { FlowSegmented(labels, "In", enabled = false) { pressed += it } }
        labels.forEach { compose.onNodeWithText(it).assertIsNotEnabled() }
        compose.onNodeWithText("Bench").performClick()
        assertTrue(pressed.isEmpty())
    }

    @Test
    fun `enabled answers press`() {
        val pressed = mutableListOf<Int>()
        compose.setContent { FlowSegmented(labels, "In") { pressed += it } }
        compose.onNodeWithText("Bench").assertIsEnabled().performClick()
        assertEquals(listOf(2), pressed)
    }

    @Test
    fun `the link line is one control that takes a tap within the 48 dp target`() {
        var taps = 0
        compose.setContent {
            Column {
                Spacer(Modifier.height(FsTarget.min))
                FlowLinkLine("60 raid · Alliance", FlowSubLink("Change") { taps += 1 }, FsType.small, FsColor.text)
                Spacer(Modifier.height(FsTarget.min))
            }
        }
        val line = compose.onNode(hasText("Change", substring = true))
        line.assert(SemanticsMatcher.keyIsDefined(SemanticsProperties.Role))
        val bounds = line.fetchSemanticsNode().boundsInRoot
        val height = bounds.height
        val reach = with(compose.density) { (FsTarget.min.toPx() - height) / 2 - 1.dp.toPx() }
        assertTrue("the line draws under the target: $height", reach > 0)
        // Above the drawn line, inside the 48 dp target: still the link.
        compose.onRoot().performTouchInput { click(Offset(bounds.center.x, bounds.top - reach)) }
        assertEquals(1, taps)
    }

    @Test
    fun `from 1_5x an item row's gain moves under its name`() {
        val look = ItemLook(id = 1, name = "Precisely Calibrated Boomstick", quality = 3, icon = null)
        compose.setContent {
            val density = LocalDensity.current
            CompositionLocalProvider(LocalDensity provides Density(density.density, LARGE)) {
                FlowItemRow(look, "1", "Onyxia's Lair", "+14.5", "DPS") {}
            }
        }
        val name = compose.onNodeWithText(look.name.orEmpty(), useUnmergedTree = true).fetchSemanticsNode().boundsInRoot
        val gain = compose.onNodeWithText("+14.5", substring = true, useUnmergedTree = true).fetchSemanticsNode().boundsInRoot
        assertTrue("gain $gain under name $name", gain.top >= name.bottom)
    }

    private companion object {
        const val LARGE = 1.5f
    }
}
