// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	CreateWallet(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Transfer(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetWalletIfc(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*WalletBalanceResp, error)
	GetWalletUsdt(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*WalletBalanceResp, error)
	GetUsdtDepositAddr(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*UsdtDepositAddrResp, error)
	ConvertCoinTips(ctx context.Context, in *ConvertCoinTipsReq, opts ...grpc.CallOption) (*ConvertCoinTipsResp, error)
	ConvertCoin(ctx context.Context, in *ConvertCoinReq, opts ...grpc.CallOption) (*ConvertCoinResp, error)
	Withdrawal(ctx context.Context, in *WithdrawalReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddIfcBalance(ctx context.Context, in *AddIfcBalanceReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTotalRebate(ctx context.Context, in *GetTotalRebateReq, opts ...grpc.CallOption) (*GetTotalRebateResp, error)
	StrategyRunNotify(ctx context.Context, in *StrategyRunNotifyReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) CreateWallet(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Transfer(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetWalletIfc(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*WalletBalanceResp, error) {
	out := new(WalletBalanceResp)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/GetWalletIfc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetWalletUsdt(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*WalletBalanceResp, error) {
	out := new(WalletBalanceResp)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/GetWalletUsdt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetUsdtDepositAddr(ctx context.Context, in *UidReq, opts ...grpc.CallOption) (*UsdtDepositAddrResp, error) {
	out := new(UsdtDepositAddrResp)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/GetUsdtDepositAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ConvertCoinTips(ctx context.Context, in *ConvertCoinTipsReq, opts ...grpc.CallOption) (*ConvertCoinTipsResp, error) {
	out := new(ConvertCoinTipsResp)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/ConvertCoinTips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ConvertCoin(ctx context.Context, in *ConvertCoinReq, opts ...grpc.CallOption) (*ConvertCoinResp, error) {
	out := new(ConvertCoinResp)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/ConvertCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Withdrawal(ctx context.Context, in *WithdrawalReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/Withdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) AddIfcBalance(ctx context.Context, in *AddIfcBalanceReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/AddIfcBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetTotalRebate(ctx context.Context, in *GetTotalRebateReq, opts ...grpc.CallOption) (*GetTotalRebateResp, error) {
	out := new(GetTotalRebateResp)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/GetTotalRebate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) StrategyRunNotify(ctx context.Context, in *StrategyRunNotifyReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.wallet.v1.Wallet/StrategyRunNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations must embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	CreateWallet(context.Context, *UidReq) (*emptypb.Empty, error)
	Transfer(context.Context, *TransferReq) (*emptypb.Empty, error)
	GetWalletIfc(context.Context, *UidReq) (*WalletBalanceResp, error)
	GetWalletUsdt(context.Context, *UidReq) (*WalletBalanceResp, error)
	GetUsdtDepositAddr(context.Context, *UidReq) (*UsdtDepositAddrResp, error)
	ConvertCoinTips(context.Context, *ConvertCoinTipsReq) (*ConvertCoinTipsResp, error)
	ConvertCoin(context.Context, *ConvertCoinReq) (*ConvertCoinResp, error)
	Withdrawal(context.Context, *WithdrawalReq) (*emptypb.Empty, error)
	AddIfcBalance(context.Context, *AddIfcBalanceReq) (*emptypb.Empty, error)
	GetTotalRebate(context.Context, *GetTotalRebateReq) (*GetTotalRebateResp, error)
	StrategyRunNotify(context.Context, *StrategyRunNotifyReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedWalletServer()
}

// UnimplementedWalletServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) CreateWallet(context.Context, *UidReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServer) Transfer(context.Context, *TransferReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedWalletServer) GetWalletIfc(context.Context, *UidReq) (*WalletBalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletIfc not implemented")
}
func (UnimplementedWalletServer) GetWalletUsdt(context.Context, *UidReq) (*WalletBalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletUsdt not implemented")
}
func (UnimplementedWalletServer) GetUsdtDepositAddr(context.Context, *UidReq) (*UsdtDepositAddrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsdtDepositAddr not implemented")
}
func (UnimplementedWalletServer) ConvertCoinTips(context.Context, *ConvertCoinTipsReq) (*ConvertCoinTipsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCoinTips not implemented")
}
func (UnimplementedWalletServer) ConvertCoin(context.Context, *ConvertCoinReq) (*ConvertCoinResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCoin not implemented")
}
func (UnimplementedWalletServer) Withdrawal(context.Context, *WithdrawalReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdrawal not implemented")
}
func (UnimplementedWalletServer) AddIfcBalance(context.Context, *AddIfcBalanceReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIfcBalance not implemented")
}
func (UnimplementedWalletServer) GetTotalRebate(context.Context, *GetTotalRebateReq) (*GetTotalRebateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalRebate not implemented")
}
func (UnimplementedWalletServer) StrategyRunNotify(context.Context, *StrategyRunNotifyReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StrategyRunNotify not implemented")
}
func (UnimplementedWalletServer) mustEmbedUnimplementedWalletServer() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateWallet(ctx, req.(*UidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Transfer(ctx, req.(*TransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetWalletIfc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetWalletIfc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/GetWalletIfc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetWalletIfc(ctx, req.(*UidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetWalletUsdt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetWalletUsdt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/GetWalletUsdt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetWalletUsdt(ctx, req.(*UidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetUsdtDepositAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetUsdtDepositAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/GetUsdtDepositAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetUsdtDepositAddr(ctx, req.(*UidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ConvertCoinTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertCoinTipsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ConvertCoinTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/ConvertCoinTips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ConvertCoinTips(ctx, req.(*ConvertCoinTipsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ConvertCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertCoinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ConvertCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/ConvertCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ConvertCoin(ctx, req.(*ConvertCoinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Withdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Withdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/Withdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Withdrawal(ctx, req.(*WithdrawalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_AddIfcBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIfcBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).AddIfcBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/AddIfcBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).AddIfcBalance(ctx, req.(*AddIfcBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetTotalRebate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalRebateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetTotalRebate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/GetTotalRebate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetTotalRebate(ctx, req.(*GetTotalRebateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_StrategyRunNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrategyRunNotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).StrategyRunNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wallet.v1.Wallet/StrategyRunNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).StrategyRunNotify(ctx, req.(*StrategyRunNotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.wallet.v1.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWallet",
			Handler:    _Wallet_CreateWallet_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Wallet_Transfer_Handler,
		},
		{
			MethodName: "GetWalletIfc",
			Handler:    _Wallet_GetWalletIfc_Handler,
		},
		{
			MethodName: "GetWalletUsdt",
			Handler:    _Wallet_GetWalletUsdt_Handler,
		},
		{
			MethodName: "GetUsdtDepositAddr",
			Handler:    _Wallet_GetUsdtDepositAddr_Handler,
		},
		{
			MethodName: "ConvertCoinTips",
			Handler:    _Wallet_ConvertCoinTips_Handler,
		},
		{
			MethodName: "ConvertCoin",
			Handler:    _Wallet_ConvertCoin_Handler,
		},
		{
			MethodName: "Withdrawal",
			Handler:    _Wallet_Withdrawal_Handler,
		},
		{
			MethodName: "AddIfcBalance",
			Handler:    _Wallet_AddIfcBalance_Handler,
		},
		{
			MethodName: "GetTotalRebate",
			Handler:    _Wallet_GetTotalRebate_Handler,
		},
		{
			MethodName: "StrategyRunNotify",
			Handler:    _Wallet_StrategyRunNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/wallet/v1/wallet.proto",
}