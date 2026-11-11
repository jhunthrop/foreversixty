package gg.foreversixty.app.designsystem

import androidx.compose.foundation.BorderStroke
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
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

/**
 * The site's buttons in round 5 (web lib/ui/button.ts): primary their flat teal with the teal edge
 * (Buy Now, one per screen at most), secondary their blue (Upgrade Now), the accent gold outline
 * (the reminder ask's Turn on), destructive the wipe outline and quiet the rule outline, each a
 * white or tone label in sentence case on the 6 dp control radius. Pressed: the Material ripple.
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
    /** The leading icon at the label's size, so it grows with the font scale (16 at 1x). */
    iconScalesWithText: Boolean = false,
) {
    val iconSize = if (iconScalesWithText) with(LocalDensity.current) { ICON_TEXT_SIZE.toDp() } else ICON_SIZE
    val shape = RoundedCornerShape(FsRadius.field)
    val size = if (fill) modifier.fillMaxWidth().heightIn(min = height) else modifier.heightIn(min = height)
    // The vertical padding keeps a wrapped label off the border at large text; one line still sits at [height].
    val padding = PaddingValues(horizontal = HORIZONTAL_PADDING, vertical = VERTICAL_PADDING)
    val content: @Composable () -> Unit = {
        if (busy) {
            CircularProgressIndicator(Modifier.size(20.dp), color = contentColor(kind), strokeWidth = 2.dp)
        } else {
            if (leadingIcon != null) {
                Icon(leadingIcon, contentDescription = null, modifier = Modifier.size(iconSize))
                Spacer(Modifier.width(FsSpace.space2))
            }
            Text(label, style = ScaledType.button, textAlign = TextAlign.Center)
        }
    }
    // A busy button keeps its colours; a disabled one dims.
    val disabled = if (busy) contentColor(kind) else contentColor(kind).copy(alpha = DISABLED_ALPHA)
    val fill = containerColor(kind)
    if (fill != null) {
        // Disabled (not busy), the fill and its edge dim to the disabled alpha so it reads as unavailable.
        val alpha = if (enabled || busy) 1f else DISABLED_ALPHA
        Button(
            onClick = onClick,
            enabled = enabled && !busy,
            shape = shape,
            contentPadding = padding,
            border = BorderStroke(1.dp, borderColor(kind).copy(alpha = alpha)),
            modifier = size,
            colors = ButtonDefaults.buttonColors(
                containerColor = fill,
                contentColor = contentColor(kind),
                disabledContainerColor = fill.copy(alpha = alpha),
                disabledContentColor = if (busy) contentColor(kind) else contentColor(kind).copy(alpha = DISABLED_ALPHA),
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
                containerColor = Color.Transparent,
                contentColor = contentColor(kind),
                disabledContainerColor = Color.Transparent,
                disabledContentColor = disabled,
            ),
        ) { content() }
    }
}

private val HORIZONTAL_PADDING = 18.dp
private val VERTICAL_PADDING = 12.dp
private val ICON_SIZE = 16.dp
private val ICON_TEXT_SIZE = 16.sp
private const val DISABLED_ALPHA = 0.5f

/** The filled kinds' ground; the outlined kinds have none. The kinds' colours are shared with the flows' small button. */
fun containerColor(kind: BrandButtonKind): Color? = when (kind) {
    BrandButtonKind.Primary -> FsColor.teal
    BrandButtonKind.Secondary -> FsColor.blue
    BrandButtonKind.Accent, BrandButtonKind.Destructive, BrandButtonKind.Quiet -> null
}

fun contentColor(kind: BrandButtonKind): Color = when (kind) {
    BrandButtonKind.Primary, BrandButtonKind.Secondary, BrandButtonKind.Quiet -> FsColor.strong
    BrandButtonKind.Accent -> FsColor.gold
    BrandButtonKind.Destructive -> FsColor.wipe
}

fun borderColor(kind: BrandButtonKind): Color = when (kind) {
    BrandButtonKind.Primary -> FsColor.tealEdge
    BrandButtonKind.Secondary -> FsColor.blueEdge
    BrandButtonKind.Quiet -> FsColor.rule
    BrandButtonKind.Destructive -> FsColor.wipe.copy(alpha = 0.45f)
    BrandButtonKind.Accent -> FsColor.gold
}
