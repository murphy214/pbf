package pbf

import (
	"encoding/binary"
	"math/rand"
	"testing"
)


func create_random_array(n int) [][]byte {
	myarr := make([][]byte,n)
	for i := 0;i < n; i++ {
		myarr[i] = EncodeVarint(uint64(rand.Int()))
	}  
	return myarr
}





var uint32e = []byte{0xa9, 0x1, 0x0, 0xfb, 0xfe, 0x6, 0x1, 0x1, 0x2, 0x1f, 0x3, 0xfc, 0xfe, 0x6, 0x4, 0xfd, 0xfe, 0x6, 0x5, 0xfe, 0xfe, 0x6, 0x6, 0xff, 0xfe, 0x6, 0x1b, 0xeb, 0xba, 0x4, 0x29, 0x80, 0xff, 0x6, 0xcc, 0x1, 0x81, 0xff, 0x6, 0x7, 0x82, 0xff, 0x6, 0x8, 0xe1, 0xf6, 0x6, 0xcd, 0x1, 0xe9, 0x2, 0xa1, 0x1, 0xf6, 0x4, 0xc9, 0x1, 0x83, 0xff, 0x6, 0x47, 0xa1, 0xc, 0xce, 0x1, 0x6b, 0x11, 0x84, 0xff, 0x6, 0xcf, 0x1, 0x85, 0xff, 0x6, 0x49, 0xe6, 0x6, 0xd0, 0x1, 0x86, 0xff, 0x6, 0x4b, 0x87, 0xff, 0x6, 0x61, 0x88, 0xff, 0x6, 0xd1, 0x1, 0x89, 0xff, 0x6, 0x51, 0xa0, 0xc, 0x62, 0xdb, 0x18, 0xd2, 0x1, 0x8a, 0xff, 0x6, 0x63, 0xdc, 0x18, 0xd3, 0x1, 0x8b, 0xff, 0x6, 0xd4, 0x1, 0x8c, 0xff, 0x6, 0x52, 0x8d, 0xff, 0x6, 0x4c, 0x8e, 0xff, 0x6, 0xd5, 0x1, 0x8f, 0xff, 0x6, 0xd6, 0x1, 0xb1, 0xc1, 0x3, 0xd7, 0x1, 0xe6, 0x6, 0xd8, 0x1, 0x90, 0xff, 0x6, 0xd9, 0x1, 0x91, 0xff, 0x6, 0xda, 0x1, 0x92, 0xff, 0x6, 0xdb, 0x1, 0x93, 0xff, 0x6, 0xdc, 0x1, 0xe4, 0x3, 0xdd, 0x1, 0xde, 0xcc, 0x2,0xa9, 0x1, 0x0, 0xfb, 0xfe, 0x6, 0x1, 0x1, 0x2, 0x1f, 0x3, 0xfc, 0xfe, 0x6, 0x4, 0xfd, 0xfe, 0x6, 0x5, 0xfe, 0xfe, 0x6, 0x6, 0xff, 0xfe, 0x6, 0x1b, 0xeb, 0xba, 0x4, 0x29, 0x80, 0xff, 0x6, 0xcc, 0x1, 0x81, 0xff, 0x6, 0x7, 0x82, 0xff, 0x6, 0x8, 0xe1, 0xf6, 0x6, 0xcd, 0x1, 0xe9, 0x2, 0xa1, 0x1, 0xf6, 0x4, 0xc9, 0x1, 0x83, 0xff, 0x6, 0x47, 0xa1, 0xc, 0xce, 0x1, 0x6b, 0x11, 0x84, 0xff, 0x6, 0xcf, 0x1, 0x85, 0xff, 0x6, 0x49, 0xe6, 0x6, 0xd0, 0x1, 0x86, 0xff, 0x6, 0x4b, 0x87, 0xff, 0x6, 0x61, 0x88, 0xff, 0x6, 0xd1, 0x1, 0x89, 0xff, 0x6, 0x51, 0xa0, 0xc, 0x62, 0xdb, 0x18, 0xd2, 0x1, 0x8a, 0xff, 0x6, 0x63, 0xdc, 0x18, 0xd3, 0x1, 0x8b, 0xff, 0x6, 0xd4, 0x1, 0x8c, 0xff, 0x6, 0x52, 0x8d, 0xff, 0x6, 0x4c, 0x8e, 0xff, 0x6, 0xd5, 0x1, 0x8f, 0xff, 0x6, 0xd6, 0x1, 0xb1, 0xc1, 0x3, 0xd7, 0x1, 0xe6, 0x6, 0xd8, 0x1, 0x90, 0xff, 0x6, 0xd9, 0x1, 0x91, 0xff, 0x6, 0xda, 0x1, 0x92, 0xff, 0x6, 0xdb, 0x1, 0x93, 0xff, 0x6, 0xdc, 0x1, 0xe4, 0x3, 0xdd, 0x1, 0xde, 0xcc, 0x2,0xa9, 0x1, 0x0, 0xfb, 0xfe, 0x6, 0x1, 0x1, 0x2, 0x1f, 0x3, 0xfc, 0xfe, 0x6, 0x4, 0xfd, 0xfe, 0x6, 0x5, 0xfe, 0xfe, 0x6, 0x6, 0xff, 0xfe, 0x6, 0x1b, 0xeb, 0xba, 0x4, 0x29, 0x80, 0xff, 0x6, 0xcc, 0x1, 0x81, 0xff, 0x6, 0x7, 0x82, 0xff, 0x6, 0x8, 0xe1, 0xf6, 0x6, 0xcd, 0x1, 0xe9, 0x2, 0xa1, 0x1, 0xf6, 0x4, 0xc9, 0x1, 0x83, 0xff, 0x6, 0x47, 0xa1, 0xc, 0xce, 0x1, 0x6b, 0x11, 0x84, 0xff, 0x6, 0xcf, 0x1, 0x85, 0xff, 0x6, 0x49, 0xe6, 0x6, 0xd0, 0x1, 0x86, 0xff, 0x6, 0x4b, 0x87, 0xff, 0x6, 0x61, 0x88, 0xff, 0x6, 0xd1, 0x1, 0x89, 0xff, 0x6, 0x51, 0xa0, 0xc, 0x62, 0xdb, 0x18, 0xd2, 0x1, 0x8a, 0xff, 0x6, 0x63, 0xdc, 0x18, 0xd3, 0x1, 0x8b, 0xff, 0x6, 0xd4, 0x1, 0x8c, 0xff, 0x6, 0x52, 0x8d, 0xff, 0x6, 0x4c, 0x8e, 0xff, 0x6, 0xd5, 0x1, 0x8f, 0xff, 0x6, 0xd6, 0x1, 0xb1, 0xc1, 0x3, 0xd7, 0x1, 0xe6, 0x6, 0xd8, 0x1, 0x90, 0xff, 0x6, 0xd9, 0x1, 0x91, 0xff, 0x6, 0xda, 0x1, 0x92, 0xff, 0x6, 0xdb, 0x1, 0x93, 0xff, 0x6, 0xdc, 0x1, 0xe4, 0x3, 0xdd, 0x1, 0xde, 0xcc, 0x2,0xa9, 0x1, 0x0, 0xfb, 0xfe, 0x6, 0x1, 0x1, 0x2, 0x1f, 0x3, 0xfc, 0xfe, 0x6, 0x4, 0xfd, 0xfe, 0x6, 0x5, 0xfe, 0xfe, 0x6, 0x6, 0xff, 0xfe, 0x6, 0x1b, 0xeb, 0xba, 0x4, 0x29, 0x80, 0xff, 0x6, 0xcc, 0x1, 0x81, 0xff, 0x6, 0x7, 0x82, 0xff, 0x6, 0x8, 0xe1, 0xf6, 0x6, 0xcd, 0x1, 0xe9, 0x2, 0xa1, 0x1, 0xf6, 0x4, 0xc9, 0x1, 0x83, 0xff, 0x6, 0x47, 0xa1, 0xc, 0xce, 0x1, 0x6b, 0x11, 0x84, 0xff, 0x6, 0xcf, 0x1, 0x85, 0xff, 0x6, 0x49, 0xe6, 0x6, 0xd0, 0x1, 0x86, 0xff, 0x6, 0x4b, 0x87, 0xff, 0x6, 0x61, 0x88, 0xff, 0x6, 0xd1, 0x1, 0x89, 0xff, 0x6, 0x51, 0xa0, 0xc, 0x62, 0xdb, 0x18, 0xd2, 0x1, 0x8a, 0xff, 0x6, 0x63, 0xdc, 0x18, 0xd3, 0x1, 0x8b, 0xff, 0x6, 0xd4, 0x1, 0x8c, 0xff, 0x6, 0x52, 0x8d, 0xff, 0x6, 0x4c, 0x8e, 0xff, 0x6, 0xd5, 0x1, 0x8f, 0xff, 0x6, 0xd6, 0x1, 0xb1, 0xc1, 0x3, 0xd7, 0x1, 0xe6, 0x6, 0xd8, 0x1, 0x90, 0xff, 0x6, 0xd9, 0x1, 0x91, 0xff, 0x6, 0xda, 0x1, 0x92, 0xff, 0x6, 0xdb, 0x1, 0x93, 0xff, 0x6, 0xdc, 0x1, 0xe4, 0x3, 0xdd, 0x1, 0xde, 0xcc, 0x2,0xa9, 0x1, 0x0, 0xfb, 0xfe, 0x6, 0x1, 0x1, 0x2, 0x1f, 0x3, 0xfc, 0xfe, 0x6, 0x4, 0xfd, 0xfe, 0x6, 0x5, 0xfe, 0xfe, 0x6, 0x6, 0xff, 0xfe, 0x6, 0x1b, 0xeb, 0xba, 0x4, 0x29, 0x80, 0xff, 0x6, 0xcc, 0x1, 0x81, 0xff, 0x6, 0x7, 0x82, 0xff, 0x6, 0x8, 0xe1, 0xf6, 0x6, 0xcd, 0x1, 0xe9, 0x2, 0xa1, 0x1, 0xf6, 0x4, 0xc9, 0x1, 0x83, 0xff, 0x6, 0x47, 0xa1, 0xc, 0xce, 0x1, 0x6b, 0x11, 0x84, 0xff, 0x6, 0xcf, 0x1, 0x85, 0xff, 0x6, 0x49, 0xe6, 0x6, 0xd0, 0x1, 0x86, 0xff, 0x6, 0x4b, 0x87, 0xff, 0x6, 0x61, 0x88, 0xff, 0x6, 0xd1, 0x1, 0x89, 0xff, 0x6, 0x51, 0xa0, 0xc, 0x62, 0xdb, 0x18, 0xd2, 0x1, 0x8a, 0xff, 0x6, 0x63, 0xdc, 0x18, 0xd3, 0x1, 0x8b, 0xff, 0x6, 0xd4, 0x1, 0x8c, 0xff, 0x6, 0x52, 0x8d, 0xff, 0x6, 0x4c, 0x8e, 0xff, 0x6, 0xd5, 0x1, 0x8f, 0xff, 0x6, 0xd6, 0x1, 0xb1, 0xc1, 0x3, 0xd7, 0x1, 0xe6, 0x6, 0xd8, 0x1, 0x90, 0xff, 0x6, 0xd9, 0x1, 0x91, 0xff, 0x6, 0xda, 0x1, 0x92, 0xff, 0x6, 0xdb, 0x1, 0x93, 0xff, 0x6, 0xdc, 0x1, 0xe4, 0x3, 0xdd, 0x1, 0xde, 0xcc, 0x2,0xa9, 0x1, 0x0, 0xfb, 0xfe, 0x6, 0x1, 0x1, 0x2, 0x1f, 0x3, 0xfc, 0xfe, 0x6, 0x4, 0xfd, 0xfe, 0x6, 0x5, 0xfe, 0xfe, 0x6, 0x6, 0xff, 0xfe, 0x6, 0x1b, 0xeb, 0xba, 0x4, 0x29, 0x80, 0xff, 0x6, 0xcc, 0x1, 0x81, 0xff, 0x6, 0x7, 0x82, 0xff, 0x6, 0x8, 0xe1, 0xf6, 0x6, 0xcd, 0x1, 0xe9, 0x2, 0xa1, 0x1, 0xf6, 0x4, 0xc9, 0x1, 0x83, 0xff, 0x6, 0x47, 0xa1, 0xc, 0xce, 0x1, 0x6b, 0x11, 0x84, 0xff, 0x6, 0xcf, 0x1, 0x85, 0xff, 0x6, 0x49, 0xe6, 0x6, 0xd0, 0x1, 0x86, 0xff, 0x6, 0x4b, 0x87, 0xff, 0x6, 0x61, 0x88, 0xff, 0x6, 0xd1, 0x1, 0x89, 0xff, 0x6, 0x51, 0xa0, 0xc, 0x62, 0xdb, 0x18, 0xd2, 0x1, 0x8a, 0xff, 0x6, 0x63, 0xdc, 0x18, 0xd3, 0x1, 0x8b, 0xff, 0x6, 0xd4, 0x1, 0x8c, 0xff, 0x6, 0x52, 0x8d, 0xff, 0x6, 0x4c, 0x8e, 0xff, 0x6, 0xd5, 0x1, 0x8f, 0xff, 0x6, 0xd6, 0x1, 0xb1, 0xc1, 0x3, 0xd7, 0x1, 0xe6, 0x6, 0xd8, 0x1, 0x90, 0xff, 0x6, 0xd9, 0x1, 0x91, 0xff, 0x6, 0xda, 0x1, 0x92, 0xff, 0x6, 0xdb, 0x1, 0x93, 0xff, 0x6, 0xdc, 0x1, 0xe4, 0x3, 0xdd, 0x1, 0xde, 0xcc, 0x2}
var pbfval = PBF{Pbf:uint32e,Length:len(uint32e)}
var int_arr = create_random_array(1000)

func TestDecodeVarint(t *testing.T) {
	expected := uint64(3234230423)

	b := make([]byte, 8)
	n := binary.PutUvarint(b, uint64(expected))
	b = b[:n]
	val := DecodeVarint(b)

	if val != expected {
		t.Errorf("DecodeVarint %d expected got %d",expected,val)
	}
}

func TestKey(t *testing.T) {
	e_k,e_t := byte(3),byte(2)
	kk,tt := Key(26) 

	if kk != e_k || tt != e_t {
		t.Errorf("TestKey %b %b expected got %b %b",e_k,e_t,kk,tt)		
	}
}


func TestReadInt32(t *testing.T) {
	expected := int32(33234)


	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, uint32(expected))
	//b = b[:n]
	val := int32(ReadInt32(b[:2]))


	if val != expected {
		t.Errorf("DecodeVarint %d expected got %d",expected,val)
	}
}



func TestReadUInt32(t *testing.T) {
	expected := uint32(33234)


	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, uint32(expected))
	//b = b[:n]
	val := uint32(ReadUInt32(b[:2]))


	if val != expected {
		t.Errorf("DecodeVarint %d expected got %d",expected,val)
	}
}


func TestReadInt32_Pbf(t *testing.T) {
	expected := int32(33234)


	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, uint32(expected))
	//b = b[:n]
	b = append(b,[]byte{2,4}...)
	pbfval := &PBF{Pbf:b,Length:len(b)}

	val := pbfval.ReadInt32()


	if val != expected {
		t.Errorf("DecodeVarint %d expected got %d",expected,val)
	}
}


func TestReadUInt32_Pbf(t *testing.T) {
	expected := uint32(33234)


	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, uint32(expected))
	//b = b[:n]
	b = append(b,[]byte{2,4}...)
	pbfval := &PBF{Pbf:b,Length:len(b)}

	val := pbfval.ReadUInt32()


	if val != expected {
		t.Errorf("DecodeVarint %d expected got %d",expected,val)
	}
}

// benchamrks every new vector tile
func Benchmark_ReadPackedUInt32_New(b *testing.B) {
        b.ReportAllocs()

        // run the Fib function b.N times
        for n := 0; n < b.N; n++ {
            pbfval.ReadPackedUInt32_2()
            pbfval.Pos = 0

        }
}


// benchamrks every new vector tile
func Benchmark_ReadPackedUInt32_Newer(b *testing.B) {
        b.ReportAllocs()

        // run the Fib function b.N times
        for n := 0; n < b.N; n++ {
            pbfval.ReadPackedUInt32()
            pbfval.Pos = 0

        }
}



// benchamrks every new vector tile
func Benchmark_ReadPackedUInt32_Newer2(b *testing.B) {
        b.ReportAllocs()

        // run the Fib function b.N times
        for n := 0; n < b.N; n++ {
            pbfval.ReadPackedUInt32_3()
            pbfval.Pos = 0

        }
}





var ii uint64
func BenchmarkDecodeVarint(b *testing.B) {
	b.ReportAllocs()

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {

		for i := 0; i < 1000; i++ {
			ii = DecodeVarint(int_arr[i])
		}
	}
}


func BenchmarkDecodeVarint2(b *testing.B) {
	b.ReportAllocs()

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {

		for i := 0; i < 1000; i++ {
			ii,_ = DecodeVarint2(int_arr[i])
		}
	}
}



