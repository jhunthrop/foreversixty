package gg.foreversixty.app.common

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Check
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Shape
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.BodyFamily
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsFontWeight
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsTextSize
import gg.foreversixty.app.designsystem.GamePalette

// The site's shared controls drawn natively (spec §4.5, tenet 10). One definition each; a feature
// never draws its own variant.

/** The `micro` role: 11 bold uppercase with tracking, for tags and section labels. */
val MicroStyle = TextStyle(
    fontFamily = BodyFamily,
    fontWeight = FsFontWeight.bodyBold,
    fontSize = FsTextSize.micro,
    letterSpacing = 1.4.sp,
)

/** The `.pill` / `.tag`: a 22 tall capsule, 11 bold uppercase, tinted by [color]. */
@Composable
fun Tag(text: String, color: Color, modifier: Modifier = Modifier, checked: Boolean = false) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Row(
        modifier
            .heightIn(min = 22.dp)
            .background(color.copy(alpha = 0.12f), shape)
            .border(1.dp, color.copy(alpha = 0.35f), shape)
            .padding(horizontal = 8.dp, vertical = 3.dp),
        horizontalArrangement = Arrangement.spacedBy(4.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        if (checked) Icon(Icons.Filled.Check, contentDescription = null, tint = color, modifier = Modifier.size(11.dp))
        Text(text.uppercase(), style = MicroStyle, color = color)
    }
}

/** The secondary button: a 4.5% white fill and the gold 17% border, 48 tall. */
@Composable
fun SecondaryButton(text: String, onClick: () -> Unit, modifier: Modifier = Modifier) {
    OutlinedButton(
        onClick = onClick,
        modifier = modifier.fillMaxWidth().heightIn(min = 48.dp),
        shape = RoundedCornerShape(FsRadius.control),
        border = BorderStroke(1.dp, FsColor.lineCard),
        colors = ButtonDefaults.outlinedButtonColors(containerColor = GamePalette.controlFill, contentColor = FsColor.strong),
    ) {
        Text(text.uppercase(), style = MicroStyle.copy(fontSize = 13.sp))
    }
}

/** A chip or pill's ground: [fill] and a 1 dp [line] border in [shape]. */
fun Modifier.pillSurface(fill: Color, line: Color, shape: Shape = RoundedCornerShape(percent = 50)): Modifier =
    background(fill, shape).border(1.dp, line, shape)
