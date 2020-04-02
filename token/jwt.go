package token

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var J *Jwt

func init() {
	sk, err := jwt.ParseRSAPrivateKeyFromPEM(sk)
	if err != nil {
		log.Fatalf("can't initialize jwt token: %v", err)
	}

	pk, err := jwt.ParseRSAPublicKeyFromPEM(pk)
	if err != nil {
		log.Fatalf("can't initialize jwt token: %v", err)
	}

	J = &Jwt{sk: sk, pk: pk}
}

// 生成token
func Token(id int) (string, error) {
	return J.Token(id, 24*time.Hour)
}

// 校验token
func Auth(token string) (int, error) {
	return J.Auth(token)
}

type Jwt struct {
	sk *rsa.PrivateKey
	pk *rsa.PublicKey
}

func (t *Jwt) Token(id int, exp time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": time.Now().Add(exp).Unix(),
		"id":  id,
	}).SignedString(t.sk)
}

func (t *Jwt) Auth(token string) (int, error) {
	keyFunc := func(jt *jwt.Token) (interface{}, error) {
		if _, ok := jt.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("jwt token method err:" + token)
		}
		return t.pk, nil
	}

	jt, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return 0, err
	}

	claims := jt.Claims.(jwt.MapClaims)
	if _, ok := claims["id"]; !ok {
		return 0, errors.New("token is not expected")
	}
	return int(claims["id"].(float64)), nil
}

var pk = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJvHTtycWu4S57V2VR5xpZ/00N
ZYklRzViXElaly8bBStZoGrZfZP/XGYKL7MWXJXz7clpiKq0luYY7tVBYyEtGfSB
uZKQHVtvTlktF8aj9VY5vyatMajSh4S4MMXLXMo061FLIJsjNldzf8dt268pDfhn
63+yoL37R6KkoYwcsQIDAQAB
-----END PUBLIC KEY-----`)
var sk = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAMm8dO3Jxa7hLntX
ZVHnGln/TQ1liSVHNWJcSVqXLxsFK1mgatl9k/9cZgovsxZclfPtyWmIqrSW5hju
1UFjIS0Z9IG5kpAdW29OWS0XxqP1Vjm/Jq0xqNKHhLgwxctcyjTrUUsgmyM2V3N/
x23brykN+Gfrf7KgvftHoqShjByxAgMBAAECgYAK5AOj+qqTYFC0E+nCXCmOBjxQ
ptSakJePMdA79cFzQfovInviNedRd9mCX0TZxjHKM2IWMmyUxziwy+PzXbiisBUX
kFm08iVWZcKihnaFM6Ovq5+unqhMTChkjfOT8xmXV2DwhZ1z4laSrpesV5FY5c79
Y6iXWSgssUNdd4Y8AQJBAPPHYMGrS2Jo60qQqyvbUyi/G+8UJUeE02ss626JzQKj
mh4cZ3Dg52RZ+WsToAGpUUfHZBA92smH6ibYwu42B1ECQQDT2YJmzrfJ+iVvz27i
b+X9AWMJSA7DlIPs+Ku9KOX9tiEr1F87hgcoySk+7uy7E46cFnIU/+zXrgvQd8aT
FidhAkBt2b5EB0hlBBpi82Xu0Vpb1iJkTZOu8q/Cb+93VClJsydTwkDqoK4kjlbS
ZXmIxh+WVMdGelkIz3I4Jx1P8pDBAkEA0K21EegWilJphhXiuBJZjjtLft0IDgfB
XDAnm5Ep3B0H19C+bje73aUph+B6OF0vYPmLLrxaZKoA4Tza0hBEgQJBAM0Hu63C
DyZOc+dM6o4iJionIXUIp3GKn+9eRZUX2qUR9/oOMoIbWyHRBDL3L5eh8EofFXf0
8EaU0VRpyuxfPsk=
-----END RSA PRIVATE KEY-----`)
