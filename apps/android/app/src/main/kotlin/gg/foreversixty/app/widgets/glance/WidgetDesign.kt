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
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.widgets.logic.TagKind

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

    fun color(color: Color): ColorProvider = ColorProvider(color)

    /**
     * A widget's cells do not grow with the font, so the type follows the system's size up to
     * [MAX_SCALE] and holds there (font scale 2.0 would otherwise push the countdown out of a 4 x 2).
     */
    private const val MAX_SCALE = 1.3f

    fun styles(fontScale: Float): WidgetStyles = WidgetStyles(if (fontScale > MAX_SCALE) MAX_SCALE / fontScale else 1f)

    fun tagColor(kind: TagKind): Color = when (kind) {
        // Kill is for `In · DPS` alone; Late, Bench and Out are muted words (spec §5.6).
        TagKind.In -> FsColor.kill
        TagKind.NotSignedUp, TagKind.ChooseSpec -> FsColor.gold
        TagKind.Cancelled -> FsColor.wipe
        TagKind.Late, TagKind.Bench, TagKind.Out, TagKind.Locked -> FsColor.muted
    }

    /**
     * Tracked caps for the micro labels (`N E X T  R A I D` on the board): Glance's TextStyle has
     * no letter-spacing, so a hair space sits between the letters. The widget's description carries
     * the plain words for TalkBack.
     */
    fun tracked(text: String): String = text.uppercase().toList().joinToString("\u200A")

    fun tagBackground(kind: TagKind): Int = when (kind) {
        TagKind.In -> R.drawable.widget_tag_kill
        TagKind.NotSignedUp, TagKind.ChooseSpec -> R.drawable.widget_tag_gold
        TagKind.Cancelled -> R.drawable.widget_tag_wipe
        TagKind.Late, TagKind.Bench, TagKind.Out, TagKind.Locked -> R.drawable.widget_tag_muted
    }
}

/** The widget type roles at one scale ([fit] shrinks sp so the system scale stops at the cap). */
internal class WidgetStyles(private val fit: Float) {
    private fun style(size: Float, color: Color, weight: FontWeight = FontWeight.Normal, family: FontFamily = FontFamily.SansSerif) =
        TextStyle(color = WidgetDesign.color(color), fontSize = (size * fit).sp, fontWeight = weight, fontFamily = family)

    /** `NEXT RAID`: micro, gold, tracked caps. */
    val micro = style(11f, FsColor.gold, FontWeight.Bold)

    /** The zone: `sec` 22, condensed caps. */
    val zone = style(22f, FsColor.strong, FontWeight.Bold, FontFamily("sans-serif-condensed"))
    val title = style(16f, FsColor.strong, FontWeight.Medium)
    val body = style(13f, FsColor.text)
    val muted = style(12f, FsColor.muted)
    val countdown = style(14f, FsColor.gold, FontWeight.Medium, FontFamily.Monospace)
    val count = style(16f, FsColor.strong, FontWeight.Medium, FontFamily.Monospace)
    val gain = style(13f, FsColor.text, FontWeight.Medium, FontFamily.Monospace)
    val link = style(14f, FsColor.gold, FontWeight.Medium)

    fun tag(kind: TagKind): TextStyle = style(11f, WidgetDesign.tagColor(kind), FontWeight.Bold)

    fun button(checked: Boolean, failed: Boolean): TextStyle =
        style(14f, if (failed) FsColor.wipe else if (checked) FsColor.strong else FsColor.nav, FontWeight.Medium)

    fun quality(color: Color, size: Float = 14f): TextStyle = style(size, color, FontWeight.Medium)
}

/** The type roles at the device's font scale, capped for the widget's fixed cells. */
@Composable
internal fun widgetStyles(): WidgetStyles = WidgetDesign.styles(LocalContext.current.resources.configuration.fontScale)
