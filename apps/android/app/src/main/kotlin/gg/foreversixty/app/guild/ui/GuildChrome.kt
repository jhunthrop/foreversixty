package gg.foreversixty.app.guild.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawWithContent
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType

/** The pushed screen's bar: back, the title in the heading face, trailing actions. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun GuildTopBar(title: String, back: () -> Unit, actions: @Composable () -> Unit = {}) {
    TopAppBar(
        title = { Text(title.uppercase(), style = FsType.sec, color = FsColor.strong, modifier = Modifier.semantics { heading() }) },
        navigationIcon = {
            IconButton(onClick = back) { Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = null, tint = FsColor.strong) }
        },
        actions = { actions() },
        colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
        // The app scaffold already pads for the status bar.
        windowInsets = WindowInsets(0, 0, 0, 0),
    )
}

/**
 * One row of a quiet card drawn item by item, so a long list stays lazy: the first row takes the
 * card's top corners and border, the last its bottom, and every other row ends in the divider.
 */
fun Modifier.cardSegment(first: Boolean, last: Boolean): Modifier {
    val radius = FsRadius.panel
    val shape = RoundedCornerShape(
        topStart = if (first) radius else 0.dp,
        topEnd = if (first) radius else 0.dp,
        bottomStart = if (last) radius else 0.dp,
        bottomEnd = if (last) radius else 0.dp,
    )
    return this
        .padding(horizontal = 16.dp)
        .fillMaxWidth()
        .clip(shape)
        .background(FsColor.raised)
        .drawWithContent {
            drawContent()
            val stroke = 1.dp.toPx()
            val r = radius.toPx()
            // The card's outline, extended past the edges a middle row does not own; the clip cuts it.
            val top = if (first) stroke / 2 else -2 * r
            val bottom = if (last) size.height - stroke / 2 else size.height + 2 * r
            drawRoundRect(
                color = FsColor.lineSoft,
                topLeft = Offset(stroke / 2, top),
                size = Size(size.width - stroke, bottom - top),
                cornerRadius = CornerRadius(r),
                style = Stroke(stroke),
            )
            if (!last) drawLine(FsColor.lineSoft, Offset(0f, size.height - stroke / 2), Offset(size.width, size.height - stroke / 2), stroke)
        }
}

/** A grey bar at a skeleton row's real size. */
@Composable
fun SkeletonBar(width: Dp, height: Dp) {
    Box(Modifier.width(width).height(height).background(FsColor.lineSoft, RoundedCornerShape(4.dp)))
}

/** The gold tile and a list row: the Raids tab's Roster and Raid history, You › Calendar. */
@Composable
fun GuildLinkRow(icon: Painter, title: String, subtitle: String?, onClick: () -> Unit) {
    Row(
        horizontalArrangement = Arrangement.spacedBy(12.dp),
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 64.dp)
            .clickable(onClick = onClick)
            .padding(horizontal = 14.dp, vertical = 10.dp)
            .semantics(mergeDescendants = true) {},
    ) {
        GoldTile(icon)
        Column(Modifier.weight(1f)) {
            Text(title, style = FsType.bodyStrong, color = FsColor.strong)
            subtitle?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
        }
        Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted, modifier = Modifier.clearAndSetSemantics {})
    }
}

/** The 36 dp gold-washed tile behind a row's glyph. */
@Composable
fun GoldTile(icon: Painter) {
    Box(
        contentAlignment = Alignment.Center,
        modifier = Modifier.size(36.dp).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.control)),
    ) { Icon(icon, contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(18.dp)) }
}
