// Code generated by protoc-gen-gogo.
// source: unit.proto
// DO NOT EDIT!

/*
	Package unit is a generated protocol buffer package.

	It is generated from these files:
		unit.proto

	It has these top-level messages:
		Key
		SourceUnit
		Info
		Resolution
*/
package unit

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Key struct {
	// Repo is the URI of the repository containing this source unit, if any.
	// The scanner tool does not need to set this field - it can be left blank,
	// to be filled in by the `srclib` tool.
	//
	// If Repo is empty, it indicates that the repository URI is
	// purposefully omitted and this field should be treated as if it
	// doesn't exist. If Repo is set to the unresolved repo sentinel
	// value, then it indicates that repository is unknown, but this
	// field value can be used.
	Repo string `protobuf:"bytes,1,opt,name=Repo,proto3" json:"Repo,omitempty"`
	// CommitID is the commit ID of the repository containing this
	// source unit, if any. The scanner tool need not fill this in; it
	// should be left blank, to be filled in by the `srclib` tool.
	CommitID string `protobuf:"bytes,2,opt,name=CommitID,proto3" json:"CommitID,omitempty"`
	// Version is the unresolved source unit version (e.g., "v1.2.3").
	// When empty, it indicates that no version is specified for the
	// source unit. Currently, this field is unused, but can still be
	// set for the sake of posterity.
	Version string `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	// Type is the type of source unit this represents, such as "GoPackage".
	Type string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	// Name is an opaque identifier for this source unit that MUST be unique
	// among all other source units of the same type in the same repository.
	//
	// Two source units of different types in a repository may have the same name.
	// To obtain an identifier for a source unit that is guaranteed to be unique
	// repository-wide, use the ID method.
	Name string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}

type SourceUnit struct {
	Key  `protobuf:"bytes,1,opt,name=key,embedded=key" json:"key"`
	Info `protobuf:"bytes,2,opt,name=info,embedded=info" json:"info"`
}

func (m *SourceUnit) Reset()         { *m = SourceUnit{} }
func (m *SourceUnit) String() string { return proto.CompactTextString(m) }
func (*SourceUnit) ProtoMessage()    {}

type Info struct {
	// Files is all of the files that make up this source unit. Filepaths should
	// be relative to the repository root.
	Files []string `protobuf:"bytes,1,rep,name=Files" json:"Files,omitempty"`
	// Dir is the root directory of this source unit. It is optional and maybe
	// empty.
	Dir string `protobuf:"bytes,2,opt,name=Dir,proto3" json:"Dir,omitempty"`
	// Dependencies is a list of dependencies that this source unit has. The
	// schema for these dependencies is internal to the scanner that produced
	// this source unit. The dependency resolver is expected to know how to
	// interpret this schema.
	//
	// The dependency information stored in this field should be able to be very
	// quickly determined by the scanner. The scanner should not perform any
	// dependency resolution on these entries. This is because the scanner is
	// run frequently and should execute very quickly, and dependency resolution
	// is often slow (requiring network access, etc.).
	Dependencies []*Key `protobuf:"bytes,3,rep,name=Dependencies" json:"Dependencies,omitempty"`
	// Data is additional data dumped by the scanner about this source unit. It
	// typically holds information that the scanner wants to make available to
	// other components in the toolchain (grapher, dep resolver, etc.).
	Data []byte `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	// Config is an arbitrary key-value property map. The Config map from the
	// tree config is copied verbatim to each source unit. It can be used to
	// pass options from the Srcfile to tools.
	//
	// DEPRECATED
	Config map[string]string `protobuf:"bytes,5,rep,name=Config" json:"Config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Ops is a deprecated field kept around for backcompat purposes. It
	// can be removed once the "graph-all" option has been removed.
	//
	// DEPRECATED
	Ops map[string][]byte `protobuf:"bytes,6,rep,name=Ops" json:"Ops,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}

type Resolution struct {
	Raw      Key `protobuf:"bytes,1,opt,name=raw" json:"raw"`
	Resolved Key `protobuf:"bytes,2,opt,name=resolved" json:"resolved"`
}

func (m *Resolution) Reset()         { *m = Resolution{} }
func (m *Resolution) String() string { return proto.CompactTextString(m) }
func (*Resolution) ProtoMessage()    {}

func (m *Key) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Key) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Repo) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintUnit(data, i, uint64(len(m.Repo)))
		i += copy(data[i:], m.Repo)
	}
	if len(m.CommitID) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintUnit(data, i, uint64(len(m.CommitID)))
		i += copy(data[i:], m.CommitID)
	}
	if len(m.Version) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintUnit(data, i, uint64(len(m.Version)))
		i += copy(data[i:], m.Version)
	}
	if len(m.Type) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintUnit(data, i, uint64(len(m.Type)))
		i += copy(data[i:], m.Type)
	}
	if len(m.Name) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintUnit(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	return i, nil
}

func (m *SourceUnit) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SourceUnit) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintUnit(data, i, uint64(m.Key.Size()))
	n1, err := m.Key.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintUnit(data, i, uint64(m.Info.Size()))
	n2, err := m.Info.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *Info) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Info) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Files) > 0 {
		for _, s := range m.Files {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Dir) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintUnit(data, i, uint64(len(m.Dir)))
		i += copy(data[i:], m.Dir)
	}
	if len(m.Dependencies) > 0 {
		for _, msg := range m.Dependencies {
			data[i] = 0x1a
			i++
			i = encodeVarintUnit(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Data != nil {
		if len(m.Data) > 0 {
			data[i] = 0x22
			i++
			i = encodeVarintUnit(data, i, uint64(len(m.Data)))
			i += copy(data[i:], m.Data)
		}
	}
	if len(m.Config) > 0 {
		keysForConfig := make([]string, 0, len(m.Config))
		for k, _ := range m.Config {
			keysForConfig = append(keysForConfig, k)
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForConfig)
		for _, k := range keysForConfig {
			data[i] = 0x2a
			i++
			v := m.Config[k]
			mapSize := 1 + len(k) + sovUnit(uint64(len(k))) + 1 + len(v) + sovUnit(uint64(len(v)))
			i = encodeVarintUnit(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintUnit(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintUnit(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	if len(m.Ops) > 0 {
		keysForOps := make([]string, 0, len(m.Ops))
		for k, _ := range m.Ops {
			keysForOps = append(keysForOps, k)
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForOps)
		for _, k := range keysForOps {
			data[i] = 0x32
			i++
			v := m.Ops[k]
			mapSize := 1 + len(k) + sovUnit(uint64(len(k))) + 1 + len(v) + sovUnit(uint64(len(v)))
			i = encodeVarintUnit(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintUnit(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintUnit(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	return i, nil
}

func (m *Resolution) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Resolution) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintUnit(data, i, uint64(m.Raw.Size()))
	n3, err := m.Raw.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	data[i] = 0x12
	i++
	i = encodeVarintUnit(data, i, uint64(m.Resolved.Size()))
	n4, err := m.Resolved.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeFixed64Unit(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Unit(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUnit(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Key) Size() (n int) {
	var l int
	_ = l
	l = len(m.Repo)
	if l > 0 {
		n += 1 + l + sovUnit(uint64(l))
	}
	l = len(m.CommitID)
	if l > 0 {
		n += 1 + l + sovUnit(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovUnit(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovUnit(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnit(uint64(l))
	}
	return n
}

func (m *SourceUnit) Size() (n int) {
	var l int
	_ = l
	l = m.Key.Size()
	n += 1 + l + sovUnit(uint64(l))
	l = m.Info.Size()
	n += 1 + l + sovUnit(uint64(l))
	return n
}

func (m *Info) Size() (n int) {
	var l int
	_ = l
	if len(m.Files) > 0 {
		for _, s := range m.Files {
			l = len(s)
			n += 1 + l + sovUnit(uint64(l))
		}
	}
	l = len(m.Dir)
	if l > 0 {
		n += 1 + l + sovUnit(uint64(l))
	}
	if len(m.Dependencies) > 0 {
		for _, e := range m.Dependencies {
			l = e.Size()
			n += 1 + l + sovUnit(uint64(l))
		}
	}
	if m.Data != nil {
		l = len(m.Data)
		if l > 0 {
			n += 1 + l + sovUnit(uint64(l))
		}
	}
	if len(m.Config) > 0 {
		for k, v := range m.Config {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovUnit(uint64(len(k))) + 1 + len(v) + sovUnit(uint64(len(v)))
			n += mapEntrySize + 1 + sovUnit(uint64(mapEntrySize))
		}
	}
	if len(m.Ops) > 0 {
		for k, v := range m.Ops {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovUnit(uint64(len(k))) + 1 + len(v) + sovUnit(uint64(len(v)))
			n += mapEntrySize + 1 + sovUnit(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Resolution) Size() (n int) {
	var l int
	_ = l
	l = m.Raw.Size()
	n += 1 + l + sovUnit(uint64(l))
	l = m.Resolved.Size()
	n += 1 + l + sovUnit(uint64(l))
	return n
}

func sovUnit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUnit(x uint64) (n int) {
	return sovUnit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Key) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repo = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnit(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SourceUnit) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SourceUnit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SourceUnit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Key.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnit(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Info) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Files", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Files = append(m.Files, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dir = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependencies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dependencies = append(m.Dependencies, &Key{})
			if err := m.Dependencies[len(m.Dependencies)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthUnit
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthUnit
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.Config == nil {
				m.Config = make(map[string]string)
			}
			m.Config[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ops", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthUnit
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapbyteLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapbyteLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intMapbyteLen := int(mapbyteLen)
			if intMapbyteLen < 0 {
				return ErrInvalidLengthUnit
			}
			postbytesIndex := iNdEx + intMapbyteLen
			if postbytesIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := make([]byte, mapbyteLen)
			copy(mapvalue, data[iNdEx:postbytesIndex])
			iNdEx = postbytesIndex
			if m.Ops == nil {
				m.Ops = make(map[string][]byte)
			}
			m.Ops[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnit(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Resolution) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Resolution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resolution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Raw.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resolved", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resolved.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnit(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUnit(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnit
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUnit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUnit
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUnit(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUnit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnit   = fmt.Errorf("proto: integer overflow")
)
