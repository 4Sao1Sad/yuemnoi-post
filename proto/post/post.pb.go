// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: proto/post/post.proto

package yuemnoi_post

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LendingPost message representing the structure of the LendingPost object
type LendingPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemName     string                 `protobuf:"bytes,2,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price        float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	ActiveStatus bool                   `protobuf:"varint,5,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"`
	ImageUrl     string                 `protobuf:"bytes,6,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	OwnerId      uint64                 `protobuf:"varint,7,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *LendingPost) Reset() {
	*x = LendingPost{}
	mi := &file_proto_post_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LendingPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LendingPost) ProtoMessage() {}

func (x *LendingPost) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LendingPost.ProtoReflect.Descriptor instead.
func (*LendingPost) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *LendingPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LendingPost) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *LendingPost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LendingPost) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *LendingPost) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

func (x *LendingPost) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *LendingPost) GetOwnerId() uint64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *LendingPost) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Request message for creating a new LendingPost
type CreateLendingPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemName     string  `protobuf:"bytes,1,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Description  string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price        float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	ActiveStatus bool    `protobuf:"varint,4,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"`
	ImageUrl     string  `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *CreateLendingPostRequest) Reset() {
	*x = CreateLendingPostRequest{}
	mi := &file_proto_post_post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLendingPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLendingPostRequest) ProtoMessage() {}

func (x *CreateLendingPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLendingPostRequest.ProtoReflect.Descriptor instead.
func (*CreateLendingPostRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLendingPostRequest) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *CreateLendingPostRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateLendingPostRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateLendingPostRequest) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

func (x *CreateLendingPostRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

// Response message for the CreateLendingPost service
type CreateLendingPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateLendingPostResponse) Reset() {
	*x = CreateLendingPostResponse{}
	mi := &file_proto_post_post_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLendingPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLendingPostResponse) ProtoMessage() {}

func (x *CreateLendingPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLendingPostResponse.ProtoReflect.Descriptor instead.
func (*CreateLendingPostResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLendingPostResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateLendingPostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request message for fetching LendingPost details
type GetLendingPostDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLendingPostDetailRequest) Reset() {
	*x = GetLendingPostDetailRequest{}
	mi := &file_proto_post_post_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLendingPostDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLendingPostDetailRequest) ProtoMessage() {}

func (x *GetLendingPostDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLendingPostDetailRequest.ProtoReflect.Descriptor instead.
func (*GetLendingPostDetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetLendingPostDetailRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Request message for searching LendingPost
type SearchLendingPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchString string `protobuf:"bytes,1,opt,name=search_string,json=searchString,proto3" json:"search_string,omitempty"` // Snake_case for field names is recommended in proto3
}

func (x *SearchLendingPostRequest) Reset() {
	*x = SearchLendingPostRequest{}
	mi := &file_proto_post_post_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchLendingPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLendingPostRequest) ProtoMessage() {}

func (x *SearchLendingPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLendingPostRequest.ProtoReflect.Descriptor instead.
func (*SearchLendingPostRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{4}
}

func (x *SearchLendingPostRequest) GetSearchString() string {
	if x != nil {
		return x.SearchString
	}
	return ""
}

// Response message for a list of LendingPosts
type LendingPostList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*LendingPost `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"` // Use `repeated` for arrays of LendingPost
}

func (x *LendingPostList) Reset() {
	*x = LendingPostList{}
	mi := &file_proto_post_post_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LendingPostList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LendingPostList) ProtoMessage() {}

func (x *LendingPostList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LendingPostList.ProtoReflect.Descriptor instead.
func (*LendingPostList) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{5}
}

func (x *LendingPostList) GetPosts() []*LendingPost {
	if x != nil {
		return x.Posts
	}
	return nil
}

// LendingPost message representing the structure of the BorrowingPost object
type BorrowingPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description  string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`                        // Required field
	ActiveStatus bool                   `protobuf:"varint,3,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"` // Required field, booleanepeated LendingPost item_list = 4; // List of BorrowingPost
	OwnerId      uint64                 `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`                // Required field
	Timestamp    int64                  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                           // Required field, using int64 for timestamps
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *BorrowingPost) Reset() {
	*x = BorrowingPost{}
	mi := &file_proto_post_post_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowingPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowingPost) ProtoMessage() {}

func (x *BorrowingPost) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowingPost.ProtoReflect.Descriptor instead.
func (*BorrowingPost) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{6}
}

func (x *BorrowingPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BorrowingPost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BorrowingPost) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

func (x *BorrowingPost) GetOwnerId() uint64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *BorrowingPost) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BorrowingPost) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Request message for creating a new BorrowingPost
type CreateBorrowingPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description  string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`                        // Optional: Description of the item
	ActiveStatus bool   `protobuf:"varint,2,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"` // Required: Active status of the post
	OwnerId      uint64 `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`                // Required: ID of the owner (can be a UUID)
}

func (x *CreateBorrowingPostRequest) Reset() {
	*x = CreateBorrowingPostRequest{}
	mi := &file_proto_post_post_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBorrowingPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBorrowingPostRequest) ProtoMessage() {}

func (x *CreateBorrowingPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBorrowingPostRequest.ProtoReflect.Descriptor instead.
func (*CreateBorrowingPostRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{7}
}

func (x *CreateBorrowingPostRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBorrowingPostRequest) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

func (x *CreateBorrowingPostRequest) GetOwnerId() uint64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

// Response message for the CreateBorrowingPost service
type CreateBorrowingPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          // The autogenerated ID of the created post
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // A success message or error details
}

func (x *CreateBorrowingPostResponse) Reset() {
	*x = CreateBorrowingPostResponse{}
	mi := &file_proto_post_post_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBorrowingPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBorrowingPostResponse) ProtoMessage() {}

func (x *CreateBorrowingPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBorrowingPostResponse.ProtoReflect.Descriptor instead.
func (*CreateBorrowingPostResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{8}
}

func (x *CreateBorrowingPostResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateBorrowingPostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetBorrowingPostDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBorrowingPostDetailRequest) Reset() {
	*x = GetBorrowingPostDetailRequest{}
	mi := &file_proto_post_post_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBorrowingPostDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBorrowingPostDetailRequest) ProtoMessage() {}

func (x *GetBorrowingPostDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBorrowingPostDetailRequest.ProtoReflect.Descriptor instead.
func (*GetBorrowingPostDetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{9}
}

func (x *GetBorrowingPostDetailRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchBorrowingPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchString string `protobuf:"bytes,1,opt,name=searchString,proto3" json:"searchString,omitempty"`
}

func (x *SearchBorrowingPostRequest) Reset() {
	*x = SearchBorrowingPostRequest{}
	mi := &file_proto_post_post_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchBorrowingPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBorrowingPostRequest) ProtoMessage() {}

func (x *SearchBorrowingPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBorrowingPostRequest.ProtoReflect.Descriptor instead.
func (*SearchBorrowingPostRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{10}
}

func (x *SearchBorrowingPostRequest) GetSearchString() string {
	if x != nil {
		return x.SearchString
	}
	return ""
}

type BorrowingPostList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BorrowingPost []*BorrowingPost `protobuf:"bytes,1,rep,name=borrowing_post,json=borrowingPost,proto3" json:"borrowing_post,omitempty"`
}

func (x *BorrowingPostList) Reset() {
	*x = BorrowingPostList{}
	mi := &file_proto_post_post_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowingPostList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowingPostList) ProtoMessage() {}

func (x *BorrowingPostList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowingPostList.ProtoReflect.Descriptor instead.
func (*BorrowingPostList) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{11}
}

func (x *BorrowingPostList) GetBorrowingPost() []*BorrowingPost {
	if x != nil {
		return x.BorrowingPost
	}
	return nil
}

type GetBorrowingPostByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBorrowingPostByIdRequest) Reset() {
	*x = GetBorrowingPostByIdRequest{}
	mi := &file_proto_post_post_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBorrowingPostByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBorrowingPostByIdRequest) ProtoMessage() {}

func (x *GetBorrowingPostByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBorrowingPostByIdRequest.ProtoReflect.Descriptor instead.
func (*GetBorrowingPostByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{12}
}

func (x *GetBorrowingPostByIdRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_post_post_proto protoreflect.FileDescriptor

var file_proto_post_post_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x0b, 0x4c, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x45, 0x0a, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x2d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x3f, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x22, 0x35, 0x0a, 0x0f, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x7e, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40,
	0x0a, 0x1a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x22, 0x4a, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x0d, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xec, 0x01, 0x0a, 0x12,
	0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x32, 0x80, 0x02, 0x0a, 0x14, 0x42,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x1e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34, 0x53, 0x61, 0x6f,
	0x31, 0x53, 0x61, 0x64, 0x2f, 0x79, 0x75, 0x65, 0x6d, 0x6e, 0x6f, 0x69, 0x2d, 0x70, 0x6f, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_post_post_proto_rawDescOnce sync.Once
	file_proto_post_post_proto_rawDescData = file_proto_post_post_proto_rawDesc
)

func file_proto_post_post_proto_rawDescGZIP() []byte {
	file_proto_post_post_proto_rawDescOnce.Do(func() {
		file_proto_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_post_post_proto_rawDescData)
	})
	return file_proto_post_post_proto_rawDescData
}

var file_proto_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_post_post_proto_goTypes = []any{
	(*LendingPost)(nil),                   // 0: LendingPost
	(*CreateLendingPostRequest)(nil),      // 1: CreateLendingPostRequest
	(*CreateLendingPostResponse)(nil),     // 2: CreateLendingPostResponse
	(*GetLendingPostDetailRequest)(nil),   // 3: GetLendingPostDetailRequest
	(*SearchLendingPostRequest)(nil),      // 4: SearchLendingPostRequest
	(*LendingPostList)(nil),               // 5: LendingPostList
	(*BorrowingPost)(nil),                 // 6: BorrowingPost
	(*CreateBorrowingPostRequest)(nil),    // 7: CreateBorrowingPostRequest
	(*CreateBorrowingPostResponse)(nil),   // 8: CreateBorrowingPostResponse
	(*GetBorrowingPostDetailRequest)(nil), // 9: GetBorrowingPostDetailRequest
	(*SearchBorrowingPostRequest)(nil),    // 10: SearchBorrowingPostRequest
	(*BorrowingPostList)(nil),             // 11: BorrowingPostList
	(*GetBorrowingPostByIdRequest)(nil),   // 12: GetBorrowingPostByIdRequest
	(*timestamppb.Timestamp)(nil),         // 13: google.protobuf.Timestamp
}
var file_proto_post_post_proto_depIdxs = []int32{
	13, // 0: LendingPost.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 1: LendingPostList.posts:type_name -> LendingPost
	13, // 2: BorrowingPost.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 3: BorrowingPostList.borrowing_post:type_name -> BorrowingPost
	1,  // 4: LendingPostService.CreateLendingPost:input_type -> CreateLendingPostRequest
	3,  // 5: LendingPostService.GetLendingPostDetail:input_type -> GetLendingPostDetailRequest
	4,  // 6: LendingPostService.SearchLendingPost:input_type -> SearchLendingPostRequest
	7,  // 7: BorrowingPostService.CreateBorrowingPost:input_type -> CreateBorrowingPostRequest
	9,  // 8: BorrowingPostService.GetBorrowingPostDetail:input_type -> GetBorrowingPostDetailRequest
	10, // 9: BorrowingPostService.SearchBorrowingPost:input_type -> SearchBorrowingPostRequest
	2,  // 10: LendingPostService.CreateLendingPost:output_type -> CreateLendingPostResponse
	0,  // 11: LendingPostService.GetLendingPostDetail:output_type -> LendingPost
	5,  // 12: LendingPostService.SearchLendingPost:output_type -> LendingPostList
	8,  // 13: BorrowingPostService.CreateBorrowingPost:output_type -> CreateBorrowingPostResponse
	6,  // 14: BorrowingPostService.GetBorrowingPostDetail:output_type -> BorrowingPost
	11, // 15: BorrowingPostService.SearchBorrowingPost:output_type -> BorrowingPostList
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_post_post_proto_init() }
func file_proto_post_post_proto_init() {
	if File_proto_post_post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_post_post_proto_goTypes,
		DependencyIndexes: file_proto_post_post_proto_depIdxs,
		MessageInfos:      file_proto_post_post_proto_msgTypes,
	}.Build()
	File_proto_post_post_proto = out.File
	file_proto_post_post_proto_rawDesc = nil
	file_proto_post_post_proto_goTypes = nil
	file_proto_post_post_proto_depIdxs = nil
}
