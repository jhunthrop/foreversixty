package gg.foreversixty.app.designsystem

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Typography
import androidx.compose.material3.darkColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.text.TextStyle

/**
 * Material 3 mapped onto the house tokens. The site is dark only, so the app is too.
 * TODO(lane T): Barlow, Barlow Condensed and Cinzel (wordmark) font families once the
 * generator ships them; the system sans stands in until then.
 */
private val ForeverColors = darkColorScheme(
    primary = Tokens.Colors.Gold,
    onPrimary = Tokens.Colors.Bg,
    primaryContainer = Tokens.Colors.GoldDeep,
    onPrimaryContainer = Tokens.Colors.Strong,
    secondary = Tokens.Colors.Gold,
    onSecondary = Tokens.Colors.Bg,
    secondaryContainer = Tokens.Colors.CardTop,
    onSecondaryContainer = Tokens.Colors.Gold,
    tertiary = Tokens.Colors.Ember,
    background = Tokens.Colors.Bg,
    onBackground = Tokens.Colors.Text,
    surface = Tokens.Colors.Raised,
    onSurface = Tokens.Colors.Text,
    surfaceVariant = Tokens.Colors.CardTop,
    onSurfaceVariant = Tokens.Colors.Muted,
    surfaceContainer = Tokens.Colors.Raised,
    outline = Tokens.Colors.Line,
    error = Tokens.Colors.Wipe,
)

private val ForeverTypography = Typography(
    headlineMedium = TextStyle(fontSize = Tokens.TypeScale.H1),
    headlineSmall = TextStyle(fontSize = Tokens.TypeScale.H2),
    titleMedium = TextStyle(fontSize = Tokens.TypeScale.CardTitle),
    bodyLarge = TextStyle(fontSize = Tokens.TypeScale.Body, lineHeight = Tokens.TypeScale.BodyLineHeight),
    bodyMedium = TextStyle(fontSize = Tokens.TypeScale.Body, lineHeight = Tokens.TypeScale.BodyLineHeight),
    labelMedium = TextStyle(fontSize = Tokens.TypeScale.Meta),
    labelSmall = TextStyle(fontSize = Tokens.TypeScale.Micro),
)

@Composable
fun ForeverTheme(content: @Composable () -> Unit) {
    MaterialTheme(colorScheme = ForeverColors, typography = ForeverTypography, content = content)
}
