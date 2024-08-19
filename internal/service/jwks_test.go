package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ffauzann/CAI-account/client"
	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/common/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestJwks(t *testing.T) {
	var (
		ctx    = context.Background()
		logger = logger.Setup(logger.EnvTesting)
		config = &model.AppConfig{
			Jwt: model.JwtConfig{
				AsymmetricKeys: model.JwtAsymmetricKeysConfig{
					&struct {
						Kid        string
						PrivateKey string
						PublicKey  string
					}{
						Kid: "8421bf44-1178-4414-8909-9e99a4dfb770",
						// PrivateKey: "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdDJWTkNiT3I4aTR6R29zT0xLSkVhbCtkcnc1SHFlTHppVEkwY0ZvVUR2Y2pQYVUwCmdWSnM5amNuOTgvRlpadHNpUDIrOFRxSlFaNUFrLzJCcnJXa0tiVlEvMTQxaGhVdCtSSWszR1J1T0V6TWxOM1MKcmVPWHVpSEsyU2pLU3RRSlRyYXdLdGJkRGdwRGplRVNTVmcvdC9oUndEaXAxN2taenVEc2xDc3E5d1pTR1h6bgp1aDlJUDI4aCs3Q051a0FqVnFScWszcjlnOThTTkx2ZDdmd2xYcmVFOFFGS0hlYnFyME9rMzcrbWNPVUhFM3JmClpPQXg0UDkyVGR1SHIzV21wZkxvRlF0S2Q1VVZqOWdYNW8yTmFaak9GSXUvbnY5K1NZWnVhb1VNRVRKL0h2VjUKcEFabDVvZDUzZnZ6Y25KQmNaME9iQW5mdG1UZU8wNGs5UEl0RVFJREFRQUJBb0lCQVFDc3VGRVhwQWw2YXB4aQprVGZtUFdTbHNpdDFwTU5GY3FMZVFWUTF4QUJFSCtrbXM2S0JjVG1Cb1d5WTdTc0JpS0Z0VzEwckgzQUpScHVYClJSZVBqUzV3d1h6cEpMYlA4cjU3WnVVa1U4bWlhR0g4aWZWVEk1ZlFDdWRhSWhweTRzTnBTSkVkcDRKRktORjYKbTlCM0Z3L2JtWmlVcWtqN0RDOE1NYlZkemxJR2xtaFpjckhHZyttb2NNc3hCcjY4TVgzMTZ5NjBoMWFDdTYwUwpvMHYyMzR3K0VBN1hxTUluUitoZUxRRVc3Y0c0ZUdZN2NNTnErY1NPWVRmczN1bGpVeHZLeUtocHJuN05wd1dhClJnVzhrcnVGSmJaRGJIUTJQZFhqakVXYURaZkVIU0h1bUovWmc2QTZZSlpmaGlFdDBkOWZxR2RPS3VUSDMxU3cKK0k5YnBKWkpBb0dCQU1DQnRiOWNDbExIbHllUDJhWjZnbTNaUzlobTNuM2lKOWcwS0R2U0c4NSt3OTlwdGVSTwpTLzdsT1UwRGN2KzVFMGN1UlpUdE9TR3NUY3JjWjhaM1RCZFVyeE5ZTWNxamgybEtjT0dWRlFaT3B6ZDhVN2U0CmdyaHdTb3hWdFovaHQvVnRpMi9sS3JZOVhueUZZQkRjZGdsbERaZjRIeHVmQ01hbmxhNUYzWGhiQW9HQkFQUGkKVG5raUtJQjBLS0JSb28rc0N3TUhJa0JFb0R4RC9yL0ZNcHdBdkh1TXZhSGxkc1krUVdBTUErUjloK2JPYUNYawoveVRvVkpvaGFKSVRiaU5HMDRpS2FMMXROaHNJTXJqL3U1dXM4TkVoOU5XSzZRbGZ2MDcvSklDWUUraTA4aGJGCmlVU1RMNCtFUW1sWGRrdmVtUHBSVkRzRjhBdlBFd3F0T1NuSEdJd0RBb0dBQVZHaUxpSnlTNmprWnpmOEZNRG8KSGRxTVEzcEk4ZkhYdGdwOWNCTjdiMG05QzgzTW1qalRHbmIxa29xQWdqSUJhTTV2V1pyYWRsbVkydGZ4dWhGZApLeGZBYjFCK1h0WUorbld4R2txTUwxUGduMmV4cHlPVGViSURRTHpobHF2VU45RTlVRkh3bmZrRHFiUzhPTUZaCjZheVFrRWI1NTVXS1dOb1RFM09WRmRzQ2dZQnJsRjlEUmRNUjNxdHhGTEdkcUtsdTIzMjdWY3BNNnoxN2dGUXoKeG90ZUFKWkJ6UU9ZclN1UFg1MXo4LyszeTBMYnZHamo4ZXduMVNiWWtPT2JnZ21iaUZwdGZMaEtNbEtWa3BGQwpPWVk4NmpxaTI5U3lBdDlUekc1Z256VGhDTGhsWFJ1UStWQVlnYUg5Nzh2SjZkWVhUVHJYa21YeC81VUp0NkdvCmtSOTkyd0tCZ0NUcGRJMU9DcnhBVEZ0YklqQXYvSHpFcjFmRGJ5ZnEzQkxWak5BYkV0bWRVamp0L2ZrU1VSSmcKQzZyRTJXTXUvQUZlQVNGSGVqb0tFQ2NnNWFJZ1U0ZTFUQm9paTVMTjNTdTJyTmFIeStxYmVhMWlaVU4zNTkybwp4amJnY1JBek16Q3JTYnkrTzhEUXFSdFJYM0xFczR5cjFlaUZyTnpNWkR0MXlvY05XeGZiCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t",
						PublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF0MlZOQ2JPcjhpNHpHb3NPTEtKRQphbCtkcnc1SHFlTHppVEkwY0ZvVUR2Y2pQYVUwZ1ZKczlqY245OC9GWlp0c2lQMis4VHFKUVo1QWsvMkJycldrCktiVlEvMTQxaGhVdCtSSWszR1J1T0V6TWxOM1NyZU9YdWlISzJTaktTdFFKVHJhd0t0YmREZ3BEamVFU1NWZy8KdC9oUndEaXAxN2taenVEc2xDc3E5d1pTR1h6bnVoOUlQMjhoKzdDTnVrQWpWcVJxazNyOWc5OFNOTHZkN2Z3bApYcmVFOFFGS0hlYnFyME9rMzcrbWNPVUhFM3JmWk9BeDRQOTJUZHVIcjNXbXBmTG9GUXRLZDVVVmo5Z1g1bzJOCmFaak9GSXUvbnY5K1NZWnVhb1VNRVRKL0h2VjVwQVpsNW9kNTNmdnpjbkpCY1owT2JBbmZ0bVRlTzA0azlQSXQKRVFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t",
					}, {
						Kid: "978b94f3-c2c1-414b-af9e-504bc45e1604",
						// PrivateKey: "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBenE1ZTY0eTZXc2xOYWFBLzdjQkxNZTNBY0VrNWVsRU9PdDBQVGdBWGdOU1hWdnFICjE1NEVqZlMrYUc1TnZYVzIvQXFxUzBZUmY4ZzNQSllLaHF0aU9XQzQyVjA0cjJBaXZqVWpqWjdWMDY1dXh2eWsKZm5SQ1hXcE5EdS9BL2xHcUg3bzFlNm1RZEd5Vy9UNXM1amR5SFJuY2I2Q0d0TnF3VFluZzA5TzUrbUJCUzdLdgpJRHFnOHNMZHFVTEh6SkRYSDlXaEYrclp5UFI5R0xnb2xFZys5U3liYVhiV3VWWkZYOUpjUTRnRkM3OVlUUlRmCnprdlorMTR5aWJxdGxVRnJVVUVQMURmVCtUV0dWVFU2UDl3UHIxWENlUENVOS9PR1dFUGlOMk16YU5NeWRPVXMKNkJ4ZGNkaklMSFpoUzUxZEhUaU9XcXk1YUY3UFd0aTdkZUFnWHdJREFRQUJBb0lCQUhYVWZYTUcyUnQzRm5ZNAprUm5IZmxjcHQ0T01pNE5MZ0xSWVlTaFQ3eEpZb1N0S1MzWEd0Y3dFa3lWUWRXdWxGN3hiakRpNzZyQVNBa085Ck9xVUtRa1o1K1FpYkYvMEw3dUxId3N3em1LNUZEUXpPN2l6VnRSd3l4Vm5Wb0E2ZG1rTGFVekY4TzBuVXVzUUgKK2VmS0Jubkd5NkNzUVFBTWlXUzdUWDBXZ1RuVzdqdWlhOEhWN1gxcGtMc3lPYjduaXNGdEkvWjNJN3pqYnlQRwpUeEhwTDZqdk1DYUVxSXFZeWZndmNYKytoKzNhR1pkL2taNzNSOHd4NWY3UTlxVEJVRTBBVFk3b2trdEt5Tm1mClZzVzAwTnhHUjg5bWNpeFNySjBvOVpsMjhwR3o0blRMdGtDdXdIcDBvWkREc0g4Vld4REtlMHdOeFY1SFVVZzgKeDNOOGJja0NnWUVBNWJFaTBSRFFZNkh0N3FzbStSV0l6UnlqaGJkL0lEZy9IdVVCanhoZVhSK01CdUhMaGpBVgpjNFRDamptTHE1TDhrNVlNbjFkOWpQTzlRcFQ2MXJRbGNpNkxrU1JHOE1EQWpFZTk2VWNwY00vaWN4MkdsaW1DCkg5Q3p0Ty95Qnh0bm9EL1Y5K2l2TXJybHIxNHVFSUYzbE9OYUdzOXhqWDlRU25BVHJoT0ZqOTBDZ1lFQTVscUgKUkE5SHRpbm1hVWlKOW03eFZielgzRTEzNHA3VjRzMHViOUwvTnVYQjJEaExLRTdlSER6b2MyckxEY1BDWWJYdgpVeE9VaXU1L3E4aWFxbTh0OVdEeUZtT25QTDdxQkRsK2FYV0oyeUJGL0VUTGVqVURrWGFpdEp4a2ZZWmNqdGFpCmp5OFgrb1Mwa1hLaDVsWWJ0WWlhMjlZMWRFNFowaGMrWW1Dd2kyc0NnWUJNQ3YzczJ6VXlsd3lQcElndGxLeUsKdzMxN3FvbGk0Rnc5WFRITDd4Um1uaWdjcXlwWFRabjhlYXB6cmFlSThRdS96TUIzREY4YmlDSlRaY0U1emNCTAo4Zzd3eVdMWEYrbG5SK1VlMHhsc0tOYmVwNXJFSWcvYmVwdlVQbEFSZkVndGJKVHBFMWJWWTd6Zzl6d202TVh2Ck8rbTcwSXZXZlp6V1dBNmI1Z2lrM1FLQmdRQ2dlbWNMOGowNldqeHNFcDRTc2IydHhtNzN5bngvdzdvc1ZGZEsKamt0QW1wWU9KVUd1NWlyMHJ6WWQya04xNW5CUUZDaVR6dXRJVXRsYyt0ZE5iMEVzVm9xeE1HRzgveUFYSzFxMgo2NnpMbk5ORTZVUjhjZHgzaFpmYWphMndpVitJSDEyR3RQOVdod2NvNFJ5TFZHZmZjemkzZUtjcXpKSk9lK1Q1Ck1uVXIyd0tCZ0N5d3VvNlRreFZ4S3k4NHZ4OFN1RXVHL0xWWDI1NllaUGpuc0VZMTJmZDF0NXpaNU5hQlZSaGQKcitDdkk4UmZyNzdQZ1JMd2VFQzVKMlBiUUF1dTVOc1pzS3E0ZWoxNFNkZ1JoYnUzcDNmZ05HQTF6THljcDB1aQpHdWt4dnp6Q1Q0TmFZQ0szWGtmcndZU0hlNnUrbDF2RWZzV3Q0TWxsckh6RDJ0NzF4bGllCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t",
						PublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF6cTVlNjR5NldzbE5hYUEvN2NCTApNZTNBY0VrNWVsRU9PdDBQVGdBWGdOU1hWdnFIMTU0RWpmUythRzVOdlhXMi9BcXFTMFlSZjhnM1BKWUtocXRpCk9XQzQyVjA0cjJBaXZqVWpqWjdWMDY1dXh2eWtmblJDWFdwTkR1L0EvbEdxSDdvMWU2bVFkR3lXL1Q1czVqZHkKSFJuY2I2Q0d0TnF3VFluZzA5TzUrbUJCUzdLdklEcWc4c0xkcVVMSHpKRFhIOVdoRityWnlQUjlHTGdvbEVnKwo5U3liYVhiV3VWWkZYOUpjUTRnRkM3OVlUUlRmemt2WisxNHlpYnF0bFVGclVVRVAxRGZUK1RXR1ZUVTZQOXdQCnIxWENlUENVOS9PR1dFUGlOMk16YU5NeWRPVXM2QnhkY2RqSUxIWmhTNTFkSFRpT1dxeTVhRjdQV3RpN2RlQWcKWHdJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t",
					},
				},
			},
		}
		service      = New(nil, nil, nil, config, logger)
		expectedJwks = []*model.Jwk{
			{
				KeyType:   "RSA",
				KeyID:     "8421bf44-1178-4414-8909-9e99a4dfb770",
				Usage:     "sig",
				Algorithm: "RS256",
				Modulus:   "t2VNCbOr8i4zGosOLKJEal-drw5HqeLziTI0cFoUDvcjPaU0gVJs9jcn98_FZZtsiP2-8TqJQZ5Ak_2BrrWkKbVQ_141hhUt-RIk3GRuOEzMlN3SreOXuiHK2SjKStQJTrawKtbdDgpDjeESSVg_t_hRwDip17kZzuDslCsq9wZSGXznuh9IP28h-7CNukAjVqRqk3r9g98SNLvd7fwlXreE8QFKHebqr0Ok37-mcOUHE3rfZOAx4P92TduHr3WmpfLoFQtKd5UVj9gX5o2NaZjOFIu_nv9-SYZuaoUMETJ_HvV5pAZl5od53fvzcnJBcZ0ObAnftmTeO04k9PItEQ",
				Exponent:  "AQAB",
			},
			{
				KeyType:   "RSA",
				KeyID:     "978b94f3-c2c1-414b-af9e-504bc45e1604",
				Usage:     "sig",
				Algorithm: "RS256",
				Modulus:   "zq5e64y6WslNaaA_7cBLMe3AcEk5elEOOt0PTgAXgNSXVvqH154EjfS-aG5NvXW2_AqqS0YRf8g3PJYKhqtiOWC42V04r2AivjUjjZ7V065uxvykfnRCXWpNDu_A_lGqH7o1e6mQdGyW_T5s5jdyHRncb6CGtNqwTYng09O5-mBBS7KvIDqg8sLdqULHzJDXH9WhF-rZyPR9GLgolEg-9SybaXbWuVZFX9JcQ4gFC79YTRTfzkvZ-14yibqtlUFrUUEP1DfT-TWGVTU6P9wPr1XCePCU9_OGWEPiN2MzaNMydOUs6BxdcdjILHZhS51dHTiOWqy5aF7PWti7deAgXw",
				Exponent:  "AQAB",
			},
		}
	)

	jwks, err := service.Jwks(ctx)
	assert.Equal(t, expectedJwks, jwks)
	assert.Equal(t, nil, err)
}

func TestVerifyWithJwks(t *testing.T) {
	var (
		ctx    = context.Background()
		logger = logger.Setup(logger.EnvTesting)
		config = &model.AppConfig{
			Jwt: model.JwtConfig{
				AsymmetricKeys: model.JwtAsymmetricKeysConfig{
					&struct {
						Kid        string
						PrivateKey string
						PublicKey  string
					}{
						Kid: "8421bf44-1178-4414-8909-9e99a4dfb770",
						// PrivateKey: "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdDJWTkNiT3I4aTR6R29zT0xLSkVhbCtkcnc1SHFlTHppVEkwY0ZvVUR2Y2pQYVUwCmdWSnM5amNuOTgvRlpadHNpUDIrOFRxSlFaNUFrLzJCcnJXa0tiVlEvMTQxaGhVdCtSSWszR1J1T0V6TWxOM1MKcmVPWHVpSEsyU2pLU3RRSlRyYXdLdGJkRGdwRGplRVNTVmcvdC9oUndEaXAxN2taenVEc2xDc3E5d1pTR1h6bgp1aDlJUDI4aCs3Q051a0FqVnFScWszcjlnOThTTkx2ZDdmd2xYcmVFOFFGS0hlYnFyME9rMzcrbWNPVUhFM3JmClpPQXg0UDkyVGR1SHIzV21wZkxvRlF0S2Q1VVZqOWdYNW8yTmFaak9GSXUvbnY5K1NZWnVhb1VNRVRKL0h2VjUKcEFabDVvZDUzZnZ6Y25KQmNaME9iQW5mdG1UZU8wNGs5UEl0RVFJREFRQUJBb0lCQVFDc3VGRVhwQWw2YXB4aQprVGZtUFdTbHNpdDFwTU5GY3FMZVFWUTF4QUJFSCtrbXM2S0JjVG1Cb1d5WTdTc0JpS0Z0VzEwckgzQUpScHVYClJSZVBqUzV3d1h6cEpMYlA4cjU3WnVVa1U4bWlhR0g4aWZWVEk1ZlFDdWRhSWhweTRzTnBTSkVkcDRKRktORjYKbTlCM0Z3L2JtWmlVcWtqN0RDOE1NYlZkemxJR2xtaFpjckhHZyttb2NNc3hCcjY4TVgzMTZ5NjBoMWFDdTYwUwpvMHYyMzR3K0VBN1hxTUluUitoZUxRRVc3Y0c0ZUdZN2NNTnErY1NPWVRmczN1bGpVeHZLeUtocHJuN05wd1dhClJnVzhrcnVGSmJaRGJIUTJQZFhqakVXYURaZkVIU0h1bUovWmc2QTZZSlpmaGlFdDBkOWZxR2RPS3VUSDMxU3cKK0k5YnBKWkpBb0dCQU1DQnRiOWNDbExIbHllUDJhWjZnbTNaUzlobTNuM2lKOWcwS0R2U0c4NSt3OTlwdGVSTwpTLzdsT1UwRGN2KzVFMGN1UlpUdE9TR3NUY3JjWjhaM1RCZFVyeE5ZTWNxamgybEtjT0dWRlFaT3B6ZDhVN2U0CmdyaHdTb3hWdFovaHQvVnRpMi9sS3JZOVhueUZZQkRjZGdsbERaZjRIeHVmQ01hbmxhNUYzWGhiQW9HQkFQUGkKVG5raUtJQjBLS0JSb28rc0N3TUhJa0JFb0R4RC9yL0ZNcHdBdkh1TXZhSGxkc1krUVdBTUErUjloK2JPYUNYawoveVRvVkpvaGFKSVRiaU5HMDRpS2FMMXROaHNJTXJqL3U1dXM4TkVoOU5XSzZRbGZ2MDcvSklDWUUraTA4aGJGCmlVU1RMNCtFUW1sWGRrdmVtUHBSVkRzRjhBdlBFd3F0T1NuSEdJd0RBb0dBQVZHaUxpSnlTNmprWnpmOEZNRG8KSGRxTVEzcEk4ZkhYdGdwOWNCTjdiMG05QzgzTW1qalRHbmIxa29xQWdqSUJhTTV2V1pyYWRsbVkydGZ4dWhGZApLeGZBYjFCK1h0WUorbld4R2txTUwxUGduMmV4cHlPVGViSURRTHpobHF2VU45RTlVRkh3bmZrRHFiUzhPTUZaCjZheVFrRWI1NTVXS1dOb1RFM09WRmRzQ2dZQnJsRjlEUmRNUjNxdHhGTEdkcUtsdTIzMjdWY3BNNnoxN2dGUXoKeG90ZUFKWkJ6UU9ZclN1UFg1MXo4LyszeTBMYnZHamo4ZXduMVNiWWtPT2JnZ21iaUZwdGZMaEtNbEtWa3BGQwpPWVk4NmpxaTI5U3lBdDlUekc1Z256VGhDTGhsWFJ1UStWQVlnYUg5Nzh2SjZkWVhUVHJYa21YeC81VUp0NkdvCmtSOTkyd0tCZ0NUcGRJMU9DcnhBVEZ0YklqQXYvSHpFcjFmRGJ5ZnEzQkxWak5BYkV0bWRVamp0L2ZrU1VSSmcKQzZyRTJXTXUvQUZlQVNGSGVqb0tFQ2NnNWFJZ1U0ZTFUQm9paTVMTjNTdTJyTmFIeStxYmVhMWlaVU4zNTkybwp4amJnY1JBek16Q3JTYnkrTzhEUXFSdFJYM0xFczR5cjFlaUZyTnpNWkR0MXlvY05XeGZiCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t",
						PublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF0MlZOQ2JPcjhpNHpHb3NPTEtKRQphbCtkcnc1SHFlTHppVEkwY0ZvVUR2Y2pQYVUwZ1ZKczlqY245OC9GWlp0c2lQMis4VHFKUVo1QWsvMkJycldrCktiVlEvMTQxaGhVdCtSSWszR1J1T0V6TWxOM1NyZU9YdWlISzJTaktTdFFKVHJhd0t0YmREZ3BEamVFU1NWZy8KdC9oUndEaXAxN2taenVEc2xDc3E5d1pTR1h6bnVoOUlQMjhoKzdDTnVrQWpWcVJxazNyOWc5OFNOTHZkN2Z3bApYcmVFOFFGS0hlYnFyME9rMzcrbWNPVUhFM3JmWk9BeDRQOTJUZHVIcjNXbXBmTG9GUXRLZDVVVmo5Z1g1bzJOCmFaak9GSXUvbnY5K1NZWnVhb1VNRVRKL0h2VjVwQVpsNW9kNTNmdnpjbkpCY1owT2JBbmZ0bVRlTzA0azlQSXQKRVFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t",
					}, {
						Kid: "978b94f3-c2c1-414b-af9e-504bc45e1604",
						// PrivateKey: "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBenE1ZTY0eTZXc2xOYWFBLzdjQkxNZTNBY0VrNWVsRU9PdDBQVGdBWGdOU1hWdnFICjE1NEVqZlMrYUc1TnZYVzIvQXFxUzBZUmY4ZzNQSllLaHF0aU9XQzQyVjA0cjJBaXZqVWpqWjdWMDY1dXh2eWsKZm5SQ1hXcE5EdS9BL2xHcUg3bzFlNm1RZEd5Vy9UNXM1amR5SFJuY2I2Q0d0TnF3VFluZzA5TzUrbUJCUzdLdgpJRHFnOHNMZHFVTEh6SkRYSDlXaEYrclp5UFI5R0xnb2xFZys5U3liYVhiV3VWWkZYOUpjUTRnRkM3OVlUUlRmCnprdlorMTR5aWJxdGxVRnJVVUVQMURmVCtUV0dWVFU2UDl3UHIxWENlUENVOS9PR1dFUGlOMk16YU5NeWRPVXMKNkJ4ZGNkaklMSFpoUzUxZEhUaU9XcXk1YUY3UFd0aTdkZUFnWHdJREFRQUJBb0lCQUhYVWZYTUcyUnQzRm5ZNAprUm5IZmxjcHQ0T01pNE5MZ0xSWVlTaFQ3eEpZb1N0S1MzWEd0Y3dFa3lWUWRXdWxGN3hiakRpNzZyQVNBa085Ck9xVUtRa1o1K1FpYkYvMEw3dUxId3N3em1LNUZEUXpPN2l6VnRSd3l4Vm5Wb0E2ZG1rTGFVekY4TzBuVXVzUUgKK2VmS0Jubkd5NkNzUVFBTWlXUzdUWDBXZ1RuVzdqdWlhOEhWN1gxcGtMc3lPYjduaXNGdEkvWjNJN3pqYnlQRwpUeEhwTDZqdk1DYUVxSXFZeWZndmNYKytoKzNhR1pkL2taNzNSOHd4NWY3UTlxVEJVRTBBVFk3b2trdEt5Tm1mClZzVzAwTnhHUjg5bWNpeFNySjBvOVpsMjhwR3o0blRMdGtDdXdIcDBvWkREc0g4Vld4REtlMHdOeFY1SFVVZzgKeDNOOGJja0NnWUVBNWJFaTBSRFFZNkh0N3FzbStSV0l6UnlqaGJkL0lEZy9IdVVCanhoZVhSK01CdUhMaGpBVgpjNFRDamptTHE1TDhrNVlNbjFkOWpQTzlRcFQ2MXJRbGNpNkxrU1JHOE1EQWpFZTk2VWNwY00vaWN4MkdsaW1DCkg5Q3p0Ty95Qnh0bm9EL1Y5K2l2TXJybHIxNHVFSUYzbE9OYUdzOXhqWDlRU25BVHJoT0ZqOTBDZ1lFQTVscUgKUkE5SHRpbm1hVWlKOW03eFZielgzRTEzNHA3VjRzMHViOUwvTnVYQjJEaExLRTdlSER6b2MyckxEY1BDWWJYdgpVeE9VaXU1L3E4aWFxbTh0OVdEeUZtT25QTDdxQkRsK2FYV0oyeUJGL0VUTGVqVURrWGFpdEp4a2ZZWmNqdGFpCmp5OFgrb1Mwa1hLaDVsWWJ0WWlhMjlZMWRFNFowaGMrWW1Dd2kyc0NnWUJNQ3YzczJ6VXlsd3lQcElndGxLeUsKdzMxN3FvbGk0Rnc5WFRITDd4Um1uaWdjcXlwWFRabjhlYXB6cmFlSThRdS96TUIzREY4YmlDSlRaY0U1emNCTAo4Zzd3eVdMWEYrbG5SK1VlMHhsc0tOYmVwNXJFSWcvYmVwdlVQbEFSZkVndGJKVHBFMWJWWTd6Zzl6d202TVh2Ck8rbTcwSXZXZlp6V1dBNmI1Z2lrM1FLQmdRQ2dlbWNMOGowNldqeHNFcDRTc2IydHhtNzN5bngvdzdvc1ZGZEsKamt0QW1wWU9KVUd1NWlyMHJ6WWQya04xNW5CUUZDaVR6dXRJVXRsYyt0ZE5iMEVzVm9xeE1HRzgveUFYSzFxMgo2NnpMbk5ORTZVUjhjZHgzaFpmYWphMndpVitJSDEyR3RQOVdod2NvNFJ5TFZHZmZjemkzZUtjcXpKSk9lK1Q1Ck1uVXIyd0tCZ0N5d3VvNlRreFZ4S3k4NHZ4OFN1RXVHL0xWWDI1NllaUGpuc0VZMTJmZDF0NXpaNU5hQlZSaGQKcitDdkk4UmZyNzdQZ1JMd2VFQzVKMlBiUUF1dTVOc1pzS3E0ZWoxNFNkZ1JoYnUzcDNmZ05HQTF6THljcDB1aQpHdWt4dnp6Q1Q0TmFZQ0szWGtmcndZU0hlNnUrbDF2RWZzV3Q0TWxsckh6RDJ0NzF4bGllCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t",
						PublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF6cTVlNjR5NldzbE5hYUEvN2NCTApNZTNBY0VrNWVsRU9PdDBQVGdBWGdOU1hWdnFIMTU0RWpmUythRzVOdlhXMi9BcXFTMFlSZjhnM1BKWUtocXRpCk9XQzQyVjA0cjJBaXZqVWpqWjdWMDY1dXh2eWtmblJDWFdwTkR1L0EvbEdxSDdvMWU2bVFkR3lXL1Q1czVqZHkKSFJuY2I2Q0d0TnF3VFluZzA5TzUrbUJCUzdLdklEcWc4c0xkcVVMSHpKRFhIOVdoRityWnlQUjlHTGdvbEVnKwo5U3liYVhiV3VWWkZYOUpjUTRnRkM3OVlUUlRmemt2WisxNHlpYnF0bFVGclVVRVAxRGZUK1RXR1ZUVTZQOXdQCnIxWENlUENVOS9PR1dFUGlOMk16YU5NeWRPVXM2QnhkY2RqSUxIWmhTNTFkSFRpT1dxeTVhRjdQV3RpN2RlQWcKWHdJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t",
					},
				},
			},
		}
		service        = New(nil, nil, nil, config, logger)
		expectedClaims = &client.Claims{
			Claims: model.Claims{
				UserId:      28,
				Name:        "backend",
				Email:       "it@example.com",
				PhoneNumber: "+6281222021792",
				TokenType:   constant.TokenTypeAccess,
				Extended:    false,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "kopdig.k-ertas.com",
					Subject:   "it@example.com",
					ExpiresAt: jwt.NewNumericDate(time.Date(2026, time.April, 15, 13, 43, 53, 0, time.Local)),
				},
			},
		}
	)

	util.SetLogger(logger)

	jwks, err := service.Jwks(ctx)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	tokenString := "eyJhbGciOiJSUzI1NiIsImtpZCI6Ijg0MjFiZjQ0LTExNzgtNDQxNC04OTA5LTllOTlhNGRmYjc3MCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyOCwibmFtZSI6ImJhY2tlbmQiLCJlbWFpbCI6Iml0QGV4YW1wbGUuY29tIiwicGhvbmVfbnVtYmVyIjoiKzYyODEyMjIwMjE3OTIiLCJyb2xlX2lkIjo1LCJncm91cHMiOlsiTktELUJERyJdLCJ0b2tlbl90eXBlIjowLCJleHRlbmRlZCI6ZmFsc2UsImlzcyI6ImtvcGRpZy5rLWVydGFzLmNvbSIsInN1YiI6Iml0QGV4YW1wbGUuY29tIiwiZXhwIjoxNzc2MjM1NDMzfQ.DQFfpqdXAAS-BV9OXgXxcHVb8SbYsfky974O4p3XL_DCHLDqo3k73GfenOjC_AwLmj5gbq8UAhDkgKqRBby1sqUa9ZVirVF8Dke7y3R_bv3tA4oVQfbcNBUfNaDnatbPcj7Xc1h5nLS1MVQhdHinBmt0W-WDbXwQZR-syUO0mGifIRqMgEiNAuuyAVMFu9GDqawnS-WngbGLV1IFPDaQqafdiFatfevOd35EinQsIIk4rpPrxongnVA4QSBYwlxqVruei9WmW1cBJqUDLIvp6Trv0lO7n0AV6nBWKdLnz4cr8Ne9bc7vcPU66Gt7cVjjPN6O4gj5iG9hgbCgD7glzQ"
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, &client.Claims{})
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	for _, v := range jwks {
		if v.KeyID == token.Header["kid"] {
			nBytes, err := base64.RawURLEncoding.DecodeString(jwks[0].Modulus)
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			eBytes, err := base64.RawURLEncoding.DecodeString(jwks[0].Exponent)
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			publicKey := &rsa.PublicKey{
				N: new(big.Int).SetBytes(nBytes),
				E: int(new(big.Int).SetBytes(eBytes).Int64()),
			}

			token, err := jwt.ParseWithClaims(tokenString, &client.Claims{}, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return publicKey, nil
			})
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			if claims, ok := token.Claims.(*client.Claims); ok && token.Valid {
				assert.Equal(t, expectedClaims, claims)
				return
			}
		}
	}

	assert.FailNow(t, "no matching jwks")
}

func TestGenerateRSA(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.Equal(t, nil, err)

	privateKeyPEM := exportRSAPrivateKeyToPEM(privateKey)
	publicKeyPEM, err := exportRSAPublicKeyToPEM(&privateKey.PublicKey)
	assert.Equal(t, nil, err)

	fmt.Println(privateKeyPEM)
	fmt.Println(publicKeyPEM)

	encodedPrivateKeyPEM := base64.StdEncoding.EncodeToString([]byte(privateKeyPEM))
	encodedPublicKeyPEM := base64.StdEncoding.EncodeToString([]byte(publicKeyPEM))

	fmt.Println(encodedPrivateKeyPEM)
	fmt.Println()
	fmt.Println(encodedPublicKeyPEM)
}

func TestJwkPlayground(t *testing.T) { //nolint
	var (
		ctx    = context.Background()
		logger = logger.Setup(logger.EnvTesting)
		s      = service{
			config: &model.AppConfig{
				Jwt: model.JwtConfig{
					AsymmetricKeys: model.JwtAsymmetricKeysConfig{
						&struct {
							Kid        string
							PrivateKey string
							PublicKey  string
						}{
							Kid:        "8421bf44-1178-4414-8909-9e99a4dfb770",
							PrivateKey: "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUJmaFlXTlFUNHNCNWgvdVlRRDc1L3l5Rlc2elRqY0FiMWlrL0VWeVRUanFpTnNFYVNiCnpJRzBNWWRaSjVzQ0lWZ2MyYk03RjRNc3NNdHg0N1ZVRE0xeFh3REdmaERoWDFyaUN4YmQyTU1QYWpwVmdWanUKMDlUTGs3QWNaRkhNOTBiOWtDbjhsMkVSbncyWnpoYkdrUVloYlRmUUNzdDlPVHNTaGh6aUpubk9iRlVaUWF0bgpkTVdsOHVaR0NRNVl2U0N1ZDBVSC83ZUN4aW96WTV3NVR0SmcvT1M4OXpyY25Td2tpU1gzdlZxS2ZpaDNPRHh6CnV3STJXVzVmQnlUWEVHcmZvcXpXWUEwcDVFUkFSMmtaRmNyY1NZZHJzUGhWQnE1VGoyUnBkbno5TnUrSzNqa1kKZkppWG5ZL1NOQXQ1WXZZZWFXRkZxOVZxZ210cis5WE1sWnQvQWdNQkFBRUNnZ0VBSDlNeG90Vm83R3gvYjVhVQptR2NlNkZtOHptY1BtYVZ1dnNoQm01dVU1b0ZLR2ZocTJvbXhjU0xMVUhYMG5Db1YzRTdmKzBFak1DR2JOcy9DCkcwWEVzUkFSQnhEN2VNczNVWGFXWU1XV2Y2MUowREV2T3lzU2k5MGg0T08vcVVWOXZuOW9yY0tWMGJRbmFPWVUKQ29aSS81d082MTZkVzVSVXpTQW53V1ZHVWZLM2YzWjRuOHh6S2d2bitCUmJkUzdDNDNReC9GMFA1QzdabWVlVwptQWZuNUtyVUUyMG94WjJKeGU2Q1l4RHZjclk1QTh5aGp2NkJzV2FxWVlOcW50YXN0YXExTVpPY2tQYktsSFdQCllBcEdXdFUyTkN4VWRFK2VtRmc3K0RhZ0dNR0pjb0xHeDIydysxU2dQTGlvMHpQT2NjWjhDZkNWdHl4aXJtcncKSDJqVUNRS0JnUUNwcGRMLzBrNExGVHRWQThjNHlzalljR2szMmgvSXNMTUNHOXZLa0lXK29NSm5nNU9HK2tiLwpSTDcrVlRXV2lSdi93YXJMV2p2VmhsUDBNV2ZqZjJJVXdmYTNGeWwwUDB3K0pyT2U0M1c4RzZhTmlUU2I4VFhSCnNhdWhGLzJ1MFBmbXlucWt6SGJHMzFSMmNUZVVrM2hSTlRaa2g4VHBwUlVTaEZyT0wzUHp5d0tCZ1FDUUpKWTMKSER6ZFczRG5wOE1sbWdjV1RsekdaS0JyWFd2T2VXb2FrTUdyRk5XUnhoMEp1MmlIN21nVHpXYjlmbWRKSm1vcgp3VGxLS2Q0b0VpU0pGSnBVREdXZHBiRVhVaXgrY21ZZ0xUd1hkbDVJZE5oN3FhNFp0R1NEOTNVZXk5MTJVNE81ClVlUWVhSUdkaFBJaEFlNGtaT1F0d1pCK1Z1YXBCK3dWRkxoSW5RS0JnUUNiV3pBTzloaGlMZDlYeTAzMXhENkoKZHVma0xleE5iUU9CT3VIY2J0MEw1VXdpWDJ3S2Y4ZmtuS0FMYVJ6WjdsV2xzVVVuVkVyWEQxeHlrNHYvMmZlSAo2dGgwY3RHVGt5UFBCc0lYRDFZU0hZQTR2UjFnY1ZSSDQ5eTRlYS9uRjVidDB4N2RMQ0Rabmt0SzdBTnFIR0ppCmU4aUQ1NUY4SmFGV2c3NWtjekJNWVFLQmdEZyszcExRcFB0blpBNHhDMWdQMjNZYnk5M3FoQ0tCQ01FLzVXUksKV2hnTkFDMXExZ2ZuSmluc29KWWhqMitaTkdwNTMvSUU2dnNDalZxcmdiQXY1dXluRGJ2UFhPUVJ2NlR6dE9BWApacHh0SnVzMUZRaGtOTGg1Q01QcCtyeXlwazgyMVc2cUFzN096czBOaElIV3cvdFZseWczb005N3ozUGowSDZGCllFZU5Bb0dBRDFVNVJzTmEyMElvdGE3UWxJNkdaZ2lTTDJCdjhzU2hST042OFp4MGhYRm5IVlAva0FqUTNZUDgKeERzUFVQL1p3dkVweFBNeG9rNkp1VStzaTVVVjQ1QTRydW5MUCtSd3dVQXlKbzg3WTZhZ3htMGkxS3hSUzk2RgpsZVlHZEJNdGV1bG5GNWozVC9DSmVMTWZLZTN1RWxzWjF4TTRkQks2cDVFVEZpYXF5QXM9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t",
							PublicKey:  "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklUQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FRNEFNSUlCQ1FLQ0FRQmZoWVdOUVQ0c0I1aC91WVFENzUveQp5Rlc2elRqY0FiMWlrL0VWeVRUanFpTnNFYVNieklHME1ZZFpKNXNDSVZnYzJiTTdGNE1zc010eDQ3VlVETTF4Clh3REdmaERoWDFyaUN4YmQyTU1QYWpwVmdWanUwOVRMazdBY1pGSE05MGI5a0NuOGwyRVJudzJaemhiR2tRWWgKYlRmUUNzdDlPVHNTaGh6aUpubk9iRlVaUWF0bmRNV2w4dVpHQ1E1WXZTQ3VkMFVILzdlQ3hpb3pZNXc1VHRKZwovT1M4OXpyY25Td2tpU1gzdlZxS2ZpaDNPRHh6dXdJMldXNWZCeVRYRUdyZm9xeldZQTBwNUVSQVIya1pGY3JjClNZZHJzUGhWQnE1VGoyUnBkbno5TnUrSzNqa1lmSmlYblkvU05BdDVZdlllYVdGRnE5VnFnbXRyKzlYTWxadC8KQWdNQkFBRT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t",
						},
					},
				},
			},
		}
		jwks = []*model.Jwk{}
	)

	util.SetLogger(logger)

	for _, v := range s.config.Jwt.AsymmetricKeys {
		b, err := base64.StdEncoding.DecodeString(v.PublicKey)
		if err != nil {
			util.LogContext(ctx).Error(err.Error())
			return
		}

		pub, err := jwt.ParseRSAPublicKeyFromPEM(b)
		if err != nil {
			util.LogContext(ctx).Error(err.Error())
			return
		}

		// n := pub.N.Bytes()
		// e := make([]byte, 4)
		// e[3] = byte(pub.E & 0xff)
		// e[2] = byte((pub.E >> 8) & 0xff)
		// e[1] = byte((pub.E >> 16) & 0xff)
		// e[0] = byte((pub.E >> 24) & 0xff)

		// jwks = append(jwks, &model.Jwk{
		// 	KeyType:   "RSA",
		// 	KeyID:     v.Kid,
		// 	Usage:     "sig",
		// 	Algorithm: "RS256",
		// 	Modulus:   base64.RawURLEncoding.EncodeToString(n),
		// 	Exponent:  base64.RawURLEncoding.EncodeToString(e[3:]),
		// })

		n := pub.N.Bytes()
		e := make([]byte, 4)
		e[3] = byte(pub.E & 0xff)
		e[2] = byte((pub.E >> 8) & 0xff)
		e[1] = byte((pub.E >> 16) & 0xff)
		e[0] = byte((pub.E >> 24) & 0xff)

		jwks = append(jwks, &model.Jwk{
			KeyType:   "RSA",
			KeyID:     v.Kid,
			Usage:     "sig",
			Algorithm: "RS256",
			Modulus:   base64.RawURLEncoding.EncodeToString(pub.N.Bytes()),
			Exponent:  base64.RawURLEncoding.EncodeToString(big.NewInt(int64(pub.E)).Bytes()),
		})

		util.LogContext(ctx).Info(string(n))
		util.LogContext(ctx).Info(string(e[3:]))
	}

	b, err := json.Marshal(jwks[0])
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Parse jwk
	jwkss, err := jwk.ParseKey(b)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	util.LogContext(ctx).Info("jwks", zap.Any("k", jwkss))

	// Idk
	pub, err := jwk.PublicKeyOf(jwkss)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	util.LogContext(ctx).Info("pub", zap.Any("k", pub))

	// public key
	bPublicKey, err := base64.StdEncoding.DecodeString(s.config.Jwt.AsymmetricKeys[0].PublicKey)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(bPublicKey)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	util.LogContext(ctx).Info("pub", zap.Any("k", publicKey))

	bn, err := base64.RawURLEncoding.DecodeString(jwks[0].Modulus)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	be, err := base64.RawURLEncoding.DecodeString(jwks[0].Exponent)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	publicKey2 := &rsa.PublicKey{
		N: new(big.Int).SetBytes(bn),
		E: int(new(big.Int).SetBytes(be).Int64()),
	}
	util.LogContext(ctx).Info("pub", zap.Any("k", publicKey2))

	// Test parse
	pk, err := exportRSAPublicKeyToPEM(publicKey2)
	util.LogContext(ctx).Info(pk)

	// Verify
	strToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6Ijg0MjFiZjQ0LTExNzgtNDQxNC04OTA5LTllOTlhNGRmYjc3MCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyOCwibmFtZSI6ImJhY2tlbmQiLCJlbWFpbCI6Iml0QGV4YW1wbGUuY29tIiwicGhvbmVfbnVtYmVyIjoiKzYyODEyMjIwMjE3OTIiLCJyb2xlX2lkIjo1LCJncm91cHMiOlsiTktELUJERyJdLCJ0b2tlbl90eXBlIjowLCJleHRlbmRlZCI6ZmFsc2UsImlzcyI6ImtvcGRpZy5rLWVydGFzLmNvbSIsInN1YiI6Iml0QGV4YW1wbGUuY29tIiwiZXhwIjoxNzc2MTc1ODg3fQ.MA2_PoIFkgT75qF3dyJHzAGLDUZhskXOBkDdwwYO0wWAUT7IQ9zDIFKWphmt8Bp0tSHbuP-LulMeRcfYKQdRLxwcOuDakYnK18g8w7DFUhcbKmC_xBdQ_QZ-HrDTErolOYpW7hO5bf-jK3VtOZ3QTcvGjRhh3fxPEsG1aoTHRm6EI2HBzQBCmS5hE9h-Cd4J2FNAZwAb5O8Ruxg5jsAQCKy6thT7OyEkyBiCeAA64-UsrkYlTlUY8og0JZI1RB7PjpvBKZgZE8aM7eo_ZrCjc7cvn2dP0zgg4nYvqUM3qlCZ5QS-Bt5-FfqibKlJbKZTuNsF0qKIszwbkSPy2Qu9Mg"
	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return publicKey2, nil
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	util.LogContext(ctx).Info("r", zap.Any("k", token))

	return
}

// Export RSA Public Key to PEM Format.
func exportRSAPublicKeyToPEM(pubKey *rsa.PublicKey) (string, error) {
	pubKeyDER, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", fmt.Errorf("error marshaling public key: %v", err)
	}
	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyDER,
	}
	pubKeyPEM := pem.EncodeToMemory(pubKeyBlock)
	return string(pubKeyPEM), nil
}

// Export RSA Private Key to PEM Format.
func exportRSAPrivateKeyToPEM(privateKey *rsa.PrivateKey) string {
	privateKeyDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDER,
	}
	privateKeyPEM := pem.EncodeToMemory(privateKeyBlock)
	return string(privateKeyPEM)
}
