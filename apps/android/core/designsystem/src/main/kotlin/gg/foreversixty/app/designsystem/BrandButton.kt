package gg.foreversixty.app.designsystem

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import java.util.Locale

/**
 * The site's buttons (spec §4.5): `PRIMARY_BUTTON_FIXED` (the gold gradient, one per screen at
 * most), `SECONDARY_BUTTON_FIXED` (4.5% white, the gold 17% border), destructive and quiet, and the
 * accent secondary (gold label and border, the reminder ask's Turn on). Pressed: the Material ripple.
 */
enum class BrandButtonKind { Primary, Secondary, Accent, Destructive, Quiet }

/** One button for every kind: full width by default; `fill = false` hugs the label. */
@Composable
fun BrandButton(
    label: String,
    kind: BrandButtonKind,
    onClick: () -> Unit,
    modifier: Modifier = Modifier,
    busy: Boolean = false,
    enabled: Boolean = true,
    height: Dp = 52.dp,
    fill: Boolean = true,
    leadingIcon: ImageVector? = null,
) {
    val shape = RoundedCornerShape(FsRadius.control)
    val size = if (fill) modifier.fillMaxWidth().heightIn(min = height) else modifier.heightIn(min = height)
    val padding = PaddingValues(horizontal = HORIZONTAL_PADDING)
    val content: @Composable () -> Unit = {
        if (busy) {
            CircularProgressIndicator(Modifier.size(20.dp), color = contentColor(kind), strokeWidth = 2.dp)
        } else {
            if (leadingIcon != null) {
                Icon(leadingIcon, contentDescription = null, modifier = Modifier.size(16.dp))
                Spacer(Modifier.width(FsSpacing.space2))
            }
            Text(caseOf(kind, label), style = FsType.button, textAlign = TextAlign.Center)
        }
    }
    // A busy button keeps its colours; a disabled one dims.
    val disabled = if (busy) contentColor(kind) else contentColor(kind).copy(alpha = DISABLED_ALPHA)
    if (kind == BrandButtonKind.Primary) {
        Button(
            onClick = onClick,
            enabled = enabled && !busy,
            shape = shape,
            contentPadding = padding,
            modifier = size.background(Brush.verticalGradient(listOf(FsColor.goldHover, FsColor.gold)), shape),
            colors = ButtonDefaults.buttonColors(
                containerColor = Color.Transparent,
                contentColor = FsColor.bg,
                disabledContainerColor = Color.Transparent,
                disabledContentColor = FsColor.bg,
            ),
        ) { content() }
    } else {
        OutlinedButton(
            onClick = onClick,
            enabled = enabled && !busy,
            shape = shape,
            contentPadding = padding,
            modifier = size,
            border = BorderStroke(1.dp, if (enabled || busy) borderColor(kind) else borderColor(kind).copy(alpha = DISABLED_ALPHA)),
            colors = ButtonDefaults.outlinedButtonColors(
                containerColor = GameColors.controlFill,
                contentColor = contentColor(kind),
                disabledContainerColor = GameColors.controlFill,
                disabledContentColor = disabled,
            ),
        ) { content() }
    }
}

private val HORIZONTAL_PADDING = 18.dp
private const val DISABLED_ALPHA = 0.38f

/** Spec §4.5: the primary button is uppercase; a secondary reads in sentence case (BiS UX verdict). */
private fun caseOf(kind: BrandButtonKind, label: String): String =
    if (kind == BrandButtonKind.Secondary || kind == BrandButtonKind.Accent) label else label.uppercase(Locale.ROOT)

private fun contentColor(kind: BrandButtonKind): Color = when (kind) {
    BrandButtonKind.Primary -> FsColor.bg
    BrandButtonKind.Secondary -> FsColor.strong
    BrandButtonKind.Accent -> FsColor.gold
    BrandButtonKind.Destructive -> FsColor.wipe
    BrandButtonKind.Quiet -> FsColor.nav
}

private fun borderColor(kind: BrandButtonKind): Color = when (kind) {
    BrandButtonKind.Destructive -> FsColor.wipe.copy(alpha = 0.45f)
    BrandButtonKind.Accent -> FsColor.gold
    else -> FsColor.lineCard
}
