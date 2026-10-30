import Auth
import DesignSystem
import Networking
import SwiftUI

/// What F1.6 shows: the rows, or the load failure.
enum DevicesState: Equatable {
    case loading
    case loaded([DeviceRow])
    case failed
}

/// F1.6 You › Devices. Pure over its state; `DevicesScreen` loads and acts.
struct DevicesView: View {
    let state: DevicesState
    /// The row whose revoke failed: `Could not revoke that device.` shows under it.
    let failedRevoke: String?
    let onRevoke: (DeviceRow) -> Void
    let onRetry: () -> Void
    @State private var confirming: DeviceRow?

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: FSSpacing.space4) {
                content
                Text(Copy.devicesFooter).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .padding(FSSpacing.space4)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.devicesTitle)
        .alert(
            confirming.map { Copy.devicesRevokeTitle(device: $0.name) } ?? "",
            isPresented: Binding(get: { confirming != nil }, set: { if !$0 { confirming = nil } }),
            presenting: confirming
        ) { row in
            Button(Copy.commonCancel, role: .cancel) {}
            Button(Copy.devicesRevoke, role: .destructive) { onRevoke(row) }
        } message: { _ in
            Text(Copy.devicesRevokeBody)
        }
    }

    @ViewBuilder
    private var content: some View {
        switch state {
        case .loading:
            VStack(spacing: 0) {
                ForEach(0..<3, id: \.self) { _ in
                    AccountRow(icon: "iphone", title: "Placeholder device", subtitle: "Last used")
                }
            }
            .redacted(reason: .placeholder)
            .accountCard()
        case let .loaded(rows):
            VStack(spacing: 0) {
                ForEach(rows) { row in
                    deviceRow(row)
                    if row.id != rows.last?.id { Divider().overlay(FSColor.lineSoft) }
                }
            }
            .accountCard()
        case .failed:
            VStack(alignment: .leading, spacing: FSSpacing.space3) {
                Text(Copy.accountDevicesLoadFailed)
                    .font(ScaledType.body)
                    .foregroundStyle(FSColor.wipe)
                    .accessibilityAddTraits(.isStaticText)
                Button(Copy.commonTryAgain, action: onRetry).buttonStyle(BrandButtonStyle(.secondary, height: 44))
            }
        }
    }

    private func deviceRow(_ row: DeviceRow) -> some View {
        VStack(alignment: .leading, spacing: FSSpacing.space1) {
            HStack {
                AccountRow(icon: "iphone", title: AccountCopy.deviceName(row), subtitle: AccountCopy.deviceMeta(row))
                if row.canRevoke {
                    Button(Copy.devicesRevoke) { confirming = row }
                        .buttonStyle(BrandButtonStyle(.destructive, height: 36))
                        .fixedSize()
                        .padding(.trailing, FSSpacing.space4)
                }
            }
            if failedRevoke == row.id {
                Text(Copy.devicesRevokeFailed)
                    .font(ScaledType.meta)
                    .foregroundStyle(FSColor.wipe)
                    .padding([.horizontal, .bottom], FSSpacing.space4)
            }
        }
    }
}

/// Loads `GET /v1/devices`, revokes with `DELETE /v1/devices/{id}`.
struct DevicesScreen: View {
    let session: SessionStore
    let onCountChange: (Int) -> Void
    @Environment(\.openURL) private var openURL
    @State private var state: DevicesState = .loading
    @State private var failedRevoke: String?

    var body: some View {
        DevicesView(state: state, failedRevoke: failedRevoke, onRevoke: revoke, onRetry: { Task { await load() } })
            .task { await load() }
    }

    private func load() async {
        do throws(APIError) {
            let rows = try await session.devices()
            withAnimation { state = .loaded(rows) }
            onCountChange(rows.count)
        } catch {
            state = .failed
        }
    }

    private func revoke(_ row: DeviceRow) {
        Task {
            switch await session.revoke(row.id) {
            case .done:
                failedRevoke = nil
                await load()
            case .failed:
                failedRevoke = row.id
            case .needsBrowser:
                openURL(AccountLinks.account)
            }
        }
    }
}
