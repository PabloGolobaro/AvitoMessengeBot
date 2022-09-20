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

type RefreshRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	// Тип OAuth flow. Строка refresh_token
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}
