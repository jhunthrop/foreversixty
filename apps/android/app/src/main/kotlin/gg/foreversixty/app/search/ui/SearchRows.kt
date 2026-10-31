package gg.foreversixty.app.search.ui

import androidx.annotation.DrawableRes
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.RowScope
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.search.logic.SearchTarget

/** Row sizes (spec F12.2): item icon 40, the others 36; rows at least 64 tall, so every target is 48+. */
object SearchLayout {
    val itemIcon = 40.dp
    val icon = 36.dp
    val tile = 32.dp
    val rowMinHeight = 64.dp
    val linkRowHeight = 56.dp
    val margin = 16.dp
    val sectionGap = 22.dp
}

/** A section label (22 above it) with an optional trailing action, as `Recent` and its `Clear`. */
@Composable
fun SectionLabel(text: String, trailing: (@Composable () -> Unit)? = null) {
    Row(
        Modifier.fillMaxWidth().padding(start = SearchLayout.margin, end = 4.dp, top = SearchLayout.sectionGap, bottom = 8.dp)
            .heightIn(min = 24.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        MicroLabel(text, Modifier.weight(1f))
        trailing?.invoke()
    }
}

/** A quiet card of rows divided by hairlines (spec F12.1, F12.2). */
@Composable
fun RowCard(count: Int, row: @Composable (Int) -> Unit) {
    Column(Modifier.padding(horizontal = SearchLayout.margin).fillMaxWidth().quietCard()) {
        repeat(count) { index ->
            if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
            row(index)
        }
    }
}

/** A tappable row, merged for TalkBack into one element read in order. */
@Composable
fun TappableRow(onClick: () -> Unit, minHeight: Dp = SearchLayout.rowMinHeight, content: @Composable RowScope.() -> Unit) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = minHeight).clickable(role = Role.Button, onClick = onClick)
            .semantics(mergeDescendants = true) {}
            .padding(horizontal = SearchLayout.margin, vertical = 12.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(14.dp),
    ) { content() }
}

/** The row's title colour: quality for an item, class for a character or guide. */
fun titleColor(target: SearchTarget): Color = when (target) {
    is SearchTarget.Item -> ItemQuality(target.item.quality).text
    is SearchTarget.Character -> classColor(target.character.className)
    is SearchTarget.Guide -> classColor(target.guide.className)
    is SearchTarget.Talent -> FsColor.strong
}

/** The row's leading art: item icon 40, ringed crest 36, spec icon 36 or talent icon 36. */
@Composable
fun TargetIcon(target: SearchTarget, build: String) {
    when (target) {
        is SearchTarget.Item -> ItemIcon(target.item.icon, target.item.quality, build, SearchLayout.itemIcon)
        is SearchTarget.Character -> ClassCrest(target.character.className, SearchLayout.icon)
        is SearchTarget.Guide -> SpecIcon(target.guide.icon, build, SearchLayout.icon)
        is SearchTarget.Talent -> TalentIcon(target.talent.icon, target.talent.name, build, SearchLayout.icon)
    }
}

/** A talent's icon on its tile, with the name's first letters while the art loads. */
@Composable
fun TalentIcon(icon: String, name: String, build: String, size: Dp) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(Modifier.size(size).clip(shape).background(FsColor.cardTop).border(1.dp, GameColors.iconLine, shape)) {
        GameIconImage(icon, build, size, fallback = name.take(2))
    }
}

/** One result or Recent row: art, the coloured title (with `You` on the account's own), the meta. */
@Composable
fun TargetRow(
    target: SearchTarget,
    title: String,
    meta: String,
    build: String,
    titleColor: Color,
    youTag: String? = null,
    chevron: Boolean = false,
    onClick: () -> Unit,
) {
    TappableRow(onClick) {
        TargetIcon(target, build)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            FlowRow(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalArrangement = Arrangement.Center) {
                Text(title, style = FsType.lead.copy(fontWeight = FontWeight.Medium), color = titleColor)
                youTag?.let { Tag(it, FsColor.gold, Modifier.align(Alignment.CenterVertically)) }
            }
            Text(meta, style = FsType.body, color = FsColor.muted)
        }
        if (chevron) Chevron()
    }
}

/** A quick link of the For card: a 32 tile with its glyph, the label, a chevron. */
@Composable
fun LinkRow(@DrawableRes icon: Int, label: String, onClick: () -> Unit) {
    TappableRow(onClick, minHeight = SearchLayout.linkRowHeight) {
        Box(
            Modifier.size(SearchLayout.tile).background(FsColor.cardTop, RoundedCornerShape(FsRadius.control)),
            contentAlignment = Alignment.Center,
        ) {
            Icon(painterResource(icon), contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(18.dp))
        }
        Text(label, style = FsType.lead.copy(fontWeight = FontWeight.Medium), color = FsColor.strong, modifier = Modifier.weight(1f))
        Chevron()
    }
}

@Composable
fun Chevron() {
    Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
}
