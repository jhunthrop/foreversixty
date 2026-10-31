package gg.foreversixty.app.home

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Check
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import java.util.Locale

/**
 * F2.7 Your gear: all 17 slots in character-frame order, worn then the pick. [open] shows an
 * item's tooltip; its second value is the item worn in that slot. Scrolling is the caller's.
 */
@Composable
fun GearList(rows: List<GearRow>, intro: String, build: String, open: (ItemLook, Int?) -> Unit = { _, _ -> }) {
    Column(Modifier.fillMaxWidth().padding(16.dp), verticalArrangement = Arrangement.spacedBy(16.dp)) {
        Text(
            stringResource(R.string.gear_title).uppercase(Locale.ROOT),
            style = FsType.title,
            color = FsColor.strong,
            modifier = Modifier.semantics { heading() },
        )
        Text(intro, style = FsType.body, color = FsColor.muted)
        Column(Modifier.fillMaxWidth().quietCard()) {
            rows.forEachIndexed { index, row ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                GearRowView(row, build, open)
            }
        }
    }
}

@Composable
private fun GearRowView(row: GearRow, build: String, open: (ItemLook, Int?) -> Unit) {
    val stacked = LocalDensity.current.fontScale >= LARGE_TEXT_SCALE
    val label = @Composable { MicroLabel(row.slotName, modifier = if (stacked) Modifier else Modifier.width(SLOT_LABEL_WIDTH)) }
    val content = @Composable {
        Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(12.dp)) {
            Worn(row, build, open, Modifier.weight(1f))
            Trailing(row, build, open)
        }
    }
    if (stacked) {
        Column(Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 12.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
            label()
            content()
        }
    } else {
        Row(
            Modifier.fillMaxWidth().heightIn(min = 72.dp).padding(horizontal = 16.dp, vertical = 12.dp),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(12.dp),
        ) {
            label()
            Column(Modifier.weight(1f)) { content() }
        }
    }
}

@Composable
private fun Worn(row: GearRow, build: String, open: (ItemLook, Int?) -> Unit, modifier: Modifier) {
    val worn = row.worn
    val pickArrow = row.pick?.let { stringResource(R.string.items_pickArrow, it.name ?: stringResource(R.string.items_notInData)) }
    val spoken = listOfNotNull(row.slotName, row.wornName, row.wornLevel, pickArrow).joinToString(", ")
    val tap = if (worn != null) Modifier.clickable(role = Role.Button) { open(worn, worn.id) } else Modifier
    Row(
        modifier.then(tap).semantics(mergeDescendants = true) { contentDescription = spoken },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ItemIcon(worn?.icon, worn?.quality, build, 36.dp, dashed = worn?.icon == null)
        Column(verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(row.wornName, style = FsType.meta, color = if (worn == null) FsColor.muted else ItemQuality(worn.quality).text)
            row.wornLevel?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
            pickArrow?.let { Text(it, style = FsType.meta, color = ItemQuality(row.pick.quality).text) }
        }
    }
}

@Composable
private fun Trailing(row: GearRow, build: String, open: (ItemLook, Int?) -> Unit) {
    val pick = row.pick
    when {
        row.isBestInSlot -> Row(
            Modifier.background(FsColor.kill.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill))
                .border(1.dp, FsColor.kill.copy(alpha = 0.35f), RoundedCornerShape(FsRadius.pill))
                .padding(horizontal = 8.dp, vertical = 4.dp),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.kill, modifier = Modifier.size(12.dp))
            Spacer(Modifier.width(4.dp))
            Text(stringResource(R.string.gear_bis).uppercase(Locale.ROOT), style = FsType.micro, color = FsColor.kill)
        }
        pick != null -> {
            val quality = ItemQuality(pick.quality).name?.let { stringResource(it) }.orEmpty()
            val spoken = stringResource(R.string.tooltip_a11yTile, pick.name.orEmpty(), quality, row.slotName)
            Row(
                Modifier.heightIn(min = 48.dp).clickable(role = Role.Button) { open(pick, row.worn?.id) }
                    .semantics(mergeDescendants = true) { contentDescription = spoken },
                verticalAlignment = Alignment.CenterVertically,
                horizontalArrangement = Arrangement.spacedBy(8.dp),
            ) {
                Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
                ItemIcon(pick.icon, pick.quality, build, 36.dp)
            }
        }
    }
}

private val SLOT_LABEL_WIDTH = 84.dp
