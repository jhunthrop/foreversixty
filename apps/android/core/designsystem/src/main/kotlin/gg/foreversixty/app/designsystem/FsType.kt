package gg.foreversixty.app.designsystem

import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.TextUnit
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp

/**
 * The spec's type roles (design/specs/2026-10-30-mobile-apps.md §4.2), in `sp` so each scales
 * with the system font size. JetBrains Mono is not bundled yet (apps/shared/fonts), so [mono]
 * uses the system mono.
 */
object FsType {
    val title = TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingHeavy, fontSize = FsTextSize.h1Phone)
    val h2 = TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingHeavy, fontSize = FsTextSize.h2)
    val stat = TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingHeavy, fontSize = FsTextSize.stat)
    val sec = TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingHeavy, fontSize = FsTextSize.cardTitle)
    val name = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyBold, fontSize = NAME_SIZE.sp)
    val lead = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodySemibold, fontSize = 17.sp, lineHeight = 24.sp)
    val body = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyRegular, fontSize = FsTextSize.body, lineHeight = FsTextSize.bodyLineHeight)
    val bodyStrong = body.copy(fontWeight = FsFontWeight.bodySemibold)
    val tooltip = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyRegular, fontSize = 16.sp, lineHeight = 22.sp)
    val meta = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyRegular, fontSize = FsTextSize.meta, lineHeight = 18.sp)
    val micro = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyBold, fontSize = FsTextSize.micro, letterSpacing = 0.14.em)
    val button = TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyBold, fontSize = 13.sp, letterSpacing = 0.07.em)
    val wordmark = TextStyle(fontFamily = DisplayFamily, fontWeight = FsFontWeight.display, fontSize = 22.sp, letterSpacing = 0.1.em)
    val mono = TextStyle(fontFamily = FontFamily.Monospace, fontWeight = FsFontWeight.bodySemibold, fontSize = FsTextSize.body)

    /** The selector name stops growing at 22 sp (spec §3.4). */
    @Composable
    fun cappedName(): TextStyle {
        val scale = LocalDensity.current.fontScale
        val size: TextUnit = ((NAME_SIZE * scale).coerceAtMost(NAME_CEILING) / scale).sp
        return name.copy(fontSize = size)
    }

    private const val NAME_SIZE = 17f
    private const val NAME_CEILING = 22f
}
