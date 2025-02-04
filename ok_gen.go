package main

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Flib) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "x":
			z.X, err = dc.ReadInt8()
			if err != nil {
				err = msgp.WrapError(err, "X")
				return
			}
		case "y":
			z.Y, err = dc.ReadInt8()
			if err != nil {
				err = msgp.WrapError(err, "Y")
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
func (z Flib) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "x"
	err = en.Append(0x82, 0xa1, 0x78)
	if err != nil {
		return
	}
	err = en.WriteInt8(z.X)
	if err != nil {
		err = msgp.WrapError(err, "X")
		return
	}
	// write "y"
	err = en.Append(0xa1, 0x79)
	if err != nil {
		return
	}
	err = en.WriteInt8(z.Y)
	if err != nil {
		err = msgp.WrapError(err, "Y")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Flib) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "x"
	o = append(o, 0x82, 0xa1, 0x78)
	o = msgp.AppendInt8(o, z.X)
	// string "y"
	o = append(o, 0xa1, 0x79)
	o = msgp.AppendInt8(o, z.Y)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Flib) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "x":
			z.X, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "X")
				return
			}
		case "y":
			z.Y, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Y")
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
func (z Flib) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int8Size + 2 + msgp.Int8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Flibbity) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "a":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "A")
				return
			}
			if cap(z.A) >= int(zb0002) {
				z.A = (z.A)[:zb0002]
			} else {
				z.A = make(Flibs, zb0002)
			}
			for za0001 := range z.A {
				var zb0003 uint32
				zb0003, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "A", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "A", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "x":
						z.A[za0001].X, err = dc.ReadInt8()
						if err != nil {
							err = msgp.WrapError(err, "A", za0001, "X")
							return
						}
					case "y":
						z.A[za0001].Y, err = dc.ReadInt8()
						if err != nil {
							err = msgp.WrapError(err, "A", za0001, "Y")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "A", za0001)
							return
						}
					}
				}
			}
		case "b":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "B")
				return
			}
			if cap(z.B) >= int(zb0004) {
				z.B = (z.B)[:zb0004]
			} else {
				z.B = make(Flibs, zb0004)
			}
			for za0002 := range z.B {
				var zb0005 uint32
				zb0005, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "B", za0002)
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "B", za0002)
						return
					}
					switch msgp.UnsafeString(field) {
					case "x":
						z.B[za0002].X, err = dc.ReadInt8()
						if err != nil {
							err = msgp.WrapError(err, "B", za0002, "X")
							return
						}
					case "y":
						z.B[za0002].Y, err = dc.ReadInt8()
						if err != nil {
							err = msgp.WrapError(err, "B", za0002, "Y")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "B", za0002)
							return
						}
					}
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
func (z *Flibbity) EncodeMsg(en *msgp.Writer) (err error) {
	// check for omitted fields
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 2 bits */
	_ = zb0001Mask
	if z.A == nil {
		zb0001Len--
		zb0001Mask |= 0x1
	}
	if z.B == nil {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}

	// skip if no fields are to be emitted
	if zb0001Len != 0 {
		if (zb0001Mask & 0x1) == 0 { // if not omitted
			// write "a"
			err = en.Append(0xa1, 0x61)
			if err != nil {
				return
			}
			err = en.WriteArrayHeader(uint32(len(z.A)))
			if err != nil {
				err = msgp.WrapError(err, "A")
				return
			}
			for za0001 := range z.A {
				// map header, size 2
				// write "x"
				err = en.Append(0x82, 0xa1, 0x78)
				if err != nil {
					return
				}
				err = en.WriteInt8(z.A[za0001].X)
				if err != nil {
					err = msgp.WrapError(err, "A", za0001, "X")
					return
				}
				// write "y"
				err = en.Append(0xa1, 0x79)
				if err != nil {
					return
				}
				err = en.WriteInt8(z.A[za0001].Y)
				if err != nil {
					err = msgp.WrapError(err, "A", za0001, "Y")
					return
				}
			}
		}
		if (zb0001Mask & 0x2) == 0 { // if not omitted
			// write "b"
			err = en.Append(0xa1, 0x62)
			if err != nil {
				return
			}
			err = en.WriteArrayHeader(uint32(len(z.B)))
			if err != nil {
				err = msgp.WrapError(err, "B")
				return
			}
			for za0002 := range z.B {
				// map header, size 2
				// write "x"
				err = en.Append(0x82, 0xa1, 0x78)
				if err != nil {
					return
				}
				err = en.WriteInt8(z.B[za0002].X)
				if err != nil {
					err = msgp.WrapError(err, "B", za0002, "X")
					return
				}
				// write "y"
				err = en.Append(0xa1, 0x79)
				if err != nil {
					return
				}
				err = en.WriteInt8(z.B[za0002].Y)
				if err != nil {
					err = msgp.WrapError(err, "B", za0002, "Y")
					return
				}
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Flibbity) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// check for omitted fields
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 2 bits */
	_ = zb0001Mask
	if z.A == nil {
		zb0001Len--
		zb0001Mask |= 0x1
	}
	if z.B == nil {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))

	// skip if no fields are to be emitted
	if zb0001Len != 0 {
		if (zb0001Mask & 0x1) == 0 { // if not omitted
			// string "a"
			o = append(o, 0xa1, 0x61)
			o = msgp.AppendArrayHeader(o, uint32(len(z.A)))
			for za0001 := range z.A {
				// map header, size 2
				// string "x"
				o = append(o, 0x82, 0xa1, 0x78)
				o = msgp.AppendInt8(o, z.A[za0001].X)
				// string "y"
				o = append(o, 0xa1, 0x79)
				o = msgp.AppendInt8(o, z.A[za0001].Y)
			}
		}
		if (zb0001Mask & 0x2) == 0 { // if not omitted
			// string "b"
			o = append(o, 0xa1, 0x62)
			o = msgp.AppendArrayHeader(o, uint32(len(z.B)))
			for za0002 := range z.B {
				// map header, size 2
				// string "x"
				o = append(o, 0x82, 0xa1, 0x78)
				o = msgp.AppendInt8(o, z.B[za0002].X)
				// string "y"
				o = append(o, 0xa1, 0x79)
				o = msgp.AppendInt8(o, z.B[za0002].Y)
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Flibbity) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "a":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "A")
				return
			}
			if cap(z.A) >= int(zb0002) {
				z.A = (z.A)[:zb0002]
			} else {
				z.A = make(Flibs, zb0002)
			}
			for za0001 := range z.A {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "A", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "A", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "x":
						z.A[za0001].X, bts, err = msgp.ReadInt8Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "A", za0001, "X")
							return
						}
					case "y":
						z.A[za0001].Y, bts, err = msgp.ReadInt8Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "A", za0001, "Y")
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "A", za0001)
							return
						}
					}
				}
			}
		case "b":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "B")
				return
			}
			if cap(z.B) >= int(zb0004) {
				z.B = (z.B)[:zb0004]
			} else {
				z.B = make(Flibs, zb0004)
			}
			for za0002 := range z.B {
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "B", za0002)
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "B", za0002)
						return
					}
					switch msgp.UnsafeString(field) {
					case "x":
						z.B[za0002].X, bts, err = msgp.ReadInt8Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "B", za0002, "X")
							return
						}
					case "y":
						z.B[za0002].Y, bts, err = msgp.ReadInt8Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "B", za0002, "Y")
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "B", za0002)
							return
						}
					}
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
func (z *Flibbity) Msgsize() (s int) {
	s = 1 + 2 + msgp.ArrayHeaderSize + (len(z.A) * (5 + msgp.Int8Size + msgp.Int8Size)) + 2 + msgp.ArrayHeaderSize + (len(z.B) * (5 + msgp.Int8Size + msgp.Int8Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Flibs) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Flibs, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, err = dc.ReadMapHeader()
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			switch msgp.UnsafeString(field) {
			case "x":
				(*z)[zb0001].X, err = dc.ReadInt8()
				if err != nil {
					err = msgp.WrapError(err, zb0001, "X")
					return
				}
			case "y":
				(*z)[zb0001].Y, err = dc.ReadInt8()
				if err != nil {
					err = msgp.WrapError(err, zb0001, "Y")
					return
				}
			default:
				err = dc.Skip()
				if err != nil {
					err = msgp.WrapError(err, zb0001)
					return
				}
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Flibs) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0004 := range z {
		// map header, size 2
		// write "x"
		err = en.Append(0x82, 0xa1, 0x78)
		if err != nil {
			return
		}
		err = en.WriteInt8(z[zb0004].X)
		if err != nil {
			err = msgp.WrapError(err, zb0004, "X")
			return
		}
		// write "y"
		err = en.Append(0xa1, 0x79)
		if err != nil {
			return
		}
		err = en.WriteInt8(z[zb0004].Y)
		if err != nil {
			err = msgp.WrapError(err, zb0004, "Y")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Flibs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0004 := range z {
		// map header, size 2
		// string "x"
		o = append(o, 0x82, 0xa1, 0x78)
		o = msgp.AppendInt8(o, z[zb0004].X)
		// string "y"
		o = append(o, 0xa1, 0x79)
		o = msgp.AppendInt8(o, z[zb0004].Y)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Flibs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Flibs, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			switch msgp.UnsafeString(field) {
			case "x":
				(*z)[zb0001].X, bts, err = msgp.ReadInt8Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, zb0001, "X")
					return
				}
			case "y":
				(*z)[zb0001].Y, bts, err = msgp.ReadInt8Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, zb0001, "Y")
					return
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					err = msgp.WrapError(err, zb0001)
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Flibs) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (len(z) * (5 + msgp.Int8Size + msgp.Int8Size))
	return
}
