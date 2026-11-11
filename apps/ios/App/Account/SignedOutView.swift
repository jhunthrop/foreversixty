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

    /// The head's top padding: clear of the toast when a notice shows.
    static let headTop: CGFloat = 48
    static let headTopWithNotice: CGFloat = 108

    var body: some View {
        // The web's signed-out page: a band head, then one Plate on the ground with the one primary.
        ZStack(alignment: .top) {
            ScrollView {
                VStack(alignment: .leading, spacing: 0) {
                    head
                    plate.padding(FSSpace.space4)
                }
            }
            .scrollBounceBehavior(.basedOnSize)
            if let notice {
                NoticeToast(text: notice, isFailure: noticeIsFailure)
                    .padding(.horizontal, FSSpace.space4)
                    .padding(.top, FSSpace.space3)
                    .accessibilitySortPriority(1)
            }
        }
        .screenGround()
    }

    private var head: some View {
        VStack(alignment: .leading, spacing: FSSpace.space4) {
            Seal(size: 64)
            Text(Copy.appName)
                .font(ScaledType.wordmark)
                .foregroundStyle(FSColor.strong)
            Text(heading)
                .font(ScaledType.title)
                .foregroundStyle(FSColor.strong)
                .accessibilityAddTraits(.isHeader)
        }
        .padding(.horizontal, FSSpace.space4)
        .padding(.top, notice == nil ? Self.headTop : Self.headTopWithNotice)
        .padding(.bottom, FSSpace.space6)
        .frame(maxWidth: .infinity, alignment: .leading)
        // The band runs up under the status bar, as every head's (N3).
        .bandUnderBars()
    }

    private var plate: some View {
        VStack(alignment: .leading, spacing: FSSpace.space4) {
            Text(line)
                .font(ScaledType.body)
                .foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            signInButton
            Button(Copy.signinPickClass, action: onPickClass)
                .buttonStyle(BrandButtonStyle(.quiet))
            Text(Copy.signinPickClassNote)
                .font(ScaledType.meta)
                .foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
        .padding(FSSpace.space4)
        .frame(maxWidth: .infinity, alignment: .leading)
        .panel()
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
            Color.clear.surfaceBand(bottomLine: false).ignoresSafeArea()
            VStack(spacing: FSSpace.space4) {
                Seal(size: 64)
                Text(Copy.signinImportTitle)
                    .font(ScaledType.h2)
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
            .padding(.horizontal, FSSpace.space4)
        }
    }
}

/// F1.3 with no characters: Home's no-characters state.
struct NoCharactersView: View {
    let onGetAddon: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            // The head on the band under the bars (N3), then the line and the one action.
            BandHead(title: Copy.signinNoCharsTitle)
            VStack(alignment: .leading, spacing: FSSpace.space4) {
                Text(Copy.signinNoCharsBody)
                    .font(ScaledType.body)
                    .foregroundStyle(FSColor.text)
                    .fixedSize(horizontal: false, vertical: true)
                Button(Copy.sharedCtaGetAddonAndCompanion, action: onGetAddon)
                    .buttonStyle(BrandButtonStyle(.secondary))
            }
            .padding(FSSpace.space4)
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .top)
        .screenGround()
    }
}
