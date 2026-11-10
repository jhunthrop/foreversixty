package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.parityItem

/**
 * Material 3's segmented buttons in the round-5 look (spec Mobile, Android): one 48 dp pill with a
 * rule between the segments, the selected one in the teal edge. The raid's four answers. Each segment
 * is a button in the parity inventory, as the mock's answers are.
 */
@Composable
fun FlowSegmented(labels: List<String>, selected: String?, modifier: Modifier = Modifier, onPress: (Int) -> Unit) {
    val shape = RoundedCornerShape(percent = 50)
    Row(
        modifier
            .fillMaxWidth()
            .height(IntrinsicSize.Min)
            .clip(shape)
            .border(1.dp, FsColor.rule, shape),
    ) {
        labels.forEachIndexed { index, label ->
            if (index > 0) Box(Modifier.width(1.dp).fillMaxHeight().background(FsColor.rule))
            val on = label == selected
            Box(
                Modifier
                    .weight(1f)
                    .height(FsTarget.min)
                    .parityItem(ParityItemKind.Button)
                    .background(if (on) FsColor.tealEdge else Color.Transparent)
                    .clickable(role = Role.Button) { onPress(index) }
                    .semantics { this.selected = on },
                contentAlignment = Alignment.Center,
            ) {
                Text(label, style = ScaledType.segment, color = FsColor.strong, maxLines = 1)
            }
        }
    }
}
