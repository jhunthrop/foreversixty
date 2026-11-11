package gg.foreversixty.app.designsystem

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Typography
import androidx.compose.material3.darkColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.TextUnit

/**
 * Material 3 mapped onto the generated tokens (Tokens.kt: FsColor, FsType, FsTextSize, FsRadius,
 * FsSpace, from design/tokens/tokens.json), in round 5's look (design/specs/2026-11-09-directions.md,
 * Mobile): the painted earth ground (Surfaces.kt `screenGround`), brown-black panels, the teal
 * primary and their blue secondary, the nav-bar teal-black behind the navigation bar. The site is dark only, so the app is too. The faces are FsFont's
 * (Open Sans for body, labels and figures; Marcellus for titles), from apps/shared/fonts through
 * the forever.shared-fonts plugin (res/font).
 */
private val ForeverColors = darkColorScheme(
    primary = FsColor.teal,
    onPrimary = FsColor.strong,
    primaryContainer = FsColor.tealEdge,
    onPrimaryContainer = FsColor.strong,
    secondary = FsColor.blue,
    onSecondary = FsColor.strong,
    secondaryContainer = FsColor.cardTop,
    onSecondaryContainer = FsColor.gold,
    tertiary = FsColor.gold,
    // The deepest well: a screen that paints nothing of its own; the painted ground is
    // `Modifier.screenGround()`, taken screen by screen as each tab lane restyles it.
    background = FsColor.bg,
    onBackground = FsColor.text,
    surface = FsColor.raised,
    onSurface = FsColor.text,
    surfaceVariant = FsColor.cardTop,
    onSurfaceVariant = FsColor.muted,
    surfaceContainer = FsColor.navBar,
    outline = FsColor.line,
    error = FsColor.wipe,
)

private fun styled(family: FontFamily, weight: FontWeight, size: TextUnit) =
    TextStyle(fontFamily = family, fontWeight = weight, fontSize = size)

private fun heading(size: TextUnit) = styled(FsFont.title, FontWeight.Normal, size)

private fun body(size: TextUnit) = styled(FsFont.body, FontWeight.Normal, size)

private val ForeverTypography = Typography(
    headlineMedium = heading(FsTextSize.h1Phone),
    headlineSmall = heading(FsTextSize.h2),
    // The M3 top app bar's title (the mock's Android bar: Marcellus 22).
    titleLarge = FsType.barTitle,
    titleMedium = heading(FsTextSize.cardTitle),
    bodyLarge = body(FsType.body.fontSize).copy(lineHeight = FsType.body.lineHeight),
    bodyMedium = body(FsType.body.fontSize).copy(lineHeight = FsType.body.lineHeight),
    labelLarge = body(FsTextSize.meta).copy(fontWeight = FontWeight.SemiBold),
    labelMedium = body(FsTextSize.meta).copy(fontWeight = FontWeight.SemiBold),
    labelSmall = body(FsTextSize.micro).copy(fontWeight = FontWeight.SemiBold),
)

@Composable
fun ForeverTheme(content: @Composable () -> Unit) {
    MaterialTheme(colorScheme = ForeverColors, typography = ForeverTypography, content = content)
}
