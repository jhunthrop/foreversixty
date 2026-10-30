import Auth
import DesignSystem
import SwiftUI

/// You › Delete account (App Store guideline 5.1.1(v)): what is deleted, in the privacy page's
/// words, then the destructive button and an explicit confirmation alert. Pure over its inputs.
struct DeleteAccountView: View {
    /// Whether the account pays for a plan: only then does the plan line show.
    let hasPlan: Bool
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

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: FSSpacing.space4) {
                Text(Copy.accountDeleteIntro).font(ScaledType.body).foregroundStyle(FSColor.text)
                VStack(alignment: .leading, spacing: FSSpacing.space3) {
                    ForEach(items, id: \.self) { item in
                        HStack(alignment: .firstTextBaseline, spacing: FSSpacing.space3) {
                            Image(systemName: "minus").foregroundStyle(FSColor.wipe).accessibilityHidden(true)
                            Text(item).font(ScaledType.body).foregroundStyle(FSColor.strong)
                                .fixedSize(horizontal: false, vertical: true)
                        }
                    }
                }
                .padding(FSSpacing.space4)
                .frame(maxWidth: .infinity, alignment: .leading)
                .accountCard()
                if hasPlan {
                    Text(Copy.accountDeletePlan).font(ScaledType.body).foregroundStyle(FSColor.text)
                }
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
            Button(Copy.accountDeleteConfirmAction, role: .destructive, action: onDelete)
        } message: {
            Text(Copy.accountDeleteConfirmBody)
        }
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
            hasPlan: session.phase.me?.entitlements.billing != nil,
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
