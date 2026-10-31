package gg.foreversixty.app.account

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.CheckCircle
import androidx.compose.material.icons.outlined.Info
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType

/** The toast at the top of F1.1: failures in wipe red, a finished deletion neutral. */
@Composable
fun NoticeToast(text: String, modifier: Modifier = Modifier, isFailure: Boolean = true) {
    val border = if (isFailure) FsColor.wipe.copy(alpha = 0.45f) else FsColor.line
    Row(
        modifier
            .fillMaxWidth()
            .background(FsColor.raised, RoundedCornerShape(FsRadius.panel))
            .border(1.dp, border, RoundedCornerShape(FsRadius.panel))
            .padding(FsSpacing.space4)
            .semantics { liveRegion = LiveRegionMode.Assertive },
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Icon(
            if (isFailure) Icons.Outlined.Info else Icons.Outlined.CheckCircle,
            contentDescription = null,
            tint = if (isFailure) FsColor.wipe else FsColor.gold,
        )
        Text(text, style = FsType.body, color = FsColor.strong)
    }
}
