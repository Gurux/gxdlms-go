package settings

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

func AesEncrypt(data []byte, secret []byte) ([]byte, error) {
	return ecbCrypt(data, secret, true)
}

func AesDecrypt(data []byte, secret []byte) ([]byte, error) {
	return ecbCrypt(data, secret, false)
}

func ecbCrypt(data []byte, secret []byte, encrypt bool) ([]byte, error) {
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("invalid AES-ECB input length %d", len(data))
	}
	if len(secret) == 0 {
		return nil, errors.New("secret is empty")
	}
	key := make([]byte, 16)
	copy(key, secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(data))
	for i := 0; i < len(data); i += aes.BlockSize {
		if encrypt {
			block.Encrypt(out[i:i+aes.BlockSize], data[i:i+aes.BlockSize])
		} else {
			block.Decrypt(out[i:i+aes.BlockSize], data[i:i+aes.BlockSize])
		}
	}
	return out, nil
}

func Secure(settings *GXDLMSSettings, cipher GXICipher, ic uint32, data []byte, secret []byte) ([]byte, error) {
	if settings == nil {
		return nil, errors.New("settings is nil")
	}
	switch settings.Authentication {
	case enums.AuthenticationHigh:
		return AesEncrypt(data, secret)
	case enums.AuthenticationHighMD5:
		h := md5.Sum(appendBytes(data, secret))
		return h[:], nil
	case enums.AuthenticationHighSHA1:
		h := sha1.Sum(appendBytes(data, secret))
		return h[:], nil
	case enums.AuthenticationHighSHA256:
		h := sha256.Sum256(appendBytes(secret, data))
		return h[:], nil
	case enums.AuthenticationHighGMAC:
		if cipher == nil {
			return nil, errors.New("cipher is nil")
		}
		p := NewAesGcmParameter(0, settings, enums.SecurityAuthentication, cipher.SecuritySuite(), uint64(ic), secret, cipher.BlockCipherKey(), cipher.AuthenticationKey())
		p.Type = CountTypeTag
		tag, err := EncryptAesGcm(p, data)
		if err != nil {
			return nil, err
		}
		out := types.NewGXByteBuffer()
		_ = out.SetUint8(byte(enums.SecurityAuthentication) | byte(cipher.SecuritySuite()))
		_ = out.SetUint32(ic)
		_ = out.Set(tag)
		return out.Array(), nil
	case enums.AuthenticationHighECDSA:
		if cipher == nil {
			return nil, errors.New("cipher is nil")
		}
		kp := cipher.SigningKeyPair()
		if kp.Value == nil {
			return nil, errors.New("signing key is not set")
		}
		sig, err := types.NewGXEcdsaFromPrivateKey(kp.Value)
		if err != nil {
			return nil, err
		}
		return sig.Sign(secret)
	default:
		return data, nil
	}
}

func appendBytes(a []byte, b []byte) []byte {
	out := make([]byte, 0, len(a)+len(b))
	out = append(out, a...)
	out = append(out, b...)
	return out
}

func GenerateChallenge(authentication enums.Authentication, size byte) []byte {
	randReader := rand.Reader
	length := int(size)
	if size == 0 || (size == 16 && authentication == enums.AuthenticationHighECDSA) {
		if authentication == enums.AuthenticationHighECDSA {
			length = 32 + pseudoRand(randReader, 32)
		} else {
			length = 8 + pseudoRand(randReader, 57)
		}
	}
	ret := make([]byte, length)
	for i := range ret {
		ret[i] = byte(pseudoRand(randReader, 0x7A))
	}
	return ret
}

func pseudoRand(r io.Reader, max int) int {
	if max <= 1 {
		return 0
	}
	b := make([]byte, 1)
	_, _ = io.ReadFull(r, b)
	return int(b[0]) % max
}

func GenerateKDF(securitySuite enums.SecuritySuite, z []byte, otherInfo []byte) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint32(1)
	if err != nil {
		return nil, err
	}
	err = bb.Set(z)
	if err != nil {
		return nil, err
	}
	err = bb.Set(otherInfo)
	if err != nil {
		return nil, err
	}
	switch securitySuite {
	case enums.SecuritySuite1:
		h := sha256.Sum256(bb.Array())
		return h[:], nil
	case enums.SecuritySuite2:
		h := sha512.Sum384(bb.Array())
		return h[:], nil
	default:
		return nil, fmt.Errorf("invalid security suite: %v", securitySuite)
	}
}

func GenerateKDFWithInfo(securitySuite enums.SecuritySuite, z []byte, algorithmID enums.AlgorithmID, partyUInfo []byte, partyVInfo []byte, suppPubInfo []byte, suppPrivInfo []byte) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.Set([]byte{0x60, 0x85, 0x74, 0x05, 0x08, 0x03, byte(algorithmID)})
	_ = bb.Set(partyUInfo)
	_ = bb.Set(partyVInfo)
	if len(suppPubInfo) != 0 {
		_ = bb.Set(suppPubInfo)
	}
	if len(suppPrivInfo) != 0 {
		_ = bb.Set(suppPrivInfo)
	}
	return GenerateKDF(securitySuite, z, bb.Array())
}

func securityControl(param *AesGcmParameter) byte {
	sc := byte(param.Security()) | byte(param.SecuritySuite)
	if param.Broacast {
		sc |= 0x40
	}
	if param.Compression {
		sc |= 0x80
	}
	return sc
}

func invocationCounterBytes(counter uint64) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(counter))
	return b
}

func nonceFrom(param *AesGcmParameter) ([]byte, error) {
	if len(param.SystemTitle()) != 8 {
		return nil, errors.New("invalid system title length")
	}
	nonce := make([]byte, 12)
	copy(nonce, param.SystemTitle())
	binary.BigEndian.PutUint32(nonce[8:], uint32(param.InvocationCounter))
	return nonce, nil
}

func buildAAD(sc byte, authKey []byte, data []byte, sec enums.Security) []byte {
	aad := make([]byte, 0, 1+len(authKey)+len(data))
	aad = append(aad, sc)
	aad = append(aad, authKey...)
	if sec == enums.SecurityAuthentication {
		aad = append(aad, data...)
	}
	return aad
}

func EncryptAesGcm(param *AesGcmParameter, plainText []byte) ([]byte, error) {
	if param == nil {
		return nil, errors.New("param is nil")
	}
	sc := securityControl(param)
	param.CountTag = nil

	block, err := aes.NewCipher(param.BlockCipherKey())
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCMWithTagSize(block, 12)
	if err != nil {
		return nil, err
	}
	nonce, err := nonceFrom(param)
	if err != nil {
		return nil, err
	}
	aad := buildAAD(sc, param.AuthenticationKey(), plainText, param.Security())

	out := types.NewGXByteBuffer()
	if param.Type == CountTypePacket {
		_ = out.SetUint8(sc)
		_ = out.Set(invocationCounterBytes(param.InvocationCounter))
	}

	switch param.Security() {
	case enums.SecurityAuthentication:
		tag := gcm.Seal(nil, nonce, nil, aad)
		param.CountTag = tag
		if param.Type == CountTypePacket || (param.Type&CountTypeData) != 0 {
			_ = out.Set(plainText)
		}
		if param.Type == CountTypePacket || (param.Type&CountTypeTag) != 0 {
			_ = out.Set(tag)
		}
	case enums.SecurityEncryption, enums.SecurityAuthenticationEncryption:
		full := gcm.Seal(nil, nonce, plainText, aad)
		ct := full[:len(full)-12]
		tag := full[len(full)-12:]
		param.CountTag = tag
		if param.Type == CountTypePacket {
			_ = out.Set(ct)
			_ = out.Set(tag)
		} else {
			if (param.Type & CountTypeData) != 0 {
				_ = out.Set(ct)
			}
			if (param.Type & CountTypeTag) != 0 {
				_ = out.Set(tag)
			}
		}
	case enums.SecurityNone:
		_ = out.Set(plainText)
	default:
		return nil, fmt.Errorf("invalid security: %v", param.Security())
	}
	return out.Array(), nil
}

func compareTag(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var v byte
	for i := range a {
		v |= a[i] ^ b[i]
	}
	return v == 0
}

func DecryptAesGcm(p *AesGcmParameter, data *types.GXByteBuffer) ([]byte, error) {
	if p == nil {
		return nil, errors.New("param is nil")
	}
	if data.Available() == 0 {
		return nil, errors.New("no data")
	}
	buf := make([]byte, data.Available())
	if err := data.Get(buf); err != nil {
		return nil, err
	}

	offset := 0
	sc := securityControl(p)
	if p.Type == CountTypePacket {
		if len(buf) < 5 {
			return nil, errors.New("invalid packet")
		}
		sc = buf[0]
		offset = 5 // security byte + invocation counter
	}

	payload := buf[offset:]
	if p.Security() == enums.SecurityNone {
		return payload, nil
	}
	if len(payload) < 12 {
		return nil, errors.New("invalid encrypted payload")
	}

	block, err := aes.NewCipher(p.BlockCipherKey())
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCMWithTagSize(block, 12)
	if err != nil {
		return nil, err
	}
	nonce, err := nonceFrom(p)
	if err != nil {
		return nil, err
	}

	switch p.Security() {
	case enums.SecurityAuthentication:
		plain := payload[:len(payload)-12]
		tag := payload[len(payload)-12:]
		aad := buildAAD(sc, p.AuthenticationKey(), plain, enums.SecurityAuthentication)
		expected := gcm.Seal(nil, nonce, nil, aad)
		if !compareTag(tag, expected) {
			return nil, errors.New("invalid authentication tag")
		}
		return plain, nil
	case enums.SecurityEncryption, enums.SecurityAuthenticationEncryption:
		aad := buildAAD(sc, p.AuthenticationKey(), nil, p.Security())
		plain, err := gcm.Open(nil, nonce, payload, aad)
		if err != nil {
			return nil, err
		}
		return plain, nil
	default:
		return nil, fmt.Errorf("invalid security: %v", p.Security())
	}
}

func GetEphemeralPublicKeyData(keyID int, ephemeralKey *types.GXPublicKey) ([]byte, error) {
	if ephemeralKey == nil {
		return nil, errors.New("ephemeral key is nil")
	}
	raw := make([]byte, len(ephemeralKey.RawValue()))
	copy(raw, ephemeralKey.RawValue())
	if len(raw) == 0 {
		return nil, errors.New("ephemeral key raw value is empty")
	}
	raw[0] = byte(keyID)
	return raw, nil
}

func GetEphemeralPublicKeySignature(keyID int, ephemeralKey *types.GXPublicKey, signKey *types.GXPrivateKey) ([]byte, error) {
	epk, err := GetEphemeralPublicKeyData(keyID, ephemeralKey)
	if err != nil {
		return nil, err
	}
	signer, err := types.NewGXEcdsaFromPrivateKey(signKey)
	if err != nil {
		return nil, err
	}
	return signer.Sign(epk)
}

func ValidateEphemeralPublicKeySignature(data []byte, sign []byte, publicSigningKey *types.GXPublicKey) (bool, error) {
	verifier, err := types.NewGXEcdsaFromPublicKey(publicSigningKey)
	if err != nil {
		return false, err
	}
	return verifier.Verify(sign, data)
}
