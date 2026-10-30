import DesignSystem
import SwiftUI

/// F1.1 signed out (Home tab), with F1.2's spinner while the browser is open and F1.4's toast.
/// Raids reuses it with its own heading and line. Pure: the snapshot tests render it directly.
struct SignedOutView: View {
    enum Variant {
        case home
        case raids
    }

    var variant: Variant = .home
    let notice: String?
    var noticeIsFailure = true
    let isSigningIn: Bool
    let onSignIn: () -> Void
    let onPickClass: () -> Void

    var body: some View {
        ZStack(alignment: .top) {
            BandBackground()
            ScrollView {
                content
                    .padding(.horizontal, FSSpacing.space4)
                    .padding(.top, notice == nil ? 60 : 120)
                    .padding(.bottom, FSSpacing.space6)
            }
            .scrollBounceBehavior(.basedOnSize)
            if let notice {
                NoticeToast(text: notice, isFailure: noticeIsFailure)
                    .padding(.horizontal, FSSpacing.space4)
                    .padding(.top, FSSpacing.space3)
                    .accessibilitySortPriority(1)
            }
        }
    }

    private var content: some View {
        VStack(alignment: .leading, spacing: FSSpacing.space4) {
            Seal(size: 64)
            Text(Copy.appName)
                .font(ScaledType.wordmark)
                .textCase(.uppercase)
                .tracking(2)
                .foregroundStyle(FSColor.strong)
            Text(heading)
                .font(ScaledType.title)
                .textCase(.uppercase)
                .foregroundStyle(FSColor.strong)
                .accessibilityAddTraits(.isHeader)
            Text(line)
                .font(ScaledType.body)
                .foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            VStack(spacing: FSSpacing.space3) {
                signInButton
                Button(Copy.signinPickClass, action: onPickClass)
                    .buttonStyle(BrandButtonStyle(.secondary))
            }
            .padding(.top, FSSpacing.space2)
            Text(Copy.signinPickClassNote)
                .font(ScaledType.meta)
                .foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
    }

    private var signInButton: some View {
        Button(action: onSignIn) {
            if isSigningIn {
                ProgressView().tint(FSColor.bg).frame(width: 20, height: 20)
            } else {
                Text(Copy.signinButton)
            }
        }
        .buttonStyle(BrandButtonStyle(.primary))
        .disabled(isSigningIn)
        .accessibilityLabel(Copy.signinButton)
    }

    private var heading: String {
        switch variant {
        case .home: Copy.signinLine
        case .raids: Copy.raidsSignedOutTitle
        }
    }

    private var line: String {
        switch variant {
        case .home: Copy.signinBody
        case .raids: Copy.raidsSignedOutBody
        }
    }
}

/// F1.3: after the exchange, until `GET /v1/me` answers.
struct ImportView: View {
    var body: some View {
        ZStack {
            BandBackground()
            VStack(spacing: FSSpacing.space4) {
                Seal(size: 64)
                Text(Copy.signinImportTitle)
                    .font(ScaledType.h2)
                    .textCase(.uppercase)
                    .foregroundStyle(FSColor.strong)
                    .multilineTextAlignment(.center)
                Text(Copy.signinImportBody)
                    .font(ScaledType.body)
                    .foregroundStyle(FSColor.text)
                    .multilineTextAlignment(.center)
                ProgressView()
                    .progressViewStyle(.linear)
                    .tint(FSColor.gold)
                    .frame(width: 160)
                    .accessibilityHidden(true)
            }
            .padding(.horizontal, FSSpacing.space4)
        }
    }
}

/// F1.3 with no characters: Home's no-characters state.
struct NoCharactersView: View {
    let onGetAddon: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: FSSpacing.space4) {
            Text(Copy.signinNoCharsTitle)
                .font(ScaledType.h2)
                .textCase(.uppercase)
                .foregroundStyle(FSColor.strong)
                .accessibilityAddTraits(.isHeader)
            Text(Copy.signinNoCharsBody)
                .font(ScaledType.body)
                .foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            Button(Copy.signinGetAddon, action: onGetAddon)
                .buttonStyle(BrandButtonStyle(.secondary))
        }
        .padding(FSSpacing.space4)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .top)
        .background(FSColor.bg)
    }
}
