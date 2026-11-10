package gg.foreversixty.app.designsystem

import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.TextUnit
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp

/**
 * The spec's type roles (design/specs/2026-10-30-mobile-apps.md §4.2), in `sp` so each scales
 * with the system font size. Round 5 (design/specs/2026-11-09-directions.md, Mobile): headings and
 * the wordmark are Marcellus 400, as on the site; body, labels and figures are Open Sans.
 * JetBrains Mono is not bundled (apps/shared/fonts), so [mono] uses the system mono.
 */
object ScaledType {
    val title = TextStyle(fontFamily = FsFont.title, fontWeight = FontWeight.Normal, fontSize = FsTextSize.h1Phone)
    val h2 = TextStyle(fontFamily = FsFont.title, fontWeight = FontWeight.Normal, fontSize = FsTextSize.h2)
    val stat = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Bold, fontSize = FsTextSize.stat)
    val sec = TextStyle(fontFamily = FsFont.title, fontWeight = FontWeight.Normal, fontSize = FsTextSize.cardTitle)
    /** A raid card's title (spec F4.1). */
    val cardHead = TextStyle(fontFamily = FsFont.title, fontWeight = FontWeight.Normal, fontSize = 22.sp)
    val name = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Bold, fontSize = NAME_SIZE.sp)
    val lead = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.SemiBold, fontSize = 17.sp, lineHeight = 24.sp)
    val body = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Normal, fontSize = FsType.body.fontSize, lineHeight = FsType.body.lineHeight)
    val bodyStrong = body.copy(fontWeight = FontWeight.SemiBold)
    val tooltip = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Normal, fontSize = 16.sp, lineHeight = 22.sp)
    val meta = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Normal, fontSize = FsTextSize.meta, lineHeight = 18.sp)
    val micro = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Bold, fontSize = FsTextSize.micro, letterSpacing = 0.14.em)
    /** A segmented control's label (the sign-up segments, spec F4.3). */
    val segment = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.SemiBold, fontSize = 14.sp)
    val button = TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Bold, fontSize = 13.sp)
    val wordmark = TextStyle(fontFamily = FsFont.title, fontWeight = FontWeight.Normal, fontSize = 22.sp)
    val mono = TextStyle(fontFamily = FontFamily.Monospace, fontWeight = FontWeight.SemiBold, fontSize = FsType.body.fontSize)
    /** Clock digits beside meta text. */
    val monoMeta = TextStyle(fontFamily = FontFamily.Monospace, fontSize = FsTextSize.meta)
    /** A date block's day number. */
    val monoDay = TextStyle(fontFamily = FontFamily.Monospace, fontWeight = FontWeight.SemiBold, fontSize = 20.sp)

    /** The selector name stops growing at 22 sp (spec §3.4). */
    @Composable
    fun cappedName(): TextStyle = capped(name)

    /** The selector's description line, capped with the name so the two lines keep their proportion (spec §3.4). */
    @Composable
    fun cappedMeta(): TextStyle = capped(meta)

    /** [style] grown with the font scale only up to the selector's ceiling (22 sp for the 17 sp name). */
    @Composable
    private fun capped(style: TextStyle): TextStyle {
        val scale = LocalDensity.current.fontScale
        val factor = scale.coerceAtMost(SELECTOR_SCALE_CEILING) / scale
        fun TextUnit.cap(): TextUnit = if (isSp) (value * factor).sp else this
        return style.copy(fontSize = style.fontSize.cap(), lineHeight = style.lineHeight.cap())
    }

    private const val NAME_SIZE = 17f
    private const val NAME_CEILING = 22f
    private const val SELECTOR_SCALE_CEILING = NAME_CEILING / NAME_SIZE
}
