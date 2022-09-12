package app

/*
	接口访问控制
	常见的两种 API 访问控制方案, 分别是 OAuth2.0 和 JWT
	- OAuth 2.0: 本质上是一个授权的行业标准协议, 提供了一整套的授权机制的
		指导标准, 常用于使用第三方登陆的情况, 提供其它第三方站点(例如用微信,
		QQ, Github 账号)关联登陆, QAuth 常常还会授予第三方应用去获取到对应
		账号的个人基本信息等等;
	- JWT: 常用于前后端分离的情况, 能够非常便捷的给 API 接口提供安全鉴权;

	TODO: jwt 包原理, 源码实现
	JWT: Json Web Token
	定义了一种紧凑且自包含的方式, 用于在各方之间作为 JSON 对象安全地传输信息;
	由于此信息是经过数字签名的, 因此可以被验证和信任, 可以使用使用 RSA 或
	ECDSA 的公用/专用密钥对对 JWT 进行签名;
	JWT 由三部分组成, 其中用"."分隔:
	- Header:头部, 用于描述元数据的对象
	- Payload:有效载荷, 存储在 JWT 中实际传输的数据
	- Signature:签名, 对前面两个部分组合（Header+Payload）进行约定算法和规则
		的签名, 而签名将会用于校验消息在整个过程中有没有被篡改, 并且对有使用
		私钥进行签名的令牌, 它还可以验证 JWT 的发送者是否它的真实身份;
	以上三部分都是 base64UrlEncode(TODO)算法转换相应的json信息后生成的;

	TODO: Payload 中不应该明文存储重要的信息, 因为可以根据token中的Payload
	信息使用base64反向解密获取到Payload中的值

	JWT的使用场景:
	通常会先在内部约定好 JWT 令牌的交流方式, 存储在 Header、Query Param、
	Cookie、Session 都有, 但最常见的是存储在 Header 中; 然后服务端提供一个
	获取 JWT 令牌的接口方法, 返回而客户端去使用, 在客户端请求其余的接口时
	需要带上所签发的 JWT 令牌, 然后服务端接口也会到约定位置上获取 JWT 令
	牌来进行鉴权处理, 以此流程来鉴定是否合法;

	TODO: jwt 包原理, 源码实现
	go的JWT实现, github.com/dgrijalva/jwt-go

*/

import (
	"blog/global"
	"blog/pkg/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppKey             string `json:"app_key"`
	AppSecret          string `json:"app_secret"`
	jwt.StandardClaims        // TODO: 熟悉 StandardClaims
	// 过期时间(ExpiresAt)是存储在 Payload 中的, 也就是 JWT 令牌一旦签发,
	// 在没有做特殊逻辑的情况下, 过期时间是不可以再度变更的,
	// 因此务必根据实际的项目情况进行设计和思考;
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// TODO: jwt 常用函数的使用
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// 解析和校验 Token
func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明, 方法内部主要是具体的解码和校验的过程
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return GetJWTSecret(), nil
		})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		// Valid 验证基于时间的声明
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
