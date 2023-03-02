package presenter

import "github.com/gofiber/fiber/v2"

const (
	NoHourForThisDay              = "ce magasin n'a pas d'horaires d'ouverture ce jour"
	ReservationOutOfOpeningHours  = "vous ne pouvez pas réserver en dehors des horaires d'ouverture du magasin"
	ReservationCanceled           = "réservation annulée"
	NotAllowedToCancelReservation = "vous n'êtes pas autoriser à annuler cette réservation"
	WrongCredential               = "email ou mot de passe incorrect"
	RessourceNotAuthorized        = "vous n'êtes pas autorisé à accéder à cette ressource"
	RessourceNotFound             = "ressource non trouvée"
	ActionNotAllowed              = "vous n'êtes pas autorisé à effectuer cette action"
	CannotAddShop                 = "vous ne pouvez pas ajouter ce magasin"
	CannotRemoveService           = "vous ne pouvez pas supprimer ce service"
	PasswordsNotMatch             = "les mots de passe ne correspondent pas"
)

func Error(
	c *fiber.Ctx,
	code int,
	err error,
) error {
	return c.Status(code).JSON(&fiber.Map{
		"statusCode": code,
		"message":    err.Error(),
	})
}

func Response(
	c *fiber.Ctx,
	code int,
	response interface{},
) error {
	return c.Status(code).JSON(response)
}
