package gg.foreversixty.app.designsystem

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Typography
import androidx.compose.material3.darkColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.Font
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.TextUnit

/**
 * Material 3 mapped onto the generated tokens (Tokens.kt: FsColor, FsFontWeight, FsTextSize,
 * FsRadius, FsSpacing, from web/src/styles/tokens.css). The site is dark only, so the app is too.
 * The faces come from apps/shared/fonts through the forever.shared-fonts plugin (res/font).
 */
private val ForeverColors = darkColorScheme(
    primary = FsColor.gold,
    onPrimary = FsColor.bg,
    primaryContainer = FsColor.goldDeep,
    onPrimaryContainer = FsColor.strong,
    secondary = FsColor.gold,
    onSecondary = FsColor.bg,
    secondaryContainer = FsColor.cardTop,
    onSecondaryContainer = FsColor.gold,
    tertiary = FsColor.ember,
    background = FsColor.bg,
    onBackground = FsColor.text,
    surface = FsColor.raised,
    onSurface = FsColor.text,
    surfaceVariant = FsColor.cardTop,
    onSurfaceVariant = FsColor.muted,
    surfaceContainer = FsColor.raised,
    outline = FsColor.line,
    error = FsColor.wipe,
)

/** Cinzel: the wordmark only. */
val DisplayFamily = FontFamily(Font(R.font.cinzel_bold, FsFontWeight.display))

/** Barlow Condensed: headings and titles. */
val HeadingFamily = FontFamily(
    Font(R.font.barlowcondensed_bold, FsFontWeight.headingBold),
    Font(R.font.barlowcondensed_extrabold, FsFontWeight.headingHeavy),
)

/** Barlow: body copy, labels and meta lines. */
val BodyFamily = FontFamily(
    Font(R.font.barlow_regular, FsFontWeight.bodyRegular),
    Font(R.font.barlow_semibold, FsFontWeight.bodySemibold),
    Font(R.font.barlow_bold, FsFontWeight.bodyBold),
)

private fun heading(size: TextUnit) =
    TextStyle(fontFamily = HeadingFamily, fontWeight = FsFontWeight.headingBold, fontSize = size)

private fun body(size: TextUnit) =
    TextStyle(fontFamily = BodyFamily, fontWeight = FsFontWeight.bodyRegular, fontSize = size)

private val ForeverTypography = Typography(
    headlineMedium = heading(FsTextSize.h1Phone),
    headlineSmall = heading(FsTextSize.h2),
    titleLarge = heading(FsTextSize.h2),
    titleMedium = heading(FsTextSize.cardTitle),
    bodyLarge = body(FsTextSize.body).copy(lineHeight = FsTextSize.bodyLineHeight),
    bodyMedium = body(FsTextSize.body).copy(lineHeight = FsTextSize.bodyLineHeight),
    labelLarge = body(FsTextSize.meta).copy(fontWeight = FsFontWeight.bodySemibold),
    labelMedium = body(FsTextSize.meta).copy(fontWeight = FsFontWeight.bodySemibold),
    labelSmall = body(FsTextSize.micro).copy(fontWeight = FsFontWeight.bodySemibold),
)

@Composable
fun ForeverTheme(content: @Composable () -> Unit) {
    MaterialTheme(colorScheme = ForeverColors, typography = ForeverTypography, content = content)
}
