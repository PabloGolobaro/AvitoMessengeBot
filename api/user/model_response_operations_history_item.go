/*
 * Информация о пользователе
 *
 * API для получения баланса кошелька пользователя, истории операций и инфорации об авторизованном пользователе **Авито API для бизнеса предоставляется согласно [Условиям использования](https://api.avito.ru/docs/public/APITermsOfServiceV1.pdf).**
 *
 * API version: 1
 * Contact: supportautoload@avito.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package user

type ResponseOperationsHistoryItem struct {
	// Сумма бонусов
	AmountBonus float64 `json:"amountBonus,omitempty"`
	// Сумма реальных денег
	AmountRub float64 `json:"amountRub,omitempty"`
	// Сумма операции всего
	AmountTotal float64 `json:"amountTotal,omitempty"`
	// ID объявления
	ItemId int32 `json:"itemId,omitempty"`
	// Название операции
	OperationName string `json:"operationName,omitempty"`
	// Тип операции
	OperationType string `json:"operationType,omitempty"`
	// Код типа услуги
	ServiceId int32 `json:"serviceId,omitempty"`
	// Название конкретного типа услуги
	ServiceName string `json:"serviceName,omitempty"`
	// Тип услуги
	ServiceType string `json:"serviceType,omitempty"`
	// Дата/время совершения операции
	UpdatedAt string `json:"updatedAt,omitempty"`
}