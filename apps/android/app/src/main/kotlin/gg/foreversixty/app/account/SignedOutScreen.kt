package gg.foreversixty.app.account

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.LinearProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType

/** Which tab draws the signed-out panel: Home's words, or Raids' (spec F1.1). */
enum class SignedOutVariant { Home, Raids }

/**
 * F1.1 signed out, with F1.2's spinner while the Custom Tab is open and F1.4's toast. Stateless:
 * the screenshot tests render it directly.
 */
@Composable
fun SignedOutScreen(
    notice: String?,
    signingIn: Boolean,
    onSignIn: () -> Unit,
    onPickClass: () -> Unit,
    variant: SignedOutVariant = SignedOutVariant.Home,
    noticeIsFailure: Boolean = true,
) {
    Box(Modifier.fillMaxSize().bandBackground()) {
        Column(
            Modifier
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .padding(horizontal = FsSpacing.space4)
                .padding(top = if (notice == null) 60.dp else 120.dp, bottom = FsSpacing.space6),
            verticalArrangement = Arrangement.spacedBy(FsSpacing.space4),
        ) {
            Seal(64.dp)
            Text(stringResource(R.string.app_name).uppercase(), style = FsType.wordmark, color = FsColor.strong)
            val (heading, line) = when (variant) {
                SignedOutVariant.Home -> R.string.signin_line to R.string.signin_body
                SignedOutVariant.Raids -> R.string.raids_signedOut_title to R.string.raids_signedOut_body
            }
            Text(
                stringResource(heading).uppercase(),
                style = FsType.title,
                color = FsColor.strong,
                modifier = Modifier.semantics { this.heading() },
            )
            Text(stringResource(line), style = FsType.body, color = FsColor.text)
            Spacer(Modifier.height(FsSpacing.space1))
            BrandButton(stringResource(R.string.signin_button), BrandButtonKind.Primary, onSignIn, busy = signingIn)
            BrandButton(stringResource(R.string.signin_pickClass), BrandButtonKind.Secondary, onPickClass)
            Text(stringResource(R.string.signin_pickClassNote), style = FsType.meta, color = FsColor.muted)
        }
        if (notice != null) {
            NoticeToast(notice, Modifier.padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space3), isFailure = noticeIsFailure)
        }
    }
}

/** F1.3: after the redirect, until `GET /v1/me` answers. */
@Composable
fun ImportScreen() {
    Column(
        Modifier.fillMaxSize().bandBackground().padding(horizontal = FsSpacing.space4),
        verticalArrangement = Arrangement.spacedBy(FsSpacing.space4, Alignment.CenterVertically),
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        Seal(64.dp)
        Text(stringResource(R.string.signin_importTitle).uppercase(), style = FsType.h2, color = FsColor.strong, textAlign = TextAlign.Center)
        Text(stringResource(R.string.signin_importBody), style = FsType.body, color = FsColor.text, textAlign = TextAlign.Center)
        LinearProgressIndicator(Modifier.width(160.dp).height(8.dp), color = FsColor.gold, trackColor = FsColor.lineSoft)
    }
}

/** F1.3 with no characters: Home's no-characters state. */
@Composable
fun NoCharactersScreen(onGetAddon: () -> Unit) {
    Column(
        Modifier.fillMaxSize().background(FsColor.bg).padding(FsSpacing.space4),
        verticalArrangement = Arrangement.spacedBy(FsSpacing.space4),
    ) {
        Text(
            stringResource(R.string.signin_noChars_title).uppercase(),
            style = FsType.h2,
            color = FsColor.strong,
            modifier = Modifier.semantics { heading() },
        )
        Text(stringResource(R.string.signin_noChars_body), style = FsType.body, color = FsColor.text)
        BrandButton(stringResource(R.string.signin_getAddon), BrandButtonKind.Secondary, onGetAddon, Modifier.fillMaxWidth())
    }
}
