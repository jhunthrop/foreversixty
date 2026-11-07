package gg.foreversixty.app.designsystem

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Typography
import androidx.compose.material3.darkColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.TextUnit

/**
 * Material 3 mapped onto the generated tokens (Tokens.kt: FsColor, FsType, FsTextSize, FsRadius,
 * FsSpace, from design/tokens/tokens.json). The site is dark only, so the app is too.
 * The faces are FsFont's (Barlow, and Marcellus for titles), from apps/shared/fonts through the
 * forever.shared-fonts plugin (res/font); headings are Barlow 700, as on the site.
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

private fun heading(size: TextUnit) =
    TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Bold, fontSize = size)

private fun body(size: TextUnit) =
    TextStyle(fontFamily = FsFont.body, fontWeight = FontWeight.Normal, fontSize = size)

private val ForeverTypography = Typography(
    headlineMedium = heading(FsTextSize.h1Phone),
    headlineSmall = heading(FsTextSize.h2),
    titleLarge = heading(FsTextSize.h2),
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
