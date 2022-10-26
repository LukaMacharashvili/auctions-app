// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Item struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartingPrice        float64  `protobuf:"fixed64,4,opt,name=startingPrice,proto3" json:"startingPrice,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	HighestBid           float64  `protobuf:"fixed64,7,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	OwnerId              int64    `protobuf:"varint,8,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	HighestBidder        int64    `protobuf:"varint,9,opt,name=highestBidder,proto3" json:"highestBidder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetStartingPrice() float64 {
	if m != nil {
		return m.StartingPrice
	}
	return 0
}

func (m *Item) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Item) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Item) GetHighestBid() float64 {
	if m != nil {
		return m.HighestBid
	}
	return 0
}

func (m *Item) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *Item) GetHighestBidder() int64 {
	if m != nil {
		return m.HighestBidder
	}
	return 0
}

type CreateRequest struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LoadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadRequest) Reset()         { *m = LoadRequest{} }
func (m *LoadRequest) String() string { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()    {}
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{3}
}

func (m *LoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadRequest.Unmarshal(m, b)
}
func (m *LoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadRequest.Marshal(b, m, deterministic)
}
func (m *LoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadRequest.Merge(m, src)
}
func (m *LoadRequest) XXX_Size() int {
	return xxx_messageInfo_LoadRequest.Size(m)
}
func (m *LoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadRequest proto.InternalMessageInfo

type LoadResponse struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadResponse) Reset()         { *m = LoadResponse{} }
func (m *LoadResponse) String() string { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()    {}
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{4}
}

func (m *LoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadResponse.Unmarshal(m, b)
}
func (m *LoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadResponse.Marshal(b, m, deterministic)
}
func (m *LoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadResponse.Merge(m, src)
}
func (m *LoadResponse) XXX_Size() int {
	return xxx_messageInfo_LoadResponse.Size(m)
}
func (m *LoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadResponse proto.InternalMessageInfo

func (m *LoadResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *LoadResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FetchRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{5}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FetchResponse struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{6}
}

func (m *FetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchResponse.Unmarshal(m, b)
}
func (m *FetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchResponse.Marshal(b, m, deterministic)
}
func (m *FetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchResponse.Merge(m, src)
}
func (m *FetchResponse) XXX_Size() int {
	return xxx_messageInfo_FetchResponse.Size(m)
}
func (m *FetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchResponse proto.InternalMessageInfo

func (m *FetchResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *FetchResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId              int64    `protobuf:"varint,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

type DeleteResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DeleteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type BidRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidRequest) Reset()         { *m = BidRequest{} }
func (m *BidRequest) String() string { return proto.CompactTextString(m) }
func (*BidRequest) ProtoMessage()    {}
func (*BidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{9}
}

func (m *BidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidRequest.Unmarshal(m, b)
}
func (m *BidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidRequest.Marshal(b, m, deterministic)
}
func (m *BidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidRequest.Merge(m, src)
}
func (m *BidRequest) XXX_Size() int {
	return xxx_messageInfo_BidRequest.Size(m)
}
func (m *BidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidRequest proto.InternalMessageInfo

func (m *BidRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BidRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BidRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type BidResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidResponse) Reset()         { *m = BidResponse{} }
func (m *BidResponse) String() string { return proto.CompactTextString(m) }
func (*BidResponse) ProtoMessage()    {}
func (*BidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{10}
}

func (m *BidResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidResponse.Unmarshal(m, b)
}
func (m *BidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidResponse.Marshal(b, m, deterministic)
}
func (m *BidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidResponse.Merge(m, src)
}
func (m *BidResponse) XXX_Size() int {
	return xxx_messageInfo_BidResponse.Size(m)
}
func (m *BidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BidResponse proto.InternalMessageInfo

func (m *BidResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BidResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FinishAuctionRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId              int64    `protobuf:"varint,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FinishAuctionRequest) Reset()         { *m = FinishAuctionRequest{} }
func (m *FinishAuctionRequest) String() string { return proto.CompactTextString(m) }
func (*FinishAuctionRequest) ProtoMessage()    {}
func (*FinishAuctionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{11}
}

func (m *FinishAuctionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishAuctionRequest.Unmarshal(m, b)
}
func (m *FinishAuctionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishAuctionRequest.Marshal(b, m, deterministic)
}
func (m *FinishAuctionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishAuctionRequest.Merge(m, src)
}
func (m *FinishAuctionRequest) XXX_Size() int {
	return xxx_messageInfo_FinishAuctionRequest.Size(m)
}
func (m *FinishAuctionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishAuctionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FinishAuctionRequest proto.InternalMessageInfo

func (m *FinishAuctionRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FinishAuctionRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

type FinishAuctionResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FinishAuctionResponse) Reset()         { *m = FinishAuctionResponse{} }
func (m *FinishAuctionResponse) String() string { return proto.CompactTextString(m) }
func (*FinishAuctionResponse) ProtoMessage()    {}
func (*FinishAuctionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{12}
}

func (m *FinishAuctionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishAuctionResponse.Unmarshal(m, b)
}
func (m *FinishAuctionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishAuctionResponse.Marshal(b, m, deterministic)
}
func (m *FinishAuctionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishAuctionResponse.Merge(m, src)
}
func (m *FinishAuctionResponse) XXX_Size() int {
	return xxx_messageInfo_FinishAuctionResponse.Size(m)
}
func (m *FinishAuctionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishAuctionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FinishAuctionResponse proto.InternalMessageInfo

func (m *FinishAuctionResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FinishAuctionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Item)(nil), "item.Item")
	proto.RegisterType((*CreateRequest)(nil), "item.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "item.CreateResponse")
	proto.RegisterType((*LoadRequest)(nil), "item.LoadRequest")
	proto.RegisterType((*LoadResponse)(nil), "item.LoadResponse")
	proto.RegisterType((*FetchRequest)(nil), "item.FetchRequest")
	proto.RegisterType((*FetchResponse)(nil), "item.FetchResponse")
	proto.RegisterType((*DeleteRequest)(nil), "item.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "item.DeleteResponse")
	proto.RegisterType((*BidRequest)(nil), "item.BidRequest")
	proto.RegisterType((*BidResponse)(nil), "item.BidResponse")
	proto.RegisterType((*FinishAuctionRequest)(nil), "item.FinishAuctionRequest")
	proto.RegisterType((*FinishAuctionResponse)(nil), "item.FinishAuctionResponse")
}

func init() {
	proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df)
}

var fileDescriptor_6007f868cf6553df = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x1d, 0xc7, 0xa5, 0xe3, 0xba, 0xa2, 0xd3, 0x50, 0xad, 0x8c, 0x14, 0x45, 0x16, 0x87,
	0x1c, 0x50, 0x81, 0xa0, 0x1e, 0x10, 0x12, 0x82, 0xd0, 0x56, 0x2a, 0xea, 0x01, 0x99, 0x1b, 0x37,
	0x37, 0x1e, 0x25, 0x2b, 0x61, 0x3b, 0xec, 0xae, 0xe1, 0xca, 0x87, 0xf2, 0x31, 0x68, 0x77, 0xed,
	0xc6, 0x1b, 0xb5, 0x42, 0x84, 0xdb, 0xce, 0x9b, 0x9d, 0x99, 0xe7, 0x37, 0x6f, 0x0d, 0xc0, 0x15,
	0x95, 0x67, 0x6b, 0x51, 0xab, 0x1a, 0x03, 0x7d, 0x4e, 0x7f, 0xf9, 0x10, 0x5c, 0x2b, 0x2a, 0xf1,
	0x08, 0x7c, 0x5e, 0x30, 0x6f, 0xe2, 0x4d, 0x07, 0x99, 0xcf, 0x0b, 0x44, 0x08, 0xaa, 0xbc, 0x24,
	0xe6, 0x4f, 0xbc, 0xe9, 0x41, 0x66, 0xce, 0x38, 0x81, 0xa8, 0x20, 0xb9, 0x10, 0x7c, 0xad, 0x78,
	0x5d, 0xb1, 0x81, 0x49, 0xf5, 0x21, 0x7c, 0x06, 0xb1, 0x54, 0xb9, 0x50, 0xbc, 0x5a, 0x7e, 0x16,
	0x7c, 0x41, 0x2c, 0x98, 0x78, 0x53, 0x2f, 0x73, 0x41, 0x1c, 0xc1, 0x90, 0x97, 0xf9, 0x92, 0xd8,
	0xd0, 0x74, 0xb0, 0x01, 0x9e, 0x42, 0x28, 0x55, 0xae, 0x1a, 0xc9, 0x42, 0x03, 0xb7, 0x11, 0x8e,
	0x01, 0x56, 0x7c, 0xb9, 0x22, 0xa9, 0xe6, 0xbc, 0x60, 0xfb, 0xa6, 0x61, 0x0f, 0x41, 0x06, 0xfb,
	0xf5, 0xcf, 0x8a, 0xc4, 0x75, 0xc1, 0x1e, 0x19, 0xfa, 0x5d, 0xa8, 0xd9, 0x6c, 0xee, 0x15, 0x24,
	0xd8, 0x81, 0xc9, 0xbb, 0x60, 0xfa, 0x02, 0xe2, 0x8f, 0x82, 0x72, 0x45, 0x19, 0x7d, 0x6f, 0x48,
	0x2a, 0x1c, 0x83, 0xd1, 0xc6, 0x88, 0x11, 0xcd, 0xe0, 0xcc, 0x88, 0xa6, 0x45, 0xca, 0xac, 0x66,
	0xef, 0xe0, 0xa8, 0x2b, 0x90, 0xeb, 0xba, 0x92, 0x7d, 0xea, 0x56, 0xc0, 0x8e, 0xfa, 0x08, 0x86,
	0x24, 0x44, 0x2d, 0x5a, 0x15, 0x6d, 0x90, 0xc6, 0x10, 0xdd, 0xd4, 0x79, 0xd1, 0x8e, 0x4b, 0x2f,
	0xe0, 0xd0, 0x86, 0x6d, 0xb3, 0xbf, 0x8c, 0x7f, 0xa0, 0xe9, 0x18, 0x0e, 0xaf, 0x48, 0x2d, 0x56,
	0xdd, 0x47, 0x6c, 0xed, 0x33, 0xbd, 0x84, 0xb8, 0xcd, 0xff, 0xd7, 0x98, 0x37, 0x10, 0x5f, 0xd0,
	0x37, 0xda, 0x88, 0xb5, 0xed, 0x9b, 0xde, 0x36, 0x7c, 0x67, 0x1b, 0x5a, 0xb6, 0xae, 0x74, 0x27,
	0xd9, 0x6e, 0x00, 0xe6, 0xbc, 0x78, 0x68, 0xee, 0x29, 0x84, 0x79, 0x59, 0x37, 0x95, 0x32, 0x45,
	0x5e, 0xd6, 0x46, 0x1a, 0x6f, 0xa4, 0xa1, 0x33, 0xb0, 0x33, 0x6c, 0x94, 0xbe, 0x85, 0xc8, 0x74,
	0xdb, 0x89, 0xca, 0x7b, 0x18, 0x5d, 0xf1, 0x8a, 0xcb, 0xd5, 0x87, 0x66, 0xa1, 0x7d, 0xff, 0xef,
	0x62, 0x5c, 0xc2, 0x93, 0xad, 0x0e, 0xbb, 0x10, 0x99, 0xfd, 0xf6, 0x21, 0xd2, 0x3b, 0xfb, 0x42,
	0xe2, 0x87, 0x7e, 0x59, 0xe7, 0x10, 0x5a, 0x6b, 0xe2, 0x89, 0x5d, 0xa8, 0xe3, 0xec, 0x64, 0xe4,
	0x82, 0x76, 0x64, 0xba, 0x87, 0xaf, 0x20, 0xd0, 0x16, 0xc4, 0x63, 0x9b, 0xef, 0xb9, 0x33, 0xc1,
	0x3e, 0xd4, 0x15, 0xbc, 0xf4, 0x70, 0x06, 0x43, 0xe3, 0x27, 0x6c, 0x2f, 0xf4, 0xcd, 0x97, 0x9c,
	0x38, 0xd8, 0xdd, 0x98, 0x73, 0x08, 0xad, 0x03, 0x3a, 0x76, 0x8e, 0x95, 0x3a, 0x76, 0xae, 0x49,
	0xd2, 0x3d, 0x7c, 0x0e, 0x03, 0xfd, 0xce, 0x1f, 0xdb, 0xf4, 0xc6, 0x03, 0xc9, 0x71, 0x0f, 0xb9,
	0xbb, 0xfd, 0x09, 0x62, 0x47, 0x59, 0x4c, 0x5a, 0x32, 0xf7, 0x2c, 0x2c, 0x79, 0x7a, 0x6f, 0xae,
	0xeb, 0x35, 0x0f, 0xbe, 0xfa, 0xeb, 0xdb, 0xdb, 0xd0, 0xfc, 0x30, 0x5f, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x47, 0xc3, 0xff, 0x04, 0x3e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (ItemService_LoadClient, error)
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error)
	FinishAuction(ctx context.Context, in *FinishAuctionRequest, opts ...grpc.CallOption) (*FinishAuctionResponse, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (ItemService_LoadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ItemService_serviceDesc.Streams[0], "/item.ItemService/Load", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemServiceLoadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemService_LoadClient interface {
	Recv() (*LoadResponse, error)
	grpc.ClientStream
}

type itemServiceLoadClient struct {
	grpc.ClientStream
}

func (x *itemServiceLoadClient) Recv() (*LoadResponse, error) {
	m := new(LoadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error) {
	out := new(BidResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) FinishAuction(ctx context.Context, in *FinishAuctionRequest, opts ...grpc.CallOption) (*FinishAuctionResponse, error) {
	out := new(FinishAuctionResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/FinishAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
type ItemServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Load(*LoadRequest, ItemService_LoadServer) error
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Bid(context.Context, *BidRequest) (*BidResponse, error)
	FinishAuction(context.Context, *FinishAuctionRequest) (*FinishAuctionResponse, error)
}

// UnimplementedItemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedItemServiceServer struct {
}

func (*UnimplementedItemServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedItemServiceServer) Load(req *LoadRequest, srv ItemService_LoadServer) error {
	return status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (*UnimplementedItemServiceServer) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedItemServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedItemServiceServer) Bid(ctx context.Context, req *BidRequest) (*BidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (*UnimplementedItemServiceServer) FinishAuction(ctx context.Context, req *FinishAuctionRequest) (*FinishAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishAuction not implemented")
}

func RegisterItemServiceServer(s *grpc.Server, srv ItemServiceServer) {
	s.RegisterService(&_ItemService_serviceDesc, srv)
}

func _ItemService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_Load_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemServiceServer).Load(m, &itemServiceLoadServer{stream})
}

type ItemService_LoadServer interface {
	Send(*LoadResponse) error
	grpc.ServerStream
}

type itemServiceLoadServer struct {
	grpc.ServerStream
}

func (x *itemServiceLoadServer) Send(m *LoadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ItemService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Bid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_FinishAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishAuctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).FinishAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/FinishAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).FinishAuction(ctx, req.(*FinishAuctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "item.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ItemService_Create_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _ItemService_Fetch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ItemService_Delete_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _ItemService_Bid_Handler,
		},
		{
			MethodName: "FinishAuction",
			Handler:    _ItemService_FinishAuction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Load",
			Handler:       _ItemService_Load_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "item.proto",
}