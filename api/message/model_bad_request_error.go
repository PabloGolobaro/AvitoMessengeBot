/*
 * Мессенджер
 *
 * API Мессенджера - набор методов для получения списка чатов пользователя на Авито, получения сообщений в чате, отправки сообщения в чат и другие Через API Мессенджера можно организовать интеграцию между мессенджером Авито и сторонней системой в обе стороны  **Авито API для бизнеса предоставляется согласно [Условиям использования](https://api.avito.ru/docs/public/APITermsOfServiceV1.pdf).**
 *
 * API version: 1
 * Contact: supportautoload@avito.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package message

type BadRequestError struct {
	Error_ *BadRequestErrorError `json:"error,omitempty"`
}
