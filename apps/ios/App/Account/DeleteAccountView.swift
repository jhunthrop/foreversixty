import Auth
import DesignSystem
import SwiftUI

/// You › Delete account (App Store guideline 5.1.1(v)): what is deleted and what is kept, in the
/// web dialog's words; what else ends for this account (its plan, founder price, seat); then the
/// typed confirmation (the BattleTag, or delete) and the destructive button, enabled only once the
/// field matches by the server's rule (`DeletionConfirmation`). Pure over its inputs.
struct DeleteAccountView: View {
    /// The lines that apply to this account (`DeletionConsequence.of`), in order.
    let consequences: [DeletionConsequence]
    let battleTag: String?
    @Binding var typed: String
    let isDeleting: Bool
    /// The last attempt's outcome when it did not delete; nil before the first try.
    let outcome: AccountActionOutcome?
    let onDelete: () -> Void

    static let deletedItems = [
        Copy.sharedAccountDeleteDeletedSubscription,
        Copy.sharedAccountDeleteDeletedIdentity,
        Copy.sharedAccountDeleteDeletedDiscord,
        Copy.sharedAccountDeleteDeletedPrivate,
        Copy.sharedAccountDeleteDeletedGuild,
    ]

    static let keptItems = [
        Copy.sharedAccountDeleteKeptBilling,
        Copy.sharedAccountDeleteKeptReports,
        Copy.sharedAccountDeleteKeptCharacters,
        Copy.sharedAccountDeleteKeptBuilds,
    ]

    private var consequenceLines: [String] { consequences.map(AccountCopy.consequence) }
    private var canDelete: Bool { !isDeleting && DeletionConfirmation.matches(typed, battleTag: battleTag) }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                BandHead(title: Copy.sharedAccountDeleteTitle)
                form
            }
        }
        .screenGround()
        .pushedBandBar(Copy.sharedAccountDeleteTitle)
    }

    private var form: some View {
        VStack(alignment: .leading, spacing: FSSpace.space4) {
            Text(Copy.sharedAccountDeleteIntro).font(ScaledType.body).foregroundStyle(FSColor.strong)
            group(Copy.sharedAccountDeleteDeletedTitle, Self.deletedItems, glyph: "minus")
            group(Copy.sharedAccountDeleteKeptTitle, Self.keptItems, glyph: "checkmark", tint: FSColor.muted)
            if !consequenceLines.isEmpty {
                group(Copy.sharedAccountDeleteLoseTitle, consequenceLines, glyph: "exclamationmark.circle")
            }
            confirmation
            Text(Copy.sharedAccountDeleteEmail).font(ScaledType.meta).foregroundStyle(FSColor.mutedGround)
        }
        .padding(FSSpace.space4)
    }

    private var confirmLabel: String { AccountCopy.deletionPrompt(battleTag: battleTag) }

    /// The confirmation on one Plate: the field, its failure in wipe and the button read on the
    /// panel's dark, never loose on the painted ground.
    private var confirmation: some View {
        VStack(alignment: .leading, spacing: FSSpace.space3) {
            confirmField
            if let failure = AccountCopy.deletionFailure(outcome) {
                Text(failure).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                    .fixedSize(horizontal: false, vertical: true)
            }
            Button(action: onDelete) {
                if isDeleting {
                    ProgressView().tint(FSColor.wipe).frame(width: 20, height: 20)
                } else {
                    Text(Copy.sharedAccountDeleteButton)
                }
            }
            .buttonStyle(BrandButtonStyle(.destructive))
            .disabled(!canDelete)
            .accessibilityLabel(Copy.sharedAccountDeleteButton)
        }
        .padding(FSSpace.space4)
        .frame(maxWidth: .infinity, alignment: .leading)
        .panel()
    }

    private var confirmField: some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            Label12(confirmLabel, color: FSColor.text)
            TextField("", text: $typed)
                .font(ScaledType.body)
                .foregroundStyle(FSColor.strong)
                .textInputAutocapitalization(.never)
                .autocorrectionDisabled()
                .submitLabel(.done)
                .disabled(isDeleting)
                .padding(.horizontal, 14)
                .frame(minHeight: 48)
                .background(Color.black.opacity(0.35), in: .rect(cornerRadius: FSRadius.legacyControl))
                .overlay(RoundedRectangle(cornerRadius: FSRadius.legacyControl).stroke(
                    outcome == .confirmMismatch ? FSColor.wipe : FSColor.lineCard
                ))
                .accessibilityLabel(confirmLabel)
            if outcome == .confirmMismatch {
                Text(AccountCopy.deletionMismatch(battleTag: battleTag)).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
            }
        }
    }

    private func group(_ title: String, _ lines: [String], glyph: String, tint: Color = FSColor.wipe) -> some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            SectionTitle(title)
            VStack(alignment: .leading, spacing: FSSpace.space3) {
                ForEach(lines, id: \.self) { line in
                    HStack(alignment: .firstTextBaseline, spacing: FSSpace.space3) {
                        Image(systemName: glyph).foregroundStyle(tint).accessibilityHidden(true)
                        Text(line).font(ScaledType.body).foregroundStyle(FSColor.strong)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                }
            }
            .padding(FSSpace.space4)
            .frame(maxWidth: .infinity, alignment: .leading)
            .quietCard()
        }
    }
}

/// Runs `DELETE /v1/me` with what the person typed; on success the session shows signed out.
struct DeleteAccountScreen: View {
    let session: SessionStore
    @Environment(\.openURL) private var openURL
    @State private var typed = ""
    @State private var isDeleting = false
    @State private var outcome: AccountActionOutcome?

    var body: some View {
        DeleteAccountView(
            consequences: session.phase.me.map(DeletionConsequence.of) ?? [],
            battleTag: session.phase.me?.user.battletag,
            typed: $typed,
            isDeleting: isDeleting,
            outcome: outcome,
            onDelete: delete
        )
        .onChange(of: typed) { _, _ in
            if outcome == .confirmMismatch { outcome = nil }
        }
    }

    private func delete() {
        isDeleting = true
        Task {
            let result = await session.deleteAccount(confirm: typed)
            isDeleting = false
            outcome = result == .done ? nil : result
            if result == .needsBrowser { openURL(AccountLinks.account) }
        }
    }
}
