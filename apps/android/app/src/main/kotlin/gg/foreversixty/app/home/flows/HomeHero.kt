package gg.foreversixty.app.home.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Shadow
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.BandKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSize
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowBandPage
import gg.foreversixty.app.flows.FlowPageState
import gg.foreversixty.app.flows.lineBox
import gg.foreversixty.app.home.LARGE_TEXT_SCALE

/**
 * Home's page in the round-5 look (spec Mobile, mock m-android-home): the camp-road hero behind the
 * bar and the head, then the blocks on the painted ground, 16 under the hero, 14 at the sides and 20
 * apart. Parity regions are measured from the bottom of [topBar], as the mock's are from its nav bar.
 */
@Composable
fun HomePage(topBar: @Composable () -> Unit, page: FlowPageState, head: @Composable () -> Unit, content: @Composable ColumnScope.() -> Unit) {
    FlowBandPage(topBar, BandKind.Hero, head, page = page, content = content)
}

/**
 * One head on every Home (spec 4.1): the class crest, the name in the class colour in the display
 * face and the fact line, on the hero. Without a [className] (no character yet) the title stands alone.
 */
@Composable
fun HomeHeroHead(title: String, facts: List<String>, className: String?) {
    Column(Modifier.fillMaxWidth().padding(start = FsSpace.space4, end = FsSpace.space4, top = HEAD_TOP, bottom = HEAD_BOTTOM).parityRegion(ParityRegion.HEAD)) {
        Row(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE), verticalAlignment = Alignment.CenterVertically) {
            if (className != null) {
                ClassCrest(className, FsSize.crest56, Modifier.parityItem(ParityItemKind.Image))
                Spacer(Modifier.width(FsSpace.space3))
            }
            Column(Modifier.weight(1f)) {
                val color = if (className != null) classColor(className) else FsColor.cream
                // Marcellus at 30 is taller than its 34 line; the box keeps the role's line (heroName 30/34).
                Text(title, style = NAME, color = color, modifier = Modifier.lineBox(NAME.lineHeight).semantics { heading() })
                if (facts.isNotEmpty()) {
                    val lines = heroFactLines(facts, stacked = LocalDensity.current.fontScale >= LARGE_TEXT_SCALE)
                    Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_META)) {
                        lines.forEach { Text(it, style = FACTS, color = FsColor.text) }
                    }
                }
            }
        }
    }
}

/**
 * The hero's fact lines: one line on the dot (`Level 58 Fury Warrior · 2 levels to the 60 list`), or at
 * large text one fact to a line, so the break falls at the dot and never strands `· 2` (round 4 row 8).
 */
internal fun heroFactLines(facts: List<String>, stacked: Boolean): List<String> =
    if (stacked) facts else listOf(facts.joinToString(FACT_SEPARATOR))

private val HEAD_TOP = 6.dp
private val HEAD_BOTTOM = 18.dp
private const val FACT_SEPARATOR = " · "

/** The display face over the art carries the mock's drop shadow (`0 1px 2px` black at .45). */
private val TITLE_SHADOW = Shadow(Color.Black.copy(alpha = 0.45f), Offset(0f, 2f), 4f)
private val NAME = FsType.heroName.copy(shadow = TITLE_SHADOW)
private val FACTS = FsType.factLine

/** A heading on a panel or the ground (the raid card's zone, a section's title) wears the same shadow. */
val HOME_CARD_TITLE = FsType.cardTitle.copy(shadow = TITLE_SHADOW)
