package util

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/ffauzann/CAI-account/client"
	"github.com/ffauzann/CAI-account/internal/model"
	authCtx "github.com/ffauzann/common/auth/jwt/ctxval"
	"github.com/golang-jwt/jwt/v5"
)

// func ExtractClaimsFromString(ctx context.Context, strToken, signingKey string) (*model.Claims, bool) {
// 	hmacSecret := []byte(signingKey)

// 	token, err := jwt.ParseWithClaims(strToken, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return hmacSecret, nil
// 	})
// 	if err != nil {
// 		LogContext(ctx).Error(err.Error())
// 		return nil, false
// 	}

// 	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
// 		LogContext(ctx).Error(fmt.Sprintf("%v", token.Valid))
// 		return claims, true
// 	}

// 	LogContext(ctx).Error(fmt.Sprintf("%v", token.Claims))
// 	return nil, false
// }

func ExtractClaimsFromString(ctx context.Context, tokenString string, jwks []*model.Jwk) (claims *model.Claims, ok bool) {
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, &model.Claims{})
	if err != nil {
		LogContext(ctx).Error(err.Error())
		return
	}

	var publicKey *rsa.PublicKey
	for _, v := range jwks {
		if v.KeyID == token.Header["kid"] {
			nBytes, err := base64.RawURLEncoding.DecodeString(v.Modulus)
			if err != nil {
				LogContext(ctx).Error(err.Error())
				return
			}

			eBytes, err := base64.RawURLEncoding.DecodeString(v.Exponent)
			if err != nil {
				LogContext(ctx).Error(err.Error())
				return
			}

			publicKey = &rsa.PublicKey{
				N: new(big.Int).SetBytes(nBytes),
				E: int(new(big.Int).SetBytes(eBytes).Int64()),
			}
			break
		}
	}

	jwtToken, err := jwt.NewParser(
		jwt.WithIssuedAt(),
	).ParseWithClaims(tokenString, &model.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		LogContext(ctx).Error(err.Error())
		return
	}

	if jwtToken == nil {
		LogContext(ctx).Error("nil jwtToken")
		return
	}

	claims, ok = jwtToken.Claims.(*model.Claims)
	if ok && jwtToken.Valid {
		return claims, true
	}

	return nil, false
}

func ClaimsFromContext(ctx context.Context) (claims *client.Claims, ok bool) {
	iClaims, ok := authCtx.GetUserInfo(ctx)
	if !ok {
		return
	}

	claims, ok = iClaims.(*client.Claims)
	return
}
