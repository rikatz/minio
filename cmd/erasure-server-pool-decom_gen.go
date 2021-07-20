package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DrainInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "st":
			z.StartTime, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "ssm":
			z.StartSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "StartSize")
				return
			}
		case "d":
			z.Duration, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "csm":
			z.CurrentSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "CurrentSize")
				return
			}
		case "cmp":
			z.Complete, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Complete")
				return
			}
		case "fl":
			z.Failed, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Failed")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *DrainInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "st"
	err = en.Append(0x86, 0xa2, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteTime(z.StartTime)
	if err != nil {
		err = msgp.WrapError(err, "StartTime")
		return
	}
	// write "ssm"
	err = en.Append(0xa3, 0x73, 0x73, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.StartSize)
	if err != nil {
		err = msgp.WrapError(err, "StartSize")
		return
	}
	// write "d"
	err = en.Append(0xa1, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Duration)
	if err != nil {
		err = msgp.WrapError(err, "Duration")
		return
	}
	// write "csm"
	err = en.Append(0xa3, 0x63, 0x73, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.CurrentSize)
	if err != nil {
		err = msgp.WrapError(err, "CurrentSize")
		return
	}
	// write "cmp"
	err = en.Append(0xa3, 0x63, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Complete)
	if err != nil {
		err = msgp.WrapError(err, "Complete")
		return
	}
	// write "fl"
	err = en.Append(0xa2, 0x66, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Failed)
	if err != nil {
		err = msgp.WrapError(err, "Failed")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DrainInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "st"
	o = append(o, 0x86, 0xa2, 0x73, 0x74)
	o = msgp.AppendTime(o, z.StartTime)
	// string "ssm"
	o = append(o, 0xa3, 0x73, 0x73, 0x6d)
	o = msgp.AppendInt64(o, z.StartSize)
	// string "d"
	o = append(o, 0xa1, 0x64)
	o = msgp.AppendInt64(o, z.Duration)
	// string "csm"
	o = append(o, 0xa3, 0x63, 0x73, 0x6d)
	o = msgp.AppendInt64(o, z.CurrentSize)
	// string "cmp"
	o = append(o, 0xa3, 0x63, 0x6d, 0x70)
	o = msgp.AppendBool(o, z.Complete)
	// string "fl"
	o = append(o, 0xa2, 0x66, 0x6c)
	o = msgp.AppendBool(o, z.Failed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DrainInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "st":
			z.StartTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "ssm":
			z.StartSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartSize")
				return
			}
		case "d":
			z.Duration, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "csm":
			z.CurrentSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CurrentSize")
				return
			}
		case "cmp":
			z.Complete, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Complete")
				return
			}
		case "fl":
			z.Failed, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Failed")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DrainInfo) Msgsize() (s int) {
	s = 1 + 3 + msgp.TimeSize + 4 + msgp.Int64Size + 2 + msgp.Int64Size + 4 + msgp.Int64Size + 4 + msgp.BoolSize + 3 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PoolInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.ID, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "cli":
			z.CmdLine, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "CmdLine")
				return
			}
		case "lu":
			z.LastUpdate, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "dr":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Drain")
					return
				}
				z.Drain = nil
			} else {
				if z.Drain == nil {
					z.Drain = new(DrainInfo)
				}
				err = z.Drain.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Drain")
					return
				}
			}
		case "sp":
			z.Suspend, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Suspend")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// omitempty: check for empty values
	zb0001Len := uint32(5)
	var zb0001Mask uint8 /* 5 bits */
	if z.CmdLine == "" {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if z.Drain == nil {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "id"
	err = en.Append(0xa2, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt(z.ID)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	if (zb0001Mask & 0x2) == 0 { // if not empty
		// write "cli"
		err = en.Append(0xa3, 0x63, 0x6c, 0x69)
		if err != nil {
			return
		}
		err = en.WriteString(z.CmdLine)
		if err != nil {
			err = msgp.WrapError(err, "CmdLine")
			return
		}
	}
	// write "lu"
	err = en.Append(0xa2, 0x6c, 0x75)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LastUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastUpdate")
		return
	}
	if (zb0001Mask & 0x8) == 0 { // if not empty
		// write "dr"
		err = en.Append(0xa2, 0x64, 0x72)
		if err != nil {
			return
		}
		if z.Drain == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Drain.EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Drain")
				return
			}
		}
	}
	// write "sp"
	err = en.Append(0xa2, 0x73, 0x70)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Suspend)
	if err != nil {
		err = msgp.WrapError(err, "Suspend")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(5)
	var zb0001Mask uint8 /* 5 bits */
	if z.CmdLine == "" {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if z.Drain == nil {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt(o, z.ID)
	if (zb0001Mask & 0x2) == 0 { // if not empty
		// string "cli"
		o = append(o, 0xa3, 0x63, 0x6c, 0x69)
		o = msgp.AppendString(o, z.CmdLine)
	}
	// string "lu"
	o = append(o, 0xa2, 0x6c, 0x75)
	o = msgp.AppendTime(o, z.LastUpdate)
	if (zb0001Mask & 0x8) == 0 { // if not empty
		// string "dr"
		o = append(o, 0xa2, 0x64, 0x72)
		if z.Drain == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Drain.MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Drain")
				return
			}
		}
	}
	// string "sp"
	o = append(o, 0xa2, 0x73, 0x70)
	o = msgp.AppendBool(o, z.Suspend)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "cli":
			z.CmdLine, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CmdLine")
				return
			}
		case "lu":
			z.LastUpdate, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "dr":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Drain = nil
			} else {
				if z.Drain == nil {
					z.Drain = new(DrainInfo)
				}
				bts, err = z.Drain.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Drain")
					return
				}
			}
		case "sp":
			z.Suspend, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Suspend")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolInfo) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 4 + msgp.StringPrefixSize + len(z.CmdLine) + 3 + msgp.TimeSize + 3
	if z.Drain == nil {
		s += msgp.NilSize
	} else {
		s += z.Drain.Msgsize()
	}
	s += 3 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *poolMeta) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "pls":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Pools")
				return
			}
			if cap(z.Pools) >= int(zb0002) {
				z.Pools = (z.Pools)[:zb0002]
			} else {
				z.Pools = make([]PoolInfo, zb0002)
			}
			for za0001 := range z.Pools {
				err = z.Pools[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Pools", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *poolMeta) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "v"
	err = en.Append(0x82, 0xa1, 0x76)
	if err != nil {
		return
	}
	err = en.WriteString(z.Version)
	if err != nil {
		err = msgp.WrapError(err, "Version")
		return
	}
	// write "pls"
	err = en.Append(0xa3, 0x70, 0x6c, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Pools)))
	if err != nil {
		err = msgp.WrapError(err, "Pools")
		return
	}
	for za0001 := range z.Pools {
		err = z.Pools[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Pools", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *poolMeta) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "v"
	o = append(o, 0x82, 0xa1, 0x76)
	o = msgp.AppendString(o, z.Version)
	// string "pls"
	o = append(o, 0xa3, 0x70, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Pools)))
	for za0001 := range z.Pools {
		o, err = z.Pools[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Pools", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *poolMeta) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "pls":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Pools")
				return
			}
			if cap(z.Pools) >= int(zb0002) {
				z.Pools = (z.Pools)[:zb0002]
			} else {
				z.Pools = make([]PoolInfo, zb0002)
			}
			for za0001 := range z.Pools {
				bts, err = z.Pools[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Pools", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *poolMeta) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Version) + 4 + msgp.ArrayHeaderSize
	for za0001 := range z.Pools {
		s += z.Pools[za0001].Msgsize()
	}
	return
}
