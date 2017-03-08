package encode_test

import (
	"testing"
	"fmt"
	. "github.com/firelyu/gotour/encode"
	"strconv"
)

type encodeBase64 struct {
	Raw	[]byte
	Encode []byte
}

var base64List []*encodeBase64

type validBase64 struct {
	Raw []byte
	Valid bool
}

var validBase64List []*validBase64

func init()  {
	raw := "any carnal pleasure."
	encode := "YW55IGNhcm5hbCBwbGVhc3VyZS4="
	base64 := &encodeBase64{Raw:[]byte(raw), Encode:[]byte(encode)}
	base64List = append(base64List, base64)

	raw = "any carnal pleasure"
	encode = "YW55IGNhcm5hbCBwbGVhc3VyZQ=="
	base64 = &encodeBase64{Raw:[]byte(raw), Encode:[]byte(encode)}
	base64List = append(base64List, base64)

	raw = "any carnal pleasur"
	encode = "YW55IGNhcm5hbCBwbGVhc3Vy"
	base64 = &encodeBase64{Raw:[]byte(raw), Encode:[]byte(encode)}
	base64List = append(base64List, base64)

	raw = "any carnal pleasu"
	encode = "YW55IGNhcm5hbCBwbGVhc3U="
	base64 = &encodeBase64{Raw:[]byte(raw), Encode:[]byte(encode)}
	base64List = append(base64List, base64)

	raw = "any carnal pleas"
	encode = "YW55IGNhcm5hbCBwbGVhcw=="
	base64 = &encodeBase64{Raw:[]byte(raw), Encode:[]byte(encode)}
	base64List = append(base64List, base64)

	raw = "Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure."
	encode = "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlzIHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0aGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdlLCBleGNlZWRzIHRoZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4="
	base64 = &encodeBase64{Raw:[]byte(raw), Encode:[]byte(encode)}
	base64List = append(base64List, base64)
	
	raw = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	valid := &validBase64{Raw:[]byte(raw), Valid:true}
	validBase64List = append(validBase64List, valid)

	raw = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	valid = &validBase64{Raw:[]byte(raw), Valid:false}
	validBase64List = append(validBase64List, valid)

	raw = "ABCD=FGH"
	valid = &validBase64{Raw:[]byte(raw), Valid:false}
	validBase64List = append(validBase64List, valid)

	raw = "ABC==FGH"
	valid = &validBase64{Raw:[]byte(raw), Valid:false}
	validBase64List = append(validBase64List, valid)

	raw = "AB===FGH"
	valid = &validBase64{Raw:[]byte(raw), Valid:false}
	validBase64List = append(validBase64List, valid)

	raw = "ABCDEF=H"
	valid = &validBase64{Raw:[]byte(raw), Valid:false}
	validBase64List = append(validBase64List, valid)

	raw = "----"
	valid = &validBase64{Raw:[]byte(raw), Valid:false}
	validBase64List = append(validBase64List, valid)

	raw = "...."
	valid = &validBase64{Raw:[]byte(raw), Valid:false}
	validBase64List = append(validBase64List, valid)
	
}

func dumpError(t *testing.T, base64 *encodeBase64) {
	t.Errorf("The raw is %s\n", string(base64.Raw))
	t.Errorf("The expected encode is %s\n", string(base64.Encode))
	t.Errorf("The error encode is %s\n", string(EncodeBase64(base64.Raw)))
	t.FailNow()
}

func TestEncodeBase64(t *testing.T)  {
	for _, base64 := range base64List {
		out := EncodeBase64(base64.Raw)

		if len(out) % 4 != 0 {
			t.Errorf("The expected length is round 4, the output length is %d\n",
				len(out))
			dumpError(t, base64)
		}

		if len(base64.Encode) != len(out) {
			t.Errorf("The expected length is %d, the output length is %d\n",
				len(base64.Raw), len(out))
			dumpError(t, base64)
		}

		for i := 0; i < len(out); i++ {
			if out[i] != base64.Encode[i] {
				t.Errorf("The %d char is not right.", i)
				t.Errorf("The expected char is %c, the output is %c",
					base64.Encode[i], out[i])
				dumpError(t, base64)
			}
		}
	}
}

func TestValidBase64(t *testing.T)  {
	for _, v := range validBase64List {
		res, _ := ValidBase64(v.Raw)
		if res != v.Valid {
			t.Errorf("The following string is expected %t, but the output is %t\n",
				v.Valid, res)
			t.Errorf("%s\n", v.Raw)
			t.Failed()
		}
	}
}

func TestDecodeBase64(t *testing.T)  {
	for _, base64 := range base64List {
		pass := true
		raw, _ := DecodeBase64(base64.Encode)

		if len(raw) != len(base64.Raw) {
			pass = false
			goto errout
		}

		for i := 0; i < len(raw); i++ {
			if raw[i] != base64.Raw[i] {
				t.Errorf("The %d char is not right.", i)
				t.Errorf("The expected char is %c, the output is %c",
					base64.Raw[i], raw[i])
				pass = false
			}
		}
errout:
		if pass != true {
			t.Errorf("The encode is %s\n", string(base64.Encode))
			t.Errorf("The expected raw is %s\n", string(base64.Raw))
			t.Errorf("The output raw is %s\n", string(raw))
			t.FailNow()
		}

		fmt.Printf("The encode is %s\n", string(base64.Encode))
		fmt.Printf("The expected raw is >%s<\n", string(base64.Raw))
		fmt.Printf("The output raw is   >%s<\n", string(raw))
		fmt.Println()
	}
}

func benchmarkBase(b *testing.B, raw []byte)  {
	for i := 0; i < b.N; i ++ {
		EncodeBase64(raw)
	}
}

func BenchmarkEncodeBase64 (b *testing.B) {
	for _, base64 := range base64List {
		b.Run("len(raw) is " + strconv.Itoa(len(base64.Raw)), func (b *testing.B){
			benchmarkBase(b, base64.Raw)
		})
	}
}