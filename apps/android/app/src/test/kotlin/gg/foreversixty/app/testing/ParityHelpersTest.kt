package gg.foreversixty.app.testing

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.material3.Text
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.designsystem.LocalParityRecorder
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRecorder
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config

/** The parity helpers themselves, on a small screen with known regions, before any screen carries them. */
@RunWith(AndroidJUnit4::class)
@Config(qualifiers = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav")
class ParityHelpersTest {
    @get:Rule val compose = createComposeRule()

    private fun render(): ParityRegions.Capture {
        val recorder = ParityRecorder()
        var density = 1f
        compose.setContent {
            density = LocalDensity.current.density
            CompositionLocalProvider(LocalParityRecorder provides recorder) {
                Column(Modifier.parityContent().padding(start = 16.dp)) {
                    Text("Play your class better.", Modifier.width(361.dp).height(40.dp).parityRegion(ParityRegion.HEAD_TITLE))
                    Column(Modifier.parityRegion(ParityRegion.section("Pick your class"))) {
                        listOf("Warrior", "Hunter").forEach { name ->
                            Row(Modifier.width(380.dp).height(56.dp).parityItem(ParityItemKind.Row)) { Text(name) }
                        }
                    }
                }
            }
        }
        compose.waitForIdle()
        return ParityRegions.capture(recorder, density)
    }

    private fun reference(vararg regions: Pair<String, ParityRegions.Region>) =
        ParityRegions.Reference(412f, 915f, mapOf(*regions))

    @Test fun sectionRegionsUseTheWebSlug() {
        assertEquals("section-what-to-stack", ParityRegion.section("What to stack"))
        assertEquals("section-top-3-for-fury-warrior-60", ParityRegion.section("Top 3 for Fury Warrior, 60"))
    }

    @Test fun aMatchingScreenHolds() {
        val want = reference(
            ParityRegion.HEAD_TITLE to ParityRegions.Region(listOf(16f, 0f, 361f, 40f), emptyMap()),
            "section-pick-your-class" to ParityRegions.Region(listOf(16f, 40f, 380f, 112f), mapOf("row" to 2)),
        )
        assertEquals(emptyList<String>(), ParityRegions.compare(render(), want))
    }

    @Test fun driftAndMissingRegionsFail() {
        val want = reference(
            ParityRegion.HEAD_TITLE to ParityRegions.Region(listOf(16f, 30f, 361f, 40f), emptyMap()),
            "section-pick-your-class" to ParityRegions.Region(listOf(16f, 40f, 380f, 112f), mapOf("row" to 3)),
            ParityRegion.COUNTDOWNS to ParityRegions.Region(listOf(16f, 200f, 361f, 80f), emptyMap()),
        )
        assertEquals(
            listOf(
                "countdowns: in the mock, missing from the build",
                "head-title: y 0, the mock has 30 (within 8)",
                "section-pick-your-class: inventory row 2, the mock has row 3",
            ),
            ParityRegions.compare(render(), want),
        )
    }

    @Test fun textReadsInOrder() {
        render()
        assertEquals(listOf("Play your class better.", "Warrior", "Hunter"), ParityText.lines(compose))
    }

    @Test fun regionFilesParse() {
        val parsed = ParityRegions.parse("""{"device":{"width":412,"height":915},"regions":{"head":{"box":[16,0,380,271],"inventory":{"image":1}}}}""")
        assertEquals(ParityRegions.Region(listOf(16f, 0f, 380f, 271f), mapOf("image" to 1)), parsed.regions["head"])
    }
}
