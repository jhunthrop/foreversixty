package gg.foreversixty.app.account

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Info
import androidx.compose.material.icons.filled.Lock
import androidx.compose.material.icons.filled.Person
import androidx.compose.material3.AlertDialog
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.graphics.vector.rememberVectorPainter
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.Seal
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.premium.PlanRow
import gg.foreversixty.app.premium.PlanSummary

/** The current character and where it plays (`Obnoxious Yell`, `US · PvP`). */
data class CharacterLine(val name: String, val className: String?, val place: String)

/**
 * What the You tab shows, already worded (spec F8.1: the account head; Plan, Character and
 * Devices; Privacy, Terms and About; Sign out; Delete account; the version line).
 */
data class YouSummary(
    val battletag: String,
    /** Lane F8: the plan in the site's words. */
    val plan: PlanSummary,
    val character: CharacterLine?,
    /** `This phone and 2 more`, null until the device list answers. */
    val devices: String?,
    /** `Version 0.1.0 (1) · data 1.60.1.70658`. */
    val version: String,
)

/** What the You tab's rows open. */
class YouActions(
    /** Lane F8: `See plans` opens foreversixty.gg/premium; a Premium row opens the Plan screen. */
    val onSeePlans: () -> Unit,
    val onPlan: () -> Unit,
    val onDevices: () -> Unit,
    val onPrivacy: () -> Unit,
    val onTerms: () -> Unit,
    val onAbout: () -> Unit,
    val onDeleteAccount: () -> Unit,
    val onSignOut: () -> Unit,
    /** Lane F16: You › Calendar; null hides the row while `guild_extras` is off. */
    val onCalendar: (() -> Unit)? = null,
    /** Lane F15: You › Offline data; null hides the row while `system` is off. */
    val onOfflineData: (() -> Unit)? = null,
)

/**
 * The You tab. F7 and F8 add the Notifications, Characters and Discord rows. Stateless apart
 * from the sign-out confirmation.
 */
@Composable
fun YouScreen(
    summary: YouSummary,
    actions: YouActions,
    mayLink: Boolean = false,
    confirmingSignOut: Boolean = false,
    /** Lane F14: the Live surfaces card, drawn while the `widgets` switch is on. */
    liveSurfaces: (@Composable () -> Unit)? = null,
) {
    var confirming by rememberSaveable { mutableStateOf(confirmingSignOut) }
    AccountPage(stringResource(R.string.you_title)) {
        Row(horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3), verticalAlignment = Alignment.CenterVertically) {
            Tile(rememberVectorPainter(Icons.Filled.Person), 44.dp)
            Column(Modifier.weight(1f)) {
                Text(summary.battletag, style = FsType.name, color = FsColor.strong)
                Text(stringResource(R.string.you_signedInWith), style = FsType.meta, color = FsColor.muted)
            }
            if (summary.plan.isPremium) Tag(stringResource(R.string.plan_premium), FsColor.gold)
        }
        AccountCard {
            PlanRow(summary.plan, mayLink, actions.onSeePlans, actions.onPlan)
            summary.character?.let {
                HorizontalDivider(color = FsColor.lineSoft)
                CharacterRow(it)
            }
            HorizontalDivider(color = FsColor.lineSoft)
            AccountRow(painterResource(R.drawable.ic_smartphone), stringResource(R.string.you_devices), subtitle = summary.devices, chevron = true, onClick = actions.onDevices)
            actions.onCalendar?.let { onCalendar ->
                HorizontalDivider(color = FsColor.lineSoft)
                AccountRow(
                    painterResource(R.drawable.ic_calendar),
                    stringResource(R.string.calendar_title),
                    subtitle = stringResource(R.string.calendar_youRow_sub),
                    chevron = true,
                    onClick = onCalendar,
                )
            }
            actions.onOfflineData?.let { onOfflineData ->
                HorizontalDivider(color = FsColor.lineSoft)
                AccountRow(
                    painterResource(R.drawable.ic_cloud_download),
                    stringResource(R.string.offlineData_title),
                    subtitle = stringResource(R.string.offlineData_youRow_sub),
                    chevron = true,
                    onClick = onOfflineData,
                )
            }
        }
        liveSurfaces?.invoke()
        AccountCard {
            AccountRow(rememberVectorPainter(Icons.Filled.Lock), stringResource(R.string.you_privacy), external = true, onClick = actions.onPrivacy)
            HorizontalDivider(color = FsColor.lineSoft)
            AccountRow(rememberVectorPainter(Icons.Filled.Info), stringResource(R.string.you_terms), external = true, onClick = actions.onTerms)
            HorizontalDivider(color = FsColor.lineSoft)
            AccountRow(rememberVectorPainter(Icons.Filled.Info), stringResource(R.string.you_about), chevron = true, onClick = actions.onAbout)
        }
        BrandButton(stringResource(R.string.you_signOut), BrandButtonKind.Quiet, { confirming = true }, height = 48.dp)
        TextButton(actions.onDeleteAccount, Modifier.fillMaxWidth().heightIn(min = 48.dp)) {
            Text(stringResource(R.string.account_delete_row), style = FsType.bodyStrong, color = FsColor.wipe)
        }
        Text(summary.version, style = FsType.micro, color = FsColor.muted, modifier = Modifier.fillMaxWidth(), textAlign = TextAlign.Center)
    }
    if (confirming) {
        ConfirmDialog(
            title = stringResource(R.string.you_signOut_confirm, stringResource(R.string.account_deviceWord)),
            body = null,
            action = stringResource(R.string.account_signOut_action),
            onConfirm = {
                confirming = false
                actions.onSignOut()
            },
            onDismiss = { confirming = false },
        )
    }
}

@Composable
private fun CharacterRow(character: CharacterLine) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = 56.dp).padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        ClassCrest(character.className, 32.dp)
        Column {
            Text(stringResource(R.string.account_you_character), style = FsType.bodyStrong, color = FsColor.strong)
            Text(
                buildAnnotatedString {
                    withStyle(SpanStyle(color = classColor(character.className))) { append(character.name) }
                    withStyle(SpanStyle(color = FsColor.muted)) { append(" · ${character.place}") }
                },
                style = FsType.meta,
            )
        }
    }
}

/** You › About Forever Sixty: the version, the build and the site footer's Blizzard line. */
@Composable
fun AboutScreen(version: String) {
    AccountPage(stringResource(R.string.you_about)) {
        Seal(64.dp)
        Text(stringResource(R.string.app_name).uppercase(), style = FsType.wordmark, color = FsColor.strong)
        Text(version, style = FsType.body, color = FsColor.text)
        Text(stringResource(R.string.about_notAffiliated), style = FsType.meta, color = FsColor.muted)
    }
}

/** A pushed or root account page: its title, then the content on the ground, scrolling. */
@Composable
fun AccountPage(title: String, content: @Composable ColumnScope.() -> Unit) {
    Column(
        Modifier
            .fillMaxSize()
            .background(FsColor.bg)
            .verticalScroll(rememberScrollState())
            .padding(FsSpacing.space4),
        verticalArrangement = Arrangement.spacedBy(FsSpacing.space5),
    ) {
        Text(title.uppercase(), style = FsType.title, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        content()
    }
}

@Composable
fun AccountCard(content: @Composable ColumnScope.() -> Unit) {
    Column(Modifier.fillMaxWidth().quietCard(), content = content)
}

@Composable
private fun Tile(icon: Painter, size: androidx.compose.ui.unit.Dp, tint: Color = FsColor.gold) {
    Box(
        Modifier.size(size).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill)),
        contentAlignment = Alignment.Center,
    ) { Icon(icon, contentDescription = null, tint = tint, modifier = Modifier.size(20.dp)) }
}

/** One list row: a gold tile, the title and its sub line, a trailing value, tag or chevron. */
@Composable
fun AccountRow(
    icon: Painter,
    title: String,
    subtitle: String? = null,
    trailing: String? = null,
    tag: String? = null,
    chevron: Boolean = false,
    /** Opens the browser: the external glyph instead of the chevron. */
    external: Boolean = false,
    tint: Color = FsColor.strong,
    onClick: (() -> Unit)? = null,
    end: @Composable () -> Unit = {},
) {
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = 56.dp)
            .let { if (onClick != null) it.clickable(onClick = onClick) else it }
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Tile(icon, 32.dp, if (tint == FsColor.strong) FsColor.gold else tint)
        Column(Modifier.weight(1f)) {
            Text(title, style = FsType.bodyStrong, color = tint)
            subtitle?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
        }
        tag?.let { Tag(it, FsColor.gold) }
        trailing?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
        end()
        if (external) Icon(painterResource(R.drawable.ic_open_external), contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(18.dp))
        if (chevron) Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
    }
}

/** The destructive confirmation (spec §3.2: `AlertDialog`). */
@Composable
fun ConfirmDialog(title: String, body: String?, action: String, onConfirm: () -> Unit, onDismiss: () -> Unit) {
    AlertDialog(
        onDismissRequest = onDismiss,
        containerColor = FsColor.raised,
        title = { Text(title, style = FsType.sec, color = FsColor.strong) },
        text = body?.let { { Text(it, style = FsType.body, color = FsColor.text) } },
        confirmButton = { TextButton(onConfirm) { Text(action, style = FsType.bodyStrong, color = FsColor.wipe) } },
        dismissButton = { TextButton(onDismiss) { Text(stringResource(R.string.common_cancel), style = FsType.bodyStrong, color = FsColor.nav) } },
    )
}
