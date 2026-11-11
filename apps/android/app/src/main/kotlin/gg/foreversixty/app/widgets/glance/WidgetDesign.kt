package gg.foreversixty.app.widgets.glance

import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.Color
import androidx.glance.LocalContext
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.DpSize
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.glance.text.FontFamily
import androidx.glance.text.FontWeight
import androidx.glance.text.TextStyle
import androidx.glance.unit.ColorProvider
import gg.foreversixty.app.designsystem.FsColor

/**
 * The home-screen widget container and its type at widget sizes (spec §4 `Home-screen widget`):
 * the card gradient, a 17% gold border, radius 28, and only v1's type roles. A widget draws in the
 * launcher's process, so the brand faces fall back to the system's sans and condensed families.
 */
internal object WidgetDesign {
    val small = DpSize(110.dp, 110.dp)
    val medium = DpSize(250.dp, 110.dp)

    val padding: Dp = 14.dp
    val buttonHeight: Dp = 48.dp
    val sealSize: Dp = 16.dp

    /** The gap between the card's lines. */
    val cardGap: Dp = 3.dp

    fun color(color: Color): ColorProvider = ColorProvider(color)

    /**
     * A widget's cells do not grow with the font, so the type follows the system's size up to
     * [MAX_SCALE] and holds there (font scale 2.0 would otherwise push the countdown out of a 4 x 2).
     */
    private const val MAX_SCALE = 1.3f

    fun styles(fontScale: Float): WidgetStyles = WidgetStyles(if (fontScale > MAX_SCALE) MAX_SCALE / fontScale else 1f, 1f / fontScale)

}

/**
 * The widget type roles at one scale ([fit] shrinks sp so the system scale stops at the cap;
 * [unscaled] undoes the system scale for the big figure, already 30 and the one line that would
 * otherwise push the card's line out of a 4 x 2).
 */
internal class WidgetStyles(private val fit: Float, private val unscaled: Float) {
    private fun style(size: Float, color: Color, weight: FontWeight = FontWeight.Normal, family: FontFamily = FontFamily.SansSerif) =
        TextStyle(color = WidgetDesign.color(color), fontSize = (size * fit).sp, fontWeight = weight, fontFamily = family)

    /** `Next raid`: the 12 label, gold, sentence case (ruling N1). */
    val micro = style(12f, FsColor.gold, FontWeight.Bold)

    val title = style(16f, FsColor.strong, FontWeight.Medium)
    val body = style(13f, FsColor.text)
    val muted = style(12f, FsColor.muted)
    val link = style(14f, FsColor.gold, FontWeight.Medium)

    /** The card's four lines (mocks `widgets-guild`, `-noguild`, `-open`): label 12, big 30, name 15, line 13. */
    val cardKey = style(12f, FsColor.muted, FontWeight.Bold)
    val cardName = style(15f, FsColor.strong, FontWeight.Bold)
    val cardLine = style(13f, FsColor.muted)

    fun cardBig(color: Color = FsColor.strong): TextStyle =
        TextStyle(color = WidgetDesign.color(color), fontSize = (BIG * unscaled).sp, fontWeight = FontWeight.Bold, fontFamily = FontFamily.SansSerif)

    fun button(checked: Boolean, failed: Boolean): TextStyle =
        style(14f, if (failed) FsColor.wipe else if (checked) FsColor.strong else FsColor.nav, FontWeight.Medium)

    private companion object {
        const val BIG = 30f
    }
}

/** The type roles at the device's font scale, capped for the widget's fixed cells. */
@Composable
internal fun widgetStyles(): WidgetStyles = WidgetDesign.styles(LocalContext.current.resources.configuration.fontScale)
