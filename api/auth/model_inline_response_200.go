/*
 * Авторизация
 *
 * Получение и обновление авторизационных токенов для персональной авторизации и авторизации приложений **Авито API для бизнеса предоставляется согласно [Условиям использования](https://api.avito.ru/docs/public/APITermsOfServiceV1.pdf).**
 *
 * API version: 1
 * Contact: supportautoload@avito.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package auth

type InlineResponse200 struct {
	// Ключ для временной авторизации в системе
	AccessToken string `json:"access_token,omitempty"`
	// Время жизни ключа в секундах
	ExpiresIn float64 `json:"expires_in,omitempty"`
	// Тип ключа авторизации
	TokenType string `json:"token_type,omitempty"`
}
