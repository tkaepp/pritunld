// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: sevsnp.proto

// Package sevsnp represents an SEV-SNP attestation report and its certificate
// chain.

package sevsnp

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SevProduct_SevProductName int32

const (
	SevProduct_SEV_PRODUCT_UNKNOWN SevProduct_SevProductName = 0
	SevProduct_SEV_PRODUCT_MILAN   SevProduct_SevProductName = 1
	SevProduct_SEV_PRODUCT_GENOA   SevProduct_SevProductName = 2
)

// Enum value maps for SevProduct_SevProductName.
var (
	SevProduct_SevProductName_name = map[int32]string{
		0: "SEV_PRODUCT_UNKNOWN",
		1: "SEV_PRODUCT_MILAN",
		2: "SEV_PRODUCT_GENOA",
	}
	SevProduct_SevProductName_value = map[string]int32{
		"SEV_PRODUCT_UNKNOWN": 0,
		"SEV_PRODUCT_MILAN":   1,
		"SEV_PRODUCT_GENOA":   2,
	}
)

func (x SevProduct_SevProductName) Enum() *SevProduct_SevProductName {
	p := new(SevProduct_SevProductName)
	*p = x
	return p
}

func (x SevProduct_SevProductName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SevProduct_SevProductName) Descriptor() protoreflect.EnumDescriptor {
	return file_sevsnp_proto_enumTypes[0].Descriptor()
}

func (SevProduct_SevProductName) Type() protoreflect.EnumType {
	return &file_sevsnp_proto_enumTypes[0]
}

func (x SevProduct_SevProductName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SevProduct_SevProductName.Descriptor instead.
func (SevProduct_SevProductName) EnumDescriptor() ([]byte, []int) {
	return file_sevsnp_proto_rawDescGZIP(), []int{2, 0}
}

// Report represents an SEV-SNP ATTESTATION_REPORT, specified in SEV SNP API
//
//	documentation https://www.amd.com/system/files/TechDocs/56860.pdf
type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version         uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"` // Should be 1 for revision 1.51
	GuestSvn        uint32 `protobuf:"varint,2,opt,name=guest_svn,json=guestSvn,proto3" json:"guest_svn,omitempty"`
	Policy          uint64 `protobuf:"varint,3,opt,name=policy,proto3" json:"policy,omitempty"`
	FamilyId        []byte `protobuf:"bytes,4,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"` // Should be 16 bytes long
	ImageId         []byte `protobuf:"bytes,5,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`    // Should be 16 bytes long
	Vmpl            uint32 `protobuf:"varint,6,opt,name=vmpl,proto3" json:"vmpl,omitempty"`
	SignatureAlgo   uint32 `protobuf:"varint,7,opt,name=signature_algo,json=signatureAlgo,proto3" json:"signature_algo,omitempty"`
	CurrentTcb      uint64 `protobuf:"varint,8,opt,name=current_tcb,json=currentTcb,proto3" json:"current_tcb,omitempty"`
	PlatformInfo    uint64 `protobuf:"varint,9,opt,name=platform_info,json=platformInfo,proto3" json:"platform_info,omitempty"`
	SignerInfo      uint32 `protobuf:"varint,10,opt,name=signer_info,json=signerInfo,proto3" json:"signer_info,omitempty"`                 // AuthorKeyEn, MaskChipKey, SigningKey
	ReportData      []byte `protobuf:"bytes,11,opt,name=report_data,json=reportData,proto3" json:"report_data,omitempty"`                  // Should be 64 bytes long
	Measurement     []byte `protobuf:"bytes,12,opt,name=measurement,proto3" json:"measurement,omitempty"`                                  // Should be 48 bytes long
	HostData        []byte `protobuf:"bytes,13,opt,name=host_data,json=hostData,proto3" json:"host_data,omitempty"`                        // Should be 32 bytes long
	IdKeyDigest     []byte `protobuf:"bytes,14,opt,name=id_key_digest,json=idKeyDigest,proto3" json:"id_key_digest,omitempty"`             // Should be 48 bytes long
	AuthorKeyDigest []byte `protobuf:"bytes,15,opt,name=author_key_digest,json=authorKeyDigest,proto3" json:"author_key_digest,omitempty"` // Should be 48 bytes long
	ReportId        []byte `protobuf:"bytes,16,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`                        // Should be 32 bytes long
	ReportIdMa      []byte `protobuf:"bytes,17,opt,name=report_id_ma,json=reportIdMa,proto3" json:"report_id_ma,omitempty"`                // Should be 32 bytes long
	ReportedTcb     uint64 `protobuf:"varint,18,opt,name=reported_tcb,json=reportedTcb,proto3" json:"reported_tcb,omitempty"`
	ChipId          []byte `protobuf:"bytes,19,opt,name=chip_id,json=chipId,proto3" json:"chip_id,omitempty"` // Should be 64 bytes long
	CommittedTcb    uint64 `protobuf:"varint,20,opt,name=committed_tcb,json=committedTcb,proto3" json:"committed_tcb,omitempty"`
	// Each build, minor, major triple should be packed together in a uint32
	// packed together at 7:0, 15:8, 23:16 respectively
	CurrentBuild   uint32 `protobuf:"varint,21,opt,name=current_build,json=currentBuild,proto3" json:"current_build,omitempty"`
	CurrentMinor   uint32 `protobuf:"varint,22,opt,name=current_minor,json=currentMinor,proto3" json:"current_minor,omitempty"`
	CurrentMajor   uint32 `protobuf:"varint,23,opt,name=current_major,json=currentMajor,proto3" json:"current_major,omitempty"`
	CommittedBuild uint32 `protobuf:"varint,24,opt,name=committed_build,json=committedBuild,proto3" json:"committed_build,omitempty"`
	CommittedMinor uint32 `protobuf:"varint,25,opt,name=committed_minor,json=committedMinor,proto3" json:"committed_minor,omitempty"`
	CommittedMajor uint32 `protobuf:"varint,26,opt,name=committed_major,json=committedMajor,proto3" json:"committed_major,omitempty"`
	LaunchTcb      uint64 `protobuf:"varint,27,opt,name=launch_tcb,json=launchTcb,proto3" json:"launch_tcb,omitempty"`
	Signature      []byte `protobuf:"bytes,28,opt,name=signature,proto3" json:"signature,omitempty"` // Should be 512 bytes long
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sevsnp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_sevsnp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_sevsnp_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Report) GetGuestSvn() uint32 {
	if x != nil {
		return x.GuestSvn
	}
	return 0
}

func (x *Report) GetPolicy() uint64 {
	if x != nil {
		return x.Policy
	}
	return 0
}

func (x *Report) GetFamilyId() []byte {
	if x != nil {
		return x.FamilyId
	}
	return nil
}

func (x *Report) GetImageId() []byte {
	if x != nil {
		return x.ImageId
	}
	return nil
}

func (x *Report) GetVmpl() uint32 {
	if x != nil {
		return x.Vmpl
	}
	return 0
}

func (x *Report) GetSignatureAlgo() uint32 {
	if x != nil {
		return x.SignatureAlgo
	}
	return 0
}

func (x *Report) GetCurrentTcb() uint64 {
	if x != nil {
		return x.CurrentTcb
	}
	return 0
}

func (x *Report) GetPlatformInfo() uint64 {
	if x != nil {
		return x.PlatformInfo
	}
	return 0
}

func (x *Report) GetSignerInfo() uint32 {
	if x != nil {
		return x.SignerInfo
	}
	return 0
}

func (x *Report) GetReportData() []byte {
	if x != nil {
		return x.ReportData
	}
	return nil
}

func (x *Report) GetMeasurement() []byte {
	if x != nil {
		return x.Measurement
	}
	return nil
}

func (x *Report) GetHostData() []byte {
	if x != nil {
		return x.HostData
	}
	return nil
}

func (x *Report) GetIdKeyDigest() []byte {
	if x != nil {
		return x.IdKeyDigest
	}
	return nil
}

func (x *Report) GetAuthorKeyDigest() []byte {
	if x != nil {
		return x.AuthorKeyDigest
	}
	return nil
}

func (x *Report) GetReportId() []byte {
	if x != nil {
		return x.ReportId
	}
	return nil
}

func (x *Report) GetReportIdMa() []byte {
	if x != nil {
		return x.ReportIdMa
	}
	return nil
}

func (x *Report) GetReportedTcb() uint64 {
	if x != nil {
		return x.ReportedTcb
	}
	return 0
}

func (x *Report) GetChipId() []byte {
	if x != nil {
		return x.ChipId
	}
	return nil
}

func (x *Report) GetCommittedTcb() uint64 {
	if x != nil {
		return x.CommittedTcb
	}
	return 0
}

func (x *Report) GetCurrentBuild() uint32 {
	if x != nil {
		return x.CurrentBuild
	}
	return 0
}

func (x *Report) GetCurrentMinor() uint32 {
	if x != nil {
		return x.CurrentMinor
	}
	return 0
}

func (x *Report) GetCurrentMajor() uint32 {
	if x != nil {
		return x.CurrentMajor
	}
	return 0
}

func (x *Report) GetCommittedBuild() uint32 {
	if x != nil {
		return x.CommittedBuild
	}
	return 0
}

func (x *Report) GetCommittedMinor() uint32 {
	if x != nil {
		return x.CommittedMinor
	}
	return 0
}

func (x *Report) GetCommittedMajor() uint32 {
	if x != nil {
		return x.CommittedMajor
	}
	return 0
}

func (x *Report) GetLaunchTcb() uint64 {
	if x != nil {
		return x.LaunchTcb
	}
	return 0
}

func (x *Report) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type CertificateChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The versioned chip endorsement key's certificate for the
	// key that signed this report.
	VcekCert []byte `protobuf:"bytes,1,opt,name=vcek_cert,json=vcekCert,proto3" json:"vcek_cert,omitempty"`
	// The versioned loaded endorsement key's certificate for the
	// key that signed this report.
	VlekCert []byte `protobuf:"bytes,6,opt,name=vlek_cert,json=vlekCert,proto3" json:"vlek_cert,omitempty"`
	// The AMD SEV or AMD SEV-VLEK certificate that signed the V?EK cert.
	AskCert []byte `protobuf:"bytes,2,opt,name=ask_cert,json=askCert,proto3" json:"ask_cert,omitempty"`
	// The AMD Root key certificate (signs the ASK cert).
	ArkCert []byte `protobuf:"bytes,3,opt,name=ark_cert,json=arkCert,proto3" json:"ark_cert,omitempty"`
	// A certificate the host may inject to endorse the measurement of the
	// firmware.
	//
	// Deprecated: Marked as deprecated in sevsnp.proto.
	FirmwareCert []byte `protobuf:"bytes,4,opt,name=firmware_cert,json=firmwareCert,proto3" json:"firmware_cert,omitempty"`
	// Non-standard certificates the host may inject.
	Extras map[string][]byte `protobuf:"bytes,7,rep,name=extras,proto3" json:"extras,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CertificateChain) Reset() {
	*x = CertificateChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sevsnp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateChain) ProtoMessage() {}

func (x *CertificateChain) ProtoReflect() protoreflect.Message {
	mi := &file_sevsnp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateChain.ProtoReflect.Descriptor instead.
func (*CertificateChain) Descriptor() ([]byte, []int) {
	return file_sevsnp_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateChain) GetVcekCert() []byte {
	if x != nil {
		return x.VcekCert
	}
	return nil
}

func (x *CertificateChain) GetVlekCert() []byte {
	if x != nil {
		return x.VlekCert
	}
	return nil
}

func (x *CertificateChain) GetAskCert() []byte {
	if x != nil {
		return x.AskCert
	}
	return nil
}

func (x *CertificateChain) GetArkCert() []byte {
	if x != nil {
		return x.ArkCert
	}
	return nil
}

// Deprecated: Marked as deprecated in sevsnp.proto.
func (x *CertificateChain) GetFirmwareCert() []byte {
	if x != nil {
		return x.FirmwareCert
	}
	return nil
}

func (x *CertificateChain) GetExtras() map[string][]byte {
	if x != nil {
		return x.Extras
	}
	return nil
}

// The CPUID[EAX=1] version information includes product info as described in
// the AMD KDS specification. The product name, model, and stepping values are
// important for determining the required parameters to KDS when requesting the
// endorsement key's certificate.
type SevProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name SevProduct_SevProductName `protobuf:"varint,1,opt,name=name,proto3,enum=sevsnp.SevProduct_SevProductName" json:"name,omitempty"`
	// Deprecated: Marked as deprecated in sevsnp.proto.
	Stepping        uint32                `protobuf:"varint,2,opt,name=stepping,proto3" json:"stepping,omitempty"` // Must be a 4-bit number
	MachineStepping *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=machine_stepping,json=machineStepping,proto3" json:"machine_stepping,omitempty"`
}

func (x *SevProduct) Reset() {
	*x = SevProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sevsnp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SevProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SevProduct) ProtoMessage() {}

func (x *SevProduct) ProtoReflect() protoreflect.Message {
	mi := &file_sevsnp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SevProduct.ProtoReflect.Descriptor instead.
func (*SevProduct) Descriptor() ([]byte, []int) {
	return file_sevsnp_proto_rawDescGZIP(), []int{2}
}

func (x *SevProduct) GetName() SevProduct_SevProductName {
	if x != nil {
		return x.Name
	}
	return SevProduct_SEV_PRODUCT_UNKNOWN
}

// Deprecated: Marked as deprecated in sevsnp.proto.
func (x *SevProduct) GetStepping() uint32 {
	if x != nil {
		return x.Stepping
	}
	return 0
}

func (x *SevProduct) GetMachineStepping() *wrappers.UInt32Value {
	if x != nil {
		return x.MachineStepping
	}
	return nil
}

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report           *Report           `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	CertificateChain *CertificateChain `protobuf:"bytes,2,opt,name=certificate_chain,json=certificateChain,proto3" json:"certificate_chain,omitempty"`
	Product          *SevProduct       `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sevsnp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_sevsnp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_sevsnp_proto_rawDescGZIP(), []int{3}
}

func (x *Attestation) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *Attestation) GetCertificateChain() *CertificateChain {
	if x != nil {
		return x.CertificateChain
	}
	return nil
}

func (x *Attestation) GetProduct() *SevProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_sevsnp_proto protoreflect.FileDescriptor

var file_sevsnp_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x07, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x53, 0x76, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6d, 0x70, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x76, 0x6d, 0x70, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41,
	0x6c, 0x67, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x63, 0x62, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x54, 0x63, 0x62, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x64,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x69, 0x64, 0x4b, 0x65, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x4b, 0x65, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x4d, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x63, 0x62, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x63, 0x62, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63,
	0x68, 0x69, 0x70, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x63, 0x62, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x54, 0x63, 0x62, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x6f, 0x72,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d,
	0x69, 0x6e, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f,
	0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x4d,
	0x61, 0x6a, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x74,
	0x63, 0x62, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x54, 0x63, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0xa4, 0x02, 0x0a, 0x10, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x63, 0x65, 0x6b, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x76, 0x63, 0x65, 0x6b, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6c, 0x65, 0x6b, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x76, 0x6c, 0x65, 0x6b, 0x43, 0x65, 0x72, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x61, 0x73, 0x6b, 0x43, 0x65, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x72, 0x6b, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61,
	0x72, 0x6b, 0x43, 0x65, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x43, 0x65, 0x72, 0x74, 0x12,
	0x3c, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x45, 0x78, 0x74, 0x72, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x85, 0x02, 0x0a, 0x0a, 0x53, 0x65, 0x76,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x2e, 0x53,
	0x65, 0x76, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x76, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x08, 0x73, 0x74, 0x65, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x47,
	0x0a, 0x10, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x65, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x57, 0x0a, 0x0e, 0x53, 0x65, 0x76, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x56,
	0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x56, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x4d, 0x49, 0x4c, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x56,
	0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x47, 0x45, 0x4e, 0x4f, 0x41, 0x10, 0x02,
	0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x10, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x2e, 0x53, 0x65, 0x76, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x76, 0x2d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x76, 0x73, 0x6e, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sevsnp_proto_rawDescOnce sync.Once
	file_sevsnp_proto_rawDescData = file_sevsnp_proto_rawDesc
)

func file_sevsnp_proto_rawDescGZIP() []byte {
	file_sevsnp_proto_rawDescOnce.Do(func() {
		file_sevsnp_proto_rawDescData = protoimpl.X.CompressGZIP(file_sevsnp_proto_rawDescData)
	})
	return file_sevsnp_proto_rawDescData
}

var file_sevsnp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sevsnp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sevsnp_proto_goTypes = []interface{}{
	(SevProduct_SevProductName)(0), // 0: sevsnp.SevProduct.SevProductName
	(*Report)(nil),                 // 1: sevsnp.Report
	(*CertificateChain)(nil),       // 2: sevsnp.CertificateChain
	(*SevProduct)(nil),             // 3: sevsnp.SevProduct
	(*Attestation)(nil),            // 4: sevsnp.Attestation
	nil,                            // 5: sevsnp.CertificateChain.ExtrasEntry
	(*wrappers.UInt32Value)(nil),   // 6: google.protobuf.UInt32Value
}
var file_sevsnp_proto_depIdxs = []int32{
	5, // 0: sevsnp.CertificateChain.extras:type_name -> sevsnp.CertificateChain.ExtrasEntry
	0, // 1: sevsnp.SevProduct.name:type_name -> sevsnp.SevProduct.SevProductName
	6, // 2: sevsnp.SevProduct.machine_stepping:type_name -> google.protobuf.UInt32Value
	1, // 3: sevsnp.Attestation.report:type_name -> sevsnp.Report
	2, // 4: sevsnp.Attestation.certificate_chain:type_name -> sevsnp.CertificateChain
	3, // 5: sevsnp.Attestation.product:type_name -> sevsnp.SevProduct
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_sevsnp_proto_init() }
func file_sevsnp_proto_init() {
	if File_sevsnp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sevsnp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sevsnp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateChain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sevsnp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SevProduct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sevsnp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sevsnp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sevsnp_proto_goTypes,
		DependencyIndexes: file_sevsnp_proto_depIdxs,
		EnumInfos:         file_sevsnp_proto_enumTypes,
		MessageInfos:      file_sevsnp_proto_msgTypes,
	}.Build()
	File_sevsnp_proto = out.File
	file_sevsnp_proto_rawDesc = nil
	file_sevsnp_proto_goTypes = nil
	file_sevsnp_proto_depIdxs = nil
}