import Auth
import DesignSystem
import SwiftUI

/// You › Delete account (App Store guideline 5.1.1(v)): what is deleted, in the privacy page's
/// words; what else ends for this account (its plan, founder price, guild plan, seat); what
/// happens to guild reports; then the destructive button and an explicit confirmation alert
/// that repeats the consequences. Pure over its inputs.
struct DeleteAccountView: View {
    /// The lines that apply to this account (`DeletionConsequence.of`), in order.
    let consequences: [DeletionConsequence]
    let isDeleting: Bool
    let failed: Bool
    let onDelete: () -> Void
    @State private var confirming = false

    private var items: [String] {
        [
            Copy.accountDeleteItemIdentity,
            Copy.accountDeleteItemCharacters,
            Copy.accountDeleteItemLogs,
            Copy.accountDeleteItemSims,
            Copy.accountDeleteItemLinks,
        ]
    }

    private var consequenceLines: [String] { consequences.map(AccountCopy.consequence) }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: FSSpacing.space4) {
                Text(Copy.accountDeleteIntro).font(ScaledType.body).foregroundStyle(FSColor.text)
                bulletCard(items, glyph: "minus")
                if !consequenceLines.isEmpty {
                    Text(Copy.accountDeleteLoseTitle)
                        .font(ScaledType.micro)
                        .textCase(.uppercase)
                        .tracking(1.5)
                        .foregroundStyle(FSColor.muted)
                        .padding(.top, FSSpacing.space2)
                    bulletCard(consequenceLines, glyph: "exclamationmark.circle")
                }
                // Pending the owner's decision on guild reports: one key, account.delete.reports.
                Text(Copy.accountDeleteReports).font(ScaledType.body).foregroundStyle(FSColor.text)
                    .fixedSize(horizontal: false, vertical: true)
                Text(Copy.accountDeleteTiming).font(ScaledType.body).foregroundStyle(FSColor.text)
                    .fixedSize(horizontal: false, vertical: true)
                if failed {
                    Text(Copy.accountDeleteFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                }
                Button { confirming = true } label: {
                    if isDeleting {
                        ProgressView().tint(FSColor.wipe).frame(width: 20, height: 20)
                    } else {
                        Text(Copy.accountDeleteButton)
                    }
                }
                .buttonStyle(BrandButtonStyle(.destructive))
                .disabled(isDeleting)
                .accessibilityLabel(Copy.accountDeleteButton)
                Text(Copy.accountDeleteEmail).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .padding(FSSpacing.space4)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.accountDeleteTitle)
        .alert(Copy.accountDeleteConfirmTitle, isPresented: $confirming) {
            Button(Copy.commonCancel, role: .cancel) {}
            Button(Copy.accountDeleteButton, role: .destructive, action: onDelete)
        } message: {
            Text(([Copy.accountDeleteConfirmBody] + consequenceLines).joined(separator: "\n\n"))
        }
    }

    private func bulletCard(_ lines: [String], glyph: String) -> some View {
        VStack(alignment: .leading, spacing: FSSpacing.space3) {
            ForEach(lines, id: \.self) { line in
                HStack(alignment: .firstTextBaseline, spacing: FSSpacing.space3) {
                    Image(systemName: glyph).foregroundStyle(FSColor.wipe).accessibilityHidden(true)
                    Text(line).font(ScaledType.body).foregroundStyle(FSColor.strong)
                        .fixedSize(horizontal: false, vertical: true)
                }
            }
        }
        .padding(FSSpacing.space4)
        .frame(maxWidth: .infinity, alignment: .leading)
        .accountCard()
    }
}

/// Runs `DELETE /v1/me` after the alert; on success the session shows signed out with the date.
struct DeleteAccountScreen: View {
    let session: SessionStore
    @Environment(\.openURL) private var openURL
    @State private var isDeleting = false
    @State private var failed = false

    var body: some View {
        DeleteAccountView(
            consequences: session.phase.me.map(DeletionConsequence.of) ?? [],
            isDeleting: isDeleting,
            failed: failed,
            onDelete: delete
        )
    }

    private func delete() {
        isDeleting = true
        Task {
            let outcome = await session.deleteAccount()
            isDeleting = false
            failed = outcome == .failed
            if outcome == .needsBrowser { openURL(AccountLinks.account) }
        }
    }
}
